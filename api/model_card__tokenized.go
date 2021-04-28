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

// CardTokenized A mini format version of the card.
type CardTokenized struct {
	// `payment-method`.
	Type *string `json:"type,omitempty"`
	// The unique ID of the payment method.
	Id *string `json:"id,omitempty"`
	// `card`.
	Method *string `json:"method,omitempty"`
	Details *CardDetails `json:"details,omitempty"`
}

// NewCardTokenized instantiates a new CardTokenized object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardTokenized() *CardTokenized {
	this := CardTokenized{}
	return &this
}

// NewCardTokenizedWithDefaults instantiates a new CardTokenized object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardTokenizedWithDefaults() *CardTokenized {
	this := CardTokenized{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CardTokenized) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTokenized) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CardTokenized) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CardTokenized) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CardTokenized) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTokenized) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CardTokenized) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CardTokenized) SetId(v string) {
	o.Id = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *CardTokenized) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTokenized) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *CardTokenized) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *CardTokenized) SetMethod(v string) {
	o.Method = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *CardTokenized) GetDetails() CardDetails {
	if o == nil || o.Details == nil {
		var ret CardDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardTokenized) GetDetailsOk() (*CardDetails, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *CardTokenized) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given CardDetails and assigns it to the Details field.
func (o *CardTokenized) SetDetails(v CardDetails) {
	o.Details = &v
}

func (o CardTokenized) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableCardTokenized struct {
	value *CardTokenized
	isSet bool
}

func (v NullableCardTokenized) Get() *CardTokenized {
	return v.value
}

func (v *NullableCardTokenized) Set(val *CardTokenized) {
	v.value = val
	v.isSet = true
}

func (v NullableCardTokenized) IsSet() bool {
	return v.isSet
}

func (v *NullableCardTokenized) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardTokenized(val *CardTokenized) *NullableCardTokenized {
	return &NullableCardTokenized{value: val, isSet: true}
}

func (v NullableCardTokenized) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardTokenized) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

