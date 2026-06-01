package processing

import (
	"context"
	"testing"

	"github.com/gr4vy/gr4vy-go/test/harness"
)

func TestCheckoutSessionLifecycle(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	session, err := m.Client.CheckoutSessions.Create(ctx, nil, nil)
	if err != nil {
		t.Fatalf("create checkout session: %v", err)
	}
	if session.ID == "" {
		t.Fatal("checkout session id empty")
	}

	got, err := m.Client.CheckoutSessions.Get(ctx, session.ID, nil)
	if err != nil {
		t.Fatalf("get checkout session: %v", err)
	}
	if got.ID != session.ID {
		t.Fatalf("expected %s, got %s", session.ID, got.ID)
	}

	// Securing the approving card into the session must return 204.
	if err := harness.PutCard(session.ID); err != nil {
		t.Fatalf("put card: %v", err)
	}

	if err := m.Client.CheckoutSessions.Delete(ctx, session.ID, nil); err != nil {
		t.Fatalf("delete checkout session: %v", err)
	}
}

func TestCheckoutSessionGetMissingIsReached(t *testing.T) {
	m := harness.Merchant(t)
	harness.Reaches(t, "checkout_sessions.get(missing)", func() error {
		_, err := m.Client.CheckoutSessions.Get(context.Background(), harness.MissingID, nil)
		return err
	})
}
