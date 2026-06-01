package backoffice

import (
	"context"
	"testing"

	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/models/operations"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

// TestAccountUpdaterJobIsReached creates a real stored card and submits it to
// the account-updater. The mock environment has no account-updater service, so
// the job submission is reached and cleanly rejected rather than succeeding.
func TestAccountUpdaterJobIsReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	card := harness.ApprovingCard()
	method, err := m.Client.PaymentMethods.Create(ctx, operations.Body{
		CardPaymentMethodCreate: &card,
	}, nil)
	if err != nil {
		t.Fatalf("create payment method: %v", err)
	}

	harness.Reaches(t, "account_updater.jobs.create", func() error {
		_, err := m.Client.AccountUpdater.Jobs.Create(ctx, components.AccountUpdaterJobCreate{
			PaymentMethodIds: []string{method.ID},
		}, nil)
		return err
	})
}
