/*
Gr4vy API

Welcome to the Gr4vy API reference documentation. Our API is still very much a work in product and subject to change.

API version: 1.1.0-beta
Contact: code@gr4vy.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Openapi

import (
	"encoding/json"
)

// PaymentServices A list of active, configured payment services.
type PaymentServices struct {
	Items *[]PaymentService `json:"items,omitempty"`
	// The limit applied to request. This represents the number of items that are at maximum returned by this request.
	Limit *int32 `json:"limit,omitempty"`
	// The cursor that represents the next page of results. Use the `cursor` query parameter to fetch this page of items.
	NextCursor NullableString `json:"next_cursor,omitempty"`
	// The cursor that represents the next page of results. Use the `cursor` query parameter to fetch this page of items.
	PreviousCursor NullableString `json:"previous_cursor,omitempty"`
}

// NewPaymentServices instantiates a new PaymentServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentServices() *PaymentServices {
	this := PaymentServices{}
	var limit int32 = 20
	this.Limit = &limit
	return &this
}

// NewPaymentServicesWithDefaults instantiates a new PaymentServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentServicesWithDefaults() *PaymentServices {
	this := PaymentServices{}
	var limit int32 = 20
	this.Limit = &limit
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PaymentServices) GetItems() []PaymentService {
	if o == nil || o.Items == nil {
		var ret []PaymentService
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServices) GetItemsOk() (*[]PaymentService, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PaymentServices) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []PaymentService and assigns it to the Items field.
func (o *PaymentServices) SetItems(v []PaymentService) {
	o.Items = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PaymentServices) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServices) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PaymentServices) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PaymentServices) SetLimit(v int32) {
	o.Limit = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServices) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}
	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServices) GetNextCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// HasNextCursor returns a boolean if a field has been set.
func (o *PaymentServices) HasNextCursor() bool {
	if o != nil && o.NextCursor.IsSet() {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given NullableString and assigns it to the NextCursor field.
func (o *PaymentServices) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}
// SetNextCursorNil sets the value for NextCursor to be an explicit nil
func (o *PaymentServices) SetNextCursorNil() {
	o.NextCursor.Set(nil)
}

// UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
func (o *PaymentServices) UnsetNextCursor() {
	o.NextCursor.Unset()
}

// GetPreviousCursor returns the PreviousCursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServices) GetPreviousCursor() string {
	if o == nil || o.PreviousCursor.Get() == nil {
		var ret string
		return ret
	}
	return *o.PreviousCursor.Get()
}

// GetPreviousCursorOk returns a tuple with the PreviousCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServices) GetPreviousCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreviousCursor.Get(), o.PreviousCursor.IsSet()
}

// HasPreviousCursor returns a boolean if a field has been set.
func (o *PaymentServices) HasPreviousCursor() bool {
	if o != nil && o.PreviousCursor.IsSet() {
		return true
	}

	return false
}

// SetPreviousCursor gets a reference to the given NullableString and assigns it to the PreviousCursor field.
func (o *PaymentServices) SetPreviousCursor(v string) {
	o.PreviousCursor.Set(&v)
}
// SetPreviousCursorNil sets the value for PreviousCursor to be an explicit nil
func (o *PaymentServices) SetPreviousCursorNil() {
	o.PreviousCursor.Set(nil)
}

// UnsetPreviousCursor ensures that no value is present for PreviousCursor, not even an explicit nil
func (o *PaymentServices) UnsetPreviousCursor() {
	o.PreviousCursor.Unset()
}

func (o PaymentServices) MarshalJSON() ([]byte, error) {
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

type NullablePaymentServices struct {
	value *PaymentServices
	isSet bool
}

func (v NullablePaymentServices) Get() *PaymentServices {
	return v.value
}

func (v *NullablePaymentServices) Set(val *PaymentServices) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentServices) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentServices(val *PaymentServices) *NullablePaymentServices {
	return &NullablePaymentServices{value: val, isSet: true}
}

func (v NullablePaymentServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


