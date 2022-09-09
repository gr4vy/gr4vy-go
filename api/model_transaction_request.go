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

// TransactionRequest A request to create a transaction.
type TransactionRequest struct {
	// The monetary amount to create an authorization for, in the smallest currency unit for the given currency, for example `1299` cents to create an authorization for `$12.99`.  If the `intent` is set to `capture`, an amount greater than zero must be supplied.
	Amount int32 `json:"amount"`
	// A supported ISO-4217 currency code.  For redirect requests, this value must match the one specified for `currency` in `payment_method`. 
	Currency string `json:"currency"`
	// The 2-letter ISO code of the country of the transaction. This is used to filter the payment services that is used to process the transaction.  If this value is provided for redirect requests and it's not `null`, it must match the one specified for `country` in `payment_method`. Otherwise, the value specified for `country` in `payment_method` will be assumed implicitly. 
	Country NullableString `json:"country,omitempty"`
	PaymentMethod TransactionPaymentMethodRequest `json:"payment_method"`
	// Whether or not to also try and store the payment method with us so that it can be used again for future use. This is only supported for payment methods that support this feature. There are also a few restrictions on how the flag may be set:  * The flag has to be set to `true` when the `payment_source` is set to `recurring` or `installment`, and `merchant_initiated` is set to `false`.  * The flag has to be set to `false` (or not set) when using a previously tokenized payment method.
	Store *bool `json:"store,omitempty"`
	// Defines the intent of this API call. This determines the desired initial state of the transaction.  * `authorize` - (Default) Optionally approves and then authorizes a transaction but does not capture the funds. * `capture` - Optionally approves and then authorizes and captures the funds of the transaction.
	Intent *string `json:"intent,omitempty"`
	// An external identifier that can be used to match the transaction against your own records.
	ExternalIdentifier NullableString `json:"external_identifier,omitempty"`
	ThreeDSecureData *ThreeDSecureDataV1V2 `json:"three_d_secure_data,omitempty"`
	// Indicates whether the transaction was initiated by the merchant (true) or customer (false).
	MerchantInitiated *bool `json:"merchant_initiated,omitempty"`
	// The source of the transaction. Defaults to `ecommerce`.
	PaymentSource *string `json:"payment_source,omitempty"`
	// Indicates whether the transaction represents a subsequent payment coming from a setup recurring payment. Please note there are some restrictions on how this flag may be used.  The flag can only be `false` (or not set) when the transaction meets one of the following criteria:  * It is not `merchant_initiated`. * `payment_source` is set to `card_on_file`.  The flag can only be set to `true` when the transaction meets one of the following criteria:  * It is not `merchant_initiated`. * `payment_source` is set to `recurring` or `installment` and `merchant_initiated` is set to `true`. * `payment_source` is set to `card_on_file`.
	IsSubsequentPayment *bool `json:"is_subsequent_payment,omitempty"`
	// Any additional information about the transaction that you would like to store as key-value pairs. This data is passed to payment service providers that support it. Please visit https://gr4vy.com/docs/ under `Connections` for more information on how specific providers support metadata.
	Metadata *map[string]string `json:"metadata,omitempty"`
	StatementDescriptor *StatementDescriptor `json:"statement_descriptor,omitempty"`
	// An array of cart items that represents the line items of a transaction.
	CartItems *[]CartItem `json:"cart_items,omitempty"`
	// A scheme's transaction identifier to use in connecting a merchant initiated transaction to a previous customer initiated transaction.  If not provided, and a qualifying customer initiated transaction has been previously made, then Gr4vy will populate this value with the identifier returned for that transaction.  e.g. the Visa Transaction Identifier, or Mastercard Trace ID.
	PreviousSchemeTransactionId NullableString `json:"previous_scheme_transaction_id,omitempty"`
	BrowserInfo *BrowserInfo `json:"browser_info,omitempty"`
}

// NewTransactionRequest instantiates a new TransactionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRequest(amount int32, currency string, paymentMethod TransactionPaymentMethodRequest) *TransactionRequest {
	this := TransactionRequest{}
	this.Amount = amount
	this.Currency = currency
	this.PaymentMethod = paymentMethod
	var store bool = false
	this.Store = &store
	var intent string = "authorize"
	this.Intent = &intent
	var merchantInitiated bool = false
	this.MerchantInitiated = &merchantInitiated
	var isSubsequentPayment bool = false
	this.IsSubsequentPayment = &isSubsequentPayment
	var previousSchemeTransactionId string = "null"
	this.PreviousSchemeTransactionId = *NewNullableString(&previousSchemeTransactionId)
	return &this
}

