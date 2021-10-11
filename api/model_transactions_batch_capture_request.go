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

// TransactionsBatchCaptureRequest A request to capture multiple previously authorized transactions.
type TransactionsBatchCaptureRequest struct {
	// The (partial) amount to capture.  When left blank, this will capture the entire amount.
	Amount int32 `json:"amount"`
	// A supported ISO-4217 currency code. 
	Currency string `json:"currency"`
	// An external identifier that can be used to match the transaction against your own records.
	ExternalIdentifier *string `json:"external_identifier,omitempty"`
	// The ID of the transaction to capture.
	TransactionId string `json:"transaction_id"`
}

// NewTransactionsBatchCaptureRequest instantiates a new TransactionsBatchCaptureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsBatchCaptureRequest(amount int32, currency string, transactionId string) *TransactionsBatchCaptureRequest {
	this := TransactionsBatchCaptureRequest{}
	this.Amount = amount
	this.Currency = currency
	this.TransactionId = transactionId
	return &this
}

// NewTransactionsBatchCaptureRequestWithDefaults instantiates a new TransactionsBatchCaptureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsBatchCaptureRequestWithDefaults() *TransactionsBatchCaptureRequest {
	this := TransactionsBatchCaptureRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *TransactionsBatchCaptureRequest) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransactionsBatchCaptureRequest) GetAmountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransactionsBatchCaptureRequest) SetAmount(v int32) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *TransactionsBatchCaptureRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TransactionsBatchCaptureRequest) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TransactionsBatchCaptureRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetExternalIdentifier returns the ExternalIdentifier field value if set, zero value otherwise.
func (o *TransactionsBatchCaptureRequest) GetExternalIdentifier() string {
	if o == nil || o.ExternalIdentifier == nil {
		var ret string
		return ret
	}
	return *o.ExternalIdentifier
}

// GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsBatchCaptureRequest) GetExternalIdentifierOk() (*string, bool) {
	if o == nil || o.ExternalIdentifier == nil {
		return nil, false
	}
	return o.ExternalIdentifier, true
}

// HasExternalIdentifier returns a boolean if a field has been set.
func (o *TransactionsBatchCaptureRequest) HasExternalIdentifier() bool {
	if o != nil && o.ExternalIdentifier != nil {
		return true
	}

	return false
}

// SetExternalIdentifier gets a reference to the given string and assigns it to the ExternalIdentifier field.
func (o *TransactionsBatchCaptureRequest) SetExternalIdentifier(v string) {
	o.ExternalIdentifier = &v
}

// GetTransactionId returns the TransactionId field value
func (o *TransactionsBatchCaptureRequest) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *TransactionsBatchCaptureRequest) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *TransactionsBatchCaptureRequest) SetTransactionId(v string) {
	o.TransactionId = v
}

func (o TransactionsBatchCaptureRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if o.ExternalIdentifier != nil {
		toSerialize["external_identifier"] = o.ExternalIdentifier
	}
	if true {
		toSerialize["transaction_id"] = o.TransactionId
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsBatchCaptureRequest struct {
	value *TransactionsBatchCaptureRequest
	isSet bool
}

func (v NullableTransactionsBatchCaptureRequest) Get() *TransactionsBatchCaptureRequest {
	return v.value
}

func (v *NullableTransactionsBatchCaptureRequest) Set(val *TransactionsBatchCaptureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsBatchCaptureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsBatchCaptureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsBatchCaptureRequest(val *TransactionsBatchCaptureRequest) *NullableTransactionsBatchCaptureRequest {
	return &NullableTransactionsBatchCaptureRequest{value: val, isSet: true}
}

func (v NullableTransactionsBatchCaptureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsBatchCaptureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


