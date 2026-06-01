package backoffice

import (
	"context"
	"testing"

	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

func TestThreeDsScenariosList(t *testing.T) {
	m := harness.Merchant(t)
	if _, err := m.Client.ThreeDsScenarios.List(context.Background(), nil, nil, nil); err != nil {
		t.Fatalf("list 3DS scenarios: %v", err)
	}
}

// TestThreeDsScenarioCreateIsReached submits a 3DS test scenario. A configured
// 3DS service is required for this to succeed, so we assert the endpoint is
// reached and the request body serialises correctly.
func TestThreeDsScenarioCreateIsReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	harness.Reaches(t, "three_ds_scenarios.create", func() error {
		_, err := m.Client.ThreeDsScenarios.Create(ctx, components.ThreeDSecureScenarioCreate{
			Conditions: components.ThreeDSecureScenarioConditions{
				Amount: gr4vygo.Int64(1000),
			},
			Outcome: components.ThreeDSecureScenarioOutcome{
				Authentication: components.ThreeDSecureScenarioOutcomeAuthentication{
					TransactionStatus: components.ThreeDSecureScenarioOutcomeAuthenticationTransactionStatusY,
				},
			},
		}, nil)
		return err
	})
}
