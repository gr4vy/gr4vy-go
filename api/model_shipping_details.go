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

// ShippingDetails A list of shipping details.
type ShippingDetails struct {
	// A list of shipping details.
	Items *[]ShippingDetail `json:"items,omitempty"`
}

// NewShippingDetails instantiates a new ShippingDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingDetails() *ShippingDetails {
	this := ShippingDetails{}
	return &this
}

// NewShippingDetailsWithDefaults instantiates a new ShippingDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingDetailsWithDefaults() *ShippingDetails {
	this := ShippingDetails{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ShippingDetails) GetItems() []ShippingDetail {
	if o == nil || o.Items == nil {
		var ret []ShippingDetail
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingDetails) GetItemsOk() (*[]ShippingDetail, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ShippingDetails) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ShippingDetail and assigns it to the Items field.
func (o *ShippingDetails) SetItems(v []ShippingDetail) {
	o.Items = &v
}

func (o ShippingDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableShippingDetails struct {
	value *ShippingDetails
	isSet bool
}

func (v NullableShippingDetails) Get() *ShippingDetails {
	return v.value
}

func (v *NullableShippingDetails) Set(val *ShippingDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingDetails(val *ShippingDetails) *NullableShippingDetails {
	return &NullableShippingDetails{value: val, isSet: true}
}

func (v NullableShippingDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


