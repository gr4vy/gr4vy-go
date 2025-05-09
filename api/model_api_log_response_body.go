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

// ApiLogResponseBody The JSON response body for the log entry.
type ApiLogResponseBody struct {
	// The error code.
	Code *string `json:"code,omitempty"`
	// The error message.
	Message *string `json:"message,omitempty"`
	// The HTTP error code.
	Status *float32 `json:"status,omitempty"`
	// Type of the log entry.
	Type *string `json:"type,omitempty"`
	Details *ApiLogResponseBodyDetails `json:"details,omitempty"`
}

// NewApiLogResponseBody instantiates a new ApiLogResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLogResponseBody() *ApiLogResponseBody {
	this := ApiLogResponseBody{}
	return &this
}

// NewApiLogResponseBodyWithDefaults instantiates a new ApiLogResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLogResponseBodyWithDefaults() *ApiLogResponseBody {
	this := ApiLogResponseBody{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApiLogResponseBody) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLogResponseBody) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApiLogResponseBody) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ApiLogResponseBody) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApiLogResponseBody) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLogResponseBody) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApiLogResponseBody) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApiLogResponseBody) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiLogResponseBody) GetStatus() float32 {
	if o == nil || o.Status == nil {
		var ret float32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLogResponseBody) GetStatusOk() (*float32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiLogResponseBody) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given float32 and assigns it to the Status field.
func (o *ApiLogResponseBody) SetStatus(v float32) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiLogResponseBody) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLogResponseBody) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiLogResponseBody) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiLogResponseBody) SetType(v string) {
	o.Type = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ApiLogResponseBody) GetDetails() ApiLogResponseBodyDetails {
	if o == nil || o.Details == nil {
		var ret ApiLogResponseBodyDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLogResponseBody) GetDetailsOk() (*ApiLogResponseBodyDetails, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ApiLogResponseBody) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given ApiLogResponseBodyDetails and assigns it to the Details field.
func (o *ApiLogResponseBody) SetDetails(v ApiLogResponseBodyDetails) {
	o.Details = &v
}

func (o ApiLogResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableApiLogResponseBody struct {
	value *ApiLogResponseBody
	isSet bool
}

func (v NullableApiLogResponseBody) Get() *ApiLogResponseBody {
	return v.value
}

func (v *NullableApiLogResponseBody) Set(val *ApiLogResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLogResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLogResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLogResponseBody(val *ApiLogResponseBody) *NullableApiLogResponseBody {
	return &NullableApiLogResponseBody{value: val, isSet: true}
}

func (v NullableApiLogResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLogResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


