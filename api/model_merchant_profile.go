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

// MerchantProfile Merchant profile for the different card schemes.
type MerchantProfile struct {
	// Merchant profile for Amex.
	Amex NullableMerchantProfileScheme `json:"amex,omitempty"`
	// Merchant profile for Discover.
	Discover NullableMerchantProfileScheme `json:"discover,omitempty"`
	// Merchant profile for Mastercard.
	Mastercard NullableMerchantProfileScheme `json:"mastercard,omitempty"`
	// Merchant profile for Visa.
	Visa NullableMerchantProfileScheme `json:"visa,omitempty"`
}

// NewMerchantProfile instantiates a new MerchantProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantProfile() *MerchantProfile {
	this := MerchantProfile{}
	return &this
}

// NewMerchantProfileWithDefaults instantiates a new MerchantProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantProfileWithDefaults() *MerchantProfile {
	this := MerchantProfile{}
	return &this
}

// GetAmex returns the Amex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantProfile) GetAmex() MerchantProfileScheme {
	if o == nil || o.Amex.Get() == nil {
		var ret MerchantProfileScheme
		return ret
	}
	return *o.Amex.Get()
}

// GetAmexOk returns a tuple with the Amex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantProfile) GetAmexOk() (*MerchantProfileScheme, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Amex.Get(), o.Amex.IsSet()
}

// HasAmex returns a boolean if a field has been set.
func (o *MerchantProfile) HasAmex() bool {
	if o != nil && o.Amex.IsSet() {
		return true
	}

	return false
}

// SetAmex gets a reference to the given NullableMerchantProfileScheme and assigns it to the Amex field.
func (o *MerchantProfile) SetAmex(v MerchantProfileScheme) {
	o.Amex.Set(&v)
}
// SetAmexNil sets the value for Amex to be an explicit nil
func (o *MerchantProfile) SetAmexNil() {
	o.Amex.Set(nil)
}

// UnsetAmex ensures that no value is present for Amex, not even an explicit nil
func (o *MerchantProfile) UnsetAmex() {
	o.Amex.Unset()
}

// GetDiscover returns the Discover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantProfile) GetDiscover() MerchantProfileScheme {
	if o == nil || o.Discover.Get() == nil {
		var ret MerchantProfileScheme
		return ret
	}
	return *o.Discover.Get()
}

// GetDiscoverOk returns a tuple with the Discover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantProfile) GetDiscoverOk() (*MerchantProfileScheme, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Discover.Get(), o.Discover.IsSet()
}

// HasDiscover returns a boolean if a field has been set.
func (o *MerchantProfile) HasDiscover() bool {
	if o != nil && o.Discover.IsSet() {
		return true
	}

	return false
}

// SetDiscover gets a reference to the given NullableMerchantProfileScheme and assigns it to the Discover field.
func (o *MerchantProfile) SetDiscover(v MerchantProfileScheme) {
	o.Discover.Set(&v)
}
// SetDiscoverNil sets the value for Discover to be an explicit nil
func (o *MerchantProfile) SetDiscoverNil() {
	o.Discover.Set(nil)
}

// UnsetDiscover ensures that no value is present for Discover, not even an explicit nil
func (o *MerchantProfile) UnsetDiscover() {
	o.Discover.Unset()
}

// GetMastercard returns the Mastercard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantProfile) GetMastercard() MerchantProfileScheme {
	if o == nil || o.Mastercard.Get() == nil {
		var ret MerchantProfileScheme
		return ret
	}
	return *o.Mastercard.Get()
}

// GetMastercardOk returns a tuple with the Mastercard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantProfile) GetMastercardOk() (*MerchantProfileScheme, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Mastercard.Get(), o.Mastercard.IsSet()
}

// HasMastercard returns a boolean if a field has been set.
func (o *MerchantProfile) HasMastercard() bool {
	if o != nil && o.Mastercard.IsSet() {
		return true
	}

	return false
}

// SetMastercard gets a reference to the given NullableMerchantProfileScheme and assigns it to the Mastercard field.
func (o *MerchantProfile) SetMastercard(v MerchantProfileScheme) {
	o.Mastercard.Set(&v)
}
// SetMastercardNil sets the value for Mastercard to be an explicit nil
func (o *MerchantProfile) SetMastercardNil() {
	o.Mastercard.Set(nil)
}

// UnsetMastercard ensures that no value is present for Mastercard, not even an explicit nil
func (o *MerchantProfile) UnsetMastercard() {
	o.Mastercard.Unset()
}

// GetVisa returns the Visa field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantProfile) GetVisa() MerchantProfileScheme {
	if o == nil || o.Visa.Get() == nil {
		var ret MerchantProfileScheme
		return ret
	}
	return *o.Visa.Get()
}

// GetVisaOk returns a tuple with the Visa field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantProfile) GetVisaOk() (*MerchantProfileScheme, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Visa.Get(), o.Visa.IsSet()
}

// HasVisa returns a boolean if a field has been set.
func (o *MerchantProfile) HasVisa() bool {
	if o != nil && o.Visa.IsSet() {
		return true
	}

	return false
}

// SetVisa gets a reference to the given NullableMerchantProfileScheme and assigns it to the Visa field.
func (o *MerchantProfile) SetVisa(v MerchantProfileScheme) {
	o.Visa.Set(&v)
}
// SetVisaNil sets the value for Visa to be an explicit nil
func (o *MerchantProfile) SetVisaNil() {
	o.Visa.Set(nil)
}

// UnsetVisa ensures that no value is present for Visa, not even an explicit nil
func (o *MerchantProfile) UnsetVisa() {
	o.Visa.Unset()
}

func (o MerchantProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amex.IsSet() {
		toSerialize["amex"] = o.Amex.Get()
	}
	if o.Discover.IsSet() {
		toSerialize["discover"] = o.Discover.Get()
	}
	if o.Mastercard.IsSet() {
		toSerialize["mastercard"] = o.Mastercard.Get()
	}
	if o.Visa.IsSet() {
		toSerialize["visa"] = o.Visa.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMerchantProfile struct {
	value *MerchantProfile
	isSet bool
}

func (v NullableMerchantProfile) Get() *MerchantProfile {
	return v.value
}

func (v *NullableMerchantProfile) Set(val *MerchantProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantProfile(val *MerchantProfile) *NullableMerchantProfile {
	return &NullableMerchantProfile{value: val, isSet: true}
}

func (v NullableMerchantProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


