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

// ConnectionOptionsAdyenSepa Additional options to be passed through to Adyen when processing SEPA transactions.
type ConnectionOptionsAdyenSepa struct {
	// Enabled Adyen's auto-rescue feature.
	AutoRescue *bool `json:"autoRescue,omitempty"`
	// Defines the number of days to auto-retry a payment for if `autoRescue` is enabled.
	MaxDaysToRescue NullableInt32 `json:"maxDaysToRescue,omitempty"`
	// Defines the Adyen auto-rescue test scenario to invoke.
	AutoRescueSepaScenario NullableString `json:"autoRescueSepaScenario,omitempty"`
	// Defines the type of chargeback that you want to simulate.
	OwnerName NullableString `json:"ownerName,omitempty"`
}

// NewConnectionOptionsAdyenSepa instantiates a new ConnectionOptionsAdyenSepa object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionOptionsAdyenSepa() *ConnectionOptionsAdyenSepa {
	this := ConnectionOptionsAdyenSepa{}
	var autoRescue bool = false
	this.AutoRescue = &autoRescue
	var autoRescueSepaScenario string = "null"
	this.AutoRescueSepaScenario = *NewNullableString(&autoRescueSepaScenario)
	var ownerName string = "null"
	this.OwnerName = *NewNullableString(&ownerName)
	return &this
}

// NewConnectionOptionsAdyenSepaWithDefaults instantiates a new ConnectionOptionsAdyenSepa object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionOptionsAdyenSepaWithDefaults() *ConnectionOptionsAdyenSepa {
	this := ConnectionOptionsAdyenSepa{}
	var autoRescue bool = false
	this.AutoRescue = &autoRescue
	var autoRescueSepaScenario string = "null"
	this.AutoRescueSepaScenario = *NewNullableString(&autoRescueSepaScenario)
	var ownerName string = "null"
	this.OwnerName = *NewNullableString(&ownerName)
	return &this
}

// GetAutoRescue returns the AutoRescue field value if set, zero value otherwise.
func (o *ConnectionOptionsAdyenSepa) GetAutoRescue() bool {
	if o == nil || o.AutoRescue == nil {
		var ret bool
		return ret
	}
	return *o.AutoRescue
}

// GetAutoRescueOk returns a tuple with the AutoRescue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionOptionsAdyenSepa) GetAutoRescueOk() (*bool, bool) {
	if o == nil || o.AutoRescue == nil {
		return nil, false
	}
	return o.AutoRescue, true
}

// HasAutoRescue returns a boolean if a field has been set.
func (o *ConnectionOptionsAdyenSepa) HasAutoRescue() bool {
	if o != nil && o.AutoRescue != nil {
		return true
	}

	return false
}

// SetAutoRescue gets a reference to the given bool and assigns it to the AutoRescue field.
func (o *ConnectionOptionsAdyenSepa) SetAutoRescue(v bool) {
	o.AutoRescue = &v
}

// GetMaxDaysToRescue returns the MaxDaysToRescue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionOptionsAdyenSepa) GetMaxDaysToRescue() int32 {
	if o == nil || o.MaxDaysToRescue.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxDaysToRescue.Get()
}

// GetMaxDaysToRescueOk returns a tuple with the MaxDaysToRescue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionOptionsAdyenSepa) GetMaxDaysToRescueOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxDaysToRescue.Get(), o.MaxDaysToRescue.IsSet()
}

// HasMaxDaysToRescue returns a boolean if a field has been set.
func (o *ConnectionOptionsAdyenSepa) HasMaxDaysToRescue() bool {
	if o != nil && o.MaxDaysToRescue.IsSet() {
		return true
	}

	return false
}

// SetMaxDaysToRescue gets a reference to the given NullableInt32 and assigns it to the MaxDaysToRescue field.
func (o *ConnectionOptionsAdyenSepa) SetMaxDaysToRescue(v int32) {
	o.MaxDaysToRescue.Set(&v)
}
// SetMaxDaysToRescueNil sets the value for MaxDaysToRescue to be an explicit nil
func (o *ConnectionOptionsAdyenSepa) SetMaxDaysToRescueNil() {
	o.MaxDaysToRescue.Set(nil)
}

