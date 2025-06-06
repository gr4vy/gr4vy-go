// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gr4vy/gr4vy-go/internal/utils"
)

type SpecType string

const (
	SpecTypeDetailedSettlement SpecType = "detailed_settlement"
	SpecTypeTransactionRetries SpecType = "transaction_retries"
	SpecTypeTransactions       SpecType = "transactions"
)

// Spec - The report specification.
type Spec struct {
	TransactionsReportSpec       *TransactionsReportSpec       `queryParam:"inline"`
	TransactionRetriesReportSpec *TransactionRetriesReportSpec `queryParam:"inline"`
	DetailedSettlementReportSpec *DetailedSettlementReportSpec `queryParam:"inline"`

	Type SpecType
}

func CreateSpecDetailedSettlement(detailedSettlement DetailedSettlementReportSpec) Spec {
	typ := SpecTypeDetailedSettlement

	return Spec{
		DetailedSettlementReportSpec: &detailedSettlement,
		Type:                         typ,
	}
}

func CreateSpecTransactionRetries(transactionRetries TransactionRetriesReportSpec) Spec {
	typ := SpecTypeTransactionRetries

	return Spec{
		TransactionRetriesReportSpec: &transactionRetries,
		Type:                         typ,
	}
}

func CreateSpecTransactions(transactions TransactionsReportSpec) Spec {
	typ := SpecTypeTransactions

	return Spec{
		TransactionsReportSpec: &transactions,
		Type:                   typ,
	}
}

func (u *Spec) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		Model string `json:"model"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.Model {
	case "detailed_settlement":
		detailedSettlementReportSpec := new(DetailedSettlementReportSpec)
		if err := utils.UnmarshalJSON(data, &detailedSettlementReportSpec, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Model == detailed_settlement) type DetailedSettlementReportSpec within Spec: %w", string(data), err)
		}

		u.DetailedSettlementReportSpec = detailedSettlementReportSpec
		u.Type = SpecTypeDetailedSettlement
		return nil
	case "transaction_retries":
		transactionRetriesReportSpec := new(TransactionRetriesReportSpec)
		if err := utils.UnmarshalJSON(data, &transactionRetriesReportSpec, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Model == transaction_retries) type TransactionRetriesReportSpec within Spec: %w", string(data), err)
		}

		u.TransactionRetriesReportSpec = transactionRetriesReportSpec
		u.Type = SpecTypeTransactionRetries
		return nil
	case "transactions":
		transactionsReportSpec := new(TransactionsReportSpec)
		if err := utils.UnmarshalJSON(data, &transactionsReportSpec, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Model == transactions) type TransactionsReportSpec within Spec: %w", string(data), err)
		}

		u.TransactionsReportSpec = transactionsReportSpec
		u.Type = SpecTypeTransactions
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Spec", string(data))
}

func (u Spec) MarshalJSON() ([]byte, error) {
	if u.TransactionsReportSpec != nil {
		return utils.MarshalJSON(u.TransactionsReportSpec, "", true)
	}

	if u.TransactionRetriesReportSpec != nil {
		return utils.MarshalJSON(u.TransactionRetriesReportSpec, "", true)
	}

	if u.DetailedSettlementReportSpec != nil {
		return utils.MarshalJSON(u.DetailedSettlementReportSpec, "", true)
	}

	return nil, errors.New("could not marshal union type Spec: all fields are null")
}

type ReportCreate struct {
	// The name of the report.
	Name string `json:"name"`
	// A description of the report.
	Description *string        `json:"description,omitempty"`
	Schedule    ReportSchedule `json:"schedule"`
	// Whether the report schedule is enabled.
	ScheduleEnabled bool `json:"schedule_enabled"`
	// The timezone for the report schedule.
	ScheduleTimezone *string `default:"Etc/UTC" json:"schedule_timezone"`
	// The report specification.
	Spec Spec `json:"spec"`
}

func (r ReportCreate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ReportCreate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ReportCreate) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ReportCreate) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ReportCreate) GetSchedule() ReportSchedule {
	if o == nil {
		return ReportSchedule("")
	}
	return o.Schedule
}

func (o *ReportCreate) GetScheduleEnabled() bool {
	if o == nil {
		return false
	}
	return o.ScheduleEnabled
}

func (o *ReportCreate) GetScheduleTimezone() *string {
	if o == nil {
		return nil
	}
	return o.ScheduleTimezone
}

func (o *ReportCreate) GetSpec() Spec {
	if o == nil {
		return Spec{}
	}
	return o.Spec
}

func (o *ReportCreate) GetSpecDetailedSettlement() *DetailedSettlementReportSpec {
	return o.GetSpec().DetailedSettlementReportSpec
}

func (o *ReportCreate) GetSpecTransactionRetries() *TransactionRetriesReportSpec {
	return o.GetSpec().TransactionRetriesReportSpec
}

func (o *ReportCreate) GetSpecTransactions() *TransactionsReportSpec {
	return o.GetSpec().TransactionsReportSpec
}
