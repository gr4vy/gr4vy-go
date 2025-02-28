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

// TransactionSummary A transaction record.
type TransactionSummary struct {
	// The type of this resource. Is always `transaction`.
	Type *string `json:"type,omitempty"`
	// The unique identifier for this transaction.
	Id *string `json:"id,omitempty"`
	// The authorized amount for this transaction. This can be more than the actual captured amount and part of this amount may be refunded.
	Amount *int32 `json:"amount,omitempty"`
	// The amount for this transaction that has been authorized for the `payment_method`. This can be less than the `amount` if gift cards were used.
	AuthorizedAmount *int32 `json:"authorized_amount,omitempty"`
	// The buyer used for this transaction.
	Buyer NullableBuyerSnapshot `json:"buyer,omitempty"`
	// The captured amount for this transaction. This can be the full value of the `authorized_amount` or less.
	CapturedAmount *int32 `json:"captured_amount,omitempty"`
	// The identifier for the checkout session this transaction is associated with.
	CheckoutSessionId NullableString `json:"checkout_session_id,omitempty"`
	// The 2-letter ISO code of the country of the transaction. This is used to filter the payment services that is used to process the transaction. 
	Country NullableString `json:"country,omitempty"`
	// The date and time when this transaction was created in our system.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The currency code for this transaction.
	Currency *string `json:"currency,omitempty"`
	// An external identifier that can be used to match the transaction against your own records.
	ExternalIdentifier NullableString `json:"external_identifier,omitempty"`
	// The gift cards redeemed for this transaction.
	GiftCardRedemptions *[]GiftCardRedemption `json:"gift_card_redemptions,omitempty"`
	// The name of the instrument used to process the transaction. 
	InstrumentType NullableString `json:"instrument_type,omitempty"`
	// The original `intent` used when the transaction was [created](#operation/authorize-new-transaction).
	Intent *string `json:"intent,omitempty"`
	// The ID of the merchant account to which this transaction belongs to.
	MerchantAccountId *string `json:"merchant_account_id,omitempty"`
	Method NullableString `json:"method,omitempty"`
	// The payment method used for this transaction.
	PaymentMethod NullablePaymentMethodSnapshot `json:"payment_method,omitempty"`
	// The payment service used for this transaction.
	PaymentService NullablePaymentServiceSnapshot `json:"payment_service,omitempty"`
	// Whether a manual review is pending.
	PendingReview *bool `json:"pending_review,omitempty"`
	// This is the response code received from the payment service. This can be set to any value and is not standardized across different payment services.
	RawResponseCode NullableString `json:"raw_response_code,omitempty"`
	// This is the response description received from the payment service. This can be set to any value and is not standardized across different payment services.
	RawResponseDescription NullableString `json:"raw_response_description,omitempty"`
	// The base62 encoded transaction ID. This represents a shorter version of this transaction's `id` which is sent to payment services, anti-fraud services, and other connectors. You can use this ID to reconcile a payment service's transaction against our system.  This ID is sent instead of the transaction ID because not all services support 36 digit identifiers.
	ReconciliationId *string `json:"reconciliation_id,omitempty"`
	// The refunded amount for this transaction. This can be the full value of the `captured_amount` or less.
	RefundedAmount *int32 `json:"refunded_amount,omitempty"`
	// The status of the transaction. The status may change over time as asynchronous processing events occur.
	Status *string `json:"status,omitempty"`
	// Defines when the transaction was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The currency of this transaction's settlement in ISO 4217 three-letter code format.
	SettledCurrency NullableString `json:"settled_currency,omitempty"`
	// The net amount settled for this transaction.
	SettledAmount *int32 `json:"settled_amount,omitempty"`
	// Indicates whether this transaction has been settled.
	Settled *bool `json:"settled,omitempty"`
}

// NewTransactionSummary instantiates a new TransactionSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionSummary() *TransactionSummary {
	this := TransactionSummary{}
	return &this
}

