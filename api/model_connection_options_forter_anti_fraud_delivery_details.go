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

// ConnectionOptionsForterAntiFraudDeliveryDetails General data regarding item such as name, price, etc.
type ConnectionOptionsForterAntiFraudDeliveryDetails struct {
	// Value to populate the `deliveryType` field for this cart item. This overrides the type set at the wider level.  Represents the type of delivery. This can be set to `PHYSICAL` for any type of shipped goods, `DIGITAL` for non-shipped goods (services, gift cards etc.), or `HYBRID` for others.
	DeliveryType NullableString `json:"delivery_type,omitempty"`
	// Value to populate the `deliveryMethod` field for this cart item. This overrides the method set at the wider level.  Represents the delivery method chosen by customer such as postal service, email, in game transfer, etc.
	DeliveryMethod *string `json:"delivery_method,omitempty"`
}

// NewConnectionOptionsForterAntiFraudDeliveryDetails instantiates a new ConnectionOptionsForterAntiFraudDeliveryDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionOptionsForterAntiFraudDeliveryDetails() *ConnectionOptionsForterAntiFraudDeliveryDetails {
	this := ConnectionOptionsForterAntiFraudDeliveryDetails{}
	return &this
}

// NewConnectionOptionsForterAntiFraudDeliveryDetailsWithDefaults instantiates a new ConnectionOptionsForterAntiFraudDeliveryDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionOptionsForterAntiFraudDeliveryDetailsWithDefaults() *ConnectionOptionsForterAntiFraudDeliveryDetails {
	this := ConnectionOptionsForterAntiFraudDeliveryDetails{}
	return &this
}

// GetDeliveryType returns the DeliveryType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionOptionsForterAntiFraudDeliveryDetails) GetDeliveryType() string {
	if o == nil || o.DeliveryType.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeliveryType.Get()
}

// GetDeliveryTypeOk returns a tuple with the DeliveryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionOptionsForterAntiFraudDeliveryDetails) GetDeliveryTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeliveryType.Get(), o.DeliveryType.IsSet()
}

// HasDeliveryType returns a boolean if a field has been set.
func (o *ConnectionOptionsForterAntiFraudDeliveryDetails) HasDeliveryType() bool {
	if o != nil && o.DeliveryType.IsSet() {
		return true
	}

	return false
}

// SetDeliveryType gets a reference to the given NullableString and assigns it to the DeliveryType field.
func (o *ConnectionOptionsForterAntiFraudDeliveryDetails) SetDeliveryType(v string) {
	o.DeliveryType.Set(&v)
}
// SetDeliveryTypeNil sets the value for DeliveryType to be an explicit nil
func (o *ConnectionOptionsForterAntiFraudDeliveryDetails) SetDeliveryTypeNil() {
	o.DeliveryType.Set(nil)
}

// UnsetDeliveryType ensures that no value is present for DeliveryType, not even an explicit nil
func (o *ConnectionOptionsForterAntiFraudDeliveryDetails) UnsetDeliveryType() {
	o.DeliveryType.Unset()
}

// GetDeliveryMethod returns the DeliveryMethod field value if set, zero value otherwise.
func (o *ConnectionOptionsForterAntiFraudDeliveryDetails) GetDeliveryMethod() string {
	if o == nil || o.DeliveryMethod == nil {
		var ret string
		return ret
	}
	return *o.DeliveryMethod
}

// GetDeliveryMethodOk returns a tuple with the DeliveryMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionOptionsForterAntiFraudDeliveryDetails) GetDeliveryMethodOk() (*string, bool) {
	if o == nil || o.DeliveryMethod == nil {
		return nil, false
	}
	return o.DeliveryMethod, true
}

// HasDeliveryMethod returns a boolean if a field has been set.
func (o *ConnectionOptionsForterAntiFraudDeliveryDetails) HasDeliveryMethod() bool {
	if o != nil && o.DeliveryMethod != nil {
		return true
	}

	return false
}

// SetDeliveryMethod gets a reference to the given string and assigns it to the DeliveryMethod field.
func (o *ConnectionOptionsForterAntiFraudDeliveryDetails) SetDeliveryMethod(v string) {
	o.DeliveryMethod = &v
}

func (o ConnectionOptionsForterAntiFraudDeliveryDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeliveryType.IsSet() {
		toSerialize["delivery_type"] = o.DeliveryType.Get()
	}
	if o.DeliveryMethod != nil {
		toSerialize["delivery_method"] = o.DeliveryMethod
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionOptionsForterAntiFraudDeliveryDetails struct {
	value *ConnectionOptionsForterAntiFraudDeliveryDetails
	isSet bool
}

func (v NullableConnectionOptionsForterAntiFraudDeliveryDetails) Get() *ConnectionOptionsForterAntiFraudDeliveryDetails {
	return v.value
}

func (v *NullableConnectionOptionsForterAntiFraudDeliveryDetails) Set(val *ConnectionOptionsForterAntiFraudDeliveryDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionOptionsForterAntiFraudDeliveryDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionOptionsForterAntiFraudDeliveryDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionOptionsForterAntiFraudDeliveryDetails(val *ConnectionOptionsForterAntiFraudDeliveryDetails) *NullableConnectionOptionsForterAntiFraudDeliveryDetails {
	return &NullableConnectionOptionsForterAntiFraudDeliveryDetails{value: val, isSet: true}
}

func (v NullableConnectionOptionsForterAntiFraudDeliveryDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionOptionsForterAntiFraudDeliveryDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


