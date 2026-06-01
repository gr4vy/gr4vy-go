package processing

import (
	"context"
	"testing"

	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/models/operations"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

// TestBuyerShippingDetailsCRUD covers the full shipping-details surface for a
// buyer: create, list, get, update, delete.
func TestBuyerShippingDetailsCRUD(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	buyer, err := m.Client.Buyers.Create(ctx, components.BuyerCreate{
		ExternalIdentifier: gr4vygo.String(harness.UniqueID("ship-buyer")),
	}, nil)
	if err != nil {
		t.Fatalf("create buyer: %v", err)
	}
	buyerID := *buyer.ID

	shipping, err := m.Client.Buyers.ShippingDetails.Create(ctx, buyerID, components.ShippingDetailsCreate{
		FirstName: gr4vygo.String("Jane"),
		LastName:  gr4vygo.String("Doe"),
		Address:   harness.SampleAddress(),
	}, nil)
	if err != nil {
		t.Fatalf("create shipping details: %v", err)
	}
	shippingID := *shipping.ID

	if _, err := m.Client.Buyers.ShippingDetails.List(ctx, buyerID, nil); err != nil {
		t.Fatalf("list shipping details: %v", err)
	}
	if _, err := m.Client.Buyers.ShippingDetails.Get(ctx, buyerID, shippingID, nil); err != nil {
		t.Fatalf("get shipping details: %v", err)
	}
	if _, err := m.Client.Buyers.ShippingDetails.Update(ctx, buyerID, shippingID, components.ShippingDetailsUpdate{
		LastName: gr4vygo.String("Smith"),
	}, nil); err != nil {
		t.Fatalf("update shipping details: %v", err)
	}
	if err := m.Client.Buyers.ShippingDetails.Delete(ctx, buyerID, shippingID, nil); err != nil {
		t.Fatalf("delete shipping details: %v", err)
	}
}

func TestBuyerGiftCardsList(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	buyer, err := m.Client.Buyers.Create(ctx, components.BuyerCreate{
		ExternalIdentifier: gr4vygo.String(harness.UniqueID("gc-buyer")),
	}, nil)
	if err != nil {
		t.Fatalf("create buyer: %v", err)
	}

	// GET /buyers/gift-cards — a 2xx even when the buyer has none.
	if _, err := m.Client.Buyers.GiftCards.List(ctx, operations.ListBuyerGiftCardsRequest{
		BuyerID: buyer.ID,
	}); err != nil {
		t.Fatalf("list buyer gift cards: %v", err)
	}
}
