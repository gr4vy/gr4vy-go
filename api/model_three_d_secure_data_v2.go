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

// ThreeDSecureDataV2 struct for ThreeDSecureDataV2
type ThreeDSecureDataV2 struct {
	// The cardholder authentication value or AAV.
	Cavv string `json:"cavv"`
	// The ecommerce indicator for the 3DS transaction.
	Eci string `json:"eci"`
	// The version of 3-D Secure that was used.
	Version string `json:"version"`
	// The transaction status received as part of the authentication request.
	DirectoryResponse string `json:"directory_response"`
	// The scheme/brand of the card that is used for 3-D Secure.
	Scheme NullableString `json:"scheme,omitempty"`
	// The transaction status after a the 3DS challenge. This will be null in case of a frictionless 3DS flow.
	AuthenticationResponse NullableString `json:"authentication_response,omitempty"`
	// The transaction identifier.
	DirectoryTransactionId string `json:"directory_transaction_id"`
}

// NewThreeDSecureDataV2 instantiates a new ThreeDSecureDataV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDSecureDataV2(cavv string, eci string, version string, directoryResponse string, directoryTransactionId string) *ThreeDSecureDataV2 {
	this := ThreeDSecureDataV2{}
	this.Cavv = cavv
	this.Eci = eci
	this.Version = version
	this.DirectoryResponse = directoryResponse
	this.DirectoryTransactionId = directoryTransactionId
	return &this
}

// NewThreeDSecureDataV2WithDefaults instantiates a new ThreeDSecureDataV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDSecureDataV2WithDefaults() *ThreeDSecureDataV2 {
	this := ThreeDSecureDataV2{}
	return &this
}

// GetCavv returns the Cavv field value
func (o *ThreeDSecureDataV2) GetCavv() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cavv
}

// GetCavvOk returns a tuple with the Cavv field value
// and a boolean to check if the value has been set.
func (o *ThreeDSecureDataV2) GetCavvOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cavv, true
}

// SetCavv sets field value
func (o *ThreeDSecureDataV2) SetCavv(v string) {
	o.Cavv = v
}

// GetEci returns the Eci field value
func (o *ThreeDSecureDataV2) GetEci() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Eci
}

// GetEciOk returns a tuple with the Eci field value
// and a boolean to check if the value has been set.
func (o *ThreeDSecureDataV2) GetEciOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Eci, true
}

// SetEci sets field value
func (o *ThreeDSecureDataV2) SetEci(v string) {
	o.Eci = v
}

// GetVersion returns the Version field value
func (o *ThreeDSecureDataV2) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ThreeDSecureDataV2) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ThreeDSecureDataV2) SetVersion(v string) {
	o.Version = v
}

// GetDirectoryResponse returns the DirectoryResponse field value
func (o *ThreeDSecureDataV2) GetDirectoryResponse() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DirectoryResponse
}

// GetDirectoryResponseOk returns a tuple with the DirectoryResponse field value
// and a boolean to check if the value has been set.
func (o *ThreeDSecureDataV2) GetDirectoryResponseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DirectoryResponse, true
}

// SetDirectoryResponse sets field value
func (o *ThreeDSecureDataV2) SetDirectoryResponse(v string) {
	o.DirectoryResponse = v
}

// GetScheme returns the Scheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSecureDataV2) GetScheme() string {
	if o == nil || o.Scheme.Get() == nil {
		var ret string
		return ret
	}
	return *o.Scheme.Get()
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSecureDataV2) GetSchemeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Scheme.Get(), o.Scheme.IsSet()
}

// HasScheme returns a boolean if a field has been set.
func (o *ThreeDSecureDataV2) HasScheme() bool {
	if o != nil && o.Scheme.IsSet() {
		return true
	}

	return false
}

// SetScheme gets a reference to the given NullableString and assigns it to the Scheme field.
func (o *ThreeDSecureDataV2) SetScheme(v string) {
	o.Scheme.Set(&v)
}
// SetSchemeNil sets the value for Scheme to be an explicit nil
func (o *ThreeDSecureDataV2) SetSchemeNil() {
	o.Scheme.Set(nil)
}

// UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
func (o *ThreeDSecureDataV2) UnsetScheme() {
	o.Scheme.Unset()
}

// GetAuthenticationResponse returns the AuthenticationResponse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSecureDataV2) GetAuthenticationResponse() string {
	if o == nil || o.AuthenticationResponse.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationResponse.Get()
}

// GetAuthenticationResponseOk returns a tuple with the AuthenticationResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSecureDataV2) GetAuthenticationResponseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AuthenticationResponse.Get(), o.AuthenticationResponse.IsSet()
}

// HasAuthenticationResponse returns a boolean if a field has been set.
func (o *ThreeDSecureDataV2) HasAuthenticationResponse() bool {
	if o != nil && o.AuthenticationResponse.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationResponse gets a reference to the given NullableString and assigns it to the AuthenticationResponse field.
func (o *ThreeDSecureDataV2) SetAuthenticationResponse(v string) {
	o.AuthenticationResponse.Set(&v)
}
// SetAuthenticationResponseNil sets the value for AuthenticationResponse to be an explicit nil
func (o *ThreeDSecureDataV2) SetAuthenticationResponseNil() {
	o.AuthenticationResponse.Set(nil)
}

// UnsetAuthenticationResponse ensures that no value is present for AuthenticationResponse, not even an explicit nil
func (o *ThreeDSecureDataV2) UnsetAuthenticationResponse() {
	o.AuthenticationResponse.Unset()
}

// GetDirectoryTransactionId returns the DirectoryTransactionId field value
func (o *ThreeDSecureDataV2) GetDirectoryTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DirectoryTransactionId
}

// GetDirectoryTransactionIdOk returns a tuple with the DirectoryTransactionId field value
// and a boolean to check if the value has been set.
func (o *ThreeDSecureDataV2) GetDirectoryTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DirectoryTransactionId, true
}

// SetDirectoryTransactionId sets field value
func (o *ThreeDSecureDataV2) SetDirectoryTransactionId(v string) {
	o.DirectoryTransactionId = v
}

func (o ThreeDSecureDataV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cavv"] = o.Cavv
	}
	if true {
		toSerialize["eci"] = o.Eci
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["directory_response"] = o.DirectoryResponse
	}
	if o.Scheme.IsSet() {
		toSerialize["scheme"] = o.Scheme.Get()
	}
	if o.AuthenticationResponse.IsSet() {
		toSerialize["authentication_response"] = o.AuthenticationResponse.Get()
	}
	if true {
		toSerialize["directory_transaction_id"] = o.DirectoryTransactionId
	}
	return json.Marshal(toSerialize)
}

type NullableThreeDSecureDataV2 struct {
	value *ThreeDSecureDataV2
	isSet bool
}

func (v NullableThreeDSecureDataV2) Get() *ThreeDSecureDataV2 {
	return v.value
}

func (v *NullableThreeDSecureDataV2) Set(val *ThreeDSecureDataV2) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDSecureDataV2) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDSecureDataV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDSecureDataV2(val *ThreeDSecureDataV2) *NullableThreeDSecureDataV2 {
	return &NullableThreeDSecureDataV2{value: val, isSet: true}
}

func (v NullableThreeDSecureDataV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDSecureDataV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


