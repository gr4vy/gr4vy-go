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

// TransactionGiftCardStoredRequest Create a transaction with a stored gift card.
type TransactionGiftCardStoredRequest struct {
	// The ID of the gift card to check a balance for.
	Id string `json:"id"`
	// The monetary amount to charge for this gift card, in the smallest currency unit for the given currency, for example `1299` cents to create an authorization for `$12.99`.  All gift card amounts are subtracted from the total transaction amount. The remainder is charged to the provided `payment_method`.
	Amount int32 `json:"amount"`
}

// NewTransactionGiftCardStoredRequest instantiates a new TransactionGiftCardStoredRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionGiftCardStoredRequest(id string, amount int32) *TransactionGiftCardStoredRequest {
	this := TransactionGiftCardStoredRequest{}
	this.Id = id
	this.Amount = amount
	return &this
}

// NewTransactionGiftCardStoredRequestWithDefaults instantiates a new TransactionGiftCardStoredRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionGiftCardStoredRequestWithDefaults() *TransactionGiftCardStoredRequest {
	this := TransactionGiftCardStoredRequest{}
	return &this
}

// GetId returns the Id field value
func (o *TransactionGiftCardStoredRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransactionGiftCardStoredRequest) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransactionGiftCardStoredRequest) SetId(v string) {
	o.Id = v
}

// GetAmount returns the Amount field value
func (o *TransactionGiftCardStoredRequest) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransactionGiftCardStoredRequest) GetAmountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransactionGiftCardStoredRequest) SetAmount(v int32) {
	o.Amount = v
}

func (o TransactionGiftCardStoredRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionGiftCardStoredRequest struct {
	value *TransactionGiftCardStoredRequest
	isSet bool
}

func (v NullableTransactionGiftCardStoredRequest) Get() *TransactionGiftCardStoredRequest {
	return v.value
}

func (v *NullableTransactionGiftCardStoredRequest) Set(val *TransactionGiftCardStoredRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionGiftCardStoredRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionGiftCardStoredRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionGiftCardStoredRequest(val *TransactionGiftCardStoredRequest) *NullableTransactionGiftCardStoredRequest {
	return &NullableTransactionGiftCardStoredRequest{value: val, isSet: true}
}

func (v NullableTransactionGiftCardStoredRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionGiftCardStoredRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


