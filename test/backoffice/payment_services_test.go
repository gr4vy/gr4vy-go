package backoffice

import (
	"context"
	"testing"

	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/models/operations"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

func TestPaymentServicesList(t *testing.T) {
	m := harness.Merchant(t)
	page, err := m.Client.PaymentServices.List(context.Background(), operations.ListPaymentServicesRequest{})
	if err != nil {
		t.Fatalf("list payment services: %v", err)
	}
	if page == nil {
		t.Fatal("list payment services returned nil")
	}
}

func TestPaymentServiceDefinitions(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	if _, err := m.Client.PaymentServiceDefinitions.List(ctx, nil, nil); err != nil {
		t.Fatalf("list payment service definitions: %v", err)
	}

	def, err := m.Client.PaymentServiceDefinitions.Get(ctx, "mock-card")
	if err != nil {
		t.Fatalf("get mock-card definition: %v", err)
	}
	if def.ID != "mock-card" {
		t.Fatalf("expected mock-card, got %s", def.ID)
	}
}

func TestPaymentServiceCreateAndDelete(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	created, err := m.Client.PaymentServices.Create(ctx, components.PaymentServiceCreate{
		DisplayName:                harness.UniqueID("svc"),
		PaymentServiceDefinitionID: "mock-card",
		AcceptedCurrencies:         []string{"USD"},
		AcceptedCountries:          []string{"US"},
		Fields: []components.Field{
			{Key: "merchant_id", Value: m.MerchantAccountID},
		},
	}, nil)
	if err != nil {
		t.Fatalf("create payment service: %v", err)
	}
	if created.ID == nil || *created.ID == "" {
		t.Fatal("payment service id empty")
	}

	// Exercise get / update / session against the real service before deleting.
	if _, err := m.Client.PaymentServices.Get(ctx, *created.ID, nil); err != nil {
		t.Fatalf("get payment service: %v", err)
	}
	if _, err := m.Client.PaymentServices.Update(ctx, *created.ID, components.PaymentServiceUpdate{
		DisplayName: gr4vyStr("renamed"),
	}, nil); err != nil {
		t.Fatalf("update payment service: %v", err)
	}
	harness.Reaches(t, "payment_services.session", func() error {
		_, err := m.Client.PaymentServices.Session(ctx, *created.ID, map[string]any{}, nil)
		return err
	})

	if err := m.Client.PaymentServices.Delete(ctx, *created.ID, nil); err != nil {
		t.Fatalf("delete payment service: %v", err)
	}
}

func TestPaymentServiceVerifyIsReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	harness.Reaches(t, "payment_services.verify", func() error {
		_, err := m.Client.PaymentServices.Verify(ctx, components.VerifyCredentials{
			PaymentServiceDefinitionID: "mock-card",
			Fields: []components.Field{
				{Key: "merchant_id", Value: m.MerchantAccountID},
			},
		}, nil)
		return err
	})
}

func TestPaymentServiceDefinitionSessionIsReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	harness.Reaches(t, "payment_service_definitions.session", func() error {
		_, err := m.Client.PaymentServiceDefinitions.Session(ctx, "mock-card", map[string]any{}, nil)
		return err
	})
}
