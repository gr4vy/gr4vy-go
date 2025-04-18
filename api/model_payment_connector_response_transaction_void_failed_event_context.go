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

// PaymentConnectorResponseTransactionVoidFailedEventContext Additional context for this event.
type PaymentConnectorResponseTransactionVoidFailedEventContext struct {
	// The unique ID of the payment service used.
	PaymentServiceId *string `json:"payment_service_id,omitempty"`
	// The display name of the payment service used.
	PaymentServiceDisplayName *string `json:"payment_service_display_name,omitempty"`
	// The payment service definition used.
	PaymentServiceDefinitionId *string `json:"payment_service_definition_id,omitempty"`
}

// NewPaymentConnectorResponseTransactionVoidFailedEventContext instantiates a new PaymentConnectorResponseTransactionVoidFailedEventContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentConnectorResponseTransactionVoidFailedEventContext() *PaymentConnectorResponseTransactionVoidFailedEventContext {
	this := PaymentConnectorResponseTransactionVoidFailedEventContext{}
	return &this
}

// NewPaymentConnectorResponseTransactionVoidFailedEventContextWithDefaults instantiates a new PaymentConnectorResponseTransactionVoidFailedEventContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentConnectorResponseTransactionVoidFailedEventContextWithDefaults() *PaymentConnectorResponseTransactionVoidFailedEventContext {
	this := PaymentConnectorResponseTransactionVoidFailedEventContext{}
	return &this
}

// GetPaymentServiceId returns the PaymentServiceId field value if set, zero value otherwise.
func (o *PaymentConnectorResponseTransactionVoidFailedEventContext) GetPaymentServiceId() string {
	if o == nil || o.PaymentServiceId == nil {
		var ret string
		return ret
	}
	return *o.PaymentServiceId
}

// GetPaymentServiceIdOk returns a tuple with the PaymentServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentConnectorResponseTransactionVoidFailedEventContext) GetPaymentServiceIdOk() (*string, bool) {
	if o == nil || o.PaymentServiceId == nil {
		return nil, false
	}
	return o.PaymentServiceId, true
}

// HasPaymentServiceId returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionVoidFailedEventContext) HasPaymentServiceId() bool {
	if o != nil && o.PaymentServiceId != nil {
		return true
	}

	return false
}

// SetPaymentServiceId gets a reference to the given string and assigns it to the PaymentServiceId field.
func (o *PaymentConnectorResponseTransactionVoidFailedEventContext) SetPaymentServiceId(v string) {
	o.PaymentServiceId = &v
}

// GetPaymentServiceDisplayName returns the PaymentServiceDisplayName field value if set, zero value otherwise.
func (o *PaymentConnectorResponseTransactionVoidFailedEventContext) GetPaymentServiceDisplayName() string {
	if o == nil || o.PaymentServiceDisplayName == nil {
		var ret string
		return ret
	}
	return *o.PaymentServiceDisplayName
}

// GetPaymentServiceDisplayNameOk returns a tuple with the PaymentServiceDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentConnectorResponseTransactionVoidFailedEventContext) GetPaymentServiceDisplayNameOk() (*string, bool) {
	if o == nil || o.PaymentServiceDisplayName == nil {
		return nil, false
	}
	return o.PaymentServiceDisplayName, true
}

// HasPaymentServiceDisplayName returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionVoidFailedEventContext) HasPaymentServiceDisplayName() bool {
	if o != nil && o.PaymentServiceDisplayName != nil {
		return true
	}

	return false
}

// SetPaymentServiceDisplayName gets a reference to the given string and assigns it to the PaymentServiceDisplayName field.
func (o *PaymentConnectorResponseTransactionVoidFailedEventContext) SetPaymentServiceDisplayName(v string) {
	o.PaymentServiceDisplayName = &v
}

// GetPaymentServiceDefinitionId returns the PaymentServiceDefinitionId field value if set, zero value otherwise.
func (o *PaymentConnectorResponseTransactionVoidFailedEventContext) GetPaymentServiceDefinitionId() string {
	if o == nil || o.PaymentServiceDefinitionId == nil {
		var ret string
		return ret
	}
	return *o.PaymentServiceDefinitionId
}

// GetPaymentServiceDefinitionIdOk returns a tuple with the PaymentServiceDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentConnectorResponseTransactionVoidFailedEventContext) GetPaymentServiceDefinitionIdOk() (*string, bool) {
	if o == nil || o.PaymentServiceDefinitionId == nil {
		return nil, false
	}
	return o.PaymentServiceDefinitionId, true
}

// HasPaymentServiceDefinitionId returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionVoidFailedEventContext) HasPaymentServiceDefinitionId() bool {
	if o != nil && o.PaymentServiceDefinitionId != nil {
		return true
	}

	return false
}

// SetPaymentServiceDefinitionId gets a reference to the given string and assigns it to the PaymentServiceDefinitionId field.
func (o *PaymentConnectorResponseTransactionVoidFailedEventContext) SetPaymentServiceDefinitionId(v string) {
	o.PaymentServiceDefinitionId = &v
}

func (o PaymentConnectorResponseTransactionVoidFailedEventContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentServiceId != nil {
		toSerialize["payment_service_id"] = o.PaymentServiceId
	}
	if o.PaymentServiceDisplayName != nil {
		toSerialize["payment_service_display_name"] = o.PaymentServiceDisplayName
	}
	if o.PaymentServiceDefinitionId != nil {
		toSerialize["payment_service_definition_id"] = o.PaymentServiceDefinitionId
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentConnectorResponseTransactionVoidFailedEventContext struct {
	value *PaymentConnectorResponseTransactionVoidFailedEventContext
	isSet bool
}

func (v NullablePaymentConnectorResponseTransactionVoidFailedEventContext) Get() *PaymentConnectorResponseTransactionVoidFailedEventContext {
	return v.value
}

func (v *NullablePaymentConnectorResponseTransactionVoidFailedEventContext) Set(val *PaymentConnectorResponseTransactionVoidFailedEventContext) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentConnectorResponseTransactionVoidFailedEventContext) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentConnectorResponseTransactionVoidFailedEventContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentConnectorResponseTransactionVoidFailedEventContext(val *PaymentConnectorResponseTransactionVoidFailedEventContext) *NullablePaymentConnectorResponseTransactionVoidFailedEventContext {
	return &NullablePaymentConnectorResponseTransactionVoidFailedEventContext{value: val, isSet: true}
}

func (v NullablePaymentConnectorResponseTransactionVoidFailedEventContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentConnectorResponseTransactionVoidFailedEventContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


