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

// PaymentServiceUpdateFields A field containing a key-value pair for a required field defined by the service for this payment service.
type PaymentServiceUpdateFields struct {
	// The key of the field to set a value for.
	Key string `json:"key"`
	// The value of a field to set.
	Value string `json:"value"`
}

// NewPaymentServiceUpdateFields instantiates a new PaymentServiceUpdateFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentServiceUpdateFields(key string, value string) *PaymentServiceUpdateFields {
	this := PaymentServiceUpdateFields{}
	this.Key = key
	this.Value = value
	return &this
}

// NewPaymentServiceUpdateFieldsWithDefaults instantiates a new PaymentServiceUpdateFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentServiceUpdateFieldsWithDefaults() *PaymentServiceUpdateFields {
	this := PaymentServiceUpdateFields{}
	return &this
}

// GetKey returns the Key field value
func (o *PaymentServiceUpdateFields) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *PaymentServiceUpdateFields) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *PaymentServiceUpdateFields) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *PaymentServiceUpdateFields) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PaymentServiceUpdateFields) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PaymentServiceUpdateFields) SetValue(v string) {
	o.Value = v
}

func (o PaymentServiceUpdateFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentServiceUpdateFields struct {
	value *PaymentServiceUpdateFields
	isSet bool
}

func (v NullablePaymentServiceUpdateFields) Get() *PaymentServiceUpdateFields {
	return v.value
}

func (v *NullablePaymentServiceUpdateFields) Set(val *PaymentServiceUpdateFields) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentServiceUpdateFields) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentServiceUpdateFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentServiceUpdateFields(val *PaymentServiceUpdateFields) *NullablePaymentServiceUpdateFields {
	return &NullablePaymentServiceUpdateFields{value: val, isSet: true}
}

func (v NullablePaymentServiceUpdateFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentServiceUpdateFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


