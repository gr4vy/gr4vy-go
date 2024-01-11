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

// ConnectionOptionsCybersourceAntiFraud Additional options for Cybersource Decision Manager (anti-fraud).
type ConnectionOptionsCybersourceAntiFraud struct {
	// An override for the merchant ID configured for the connector, used in combination with meta keys.
	MetaKeyMerchantId NullableString `json:"meta_key_merchant_id,omitempty"`
	// This is a key-value object for merchant defined data. Each key needs to be a numeric string identifying the MDD field to set. For example, for field 1 set the key to \"1\".
	MerchantDefinedData *map[string]string `json:"merchant_defined_data,omitempty"`
}

// NewConnectionOptionsCybersourceAntiFraud instantiates a new ConnectionOptionsCybersourceAntiFraud object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionOptionsCybersourceAntiFraud() *ConnectionOptionsCybersourceAntiFraud {
	this := ConnectionOptionsCybersourceAntiFraud{}
	return &this
}

// NewConnectionOptionsCybersourceAntiFraudWithDefaults instantiates a new ConnectionOptionsCybersourceAntiFraud object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionOptionsCybersourceAntiFraudWithDefaults() *ConnectionOptionsCybersourceAntiFraud {
	this := ConnectionOptionsCybersourceAntiFraud{}
	return &this
}

// GetMetaKeyMerchantId returns the MetaKeyMerchantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionOptionsCybersourceAntiFraud) GetMetaKeyMerchantId() string {
	if o == nil || o.MetaKeyMerchantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.MetaKeyMerchantId.Get()
}

// GetMetaKeyMerchantIdOk returns a tuple with the MetaKeyMerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionOptionsCybersourceAntiFraud) GetMetaKeyMerchantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MetaKeyMerchantId.Get(), o.MetaKeyMerchantId.IsSet()
}

// HasMetaKeyMerchantId returns a boolean if a field has been set.
func (o *ConnectionOptionsCybersourceAntiFraud) HasMetaKeyMerchantId() bool {
	if o != nil && o.MetaKeyMerchantId.IsSet() {
		return true
	}

	return false
}

// SetMetaKeyMerchantId gets a reference to the given NullableString and assigns it to the MetaKeyMerchantId field.
func (o *ConnectionOptionsCybersourceAntiFraud) SetMetaKeyMerchantId(v string) {
	o.MetaKeyMerchantId.Set(&v)
}
// SetMetaKeyMerchantIdNil sets the value for MetaKeyMerchantId to be an explicit nil
func (o *ConnectionOptionsCybersourceAntiFraud) SetMetaKeyMerchantIdNil() {
	o.MetaKeyMerchantId.Set(nil)
}

// UnsetMetaKeyMerchantId ensures that no value is present for MetaKeyMerchantId, not even an explicit nil
func (o *ConnectionOptionsCybersourceAntiFraud) UnsetMetaKeyMerchantId() {
	o.MetaKeyMerchantId.Unset()
}

// GetMerchantDefinedData returns the MerchantDefinedData field value if set, zero value otherwise.
func (o *ConnectionOptionsCybersourceAntiFraud) GetMerchantDefinedData() map[string]string {
	if o == nil || o.MerchantDefinedData == nil {
		var ret map[string]string
		return ret
	}
	return *o.MerchantDefinedData
}

// GetMerchantDefinedDataOk returns a tuple with the MerchantDefinedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionOptionsCybersourceAntiFraud) GetMerchantDefinedDataOk() (*map[string]string, bool) {
	if o == nil || o.MerchantDefinedData == nil {
		return nil, false
	}
	return o.MerchantDefinedData, true
}

// HasMerchantDefinedData returns a boolean if a field has been set.
func (o *ConnectionOptionsCybersourceAntiFraud) HasMerchantDefinedData() bool {
	if o != nil && o.MerchantDefinedData != nil {
		return true
	}

	return false
}

// SetMerchantDefinedData gets a reference to the given map[string]string and assigns it to the MerchantDefinedData field.
func (o *ConnectionOptionsCybersourceAntiFraud) SetMerchantDefinedData(v map[string]string) {
	o.MerchantDefinedData = &v
}

func (o ConnectionOptionsCybersourceAntiFraud) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MetaKeyMerchantId.IsSet() {
		toSerialize["meta_key_merchant_id"] = o.MetaKeyMerchantId.Get()
	}
	if o.MerchantDefinedData != nil {
		toSerialize["merchant_defined_data"] = o.MerchantDefinedData
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionOptionsCybersourceAntiFraud struct {
	value *ConnectionOptionsCybersourceAntiFraud
	isSet bool
}

func (v NullableConnectionOptionsCybersourceAntiFraud) Get() *ConnectionOptionsCybersourceAntiFraud {
	return v.value
}

func (v *NullableConnectionOptionsCybersourceAntiFraud) Set(val *ConnectionOptionsCybersourceAntiFraud) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionOptionsCybersourceAntiFraud) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionOptionsCybersourceAntiFraud) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionOptionsCybersourceAntiFraud(val *ConnectionOptionsCybersourceAntiFraud) *NullableConnectionOptionsCybersourceAntiFraud {
	return &NullableConnectionOptionsCybersourceAntiFraud{value: val, isSet: true}
}

func (v NullableConnectionOptionsCybersourceAntiFraud) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionOptionsCybersourceAntiFraud) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


