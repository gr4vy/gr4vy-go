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

// DigitalWalletClickToPayTokenDecryptedContext Click to Pay decrypted token context.
type DigitalWalletClickToPayTokenDecryptedContext struct {
	// Correlation ID for transaction.
	CorrelationId *string `json:"correlation_id,omitempty"`
	// Merchant Checkout Transaction Identifier which links the client-side JavaScript calls and server-side API calls for a specific transaction.
	MerchantTransactionId *string `json:"merchant_transaction_id,omitempty"`
	// The type of payment instrument.
	Type *string `json:"type,omitempty"`
	// Expiration of the card/token.
	ExpirationDate *string `json:"expiration_date,omitempty"`
	// Online payment cryptogram, as defined by 3-D Secure.
	HasCryptogram *bool `json:"has_cryptogram,omitempty"`
	// The cardholder name.
	CardholderName NullableString `json:"cardholder_name,omitempty"`
	// First six digits of underlying card.
	CardBin NullableString `json:"card_bin,omitempty"`
	// Last four digits of underlying card.
	CardLastFour NullableString `json:"card_last_four,omitempty"`
	// Expiration date of underlying card.
	CardExpirationDate NullableString `json:"card_expiration_date,omitempty"`
	// Card type.
	CardType NullableString `json:"card_type,omitempty"`
	// Address line 1.
	BillingLine1 NullableString `json:"billing_line1,omitempty"`
	// Address line 2.
	BillingLine2 NullableString `json:"billing_line2,omitempty"`
	// Address city.
	BillingCity NullableString `json:"billing_city,omitempty"`
	// Address state.
	BillingState NullableString `json:"billing_state,omitempty"`
	// Address zip/postal code.
	BillingZip NullableString `json:"billing_zip,omitempty"`
	// ISO 3166-1 alpha 2 address country code.
	BillingCountryCode NullableString `json:"billing_country_code,omitempty"`
}

// NewDigitalWalletClickToPayTokenDecryptedContext instantiates a new DigitalWalletClickToPayTokenDecryptedContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigitalWalletClickToPayTokenDecryptedContext() *DigitalWalletClickToPayTokenDecryptedContext {
	this := DigitalWalletClickToPayTokenDecryptedContext{}
	return &this
}

// NewDigitalWalletClickToPayTokenDecryptedContextWithDefaults instantiates a new DigitalWalletClickToPayTokenDecryptedContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigitalWalletClickToPayTokenDecryptedContextWithDefaults() *DigitalWalletClickToPayTokenDecryptedContext {
	this := DigitalWalletClickToPayTokenDecryptedContext{}
	return &this
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCorrelationId() string {
	if o == nil || o.CorrelationId == nil {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCorrelationIdOk() (*string, bool) {
	if o == nil || o.CorrelationId == nil {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) HasCorrelationId() bool {
	if o != nil && o.CorrelationId != nil {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetMerchantTransactionId returns the MerchantTransactionId field value if set, zero value otherwise.
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetMerchantTransactionId() string {
	if o == nil || o.MerchantTransactionId == nil {
		var ret string
		return ret
	}
	return *o.MerchantTransactionId
}

// GetMerchantTransactionIdOk returns a tuple with the MerchantTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetMerchantTransactionIdOk() (*string, bool) {
	if o == nil || o.MerchantTransactionId == nil {
		return nil, false
	}
	return o.MerchantTransactionId, true
}

// HasMerchantTransactionId returns a boolean if a field has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) HasMerchantTransactionId() bool {
	if o != nil && o.MerchantTransactionId != nil {
		return true
	}

	return false
}

// SetMerchantTransactionId gets a reference to the given string and assigns it to the MerchantTransactionId field.
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetMerchantTransactionId(v string) {
	o.MerchantTransactionId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetType(v string) {
	o.Type = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetExpirationDate() string {
	if o == nil || o.ExpirationDate == nil {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetExpirationDateOk() (*string, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate != nil {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetHasCryptogram returns the HasCryptogram field value if set, zero value otherwise.
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetHasCryptogram() bool {
	if o == nil || o.HasCryptogram == nil {
		var ret bool
		return ret
	}
	return *o.HasCryptogram
}

// GetHasCryptogramOk returns a tuple with the HasCryptogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetHasCryptogramOk() (*bool, bool) {
	if o == nil || o.HasCryptogram == nil {
		return nil, false
	}
	return o.HasCryptogram, true
}

// HasHasCryptogram returns a boolean if a field has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) HasHasCryptogram() bool {
	if o != nil && o.HasCryptogram != nil {
		return true
	}

	return false
}

// SetHasCryptogram gets a reference to the given bool and assigns it to the HasCryptogram field.
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetHasCryptogram(v bool) {
	o.HasCryptogram = &v
}

// GetCardholderName returns the CardholderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCardholderName() string {
	if o == nil || o.CardholderName.Get() == nil {
		var ret string
		return ret
	}
	return *o.CardholderName.Get()
}

// GetCardholderNameOk returns a tuple with the CardholderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCardholderNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CardholderName.Get(), o.CardholderName.IsSet()
}

// HasCardholderName returns a boolean if a field has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) HasCardholderName() bool {
	if o != nil && o.CardholderName.IsSet() {
		return true
	}

	return false
}

// SetCardholderName gets a reference to the given NullableString and assigns it to the CardholderName field.
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCardholderName(v string) {
	o.CardholderName.Set(&v)
}
// SetCardholderNameNil sets the value for CardholderName to be an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCardholderNameNil() {
	o.CardholderName.Set(nil)
}

// UnsetCardholderName ensures that no value is present for CardholderName, not even an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetCardholderName() {
	o.CardholderName.Unset()
}

// GetCardBin returns the CardBin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCardBin() string {
	if o == nil || o.CardBin.Get() == nil {
		var ret string
		return ret
	}
	return *o.CardBin.Get()
}

// GetCardBinOk returns a tuple with the CardBin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCardBinOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CardBin.Get(), o.CardBin.IsSet()
}

