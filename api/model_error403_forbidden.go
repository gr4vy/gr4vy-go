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

// Error403Forbidden Forbidden Error (HTTP 403).
type Error403Forbidden struct {
	// `error`.
	Type *string `json:"type,omitempty"`
	// `forbidden`.
	Code *string `json:"code,omitempty"`
	// `403`.
	Status *int32 `json:"status,omitempty"`
	// Invalid credentials.
	Message *string `json:"message,omitempty"`
	// A list of detail objects that further clarify the reason for the error. Not every error supports more detail.
	Details *[]ErrorDetail `json:"details,omitempty"`
}

// NewError403Forbidden instantiates a new Error403Forbidden object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError403Forbidden() *Error403Forbidden {
	this := Error403Forbidden{}
	return &this
}

// NewError403ForbiddenWithDefaults instantiates a new Error403Forbidden object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewError403ForbiddenWithDefaults() *Error403Forbidden {
	this := Error403Forbidden{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Error403Forbidden) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error403Forbidden) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Error403Forbidden) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Error403Forbidden) SetType(v string) {
	o.Type = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Error403Forbidden) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error403Forbidden) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Error403Forbidden) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Error403Forbidden) SetCode(v string) {
	o.Code = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Error403Forbidden) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error403Forbidden) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Error403Forbidden) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *Error403Forbidden) SetStatus(v int32) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Error403Forbidden) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error403Forbidden) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Error403Forbidden) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Error403Forbidden) SetMessage(v string) {
	o.Message = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *Error403Forbidden) GetDetails() []ErrorDetail {
	if o == nil || o.Details == nil {
		var ret []ErrorDetail
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error403Forbidden) GetDetailsOk() (*[]ErrorDetail, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *Error403Forbidden) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []ErrorDetail and assigns it to the Details field.
func (o *Error403Forbidden) SetDetails(v []ErrorDetail) {
	o.Details = &v
}

func (o Error403Forbidden) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableError403Forbidden struct {
	value *Error403Forbidden
	isSet bool
}

func (v NullableError403Forbidden) Get() *Error403Forbidden {
	return v.value
}

func (v *NullableError403Forbidden) Set(val *Error403Forbidden) {
	v.value = val
	v.isSet = true
}

func (v NullableError403Forbidden) IsSet() bool {
	return v.isSet
}

func (v *NullableError403Forbidden) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError403Forbidden(val *Error403Forbidden) *NullableError403Forbidden {
	return &NullableError403Forbidden{value: val, isSet: true}
}

func (v NullableError403Forbidden) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError403Forbidden) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