// NewTransactionSummaryWithDefaults instantiates a new TransactionSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionSummaryWithDefaults() *TransactionSummary {
	this := TransactionSummary{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransactionSummary) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransactionSummary) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransactionSummary) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransactionSummary) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransactionSummary) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TransactionSummary) SetId(v string) {
	o.Id = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *TransactionSummary) GetAmount() int32 {
	if o == nil || o.Amount == nil {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetAmountOk() (*int32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *TransactionSummary) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *TransactionSummary) SetAmount(v int32) {
	o.Amount = &v
}

// GetAuthorizedAmount returns the AuthorizedAmount field value if set, zero value otherwise.
func (o *TransactionSummary) GetAuthorizedAmount() int32 {
	if o == nil || o.AuthorizedAmount == nil {
		var ret int32
		return ret
	}
	return *o.AuthorizedAmount
}

// GetAuthorizedAmountOk returns a tuple with the AuthorizedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetAuthorizedAmountOk() (*int32, bool) {
	if o == nil || o.AuthorizedAmount == nil {
		return nil, false
	}
	return o.AuthorizedAmount, true
}

// HasAuthorizedAmount returns a boolean if a field has been set.
func (o *TransactionSummary) HasAuthorizedAmount() bool {
	if o != nil && o.AuthorizedAmount != nil {
		return true
	}

	return false
}

// SetAuthorizedAmount gets a reference to the given int32 and assigns it to the AuthorizedAmount field.
func (o *TransactionSummary) SetAuthorizedAmount(v int32) {
	o.AuthorizedAmount = &v
}

// GetBuyer returns the Buyer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionSummary) GetBuyer() BuyerSnapshot {
	if o == nil || o.Buyer.Get() == nil {
		var ret BuyerSnapshot
		return ret
	}
	return *o.Buyer.Get()
}

// GetBuyerOk returns a tuple with the Buyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionSummary) GetBuyerOk() (*BuyerSnapshot, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Buyer.Get(), o.Buyer.IsSet()
}

// HasBuyer returns a boolean if a field has been set.
func (o *TransactionSummary) HasBuyer() bool {
	if o != nil && o.Buyer.IsSet() {
		return true
	}

	return false
}

// SetBuyer gets a reference to the given NullableBuyerSnapshot and assigns it to the Buyer field.
func (o *TransactionSummary) SetBuyer(v BuyerSnapshot) {
	o.Buyer.Set(&v)
}
// SetBuyerNil sets the value for Buyer to be an explicit nil
func (o *TransactionSummary) SetBuyerNil() {
	o.Buyer.Set(nil)
}

// UnsetBuyer ensures that no value is present for Buyer, not even an explicit nil
func (o *TransactionSummary) UnsetBuyer() {
	o.Buyer.Unset()
}

// GetCapturedAmount returns the CapturedAmount field value if set, zero value otherwise.
func (o *TransactionSummary) GetCapturedAmount() int32 {
	if o == nil || o.CapturedAmount == nil {
		var ret int32
		return ret
	}
	return *o.CapturedAmount
}

// GetCapturedAmountOk returns a tuple with the CapturedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetCapturedAmountOk() (*int32, bool) {
	if o == nil || o.CapturedAmount == nil {
		return nil, false
	}
	return o.CapturedAmount, true
}

// HasCapturedAmount returns a boolean if a field has been set.
func (o *TransactionSummary) HasCapturedAmount() bool {
	if o != nil && o.CapturedAmount != nil {
		return true
	}

	return false
}

// SetCapturedAmount gets a reference to the given int32 and assigns it to the CapturedAmount field.
func (o *TransactionSummary) SetCapturedAmount(v int32) {
	o.CapturedAmount = &v
}

// GetCheckoutSessionId returns the CheckoutSessionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionSummary) GetCheckoutSessionId() string {
	if o == nil || o.CheckoutSessionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CheckoutSessionId.Get()
}

// GetCheckoutSessionIdOk returns a tuple with the CheckoutSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionSummary) GetCheckoutSessionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CheckoutSessionId.Get(), o.CheckoutSessionId.IsSet()
}

// HasCheckoutSessionId returns a boolean if a field has been set.
func (o *TransactionSummary) HasCheckoutSessionId() bool {
	if o != nil && o.CheckoutSessionId.IsSet() {
		return true
	}

	return false
}