// HasCardBin returns a boolean if a field has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) HasCardBin() bool {
	if o != nil && o.CardBin.IsSet() {
		return true
	}

	return false
}

// SetCardBin gets a reference to the given NullableString and assigns it to the CardBin field.
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCardBin(v string) {
	o.CardBin.Set(&v)
}
// SetCardBinNil sets the value for CardBin to be an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCardBinNil() {
	o.CardBin.Set(nil)
}

// UnsetCardBin ensures that no value is present for CardBin, not even an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetCardBin() {
	o.CardBin.Unset()
}

// GetCardLastFour returns the CardLastFour field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCardLastFour() string {
	if o == nil || o.CardLastFour.Get() == nil {
		var ret string
		return ret
	}
	return *o.CardLastFour.Get()
}

// GetCardLastFourOk returns a tuple with the CardLastFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCardLastFourOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CardLastFour.Get(), o.CardLastFour.IsSet()
}

// HasCardLastFour returns a boolean if a field has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) HasCardLastFour() bool {
	if o != nil && o.CardLastFour.IsSet() {
		return true
	}

	return false
}

// SetCardLastFour gets a reference to the given NullableString and assigns it to the CardLastFour field.
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCardLastFour(v string) {
	o.CardLastFour.Set(&v)
}
// SetCardLastFourNil sets the value for CardLastFour to be an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCardLastFourNil() {
	o.CardLastFour.Set(nil)
}

// UnsetCardLastFour ensures that no value is present for CardLastFour, not even an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetCardLastFour() {
	o.CardLastFour.Unset()
}

// GetCardExpirationDate returns the CardExpirationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCardExpirationDate() string {
	if o == nil || o.CardExpirationDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.CardExpirationDate.Get()
}

// GetCardExpirationDateOk returns a tuple with the CardExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCardExpirationDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CardExpirationDate.Get(), o.CardExpirationDate.IsSet()
}

// HasCardExpirationDate returns a boolean if a field has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) HasCardExpirationDate() bool {
	if o != nil && o.CardExpirationDate.IsSet() {
		return true
	}

	return false
}

// SetCardExpirationDate gets a reference to the given NullableString and assigns it to the CardExpirationDate field.
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCardExpirationDate(v string) {
	o.CardExpirationDate.Set(&v)
}
// SetCardExpirationDateNil sets the value for CardExpirationDate to be an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCardExpirationDateNil() {
	o.CardExpirationDate.Set(nil)
}

// UnsetCardExpirationDate ensures that no value is present for CardExpirationDate, not even an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetCardExpirationDate() {
	o.CardExpirationDate.Unset()
}

// GetCardType returns the CardType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCardType() string {
	if o == nil || o.CardType.Get() == nil {
		var ret string
		return ret
	}
	return *o.CardType.Get()
}

// GetCardTypeOk returns a tuple with the CardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetCardTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CardType.Get(), o.CardType.IsSet()
}

// HasCardType returns a boolean if a field has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) HasCardType() bool {
	if o != nil && o.CardType.IsSet() {
		return true
	}

	return false
}

