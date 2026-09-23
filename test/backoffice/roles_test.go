package backoffice

import (
	"context"
	"testing"

	"github.com/gr4vy/gr4vy-go/test/harness"
)

func TestRolesList(t *testing.T) {
	m := harness.Merchant(t)
	page, err := m.Client.Roles.List(context.Background(), nil, nil)
	if err != nil {
		t.Fatalf("list roles: %v", err)
	}
	if page == nil {
		t.Fatal("list roles returned nil")
	}
	for i, role := range page.Result.GetItems() {
		if got := role.GetType(); got == nil || *got != "role" {
			t.Errorf("item %d: type = %v, want role", i, got)
		}
		if role.GetID() == "" || role.GetSlug() == "" {
			t.Errorf("item %d: missing id or slug: %+v", i, role)
		}
	}
}
