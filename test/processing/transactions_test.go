// Package processing holds per-resource CRUD and action coverage for the
// payment-processing surface. It is one CI shard; see TESTING.md.
package processing

import (
	"context"
	"testing"

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