// SetCardType gets a reference to the given NullableString and assigns it to the CardType field.
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCardType(v string) {
	o.CardType.Set(&v)
}
// SetCardTypeNil sets the value for CardType to be an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetCardTypeNil() {
	o.CardType.Set(nil)
}

// UnsetCardType ensures that no value is present for CardType, not even an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetCardType() {
	o.CardType.Unset()
}

// GetBillingLine1 returns the BillingLine1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingLine1() string {
	if o == nil || o.BillingLine1.Get() == nil {
		var ret string
		return ret
	}
	return *o.BillingLine1.Get()
}

// GetBillingLine1Ok returns a tuple with the BillingLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingLine1Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BillingLine1.Get(), o.BillingLine1.IsSet()
}

// HasBillingLine1 returns a boolean if a field has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) HasBillingLine1() bool {
	if o != nil && o.BillingLine1.IsSet() {
		return true
	}

	return false
}

// SetBillingLine1 gets a reference to the given NullableString and assigns it to the BillingLine1 field.
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingLine1(v string) {
	o.BillingLine1.Set(&v)
}
// SetBillingLine1Nil sets the value for BillingLine1 to be an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingLine1Nil() {
	o.BillingLine1.Set(nil)
}

// UnsetBillingLine1 ensures that no value is present for BillingLine1, not even an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetBillingLine1() {
	o.BillingLine1.Unset()
}

// GetBillingLine2 returns the BillingLine2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingLine2() string {
	if o == nil || o.BillingLine2.Get() == nil {
		var ret string
		return ret
	}
	return *o.BillingLine2.Get()
}

// GetBillingLine2Ok returns a tuple with the BillingLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingLine2Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BillingLine2.Get(), o.BillingLine2.IsSet()
}

// HasBillingLine2 returns a boolean if a field has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) HasBillingLine2() bool {
	if o != nil && o.BillingLine2.IsSet() {
		return true
	}

	return false
}

// SetBillingLine2 gets a reference to the given NullableString and assigns it to the BillingLine2 field.
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingLine2(v string) {
	o.BillingLine2.Set(&v)
}
// SetBillingLine2Nil sets the value for BillingLine2 to be an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingLine2Nil() {
	o.BillingLine2.Set(nil)
}

// UnsetBillingLine2 ensures that no value is present for BillingLine2, not even an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetBillingLine2() {
	o.BillingLine2.Unset()
}

// GetBillingCity returns the BillingCity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingCity() string {
	if o == nil || o.BillingCity.Get() == nil {
		var ret string
		return ret
	}
	return *o.BillingCity.Get()
}

// GetBillingCityOk returns a tuple with the BillingCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BillingCity.Get(), o.BillingCity.IsSet()
}

// HasBillingCity returns a boolean if a field has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) HasBillingCity() bool {
	if o != nil && o.BillingCity.IsSet() {
		return true
	}

	return false
}

// SetBillingCity gets a reference to the given NullableString and assigns it to the BillingCity field.
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingCity(v string) {
	o.BillingCity.Set(&v)
}
// SetBillingCityNil sets the value for BillingCity to be an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingCityNil() {
	o.BillingCity.Set(nil)
}

// UnsetBillingCity ensures that no value is present for BillingCity, not even an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetBillingCity() {
	o.BillingCity.Unset()
}

// GetBillingState returns the BillingState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingState() string {
	if o == nil || o.BillingState.Get() == nil {
		var ret string
		return ret
	}
	return *o.BillingState.Get()
}

// GetBillingStateOk returns a tuple with the BillingState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BillingState.Get(), o.BillingState.IsSet()
}

// HasBillingState returns a boolean if a field has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) HasBillingState() bool {
	if o != nil && o.BillingState.IsSet() {
		return true
	}

	return false
}

// SetBillingState gets a reference to the given NullableString and assigns it to the BillingState field.
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingState(v string) {
	o.BillingState.Set(&v)
}
// SetBillingStateNil sets the value for BillingState to be an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingStateNil() {
	o.BillingState.Set(nil)
}

// UnsetBillingState ensures that no value is present for BillingState, not even an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetBillingState() {
	o.BillingState.Unset()
}

// GetBillingZip returns the BillingZip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingZip() string {
	if o == nil || o.BillingZip.Get() == nil {
		var ret string
		return ret
	}
	return *o.BillingZip.Get()
}

// GetBillingZipOk returns a tuple with the BillingZip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingZipOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BillingZip.Get(), o.BillingZip.IsSet()
}

// HasBillingZip returns a boolean if a field has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) HasBillingZip() bool {
	if o != nil && o.BillingZip.IsSet() {
		return true
	}

	return false
}

