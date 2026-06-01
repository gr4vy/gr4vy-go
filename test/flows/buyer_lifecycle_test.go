package flows

import (
	"context"
	"testing"

	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/models/operations"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

func TestBuyerWithStoredMethodAndShipping(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	// Create a buyer.
	buyer, err := m.Client.Buyers.Create(ctx, components.BuyerCreate{
		DisplayName:        gr4vygo.String("Jane Doe"),
		ExternalIdentifier: gr4vygo.String(harness.UniqueID("buyer")),
		BillingDetails:     harness.SampleBillingDetails(),
	}, nil)
	if err != nil {
		t.Fatalf("create buyer: %v", err)
	}
	if buyer.ID == nil || *buyer.ID == "" {
		t.Fatal("buyer id is empty")
	}
	buyerID := *buyer.ID

	// Store a card against the buyer.
	card := harness.ApprovingCard()
	card.BuyerID = &buyerID
	method, err := m.Client.PaymentMethods.Create(ctx, operations.Body{
		CardPaymentMethodCreate: &card,
	}, nil)
	if err != nil {
		t.Fatalf("store payment method: %v", err)
	}
	if method.ID == "" {
		t.Fatal("payment method id is empty")
	}

	// Add shipping details for the buyer.
	shipping, err := m.Client.Buyers.ShippingDetails.Create(ctx, buyerID, components.ShippingDetailsCreate{
		FirstName: gr4vygo.String("Jane"),
		LastName:  gr4vygo.String("Doe"),
		Address:   harness.SampleAddress(),
	}, nil)
	if err != nil {
		t.Fatalf("create shipping details: %v", err)
	}
	if shipping.ID == nil || *shipping.ID == "" {
		t.Fatal("shipping details id is empty")
	}

	// The buyer's stored methods should be listable.
	methods, err := m.Client.Buyers.PaymentMethods.List(ctx, operations.ListBuyerPaymentMethodsRequest{
		BuyerID: &buyerID,
	})
	if err != nil {
		t.Fatalf("list buyer payment methods: %v", err)
	}
	if methods == nil {
		t.Fatal("list buyer payment methods returned nil")
	}
}
