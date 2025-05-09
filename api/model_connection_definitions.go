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

// ConnectionDefinitions A list of available connection definitions.
type ConnectionDefinitions struct {
	Items *[]ConnectionDefinition `json:"items,omitempty"`
}

// NewConnectionDefinitions instantiates a new ConnectionDefinitions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionDefinitions() *ConnectionDefinitions {
	this := ConnectionDefinitions{}
	return &this
}

// NewConnectionDefinitionsWithDefaults instantiates a new ConnectionDefinitions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionDefinitionsWithDefaults() *ConnectionDefinitions {
	this := ConnectionDefinitions{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ConnectionDefinitions) GetItems() []ConnectionDefinition {
	if o == nil || o.Items == nil {
		var ret []ConnectionDefinition
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDefinitions) GetItemsOk() (*[]ConnectionDefinition, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ConnectionDefinitions) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ConnectionDefinition and assigns it to the Items field.
func (o *ConnectionDefinitions) SetItems(v []ConnectionDefinition) {
	o.Items = &v
}

func (o ConnectionDefinitions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionDefinitions struct {
	value *ConnectionDefinitions
	isSet bool
}

func (v NullableConnectionDefinitions) Get() *ConnectionDefinitions {
	return v.value
}

func (v *NullableConnectionDefinitions) Set(val *ConnectionDefinitions) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionDefinitions) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionDefinitions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionDefinitions(val *ConnectionDefinitions) *NullableConnectionDefinitions {
	return &NullableConnectionDefinitions{value: val, isSet: true}
}

func (v NullableConnectionDefinitions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionDefinitions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


