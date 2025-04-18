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

// CheckoutSessionRequest Details to register a new Checkout Session payment method.
type CheckoutSessionRequest struct {
	// `checkout-session`.
	Method string `json:"method"`
	// The ID of the Checkout Session.
	Id string `json:"id"`
	// The redirect URL to redirect a buyer to after they have authorized their transaction or payment method. This only applies to payment methods that require buyer approval.
	RedirectUrl NullableString `json:"redirect_url,omitempty"`
	// An external identifier that can be used to match the card against your own records.
	ExternalIdentifier NullableString `json:"external_identifier,omitempty"`
	// The ID of the buyer to associate this payment method to. If this field is provided then the `buyer_external_identifier` field needs to be unset.
	BuyerId *string `json:"buyer_id,omitempty"`
	// The `external_identifier` of the buyer to associate this payment method to. If this field is provided then the `buyer_id` field needs to be unset.
	BuyerExternalIdentifier *string `json:"buyer_external_identifier,omitempty"`
}

// NewCheckoutSessionRequest instantiates a new CheckoutSessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutSessionRequest(method string, id string) *CheckoutSessionRequest {
	this := CheckoutSessionRequest{}
	this.Method = method
	this.Id = id
	return &this
}

// NewCheckoutSessionRequestWithDefaults instantiates a new CheckoutSessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutSessionRequestWithDefaults() *CheckoutSessionRequest {
	this := CheckoutSessionRequest{}
	return &this
}

// GetMethod returns the Method field value
func (o *CheckoutSessionRequest) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *CheckoutSessionRequest) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *CheckoutSessionRequest) SetMethod(v string) {
	o.Method = v
}

// GetId returns the Id field value
func (o *CheckoutSessionRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CheckoutSessionRequest) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CheckoutSessionRequest) SetId(v string) {
	o.Id = v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckoutSessionRequest) GetRedirectUrl() string {
	if o == nil || o.RedirectUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.RedirectUrl.Get()
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSessionRequest) GetRedirectUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RedirectUrl.Get(), o.RedirectUrl.IsSet()
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *CheckoutSessionRequest) HasRedirectUrl() bool {
	if o != nil && o.RedirectUrl.IsSet() {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given NullableString and assigns it to the RedirectUrl field.
func (o *CheckoutSessionRequest) SetRedirectUrl(v string) {
	o.RedirectUrl.Set(&v)
}
// SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil
func (o *CheckoutSessionRequest) SetRedirectUrlNil() {
	o.RedirectUrl.Set(nil)
}

// UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
func (o *CheckoutSessionRequest) UnsetRedirectUrl() {
	o.RedirectUrl.Unset()
}

// GetExternalIdentifier returns the ExternalIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckoutSessionRequest) GetExternalIdentifier() string {
	if o == nil || o.ExternalIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalIdentifier.Get()
}

// GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSessionRequest) GetExternalIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalIdentifier.Get(), o.ExternalIdentifier.IsSet()
}

// HasExternalIdentifier returns a boolean if a field has been set.
func (o *CheckoutSessionRequest) HasExternalIdentifier() bool {
	if o != nil && o.ExternalIdentifier.IsSet() {
		return true
	}

	return false
}

// SetExternalIdentifier gets a reference to the given NullableString and assigns it to the ExternalIdentifier field.
func (o *CheckoutSessionRequest) SetExternalIdentifier(v string) {
	o.ExternalIdentifier.Set(&v)
}
// SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil
func (o *CheckoutSessionRequest) SetExternalIdentifierNil() {
	o.ExternalIdentifier.Set(nil)
}

// UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
func (o *CheckoutSessionRequest) UnsetExternalIdentifier() {
	o.ExternalIdentifier.Unset()
}

// GetBuyerId returns the BuyerId field value if set, zero value otherwise.
func (o *CheckoutSessionRequest) GetBuyerId() string {
	if o == nil || o.BuyerId == nil {
		var ret string
		return ret
	}
	return *o.BuyerId
}

// GetBuyerIdOk returns a tuple with the BuyerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionRequest) GetBuyerIdOk() (*string, bool) {
	if o == nil || o.BuyerId == nil {
		return nil, false
	}
	return o.BuyerId, true
}

// HasBuyerId returns a boolean if a field has been set.
func (o *CheckoutSessionRequest) HasBuyerId() bool {
	if o != nil && o.BuyerId != nil {
		return true
	}

	return false
}

// SetBuyerId gets a reference to the given string and assigns it to the BuyerId field.
func (o *CheckoutSessionRequest) SetBuyerId(v string) {
	o.BuyerId = &v
}

// GetBuyerExternalIdentifier returns the BuyerExternalIdentifier field value if set, zero value otherwise.
func (o *CheckoutSessionRequest) GetBuyerExternalIdentifier() string {
	if o == nil || o.BuyerExternalIdentifier == nil {
		var ret string
		return ret
	}
	return *o.BuyerExternalIdentifier
}

// GetBuyerExternalIdentifierOk returns a tuple with the BuyerExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionRequest) GetBuyerExternalIdentifierOk() (*string, bool) {
	if o == nil || o.BuyerExternalIdentifier == nil {
		return nil, false
	}
	return o.BuyerExternalIdentifier, true
}

// HasBuyerExternalIdentifier returns a boolean if a field has been set.
func (o *CheckoutSessionRequest) HasBuyerExternalIdentifier() bool {
	if o != nil && o.BuyerExternalIdentifier != nil {
		return true
	}

	return false
}

// SetBuyerExternalIdentifier gets a reference to the given string and assigns it to the BuyerExternalIdentifier field.
func (o *CheckoutSessionRequest) SetBuyerExternalIdentifier(v string) {
	o.BuyerExternalIdentifier = &v
}

func (o CheckoutSessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.RedirectUrl.IsSet() {
		toSerialize["redirect_url"] = o.RedirectUrl.Get()
	}
	if o.ExternalIdentifier.IsSet() {
		toSerialize["external_identifier"] = o.ExternalIdentifier.Get()
	}
	if o.BuyerId != nil {
		toSerialize["buyer_id"] = o.BuyerId
	}
	if o.BuyerExternalIdentifier != nil {
		toSerialize["buyer_external_identifier"] = o.BuyerExternalIdentifier
	}
	return json.Marshal(toSerialize)
}

type NullableCheckoutSessionRequest struct {
	value *CheckoutSessionRequest
	isSet bool
}

func (v NullableCheckoutSessionRequest) Get() *CheckoutSessionRequest {
	return v.value
}

func (v *NullableCheckoutSessionRequest) Set(val *CheckoutSessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionRequest(val *CheckoutSessionRequest) *NullableCheckoutSessionRequest {
	return &NullableCheckoutSessionRequest{value: val, isSet: true}
}

func (v NullableCheckoutSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


