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

// PaymentServiceDefinitionSupportedFeatures Features supported by the payment definition.
type PaymentServiceDefinitionSupportedFeatures struct {
	// Supports storing a payment method via tokenization.
	PaymentMethodTokenization *bool `json:"payment_method_tokenization,omitempty"`
	// Supports hosted 3-D Secure with a redirect.
	ThreeDSecureHosted *bool `json:"three_d_secure_hosted,omitempty"`
	// Supports passing 3-D Secure data to the underlying processor.
	ThreeDSecurePassThrough *bool `json:"three_d_secure_pass_through,omitempty"`
}

// NewPaymentServiceDefinitionSupportedFeatures instantiates a new PaymentServiceDefinitionSupportedFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentServiceDefinitionSupportedFeatures() *PaymentServiceDefinitionSupportedFeatures {
	this := PaymentServiceDefinitionSupportedFeatures{}
	return &this
}

// NewPaymentServiceDefinitionSupportedFeaturesWithDefaults instantiates a new PaymentServiceDefinitionSupportedFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentServiceDefinitionSupportedFeaturesWithDefaults() *PaymentServiceDefinitionSupportedFeatures {
	this := PaymentServiceDefinitionSupportedFeatures{}
	return &this
}

// GetPaymentMethodTokenization returns the PaymentMethodTokenization field value if set, zero value otherwise.
func (o *PaymentServiceDefinitionSupportedFeatures) GetPaymentMethodTokenization() bool {
	if o == nil || o.PaymentMethodTokenization == nil {
		var ret bool
		return ret
	}
	return *o.PaymentMethodTokenization
}

// GetPaymentMethodTokenizationOk returns a tuple with the PaymentMethodTokenization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceDefinitionSupportedFeatures) GetPaymentMethodTokenizationOk() (*bool, bool) {
	if o == nil || o.PaymentMethodTokenization == nil {
		return nil, false
	}
	return o.PaymentMethodTokenization, true
}

// HasPaymentMethodTokenization returns a boolean if a field has been set.
func (o *PaymentServiceDefinitionSupportedFeatures) HasPaymentMethodTokenization() bool {
	if o != nil && o.PaymentMethodTokenization != nil {
		return true
	}

	return false
}

// SetPaymentMethodTokenization gets a reference to the given bool and assigns it to the PaymentMethodTokenization field.
func (o *PaymentServiceDefinitionSupportedFeatures) SetPaymentMethodTokenization(v bool) {
	o.PaymentMethodTokenization = &v
}

// GetThreeDSecureHosted returns the ThreeDSecureHosted field value if set, zero value otherwise.
func (o *PaymentServiceDefinitionSupportedFeatures) GetThreeDSecureHosted() bool {
	if o == nil || o.ThreeDSecureHosted == nil {
		var ret bool
		return ret
	}
	return *o.ThreeDSecureHosted
}

// GetThreeDSecureHostedOk returns a tuple with the ThreeDSecureHosted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceDefinitionSupportedFeatures) GetThreeDSecureHostedOk() (*bool, bool) {
	if o == nil || o.ThreeDSecureHosted == nil {
		return nil, false
	}
	return o.ThreeDSecureHosted, true
}

// HasThreeDSecureHosted returns a boolean if a field has been set.
func (o *PaymentServiceDefinitionSupportedFeatures) HasThreeDSecureHosted() bool {
	if o != nil && o.ThreeDSecureHosted != nil {
		return true
	}

	return false
}

// SetThreeDSecureHosted gets a reference to the given bool and assigns it to the ThreeDSecureHosted field.
func (o *PaymentServiceDefinitionSupportedFeatures) SetThreeDSecureHosted(v bool) {
	o.ThreeDSecureHosted = &v
}

// GetThreeDSecurePassThrough returns the ThreeDSecurePassThrough field value if set, zero value otherwise.
func (o *PaymentServiceDefinitionSupportedFeatures) GetThreeDSecurePassThrough() bool {
	if o == nil || o.ThreeDSecurePassThrough == nil {
		var ret bool
		return ret
	}
	return *o.ThreeDSecurePassThrough
}

// GetThreeDSecurePassThroughOk returns a tuple with the ThreeDSecurePassThrough field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceDefinitionSupportedFeatures) GetThreeDSecurePassThroughOk() (*bool, bool) {
	if o == nil || o.ThreeDSecurePassThrough == nil {
		return nil, false
	}
	return o.ThreeDSecurePassThrough, true
}

// HasThreeDSecurePassThrough returns a boolean if a field has been set.
func (o *PaymentServiceDefinitionSupportedFeatures) HasThreeDSecurePassThrough() bool {
	if o != nil && o.ThreeDSecurePassThrough != nil {
		return true
	}

	return false
}

// SetThreeDSecurePassThrough gets a reference to the given bool and assigns it to the ThreeDSecurePassThrough field.
func (o *PaymentServiceDefinitionSupportedFeatures) SetThreeDSecurePassThrough(v bool) {
	o.ThreeDSecurePassThrough = &v
}

func (o PaymentServiceDefinitionSupportedFeatures) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentMethodTokenization != nil {
		toSerialize["payment_method_tokenization"] = o.PaymentMethodTokenization
	}
	if o.ThreeDSecureHosted != nil {
		toSerialize["three_d_secure_hosted"] = o.ThreeDSecureHosted
	}
	if o.ThreeDSecurePassThrough != nil {
		toSerialize["three_d_secure_pass_through"] = o.ThreeDSecurePassThrough
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentServiceDefinitionSupportedFeatures struct {
	value *PaymentServiceDefinitionSupportedFeatures
	isSet bool
}

func (v NullablePaymentServiceDefinitionSupportedFeatures) Get() *PaymentServiceDefinitionSupportedFeatures {
	return v.value
}

func (v *NullablePaymentServiceDefinitionSupportedFeatures) Set(val *PaymentServiceDefinitionSupportedFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentServiceDefinitionSupportedFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentServiceDefinitionSupportedFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentServiceDefinitionSupportedFeatures(val *PaymentServiceDefinitionSupportedFeatures) *NullablePaymentServiceDefinitionSupportedFeatures {
	return &NullablePaymentServiceDefinitionSupportedFeatures{value: val, isSet: true}
}

func (v NullablePaymentServiceDefinitionSupportedFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentServiceDefinitionSupportedFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


