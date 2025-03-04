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

// TransactionNetworkTokenApplePayRequest Details for an Apple Pay decrypted token payment method.
type TransactionNetworkTokenApplePayRequest struct {
	// `network-token`.
	Method string `json:"method"`
	// The value of the decrypted Apple Pay token.
	Token string `json:"token"`
	// The expiration date of the network token, formatted `MM/YY`.
	ExpirationDate string `json:"expiration_date"`
	// The cryptogram of the network token.
	Cryptogram NullableString `json:"cryptogram,omitempty"`
	// The ecommerce indicator for 3D-Secure.
	Eci *string `json:"eci,omitempty"`
	CardSource string `json:"card_source"`
	// Last four digits of card PAN.
	CardSuffix NullableString `json:"card_suffix,omitempty"`
	// The scheme/brand of the card.
	CardScheme NullableString `json:"card_scheme,omitempty"`
	// The cardholder name.
	CardholderName NullableString `json:"cardholder_name,omitempty"`
}

// NewTransactionNetworkTokenApplePayRequest instantiates a new TransactionNetworkTokenApplePayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionNetworkTokenApplePayRequest(method string, token string, expirationDate string, cardSource string) *TransactionNetworkTokenApplePayRequest {
	this := TransactionNetworkTokenApplePayRequest{}
	this.Method = method
	this.Token = token
	this.ExpirationDate = expirationDate
	this.CardSource = cardSource
	return &this
}

// NewTransactionNetworkTokenApplePayRequestWithDefaults instantiates a new TransactionNetworkTokenApplePayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionNetworkTokenApplePayRequestWithDefaults() *TransactionNetworkTokenApplePayRequest {
	this := TransactionNetworkTokenApplePayRequest{}
	return &this
}

// GetMethod returns the Method field value
func (o *TransactionNetworkTokenApplePayRequest) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *TransactionNetworkTokenApplePayRequest) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *TransactionNetworkTokenApplePayRequest) SetMethod(v string) {
	o.Method = v
}

// GetToken returns the Token field value
func (o *TransactionNetworkTokenApplePayRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *TransactionNetworkTokenApplePayRequest) GetTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *TransactionNetworkTokenApplePayRequest) SetToken(v string) {
	o.Token = v
}

// GetExpirationDate returns the ExpirationDate field value
func (o *TransactionNetworkTokenApplePayRequest) GetExpirationDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value
// and a boolean to check if the value has been set.
func (o *TransactionNetworkTokenApplePayRequest) GetExpirationDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpirationDate, true
}

// SetExpirationDate sets field value
func (o *TransactionNetworkTokenApplePayRequest) SetExpirationDate(v string) {
	o.ExpirationDate = v
}

// GetCryptogram returns the Cryptogram field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionNetworkTokenApplePayRequest) GetCryptogram() string {
	if o == nil || o.Cryptogram.Get() == nil {
		var ret string
		return ret
	}
	return *o.Cryptogram.Get()
}

// GetCryptogramOk returns a tuple with the Cryptogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionNetworkTokenApplePayRequest) GetCryptogramOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Cryptogram.Get(), o.Cryptogram.IsSet()
}

// HasCryptogram returns a boolean if a field has been set.
func (o *TransactionNetworkTokenApplePayRequest) HasCryptogram() bool {
	if o != nil && o.Cryptogram.IsSet() {
		return true
	}

	return false
}

// SetCryptogram gets a reference to the given NullableString and assigns it to the Cryptogram field.
func (o *TransactionNetworkTokenApplePayRequest) SetCryptogram(v string) {
	o.Cryptogram.Set(&v)
}
// SetCryptogramNil sets the value for Cryptogram to be an explicit nil
func (o *TransactionNetworkTokenApplePayRequest) SetCryptogramNil() {
	o.Cryptogram.Set(nil)
}

// UnsetCryptogram ensures that no value is present for Cryptogram, not even an explicit nil
func (o *TransactionNetworkTokenApplePayRequest) UnsetCryptogram() {
	o.Cryptogram.Unset()
}

// GetEci returns the Eci field value if set, zero value otherwise.
func (o *TransactionNetworkTokenApplePayRequest) GetEci() string {
	if o == nil || o.Eci == nil {
		var ret string
		return ret
	}
	return *o.Eci
}

// GetEciOk returns a tuple with the Eci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionNetworkTokenApplePayRequest) GetEciOk() (*string, bool) {
	if o == nil || o.Eci == nil {
		return nil, false
	}
	return o.Eci, true
}

// HasEci returns a boolean if a field has been set.
func (o *TransactionNetworkTokenApplePayRequest) HasEci() bool {
	if o != nil && o.Eci != nil {
		return true
	}

	return false
}

// SetEci gets a reference to the given string and assigns it to the Eci field.
func (o *TransactionNetworkTokenApplePayRequest) SetEci(v string) {
	o.Eci = &v
}

// GetCardSource returns the CardSource field value
func (o *TransactionNetworkTokenApplePayRequest) GetCardSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardSource
}

// GetCardSourceOk returns a tuple with the CardSource field value
// and a boolean to check if the value has been set.
func (o *TransactionNetworkTokenApplePayRequest) GetCardSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CardSource, true
}

// SetCardSource sets field value
func (o *TransactionNetworkTokenApplePayRequest) SetCardSource(v string) {
	o.CardSource = v
}

