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

// ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount A monetary amount in USD or local currency.
type ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount struct {
	// Transaction amount in USD.
	AmountUsd *string `json:"amount_usd,omitempty"`
	// Transaction amount in currency chosen by the buyer.
	AmountLocalCurrency *string `json:"amount_local_currency,omitempty"`
	// Transaction currency chosen by the buyer, 3-letter ISO-4217 format currency code.
	Currency *string `json:"currency,omitempty"`
}

// NewConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount instantiates a new ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount() *ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount {
	this := ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount{}
	return &this
}

// NewConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmountWithDefaults instantiates a new ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmountWithDefaults() *ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount {
	this := ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount{}
	return &this
}

// GetAmountUsd returns the AmountUsd field value if set, zero value otherwise.
func (o *ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) GetAmountUsd() string {
	if o == nil || o.AmountUsd == nil {
		var ret string
		return ret
	}
	return *o.AmountUsd
}

// GetAmountUsdOk returns a tuple with the AmountUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) GetAmountUsdOk() (*string, bool) {
	if o == nil || o.AmountUsd == nil {
		return nil, false
	}
	return o.AmountUsd, true
}

// HasAmountUsd returns a boolean if a field has been set.
func (o *ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) HasAmountUsd() bool {
	if o != nil && o.AmountUsd != nil {
		return true
	}

	return false
}

// SetAmountUsd gets a reference to the given string and assigns it to the AmountUsd field.
func (o *ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) SetAmountUsd(v string) {
	o.AmountUsd = &v
}

// GetAmountLocalCurrency returns the AmountLocalCurrency field value if set, zero value otherwise.
func (o *ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) GetAmountLocalCurrency() string {
	if o == nil || o.AmountLocalCurrency == nil {
		var ret string
		return ret
	}
	return *o.AmountLocalCurrency
}

// GetAmountLocalCurrencyOk returns a tuple with the AmountLocalCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) GetAmountLocalCurrencyOk() (*string, bool) {
	if o == nil || o.AmountLocalCurrency == nil {
		return nil, false
	}
	return o.AmountLocalCurrency, true
}

// HasAmountLocalCurrency returns a boolean if a field has been set.
func (o *ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) HasAmountLocalCurrency() bool {
	if o != nil && o.AmountLocalCurrency != nil {
		return true
	}

	return false
}

// SetAmountLocalCurrency gets a reference to the given string and assigns it to the AmountLocalCurrency field.
func (o *ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) SetAmountLocalCurrency(v string) {
	o.AmountLocalCurrency = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) SetCurrency(v string) {
	o.Currency = &v
}

func (o ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AmountUsd != nil {
		toSerialize["amount_usd"] = o.AmountUsd
	}
	if o.AmountLocalCurrency != nil {
		toSerialize["amount_local_currency"] = o.AmountLocalCurrency
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount struct {
	value *ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount
	isSet bool
}

func (v NullableConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) Get() *ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount {
	return v.value
}

func (v *NullableConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) Set(val *ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount(val *ConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) *NullableConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount {
	return &NullableConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount{value: val, isSet: true}
}

func (v NullableConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionOptionsForterAntiFraudTotalDiscountCouponDiscountAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


