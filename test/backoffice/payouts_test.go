package backoffice

import (
	"context"
	"testing"

	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/models/operations"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

func TestPayoutsList(t *testing.T) {
	m := harness.Merchant(t)
	if _, err := m.Client.Payouts.List(context.Background(), operations.ListPayoutsRequest{}); err != nil {
		t.Fatalf("list payouts: %v", err)
	}
}

func TestPayoutGetMissingIsReached(t *testing.T) {
	m := harness.Merchant(t)
	harness.Reaches(t, "payouts.get(missing)", func() error {
		_, err := m.Client.Payouts.Get(context.Background(), harness.MissingID, nil)
		return err
	})
}

func TestPayoutCreateIsReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	// A real payout requires a payout-capable PSP, which the mock environment
	// does not provide, so we assert the endpoint is reached and rejected.
	harness.Reaches(t, "payouts.create", func() error {
		_, err := m.Client.Payouts.Create(ctx, components.PayoutCreate{
			Amount:           1299,
			Currency:         "USD",
			PaymentServiceID: harness.MissingID,
			PaymentMethod: components.CreatePayoutCreatePaymentMethodPaymentMethodCard(components.PaymentMethodCard{
				Number:         harness.ApprovingCardNumber,
				ExpirationDate: harness.CardExpirationDate,
			}),
		}, nil)
		return err
	})
}