// GetCardSuffix returns the CardSuffix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionNetworkTokenApplePayRequest) GetCardSuffix() string {
	if o == nil || o.CardSuffix.Get() == nil {
		var ret string
		return ret
	}
	return *o.CardSuffix.Get()
}

// GetCardSuffixOk returns a tuple with the CardSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionNetworkTokenApplePayRequest) GetCardSuffixOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CardSuffix.Get(), o.CardSuffix.IsSet()
}

// HasCardSuffix returns a boolean if a field has been set.
func (o *TransactionNetworkTokenApplePayRequest) HasCardSuffix() bool {
	if o != nil && o.CardSuffix.IsSet() {
		return true
	}

	return false
}

// SetCardSuffix gets a reference to the given NullableString and assigns it to the CardSuffix field.
func (o *TransactionNetworkTokenApplePayRequest) SetCardSuffix(v string) {
	o.CardSuffix.Set(&v)
}
// SetCardSuffixNil sets the value for CardSuffix to be an explicit nil
func (o *TransactionNetworkTokenApplePayRequest) SetCardSuffixNil() {
	o.CardSuffix.Set(nil)
}

// UnsetCardSuffix ensures that no value is present for CardSuffix, not even an explicit nil
func (o *TransactionNetworkTokenApplePayRequest) UnsetCardSuffix() {
	o.CardSuffix.Unset()
}

// GetCardScheme returns the CardScheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionNetworkTokenApplePayRequest) GetCardScheme() string {
	if o == nil || o.CardScheme.Get() == nil {
		var ret string
		return ret
	}
	return *o.CardScheme.Get()
}

// GetCardSchemeOk returns a tuple with the CardScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionNetworkTokenApplePayRequest) GetCardSchemeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CardScheme.Get(), o.CardScheme.IsSet()
}

// HasCardScheme returns a boolean if a field has been set.
func (o *TransactionNetworkTokenApplePayRequest) HasCardScheme() bool {
	if o != nil && o.CardScheme.IsSet() {
		return true
	}

	return false
}

// SetCardScheme gets a reference to the given NullableString and assigns it to the CardScheme field.
func (o *TransactionNetworkTokenApplePayRequest) SetCardScheme(v string) {
	o.CardScheme.Set(&v)
}
// SetCardSchemeNil sets the value for CardScheme to be an explicit nil
func (o *TransactionNetworkTokenApplePayRequest) SetCardSchemeNil() {
	o.CardScheme.Set(nil)
}

// UnsetCardScheme ensures that no value is present for CardScheme, not even an explicit nil
func (o *TransactionNetworkTokenApplePayRequest) UnsetCardScheme() {
	o.CardScheme.Unset()
}

// GetCardholderName returns the CardholderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionNetworkTokenApplePayRequest) GetCardholderName() string {
	if o == nil || o.CardholderName.Get() == nil {
		var ret string
		return ret
	}
	return *o.CardholderName.Get()
}

// GetCardholderNameOk returns a tuple with the CardholderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionNetworkTokenApplePayRequest) GetCardholderNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CardholderName.Get(), o.CardholderName.IsSet()
}

// HasCardholderName returns a boolean if a field has been set.
func (o *TransactionNetworkTokenApplePayRequest) HasCardholderName() bool {
	if o != nil && o.CardholderName.IsSet() {
		return true
	}

	return false
}

// SetCardholderName gets a reference to the given NullableString and assigns it to the CardholderName field.
func (o *TransactionNetworkTokenApplePayRequest) SetCardholderName(v string) {
	o.CardholderName.Set(&v)
}
// SetCardholderNameNil sets the value for CardholderName to be an explicit nil
func (o *TransactionNetworkTokenApplePayRequest) SetCardholderNameNil() {
	o.CardholderName.Set(nil)
}

// UnsetCardholderName ensures that no value is present for CardholderName, not even an explicit nil
func (o *TransactionNetworkTokenApplePayRequest) UnsetCardholderName() {
	o.CardholderName.Unset()
}

func (o TransactionNetworkTokenApplePayRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	if o.Cryptogram.IsSet() {
		toSerialize["cryptogram"] = o.Cryptogram.Get()
	}
	if o.Eci != nil {
		toSerialize["eci"] = o.Eci
	}
	if true {
		toSerialize["card_source"] = o.CardSource
	}
	if o.CardSuffix.IsSet() {
		toSerialize["card_suffix"] = o.CardSuffix.Get()
	}
	if o.CardScheme.IsSet() {
		toSerialize["card_scheme"] = o.CardScheme.Get()
	}
	if o.CardholderName.IsSet() {
		toSerialize["cardholder_name"] = o.CardholderName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionNetworkTokenApplePayRequest struct {
	value *TransactionNetworkTokenApplePayRequest
	isSet bool
}

func (v NullableTransactionNetworkTokenApplePayRequest) Get() *TransactionNetworkTokenApplePayRequest {
	return v.value
}

func (v *NullableTransactionNetworkTokenApplePayRequest) Set(val *TransactionNetworkTokenApplePayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionNetworkTokenApplePayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionNetworkTokenApplePayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionNetworkTokenApplePayRequest(val *TransactionNetworkTokenApplePayRequest) *NullableTransactionNetworkTokenApplePayRequest {
	return &NullableTransactionNetworkTokenApplePayRequest{value: val, isSet: true}
}

func (v NullableTransactionNetworkTokenApplePayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionNetworkTokenApplePayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


