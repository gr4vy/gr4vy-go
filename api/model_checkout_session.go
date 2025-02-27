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

// CheckoutSession A short-lived checkout session.
type CheckoutSession struct {
	// `checkout-session`.
	Type *string `json:"type,omitempty"`
	// The ID of the Checkout Session.
	Id *string `json:"id,omitempty"`
	// The date and time when the Checkout Session will expire. By default this will be set to 1 hour from the date of creation.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// An array of cart items that represents the line items of a transaction.
	CartItems []CartItem `json:"cart_items,omitempty"`
	// Any additional information about the transaction that you would like to store as key-value pairs. This data is passed to payment service providers that support it.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Contains information about an airline travel, if applicable.
	Airline NullableAirline `json:"airline,omitempty"`
	PaymentMethod NullableCheckoutSessionPaymentMethod `json:"payment_method,omitempty"`
}

// NewCheckoutSession instantiates a new CheckoutSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutSession() *CheckoutSession {
	this := CheckoutSession{}
	return &this
}

// NewCheckoutSessionWithDefaults instantiates a new CheckoutSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutSessionWithDefaults() *CheckoutSession {
	this := CheckoutSession{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CheckoutSession) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CheckoutSession) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CheckoutSession) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CheckoutSession) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CheckoutSession) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CheckoutSession) SetId(v string) {
	o.Id = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *CheckoutSession) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CheckoutSession) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *CheckoutSession) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetCartItems returns the CartItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckoutSession) GetCartItems() []CartItem {
	if o == nil  {
		var ret []CartItem
		return ret
	}
	return o.CartItems
}

// GetCartItemsOk returns a tuple with the CartItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSession) GetCartItemsOk() (*[]CartItem, bool) {
	if o == nil || o.CartItems == nil {
		return nil, false
	}
	return &o.CartItems, true
}

// HasCartItems returns a boolean if a field has been set.
func (o *CheckoutSession) HasCartItems() bool {
	if o != nil && o.CartItems != nil {
		return true
	}

	return false
}

// SetCartItems gets a reference to the given []CartItem and assigns it to the CartItems field.
func (o *CheckoutSession) SetCartItems(v []CartItem) {
	o.CartItems = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckoutSession) GetMetadata() map[string]string {
	if o == nil  {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSession) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CheckoutSession) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CheckoutSession) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetAirline returns the Airline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckoutSession) GetAirline() Airline {
	if o == nil || o.Airline.Get() == nil {
		var ret Airline
		return ret
	}
	return *o.Airline.Get()
}

// GetAirlineOk returns a tuple with the Airline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSession) GetAirlineOk() (*Airline, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Airline.Get(), o.Airline.IsSet()
}

// HasAirline returns a boolean if a field has been set.
func (o *CheckoutSession) HasAirline() bool {
	if o != nil && o.Airline.IsSet() {
		return true
	}

	return false
}

// SetAirline gets a reference to the given NullableAirline and assigns it to the Airline field.
func (o *CheckoutSession) SetAirline(v Airline) {
	o.Airline.Set(&v)
}
// SetAirlineNil sets the value for Airline to be an explicit nil
func (o *CheckoutSession) SetAirlineNil() {
	o.Airline.Set(nil)
}

// UnsetAirline ensures that no value is present for Airline, not even an explicit nil
func (o *CheckoutSession) UnsetAirline() {
	o.Airline.Unset()
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckoutSession) GetPaymentMethod() CheckoutSessionPaymentMethod {
	if o == nil || o.PaymentMethod.Get() == nil {
		var ret CheckoutSessionPaymentMethod
		return ret
	}
	return *o.PaymentMethod.Get()
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSession) GetPaymentMethodOk() (*CheckoutSessionPaymentMethod, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentMethod.Get(), o.PaymentMethod.IsSet()
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *CheckoutSession) HasPaymentMethod() bool {
	if o != nil && o.PaymentMethod.IsSet() {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given NullableCheckoutSessionPaymentMethod and assigns it to the PaymentMethod field.
func (o *CheckoutSession) SetPaymentMethod(v CheckoutSessionPaymentMethod) {
	o.PaymentMethod.Set(&v)
}
// SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil
func (o *CheckoutSession) SetPaymentMethodNil() {
	o.PaymentMethod.Set(nil)
}

// UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil
func (o *CheckoutSession) UnsetPaymentMethod() {
	o.PaymentMethod.Unset()
}

func (o CheckoutSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.CartItems != nil {
		toSerialize["cart_items"] = o.CartItems
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Airline.IsSet() {
		toSerialize["airline"] = o.Airline.Get()
	}
	if o.PaymentMethod.IsSet() {
		toSerialize["payment_method"] = o.PaymentMethod.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCheckoutSession struct {
	value *CheckoutSession
	isSet bool
}

func (v NullableCheckoutSession) Get() *CheckoutSession {
	return v.value
}

func (v *NullableCheckoutSession) Set(val *CheckoutSession) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSession) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSession(val *CheckoutSession) *NullableCheckoutSession {
	return &NullableCheckoutSession{value: val, isSet: true}
}

func (v NullableCheckoutSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


