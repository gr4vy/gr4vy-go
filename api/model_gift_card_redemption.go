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

// GiftCardRedemption A redemption of a gift card used in a transaction.
type GiftCardRedemption struct {
	// The type of this resource.
	Type *string `json:"type,omitempty"`
	// The ID of this gift card redemption. This may be `null` if the no redemption happened.
	Id NullableString `json:"id,omitempty"`
	// The status of the gift card redemption for the `payment_method`.
	Status *string `json:"status,omitempty"`
	// The amount redeemed for this gift card.
	Amount *int32 `json:"amount,omitempty"`
	// The amount refunded for this gift card. This can not be larger than `amount`.
	RefundedAmount *int32 `json:"refunded_amount,omitempty"`
	// The gift card service's unique ID for the redemption.
	GiftCardServiceRedemptionId NullableString `json:"gift_card_service_redemption_id,omitempty"`
	// If this gift card redemption resulted in an error, this will contain the internal code for the error.
	ErrorCode NullableString `json:"error_code,omitempty"`
	// If this gift card redemption resulted in an error, this will contain the raw error code received from the gift card provider.
	RawErrorCode NullableString `json:"raw_error_code,omitempty"`
	// If this gift card redemption resulted in an error, this will contain the raw error message received from the gift card provider.
	RawErrorMessage NullableString `json:"raw_error_message,omitempty"`
	GiftCard *GiftCardSnapshot `json:"gift_card,omitempty"`
}

// NewGiftCardRedemption instantiates a new GiftCardRedemption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardRedemption() *GiftCardRedemption {
	this := GiftCardRedemption{}
	return &this
}

// NewGiftCardRedemptionWithDefaults instantiates a new GiftCardRedemption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardRedemptionWithDefaults() *GiftCardRedemption {
	this := GiftCardRedemption{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GiftCardRedemption) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardRedemption) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GiftCardRedemption) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GiftCardRedemption) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftCardRedemption) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardRedemption) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *GiftCardRedemption) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *GiftCardRedemption) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *GiftCardRedemption) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *GiftCardRedemption) UnsetId() {
	o.Id.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GiftCardRedemption) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardRedemption) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GiftCardRedemption) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GiftCardRedemption) SetStatus(v string) {
	o.Status = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GiftCardRedemption) GetAmount() int32 {
	if o == nil || o.Amount == nil {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardRedemption) GetAmountOk() (*int32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GiftCardRedemption) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *GiftCardRedemption) SetAmount(v int32) {
	o.Amount = &v
}

// GetRefundedAmount returns the RefundedAmount field value if set, zero value otherwise.
func (o *GiftCardRedemption) GetRefundedAmount() int32 {
	if o == nil || o.RefundedAmount == nil {
		var ret int32
		return ret
	}
	return *o.RefundedAmount
}

// GetRefundedAmountOk returns a tuple with the RefundedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardRedemption) GetRefundedAmountOk() (*int32, bool) {
	if o == nil || o.RefundedAmount == nil {
		return nil, false
	}
	return o.RefundedAmount, true
}

// HasRefundedAmount returns a boolean if a field has been set.
func (o *GiftCardRedemption) HasRefundedAmount() bool {
	if o != nil && o.RefundedAmount != nil {
		return true
	}

	return false
}

// SetRefundedAmount gets a reference to the given int32 and assigns it to the RefundedAmount field.
func (o *GiftCardRedemption) SetRefundedAmount(v int32) {
	o.RefundedAmount = &v
}

// GetGiftCardServiceRedemptionId returns the GiftCardServiceRedemptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftCardRedemption) GetGiftCardServiceRedemptionId() string {
	if o == nil || o.GiftCardServiceRedemptionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.GiftCardServiceRedemptionId.Get()
}

// GetGiftCardServiceRedemptionIdOk returns a tuple with the GiftCardServiceRedemptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardRedemption) GetGiftCardServiceRedemptionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GiftCardServiceRedemptionId.Get(), o.GiftCardServiceRedemptionId.IsSet()
}

// HasGiftCardServiceRedemptionId returns a boolean if a field has been set.
func (o *GiftCardRedemption) HasGiftCardServiceRedemptionId() bool {
	if o != nil && o.GiftCardServiceRedemptionId.IsSet() {
		return true
	}

	return false
}

// SetGiftCardServiceRedemptionId gets a reference to the given NullableString and assigns it to the GiftCardServiceRedemptionId field.
func (o *GiftCardRedemption) SetGiftCardServiceRedemptionId(v string) {
	o.GiftCardServiceRedemptionId.Set(&v)
}
// SetGiftCardServiceRedemptionIdNil sets the value for GiftCardServiceRedemptionId to be an explicit nil
func (o *GiftCardRedemption) SetGiftCardServiceRedemptionIdNil() {
	o.GiftCardServiceRedemptionId.Set(nil)
}

// UnsetGiftCardServiceRedemptionId ensures that no value is present for GiftCardServiceRedemptionId, not even an explicit nil
func (o *GiftCardRedemption) UnsetGiftCardServiceRedemptionId() {
	o.GiftCardServiceRedemptionId.Unset()
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftCardRedemption) GetErrorCode() string {
	if o == nil || o.ErrorCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardRedemption) GetErrorCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *GiftCardRedemption) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableString and assigns it to the ErrorCode field.