// UnsetMaxDaysToRescue ensures that no value is present for MaxDaysToRescue, not even an explicit nil
func (o *ConnectionOptionsAdyenSepa) UnsetMaxDaysToRescue() {
	o.MaxDaysToRescue.Unset()
}

// GetAutoRescueSepaScenario returns the AutoRescueSepaScenario field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionOptionsAdyenSepa) GetAutoRescueSepaScenario() string {
	if o == nil || o.AutoRescueSepaScenario.Get() == nil {
		var ret string
		return ret
	}
	return *o.AutoRescueSepaScenario.Get()
}

// GetAutoRescueSepaScenarioOk returns a tuple with the AutoRescueSepaScenario field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionOptionsAdyenSepa) GetAutoRescueSepaScenarioOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AutoRescueSepaScenario.Get(), o.AutoRescueSepaScenario.IsSet()
}

// HasAutoRescueSepaScenario returns a boolean if a field has been set.
func (o *ConnectionOptionsAdyenSepa) HasAutoRescueSepaScenario() bool {
	if o != nil && o.AutoRescueSepaScenario.IsSet() {
		return true
	}

	return false
}

// SetAutoRescueSepaScenario gets a reference to the given NullableString and assigns it to the AutoRescueSepaScenario field.
func (o *ConnectionOptionsAdyenSepa) SetAutoRescueSepaScenario(v string) {
	o.AutoRescueSepaScenario.Set(&v)
}
// SetAutoRescueSepaScenarioNil sets the value for AutoRescueSepaScenario to be an explicit nil
func (o *ConnectionOptionsAdyenSepa) SetAutoRescueSepaScenarioNil() {
	o.AutoRescueSepaScenario.Set(nil)
}

// UnsetAutoRescueSepaScenario ensures that no value is present for AutoRescueSepaScenario, not even an explicit nil
func (o *ConnectionOptionsAdyenSepa) UnsetAutoRescueSepaScenario() {
	o.AutoRescueSepaScenario.Unset()
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionOptionsAdyenSepa) GetOwnerName() string {
	if o == nil || o.OwnerName.Get() == nil {
		var ret string
		return ret
	}
	return *o.OwnerName.Get()
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionOptionsAdyenSepa) GetOwnerNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OwnerName.Get(), o.OwnerName.IsSet()
}

// HasOwnerName returns a boolean if a field has been set.
func (o *ConnectionOptionsAdyenSepa) HasOwnerName() bool {
	if o != nil && o.OwnerName.IsSet() {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given NullableString and assigns it to the OwnerName field.
func (o *ConnectionOptionsAdyenSepa) SetOwnerName(v string) {
	o.OwnerName.Set(&v)
}
// SetOwnerNameNil sets the value for OwnerName to be an explicit nil
func (o *ConnectionOptionsAdyenSepa) SetOwnerNameNil() {
	o.OwnerName.Set(nil)
}

// UnsetOwnerName ensures that no value is present for OwnerName, not even an explicit nil
func (o *ConnectionOptionsAdyenSepa) UnsetOwnerName() {
	o.OwnerName.Unset()
}

func (o ConnectionOptionsAdyenSepa) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoRescue != nil {
		toSerialize["autoRescue"] = o.AutoRescue
	}
	if o.MaxDaysToRescue.IsSet() {
		toSerialize["maxDaysToRescue"] = o.MaxDaysToRescue.Get()
	}
	if o.AutoRescueSepaScenario.IsSet() {
		toSerialize["autoRescueSepaScenario"] = o.AutoRescueSepaScenario.Get()
	}
	if o.OwnerName.IsSet() {
		toSerialize["ownerName"] = o.OwnerName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionOptionsAdyenSepa struct {
	value *ConnectionOptionsAdyenSepa
	isSet bool
}

func (v NullableConnectionOptionsAdyenSepa) Get() *ConnectionOptionsAdyenSepa {
	return v.value
}

func (v *NullableConnectionOptionsAdyenSepa) Set(val *ConnectionOptionsAdyenSepa) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionOptionsAdyenSepa) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionOptionsAdyenSepa) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionOptionsAdyenSepa(val *ConnectionOptionsAdyenSepa) *NullableConnectionOptionsAdyenSepa {
	return &NullableConnectionOptionsAdyenSepa{value: val, isSet: true}
}

func (v NullableConnectionOptionsAdyenSepa) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionOptionsAdyenSepa) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