// SetCheckoutSessionId gets a reference to the given NullableString and assigns it to the CheckoutSessionId field.
func (o *TransactionSummary) SetCheckoutSessionId(v string) {
	o.CheckoutSessionId.Set(&v)
}
// SetCheckoutSessionIdNil sets the value for CheckoutSessionId to be an explicit nil
func (o *TransactionSummary) SetCheckoutSessionIdNil() {
	o.CheckoutSessionId.Set(nil)
}

// UnsetCheckoutSessionId ensures that no value is present for CheckoutSessionId, not even an explicit nil
func (o *TransactionSummary) UnsetCheckoutSessionId() {
	o.CheckoutSessionId.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionSummary) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionSummary) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *TransactionSummary) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *TransactionSummary) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *TransactionSummary) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *TransactionSummary) UnsetCountry() {
	o.Country.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TransactionSummary) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TransactionSummary) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TransactionSummary) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *TransactionSummary) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *TransactionSummary) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *TransactionSummary) SetCurrency(v string) {
	o.Currency = &v
}

// GetExternalIdentifier returns the ExternalIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionSummary) GetExternalIdentifier() string {
	if o == nil || o.ExternalIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalIdentifier.Get()
}

// GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionSummary) GetExternalIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalIdentifier.Get(), o.ExternalIdentifier.IsSet()
}

// HasExternalIdentifier returns a boolean if a field has been set.
func (o *TransactionSummary) HasExternalIdentifier() bool {
	if o != nil && o.ExternalIdentifier.IsSet() {
		return true
	}

	return false
}

// SetExternalIdentifier gets a reference to the given NullableString and assigns it to the ExternalIdentifier field.
func (o *TransactionSummary) SetExternalIdentifier(v string) {
	o.ExternalIdentifier.Set(&v)
}
// SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil
func (o *TransactionSummary) SetExternalIdentifierNil() {
	o.ExternalIdentifier.Set(nil)
}

// UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
func (o *TransactionSummary) UnsetExternalIdentifier() {
	o.ExternalIdentifier.Unset()
}

// GetGiftCardRedemptions returns the GiftCardRedemptions field value if set, zero value otherwise.
func (o *TransactionSummary) GetGiftCardRedemptions() []GiftCardRedemption {
	if o == nil || o.GiftCardRedemptions == nil {
		var ret []GiftCardRedemption
		return ret
	}
	return *o.GiftCardRedemptions
}

// GetGiftCardRedemptionsOk returns a tuple with the GiftCardRedemptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetGiftCardRedemptionsOk() (*[]GiftCardRedemption, bool) {
	if o == nil || o.GiftCardRedemptions == nil {
		return nil, false
	}
	return o.GiftCardRedemptions, true
}

// HasGiftCardRedemptions returns a boolean if a field has been set.
func (o *TransactionSummary) HasGiftCardRedemptions() bool {
	if o != nil && o.GiftCardRedemptions != nil {
		return true
	}

	return false
}

// SetGiftCardRedemptions gets a reference to the given []GiftCardRedemption and assigns it to the GiftCardRedemptions field.
func (o *TransactionSummary) SetGiftCardRedemptions(v []GiftCardRedemption) {
	o.GiftCardRedemptions = &v
}

// GetInstrumentType returns the InstrumentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionSummary) GetInstrumentType() string {
	if o == nil || o.InstrumentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.InstrumentType.Get()
}

// GetInstrumentTypeOk returns a tuple with the InstrumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionSummary) GetInstrumentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InstrumentType.Get(), o.InstrumentType.IsSet()
}

// HasInstrumentType returns a boolean if a field has been set.
func (o *TransactionSummary) HasInstrumentType() bool {
	if o != nil && o.InstrumentType.IsSet() {
		return true
	}

	return false
}

