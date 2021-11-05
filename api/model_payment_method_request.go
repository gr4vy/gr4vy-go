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

// PaymentMethodRequest Payment method details used to register a new payment method.
type PaymentMethodRequest struct {
	Method string `json:"method"`
	// The 15-16 digit number for this credit card as it can be found on the front of the card.  If a card has been stored with us previously, this number will represent the unique tokenized card ID provided via our API.
	Number *string `json:"number,omitempty"`
	// The expiration date of the card, formatted `MM/YY`. If a card has been previously stored with us this value is optional.  If the `number` of this card represents a tokenized card, then this value is ignored.
	ExpirationDate *string `json:"expiration_date,omitempty"`
	// The 3 or 4 digit security code often found on the card. This often referred to as the CVV or CVD.  If the `number` of this card represents a tokenized card, then this value is ignored.
	SecurityCode *string `json:"security_code,omitempty"`
	// An external identifier that can be used to match the card against your own records.
	ExternalIdentifier NullableString `json:"external_identifier,omitempty"`
	// The ID of the buyer to associate this payment method to. If this field is provided then the `buyer_external_identifier` field needs to be unset.
	BuyerId *string `json:"buyer_id,omitempty"`
	// The `external_identifier` of the buyer to associate this payment method to. If this field is provided then the `buyer_id` field needs to be unset.
	BuyerExternalIdentifier *string `json:"buyer_external_identifier,omitempty"`
	// The redirect URL to redirect a buyer to after they have authorized their transaction or payment method. This only applies to payment methods that require buyer approval.
	RedirectUrl *string `json:"redirect_url,omitempty"`
	// The ISO-4217 currency code to store this payment method for. This is used to select the payment service to use.  This only applies to `redirect` mode payment methods like `gocardless`.
	Currency *string `json:"currency,omitempty"`
	// The 2-letter ISO code of the country to store this payment method for. This is used to select the payment service to use.  This only applies to `redirect` mode payment methods like `gocardless`.
	Country *string `json:"country,omitempty"`
	// Defines the environment to store this payment method in. Setting this to anything other than `production` will force Gr4vy to use a payment a service configured for that environment.
	Environment *string `json:"environment,omitempty"`
}

// NewPaymentMethodRequest instantiates a new PaymentMethodRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodRequest(method string) *PaymentMethodRequest {
	this := PaymentMethodRequest{}
	this.Method = method
	return &this
}

// NewPaymentMethodRequestWithDefaults instantiates a new PaymentMethodRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodRequestWithDefaults() *PaymentMethodRequest {
	this := PaymentMethodRequest{}
	return &this
}

// GetMethod returns the Method field value
func (o *PaymentMethodRequest) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodRequest) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *PaymentMethodRequest) SetMethod(v string) {
	o.Method = v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *PaymentMethodRequest) GetNumber() string {
	if o == nil || o.Number == nil {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodRequest) GetNumberOk() (*string, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *PaymentMethodRequest) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *PaymentMethodRequest) SetNumber(v string) {
	o.Number = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *PaymentMethodRequest) GetExpirationDate() string {
	if o == nil || o.ExpirationDate == nil {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodRequest) GetExpirationDateOk() (*string, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *PaymentMethodRequest) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate != nil {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *PaymentMethodRequest) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetSecurityCode returns the SecurityCode field value if set, zero value otherwise.
func (o *PaymentMethodRequest) GetSecurityCode() string {
	if o == nil || o.SecurityCode == nil {
		var ret string
		return ret
	}
	return *o.SecurityCode
}

// GetSecurityCodeOk returns a tuple with the SecurityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodRequest) GetSecurityCodeOk() (*string, bool) {
	if o == nil || o.SecurityCode == nil {
		return nil, false
	}
	return o.SecurityCode, true
}

// HasSecurityCode returns a boolean if a field has been set.
func (o *PaymentMethodRequest) HasSecurityCode() bool {
	if o != nil && o.SecurityCode != nil {
		return true
	}

	return false
}

// SetSecurityCode gets a reference to the given string and assigns it to the SecurityCode field.
func (o *PaymentMethodRequest) SetSecurityCode(v string) {
	o.SecurityCode = &v
}

// GetExternalIdentifier returns the ExternalIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodRequest) GetExternalIdentifier() string {
	if o == nil || o.ExternalIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalIdentifier.Get()
}

// GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodRequest) GetExternalIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalIdentifier.Get(), o.ExternalIdentifier.IsSet()
}

