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

// PaymentOptions A list of payment options.
type PaymentOptions struct {
	Items *[]PaymentOption `json:"items,omitempty"`
}

// NewPaymentOptions instantiates a new PaymentOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentOptions() *PaymentOptions {
	this := PaymentOptions{}
	return &this
}

// NewPaymentOptionsWithDefaults instantiates a new PaymentOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentOptionsWithDefaults() *PaymentOptions {
	this := PaymentOptions{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PaymentOptions) GetItems() []PaymentOption {
	if o == nil || o.Items == nil {
		var ret []PaymentOption
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentOptions) GetItemsOk() (*[]PaymentOption, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PaymentOptions) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []PaymentOption and assigns it to the Items field.
func (o *PaymentOptions) SetItems(v []PaymentOption) {
	o.Items = &v
}

func (o PaymentOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentOptions struct {
	value *PaymentOptions
	isSet bool
}

func (v NullablePaymentOptions) Get() *PaymentOptions {
	return v.value
}

func (v *NullablePaymentOptions) Set(val *PaymentOptions) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentOptions) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentOptions(val *PaymentOptions) *NullablePaymentOptions {
	return &NullablePaymentOptions{value: val, isSet: true}
}

func (v NullablePaymentOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


