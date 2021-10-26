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

// PaymentOption An available payment option for a locale.
type PaymentOption struct {
	// `payment-option`.
	Type *string `json:"type,omitempty"`
	// The type of payment method that is available.
	Method *string `json:"method,omitempty"`
	// An icon to display for the payment option.
	IconUrl NullableString `json:"icon_url,omitempty"`
	// The mode of how the payment option should be displayed.
	Mode *string `json:"mode,omitempty"`
	// A label that describes this payment option. This label is returned in the language defined by the `locale` query parameter. The label can be used to display a list of payment options to the buyer in their language.
	Label *string `json:"label,omitempty"`
	// A flag to indicate if storing the payment method is supported.
	CanStorePaymentMethod *bool `json:"can_store_payment_method,omitempty"`
}

// NewPaymentOption instantiates a new PaymentOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentOption() *PaymentOption {
	this := PaymentOption{}
	return &this
}

// NewPaymentOptionWithDefaults instantiates a new PaymentOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentOptionWithDefaults() *PaymentOption {
	this := PaymentOption{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentOption) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentOption) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentOption) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentOption) SetType(v string) {
	o.Type = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *PaymentOption) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentOption) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *PaymentOption) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *PaymentOption) SetMethod(v string) {
	o.Method = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentOption) GetIconUrl() string {
	if o == nil || o.IconUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentOption) GetIconUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *PaymentOption) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *PaymentOption) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}
// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *PaymentOption) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *PaymentOption) UnsetIconUrl() {
	o.IconUrl.Unset()
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *PaymentOption) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentOption) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *PaymentOption) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *PaymentOption) SetMode(v string) {
	o.Mode = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PaymentOption) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentOption) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PaymentOption) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PaymentOption) SetLabel(v string) {
	o.Label = &v
}

// GetCanStorePaymentMethod returns the CanStorePaymentMethod field value if set, zero value otherwise.
func (o *PaymentOption) GetCanStorePaymentMethod() bool {
	if o == nil || o.CanStorePaymentMethod == nil {
		var ret bool
		return ret
	}
	return *o.CanStorePaymentMethod
}

// GetCanStorePaymentMethodOk returns a tuple with the CanStorePaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentOption) GetCanStorePaymentMethodOk() (*bool, bool) {
	if o == nil || o.CanStorePaymentMethod == nil {
		return nil, false
	}
	return o.CanStorePaymentMethod, true
}

// HasCanStorePaymentMethod returns a boolean if a field has been set.
func (o *PaymentOption) HasCanStorePaymentMethod() bool {
	if o != nil && o.CanStorePaymentMethod != nil {
		return true
	}

	return false
}

// SetCanStorePaymentMethod gets a reference to the given bool and assigns it to the CanStorePaymentMethod field.
func (o *PaymentOption) SetCanStorePaymentMethod(v bool) {
	o.CanStorePaymentMethod = &v
}

func (o PaymentOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.IconUrl.IsSet() {
		toSerialize["icon_url"] = o.IconUrl.Get()
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.CanStorePaymentMethod != nil {
		toSerialize["can_store_payment_method"] = o.CanStorePaymentMethod
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentOption struct {
	value *PaymentOption
	isSet bool
}

func (v NullablePaymentOption) Get() *PaymentOption {
	return v.value
}

func (v *NullablePaymentOption) Set(val *PaymentOption) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentOption) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentOption(val *PaymentOption) *NullablePaymentOption {
	return &NullablePaymentOption{value: val, isSet: true}
}

func (v NullablePaymentOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


