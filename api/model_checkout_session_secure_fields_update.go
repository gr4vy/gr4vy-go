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

// CheckoutSessionSecureFieldsUpdate A request to update the secure fields of a checkout session.
type CheckoutSessionSecureFieldsUpdate struct {
	PaymentMethod *CheckoutSessionFieldsPaymentMethod `json:"payment_method,omitempty"`
}

// NewCheckoutSessionSecureFieldsUpdate instantiates a new CheckoutSessionSecureFieldsUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutSessionSecureFieldsUpdate() *CheckoutSessionSecureFieldsUpdate {
	this := CheckoutSessionSecureFieldsUpdate{}
	return &this
}

// NewCheckoutSessionSecureFieldsUpdateWithDefaults instantiates a new CheckoutSessionSecureFieldsUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutSessionSecureFieldsUpdateWithDefaults() *CheckoutSessionSecureFieldsUpdate {
	this := CheckoutSessionSecureFieldsUpdate{}
	return &this
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *CheckoutSessionSecureFieldsUpdate) GetPaymentMethod() CheckoutSessionFieldsPaymentMethod {
	if o == nil || o.PaymentMethod == nil {
		var ret CheckoutSessionFieldsPaymentMethod
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionSecureFieldsUpdate) GetPaymentMethodOk() (*CheckoutSessionFieldsPaymentMethod, bool) {
	if o == nil || o.PaymentMethod == nil {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *CheckoutSessionSecureFieldsUpdate) HasPaymentMethod() bool {
	if o != nil && o.PaymentMethod != nil {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given CheckoutSessionFieldsPaymentMethod and assigns it to the PaymentMethod field.
func (o *CheckoutSessionSecureFieldsUpdate) SetPaymentMethod(v CheckoutSessionFieldsPaymentMethod) {
	o.PaymentMethod = &v
}

func (o CheckoutSessionSecureFieldsUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentMethod != nil {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	return json.Marshal(toSerialize)
}

type NullableCheckoutSessionSecureFieldsUpdate struct {
	value *CheckoutSessionSecureFieldsUpdate
	isSet bool
}

func (v NullableCheckoutSessionSecureFieldsUpdate) Get() *CheckoutSessionSecureFieldsUpdate {
	return v.value
}

func (v *NullableCheckoutSessionSecureFieldsUpdate) Set(val *CheckoutSessionSecureFieldsUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionSecureFieldsUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionSecureFieldsUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionSecureFieldsUpdate(val *CheckoutSessionSecureFieldsUpdate) *NullableCheckoutSessionSecureFieldsUpdate {
	return &NullableCheckoutSessionSecureFieldsUpdate{value: val, isSet: true}
}

func (v NullableCheckoutSessionSecureFieldsUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionSecureFieldsUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


