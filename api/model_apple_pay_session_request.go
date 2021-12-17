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

// ApplePaySessionRequest Initiates a new session with Apple Pay.
type ApplePaySessionRequest struct {
	// Validation URL obtained from the event passed to a `onvalidatemerchant` callback.
	ValidationUrl string `json:"validation_url"`
	// Fully qualified domain name of the merchant.
	DomainName string `json:"domain_name"`
}

// NewApplePaySessionRequest instantiates a new ApplePaySessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplePaySessionRequest(validationUrl string, domainName string) *ApplePaySessionRequest {
	this := ApplePaySessionRequest{}
	this.ValidationUrl = validationUrl
	this.DomainName = domainName
	return &this
}

// NewApplePaySessionRequestWithDefaults instantiates a new ApplePaySessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplePaySessionRequestWithDefaults() *ApplePaySessionRequest {
	this := ApplePaySessionRequest{}
	return &this
}

// GetValidationUrl returns the ValidationUrl field value
func (o *ApplePaySessionRequest) GetValidationUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidationUrl
}

// GetValidationUrlOk returns a tuple with the ValidationUrl field value
// and a boolean to check if the value has been set.
func (o *ApplePaySessionRequest) GetValidationUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ValidationUrl, true
}

// SetValidationUrl sets field value
func (o *ApplePaySessionRequest) SetValidationUrl(v string) {
	o.ValidationUrl = v
}

// GetDomainName returns the DomainName field value
func (o *ApplePaySessionRequest) GetDomainName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value
// and a boolean to check if the value has been set.
func (o *ApplePaySessionRequest) GetDomainNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DomainName, true
}

// SetDomainName sets field value
func (o *ApplePaySessionRequest) SetDomainName(v string) {
	o.DomainName = v
}

func (o ApplePaySessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["validation_url"] = o.ValidationUrl
	}
	if true {
		toSerialize["domain_name"] = o.DomainName
	}
	return json.Marshal(toSerialize)
}

type NullableApplePaySessionRequest struct {
	value *ApplePaySessionRequest
	isSet bool
}

func (v NullableApplePaySessionRequest) Get() *ApplePaySessionRequest {
	return v.value
}

func (v *NullableApplePaySessionRequest) Set(val *ApplePaySessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplePaySessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplePaySessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplePaySessionRequest(val *ApplePaySessionRequest) *NullableApplePaySessionRequest {
	return &NullableApplePaySessionRequest{value: val, isSet: true}
}

func (v NullableApplePaySessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplePaySessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

