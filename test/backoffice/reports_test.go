package backoffice

import (
	"context"
	"testing"

	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/models/operations"
	"github.com/gr4vy/gr4vy-go/test/harness"
)

func TestReportsListEndpoints(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	if _, err := m.Client.Reports.List(ctx, operations.ListReportsRequest{}); err != nil {
		t.Fatalf("list reports: %v", err)
	}
	if _, err := m.Client.ReportExecutions.List(ctx, operations.ListAllReportExecutionsRequest{}); err != nil {
		t.Fatalf("list report executions: %v", err)
	}
}

func TestReportGetAndPutAreReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	harness.Reaches(t, "reports.get(missing)", func() error {
		_, err := m.Client.Reports.Get(ctx, harness.MissingID, nil)
		return err
	})
	harness.Reaches(t, "reports.put(missing)", func() error {
		_, err := m.Client.Reports.Put(ctx, harness.MissingID, components.ReportUpdate{
			Name: gr4vyStr("Renamed"),
		}, nil)
		return err
	})
}

func TestReportExecutionsAreReached(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	harness.Reaches(t, "reports.executions.list", func() error {
		_, err := m.Client.Reports.Executions.List(ctx, harness.MissingID, nil, nil, nil)
		return err
	})
	harness.Reaches(t, "reports.executions.get", func() error {
		_, err := m.Client.Reports.Executions.Get(ctx, harness.MissingID, nil)
		return err
	})
	harness.Reaches(t, "reports.executions.url", func() error {
		_, err := m.Client.Reports.Executions.URL(ctx, harness.MissingID, harness.MissingID, nil, nil)
		return err
	})
}

// TestReportCreateIsHappyPath is a real 2xx create: the Go SDK marshals the
// discriminated `spec` union cleanly (the model is injected via the const tag),
// so creating a transactions report reaches a successful response.
func TestReportCreateIsHappyPath(t *testing.T) {
	m := harness.Merchant(t)
	ctx := context.Background()

	spec := components.CreateSpecTransactions(components.TransactionsReportSpec{
		Params: map[string]any{
			"fields":  []string{"id", "status"},
			"filters": map[string]any{"status": []string{"capture_succeeded"}},
			"sort":    []map[string]any{{"field": "created_at", "order": "desc"}},
		},
	})
	report, err := m.Client.Reports.Create(ctx, components.ReportCreate{
		Name:            harness.UniqueID("report"),
		Schedule:        components.ReportScheduleDaily,
		ScheduleEnabled: false,
		Spec:            spec,
	}, nil)
	if err != nil {
		t.Fatalf("create report: %v", err)
	}
	if report.ID == "" {
		t.Fatal("report id empty")
	}
}

func gr4vyStr(s string) *string { return &s }
