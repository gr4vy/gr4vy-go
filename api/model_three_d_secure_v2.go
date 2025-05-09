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

// ThreeDSecureV2 struct for ThreeDSecureV2
type ThreeDSecureV2 struct {
	// The cardholder authentication value or AAV.
	Cavv NullableString `json:"cavv,omitempty"`
	// The ecommerce indicator for the 3DS transaction.
	Eci NullableString `json:"eci,omitempty"`
	// The version of 3-D Secure that was used.
	Version *string `json:"version,omitempty"`
	// The transaction status after a the 3DS challenge. This will be null in case of a frictionless 3DS flow.
	AuthenticationResponse NullableString `json:"authentication_response,omitempty"`
	// The transaction status received as part of the authentication request.
	DirectoryResponse NullableString `json:"directory_response,omitempty"`
	// The transaction identifier.
	DirectoryTransactionId NullableString `json:"directory_transaction_id,omitempty"`
	// The reason identifier for a declined transaction.
	TransactionReason NullableString `json:"transaction_reason,omitempty"`
}

// NewThreeDSecureV2 instantiates a new ThreeDSecureV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDSecureV2() *ThreeDSecureV2 {
	this := ThreeDSecureV2{}
	return &this
}

// NewThreeDSecureV2WithDefaults instantiates a new ThreeDSecureV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDSecureV2WithDefaults() *ThreeDSecureV2 {
	this := ThreeDSecureV2{}
	return &this
}

// GetCavv returns the Cavv field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSecureV2) GetCavv() string {
	if o == nil || o.Cavv.Get() == nil {
		var ret string
		return ret
	}
	return *o.Cavv.Get()
}

// GetCavvOk returns a tuple with the Cavv field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSecureV2) GetCavvOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Cavv.Get(), o.Cavv.IsSet()
}

// HasCavv returns a boolean if a field has been set.
func (o *ThreeDSecureV2) HasCavv() bool {
	if o != nil && o.Cavv.IsSet() {
		return true
	}

	return false
}

// SetCavv gets a reference to the given NullableString and assigns it to the Cavv field.
func (o *ThreeDSecureV2) SetCavv(v string) {
	o.Cavv.Set(&v)
}
// SetCavvNil sets the value for Cavv to be an explicit nil
func (o *ThreeDSecureV2) SetCavvNil() {
	o.Cavv.Set(nil)
}

// UnsetCavv ensures that no value is present for Cavv, not even an explicit nil
func (o *ThreeDSecureV2) UnsetCavv() {
	o.Cavv.Unset()
}

// GetEci returns the Eci field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSecureV2) GetEci() string {
	if o == nil || o.Eci.Get() == nil {
		var ret string
		return ret
	}
	return *o.Eci.Get()
}

// GetEciOk returns a tuple with the Eci field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSecureV2) GetEciOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Eci.Get(), o.Eci.IsSet()
}

// HasEci returns a boolean if a field has been set.
func (o *ThreeDSecureV2) HasEci() bool {
	if o != nil && o.Eci.IsSet() {
		return true
	}

	return false
}

// SetEci gets a reference to the given NullableString and assigns it to the Eci field.
func (o *ThreeDSecureV2) SetEci(v string) {
	o.Eci.Set(&v)
}
// SetEciNil sets the value for Eci to be an explicit nil
func (o *ThreeDSecureV2) SetEciNil() {
	o.Eci.Set(nil)
}

// UnsetEci ensures that no value is present for Eci, not even an explicit nil
func (o *ThreeDSecureV2) UnsetEci() {
	o.Eci.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ThreeDSecureV2) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDSecureV2) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ThreeDSecureV2) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ThreeDSecureV2) SetVersion(v string) {
	o.Version = &v
}

// GetAuthenticationResponse returns the AuthenticationResponse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSecureV2) GetAuthenticationResponse() string {
	if o == nil || o.AuthenticationResponse.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationResponse.Get()
}

// GetAuthenticationResponseOk returns a tuple with the AuthenticationResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSecureV2) GetAuthenticationResponseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AuthenticationResponse.Get(), o.AuthenticationResponse.IsSet()
}

// HasAuthenticationResponse returns a boolean if a field has been set.
func (o *ThreeDSecureV2) HasAuthenticationResponse() bool {
	if o != nil && o.AuthenticationResponse.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationResponse gets a reference to the given NullableString and assigns it to the AuthenticationResponse field.
func (o *ThreeDSecureV2) SetAuthenticationResponse(v string) {
	o.AuthenticationResponse.Set(&v)
}
// SetAuthenticationResponseNil sets the value for AuthenticationResponse to be an explicit nil
func (o *ThreeDSecureV2) SetAuthenticationResponseNil() {
	o.AuthenticationResponse.Set(nil)
}

// UnsetAuthenticationResponse ensures that no value is present for AuthenticationResponse, not even an explicit nil
func (o *ThreeDSecureV2) UnsetAuthenticationResponse() {
	o.AuthenticationResponse.Unset()
}

// GetDirectoryResponse returns the DirectoryResponse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSecureV2) GetDirectoryResponse() string {
	if o == nil || o.DirectoryResponse.Get() == nil {
		var ret string
		return ret
	}
	return *o.DirectoryResponse.Get()
}

