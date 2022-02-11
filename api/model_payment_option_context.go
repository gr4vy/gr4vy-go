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

// PaymentOptionContext Additional context specific to the payment option. This is currently only returned for Apple Pay and Google Pay.
type PaymentOptionContext struct {
	// Display name of the merchant as registered with the digital wallet provider.
	MerchantName *string `json:"merchant_name,omitempty"`
	// Card schemes supported by the digital wallet provider.
	SupportedSchemes *[]string `json:"supported_schemes,omitempty"`
}

// NewPaymentOptionContext instantiates a new PaymentOptionContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentOptionContext() *PaymentOptionContext {
	this := PaymentOptionContext{}
	return &this
}

// NewPaymentOptionContextWithDefaults instantiates a new PaymentOptionContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentOptionContextWithDefaults() *PaymentOptionContext {
	this := PaymentOptionContext{}
	return &this
}

// GetMerchantName returns the MerchantName field value if set, zero value otherwise.
func (o *PaymentOptionContext) GetMerchantName() string {
	if o == nil || o.MerchantName == nil {
		var ret string
		return ret
	}
	return *o.MerchantName
}

// GetMerchantNameOk returns a tuple with the MerchantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentOptionContext) GetMerchantNameOk() (*string, bool) {
	if o == nil || o.MerchantName == nil {
		return nil, false
	}
	return o.MerchantName, true
}

// HasMerchantName returns a boolean if a field has been set.
func (o *PaymentOptionContext) HasMerchantName() bool {
	if o != nil && o.MerchantName != nil {
		return true
	}

	return false
}

// SetMerchantName gets a reference to the given string and assigns it to the MerchantName field.
func (o *PaymentOptionContext) SetMerchantName(v string) {
	o.MerchantName = &v
}

// GetSupportedSchemes returns the SupportedSchemes field value if set, zero value otherwise.
func (o *PaymentOptionContext) GetSupportedSchemes() []string {
	if o == nil || o.SupportedSchemes == nil {
		var ret []string
		return ret
	}
	return *o.SupportedSchemes
}

// GetSupportedSchemesOk returns a tuple with the SupportedSchemes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentOptionContext) GetSupportedSchemesOk() (*[]string, bool) {
	if o == nil || o.SupportedSchemes == nil {
		return nil, false
	}
	return o.SupportedSchemes, true
}

// HasSupportedSchemes returns a boolean if a field has been set.
func (o *PaymentOptionContext) HasSupportedSchemes() bool {
	if o != nil && o.SupportedSchemes != nil {
		return true
	}

	return false
}

// SetSupportedSchemes gets a reference to the given []string and assigns it to the SupportedSchemes field.
func (o *PaymentOptionContext) SetSupportedSchemes(v []string) {
	o.SupportedSchemes = &v
}

func (o PaymentOptionContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MerchantName != nil {
		toSerialize["merchant_name"] = o.MerchantName
	}
	if o.SupportedSchemes != nil {
		toSerialize["supported_schemes"] = o.SupportedSchemes
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentOptionContext struct {
	value *PaymentOptionContext
	isSet bool
}

func (v NullablePaymentOptionContext) Get() *PaymentOptionContext {
	return v.value
}

func (v *NullablePaymentOptionContext) Set(val *PaymentOptionContext) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentOptionContext) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentOptionContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentOptionContext(val *PaymentOptionContext) *NullablePaymentOptionContext {
	return &NullablePaymentOptionContext{value: val, isSet: true}
}

func (v NullablePaymentOptionContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentOptionContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