// SetInstrumentType gets a reference to the given NullableString and assigns it to the InstrumentType field.
func (o *TransactionSummary) SetInstrumentType(v string) {
	o.InstrumentType.Set(&v)
}
// SetInstrumentTypeNil sets the value for InstrumentType to be an explicit nil
func (o *TransactionSummary) SetInstrumentTypeNil() {
	o.InstrumentType.Set(nil)
}

// UnsetInstrumentType ensures that no value is present for InstrumentType, not even an explicit nil
func (o *TransactionSummary) UnsetInstrumentType() {
	o.InstrumentType.Unset()
}

// GetIntent returns the Intent field value if set, zero value otherwise.
func (o *TransactionSummary) GetIntent() string {
	if o == nil || o.Intent == nil {
		var ret string
		return ret
	}
	return *o.Intent
}

// GetIntentOk returns a tuple with the Intent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetIntentOk() (*string, bool) {
	if o == nil || o.Intent == nil {
		return nil, false
	}
	return o.Intent, true
}

// HasIntent returns a boolean if a field has been set.
func (o *TransactionSummary) HasIntent() bool {
	if o != nil && o.Intent != nil {
		return true
	}

	return false
}

// SetIntent gets a reference to the given string and assigns it to the Intent field.
func (o *TransactionSummary) SetIntent(v string) {
	o.Intent = &v
}

// GetMerchantAccountId returns the MerchantAccountId field value if set, zero value otherwise.
func (o *TransactionSummary) GetMerchantAccountId() string {
	if o == nil || o.MerchantAccountId == nil {
		var ret string
		return ret
	}
	return *o.MerchantAccountId
}

// GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetMerchantAccountIdOk() (*string, bool) {
	if o == nil || o.MerchantAccountId == nil {
		return nil, false
	}
	return o.MerchantAccountId, true
}

// HasMerchantAccountId returns a boolean if a field has been set.
func (o *TransactionSummary) HasMerchantAccountId() bool {
	if o != nil && o.MerchantAccountId != nil {
		return true
	}

	return false
}

// SetMerchantAccountId gets a reference to the given string and assigns it to the MerchantAccountId field.
func (o *TransactionSummary) SetMerchantAccountId(v string) {
	o.MerchantAccountId = &v
}

// GetMethod returns the Method field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionSummary) GetMethod() string {
	if o == nil || o.Method.Get() == nil {
		var ret string
		return ret
	}
	return *o.Method.Get()
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionSummary) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Method.Get(), o.Method.IsSet()
}

// HasMethod returns a boolean if a field has been set.
func (o *TransactionSummary) HasMethod() bool {
	if o != nil && o.Method.IsSet() {
		return true
	}

	return false
}

// SetMethod gets a reference to the given NullableString and assigns it to the Method field.
func (o *TransactionSummary) SetMethod(v string) {
	o.Method.Set(&v)
}
// SetMethodNil sets the value for Method to be an explicit nil
func (o *TransactionSummary) SetMethodNil() {
	o.Method.Set(nil)
}

// UnsetMethod ensures that no value is present for Method, not even an explicit nil
func (o *TransactionSummary) UnsetMethod() {
	o.Method.Unset()
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionSummary) GetPaymentMethod() PaymentMethodSnapshot {
	if o == nil || o.PaymentMethod.Get() == nil {
		var ret PaymentMethodSnapshot
		return ret
	}
	return *o.PaymentMethod.Get()
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionSummary) GetPaymentMethodOk() (*PaymentMethodSnapshot, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentMethod.Get(), o.PaymentMethod.IsSet()
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *TransactionSummary) HasPaymentMethod() bool {
	if o != nil && o.PaymentMethod.IsSet() {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given NullablePaymentMethodSnapshot and assigns it to the PaymentMethod field.
func (o *TransactionSummary) SetPaymentMethod(v PaymentMethodSnapshot) {
	o.PaymentMethod.Set(&v)
}
// SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil
func (o *TransactionSummary) SetPaymentMethodNil() {
	o.PaymentMethod.Set(nil)
}

// UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil
func (o *TransactionSummary) UnsetPaymentMethod() {
	o.PaymentMethod.Unset()
}

// GetPaymentService returns the PaymentService field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionSummary) GetPaymentService() PaymentServiceSnapshot {
	if o == nil || o.PaymentService.Get() == nil {
		var ret PaymentServiceSnapshot
		return ret
	}
	return *o.PaymentService.Get()
}

// GetPaymentServiceOk returns a tuple with the PaymentService field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionSummary) GetPaymentServiceOk() (*PaymentServiceSnapshot, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentService.Get(), o.PaymentService.IsSet()
}

