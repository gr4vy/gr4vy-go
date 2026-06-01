// Package backoffice holds coverage for the back-office configuration surface:
// merchant accounts, payment services, payouts, audit logs, account updater,
// 3DS scenarios and reports. It is one CI shard; see TESTING.md.
package backoffice

import (
	"context"
	"testing"

	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

func TestMerchantAccountListAndGet(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	if _, err := m.Client.MerchantAccounts.List(ctx, nil, nil, nil); err != nil {
		t.Fatalf("list merchant accounts: %v", err)
	}

	got, err := m.Client.MerchantAccounts.Get(ctx, m.MerchantAccountID)
	if err != nil {
		t.Fatalf("get merchant account: %v", err)
	}
	if got.ID != m.MerchantAccountID {
		t.Fatalf("expected %s, got %s", m.MerchantAccountID, got.ID)
	}
}

func TestMerchantAccountUpdate(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	updated, err := m.Client.MerchantAccounts.Update(ctx, m.MerchantAccountID, components.MerchantAccountUpdate{
		DisplayName: gr4vygo.String("Updated " + m.MerchantAccountID),
	})
	if err != nil {
		t.Fatalf("update merchant account: %v", err)
	}
	if updated.DisplayName != "Updated "+m.MerchantAccountID {
		t.Fatalf("display name not updated: %s", updated.DisplayName)
	}
}
