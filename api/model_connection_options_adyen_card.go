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

// ConnectionOptionsAdyenCard Additional options to be passed through to Adyen when processing card transactions.
type ConnectionOptionsAdyenCard struct {
	// A key-value object representing additional data to be passed to Adyen.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
}

// NewConnectionOptionsAdyenCard instantiates a new ConnectionOptionsAdyenCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionOptionsAdyenCard() *ConnectionOptionsAdyenCard {
	this := ConnectionOptionsAdyenCard{}
	return &this
}

// NewConnectionOptionsAdyenCardWithDefaults instantiates a new ConnectionOptionsAdyenCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionOptionsAdyenCardWithDefaults() *ConnectionOptionsAdyenCard {
	this := ConnectionOptionsAdyenCard{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *ConnectionOptionsAdyenCard) GetAdditionalData() map[string]string {
	if o == nil || o.AdditionalData == nil {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionOptionsAdyenCard) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || o.AdditionalData == nil {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *ConnectionOptionsAdyenCard) HasAdditionalData() bool {
	if o != nil && o.AdditionalData != nil {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *ConnectionOptionsAdyenCard) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

func (o ConnectionOptionsAdyenCard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalData != nil {
		toSerialize["additionalData"] = o.AdditionalData
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionOptionsAdyenCard struct {
	value *ConnectionOptionsAdyenCard
	isSet bool
}

func (v NullableConnectionOptionsAdyenCard) Get() *ConnectionOptionsAdyenCard {
	return v.value
}

func (v *NullableConnectionOptionsAdyenCard) Set(val *ConnectionOptionsAdyenCard) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionOptionsAdyenCard) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionOptionsAdyenCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionOptionsAdyenCard(val *ConnectionOptionsAdyenCard) *NullableConnectionOptionsAdyenCard {
	return &NullableConnectionOptionsAdyenCard{value: val, isSet: true}
}

func (v NullableConnectionOptionsAdyenCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionOptionsAdyenCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