// HasPaymentService returns a boolean if a field has been set.
func (o *TransactionSummary) HasPaymentService() bool {
	if o != nil && o.PaymentService.IsSet() {
		return true
	}

	return false
}

// SetPaymentService gets a reference to the given NullablePaymentServiceSnapshot and assigns it to the PaymentService field.
func (o *TransactionSummary) SetPaymentService(v PaymentServiceSnapshot) {
	o.PaymentService.Set(&v)
}
// SetPaymentServiceNil sets the value for PaymentService to be an explicit nil
func (o *TransactionSummary) SetPaymentServiceNil() {
	o.PaymentService.Set(nil)
}

// UnsetPaymentService ensures that no value is present for PaymentService, not even an explicit nil
func (o *TransactionSummary) UnsetPaymentService() {
	o.PaymentService.Unset()
}

// GetPendingReview returns the PendingReview field value if set, zero value otherwise.
func (o *TransactionSummary) GetPendingReview() bool {
	if o == nil || o.PendingReview == nil {
		var ret bool
		return ret
	}
	return *o.PendingReview
}

// GetPendingReviewOk returns a tuple with the PendingReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetPendingReviewOk() (*bool, bool) {
	if o == nil || o.PendingReview == nil {
		return nil, false
	}
	return o.PendingReview, true
}

// HasPendingReview returns a boolean if a field has been set.
func (o *TransactionSummary) HasPendingReview() bool {
	if o != nil && o.PendingReview != nil {
		return true
	}

	return false
}

// SetPendingReview gets a reference to the given bool and assigns it to the PendingReview field.
func (o *TransactionSummary) SetPendingReview(v bool) {
	o.PendingReview = &v
}

// GetRawResponseCode returns the RawResponseCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionSummary) GetRawResponseCode() string {
	if o == nil || o.RawResponseCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.RawResponseCode.Get()
}

// GetRawResponseCodeOk returns a tuple with the RawResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionSummary) GetRawResponseCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RawResponseCode.Get(), o.RawResponseCode.IsSet()
}

// HasRawResponseCode returns a boolean if a field has been set.
func (o *TransactionSummary) HasRawResponseCode() bool {
	if o != nil && o.RawResponseCode.IsSet() {
		return true
	}

	return false
}

// SetRawResponseCode gets a reference to the given NullableString and assigns it to the RawResponseCode field.
func (o *TransactionSummary) SetRawResponseCode(v string) {
	o.RawResponseCode.Set(&v)
}
// SetRawResponseCodeNil sets the value for RawResponseCode to be an explicit nil
func (o *TransactionSummary) SetRawResponseCodeNil() {
	o.RawResponseCode.Set(nil)
}

// UnsetRawResponseCode ensures that no value is present for RawResponseCode, not even an explicit nil
func (o *TransactionSummary) UnsetRawResponseCode() {
	o.RawResponseCode.Unset()
}

// GetRawResponseDescription returns the RawResponseDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionSummary) GetRawResponseDescription() string {
	if o == nil || o.RawResponseDescription.Get() == nil {
		var ret string
		return ret
	}
	return *o.RawResponseDescription.Get()
}

// GetRawResponseDescriptionOk returns a tuple with the RawResponseDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionSummary) GetRawResponseDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RawResponseDescription.Get(), o.RawResponseDescription.IsSet()
}

// HasRawResponseDescription returns a boolean if a field has been set.
func (o *TransactionSummary) HasRawResponseDescription() bool {
	if o != nil && o.RawResponseDescription.IsSet() {
		return true
	}

	return false
}

