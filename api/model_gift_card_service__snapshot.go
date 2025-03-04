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

// GiftCardServiceSnapshot A snapshot of a gift card service used in a transaction.
type GiftCardServiceSnapshot struct {
	// The type of this resource.
	Type *string `json:"type,omitempty"`
	// The ID of this gift card service.
	Id *string `json:"id,omitempty"`
	// The ID of the gift card service definition used to create this service. 
	GiftCardServiceDefinitionId *string `json:"gift_card_service_definition_id,omitempty"`
	// The custom name set for this service.
	DisplayName *string `json:"display_name,omitempty"`
}

// NewGiftCardServiceSnapshot instantiates a new GiftCardServiceSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardServiceSnapshot() *GiftCardServiceSnapshot {
	this := GiftCardServiceSnapshot{}
	return &this
}

// NewGiftCardServiceSnapshotWithDefaults instantiates a new GiftCardServiceSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardServiceSnapshotWithDefaults() *GiftCardServiceSnapshot {
	this := GiftCardServiceSnapshot{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GiftCardServiceSnapshot) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardServiceSnapshot) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GiftCardServiceSnapshot) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GiftCardServiceSnapshot) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GiftCardServiceSnapshot) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardServiceSnapshot) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GiftCardServiceSnapshot) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GiftCardServiceSnapshot) SetId(v string) {
	o.Id = &v
}

// GetGiftCardServiceDefinitionId returns the GiftCardServiceDefinitionId field value if set, zero value otherwise.
func (o *GiftCardServiceSnapshot) GetGiftCardServiceDefinitionId() string {
	if o == nil || o.GiftCardServiceDefinitionId == nil {
		var ret string
		return ret
	}
	return *o.GiftCardServiceDefinitionId
}

// GetGiftCardServiceDefinitionIdOk returns a tuple with the GiftCardServiceDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardServiceSnapshot) GetGiftCardServiceDefinitionIdOk() (*string, bool) {
	if o == nil || o.GiftCardServiceDefinitionId == nil {
		return nil, false
	}
	return o.GiftCardServiceDefinitionId, true
}

// HasGiftCardServiceDefinitionId returns a boolean if a field has been set.
func (o *GiftCardServiceSnapshot) HasGiftCardServiceDefinitionId() bool {
	if o != nil && o.GiftCardServiceDefinitionId != nil {
		return true
	}

	return false
}

// SetGiftCardServiceDefinitionId gets a reference to the given string and assigns it to the GiftCardServiceDefinitionId field.
func (o *GiftCardServiceSnapshot) SetGiftCardServiceDefinitionId(v string) {
	o.GiftCardServiceDefinitionId = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *GiftCardServiceSnapshot) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardServiceSnapshot) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *GiftCardServiceSnapshot) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *GiftCardServiceSnapshot) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o GiftCardServiceSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.GiftCardServiceDefinitionId != nil {
		toSerialize["gift_card_service_definition_id"] = o.GiftCardServiceDefinitionId
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	return json.Marshal(toSerialize)
}

type NullableGiftCardServiceSnapshot struct {
	value *GiftCardServiceSnapshot
	isSet bool
}

func (v NullableGiftCardServiceSnapshot) Get() *GiftCardServiceSnapshot {
	return v.value
}

func (v *NullableGiftCardServiceSnapshot) Set(val *GiftCardServiceSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardServiceSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardServiceSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardServiceSnapshot(val *GiftCardServiceSnapshot) *NullableGiftCardServiceSnapshot {
	return &NullableGiftCardServiceSnapshot{value: val, isSet: true}
}

func (v NullableGiftCardServiceSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardServiceSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


