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

// Statuses A list of status resources.
type Statuses struct {
	// A list of authorizations.
	Items *[]Status `json:"items,omitempty"`
}

// NewStatuses instantiates a new Statuses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatuses() *Statuses {
	this := Statuses{}
	return &this
}

// NewStatusesWithDefaults instantiates a new Statuses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusesWithDefaults() *Statuses {
	this := Statuses{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Statuses) GetItems() []Status {
	if o == nil || o.Items == nil {
		var ret []Status
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Statuses) GetItemsOk() (*[]Status, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Statuses) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Status and assigns it to the Items field.
func (o *Statuses) SetItems(v []Status) {
	o.Items = &v
}

func (o Statuses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableStatuses struct {
	value *Statuses
	isSet bool
}

func (v NullableStatuses) Get() *Statuses {
	return v.value
}

func (v *NullableStatuses) Set(val *Statuses) {
	v.value = val
	v.isSet = true
}

func (v NullableStatuses) IsSet() bool {
	return v.isSet
}

func (v *NullableStatuses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatuses(val *Statuses) *NullableStatuses {
	return &NullableStatuses{value: val, isSet: true}
}

func (v NullableStatuses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatuses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