// NewTransactionRequestWithDefaults instantiates a new TransactionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRequestWithDefaults() *TransactionRequest {
	this := TransactionRequest{}
	var store bool = false
	this.Store = &store
	var intent string = "authorize"
	this.Intent = &intent
	var merchantInitiated bool = false
	this.MerchantInitiated = &merchantInitiated
	var isSubsequentPayment bool = false
	this.IsSubsequentPayment = &isSubsequentPayment
	var previousSchemeTransactionId string = "null"
	this.PreviousSchemeTransactionId = *NewNullableString(&previousSchemeTransactionId)
	return &this
}

// GetAmount returns the Amount field value
func (o *TransactionRequest) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransactionRequest) GetAmountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransactionRequest) SetAmount(v int32) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *TransactionRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TransactionRequest) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TransactionRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionRequest) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionRequest) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *TransactionRequest) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *TransactionRequest) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *TransactionRequest) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *TransactionRequest) UnsetCountry() {
	o.Country.Unset()
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *TransactionRequest) GetPaymentMethod() TransactionPaymentMethodRequest {
	if o == nil {
		var ret TransactionPaymentMethodRequest
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *TransactionRequest) GetPaymentMethodOk() (*TransactionPaymentMethodRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *TransactionRequest) SetPaymentMethod(v TransactionPaymentMethodRequest) {
	o.PaymentMethod = v
}

// GetStore returns the Store field value if set, zero value otherwise.
func (o *TransactionRequest) GetStore() bool {
	if o == nil || o.Store == nil {
		var ret bool
		return ret
	}
	return *o.Store
}

// GetStoreOk returns a tuple with the Store field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRequest) GetStoreOk() (*bool, bool) {
	if o == nil || o.Store == nil {
		return nil, false
	}
	return o.Store, true
}

// HasStore returns a boolean if a field has been set.
func (o *TransactionRequest) HasStore() bool {
	if o != nil && o.Store != nil {
		return true
	}

	return false
}

// SetStore gets a reference to the given bool and assigns it to the Store field.
func (o *TransactionRequest) SetStore(v bool) {
	o.Store = &v
}

// GetIntent returns the Intent field value if set, zero value otherwise.
func (o *TransactionRequest) GetIntent() string {
	if o == nil || o.Intent == nil {
		var ret string
		return ret
	}
	return *o.Intent
}

// GetIntentOk returns a tuple with the Intent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRequest) GetIntentOk() (*string, bool) {
	if o == nil || o.Intent == nil {
		return nil, false
	}
	return o.Intent, true
}

// HasIntent returns a boolean if a field has been set.
func (o *TransactionRequest) HasIntent() bool {
	if o != nil && o.Intent != nil {
		return true
	}

	return false
}

// SetIntent gets a reference to the given string and assigns it to the Intent field.
func (o *TransactionRequest) SetIntent(v string) {
	o.Intent = &v
}

// GetExternalIdentifier returns the ExternalIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionRequest) GetExternalIdentifier() string {
	if o == nil || o.ExternalIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalIdentifier.Get()
}

// GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionRequest) GetExternalIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalIdentifier.Get(), o.ExternalIdentifier.IsSet()
}

// HasExternalIdentifier returns a boolean if a field has been set.
func (o *TransactionRequest) HasExternalIdentifier() bool {
	if o != nil && o.ExternalIdentifier.IsSet() {
		return true
	}

	return false
}

// SetExternalIdentifier gets a reference to the given NullableString and assigns it to the ExternalIdentifier field.
func (o *TransactionRequest) SetExternalIdentifier(v string) {
	o.ExternalIdentifier.Set(&v)
}
// SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil
func (o *TransactionRequest) SetExternalIdentifierNil() {
	o.ExternalIdentifier.Set(nil)
}

// UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
func (o *TransactionRequest) UnsetExternalIdentifier() {
	o.ExternalIdentifier.Unset()
}

// GetThreeDSecureData returns the ThreeDSecureData field value if set, zero value otherwise.
func (o *TransactionRequest) GetThreeDSecureData() ThreeDSecureDataV1V2 {
	if o == nil || o.ThreeDSecureData == nil {
		var ret ThreeDSecureDataV1V2
		return ret
	}
	return *o.ThreeDSecureData
}

// GetThreeDSecureDataOk returns a tuple with the ThreeDSecureData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRequest) GetThreeDSecureDataOk() (*ThreeDSecureDataV1V2, bool) {
	if o == nil || o.ThreeDSecureData == nil {
		return nil, false
	}
	return o.ThreeDSecureData, true
}

// HasThreeDSecureData returns a boolean if a field has been set.
func (o *TransactionRequest) HasThreeDSecureData() bool {
	if o != nil && o.ThreeDSecureData != nil {
		return true
	}

	return false
}

