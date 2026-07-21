package backoffice

import (
	"context"
	"testing"

	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

// API key pairs are an account-level resource, like merchant accounts: the
// operations take no per-call merchant account id, so we drive them through the
// shard's shared client directly (m.Client) rather than scoping a request to a
// merchant account the way payment/transaction flows do.

func TestApiKeyPairsList(t *testing.T) {
	m := harness.Merchant(t)
	// Happy path: listing needs no pre-existing resource, so it reaches a 2xx.
	if _, err := m.Client.APIKeyPairs.List(context.Background(), nil, nil); err != nil {
		t.Fatalf("list api key pairs: %v", err)
	}
}

func TestApiKeyPairCreateIsReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	// Creating a usable key pair requires a real, existing role; the mock
	// environment has none we can reference, so we point at a nonexistent role
	// and assert the endpoint is reached and cleanly rejected (4xx).
	harness.Reaches(t, "api_key_pairs.create", func() error {
		_, err := m.Client.APIKeyPairs.Create(ctx, components.APIKeyPairCreate{
			DisplayName: "E2E api key pair",
			RoleIds:     []string{harness.MissingID},
		})
		return err
	})
}

func TestApiKeyPairGetMissingIsReached(t *testing.T) {
	m := harness.Merchant(t)
	// No key pair with this id exists, so a fetch reaches the endpoint and is
	// rejected with a 4xx — proving the SDK routes and serialises the request.
	harness.Reaches(t, "api_key_pairs.get(missing)", func() error {
		_, err := m.Client.APIKeyPairs.Get(context.Background(), harness.MissingID)
		return err
	})
}

func TestApiKeyPairUpdateMissingIsReached(t *testing.T) {
	m := harness.Merchant(t)
	// Updating a nonexistent key pair reaches the endpoint and is rejected (4xx).
	harness.Reaches(t, "api_key_pairs.update(missing)", func() error {
		_, err := m.Client.APIKeyPairs.Update(context.Background(), harness.MissingID, components.APIKeyPairUpdate{
			DisplayName: gr4vygo.String("Updated"),
		})
		return err
	})
}

func TestApiKeyPairDeleteMissingIsReached(t *testing.T) {
	m := harness.Merchant(t)
	// Deleting a nonexistent key pair reaches the endpoint and is rejected (4xx).
	harness.Reaches(t, "api_key_pairs.delete(missing)", func() error {
		return m.Client.APIKeyPairs.Delete(context.Background(), harness.MissingID)
	})
}
