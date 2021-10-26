/*
Gr4vy API

Welcome to the Gr4vy API reference documentation. Our API is still very much a work in product and subject to change.

API version: 1.1.0-beta
Contact: code@gr4vy.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Openapi

import (
	"encoding/json"
)

// BuyerRequest A request to create a buyer.
type BuyerRequest struct {
	// An external identifier that can be used to match the buyer against your own records. This value needs to be unique for all buyers.
	ExternalIdentifier NullableString `json:"external_identifier,omitempty"`
	// A unique name for this buyer which is used in the Gr4vy admin panel to give a buyer a human readable name.
	DisplayName NullableString `json:"display_name,omitempty"`
	// The optional billing details to create a buyer.
	BillingDetails BillingDetails `json:"billing_details,omitempty"`
}

// NewBuyerRequest instantiates a new BuyerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyerRequest() *BuyerRequest {
	this := BuyerRequest{}
	return &this
}

// NewBuyerRequestWithDefaults instantiates a new BuyerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyerRequestWithDefaults() *BuyerRequest {
	this := BuyerRequest{}
	return &this
}

// GetExternalIdentifier returns the ExternalIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BuyerRequest) GetExternalIdentifier() string {
	if o == nil || o.ExternalIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalIdentifier.Get()
}

// GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BuyerRequest) GetExternalIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalIdentifier.Get(), o.ExternalIdentifier.IsSet()
}

// HasExternalIdentifier returns a boolean if a field has been set.
func (o *BuyerRequest) HasExternalIdentifier() bool {
	if o != nil && o.ExternalIdentifier.IsSet() {
		return true
	}

	return false
}

// SetExternalIdentifier gets a reference to the given NullableString and assigns it to the ExternalIdentifier field.
func (o *BuyerRequest) SetExternalIdentifier(v string) {
	o.ExternalIdentifier.Set(&v)
}
// SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil
func (o *BuyerRequest) SetExternalIdentifierNil() {
	o.ExternalIdentifier.Set(nil)
}

// UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
func (o *BuyerRequest) UnsetExternalIdentifier() {
	o.ExternalIdentifier.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BuyerRequest) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BuyerRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *BuyerRequest) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *BuyerRequest) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *BuyerRequest) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *BuyerRequest) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetBillingDetails returns the BillingDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BuyerRequest) GetBillingDetails() BillingDetails {
	if o == nil  {
		var ret BillingDetails
		return ret
	}
	return o.BillingDetails
}

// GetBillingDetailsOk returns a tuple with the BillingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BuyerRequest) GetBillingDetailsOk() (*BillingDetails, bool) {
	if o == nil || o.BillingDetails == nil {
		return nil, false
	}
	return &o.BillingDetails, true
}

// HasBillingDetails returns a boolean if a field has been set.
func (o *BuyerRequest) HasBillingDetails() bool {
	if o != nil && o.BillingDetails != nil {
		return true
	}

	return false
}

// SetBillingDetails gets a reference to the given BillingDetails and assigns it to the BillingDetails field.
func (o *BuyerRequest) SetBillingDetails(v BillingDetails) {
	o.BillingDetails = v
}

func (o BuyerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalIdentifier.IsSet() {
		toSerialize["external_identifier"] = o.ExternalIdentifier.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["display_name"] = o.DisplayName.Get()
	}
	if o.BillingDetails != nil {
		toSerialize["billing_details"] = o.BillingDetails
	}
	return json.Marshal(toSerialize)
}

type NullableBuyerRequest struct {
	value *BuyerRequest
	isSet bool
}

func (v NullableBuyerRequest) Get() *BuyerRequest {
	return v.value
}

func (v *NullableBuyerRequest) Set(val *BuyerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyerRequest(val *BuyerRequest) *NullableBuyerRequest {
	return &NullableBuyerRequest{value: val, isSet: true}
}

func (v NullableBuyerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


