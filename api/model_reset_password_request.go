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

// ResetPasswordRequest A request to reset a password for a user.
type ResetPasswordRequest struct {
	// The email address of the user account to reset.
	EmailAddress *string `json:"email_address,omitempty"`
}

// NewResetPasswordRequest instantiates a new ResetPasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResetPasswordRequest() *ResetPasswordRequest {
	this := ResetPasswordRequest{}
	return &this
}

// NewResetPasswordRequestWithDefaults instantiates a new ResetPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResetPasswordRequestWithDefaults() *ResetPasswordRequest {
	this := ResetPasswordRequest{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *ResetPasswordRequest) GetEmailAddress() string {
	if o == nil || o.EmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetPasswordRequest) GetEmailAddressOk() (*string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *ResetPasswordRequest) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *ResetPasswordRequest) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

func (o ResetPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EmailAddress != nil {
		toSerialize["email_address"] = o.EmailAddress
	}
	return json.Marshal(toSerialize)
}

type NullableResetPasswordRequest struct {
	value *ResetPasswordRequest
	isSet bool
}

func (v NullableResetPasswordRequest) Get() *ResetPasswordRequest {
	return v.value
}

func (v *NullableResetPasswordRequest) Set(val *ResetPasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResetPasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResetPasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResetPasswordRequest(val *ResetPasswordRequest) *NullableResetPasswordRequest {
	return &NullableResetPasswordRequest{value: val, isSet: true}
}

func (v NullableResetPasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResetPasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