// HasExternalIdentifier returns a boolean if a field has been set.
func (o *PaymentMethodRequest) HasExternalIdentifier() bool {
	if o != nil && o.ExternalIdentifier.IsSet() {
		return true
	}

	return false
}

// SetExternalIdentifier gets a reference to the given NullableString and assigns it to the ExternalIdentifier field.
func (o *PaymentMethodRequest) SetExternalIdentifier(v string) {
	o.ExternalIdentifier.Set(&v)
}
// SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil
func (o *PaymentMethodRequest) SetExternalIdentifierNil() {
	o.ExternalIdentifier.Set(nil)
}

// UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
func (o *PaymentMethodRequest) UnsetExternalIdentifier() {
	o.ExternalIdentifier.Unset()
}

// GetBuyerId returns the BuyerId field value if set, zero value otherwise.
func (o *PaymentMethodRequest) GetBuyerId() string {
	if o == nil || o.BuyerId == nil {
		var ret string
		return ret
	}
	return *o.BuyerId
}

// GetBuyerIdOk returns a tuple with the BuyerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodRequest) GetBuyerIdOk() (*string, bool) {
	if o == nil || o.BuyerId == nil {
		return nil, false
	}
	return o.BuyerId, true
}

// HasBuyerId returns a boolean if a field has been set.
func (o *PaymentMethodRequest) HasBuyerId() bool {
	if o != nil && o.BuyerId != nil {
		return true
	}

	return false
}

// SetBuyerId gets a reference to the given string and assigns it to the BuyerId field.
func (o *PaymentMethodRequest) SetBuyerId(v string) {
	o.BuyerId = &v
}

// GetBuyerExternalIdentifier returns the BuyerExternalIdentifier field value if set, zero value otherwise.
func (o *PaymentMethodRequest) GetBuyerExternalIdentifier() string {
	if o == nil || o.BuyerExternalIdentifier == nil {
		var ret string
		return ret
	}
	return *o.BuyerExternalIdentifier
}

// GetBuyerExternalIdentifierOk returns a tuple with the BuyerExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodRequest) GetBuyerExternalIdentifierOk() (*string, bool) {
	if o == nil || o.BuyerExternalIdentifier == nil {
		return nil, false
	}
	return o.BuyerExternalIdentifier, true
}

// HasBuyerExternalIdentifier returns a boolean if a field has been set.
func (o *PaymentMethodRequest) HasBuyerExternalIdentifier() bool {
	if o != nil && o.BuyerExternalIdentifier != nil {
		return true
	}

	return false
}

// SetBuyerExternalIdentifier gets a reference to the given string and assigns it to the BuyerExternalIdentifier field.
func (o *PaymentMethodRequest) SetBuyerExternalIdentifier(v string) {
	o.BuyerExternalIdentifier = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *PaymentMethodRequest) GetRedirectUrl() string {
	if o == nil || o.RedirectUrl == nil {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodRequest) GetRedirectUrlOk() (*string, bool) {
	if o == nil || o.RedirectUrl == nil {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *PaymentMethodRequest) HasRedirectUrl() bool {
	if o != nil && o.RedirectUrl != nil {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *PaymentMethodRequest) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PaymentMethodRequest) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodRequest) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PaymentMethodRequest) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *PaymentMethodRequest) SetCurrency(v string) {
	o.Currency = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *PaymentMethodRequest) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodRequest) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *PaymentMethodRequest) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *PaymentMethodRequest) SetCountry(v string) {
	o.Country = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *PaymentMethodRequest) GetEnvironment() string {
	if o == nil || o.Environment == nil {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *PaymentMethodRequest) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *PaymentMethodRequest) SetEnvironment(v string) {
	o.Environment = &v
}

func (o PaymentMethodRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["method"] = o.Method
	}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.ExpirationDate != nil {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	if o.SecurityCode != nil {
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
	if o.RedirectUrl != nil {
		toSerialize["redirect_url"] = o.RedirectUrl
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentMethodRequest struct {
	value *PaymentMethodRequest
	isSet bool
}

func (v NullablePaymentMethodRequest) Get() *PaymentMethodRequest {
	return v.value
}

func (v *NullablePaymentMethodRequest) Set(val *PaymentMethodRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodRequest(val *PaymentMethodRequest) *NullablePaymentMethodRequest {
	return &NullablePaymentMethodRequest{value: val, isSet: true}
}

func (v NullablePaymentMethodRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


