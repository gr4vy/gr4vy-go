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

// CardRequest Card details to use in a transaction or to register a new payment method.
type CardRequest struct {
	// `card`.
	Method string `json:"method"`
	// The 15-16 digit number for this card as it can be found on the front of the card.
	Number string `json:"number"`
	// The expiration date of the card, formatted `MM/YY`.
	ExpirationDate string `json:"expiration_date"`
	// The 3 or 4 digit security code often found on the card. This often referred to as the CVV or CVD.
	SecurityCode string `json:"security_code"`
	// An external identifier that can be used to match the card against your own records.
	ExternalIdentifier NullableString `json:"external_identifier,omitempty"`
	// The ID of the buyer to associate this payment method to. If this field is provided then the `buyer_external_identifier` field needs to be unset.
	BuyerId *string `json:"buyer_id,omitempty"`
	// The `external_identifier` of the buyer to associate this payment method to. If this field is provided then the `buyer_id` field needs to be unset.
	BuyerExternalIdentifier *string `json:"buyer_external_identifier,omitempty"`
}

// NewCardRequest instantiates a new CardRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardRequest(method string, number string, expirationDate string, securityCode string) *CardRequest {
	this := CardRequest{}
	this.Method = method
	this.Number = number
	this.ExpirationDate = expirationDate
	this.SecurityCode = securityCode
	return &this
}

// NewCardRequestWithDefaults instantiates a new CardRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardRequestWithDefaults() *CardRequest {
	this := CardRequest{}
	return &this
}

// GetMethod returns the Method field value
func (o *CardRequest) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *CardRequest) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *CardRequest) SetMethod(v string) {
	o.Method = v
}

// GetNumber returns the Number field value
func (o *CardRequest) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *CardRequest) GetNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *CardRequest) SetNumber(v string) {
	o.Number = v
}

// GetExpirationDate returns the ExpirationDate field value
func (o *CardRequest) GetExpirationDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value
// and a boolean to check if the value has been set.
func (o *CardRequest) GetExpirationDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpirationDate, true
}

// SetExpirationDate sets field value
func (o *CardRequest) SetExpirationDate(v string) {
	o.ExpirationDate = v
}

// GetSecurityCode returns the SecurityCode field value
func (o *CardRequest) GetSecurityCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecurityCode
}

// GetSecurityCodeOk returns a tuple with the SecurityCode field value
// and a boolean to check if the value has been set.
func (o *CardRequest) GetSecurityCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecurityCode, true
}

// SetSecurityCode sets field value
func (o *CardRequest) SetSecurityCode(v string) {
	o.SecurityCode = v
}

// GetExternalIdentifier returns the ExternalIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardRequest) GetExternalIdentifier() string {
	if o == nil || o.ExternalIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalIdentifier.Get()
}

// GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardRequest) GetExternalIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalIdentifier.Get(), o.ExternalIdentifier.IsSet()
}

// HasExternalIdentifier returns a boolean if a field has been set.
func (o *CardRequest) HasExternalIdentifier() bool {
	if o != nil && o.ExternalIdentifier.IsSet() {
		return true
	}

	return false
}

// SetExternalIdentifier gets a reference to the given NullableString and assigns it to the ExternalIdentifier field.
func (o *CardRequest) SetExternalIdentifier(v string) {
	o.ExternalIdentifier.Set(&v)
}
// SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil
func (o *CardRequest) SetExternalIdentifierNil() {
	o.ExternalIdentifier.Set(nil)
}

// UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
func (o *CardRequest) UnsetExternalIdentifier() {
	o.ExternalIdentifier.Unset()
}

// GetBuyerId returns the BuyerId field value if set, zero value otherwise.
func (o *CardRequest) GetBuyerId() string {
	if o == nil || o.BuyerId == nil {
		var ret string
		return ret
	}
	return *o.BuyerId
}

// GetBuyerIdOk returns a tuple with the BuyerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardRequest) GetBuyerIdOk() (*string, bool) {
	if o == nil || o.BuyerId == nil {
		return nil, false
	}
	return o.BuyerId, true
}

// HasBuyerId returns a boolean if a field has been set.
func (o *CardRequest) HasBuyerId() bool {
	if o != nil && o.BuyerId != nil {
		return true
	}

	return false
}

// SetBuyerId gets a reference to the given string and assigns it to the BuyerId field.
func (o *CardRequest) SetBuyerId(v string) {
	o.BuyerId = &v
}

// GetBuyerExternalIdentifier returns the BuyerExternalIdentifier field value if set, zero value otherwise.
func (o *CardRequest) GetBuyerExternalIdentifier() string {
	if o == nil || o.BuyerExternalIdentifier == nil {
		var ret string
		return ret
	}
	return *o.BuyerExternalIdentifier
}

// GetBuyerExternalIdentifierOk returns a tuple with the BuyerExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardRequest) GetBuyerExternalIdentifierOk() (*string, bool) {
	if o == nil || o.BuyerExternalIdentifier == nil {
		return nil, false
	}
	return o.BuyerExternalIdentifier, true
}

// HasBuyerExternalIdentifier returns a boolean if a field has been set.
func (o *CardRequest) HasBuyerExternalIdentifier() bool {
	if o != nil && o.BuyerExternalIdentifier != nil {
		return true
	}

	return false
}

// SetBuyerExternalIdentifier gets a reference to the given string and assigns it to the BuyerExternalIdentifier field.
func (o *CardRequest) SetBuyerExternalIdentifier(v string) {
	o.BuyerExternalIdentifier = &v
}

func (o CardRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["number"] = o.Number
	}
	if true {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	if true {
		toSerialize["security_code"] = o.SecurityCode
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

type NullableCardRequest struct {
	value *CardRequest
	isSet bool
}

func (v NullableCardRequest) Get() *CardRequest {
	return v.value
}

func (v *NullableCardRequest) Set(val *CardRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCardRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCardRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardRequest(val *CardRequest) *NullableCardRequest {
	return &NullableCardRequest{value: val, isSet: true}
}

func (v NullableCardRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


