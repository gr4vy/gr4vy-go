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

// GooglePaySessionRequest Initiates a new session with Google Pay.
type GooglePaySessionRequest struct {
	// Fully qualified domain name of the merchant.
	OriginDomain string `json:"origin_domain"`
}

// NewGooglePaySessionRequest instantiates a new GooglePaySessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGooglePaySessionRequest(originDomain string) *GooglePaySessionRequest {
	this := GooglePaySessionRequest{}
	this.OriginDomain = originDomain
	return &this
}

// NewGooglePaySessionRequestWithDefaults instantiates a new GooglePaySessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGooglePaySessionRequestWithDefaults() *GooglePaySessionRequest {
	this := GooglePaySessionRequest{}
	return &this
}

// GetOriginDomain returns the OriginDomain field value
func (o *GooglePaySessionRequest) GetOriginDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginDomain
}

// GetOriginDomainOk returns a tuple with the OriginDomain field value
// and a boolean to check if the value has been set.
func (o *GooglePaySessionRequest) GetOriginDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginDomain, true
}

// SetOriginDomain sets field value
func (o *GooglePaySessionRequest) SetOriginDomain(v string) {
	o.OriginDomain = v
}

func (o GooglePaySessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["origin_domain"] = o.OriginDomain
	}
	return json.Marshal(toSerialize)
}

type NullableGooglePaySessionRequest struct {
	value *GooglePaySessionRequest
	isSet bool
}

func (v NullableGooglePaySessionRequest) Get() *GooglePaySessionRequest {
	return v.value
}

func (v *NullableGooglePaySessionRequest) Set(val *GooglePaySessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGooglePaySessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGooglePaySessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGooglePaySessionRequest(val *GooglePaySessionRequest) *NullableGooglePaySessionRequest {
	return &NullableGooglePaySessionRequest{value: val, isSet: true}
}

func (v NullableGooglePaySessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGooglePaySessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