// SetThreeDSecureData gets a reference to the given ThreeDSecureDataV1V2 and assigns it to the ThreeDSecureData field.
func (o *TransactionRequest) SetThreeDSecureData(v ThreeDSecureDataV1V2) {
	o.ThreeDSecureData = &v
}

// GetMerchantInitiated returns the MerchantInitiated field value if set, zero value otherwise.
func (o *TransactionRequest) GetMerchantInitiated() bool {
	if o == nil || o.MerchantInitiated == nil {
		var ret bool
		return ret
	}
	return *o.MerchantInitiated
}

// GetMerchantInitiatedOk returns a tuple with the MerchantInitiated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRequest) GetMerchantInitiatedOk() (*bool, bool) {
	if o == nil || o.MerchantInitiated == nil {
		return nil, false
	}
	return o.MerchantInitiated, true
}

// HasMerchantInitiated returns a boolean if a field has been set.
func (o *TransactionRequest) HasMerchantInitiated() bool {
	if o != nil && o.MerchantInitiated != nil {
		return true
	}

	return false
}

// SetMerchantInitiated gets a reference to the given bool and assigns it to the MerchantInitiated field.
func (o *TransactionRequest) SetMerchantInitiated(v bool) {
	o.MerchantInitiated = &v
}

// GetPaymentSource returns the PaymentSource field value if set, zero value otherwise.
func (o *TransactionRequest) GetPaymentSource() string {
	if o == nil || o.PaymentSource == nil {
		var ret string
		return ret
	}
	return *o.PaymentSource
}

// GetPaymentSourceOk returns a tuple with the PaymentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRequest) GetPaymentSourceOk() (*string, bool) {
	if o == nil || o.PaymentSource == nil {
		return nil, false
	}
	return o.PaymentSource, true
}

// HasPaymentSource returns a boolean if a field has been set.
func (o *TransactionRequest) HasPaymentSource() bool {
	if o != nil && o.PaymentSource != nil {
		return true
	}

	return false
}

// SetPaymentSource gets a reference to the given string and assigns it to the PaymentSource field.
func (o *TransactionRequest) SetPaymentSource(v string) {
	o.PaymentSource = &v
}

// GetIsSubsequentPayment returns the IsSubsequentPayment field value if set, zero value otherwise.
func (o *TransactionRequest) GetIsSubsequentPayment() bool {
	if o == nil || o.IsSubsequentPayment == nil {
		var ret bool
		return ret
	}
	return *o.IsSubsequentPayment
}

// GetIsSubsequentPaymentOk returns a tuple with the IsSubsequentPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRequest) GetIsSubsequentPaymentOk() (*bool, bool) {
	if o == nil || o.IsSubsequentPayment == nil {
		return nil, false
	}
	return o.IsSubsequentPayment, true
}

// HasIsSubsequentPayment returns a boolean if a field has been set.
func (o *TransactionRequest) HasIsSubsequentPayment() bool {
	if o != nil && o.IsSubsequentPayment != nil {
		return true
	}

	return false
}

// SetIsSubsequentPayment gets a reference to the given bool and assigns it to the IsSubsequentPayment field.
func (o *TransactionRequest) SetIsSubsequentPayment(v bool) {
	o.IsSubsequentPayment = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *TransactionRequest) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *TransactionRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *TransactionRequest) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetStatementDescriptor returns the StatementDescriptor field value if set, zero value otherwise.
func (o *TransactionRequest) GetStatementDescriptor() StatementDescriptor {
	if o == nil || o.StatementDescriptor == nil {
		var ret StatementDescriptor
		return ret
	}
	return *o.StatementDescriptor
}

// GetStatementDescriptorOk returns a tuple with the StatementDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRequest) GetStatementDescriptorOk() (*StatementDescriptor, bool) {
	if o == nil || o.StatementDescriptor == nil {
		return nil, false
	}
	return o.StatementDescriptor, true
}

// HasStatementDescriptor returns a boolean if a field has been set.
func (o *TransactionRequest) HasStatementDescriptor() bool {
	if o != nil && o.StatementDescriptor != nil {
		return true
	}

	return false
}

// SetStatementDescriptor gets a reference to the given StatementDescriptor and assigns it to the StatementDescriptor field.
func (o *TransactionRequest) SetStatementDescriptor(v StatementDescriptor) {
	o.StatementDescriptor = &v
}

// GetCartItems returns the CartItems field value if set, zero value otherwise.
func (o *TransactionRequest) GetCartItems() []CartItem {
	if o == nil || o.CartItems == nil {
		var ret []CartItem
		return ret
	}
	return *o.CartItems
}

// GetCartItemsOk returns a tuple with the CartItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRequest) GetCartItemsOk() (*[]CartItem, bool) {
	if o == nil || o.CartItems == nil {
		return nil, false
	}
	return o.CartItems, true
}

