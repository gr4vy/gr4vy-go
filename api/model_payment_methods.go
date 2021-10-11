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

// PaymentMethods A list of stored payment methods.
type PaymentMethods struct {
	// A list of stored payment methods.
	Items *[]PaymentMethod `json:"items,omitempty"`
	// The limit applied to request. This represents the number of items that are at maximum returned by this request.
	Limit *int32 `json:"limit,omitempty"`
	// The cursor that represents the next page of results. Use the `cursor` query parameter to fetch this page of items.
	NextCursor NullableString `json:"next_cursor,omitempty"`
	// The cursor that represents the next page of results. Use the `cursor` query parameter to fetch this page of items.
	PreviousCursor NullableString `json:"previous_cursor,omitempty"`
}

// NewPaymentMethods instantiates a new PaymentMethods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethods() *PaymentMethods {
	this := PaymentMethods{}
	var limit int32 = 20
	this.Limit = &limit
	return &this
}

// NewPaymentMethodsWithDefaults instantiates a new PaymentMethods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodsWithDefaults() *PaymentMethods {
	this := PaymentMethods{}
	var limit int32 = 20
	this.Limit = &limit
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PaymentMethods) GetItems() []PaymentMethod {
	if o == nil || o.Items == nil {
		var ret []PaymentMethod
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethods) GetItemsOk() (*[]PaymentMethod, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PaymentMethods) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []PaymentMethod and assigns it to the Items field.
func (o *PaymentMethods) SetItems(v []PaymentMethod) {
	o.Items = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PaymentMethods) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethods) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PaymentMethods) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PaymentMethods) SetLimit(v int32) {
	o.Limit = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethods) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}
	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethods) GetNextCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// HasNextCursor returns a boolean if a field has been set.
func (o *PaymentMethods) HasNextCursor() bool {
	if o != nil && o.NextCursor.IsSet() {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given NullableString and assigns it to the NextCursor field.
func (o *PaymentMethods) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}
// SetNextCursorNil sets the value for NextCursor to be an explicit nil
func (o *PaymentMethods) SetNextCursorNil() {
	o.NextCursor.Set(nil)
}

// UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
func (o *PaymentMethods) UnsetNextCursor() {
	o.NextCursor.Unset()
}

// GetPreviousCursor returns the PreviousCursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethods) GetPreviousCursor() string {
	if o == nil || o.PreviousCursor.Get() == nil {
		var ret string
		return ret
	}
	return *o.PreviousCursor.Get()
}

// GetPreviousCursorOk returns a tuple with the PreviousCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethods) GetPreviousCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreviousCursor.Get(), o.PreviousCursor.IsSet()
}

// HasPreviousCursor returns a boolean if a field has been set.
func (o *PaymentMethods) HasPreviousCursor() bool {
	if o != nil && o.PreviousCursor.IsSet() {
		return true
	}

	return false
}

// SetPreviousCursor gets a reference to the given NullableString and assigns it to the PreviousCursor field.
func (o *PaymentMethods) SetPreviousCursor(v string) {
	o.PreviousCursor.Set(&v)
}
// SetPreviousCursorNil sets the value for PreviousCursor to be an explicit nil
func (o *PaymentMethods) SetPreviousCursorNil() {
	o.PreviousCursor.Set(nil)
}

// UnsetPreviousCursor ensures that no value is present for PreviousCursor, not even an explicit nil
func (o *PaymentMethods) UnsetPreviousCursor() {
	o.PreviousCursor.Unset()
}

func (o PaymentMethods) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.NextCursor.IsSet() {
		toSerialize["next_cursor"] = o.NextCursor.Get()
	}
	if o.PreviousCursor.IsSet() {
		toSerialize["previous_cursor"] = o.PreviousCursor.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentMethods struct {
	value *PaymentMethods
	isSet bool
}

func (v NullablePaymentMethods) Get() *PaymentMethods {
	return v.value
}

func (v *NullablePaymentMethods) Set(val *PaymentMethods) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethods) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethods(val *PaymentMethods) *NullablePaymentMethods {
	return &NullablePaymentMethods{value: val, isSet: true}
}

func (v NullablePaymentMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


