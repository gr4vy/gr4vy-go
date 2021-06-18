/*
 * Gr4vy API (Beta)
 *
 * Welcome to the Gr4vy API reference documentation. Our API is still very much a work in product and subject to change.
 *
 * API version: 1.0
 * Contact: code@gr4vy.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Openapi

import (
	"encoding/json"
)

// TransactionCaptureRequest A request to capture a transaction.
type TransactionCaptureRequest struct {
	// The (partial) amount to capture.  When omitted blank, this will capture the entire amount.
	Amount *float32 `json:"amount,omitempty"`
}

// NewTransactionCaptureRequest instantiates a new TransactionCaptureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionCaptureRequest() *TransactionCaptureRequest {
	this := TransactionCaptureRequest{}
	return &this
}

// NewTransactionCaptureRequestWithDefaults instantiates a new TransactionCaptureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionCaptureRequestWithDefaults() *TransactionCaptureRequest {
	this := TransactionCaptureRequest{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *TransactionCaptureRequest) GetAmount() float32 {
	if o == nil || o.Amount == nil {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionCaptureRequest) GetAmountOk() (*float32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *TransactionCaptureRequest) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *TransactionCaptureRequest) SetAmount(v float32) {
	o.Amount = &v
}

func (o TransactionCaptureRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionCaptureRequest struct {
	value *TransactionCaptureRequest
	isSet bool
}

func (v NullableTransactionCaptureRequest) Get() *TransactionCaptureRequest {
	return v.value
}

func (v *NullableTransactionCaptureRequest) Set(val *TransactionCaptureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionCaptureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionCaptureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionCaptureRequest(val *TransactionCaptureRequest) *NullableTransactionCaptureRequest {
	return &NullableTransactionCaptureRequest{value: val, isSet: true}
}

func (v NullableTransactionCaptureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionCaptureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