// SetRawResponseDescription gets a reference to the given NullableString and assigns it to the RawResponseDescription field.
func (o *TransactionSummary) SetRawResponseDescription(v string) {
	o.RawResponseDescription.Set(&v)
}
// SetRawResponseDescriptionNil sets the value for RawResponseDescription to be an explicit nil
func (o *TransactionSummary) SetRawResponseDescriptionNil() {
	o.RawResponseDescription.Set(nil)
}

// UnsetRawResponseDescription ensures that no value is present for RawResponseDescription, not even an explicit nil
func (o *TransactionSummary) UnsetRawResponseDescription() {
	o.RawResponseDescription.Unset()
}

// GetReconciliationId returns the ReconciliationId field value if set, zero value otherwise.
func (o *TransactionSummary) GetReconciliationId() string {
	if o == nil || o.ReconciliationId == nil {
		var ret string
		return ret
	}
	return *o.ReconciliationId
}

// GetReconciliationIdOk returns a tuple with the ReconciliationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetReconciliationIdOk() (*string, bool) {
	if o == nil || o.ReconciliationId == nil {
		return nil, false
	}
	return o.ReconciliationId, true
}

// HasReconciliationId returns a boolean if a field has been set.
func (o *TransactionSummary) HasReconciliationId() bool {
	if o != nil && o.ReconciliationId != nil {
		return true
	}

	return false
}

// SetReconciliationId gets a reference to the given string and assigns it to the ReconciliationId field.
func (o *TransactionSummary) SetReconciliationId(v string) {
	o.ReconciliationId = &v
}

// GetRefundedAmount returns the RefundedAmount field value if set, zero value otherwise.
func (o *TransactionSummary) GetRefundedAmount() int32 {
	if o == nil || o.RefundedAmount == nil {
		var ret int32
		return ret
	}
	return *o.RefundedAmount
}

// GetRefundedAmountOk returns a tuple with the RefundedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetRefundedAmountOk() (*int32, bool) {
	if o == nil || o.RefundedAmount == nil {
		return nil, false
	}
	return o.RefundedAmount, true
}

// HasRefundedAmount returns a boolean if a field has been set.
func (o *TransactionSummary) HasRefundedAmount() bool {
	if o != nil && o.RefundedAmount != nil {
		return true
	}

	return false
}

// SetRefundedAmount gets a reference to the given int32 and assigns it to the RefundedAmount field.
func (o *TransactionSummary) SetRefundedAmount(v int32) {
	o.RefundedAmount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TransactionSummary) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TransactionSummary) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TransactionSummary) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TransactionSummary) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TransactionSummary) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *TransactionSummary) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetSettledCurrency returns the SettledCurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionSummary) GetSettledCurrency() string {
	if o == nil || o.SettledCurrency.Get() == nil {
		var ret string
		return ret
	}
	return *o.SettledCurrency.Get()
}

// GetSettledCurrencyOk returns a tuple with the SettledCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionSummary) GetSettledCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SettledCurrency.Get(), o.SettledCurrency.IsSet()
}

// HasSettledCurrency returns a boolean if a field has been set.
func (o *TransactionSummary) HasSettledCurrency() bool {
	if o != nil && o.SettledCurrency.IsSet() {
		return true
	}

	return false
}

// SetSettledCurrency gets a reference to the given NullableString and assigns it to the SettledCurrency field.
func (o *TransactionSummary) SetSettledCurrency(v string) {
	o.SettledCurrency.Set(&v)
}
// SetSettledCurrencyNil sets the value for SettledCurrency to be an explicit nil
func (o *TransactionSummary) SetSettledCurrencyNil() {
	o.SettledCurrency.Set(nil)
}

// UnsetSettledCurrency ensures that no value is present for SettledCurrency, not even an explicit nil
func (o *TransactionSummary) UnsetSettledCurrency() {
	o.SettledCurrency.Unset()
}

// GetSettledAmount returns the SettledAmount field value if set, zero value otherwise.
func (o *TransactionSummary) GetSettledAmount() int32 {
	if o == nil || o.SettledAmount == nil {
		var ret int32
		return ret
	}
	return *o.SettledAmount
}

