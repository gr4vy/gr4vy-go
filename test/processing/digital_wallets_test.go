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

// TestDigitalWalletByIDAreReached exercises the per-wallet endpoints. No wallet
// can be provisioned in the mock env, so these hit a missing id and 4xx —
// still reaching (and serialising correctly against) each endpoint.
func TestDigitalWalletByIDAreReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()
	dw := m.Client.DigitalWallets

	harness.Reaches(t, "digital_wallets.get", func() error {
		_, err := dw.Get(ctx, harness.MissingID, nil)
		return err
	})
	harness.Reaches(t, "digital_wallets.update", func() error {
		_, err := dw.Update(ctx, harness.MissingID, components.DigitalWalletUpdate{}, nil)
		return err
	})
	harness.Reaches(t, "digital_wallets.delete", func() error {
		return dw.Delete(ctx, harness.MissingID, nil)
	})
	harness.Reaches(t, "digital_wallets.domains.create", func() error {
		_, err := dw.Domains.Create(ctx, harness.MissingID, components.DigitalWalletDomain{
			DomainName: "example.com",
		}, nil)
		return err
	})
	harness.Reaches(t, "digital_wallets.domains.delete", func() error {
		return dw.Domains.Delete(ctx, harness.MissingID, components.DigitalWalletDomain{
			DomainName: "example.com",
		}, nil)
	})
}

// TestDigitalWalletSessionsAreReached exercises the wallet session endpoints
// (Apple/Google/Click-to-Pay/Paze). These need live wallet credentials, so they
// reach and are cleanly rejected.
func TestDigitalWalletSessionsAreReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()
	s := m.Client.DigitalWallets.Sessions

	harness.Reaches(t, "sessions.apple_pay", func() error {
		_, err := s.ApplePay(ctx, components.ApplePaySessionRequest{
			ValidationURL: "https://apple-pay-gateway.apple.com/paymentservices/startSession",
			DomainName:    "example.com",
		}, nil)
		return err
	})
	harness.Reaches(t, "sessions.google_pay", func() error {
		_, err := s.GooglePay(ctx, components.GooglePaySessionRequest{
			OriginDomain: "example.com",
		}, nil)
		return err
	})
	harness.Reaches(t, "sessions.click_to_pay", func() error {
		_, err := s.ClickToPay(ctx, components.ClickToPaySessionRequest{
			CheckoutSessionID: harness.MissingID,
		})
		return err
	})
	harness.Reaches(t, "sessions.paze", func() error {
		_, err := s.Paze(ctx, components.PazeSessionRequest{
			DomainName: gr4vygoString("example.com"),
		}, nil)
		return err
	})
	harness.Reaches(t, "sessions.paze.create", func() error {
		_, err := s.PazeMobileSessionCreate(ctx, components.PazeMobileSessionCreateRequest{}, nil)
		return err
	})
	harness.Reaches(t, "sessions.paze.review", func() error {
		_, err := s.PazeMobileSessionReview(ctx, components.PazeSessionReviewRequest{}, nil)
		return err
	})
	harness.Reaches(t, "sessions.paze.complete", func() error {
		_, err := s.PazeMobileSessionComplete(ctx, components.PazeSessionCompleteRequest{}, nil)
		return err
	})
}

func gr4vygoString(s string) *string { return &s }
