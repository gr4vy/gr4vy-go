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
	"time"
)

// GiftCardSummary A gift card stored for a buyer.
type GiftCardSummary struct {
	// The type of this resource.
	Type *string `json:"type,omitempty"`
	// The ID of this gift card.
	Id NullableString `json:"id,omitempty"`
	// The unique ID for a merchant account.
	MerchantAccountId *string `json:"merchant_account_id,omitempty"`
	// The first 6 digits of the full gift card number.
	Bin *string `json:"bin,omitempty"`
	// The 3 digits after the `bin` of the full gift card number.
	SubBin *string `json:"sub_bin,omitempty"`
	// The last 4 digits for the gift card.
	Last4 *string `json:"last4,omitempty"`
	// The date and time when this gift card expires. This is a full date/time and may be more accurate than the actual expiry date received by the gift card service.
	ExpirationDate NullableTime `json:"expiration_date,omitempty"`
	// The amount remaining on the balance for this gift card according to the gift card service. This may be `null` if the balance could not be fetched.
	Balance NullableInt32 `json:"balance,omitempty"`
	// The ISO-4217 currency code that this gift card has a balance for.
	Currency NullableString `json:"currency,omitempty"`
	// If the last balance update failed, this will contain the internal code for this error.
	BalanceErrorCode NullableString `json:"balance_error_code,omitempty"`
	// If the last balance update failed, this will contain the the raw error code received from the gift card provider.
	BalanceRawErrorCode NullableString `json:"balance_raw_error_code,omitempty"`
	// If the last balance update failed, this will contain the the raw error message received from the gift card provider.
	BalanceRawErrorMessage NullableString `json:"balance_raw_error_message,omitempty"`
}

// NewGiftCardSummary instantiates a new GiftCardSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardSummary() *GiftCardSummary {
	this := GiftCardSummary{}
	return &this
}

// NewGiftCardSummaryWithDefaults instantiates a new GiftCardSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardSummaryWithDefaults() *GiftCardSummary {
	this := GiftCardSummary{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GiftCardSummary) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardSummary) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GiftCardSummary) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GiftCardSummary) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftCardSummary) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardSummary) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *GiftCardSummary) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *GiftCardSummary) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *GiftCardSummary) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *GiftCardSummary) UnsetId() {
	o.Id.Unset()
}

// GetMerchantAccountId returns the MerchantAccountId field value if set, zero value otherwise.
func (o *GiftCardSummary) GetMerchantAccountId() string {
	if o == nil || o.MerchantAccountId == nil {
		var ret string
		return ret
	}
	return *o.MerchantAccountId
}

// GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardSummary) GetMerchantAccountIdOk() (*string, bool) {
	if o == nil || o.MerchantAccountId == nil {
		return nil, false
	}
	return o.MerchantAccountId, true
}

// HasMerchantAccountId returns a boolean if a field has been set.
func (o *GiftCardSummary) HasMerchantAccountId() bool {
	if o != nil && o.MerchantAccountId != nil {
		return true
	}

	return false
}

// SetMerchantAccountId gets a reference to the given string and assigns it to the MerchantAccountId field.
func (o *GiftCardSummary) SetMerchantAccountId(v string) {
	o.MerchantAccountId = &v
}

// GetBin returns the Bin field value if set, zero value otherwise.
func (o *GiftCardSummary) GetBin() string {
	if o == nil || o.Bin == nil {
		var ret string
		return ret
	}
	return *o.Bin
}

// GetBinOk returns a tuple with the Bin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardSummary) GetBinOk() (*string, bool) {
	if o == nil || o.Bin == nil {
		return nil, false
	}
	return o.Bin, true
}

// HasBin returns a boolean if a field has been set.
func (o *GiftCardSummary) HasBin() bool {
	if o != nil && o.Bin != nil {
		return true
	}

	return false
}

// SetBin gets a reference to the given string and assigns it to the Bin field.
func (o *GiftCardSummary) SetBin(v string) {
	o.Bin = &v
}

// GetSubBin returns the SubBin field value if set, zero value otherwise.
func (o *GiftCardSummary) GetSubBin() string {
	if o == nil || o.SubBin == nil {
		var ret string
		return ret
	}
	return *o.SubBin
}

// GetSubBinOk returns a tuple with the SubBin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardSummary) GetSubBinOk() (*string, bool) {
	if o == nil || o.SubBin == nil {
		return nil, false
	}
	return o.SubBin, true
}

