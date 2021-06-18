/*
 * Gr4vy API (Beta)
 *
 * Welcome to the Gr4vy API reference documentation. Our API is still very much a work in product and subject to change.
 *
 * API version: 1.0
 * Contact: code@gr4vy.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Openapi

import (
	"encoding/json"
)

// CardRuleCondition Generic card rule condition.
type CardRuleCondition struct {
	// The type of match made for this rule.
	Match string `json:"match"`
	// The transaction field to filter by.
	Key string `json:"key"`
	// The comparison to make to `value` property.
	Operator string `json:"operator"`
	Value string `json:"value"`
}

// NewCardRuleCondition instantiates a new CardRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardRuleCondition(match string, key string, operator string, value string) *CardRuleCondition {
	this := CardRuleCondition{}
	this.Match = match
	this.Key = key
	this.Operator = operator
	this.Value = value
	return &this
}

// NewCardRuleConditionWithDefaults instantiates a new CardRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardRuleConditionWithDefaults() *CardRuleCondition {
	this := CardRuleCondition{}
	return &this
}

// GetMatch returns the Match field value
func (o *CardRuleCondition) GetMatch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Match
}

// GetMatchOk returns a tuple with the Match field value
// and a boolean to check if the value has been set.
func (o *CardRuleCondition) GetMatchOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Match, true
}

// SetMatch sets field value
func (o *CardRuleCondition) SetMatch(v string) {
	o.Match = v
}

// GetKey returns the Key field value
func (o *CardRuleCondition) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CardRuleCondition) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *CardRuleCondition) SetKey(v string) {
	o.Key = v
}

// GetOperator returns the Operator field value
func (o *CardRuleCondition) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *CardRuleCondition) GetOperatorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *CardRuleCondition) SetOperator(v string) {
	o.Operator = v
}

// GetValue returns the Value field value
func (o *CardRuleCondition) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CardRuleCondition) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CardRuleCondition) SetValue(v string) {
	o.Value = v
}

func (o CardRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["match"] = o.Match
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["operator"] = o.Operator
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCardRuleCondition struct {
	value *CardRuleCondition
	isSet bool
}

func (v NullableCardRuleCondition) Get() *CardRuleCondition {
	return v.value
}

func (v *NullableCardRuleCondition) Set(val *CardRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableCardRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableCardRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardRuleCondition(val *CardRuleCondition) *NullableCardRuleCondition {
	return &NullableCardRuleCondition{value: val, isSet: true}
}

func (v NullableCardRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


