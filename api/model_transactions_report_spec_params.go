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
)

// TransactionsReportSpecParams Parameters used to configure the report.
type TransactionsReportSpecParams struct {
	// A list of fields for the report.
	Fields *[]string `json:"fields,omitempty"`
	Filters *TransactionsReportSpecParamsFilters `json:"filters,omitempty"`
	// A list of fields to sort the report.
	Sort *[]map[string]interface{} `json:"sort,omitempty"`
}

// NewTransactionsReportSpecParams instantiates a new TransactionsReportSpecParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsReportSpecParams() *TransactionsReportSpecParams {
	this := TransactionsReportSpecParams{}
	return &this
}

// NewTransactionsReportSpecParamsWithDefaults instantiates a new TransactionsReportSpecParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsReportSpecParamsWithDefaults() *TransactionsReportSpecParams {
	this := TransactionsReportSpecParams{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *TransactionsReportSpecParams) GetFields() []string {
	if o == nil || o.Fields == nil {
		var ret []string
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsReportSpecParams) GetFieldsOk() (*[]string, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *TransactionsReportSpecParams) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []string and assigns it to the Fields field.
func (o *TransactionsReportSpecParams) SetFields(v []string) {
	o.Fields = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *TransactionsReportSpecParams) GetFilters() TransactionsReportSpecParamsFilters {
	if o == nil || o.Filters == nil {
		var ret TransactionsReportSpecParamsFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsReportSpecParams) GetFiltersOk() (*TransactionsReportSpecParamsFilters, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *TransactionsReportSpecParams) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given TransactionsReportSpecParamsFilters and assigns it to the Filters field.
func (o *TransactionsReportSpecParams) SetFilters(v TransactionsReportSpecParamsFilters) {
	o.Filters = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *TransactionsReportSpecParams) GetSort() []map[string]interface{} {
	if o == nil || o.Sort == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsReportSpecParams) GetSortOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *TransactionsReportSpecParams) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given []map[string]interface{} and assigns it to the Sort field.
func (o *TransactionsReportSpecParams) SetSort(v []map[string]interface{}) {
	o.Sort = &v
}

func (o TransactionsReportSpecParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsReportSpecParams struct {
	value *TransactionsReportSpecParams
	isSet bool
}

func (v NullableTransactionsReportSpecParams) Get() *TransactionsReportSpecParams {
	return v.value
}

func (v *NullableTransactionsReportSpecParams) Set(val *TransactionsReportSpecParams) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsReportSpecParams) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsReportSpecParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsReportSpecParams(val *TransactionsReportSpecParams) *NullableTransactionsReportSpecParams {
	return &NullableTransactionsReportSpecParams{value: val, isSet: true}
}

func (v NullableTransactionsReportSpecParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsReportSpecParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


