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

// TransactionRedirectRequest Redirect payment method details to use in a transaction.
type TransactionRedirectRequest struct {
	// The method to use, this can be any of the methods that support redirect requests.  When storing a new payment method, only `gocardless` and `stripedd` are currently supported.
	Method string `json:"method"`
	// The redirect URL to redirect a buyer to after they have authorized their transaction.
	RedirectUrl string `json:"redirect_url"`
	// The ISO-4217 currency code to use this payment method for. This is used to select the payment service to use.
	Currency string `json:"currency"`
	// The 2-letter ISO code of the country to use this payment method for. This is used to select the payment service to use.
	Country string `json:"country"`
	// An external identifier that can be used to match the account against your own records. This can only be set if the `store` flag is set to `true`.
	ExternalIdentifier NullableString `json:"external_identifier,omitempty"`
}

// NewTransactionRedirectRequest instantiates a new TransactionRedirectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRedirectRequest(method string, redirectUrl string, currency string, country string) *TransactionRedirectRequest {
	this := TransactionRedirectRequest{}
	this.Method = method
	this.RedirectUrl = redirectUrl
	this.Currency = currency
	this.Country = country
	return &this
}

// NewTransactionRedirectRequestWithDefaults instantiates a new TransactionRedirectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRedirectRequestWithDefaults() *TransactionRedirectRequest {
	this := TransactionRedirectRequest{}
	return &this
}

// GetMethod returns the Method field value
func (o *TransactionRedirectRequest) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *TransactionRedirectRequest) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *TransactionRedirectRequest) SetMethod(v string) {
	o.Method = v
}

// GetRedirectUrl returns the RedirectUrl field value
func (o *TransactionRedirectRequest) GetRedirectUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value
// and a boolean to check if the value has been set.
func (o *TransactionRedirectRequest) GetRedirectUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RedirectUrl, true
}

// SetRedirectUrl sets field value
func (o *TransactionRedirectRequest) SetRedirectUrl(v string) {
	o.RedirectUrl = v
}

// GetCurrency returns the Currency field value
func (o *TransactionRedirectRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TransactionRedirectRequest) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TransactionRedirectRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetCountry returns the Country field value
func (o *TransactionRedirectRequest) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *TransactionRedirectRequest) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *TransactionRedirectRequest) SetCountry(v string) {
	o.Country = v
}

// GetExternalIdentifier returns the ExternalIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionRedirectRequest) GetExternalIdentifier() string {
	if o == nil || o.ExternalIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalIdentifier.Get()
}

// GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionRedirectRequest) GetExternalIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalIdentifier.Get(), o.ExternalIdentifier.IsSet()
}

// HasExternalIdentifier returns a boolean if a field has been set.
func (o *TransactionRedirectRequest) HasExternalIdentifier() bool {
	if o != nil && o.ExternalIdentifier.IsSet() {
		return true
	}

	return false
}

// SetExternalIdentifier gets a reference to the given NullableString and assigns it to the ExternalIdentifier field.
func (o *TransactionRedirectRequest) SetExternalIdentifier(v string) {
	o.ExternalIdentifier.Set(&v)
}
// SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil
func (o *TransactionRedirectRequest) SetExternalIdentifierNil() {
	o.ExternalIdentifier.Set(nil)
}

// UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
func (o *TransactionRedirectRequest) UnsetExternalIdentifier() {
	o.ExternalIdentifier.Unset()
}

func (o TransactionRedirectRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["redirect_url"] = o.RedirectUrl
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["country"] = o.Country
	}
	if o.ExternalIdentifier.IsSet() {
		toSerialize["external_identifier"] = o.ExternalIdentifier.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionRedirectRequest struct {
	value *TransactionRedirectRequest
	isSet bool
}

func (v NullableTransactionRedirectRequest) Get() *TransactionRedirectRequest {
	return v.value
}

func (v *NullableTransactionRedirectRequest) Set(val *TransactionRedirectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRedirectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRedirectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRedirectRequest(val *TransactionRedirectRequest) *NullableTransactionRedirectRequest {
	return &NullableTransactionRedirectRequest{value: val, isSet: true}
}

func (v NullableTransactionRedirectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRedirectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


