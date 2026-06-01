// Package flows holds the headline end-to-end narratives: a full transaction
// lifecycle and a full buyer lifecycle. It is one CI shard; see TESTING.md.
package flows

import (
	"context"
	"testing"

	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/models/operations"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

func TestAuthorizeCaptureRefund(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	tx := harness.Authorize(t, m, 1299, "USD")
	if tx.Status != components.TransactionStatusAuthorizationSucceeded {
		t.Fatalf("expected authorization_succeeded, got %s", tx.Status)
	}

	captured, err := m.Client.Transactions.Capture(ctx, operations.CaptureTransactionRequest{
		TransactionID:            tx.ID,
		TransactionCaptureCreate: components.TransactionCaptureCreate{Amount: gr4vyInt(1299)},
	})
	if err != nil {
		t.Fatalf("capture: %v", err)
	}
	if captured == nil {
		t.Fatal("capture returned nil response")
	}

	refund, err := m.Client.Transactions.Refunds.Create(ctx, tx.ID, components.TransactionRefundCreate{
		Amount: gr4vyInt(1299),
	}, nil, nil)
	if err != nil {
		t.Fatalf("refund: %v", err)
	}
	if refund.ID == "" {
		t.Fatal("refund id is empty")
	}
}

func TestAuthorizeThenVoid(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	tx := harness.Authorize(t, m, 1599, "USD")
	if tx.Status != components.TransactionStatusAuthorizationSucceeded {
		t.Fatalf("expected authorization_succeeded, got %s", tx.Status)
	}

	if _, err := m.Client.Transactions.Void(ctx, tx.ID, nil, nil, nil); err != nil {
		t.Fatalf("void: %v", err)
	}
}

func TestSyncIsReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	tx := harness.Authorize(t, m, 1299, "USD")
	harness.Reaches(t, "transactions.sync", func() error {
		_, err := m.Client.Transactions.Sync(ctx, tx.ID, nil)
		return err
	})
}

func gr4vyInt(i int64) *int64 { return &i }
