package processing

import (
	"context"
	"testing"

	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/models/operations"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

func TestBuyerCRUD(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	buyer, err := m.Client.Buyers.Create(ctx, components.BuyerCreate{
		DisplayName:        gr4vygo.String("Acme Inc"),
		ExternalIdentifier: gr4vygo.String(harness.UniqueID("buyer")),
	}, nil)
	if err != nil {
		t.Fatalf("create buyer: %v", err)
	}
	if buyer.ID == nil || *buyer.ID == "" {
		t.Fatal("buyer id empty")
	}
	id := *buyer.ID

	got, err := m.Client.Buyers.Get(ctx, id, nil)
	if err != nil {
		t.Fatalf("get buyer: %v", err)
	}
	if got.ID == nil || *got.ID != id {
		t.Fatal("get buyer mismatch")
	}

	updated, err := m.Client.Buyers.Update(ctx, id, components.BuyerUpdate{
		DisplayName: gr4vygo.String("Acme Corp"),
	}, nil)
	if err != nil {
		t.Fatalf("update buyer: %v", err)
	}
	if updated.DisplayName == nil || *updated.DisplayName != "Acme Corp" {
		t.Fatal("buyer display name not updated")
	}

	if _, err := m.Client.Buyers.List(ctx, operations.ListBuyersRequest{}); err != nil {
		t.Fatalf("list buyers: %v", err)
	}

	if err := m.Client.Buyers.Delete(ctx, id, nil); err != nil {
		t.Fatalf("delete buyer: %v", err)
	}
}

func TestBuyerGetMissingIsReached(t *testing.T) {
	m := harness.Merchant(t)
	harness.Reaches(t, "buyers.get(missing)", func() error {
		_, err := m.Client.Buyers.Get(context.Background(), harness.MissingID, nil)
		return err
	})
}
