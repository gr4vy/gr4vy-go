package processing

import (
	"context"
	"testing"

	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

func TestCardSchemeDefinitionsList(t *testing.T) {
	m := harness.Merchant(t)
	if _, err := m.Client.CardSchemeDefinitions.List(context.Background(), nil); err != nil {
		t.Fatalf("list card scheme definitions: %v", err)
	}
}

func TestPaymentOptionsList(t *testing.T) {
	m := harness.Merchant(t)
	// POST /payment-options evaluates checkout rules and returns the options.
	if _, err := m.Client.PaymentOptions.List(context.Background(), components.PaymentOptionRequest{
		Country:  gr4vyStrP("US"),
		Currency: gr4vyStrP("USD"),
		Amount:   gr4vyInt64P(1299),
	}, nil); err != nil {
		t.Fatalf("list payment options: %v", err)
	}
}

func gr4vyStrP(s string) *string { return &s }
func gr4vyInt64P(i int64) *int64 { return &i }
