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

// ConnectionOptionsCybersourceKcp Additional options for Cybersource KCP APM.
type ConnectionOptionsCybersourceKcp struct {
	// An override for the merchant ID configured for the connector, used in combination with meta keys.
	MetaKeyMerchantId NullableString `json:"meta_key_merchant_id,omitempty"`
	// This is a key-value object for merchant defined information. Each key needs to be a numeric string identifying the MDI field to set. For example, for field 1 set the key to \"1\".
	MerchantDefinedInformation *map[string]string `json:"merchant_defined_information,omitempty"`
	// Shipping method for the order.
	ShipToMethod NullableString `json:"ship_to_method,omitempty"`
}

// NewConnectionOptionsCybersourceKcp instantiates a new ConnectionOptionsCybersourceKcp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionOptionsCybersourceKcp() *ConnectionOptionsCybersourceKcp {
	this := ConnectionOptionsCybersourceKcp{}
	return &this
}

// NewConnectionOptionsCybersourceKcpWithDefaults instantiates a new ConnectionOptionsCybersourceKcp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionOptionsCybersourceKcpWithDefaults() *ConnectionOptionsCybersourceKcp {
	this := ConnectionOptionsCybersourceKcp{}
	return &this
}

// GetMetaKeyMerchantId returns the MetaKeyMerchantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionOptionsCybersourceKcp) GetMetaKeyMerchantId() string {
	if o == nil || o.MetaKeyMerchantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.MetaKeyMerchantId.Get()
}

// GetMetaKeyMerchantIdOk returns a tuple with the MetaKeyMerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionOptionsCybersourceKcp) GetMetaKeyMerchantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MetaKeyMerchantId.Get(), o.MetaKeyMerchantId.IsSet()
}

// HasMetaKeyMerchantId returns a boolean if a field has been set.
func (o *ConnectionOptionsCybersourceKcp) HasMetaKeyMerchantId() bool {
	if o != nil && o.MetaKeyMerchantId.IsSet() {
		return true
	}

	return false
}

// SetMetaKeyMerchantId gets a reference to the given NullableString and assigns it to the MetaKeyMerchantId field.
func (o *ConnectionOptionsCybersourceKcp) SetMetaKeyMerchantId(v string) {
	o.MetaKeyMerchantId.Set(&v)
}
// SetMetaKeyMerchantIdNil sets the value for MetaKeyMerchantId to be an explicit nil
func (o *ConnectionOptionsCybersourceKcp) SetMetaKeyMerchantIdNil() {
	o.MetaKeyMerchantId.Set(nil)
}

// UnsetMetaKeyMerchantId ensures that no value is present for MetaKeyMerchantId, not even an explicit nil
func (o *ConnectionOptionsCybersourceKcp) UnsetMetaKeyMerchantId() {
	o.MetaKeyMerchantId.Unset()
}

// GetMerchantDefinedInformation returns the MerchantDefinedInformation field value if set, zero value otherwise.
func (o *ConnectionOptionsCybersourceKcp) GetMerchantDefinedInformation() map[string]string {
	if o == nil || o.MerchantDefinedInformation == nil {
		var ret map[string]string
		return ret
	}
	return *o.MerchantDefinedInformation
}

// GetMerchantDefinedInformationOk returns a tuple with the MerchantDefinedInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionOptionsCybersourceKcp) GetMerchantDefinedInformationOk() (*map[string]string, bool) {
	if o == nil || o.MerchantDefinedInformation == nil {
		return nil, false
	}
	return o.MerchantDefinedInformation, true
}

// HasMerchantDefinedInformation returns a boolean if a field has been set.
func (o *ConnectionOptionsCybersourceKcp) HasMerchantDefinedInformation() bool {
	if o != nil && o.MerchantDefinedInformation != nil {
		return true
	}

	return false
}

// SetMerchantDefinedInformation gets a reference to the given map[string]string and assigns it to the MerchantDefinedInformation field.
func (o *ConnectionOptionsCybersourceKcp) SetMerchantDefinedInformation(v map[string]string) {
	o.MerchantDefinedInformation = &v
}

// GetShipToMethod returns the ShipToMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionOptionsCybersourceKcp) GetShipToMethod() string {
	if o == nil || o.ShipToMethod.Get() == nil {
		var ret string
		return ret
	}
	return *o.ShipToMethod.Get()
}

// GetShipToMethodOk returns a tuple with the ShipToMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionOptionsCybersourceKcp) GetShipToMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ShipToMethod.Get(), o.ShipToMethod.IsSet()
}

// HasShipToMethod returns a boolean if a field has been set.
func (o *ConnectionOptionsCybersourceKcp) HasShipToMethod() bool {
	if o != nil && o.ShipToMethod.IsSet() {
		return true
	}

	return false
}

// SetShipToMethod gets a reference to the given NullableString and assigns it to the ShipToMethod field.
func (o *ConnectionOptionsCybersourceKcp) SetShipToMethod(v string) {
	o.ShipToMethod.Set(&v)
}
// SetShipToMethodNil sets the value for ShipToMethod to be an explicit nil
func (o *ConnectionOptionsCybersourceKcp) SetShipToMethodNil() {
	o.ShipToMethod.Set(nil)
}

// UnsetShipToMethod ensures that no value is present for ShipToMethod, not even an explicit nil
func (o *ConnectionOptionsCybersourceKcp) UnsetShipToMethod() {
	o.ShipToMethod.Unset()
}

func (o ConnectionOptionsCybersourceKcp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MetaKeyMerchantId.IsSet() {
		toSerialize["meta_key_merchant_id"] = o.MetaKeyMerchantId.Get()
	}
	if o.MerchantDefinedInformation != nil {
		toSerialize["merchant_defined_information"] = o.MerchantDefinedInformation
	}
	if o.ShipToMethod.IsSet() {
		toSerialize["ship_to_method"] = o.ShipToMethod.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionOptionsCybersourceKcp struct {
	value *ConnectionOptionsCybersourceKcp
	isSet bool
}

func (v NullableConnectionOptionsCybersourceKcp) Get() *ConnectionOptionsCybersourceKcp {
	return v.value
}

func (v *NullableConnectionOptionsCybersourceKcp) Set(val *ConnectionOptionsCybersourceKcp) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionOptionsCybersourceKcp) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionOptionsCybersourceKcp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionOptionsCybersourceKcp(val *ConnectionOptionsCybersourceKcp) *NullableConnectionOptionsCybersourceKcp {
	return &NullableConnectionOptionsCybersourceKcp{value: val, isSet: true}
}

func (v NullableConnectionOptionsCybersourceKcp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionOptionsCybersourceKcp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


