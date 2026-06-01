package processing

import (
	"context"
	"testing"

	"github.com/gr4vy/gr4vy-go/models/operations"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

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