// GetSettledAmountOk returns a tuple with the SettledAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetSettledAmountOk() (*int32, bool) {
	if o == nil || o.SettledAmount == nil {
		return nil, false
	}
	return o.SettledAmount, true
}

// HasSettledAmount returns a boolean if a field has been set.
func (o *TransactionSummary) HasSettledAmount() bool {
	if o != nil && o.SettledAmount != nil {
		return true
	}

	return false
}

// SetSettledAmount gets a reference to the given int32 and assigns it to the SettledAmount field.
func (o *TransactionSummary) SetSettledAmount(v int32) {
	o.SettledAmount = &v
}

// GetSettled returns the Settled field value if set, zero value otherwise.
func (o *TransactionSummary) GetSettled() bool {
	if o == nil || o.Settled == nil {
		var ret bool
		return ret
	}
	return *o.Settled
}

// GetSettledOk returns a tuple with the Settled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetSettledOk() (*bool, bool) {
	if o == nil || o.Settled == nil {
		return nil, false
	}
	return o.Settled, true
}

// HasSettled returns a boolean if a field has been set.
func (o *TransactionSummary) HasSettled() bool {
	if o != nil && o.Settled != nil {
		return true
	}

	return false
}

// SetSettled gets a reference to the given bool and assigns it to the Settled field.
func (o *TransactionSummary) SetSettled(v bool) {
	o.Settled = &v
}

func (o TransactionSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.AuthorizedAmount != nil {
		toSerialize["authorized_amount"] = o.AuthorizedAmount
	}
	if o.Buyer.IsSet() {
		toSerialize["buyer"] = o.Buyer.Get()
	}
	if o.CapturedAmount != nil {
		toSerialize["captured_amount"] = o.CapturedAmount
	}
	if o.CheckoutSessionId.IsSet() {
		toSerialize["checkout_session_id"] = o.CheckoutSessionId.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.ExternalIdentifier.IsSet() {
		toSerialize["external_identifier"] = o.ExternalIdentifier.Get()
	}
	if o.GiftCardRedemptions != nil {
		toSerialize["gift_card_redemptions"] = o.GiftCardRedemptions
	}
	if o.InstrumentType.IsSet() {
		toSerialize["instrument_type"] = o.InstrumentType.Get()
	}
	if o.Intent != nil {
		toSerialize["intent"] = o.Intent
	}
	if o.MerchantAccountId != nil {
		toSerialize["merchant_account_id"] = o.MerchantAccountId
	}
	if o.Method.IsSet() {
		toSerialize["method"] = o.Method.Get()
	}
	if o.PaymentMethod.IsSet() {
		toSerialize["payment_method"] = o.PaymentMethod.Get()
	}
	if o.PaymentService.IsSet() {
		toSerialize["payment_service"] = o.PaymentService.Get()
	}
	if o.PendingReview != nil {
		toSerialize["pending_review"] = o.PendingReview
	}
	if o.RawResponseCode.IsSet() {
		toSerialize["raw_response_code"] = o.RawResponseCode.Get()
	}
	if o.RawResponseDescription.IsSet() {
		toSerialize["raw_response_description"] = o.RawResponseDescription.Get()
	}
	if o.ReconciliationId != nil {
		toSerialize["reconciliation_id"] = o.ReconciliationId
	}
	if o.RefundedAmount != nil {
		toSerialize["refunded_amount"] = o.RefundedAmount
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.SettledCurrency.IsSet() {
		toSerialize["settled_currency"] = o.SettledCurrency.Get()
	}
	if o.SettledAmount != nil {
		toSerialize["settled_amount"] = o.SettledAmount
	}
	if o.Settled != nil {
		toSerialize["settled"] = o.Settled
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionSummary struct {
	value *TransactionSummary
	isSet bool
}

func (v NullableTransactionSummary) Get() *TransactionSummary {
	return v.value
}

func (v *NullableTransactionSummary) Set(val *TransactionSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionSummary(val *TransactionSummary) *NullableTransactionSummary {
	return &NullableTransactionSummary{value: val, isSet: true}
}

func (v NullableTransactionSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


