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

// GiftCardBalanceNewRequest Check the balance for a non-stored gift card.
type GiftCardBalanceNewRequest struct {
	// The 16-19 digit number for this gift card.
	Number string `json:"number"`
	// The PIN for this gift card.
	Pin string `json:"pin"`
}

// NewGiftCardBalanceNewRequest instantiates a new GiftCardBalanceNewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardBalanceNewRequest(number string, pin string) *GiftCardBalanceNewRequest {
	this := GiftCardBalanceNewRequest{}
	this.Number = number
	this.Pin = pin
	return &this
}

// NewGiftCardBalanceNewRequestWithDefaults instantiates a new GiftCardBalanceNewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardBalanceNewRequestWithDefaults() *GiftCardBalanceNewRequest {
	this := GiftCardBalanceNewRequest{}
	return &this
}

// GetNumber returns the Number field value
func (o *GiftCardBalanceNewRequest) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *GiftCardBalanceNewRequest) GetNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *GiftCardBalanceNewRequest) SetNumber(v string) {
	o.Number = v
}

// GetPin returns the Pin field value
func (o *GiftCardBalanceNewRequest) GetPin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pin
}

// GetPinOk returns a tuple with the Pin field value
// and a boolean to check if the value has been set.
func (o *GiftCardBalanceNewRequest) GetPinOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pin, true
}

// SetPin sets field value
func (o *GiftCardBalanceNewRequest) SetPin(v string) {
	o.Pin = v
}

func (o GiftCardBalanceNewRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["number"] = o.Number
	}
	if true {
		toSerialize["pin"] = o.Pin
	}
	return json.Marshal(toSerialize)
}

type NullableGiftCardBalanceNewRequest struct {
	value *GiftCardBalanceNewRequest
	isSet bool
}

func (v NullableGiftCardBalanceNewRequest) Get() *GiftCardBalanceNewRequest {
	return v.value
}

func (v *NullableGiftCardBalanceNewRequest) Set(val *GiftCardBalanceNewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardBalanceNewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardBalanceNewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardBalanceNewRequest(val *GiftCardBalanceNewRequest) *NullableGiftCardBalanceNewRequest {
	return &NullableGiftCardBalanceNewRequest{value: val, isSet: true}
}

func (v NullableGiftCardBalanceNewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardBalanceNewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