// GetDirectoryResponseOk returns a tuple with the DirectoryResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSecureV2) GetDirectoryResponseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DirectoryResponse.Get(), o.DirectoryResponse.IsSet()
}

// HasDirectoryResponse returns a boolean if a field has been set.
func (o *ThreeDSecureV2) HasDirectoryResponse() bool {
	if o != nil && o.DirectoryResponse.IsSet() {
		return true
	}

	return false
}

// SetDirectoryResponse gets a reference to the given NullableString and assigns it to the DirectoryResponse field.
func (o *ThreeDSecureV2) SetDirectoryResponse(v string) {
	o.DirectoryResponse.Set(&v)
}
// SetDirectoryResponseNil sets the value for DirectoryResponse to be an explicit nil
func (o *ThreeDSecureV2) SetDirectoryResponseNil() {
	o.DirectoryResponse.Set(nil)
}

// UnsetDirectoryResponse ensures that no value is present for DirectoryResponse, not even an explicit nil
func (o *ThreeDSecureV2) UnsetDirectoryResponse() {
	o.DirectoryResponse.Unset()
}

// GetDirectoryTransactionId returns the DirectoryTransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSecureV2) GetDirectoryTransactionId() string {
	if o == nil || o.DirectoryTransactionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DirectoryTransactionId.Get()
}

// GetDirectoryTransactionIdOk returns a tuple with the DirectoryTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSecureV2) GetDirectoryTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DirectoryTransactionId.Get(), o.DirectoryTransactionId.IsSet()
}

// HasDirectoryTransactionId returns a boolean if a field has been set.
func (o *ThreeDSecureV2) HasDirectoryTransactionId() bool {
	if o != nil && o.DirectoryTransactionId.IsSet() {
		return true
	}

	return false
}

// SetDirectoryTransactionId gets a reference to the given NullableString and assigns it to the DirectoryTransactionId field.
func (o *ThreeDSecureV2) SetDirectoryTransactionId(v string) {
	o.DirectoryTransactionId.Set(&v)
}
// SetDirectoryTransactionIdNil sets the value for DirectoryTransactionId to be an explicit nil
func (o *ThreeDSecureV2) SetDirectoryTransactionIdNil() {
	o.DirectoryTransactionId.Set(nil)
}

// UnsetDirectoryTransactionId ensures that no value is present for DirectoryTransactionId, not even an explicit nil
func (o *ThreeDSecureV2) UnsetDirectoryTransactionId() {
	o.DirectoryTransactionId.Unset()
}

// GetTransactionReason returns the TransactionReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreeDSecureV2) GetTransactionReason() string {
	if o == nil || o.TransactionReason.Get() == nil {
		var ret string
		return ret
	}
	return *o.TransactionReason.Get()
}

// GetTransactionReasonOk returns a tuple with the TransactionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreeDSecureV2) GetTransactionReasonOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TransactionReason.Get(), o.TransactionReason.IsSet()
}

// HasTransactionReason returns a boolean if a field has been set.
func (o *ThreeDSecureV2) HasTransactionReason() bool {
	if o != nil && o.TransactionReason.IsSet() {
		return true
	}

	return false
}

// SetTransactionReason gets a reference to the given NullableString and assigns it to the TransactionReason field.
func (o *ThreeDSecureV2) SetTransactionReason(v string) {
	o.TransactionReason.Set(&v)
}
// SetTransactionReasonNil sets the value for TransactionReason to be an explicit nil
func (o *ThreeDSecureV2) SetTransactionReasonNil() {
	o.TransactionReason.Set(nil)
}

// UnsetTransactionReason ensures that no value is present for TransactionReason, not even an explicit nil
func (o *ThreeDSecureV2) UnsetTransactionReason() {
	o.TransactionReason.Unset()
}

func (o ThreeDSecureV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cavv.IsSet() {
		toSerialize["cavv"] = o.Cavv.Get()
	}
	if o.Eci.IsSet() {
		toSerialize["eci"] = o.Eci.Get()
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.AuthenticationResponse.IsSet() {
		toSerialize["authentication_response"] = o.AuthenticationResponse.Get()
	}
	if o.DirectoryResponse.IsSet() {
		toSerialize["directory_response"] = o.DirectoryResponse.Get()
	}
	if o.DirectoryTransactionId.IsSet() {
		toSerialize["directory_transaction_id"] = o.DirectoryTransactionId.Get()
	}
	if o.TransactionReason.IsSet() {
		toSerialize["transaction_reason"] = o.TransactionReason.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableThreeDSecureV2 struct {
	value *ThreeDSecureV2
	isSet bool
}

func (v NullableThreeDSecureV2) Get() *ThreeDSecureV2 {
	return v.value
}

func (v *NullableThreeDSecureV2) Set(val *ThreeDSecureV2) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDSecureV2) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDSecureV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDSecureV2(val *ThreeDSecureV2) *NullableThreeDSecureV2 {
	return &NullableThreeDSecureV2{value: val, isSet: true}
}

func (v NullableThreeDSecureV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDSecureV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


