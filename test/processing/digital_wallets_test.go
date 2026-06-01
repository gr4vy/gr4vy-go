package processing

import (
	"context"
	"testing"

	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

func TestDigitalWalletsList(t *testing.T) {
	m := harness.Merchant(t)
	if _, err := m.Client.DigitalWallets.List(context.Background(), nil); err != nil {
		t.Fatalf("list digital wallets: %v", err)
	}
}

func TestDigitalWalletCreateIsReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	// Registering a wallet needs real Apple/Google onboarding credentials, so
	// the mock environment rejects it — we assert the endpoint is reached.
	harness.Reaches(t, "digital_wallets.create", func() error {
		_, err := m.Client.DigitalWallets.Create(ctx, components.DigitalWalletCreate{
			Provider:     components.DigitalWalletProviderApple,
			MerchantName: "E2E Merchant",
		}, nil)
		return err
	})
}
