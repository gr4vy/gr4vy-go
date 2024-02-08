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
	"time"
)

// NetworkTokenProvisionFailed This event logs the request and response details of HTTP calls made to provision a network token, in case the call failed.
type NetworkTokenProvisionFailed struct {
	// The type of this resource. Is always `transaction-event`.
	Type *string `json:"type,omitempty"`
	// The unique identifier for this event.
	Id *string `json:"id,omitempty"`
	// The name of this resource. Is always `network-token-provision-failed`.
	Name *string `json:"name,omitempty"`
	// The date and time when this event was created in our system.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Context *NetworkTokenProvisionFailedContext `json:"context,omitempty"`
}

// NewNetworkTokenProvisionFailed instantiates a new NetworkTokenProvisionFailed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkTokenProvisionFailed() *NetworkTokenProvisionFailed {
	this := NetworkTokenProvisionFailed{}
	return &this
}

// NewNetworkTokenProvisionFailedWithDefaults instantiates a new NetworkTokenProvisionFailed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkTokenProvisionFailedWithDefaults() *NetworkTokenProvisionFailed {
	this := NetworkTokenProvisionFailed{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworkTokenProvisionFailed) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTokenProvisionFailed) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworkTokenProvisionFailed) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworkTokenProvisionFailed) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworkTokenProvisionFailed) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTokenProvisionFailed) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworkTokenProvisionFailed) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworkTokenProvisionFailed) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkTokenProvisionFailed) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTokenProvisionFailed) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkTokenProvisionFailed) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkTokenProvisionFailed) SetName(v string) {
	o.Name = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NetworkTokenProvisionFailed) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTokenProvisionFailed) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NetworkTokenProvisionFailed) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NetworkTokenProvisionFailed) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *NetworkTokenProvisionFailed) GetContext() NetworkTokenProvisionFailedContext {
	if o == nil || o.Context == nil {
		var ret NetworkTokenProvisionFailedContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkTokenProvisionFailed) GetContextOk() (*NetworkTokenProvisionFailedContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *NetworkTokenProvisionFailed) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given NetworkTokenProvisionFailedContext and assigns it to the Context field.
func (o *NetworkTokenProvisionFailed) SetContext(v NetworkTokenProvisionFailedContext) {
	o.Context = &v
}

func (o NetworkTokenProvisionFailed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkTokenProvisionFailed struct {
	value *NetworkTokenProvisionFailed
	isSet bool
}

func (v NullableNetworkTokenProvisionFailed) Get() *NetworkTokenProvisionFailed {
	return v.value
}

func (v *NullableNetworkTokenProvisionFailed) Set(val *NetworkTokenProvisionFailed) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkTokenProvisionFailed) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkTokenProvisionFailed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkTokenProvisionFailed(val *NetworkTokenProvisionFailed) *NullableNetworkTokenProvisionFailed {
	return &NullableNetworkTokenProvisionFailed{value: val, isSet: true}
}

func (v NullableNetworkTokenProvisionFailed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkTokenProvisionFailed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


