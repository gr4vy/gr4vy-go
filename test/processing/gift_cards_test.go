package processing

import (
	"context"
	"testing"

	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/models/operations"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

func TestGiftCardsListEndpoints(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	// Listing gift cards is a 2xx even when empty.
	if _, err := m.Client.GiftCards.List(ctx, operations.ListGiftCardsRequest{}); err != nil {
		t.Fatalf("list gift cards: %v", err)
	}
}

func TestGiftCardCreateIsReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	// Creating a real gift card needs a provisioned gift-card service, which
	// the mock environment does not have, so we assert the endpoint is reached.
	harness.Reaches(t, "gift_cards.create", func() error {
		_, err := m.Client.GiftCards.Create(ctx, components.GiftCardCreate{
			Number: "4111111111111111",
			Pin:    "1234",
		}, nil)
		return err
	})
}

func TestGiftCardByIDAreReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	// No gift card can be created in the mock env, so these hit a missing id.
	harness.Reaches(t, "gift_cards.get", func() error {
		_, err := m.Client.GiftCards.Get(ctx, harness.MissingID, nil)
		return err
	})
	harness.Reaches(t, "gift_cards.delete", func() error {
		return m.Client.GiftCards.Delete(ctx, harness.MissingID, nil)
	})
}

func TestGiftCardBalancesIsReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	harness.Reaches(t, "gift_cards.balances", func() error {
		_, err := m.Client.GiftCards.Balances.List(ctx, components.GiftCardBalanceRequest{
			Items: []components.Item{
				components.CreateItemGiftCardRequest(components.GiftCardRequest{
					Number: "4111111111111111",
					Pin:    "1234",
				}),
			},
		}, nil)
		return err
	})
}

// TestGiftCardActivationIsReached activates a gift card. Activation needs a
// provisioned gift-card service, so we assert the endpoint is reached.
func TestGiftCardActivationIsReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	pin := "1234"
	amount := int64(1299)
	currency := "USD"

	harness.Reaches(t, "gift_cards.activations.create", func() error {
		_, err := m.Client.GiftCards.Activations.Create(ctx, components.GiftCardActivationCreate{
			Number:   "4111111111111111",
			Pin:      &pin,
			Amount:   &amount,
			Currency: &currency,
		}, nil, nil)
		return err
	})
}

// TestGiftCardIssuanceIsReached issues a new gift card. Issuance needs a
// provisioned gift-card service with a theme, so we assert reach.
func TestGiftCardIssuanceIsReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	harness.Reaches(t, "gift_cards.issuances.create", func() error {
		_, err := m.Client.GiftCards.Issuances.Create(ctx, components.GiftCardIssuanceCreate{
			Theme:    "default",
			Amount:   1299,
			Currency: "USD",
		}, nil, nil)
		return err
	})
}
