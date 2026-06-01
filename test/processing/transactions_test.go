// Package processing holds per-resource CRUD and action coverage for the
// payment-processing surface. It is one CI shard; see TESTING.md.
package processing

import (
	"context"
	"testing"

	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/models/operations"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

func TestTransactionsList(t *testing.T) {
	m := harness.Merchant(t)
	page, err := m.Client.Transactions.List(context.Background(), operations.ListTransactionsRequest{})
	if err != nil {
		t.Fatalf("list transactions: %v", err)
	}
	if page == nil {
		t.Fatal("list transactions returned nil")
	}
}

func TestTransactionsGetMissingIsReached(t *testing.T) {
	m := harness.Merchant(t)
	harness.Reaches(t, "transactions.get(missing)", func() error {
		_, err := m.Client.Transactions.Get(context.Background(), harness.MissingID, nil)
		return err
	})
}

func TestTransactionGetAfterAuthorize(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	tx := harness.Authorize(t, m, 1299, "USD")
	got, err := m.Client.Transactions.Get(ctx, tx.ID, nil)
	if err != nil {
		t.Fatalf("get transaction: %v", err)
	}
	if got.ID != tx.ID {
		t.Fatalf("expected transaction %s, got %s", tx.ID, got.ID)
	}
}

// TestTransactionSubResources exercises the per-transaction read and action
// endpoints against a freshly authorized transaction.
func TestTransactionSubResources(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()
	tx := harness.Authorize(t, m, 1299, "USD")

	updated, err := m.Client.Transactions.Update(ctx, tx.ID, components.TransactionUpdate{}, nil)
	if err != nil {
		t.Fatalf("update transaction: %v", err)
	}
	if updated.ID != tx.ID {
		t.Fatalf("expected %s, got %s", tx.ID, updated.ID)
	}

	if _, err := m.Client.Transactions.Actions.List(ctx, tx.ID, nil); err != nil {
		t.Fatalf("list transaction actions: %v", err)
	}
	if _, err := m.Client.Transactions.Events.List(ctx, tx.ID, nil, nil, nil); err != nil {
		t.Fatalf("list transaction events: %v", err)
	}
	if _, err := m.Client.Transactions.Settlements.List(ctx, tx.ID, nil); err != nil {
		t.Fatalf("list transaction settlements: %v", err)
	}
	harness.Reaches(t, "transactions.settlements.get", func() error {
		_, err := m.Client.Transactions.Settlements.Get(ctx, tx.ID, harness.MissingID, nil)
		return err
	})
}

// TestTransactionCancelIsReached cancels a fresh authorization. Depending on
// the connector this may succeed or be cleanly rejected; either way it reaches.
func TestTransactionCancelIsReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()
	tx := harness.Authorize(t, m, 1299, "USD")

	harness.Reaches(t, "transactions.cancel", func() error {
		_, err := m.Client.Transactions.Cancel(ctx, tx.ID, nil)
		return err
	})
}
