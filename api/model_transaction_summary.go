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
	// The ID of the merchant account to which this transaction belongs to.
	MerchantAccountId *string `json:"merchant_account_id,omitempty"`
	// The status of the transaction. The status may change over time as asynchronous processing events occur.
	Status *string `json:"status,omitempty"`
	// The original `intent` used when the transaction was [created](#operation/authorize-new-transaction).
	Intent *string `json:"intent,omitempty"`
	// The authorized amount for this transaction. This can be more than the actual captured amount and part of this amount may be refunded.
	Amount *int32 `json:"amount,omitempty"`
	// The captured amount for this transaction. This can be the total or a portion of the authorized amount.
	CapturedAmount *int32 `json:"captured_amount,omitempty"`
	// The refunded amount for this transaction. This can be the total or a portion of the captured amount.
	RefundedAmount *int32 `json:"refunded_amount,omitempty"`
	// The currency code for this transaction.
	Currency *string `json:"currency,omitempty"`
	// The 2-letter ISO code of the country of the transaction. This is used to filter the payment services that is used to process the transaction. 
	Country NullableString `json:"country,omitempty"`
	// The payment method used for this transaction.
	PaymentMethod *PaymentMethodSnapshot `json:"payment_method,omitempty"`
	// The buyer used for this transaction.
	Buyer NullableBuyerSnapshot `json:"buyer,omitempty"`
	// The date and time when this transaction was created in our system.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// An external identifier that can be used to match the transaction against your own records.
	ExternalIdentifier NullableString `json:"external_identifier,omitempty"`
	// Defines when the transaction was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The payment service used for this transaction.
	PaymentService *PaymentServiceSnapshot `json:"payment_service,omitempty"`
	// Whether a manual review is pending.
	PendingReview *bool `json:"pending_review,omitempty"`
	Method *string `json:"method,omitempty"`
	// This is the response code received from the payment service. This can be set to any value and is not standardized across different payment services.
	RawResponseCode NullableString `json:"raw_response_code,omitempty"`
	// This is the response description received from the payment service. This can be set to any value and is not standardized across different payment services.
	RawResponseDescription NullableString `json:"raw_response_description,omitempty"`
	// The identifier for the checkout session this transaction is associated with.
	CheckoutSessionId *string `json:"checkout_session_id,omitempty"`
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

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *TransactionSummary) GetPaymentMethod() PaymentMethodSnapshot {
	if o == nil || o.PaymentMethod == nil {
		var ret PaymentMethodSnapshot
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetPaymentMethodOk() (*PaymentMethodSnapshot, bool) {
	if o == nil || o.PaymentMethod == nil {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *TransactionSummary) HasPaymentMethod() bool {
	if o != nil && o.PaymentMethod != nil {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given PaymentMethodSnapshot and assigns it to the PaymentMethod field.
func (o *TransactionSummary) SetPaymentMethod(v PaymentMethodSnapshot) {
	o.PaymentMethod = &v
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

// GetPaymentService returns the PaymentService field value if set, zero value otherwise.
func (o *TransactionSummary) GetPaymentService() PaymentServiceSnapshot {
	if o == nil || o.PaymentService == nil {
		var ret PaymentServiceSnapshot
		return ret
	}
	return *o.PaymentService
}

// GetPaymentServiceOk returns a tuple with the PaymentService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetPaymentServiceOk() (*PaymentServiceSnapshot, bool) {
	if o == nil || o.PaymentService == nil {
		return nil, false
	}
	return o.PaymentService, true
}

// HasPaymentService returns a boolean if a field has been set.
func (o *TransactionSummary) HasPaymentService() bool {
	if o != nil && o.PaymentService != nil {
		return true
	}

	return false
}

// SetPaymentService gets a reference to the given PaymentServiceSnapshot and assigns it to the PaymentService field.
func (o *TransactionSummary) SetPaymentService(v PaymentServiceSnapshot) {
	o.PaymentService = &v
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

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *TransactionSummary) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *TransactionSummary) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *TransactionSummary) SetMethod(v string) {
	o.Method = &v
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

// GetCheckoutSessionId returns the CheckoutSessionId field value if set, zero value otherwise.
func (o *TransactionSummary) GetCheckoutSessionId() string {
	if o == nil || o.CheckoutSessionId == nil {
		var ret string
		return ret
	}
	return *o.CheckoutSessionId
}

// GetCheckoutSessionIdOk returns a tuple with the CheckoutSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionSummary) GetCheckoutSessionIdOk() (*string, bool) {
	if o == nil || o.CheckoutSessionId == nil {
		return nil, false
	}
	return o.CheckoutSessionId, true
}

// HasCheckoutSessionId returns a boolean if a field has been set.
func (o *TransactionSummary) HasCheckoutSessionId() bool {
	if o != nil && o.CheckoutSessionId != nil {
		return true
	}

	return false
}

// SetCheckoutSessionId gets a reference to the given string and assigns it to the CheckoutSessionId field.
func (o *TransactionSummary) SetCheckoutSessionId(v string) {
	o.CheckoutSessionId = &v
}

func (o TransactionSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MerchantAccountId != nil {
		toSerialize["merchant_account_id"] = o.MerchantAccountId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Intent != nil {
		toSerialize["intent"] = o.Intent
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.CapturedAmount != nil {
		toSerialize["captured_amount"] = o.CapturedAmount
	}
	if o.RefundedAmount != nil {
		toSerialize["refunded_amount"] = o.RefundedAmount
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if o.PaymentMethod != nil {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if o.Buyer.IsSet() {
		toSerialize["buyer"] = o.Buyer.Get()
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ExternalIdentifier.IsSet() {
		toSerialize["external_identifier"] = o.ExternalIdentifier.Get()
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.PaymentService != nil {
		toSerialize["payment_service"] = o.PaymentService
	}
	if o.PendingReview != nil {
		toSerialize["pending_review"] = o.PendingReview
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.RawResponseCode.IsSet() {
		toSerialize["raw_response_code"] = o.RawResponseCode.Get()
	}
	if o.RawResponseDescription.IsSet() {
		toSerialize["raw_response_description"] = o.RawResponseDescription.Get()
	}
	if o.CheckoutSessionId != nil {
		toSerialize["checkout_session_id"] = o.CheckoutSessionId
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