// HasCartItems returns a boolean if a field has been set.
func (o *TransactionRequest) HasCartItems() bool {
	if o != nil && o.CartItems != nil {
		return true
	}

	return false
}

// SetCartItems gets a reference to the given []CartItem and assigns it to the CartItems field.
func (o *TransactionRequest) SetCartItems(v []CartItem) {
	o.CartItems = &v
}

// GetPreviousSchemeTransactionId returns the PreviousSchemeTransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionRequest) GetPreviousSchemeTransactionId() string {
	if o == nil || o.PreviousSchemeTransactionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PreviousSchemeTransactionId.Get()
}

// GetPreviousSchemeTransactionIdOk returns a tuple with the PreviousSchemeTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionRequest) GetPreviousSchemeTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreviousSchemeTransactionId.Get(), o.PreviousSchemeTransactionId.IsSet()
}

// HasPreviousSchemeTransactionId returns a boolean if a field has been set.
func (o *TransactionRequest) HasPreviousSchemeTransactionId() bool {
	if o != nil && o.PreviousSchemeTransactionId.IsSet() {
		return true
	}

	return false
}

// SetPreviousSchemeTransactionId gets a reference to the given NullableString and assigns it to the PreviousSchemeTransactionId field.
func (o *TransactionRequest) SetPreviousSchemeTransactionId(v string) {
	o.PreviousSchemeTransactionId.Set(&v)
}
// SetPreviousSchemeTransactionIdNil sets the value for PreviousSchemeTransactionId to be an explicit nil
func (o *TransactionRequest) SetPreviousSchemeTransactionIdNil() {
	o.PreviousSchemeTransactionId.Set(nil)
}

// UnsetPreviousSchemeTransactionId ensures that no value is present for PreviousSchemeTransactionId, not even an explicit nil
func (o *TransactionRequest) UnsetPreviousSchemeTransactionId() {
	o.PreviousSchemeTransactionId.Unset()
}

// GetBrowserInfo returns the BrowserInfo field value if set, zero value otherwise.
func (o *TransactionRequest) GetBrowserInfo() BrowserInfo {
	if o == nil || o.BrowserInfo == nil {
		var ret BrowserInfo
		return ret
	}
	return *o.BrowserInfo
}

// GetBrowserInfoOk returns a tuple with the BrowserInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRequest) GetBrowserInfoOk() (*BrowserInfo, bool) {
	if o == nil || o.BrowserInfo == nil {
		return nil, false
	}
	return o.BrowserInfo, true
}

// HasBrowserInfo returns a boolean if a field has been set.
func (o *TransactionRequest) HasBrowserInfo() bool {
	if o != nil && o.BrowserInfo != nil {
		return true
	}

	return false
}

// SetBrowserInfo gets a reference to the given BrowserInfo and assigns it to the BrowserInfo field.
func (o *TransactionRequest) SetBrowserInfo(v BrowserInfo) {
	o.BrowserInfo = &v
}

func (o TransactionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if true {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if o.Store != nil {
		toSerialize["store"] = o.Store
	}
	if o.Intent != nil {
		toSerialize["intent"] = o.Intent
	}
	if o.ExternalIdentifier.IsSet() {
		toSerialize["external_identifier"] = o.ExternalIdentifier.Get()
	}
	if o.ThreeDSecureData != nil {
		toSerialize["three_d_secure_data"] = o.ThreeDSecureData
	}
	if o.MerchantInitiated != nil {
		toSerialize["merchant_initiated"] = o.MerchantInitiated
	}
	if o.PaymentSource != nil {
		toSerialize["payment_source"] = o.PaymentSource
	}
	if o.IsSubsequentPayment != nil {
		toSerialize["is_subsequent_payment"] = o.IsSubsequentPayment
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.StatementDescriptor != nil {
		toSerialize["statement_descriptor"] = o.StatementDescriptor
	}
	if o.CartItems != nil {
		toSerialize["cart_items"] = o.CartItems
	}
	if o.PreviousSchemeTransactionId.IsSet() {
		toSerialize["previous_scheme_transaction_id"] = o.PreviousSchemeTransactionId.Get()
	}
	if o.BrowserInfo != nil {
		toSerialize["browser_info"] = o.BrowserInfo
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionRequest struct {
	value *TransactionRequest
	isSet bool
}

func (v NullableTransactionRequest) Get() *TransactionRequest {
	return v.value
}

func (v *NullableTransactionRequest) Set(val *TransactionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRequest(val *TransactionRequest) *NullableTransactionRequest {
	return &NullableTransactionRequest{value: val, isSet: true}
}

func (v NullableTransactionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


