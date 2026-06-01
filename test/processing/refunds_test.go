package processing

import (
	"context"
	"testing"

	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/models/operations"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

func TestRefundsGetMissingIsReached(t *testing.T) {
	m := harness.Merchant(t)
	harness.Reaches(t, "refunds.get(missing)", func() error {
		_, err := m.Client.Refunds.Get(context.Background(), harness.MissingID, nil)
		return err
	})
}

func TestTransactionRefundsList(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	tx := harness.Authorize(t, m, 1299, "USD")
	refunds, err := m.Client.Transactions.Refunds.List(ctx, tx.ID, nil)
	if err != nil {
		t.Fatalf("list transaction refunds: %v", err)
	}
	if refunds == nil {
		t.Fatal("list transaction refunds returned nil")
	}
}

func TestCaptureThenRefund(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	tx := harness.Authorize(t, m, 1299, "USD")
	if _, err := m.Client.Transactions.Capture(ctx, operations.CaptureTransactionRequest{
		TransactionID:            tx.ID,
		TransactionCaptureCreate: components.TransactionCaptureCreate{Amount: ptrInt(1299)},
	}); err != nil {
		t.Fatalf("capture: %v", err)
	}

	refund, err := m.Client.Transactions.Refunds.Create(ctx, tx.ID, components.TransactionRefundCreate{
		Amount: ptrInt(500),
	}, nil, nil)
	if err != nil {
		t.Fatalf("partial refund: %v", err)
	}
	if refund.ID == "" {
		t.Fatal("refund id empty")
	}

	got, err := m.Client.Transactions.Refunds.Get(ctx, tx.ID, refund.ID, nil)
	if err != nil {
		t.Fatalf("get refund: %v", err)
	}
	if got.ID != refund.ID {
		t.Fatalf("expected refund %s, got %s", refund.ID, got.ID)
	}
}

// TestRefundAllIsReached exercises POST /transactions/{id}/refunds/all on a
// freshly authorized (uncaptured) transaction. The connector may accept or
// reject it; either way the endpoint is reached.
func TestRefundAllIsReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()
	tx := harness.Authorize(t, m, 1299, "USD")

	harness.Reaches(t, "transactions.refunds.all", func() error {
		_, err := m.Client.Transactions.Refunds.All.Create(ctx, tx.ID, nil, nil, nil)
		return err
	})
}

func ptrInt(i int64) *int64 { return &i }
