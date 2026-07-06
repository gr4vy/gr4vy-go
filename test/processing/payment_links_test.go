package processing

import (
	"context"
	"testing"

	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/models/operations"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

func TestPaymentLinkCRUD(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	link, err := m.Client.PaymentLinks.Create(ctx, components.PaymentLinkCreate{
		Amount:             1299,
		Currency:           "USD",
		Country:            "US",
		ExternalIdentifier: strPtr(harness.UniqueID("link")),
	}, nil)
	if err != nil {
		t.Fatalf("create payment link: %v", err)
	}
	if link.ID == "" {
		t.Fatal("payment link id empty")
	}

	got, err := m.Client.PaymentLinks.Get(ctx, link.ID, nil)
	if err != nil {
		t.Fatalf("get payment link: %v", err)
	}
	if got.ID != link.ID {
		t.Fatalf("expected %s, got %s", link.ID, got.ID)
	}

	if _, err := m.Client.PaymentLinks.List(ctx, operations.ListPaymentLinksRequest{}); err != nil {
		t.Fatalf("list payment links: %v", err)
	}

	if err := m.Client.PaymentLinks.Expire(ctx, link.ID, nil); err != nil {
		t.Fatalf("expire payment link: %v", err)
	}
}

func strPtr(s string) *string { return &s }
