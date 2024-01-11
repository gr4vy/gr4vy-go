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

// Refund A refund record.  A refund is always associated with a single transaction, while a transaction can potentially have several refunds.
type Refund struct {
	// The type of this resource. Is always `refund`.
	Type *string `json:"type,omitempty"`
	// The unique ID of the refund.
	Id *string `json:"id,omitempty"`
	// The ID of the transaction associated with this refund.
	TransactionId *string `json:"transaction_id,omitempty"`
	// The payment service's unique ID for the refund.
	PaymentServiceRefundId *string `json:"payment_service_refund_id,omitempty"`
	// The status of the refund. It may change over time as asynchronous processing events occur.  - `processing` - The refund is being processed. - `succeeded` - The refund was successful. - `declined` - The refund was declined by the underlying PSP. - `failed` - The refund could not proceed due to a technical issue. - `voided` - The refund was voided and will not proceed.
	Status *string `json:"status,omitempty"`
	// The currency code for this refund. Will always match that of the associated transaction.
	Currency *string `json:"currency,omitempty"`
	// The amount requested for this refund.
	Amount *int32 `json:"amount,omitempty"`
	// The date and time when this refund was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The date and time when this refund was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The type of the instrument that was refunded.
	TargetType *string `json:"target_type,omitempty"`
	// The optional ID of the instrument that was refunded. This may be `null` if the instrument was not stored.
	TargetId NullableString `json:"target_id,omitempty"`
}

// NewRefund instantiates a new Refund object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefund() *Refund {
	this := Refund{}
	return &this
}

// NewRefundWithDefaults instantiates a new Refund object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundWithDefaults() *Refund {
	this := Refund{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Refund) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Refund) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Refund) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Refund) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Refund) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Refund) SetId(v string) {
	o.Id = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *Refund) GetTransactionId() string {
	if o == nil || o.TransactionId == nil {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetTransactionIdOk() (*string, bool) {
	if o == nil || o.TransactionId == nil {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *Refund) HasTransactionId() bool {
	if o != nil && o.TransactionId != nil {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *Refund) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetPaymentServiceRefundId returns the PaymentServiceRefundId field value if set, zero value otherwise.
func (o *Refund) GetPaymentServiceRefundId() string {
	if o == nil || o.PaymentServiceRefundId == nil {
		var ret string
		return ret
	}
	return *o.PaymentServiceRefundId
}

// GetPaymentServiceRefundIdOk returns a tuple with the PaymentServiceRefundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetPaymentServiceRefundIdOk() (*string, bool) {
	if o == nil || o.PaymentServiceRefundId == nil {
		return nil, false
	}
	return o.PaymentServiceRefundId, true
}

// HasPaymentServiceRefundId returns a boolean if a field has been set.
func (o *Refund) HasPaymentServiceRefundId() bool {
	if o != nil && o.PaymentServiceRefundId != nil {
		return true
	}

	return false
}

// SetPaymentServiceRefundId gets a reference to the given string and assigns it to the PaymentServiceRefundId field.
func (o *Refund) SetPaymentServiceRefundId(v string) {
	o.PaymentServiceRefundId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Refund) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Refund) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Refund) SetStatus(v string) {
	o.Status = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Refund) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Refund) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Refund) SetCurrency(v string) {
	o.Currency = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Refund) GetAmount() int32 {
	if o == nil || o.Amount == nil {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetAmountOk() (*int32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Refund) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *Refund) SetAmount(v int32) {
	o.Amount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Refund) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Refund) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Refund) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Refund) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Refund) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Refund) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *Refund) GetTargetType() string {
	if o == nil || o.TargetType == nil {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetTargetTypeOk() (*string, bool) {
	if o == nil || o.TargetType == nil {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *Refund) HasTargetType() bool {
	if o != nil && o.TargetType != nil {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *Refund) SetTargetType(v string) {
	o.TargetType = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Refund) GetTargetId() string {
	if o == nil || o.TargetId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TargetId.Get()
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Refund) GetTargetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TargetId.Get(), o.TargetId.IsSet()
}

// HasTargetId returns a boolean if a field has been set.
func (o *Refund) HasTargetId() bool {
	if o != nil && o.TargetId.IsSet() {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given NullableString and assigns it to the TargetId field.
func (o *Refund) SetTargetId(v string) {
	o.TargetId.Set(&v)
}
// SetTargetIdNil sets the value for TargetId to be an explicit nil
func (o *Refund) SetTargetIdNil() {
	o.TargetId.Set(nil)
}

// UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
func (o *Refund) UnsetTargetId() {
	o.TargetId.Unset()
}

func (o Refund) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TransactionId != nil {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if o.PaymentServiceRefundId != nil {
		toSerialize["payment_service_refund_id"] = o.PaymentServiceRefundId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.TargetType != nil {
		toSerialize["target_type"] = o.TargetType
	}
	if o.TargetId.IsSet() {
		toSerialize["target_id"] = o.TargetId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRefund struct {
	value *Refund
	isSet bool
}

func (v NullableRefund) Get() *Refund {
	return v.value
}

func (v *NullableRefund) Set(val *Refund) {
	v.value = val
	v.isSet = true
}

func (v NullableRefund) IsSet() bool {
	return v.isSet
}

func (v *NullableRefund) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefund(val *Refund) *NullableRefund {
	return &NullableRefund{value: val, isSet: true}
}

func (v NullableRefund) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefund) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


