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

// BuyerSnapshot Snapshot of a buyer, as used when embedded inside other resources.
type BuyerSnapshot struct {
	// The type of this resource. Is always `buyer`.
	Type *string `json:"type,omitempty"`
	// The unique Gr4vy ID for this buyer.
	Id *string `json:"id,omitempty"`
	// An external identifier that can be used to match the buyer against your own records.
	ExternalIdentifier NullableString `json:"external_identifier,omitempty"`
	// A unique name for this buyer which is used in the Gr4vy admin panel to give a buyer a human readable name.
	DisplayName NullableString `json:"display_name,omitempty"`
	// The billing details associated with the buyer, which include the address and tax ID.
	BillingDetails NullableBillingDetails `json:"billing_details,omitempty"`
}

// NewBuyerSnapshot instantiates a new BuyerSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyerSnapshot() *BuyerSnapshot {
	this := BuyerSnapshot{}
	return &this
}

// NewBuyerSnapshotWithDefaults instantiates a new BuyerSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyerSnapshotWithDefaults() *BuyerSnapshot {
	this := BuyerSnapshot{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BuyerSnapshot) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerSnapshot) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BuyerSnapshot) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BuyerSnapshot) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BuyerSnapshot) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyerSnapshot) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BuyerSnapshot) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BuyerSnapshot) SetId(v string) {
	o.Id = &v
}

// GetExternalIdentifier returns the ExternalIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BuyerSnapshot) GetExternalIdentifier() string {
	if o == nil || o.ExternalIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalIdentifier.Get()
}

// GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BuyerSnapshot) GetExternalIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalIdentifier.Get(), o.ExternalIdentifier.IsSet()
}

// HasExternalIdentifier returns a boolean if a field has been set.
func (o *BuyerSnapshot) HasExternalIdentifier() bool {
	if o != nil && o.ExternalIdentifier.IsSet() {
		return true
	}

	return false
}

// SetExternalIdentifier gets a reference to the given NullableString and assigns it to the ExternalIdentifier field.
func (o *BuyerSnapshot) SetExternalIdentifier(v string) {
	o.ExternalIdentifier.Set(&v)
}
// SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil
func (o *BuyerSnapshot) SetExternalIdentifierNil() {
	o.ExternalIdentifier.Set(nil)
}

// UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
func (o *BuyerSnapshot) UnsetExternalIdentifier() {
	o.ExternalIdentifier.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BuyerSnapshot) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BuyerSnapshot) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *BuyerSnapshot) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *BuyerSnapshot) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *BuyerSnapshot) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *BuyerSnapshot) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetBillingDetails returns the BillingDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BuyerSnapshot) GetBillingDetails() BillingDetails {
	if o == nil || o.BillingDetails.Get() == nil {
		var ret BillingDetails
		return ret
	}
	return *o.BillingDetails.Get()
}

// GetBillingDetailsOk returns a tuple with the BillingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BuyerSnapshot) GetBillingDetailsOk() (*BillingDetails, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BillingDetails.Get(), o.BillingDetails.IsSet()
}

// HasBillingDetails returns a boolean if a field has been set.
func (o *BuyerSnapshot) HasBillingDetails() bool {
	if o != nil && o.BillingDetails.IsSet() {
		return true
	}

	return false
}

// SetBillingDetails gets a reference to the given NullableBillingDetails and assigns it to the BillingDetails field.
func (o *BuyerSnapshot) SetBillingDetails(v BillingDetails) {
	o.BillingDetails.Set(&v)
}
// SetBillingDetailsNil sets the value for BillingDetails to be an explicit nil
func (o *BuyerSnapshot) SetBillingDetailsNil() {
	o.BillingDetails.Set(nil)
}

// UnsetBillingDetails ensures that no value is present for BillingDetails, not even an explicit nil
func (o *BuyerSnapshot) UnsetBillingDetails() {
	o.BillingDetails.Unset()
}

func (o BuyerSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ExternalIdentifier.IsSet() {
		toSerialize["external_identifier"] = o.ExternalIdentifier.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["display_name"] = o.DisplayName.Get()
	}
	if o.BillingDetails.IsSet() {
		toSerialize["billing_details"] = o.BillingDetails.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableBuyerSnapshot struct {
	value *BuyerSnapshot
	isSet bool
}

func (v NullableBuyerSnapshot) Get() *BuyerSnapshot {
	return v.value
}

func (v *NullableBuyerSnapshot) Set(val *BuyerSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyerSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyerSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyerSnapshot(val *BuyerSnapshot) *NullableBuyerSnapshot {
	return &NullableBuyerSnapshot{value: val, isSet: true}
}

func (v NullableBuyerSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyerSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