// SetBillingZip gets a reference to the given NullableString and assigns it to the BillingZip field.
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingZip(v string) {
	o.BillingZip.Set(&v)
}
// SetBillingZipNil sets the value for BillingZip to be an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingZipNil() {
	o.BillingZip.Set(nil)
}

// UnsetBillingZip ensures that no value is present for BillingZip, not even an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetBillingZip() {
	o.BillingZip.Unset()
}

// GetBillingCountryCode returns the BillingCountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingCountryCode() string {
	if o == nil || o.BillingCountryCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.BillingCountryCode.Get()
}

// GetBillingCountryCodeOk returns a tuple with the BillingCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigitalWalletClickToPayTokenDecryptedContext) GetBillingCountryCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BillingCountryCode.Get(), o.BillingCountryCode.IsSet()
}

// HasBillingCountryCode returns a boolean if a field has been set.
func (o *DigitalWalletClickToPayTokenDecryptedContext) HasBillingCountryCode() bool {
	if o != nil && o.BillingCountryCode.IsSet() {
		return true
	}

	return false
}

// SetBillingCountryCode gets a reference to the given NullableString and assigns it to the BillingCountryCode field.
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingCountryCode(v string) {
	o.BillingCountryCode.Set(&v)
}
// SetBillingCountryCodeNil sets the value for BillingCountryCode to be an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) SetBillingCountryCodeNil() {
	o.BillingCountryCode.Set(nil)
}

// UnsetBillingCountryCode ensures that no value is present for BillingCountryCode, not even an explicit nil
func (o *DigitalWalletClickToPayTokenDecryptedContext) UnsetBillingCountryCode() {
	o.BillingCountryCode.Unset()
}

func (o DigitalWalletClickToPayTokenDecryptedContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CorrelationId != nil {
		toSerialize["correlation_id"] = o.CorrelationId
	}
	if o.MerchantTransactionId != nil {
		toSerialize["merchant_transaction_id"] = o.MerchantTransactionId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ExpirationDate != nil {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	if o.HasCryptogram != nil {
		toSerialize["has_cryptogram"] = o.HasCryptogram
	}
	if o.CardholderName.IsSet() {
		toSerialize["cardholder_name"] = o.CardholderName.Get()
	}
	if o.CardBin.IsSet() {
		toSerialize["card_bin"] = o.CardBin.Get()
	}
	if o.CardLastFour.IsSet() {
		toSerialize["card_last_four"] = o.CardLastFour.Get()
	}
	if o.CardExpirationDate.IsSet() {
		toSerialize["card_expiration_date"] = o.CardExpirationDate.Get()
	}
	if o.CardType.IsSet() {
		toSerialize["card_type"] = o.CardType.Get()
	}
	if o.BillingLine1.IsSet() {
		toSerialize["billing_line1"] = o.BillingLine1.Get()
	}
	if o.BillingLine2.IsSet() {
		toSerialize["billing_line2"] = o.BillingLine2.Get()
	}
	if o.BillingCity.IsSet() {
		toSerialize["billing_city"] = o.BillingCity.Get()
	}
	if o.BillingState.IsSet() {
		toSerialize["billing_state"] = o.BillingState.Get()
	}
	if o.BillingZip.IsSet() {
		toSerialize["billing_zip"] = o.BillingZip.Get()
	}
	if o.BillingCountryCode.IsSet() {
		toSerialize["billing_country_code"] = o.BillingCountryCode.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDigitalWalletClickToPayTokenDecryptedContext struct {
	value *DigitalWalletClickToPayTokenDecryptedContext
	isSet bool
}

func (v NullableDigitalWalletClickToPayTokenDecryptedContext) Get() *DigitalWalletClickToPayTokenDecryptedContext {
	return v.value
}

func (v *NullableDigitalWalletClickToPayTokenDecryptedContext) Set(val *DigitalWalletClickToPayTokenDecryptedContext) {
	v.value = val
	v.isSet = true
}

func (v NullableDigitalWalletClickToPayTokenDecryptedContext) IsSet() bool {
	return v.isSet
}

func (v *NullableDigitalWalletClickToPayTokenDecryptedContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigitalWalletClickToPayTokenDecryptedContext(val *DigitalWalletClickToPayTokenDecryptedContext) *NullableDigitalWalletClickToPayTokenDecryptedContext {
	return &NullableDigitalWalletClickToPayTokenDecryptedContext{value: val, isSet: true}
}

func (v NullableDigitalWalletClickToPayTokenDecryptedContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigitalWalletClickToPayTokenDecryptedContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


