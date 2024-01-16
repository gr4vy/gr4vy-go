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

// TransactionCheckoutSessionRequest Checkout Session payment method details to use in a transaction.
type TransactionCheckoutSessionRequest struct {
	// `checkout-session`.
	Method string `json:"method"`
	// The ID of the Checkout Session.
	Id string `json:"id"`
	// An external identifier that can be used to match the card against your own records. This can only be set if the `store` flag is set to `true`.
	ExternalIdentifier NullableString `json:"external_identifier,omitempty"`
	// We strongly recommend providing a `redirect_url` either when 3-D Secure is enabled and `three_d_secure_data` is not provided, or when using connections where 3DS is enabled. This value will be appended with both a transaction ID and status (e.g. `https://example.com/callback?gr4vy_transaction_id=123 &gr4vy_transaction_status=capture_succeeded`) after 3-D Secure has completed. For those cases, if the value is not present, the transaction will be marked as failed.
	RedirectUrl NullableString `json:"redirect_url,omitempty"`
}

// NewTransactionCheckoutSessionRequest instantiates a new TransactionCheckoutSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionCheckoutSessionRequest(method string, id string) *TransactionCheckoutSessionRequest {
	this := TransactionCheckoutSessionRequest{}
	this.Method = method
	this.Id = id
	return &this
}

// NewTransactionCheckoutSessionRequestWithDefaults instantiates a new TransactionCheckoutSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionCheckoutSessionRequestWithDefaults() *TransactionCheckoutSessionRequest {
	this := TransactionCheckoutSessionRequest{}
	return &this
}

// GetMethod returns the Method field value
func (o *TransactionCheckoutSessionRequest) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *TransactionCheckoutSessionRequest) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *TransactionCheckoutSessionRequest) SetMethod(v string) {
	o.Method = v
}

// GetId returns the Id field value
func (o *TransactionCheckoutSessionRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransactionCheckoutSessionRequest) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransactionCheckoutSessionRequest) SetId(v string) {
	o.Id = v
}

// GetExternalIdentifier returns the ExternalIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionCheckoutSessionRequest) GetExternalIdentifier() string {
	if o == nil || o.ExternalIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalIdentifier.Get()
}

// GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionCheckoutSessionRequest) GetExternalIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalIdentifier.Get(), o.ExternalIdentifier.IsSet()
}

// HasExternalIdentifier returns a boolean if a field has been set.
func (o *TransactionCheckoutSessionRequest) HasExternalIdentifier() bool {
	if o != nil && o.ExternalIdentifier.IsSet() {
		return true
	}

	return false
}

// SetExternalIdentifier gets a reference to the given NullableString and assigns it to the ExternalIdentifier field.
func (o *TransactionCheckoutSessionRequest) SetExternalIdentifier(v string) {
	o.ExternalIdentifier.Set(&v)
}
// SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil
func (o *TransactionCheckoutSessionRequest) SetExternalIdentifierNil() {
	o.ExternalIdentifier.Set(nil)
}

// UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
func (o *TransactionCheckoutSessionRequest) UnsetExternalIdentifier() {
	o.ExternalIdentifier.Unset()
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionCheckoutSessionRequest) GetRedirectUrl() string {
	if o == nil || o.RedirectUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.RedirectUrl.Get()
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionCheckoutSessionRequest) GetRedirectUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RedirectUrl.Get(), o.RedirectUrl.IsSet()
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *TransactionCheckoutSessionRequest) HasRedirectUrl() bool {
	if o != nil && o.RedirectUrl.IsSet() {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given NullableString and assigns it to the RedirectUrl field.
func (o *TransactionCheckoutSessionRequest) SetRedirectUrl(v string) {
	o.RedirectUrl.Set(&v)
}
// SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil
func (o *TransactionCheckoutSessionRequest) SetRedirectUrlNil() {
	o.RedirectUrl.Set(nil)
}

// UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
func (o *TransactionCheckoutSessionRequest) UnsetRedirectUrl() {
	o.RedirectUrl.Unset()
}

func (o TransactionCheckoutSessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.ExternalIdentifier.IsSet() {
		toSerialize["external_identifier"] = o.ExternalIdentifier.Get()
	}
	if o.RedirectUrl.IsSet() {
		toSerialize["redirect_url"] = o.RedirectUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionCheckoutSessionRequest struct {
	value *TransactionCheckoutSessionRequest
	isSet bool
}

func (v NullableTransactionCheckoutSessionRequest) Get() *TransactionCheckoutSessionRequest {
	return v.value
}

func (v *NullableTransactionCheckoutSessionRequest) Set(val *TransactionCheckoutSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionCheckoutSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionCheckoutSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionCheckoutSessionRequest(val *TransactionCheckoutSessionRequest) *NullableTransactionCheckoutSessionRequest {
	return &NullableTransactionCheckoutSessionRequest{value: val, isSet: true}
}

func (v NullableTransactionCheckoutSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionCheckoutSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


