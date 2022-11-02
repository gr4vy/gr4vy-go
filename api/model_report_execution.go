/*
 * Gr4vy API
 *
 * Welcome to the Gr4vy API reference documentation. Our API is still very much a work in product and subject to change.
 *
 * API version: 1.1.0-beta
 * Contact: code@gr4vy.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Openapi

import (
	"encoding/json"
	"time"
)

// ReportExecution A report execution.
type ReportExecution struct {
	// The report associated with this report execution.
	Report *ReportSummary `json:"report,omitempty"`
	// The type of this resource. Is always `report-execution`.
	Type *string `json:"type,omitempty"`
	// The unique identifier for this report execution.
	Id *string `json:"id,omitempty"`
	// The date and time this report execution was created in our system.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The date and time this report execution was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The status of this report execution.
	Status *string `json:"status,omitempty"`
	Context *ReportExecutionSummaryContext `json:"context,omitempty"`
}

// NewReportExecution instantiates a new ReportExecution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportExecution() *ReportExecution {
	this := ReportExecution{}
	return &this
}

// NewReportExecutionWithDefaults instantiates a new ReportExecution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportExecutionWithDefaults() *ReportExecution {
	this := ReportExecution{}
	return &this
}

// GetReport returns the Report field value if set, zero value otherwise.
func (o *ReportExecution) GetReport() ReportSummary {
	if o == nil || o.Report == nil {
		var ret ReportSummary
		return ret
	}
	return *o.Report
}

// GetReportOk returns a tuple with the Report field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportExecution) GetReportOk() (*ReportSummary, bool) {
	if o == nil || o.Report == nil {
		return nil, false
	}
	return o.Report, true
}

// HasReport returns a boolean if a field has been set.
func (o *ReportExecution) HasReport() bool {
	if o != nil && o.Report != nil {
		return true
	}

	return false
}

// SetReport gets a reference to the given ReportSummary and assigns it to the Report field.
func (o *ReportExecution) SetReport(v ReportSummary) {
	o.Report = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ReportExecution) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportExecution) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ReportExecution) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ReportExecution) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReportExecution) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportExecution) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReportExecution) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReportExecution) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ReportExecution) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportExecution) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ReportExecution) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ReportExecution) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ReportExecution) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportExecution) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ReportExecution) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ReportExecution) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ReportExecution) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportExecution) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ReportExecution) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ReportExecution) SetStatus(v string) {
	o.Status = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ReportExecution) GetContext() ReportExecutionSummaryContext {
	if o == nil || o.Context == nil {
		var ret ReportExecutionSummaryContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportExecution) GetContextOk() (*ReportExecutionSummaryContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ReportExecution) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given ReportExecutionSummaryContext and assigns it to the Context field.
func (o *ReportExecution) SetContext(v ReportExecutionSummaryContext) {
	o.Context = &v
}

func (o ReportExecution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Report != nil {
		toSerialize["report"] = o.Report
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	return json.Marshal(toSerialize)
}

type NullableReportExecution struct {
	value *ReportExecution
	isSet bool
}

func (v NullableReportExecution) Get() *ReportExecution {
	return v.value
}

func (v *NullableReportExecution) Set(val *ReportExecution) {
	v.value = val
	v.isSet = true
}

func (v NullableReportExecution) IsSet() bool {
	return v.isSet
}

func (v *NullableReportExecution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportExecution(val *ReportExecution) *NullableReportExecution {
	return &NullableReportExecution{value: val, isSet: true}
}

func (v NullableReportExecution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportExecution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


