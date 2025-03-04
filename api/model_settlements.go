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

// Settlements A list of settlements.
type Settlements struct {
	// A list of settlements.
	Items *[]Settlement `json:"items,omitempty"`
}

// NewSettlements instantiates a new Settlements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettlements() *Settlements {
	this := Settlements{}
	return &this
}

// NewSettlementsWithDefaults instantiates a new Settlements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettlementsWithDefaults() *Settlements {
	this := Settlements{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Settlements) GetItems() []Settlement {
	if o == nil || o.Items == nil {
		var ret []Settlement
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Settlements) GetItemsOk() (*[]Settlement, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Settlements) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Settlement and assigns it to the Items field.
func (o *Settlements) SetItems(v []Settlement) {
	o.Items = &v
}

func (o Settlements) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableSettlements struct {
	value *Settlements
	isSet bool
}

func (v NullableSettlements) Get() *Settlements {
	return v.value
}

func (v *NullableSettlements) Set(val *Settlements) {
	v.value = val
	v.isSet = true
}

func (v NullableSettlements) IsSet() bool {
	return v.isSet
}

func (v *NullableSettlements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettlements(val *Settlements) *NullableSettlements {
	return &NullableSettlements{value: val, isSet: true}
}

func (v NullableSettlements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettlements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