// HasSubBin returns a boolean if a field has been set.
func (o *GiftCardSummary) HasSubBin() bool {
	if o != nil && o.SubBin != nil {
		return true
	}

	return false
}

// SetSubBin gets a reference to the given string and assigns it to the SubBin field.
func (o *GiftCardSummary) SetSubBin(v string) {
	o.SubBin = &v
}

// GetLast4 returns the Last4 field value if set, zero value otherwise.
func (o *GiftCardSummary) GetLast4() string {
	if o == nil || o.Last4 == nil {
		var ret string
		return ret
	}
	return *o.Last4
}

// GetLast4Ok returns a tuple with the Last4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardSummary) GetLast4Ok() (*string, bool) {
	if o == nil || o.Last4 == nil {
		return nil, false
	}
	return o.Last4, true
}

// HasLast4 returns a boolean if a field has been set.
func (o *GiftCardSummary) HasLast4() bool {
	if o != nil && o.Last4 != nil {
		return true
	}

	return false
}

// SetLast4 gets a reference to the given string and assigns it to the Last4 field.
func (o *GiftCardSummary) SetLast4(v string) {
	o.Last4 = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftCardSummary) GetExpirationDate() time.Time {
	if o == nil || o.ExpirationDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate.Get()
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardSummary) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpirationDate.Get(), o.ExpirationDate.IsSet()
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *GiftCardSummary) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate.IsSet() {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given NullableTime and assigns it to the ExpirationDate field.
func (o *GiftCardSummary) SetExpirationDate(v time.Time) {
	o.ExpirationDate.Set(&v)
}
// SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil
func (o *GiftCardSummary) SetExpirationDateNil() {
	o.ExpirationDate.Set(nil)
}

// UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
func (o *GiftCardSummary) UnsetExpirationDate() {
	o.ExpirationDate.Unset()
}

// GetBalance returns the Balance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftCardSummary) GetBalance() int32 {
	if o == nil || o.Balance.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Balance.Get()
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardSummary) GetBalanceOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Balance.Get(), o.Balance.IsSet()
}

// HasBalance returns a boolean if a field has been set.
func (o *GiftCardSummary) HasBalance() bool {
	if o != nil && o.Balance.IsSet() {
		return true
	}

	return false
}

// SetBalance gets a reference to the given NullableInt32 and assigns it to the Balance field.
func (o *GiftCardSummary) SetBalance(v int32) {
	o.Balance.Set(&v)
}
// SetBalanceNil sets the value for Balance to be an explicit nil
func (o *GiftCardSummary) SetBalanceNil() {
	o.Balance.Set(nil)
}

// UnsetBalance ensures that no value is present for Balance, not even an explicit nil
func (o *GiftCardSummary) UnsetBalance() {
	o.Balance.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftCardSummary) GetCurrency() string {
	if o == nil || o.Currency.Get() == nil {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardSummary) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *GiftCardSummary) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *GiftCardSummary) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *GiftCardSummary) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *GiftCardSummary) UnsetCurrency() {
	o.Currency.Unset()
}

// GetBalanceErrorCode returns the BalanceErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftCardSummary) GetBalanceErrorCode() string {
	if o == nil || o.BalanceErrorCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.BalanceErrorCode.Get()
}

// GetBalanceErrorCodeOk returns a tuple with the BalanceErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardSummary) GetBalanceErrorCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BalanceErrorCode.Get(), o.BalanceErrorCode.IsSet()
}

// HasBalanceErrorCode returns a boolean if a field has been set.
func (o *GiftCardSummary) HasBalanceErrorCode() bool {
	if o != nil && o.BalanceErrorCode.IsSet() {
		return true
	}

	return false
}

// SetBalanceErrorCode gets a reference to the given NullableString and assigns it to the BalanceErrorCode field.
func (o *GiftCardSummary) SetBalanceErrorCode(v string) {
	o.BalanceErrorCode.Set(&v)
}
// SetBalanceErrorCodeNil sets the value for BalanceErrorCode to be an explicit nil
func (o *GiftCardSummary) SetBalanceErrorCodeNil() {
	o.BalanceErrorCode.Set(nil)
}

// UnsetBalanceErrorCode ensures that no value is present for BalanceErrorCode, not even an explicit nil
func (o *GiftCardSummary) UnsetBalanceErrorCode() {
	o.BalanceErrorCode.Unset()
}

