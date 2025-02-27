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

// TransactionRetriesReportSpec The specification of a transaction retries report.
type TransactionRetriesReportSpec struct {
	// The model (dataset) that the data used for the report is retrieved from.
	Model string `json:"model"`
	Params TransactionRetriesReportSpecParams `json:"params"`
}

// NewTransactionRetriesReportSpec instantiates a new TransactionRetriesReportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRetriesReportSpec(model string, params TransactionRetriesReportSpecParams) *TransactionRetriesReportSpec {
	this := TransactionRetriesReportSpec{}
	this.Model = model
	this.Params = params
	return &this
}

// NewTransactionRetriesReportSpecWithDefaults instantiates a new TransactionRetriesReportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRetriesReportSpecWithDefaults() *TransactionRetriesReportSpec {
	this := TransactionRetriesReportSpec{}
	return &this
}

// GetModel returns the Model field value
func (o *TransactionRetriesReportSpec) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *TransactionRetriesReportSpec) GetModelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *TransactionRetriesReportSpec) SetModel(v string) {
	o.Model = v
}

// GetParams returns the Params field value
func (o *TransactionRetriesReportSpec) GetParams() TransactionRetriesReportSpecParams {
	if o == nil {
		var ret TransactionRetriesReportSpecParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *TransactionRetriesReportSpec) GetParamsOk() (*TransactionRetriesReportSpecParams, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *TransactionRetriesReportSpec) SetParams(v TransactionRetriesReportSpecParams) {
	o.Params = v
}

func (o TransactionRetriesReportSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["model"] = o.Model
	}
	if true {
		toSerialize["params"] = o.Params
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionRetriesReportSpec struct {
	value *TransactionRetriesReportSpec
	isSet bool
}

func (v NullableTransactionRetriesReportSpec) Get() *TransactionRetriesReportSpec {
	return v.value
}

func (v *NullableTransactionRetriesReportSpec) Set(val *TransactionRetriesReportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRetriesReportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRetriesReportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRetriesReportSpec(val *TransactionRetriesReportSpec) *NullableTransactionRetriesReportSpec {
	return &NullableTransactionRetriesReportSpec{value: val, isSet: true}
}

func (v NullableTransactionRetriesReportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRetriesReportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


