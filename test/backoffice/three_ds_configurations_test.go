package backoffice

import (
	"context"
	"testing"

	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

// TestThreeDsConfigurationsAreReached exercises the merchant-account 3DS
// configuration endpoints. List is a 2xx; create needs a configured 3DS
// acquirer (rejected in the mock env), and update/delete hit a missing id —
// all reaching their endpoints.
func TestThreeDsConfigurationsAreReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()
	cfg := m.Client.MerchantAccounts.ThreeDsConfiguration

	if _, err := cfg.List(ctx, m.MerchantAccountID, nil); err != nil {
		t.Fatalf("list 3DS configurations: %v", err)
	}

	harness.Reaches(t, "three_ds_configurations.create", func() error {
		_, err := cfg.Create(ctx, m.MerchantAccountID, components.MerchantAccountThreeDSConfigurationCreate{
			MerchantAcquirerBin:  "123456",
			MerchantAcquirerID:   "acq-1",
			MerchantName:         "E2E Merchant",
			MerchantCountryCode:  "US",
			MerchantCategoryCode: "5411",
			MerchantURL:          "https://example.com",
			Scheme:               components.CardSchemeVisa,
			Metadata:             map[string]string{},
		})
		return err
	})
	harness.Reaches(t, "three_ds_configurations.update", func() error {
		_, err := cfg.Update(ctx, m.MerchantAccountID, harness.MissingID, components.MerchantAccountThreeDSConfigurationUpdate{})
		return err
	})
	harness.Reaches(t, "three_ds_configurations.delete", func() error {
		return cfg.Delete(ctx, m.MerchantAccountID, harness.MissingID)
	})
}
