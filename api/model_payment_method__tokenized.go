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

// PaymentMethodTokenized A mini format version of a payment method.
type PaymentMethodTokenized struct {
	// `payment-method`.
	Type *string `json:"type,omitempty"`
	// The unique ID of the payment method.
	Id *string `json:"id,omitempty"`
	Method *string `json:"method,omitempty"`
	// A label for the payment method. For a `card` payment method this is the last 4 digits on the card. For others it would be the email address.
	Label *string `json:"label,omitempty"`
	// The type of the card, if the payment method is a card.
	Scheme NullableString `json:"scheme,omitempty"`
	// The expiration date for the payment method.
	ExpirationDate NullableString `json:"expiration_date,omitempty"`
	// The optional URL that the buyer needs to be redirected to to further authorize their payment.
	ApprovalUrl NullableString `json:"approval_url,omitempty"`
	// The ISO-4217 currency code that this payment method can be used for. If this value is `null` the payment method may be used for multiple currencies.
	Currency NullableString `json:"currency,omitempty"`
	// The 2-letter ISO code of the country this payment method can be used for. If this value is `null` the payment method may be used in multiple countries.
	Country NullableString `json:"country,omitempty"`
}

// NewPaymentMethodTokenized instantiates a new PaymentMethodTokenized object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodTokenized() *PaymentMethodTokenized {
	this := PaymentMethodTokenized{}
	return &this
}

// NewPaymentMethodTokenizedWithDefaults instantiates a new PaymentMethodTokenized object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodTokenizedWithDefaults() *PaymentMethodTokenized {
	this := PaymentMethodTokenized{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentMethodTokenized) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodTokenized) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentMethodTokenized) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentMethodTokenized) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaymentMethodTokenized) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodTokenized) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaymentMethodTokenized) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PaymentMethodTokenized) SetId(v string) {
	o.Id = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *PaymentMethodTokenized) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodTokenized) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *PaymentMethodTokenized) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *PaymentMethodTokenized) SetMethod(v string) {
	o.Method = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PaymentMethodTokenized) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodTokenized) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PaymentMethodTokenized) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PaymentMethodTokenized) SetLabel(v string) {
	o.Label = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodTokenized) GetScheme() string {
	if o == nil || o.Scheme.Get() == nil {
		var ret string
		return ret
	}
	return *o.Scheme.Get()
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodTokenized) GetSchemeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Scheme.Get(), o.Scheme.IsSet()
}

// HasScheme returns a boolean if a field has been set.
func (o *PaymentMethodTokenized) HasScheme() bool {
	if o != nil && o.Scheme.IsSet() {
		return true
	}

	return false
}

// SetScheme gets a reference to the given NullableString and assigns it to the Scheme field.
func (o *PaymentMethodTokenized) SetScheme(v string) {
	o.Scheme.Set(&v)
}
// SetSchemeNil sets the value for Scheme to be an explicit nil
func (o *PaymentMethodTokenized) SetSchemeNil() {
	o.Scheme.Set(nil)
}

// UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
func (o *PaymentMethodTokenized) UnsetScheme() {
	o.Scheme.Unset()
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodTokenized) GetExpirationDate() string {
	if o == nil || o.ExpirationDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExpirationDate.Get()
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodTokenized) GetExpirationDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpirationDate.Get(), o.ExpirationDate.IsSet()
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *PaymentMethodTokenized) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate.IsSet() {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given NullableString and assigns it to the ExpirationDate field.
func (o *PaymentMethodTokenized) SetExpirationDate(v string) {
	o.ExpirationDate.Set(&v)
}
// SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil
func (o *PaymentMethodTokenized) SetExpirationDateNil() {
	o.ExpirationDate.Set(nil)
}

// UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
func (o *PaymentMethodTokenized) UnsetExpirationDate() {
	o.ExpirationDate.Unset()
}

// GetApprovalUrl returns the ApprovalUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodTokenized) GetApprovalUrl() string {
	if o == nil || o.ApprovalUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApprovalUrl.Get()
}

// GetApprovalUrlOk returns a tuple with the ApprovalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodTokenized) GetApprovalUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApprovalUrl.Get(), o.ApprovalUrl.IsSet()
}

// HasApprovalUrl returns a boolean if a field has been set.
func (o *PaymentMethodTokenized) HasApprovalUrl() bool {
	if o != nil && o.ApprovalUrl.IsSet() {
		return true
	}

	return false
}

// SetApprovalUrl gets a reference to the given NullableString and assigns it to the ApprovalUrl field.
func (o *PaymentMethodTokenized) SetApprovalUrl(v string) {
	o.ApprovalUrl.Set(&v)
}
// SetApprovalUrlNil sets the value for ApprovalUrl to be an explicit nil
func (o *PaymentMethodTokenized) SetApprovalUrlNil() {
	o.ApprovalUrl.Set(nil)
}

// UnsetApprovalUrl ensures that no value is present for ApprovalUrl, not even an explicit nil
func (o *PaymentMethodTokenized) UnsetApprovalUrl() {
	o.ApprovalUrl.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodTokenized) GetCurrency() string {
	if o == nil || o.Currency.Get() == nil {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodTokenized) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *PaymentMethodTokenized) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *PaymentMethodTokenized) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *PaymentMethodTokenized) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *PaymentMethodTokenized) UnsetCurrency() {
	o.Currency.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodTokenized) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodTokenized) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *PaymentMethodTokenized) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *PaymentMethodTokenized) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *PaymentMethodTokenized) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *PaymentMethodTokenized) UnsetCountry() {
	o.Country.Unset()
}

func (o PaymentMethodTokenized) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Scheme.IsSet() {
		toSerialize["scheme"] = o.Scheme.Get()
	}
	if o.ExpirationDate.IsSet() {
		toSerialize["expiration_date"] = o.ExpirationDate.Get()
	}
	if o.ApprovalUrl.IsSet() {
		toSerialize["approval_url"] = o.ApprovalUrl.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentMethodTokenized struct {
	value *PaymentMethodTokenized
	isSet bool
}

func (v NullablePaymentMethodTokenized) Get() *PaymentMethodTokenized {
	return v.value
}

func (v *NullablePaymentMethodTokenized) Set(val *PaymentMethodTokenized) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodTokenized) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodTokenized) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodTokenized(val *PaymentMethodTokenized) *NullablePaymentMethodTokenized {
	return &NullablePaymentMethodTokenized{value: val, isSet: true}
}

func (v NullablePaymentMethodTokenized) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodTokenized) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


