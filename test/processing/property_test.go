package processing

import (
	"context"
	"testing"

	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

// TestPropertyBuyerCreate explores generated buyer inputs (mixed-case metadata,
// varied display names) from a fixed seed. The run count is bounded by
// harness.Runs to keep live-sandbox cost low. Override the seed with FC_SEED.
func TestPropertyBuyerCreate(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()
	g := harness.NewGen()

	for i := 0; i < harness.Runs; i++ {
		meta := harness.Metadata(g)
		buyer, err := m.Client.Buyers.Create(ctx, components.BuyerCreate{
			ExternalIdentifier: gr4vygo.String(harness.UniqueID("prop-buyer")),
			DisplayName:        gr4vygo.String("buyer-" + meta["camelCaseKey"]),
		}, nil)
		if err != nil {
			t.Fatalf("run %d: create buyer: %v", i, err)
		}
		if buyer.ID == nil || *buyer.ID == "" {
			t.Fatalf("run %d: buyer id empty", i)
		}
	}
}

// TestPropertyPaymentLinkAmounts exercises payment-link creation across a range
// of generated amounts and currencies.
func TestPropertyPaymentLinkAmounts(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()
	g := harness.NewGen()

	for i := 0; i < harness.Runs; i++ {
		link, err := m.Client.PaymentLinks.Create(ctx, components.PaymentLinkCreate{
			Amount:             harness.Amount(g),
			Currency:           harness.Currency(g),
			Country:            "US",
			ExternalIdentifier: gr4vygo.String(harness.UniqueID("prop-link")),
		}, nil)
		if err != nil {
			t.Fatalf("run %d: create payment link: %v", i, err)
		}
		if link.ID == "" {
			t.Fatalf("run %d: payment link id empty", i)
		}
	}
}
