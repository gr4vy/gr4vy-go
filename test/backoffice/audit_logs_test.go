package backoffice

import (
	"context"
	"testing"

	"github.com/gr4vy/gr4vy-go/models/operations"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

func TestAuditLogsList(t *testing.T) {
	m := harness.Merchant(t)
	page, err := m.Client.AuditLogs.List(context.Background(), operations.ListAuditLogsRequest{})
	if err != nil {
		t.Fatalf("list audit logs: %v", err)
	}
	if page == nil {
		t.Fatal("list audit logs returned nil")
	}
}
