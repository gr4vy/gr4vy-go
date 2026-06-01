package processing

import (
	"context"
	"testing"

	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/models/operations"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

// storeCard creates a real stored card and returns its id.
func storeCard(t *testing.T, m *harness.TestMerchant) string {
	t.Helper()
	card := harness.ApprovingCard()
	created, err := m.Client.PaymentMethods.Create(context.Background(), operations.Body{
		CardPaymentMethodCreate: &card,
	}, nil)
	if err != nil {
		t.Fatalf("create payment method: %v", err)
	}
	return created.ID
}

func TestPaymentMethodUpdate(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()
	id := storeCard(t, m)

	updated, err := m.Client.PaymentMethods.Update(ctx, id, components.PaymentMethodUpdate{
		ExpirationDate: gr4vygo.String("12/40"),
	}, nil)
	if err != nil {
		t.Fatalf("update payment method: %v", err)
	}
	if updated.ID != id {
		t.Fatalf("expected %s, got %s", id, updated.ID)
	}
}

func TestNetworkTokenEndpointsAreReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()
	id := storeCard(t, m)
	nt := m.Client.PaymentMethods.NetworkTokens

	// Provisioning a network token needs a token service; we only assert each
	// endpoint is reached and cleanly handled (a missing token id 4xxs).
	harness.Reaches(t, "network_tokens.create", func() error {
		_, err := nt.Create(ctx, id, components.NetworkTokenCreate{}, nil)
		return err
	})
	harness.Reaches(t, "network_tokens.delete", func() error {
		return nt.Delete(ctx, id, harness.MissingID, nil)
	})
	harness.Reaches(t, "network_tokens.resume", func() error {
		_, err := nt.Resume(ctx, id, harness.MissingID, nil)
		return err
	})
	harness.Reaches(t, "network_tokens.suspend", func() error {
		_, err := nt.Suspend(ctx, id, harness.MissingID, nil)
		return err
	})
	harness.Reaches(t, "network_tokens.cryptogram", func() error {
		_, err := nt.Cryptogram.Create(ctx, id, harness.MissingID, components.CryptogramCreate{}, nil)
		return err
	})
}

func TestPaymentServiceTokensAreReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()
	id := storeCard(t, m)
	pst := m.Client.PaymentMethods.PaymentServiceTokens

	// Listing is a 2xx even when empty; create/delete need a real service token,
	// so they're reach-only.
	if _, err := pst.List(ctx, id, nil, nil); err != nil {
		t.Fatalf("list payment service tokens: %v", err)
	}
	harness.Reaches(t, "payment_service_tokens.create", func() error {
		_, err := pst.Create(ctx, id, components.PaymentServiceTokenCreate{
			PaymentServiceID: harness.MissingID,
			RedirectURL:      "https://example.com/return",
		}, nil)
		return err
	})
	harness.Reaches(t, "payment_service_tokens.delete", func() error {
		return pst.Delete(ctx, id, harness.MissingID, nil)
	})
}

func TestPaymentMethodCRUD(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	card := harness.ApprovingCard()
	created, err := m.Client.PaymentMethods.Create(ctx, operations.Body{
		CardPaymentMethodCreate: &card,
	}, nil)
	if err != nil {
		t.Fatalf("create payment method: %v", err)
	}
	if created.ID == "" {
		t.Fatal("payment method id empty")
	}

	got, err := m.Client.PaymentMethods.Get(ctx, created.ID, nil)
	if err != nil {
		t.Fatalf("get payment method: %v", err)
	}
	if got.ID != created.ID {
		t.Fatalf("expected %s, got %s", created.ID, got.ID)
	}

	if _, err := m.Client.PaymentMethods.List(ctx, operations.ListPaymentMethodsRequest{}); err != nil {
		t.Fatalf("list payment methods: %v", err)
	}

	if err := m.Client.PaymentMethods.Delete(ctx, created.ID, nil); err != nil {
		t.Fatalf("delete payment method: %v", err)
	}
}

func TestNetworkTokensListIsReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	card := harness.ApprovingCard()
	created, err := m.Client.PaymentMethods.Create(ctx, operations.Body{
		CardPaymentMethodCreate: &card,
	}, nil)
	if err != nil {
		t.Fatalf("create payment method: %v", err)
	}

	// Network tokens require a provisioned token service; we only assert the
	// endpoint is reached and cleanly handled.
	harness.Reaches(t, "payment_methods.network_tokens.list", func() error {
		_, err := m.Client.PaymentMethods.NetworkTokens.List(ctx, created.ID, nil)
		return err
	})
}