func (o *GiftCardRedemption) SetErrorCode(v string) {
	o.ErrorCode.Set(&v)
}
// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *GiftCardRedemption) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *GiftCardRedemption) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetRawErrorCode returns the RawErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftCardRedemption) GetRawErrorCode() string {
	if o == nil || o.RawErrorCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.RawErrorCode.Get()
}

// GetRawErrorCodeOk returns a tuple with the RawErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardRedemption) GetRawErrorCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RawErrorCode.Get(), o.RawErrorCode.IsSet()
}

// HasRawErrorCode returns a boolean if a field has been set.
func (o *GiftCardRedemption) HasRawErrorCode() bool {
	if o != nil && o.RawErrorCode.IsSet() {
		return true
	}

	return false
}

// SetRawErrorCode gets a reference to the given NullableString and assigns it to the RawErrorCode field.
func (o *GiftCardRedemption) SetRawErrorCode(v string) {
	o.RawErrorCode.Set(&v)
}
// SetRawErrorCodeNil sets the value for RawErrorCode to be an explicit nil
func (o *GiftCardRedemption) SetRawErrorCodeNil() {
	o.RawErrorCode.Set(nil)
}

// UnsetRawErrorCode ensures that no value is present for RawErrorCode, not even an explicit nil
func (o *GiftCardRedemption) UnsetRawErrorCode() {
	o.RawErrorCode.Unset()
}

// GetRawErrorMessage returns the RawErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftCardRedemption) GetRawErrorMessage() string {
	if o == nil || o.RawErrorMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.RawErrorMessage.Get()
}

// GetRawErrorMessageOk returns a tuple with the RawErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardRedemption) GetRawErrorMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RawErrorMessage.Get(), o.RawErrorMessage.IsSet()
}

// HasRawErrorMessage returns a boolean if a field has been set.
func (o *GiftCardRedemption) HasRawErrorMessage() bool {
	if o != nil && o.RawErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetRawErrorMessage gets a reference to the given NullableString and assigns it to the RawErrorMessage field.
func (o *GiftCardRedemption) SetRawErrorMessage(v string) {
	o.RawErrorMessage.Set(&v)
}
// SetRawErrorMessageNil sets the value for RawErrorMessage to be an explicit nil
func (o *GiftCardRedemption) SetRawErrorMessageNil() {
	o.RawErrorMessage.Set(nil)
}

// UnsetRawErrorMessage ensures that no value is present for RawErrorMessage, not even an explicit nil
func (o *GiftCardRedemption) UnsetRawErrorMessage() {
	o.RawErrorMessage.Unset()
}

// GetGiftCard returns the GiftCard field value if set, zero value otherwise.
func (o *GiftCardRedemption) GetGiftCard() GiftCardSnapshot {
	if o == nil || o.GiftCard == nil {
		var ret GiftCardSnapshot
		return ret
	}
	return *o.GiftCard
}

// GetGiftCardOk returns a tuple with the GiftCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GiftCardRedemption) GetGiftCardOk() (*GiftCardSnapshot, bool) {
	if o == nil || o.GiftCard == nil {
		return nil, false
	}
	return o.GiftCard, true
}

// HasGiftCard returns a boolean if a field has been set.
func (o *GiftCardRedemption) HasGiftCard() bool {
	if o != nil && o.GiftCard != nil {
		return true
	}

	return false
}

// SetGiftCard gets a reference to the given GiftCardSnapshot and assigns it to the GiftCard field.
func (o *GiftCardRedemption) SetGiftCard(v GiftCardSnapshot) {
	o.GiftCard = &v
}

func (o GiftCardRedemption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.RefundedAmount != nil {
		toSerialize["refunded_amount"] = o.RefundedAmount
	}
	if o.GiftCardServiceRedemptionId.IsSet() {
		toSerialize["gift_card_service_redemption_id"] = o.GiftCardServiceRedemptionId.Get()
	}
	if o.ErrorCode.IsSet() {
		toSerialize["error_code"] = o.ErrorCode.Get()
	}
	if o.RawErrorCode.IsSet() {
		toSerialize["raw_error_code"] = o.RawErrorCode.Get()
	}
	if o.RawErrorMessage.IsSet() {
		toSerialize["raw_error_message"] = o.RawErrorMessage.Get()
	}
	if o.GiftCard != nil {
		toSerialize["gift_card"] = o.GiftCard
	}
	return json.Marshal(toSerialize)
}

type NullableGiftCardRedemption struct {
	value *GiftCardRedemption
	isSet bool
}

func (v NullableGiftCardRedemption) Get() *GiftCardRedemption {
	return v.value
}

func (v *NullableGiftCardRedemption) Set(val *GiftCardRedemption) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardRedemption) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardRedemption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardRedemption(val *GiftCardRedemption) *NullableGiftCardRedemption {
	return &NullableGiftCardRedemption{value: val, isSet: true}
}

func (v NullableGiftCardRedemption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardRedemption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


