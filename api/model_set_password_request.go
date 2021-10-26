/*
Gr4vy API

Welcome to the Gr4vy API reference documentation. Our API is still very much a work in product and subject to change.

API version: 1.1.0-beta
Contact: code@gr4vy.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Openapi

import (
	"encoding/json"
)

// SetPasswordRequest A request to set a password for a user.
type SetPasswordRequest struct {
	// Unique reset token valid for 7 days.
	ResetToken *string `json:"reset_token,omitempty"`
	// The password the user to log in with.
	Password *string `json:"password,omitempty"`
}

// NewSetPasswordRequest instantiates a new SetPasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetPasswordRequest() *SetPasswordRequest {
	this := SetPasswordRequest{}
	return &this
}

// NewSetPasswordRequestWithDefaults instantiates a new SetPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetPasswordRequestWithDefaults() *SetPasswordRequest {
	this := SetPasswordRequest{}
	return &this
}

// GetResetToken returns the ResetToken field value if set, zero value otherwise.
func (o *SetPasswordRequest) GetResetToken() string {
	if o == nil || o.ResetToken == nil {
		var ret string
		return ret
	}
	return *o.ResetToken
}

// GetResetTokenOk returns a tuple with the ResetToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetPasswordRequest) GetResetTokenOk() (*string, bool) {
	if o == nil || o.ResetToken == nil {
		return nil, false
	}
	return o.ResetToken, true
}

// HasResetToken returns a boolean if a field has been set.
func (o *SetPasswordRequest) HasResetToken() bool {
	if o != nil && o.ResetToken != nil {
		return true
	}

	return false
}

// SetResetToken gets a reference to the given string and assigns it to the ResetToken field.
func (o *SetPasswordRequest) SetResetToken(v string) {
	o.ResetToken = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SetPasswordRequest) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetPasswordRequest) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SetPasswordRequest) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SetPasswordRequest) SetPassword(v string) {
	o.Password = &v
}

func (o SetPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResetToken != nil {
		toSerialize["reset_token"] = o.ResetToken
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableSetPasswordRequest struct {
	value *SetPasswordRequest
	isSet bool
}

func (v NullableSetPasswordRequest) Get() *SetPasswordRequest {
	return v.value
}

func (v *NullableSetPasswordRequest) Set(val *SetPasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetPasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetPasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetPasswordRequest(val *SetPasswordRequest) *NullableSetPasswordRequest {
	return &NullableSetPasswordRequest{value: val, isSet: true}
}

func (v NullableSetPasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetPasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