// GetBalanceRawErrorCode returns the BalanceRawErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftCardSummary) GetBalanceRawErrorCode() string {
	if o == nil || o.BalanceRawErrorCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.BalanceRawErrorCode.Get()
}

// GetBalanceRawErrorCodeOk returns a tuple with the BalanceRawErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardSummary) GetBalanceRawErrorCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BalanceRawErrorCode.Get(), o.BalanceRawErrorCode.IsSet()
}

// HasBalanceRawErrorCode returns a boolean if a field has been set.
func (o *GiftCardSummary) HasBalanceRawErrorCode() bool {
	if o != nil && o.BalanceRawErrorCode.IsSet() {
		return true
	}

	return false
}

// SetBalanceRawErrorCode gets a reference to the given NullableString and assigns it to the BalanceRawErrorCode field.
func (o *GiftCardSummary) SetBalanceRawErrorCode(v string) {
	o.BalanceRawErrorCode.Set(&v)
}
// SetBalanceRawErrorCodeNil sets the value for BalanceRawErrorCode to be an explicit nil
func (o *GiftCardSummary) SetBalanceRawErrorCodeNil() {
	o.BalanceRawErrorCode.Set(nil)
}

// UnsetBalanceRawErrorCode ensures that no value is present for BalanceRawErrorCode, not even an explicit nil
func (o *GiftCardSummary) UnsetBalanceRawErrorCode() {
	o.BalanceRawErrorCode.Unset()
}

// GetBalanceRawErrorMessage returns the BalanceRawErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftCardSummary) GetBalanceRawErrorMessage() string {
	if o == nil || o.BalanceRawErrorMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.BalanceRawErrorMessage.Get()
}

// GetBalanceRawErrorMessageOk returns a tuple with the BalanceRawErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardSummary) GetBalanceRawErrorMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BalanceRawErrorMessage.Get(), o.BalanceRawErrorMessage.IsSet()
}

// HasBalanceRawErrorMessage returns a boolean if a field has been set.
func (o *GiftCardSummary) HasBalanceRawErrorMessage() bool {
	if o != nil && o.BalanceRawErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetBalanceRawErrorMessage gets a reference to the given NullableString and assigns it to the BalanceRawErrorMessage field.
func (o *GiftCardSummary) SetBalanceRawErrorMessage(v string) {
	o.BalanceRawErrorMessage.Set(&v)
}
// SetBalanceRawErrorMessageNil sets the value for BalanceRawErrorMessage to be an explicit nil
func (o *GiftCardSummary) SetBalanceRawErrorMessageNil() {
	o.BalanceRawErrorMessage.Set(nil)
}

// UnsetBalanceRawErrorMessage ensures that no value is present for BalanceRawErrorMessage, not even an explicit nil
func (o *GiftCardSummary) UnsetBalanceRawErrorMessage() {
	o.BalanceRawErrorMessage.Unset()
}

func (o GiftCardSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.MerchantAccountId != nil {
		toSerialize["merchant_account_id"] = o.MerchantAccountId
	}
	if o.Bin != nil {
		toSerialize["bin"] = o.Bin
	}
	if o.SubBin != nil {
		toSerialize["sub_bin"] = o.SubBin
	}
	if o.Last4 != nil {
		toSerialize["last4"] = o.Last4
	}
	if o.ExpirationDate.IsSet() {
		toSerialize["expiration_date"] = o.ExpirationDate.Get()
	}
	if o.Balance.IsSet() {
		toSerialize["balance"] = o.Balance.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.BalanceErrorCode.IsSet() {
		toSerialize["balance_error_code"] = o.BalanceErrorCode.Get()
	}
	if o.BalanceRawErrorCode.IsSet() {
		toSerialize["balance_raw_error_code"] = o.BalanceRawErrorCode.Get()
	}
	if o.BalanceRawErrorMessage.IsSet() {
		toSerialize["balance_raw_error_message"] = o.BalanceRawErrorMessage.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGiftCardSummary struct {
	value *GiftCardSummary
	isSet bool
}

func (v NullableGiftCardSummary) Get() *GiftCardSummary {
	return v.value
}

func (v *NullableGiftCardSummary) Set(val *GiftCardSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardSummary(val *GiftCardSummary) *NullableGiftCardSummary {
	return &NullableGiftCardSummary{value: val, isSet: true}
}

func (v NullableGiftCardSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


