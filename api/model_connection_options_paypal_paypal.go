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

// ConnectionOptionsPaypalPaypal Additional options to be passed through to PayPal when processing transactions.
type ConnectionOptionsPaypalPaypal struct {
	// An array with key-value objects representing additional data to be passed to PayPal.
	AdditionalData *[]ConnectionOptionsPaypalPaypalAdditionalData `json:"additional_data,omitempty"`
}

// NewConnectionOptionsPaypalPaypal instantiates a new ConnectionOptionsPaypalPaypal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionOptionsPaypalPaypal() *ConnectionOptionsPaypalPaypal {
	this := ConnectionOptionsPaypalPaypal{}
	return &this
}

// NewConnectionOptionsPaypalPaypalWithDefaults instantiates a new ConnectionOptionsPaypalPaypal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionOptionsPaypalPaypalWithDefaults() *ConnectionOptionsPaypalPaypal {
	this := ConnectionOptionsPaypalPaypal{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *ConnectionOptionsPaypalPaypal) GetAdditionalData() []ConnectionOptionsPaypalPaypalAdditionalData {
	if o == nil || o.AdditionalData == nil {
		var ret []ConnectionOptionsPaypalPaypalAdditionalData
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionOptionsPaypalPaypal) GetAdditionalDataOk() (*[]ConnectionOptionsPaypalPaypalAdditionalData, bool) {
	if o == nil || o.AdditionalData == nil {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *ConnectionOptionsPaypalPaypal) HasAdditionalData() bool {
	if o != nil && o.AdditionalData != nil {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given []ConnectionOptionsPaypalPaypalAdditionalData and assigns it to the AdditionalData field.
func (o *ConnectionOptionsPaypalPaypal) SetAdditionalData(v []ConnectionOptionsPaypalPaypalAdditionalData) {
	o.AdditionalData = &v
}

func (o ConnectionOptionsPaypalPaypal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalData != nil {
		toSerialize["additional_data"] = o.AdditionalData
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionOptionsPaypalPaypal struct {
	value *ConnectionOptionsPaypalPaypal
	isSet bool
}

func (v NullableConnectionOptionsPaypalPaypal) Get() *ConnectionOptionsPaypalPaypal {
	return v.value
}

func (v *NullableConnectionOptionsPaypalPaypal) Set(val *ConnectionOptionsPaypalPaypal) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionOptionsPaypalPaypal) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionOptionsPaypalPaypal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionOptionsPaypalPaypal(val *ConnectionOptionsPaypalPaypal) *NullableConnectionOptionsPaypalPaypal {
	return &NullableConnectionOptionsPaypalPaypal{value: val, isSet: true}
}

func (v NullableConnectionOptionsPaypalPaypal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionOptionsPaypalPaypal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


