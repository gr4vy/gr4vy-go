/*
 * Gr4vy API (Beta)
 *
 * Welcome to the Gr4vy API reference documentation. Our API is still very much a work in product and subject to change.
 *
 * API version: 1.0
 * Contact: code@gr4vy.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Openapi

import (
	"encoding/json"
)

// PaymentServiceDefinition An available payment service that can be configured.
type PaymentServiceDefinition struct {
	// The ID of the payment service. This is the underlying provider followed by a dash followed by the payment method ID.
	Id *string `json:"id,omitempty"`
	// `payment-service-definition`.
	Type *string `json:"type,omitempty"`
	// The display name of this service.
	DisplayName *string `json:"display_name,omitempty"`
	// The ID of the payment method that this services handles.
	Method *string `json:"method,omitempty"`
	// A list of fields that need to be submitted when activating the payment. service.
	Fields *[]PaymentServiceDefinitionFields `json:"fields,omitempty"`
	// A list of three-letter ISO currency codes that this service supports.
	SupportedCurrencies *[]string `json:"supported_currencies,omitempty"`
	// A list of two-letter ISO country codes that this service supports.
	SupportedCountries *[]string `json:"supported_countries,omitempty"`
}

// NewPaymentServiceDefinition instantiates a new PaymentServiceDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentServiceDefinition() *PaymentServiceDefinition {
	this := PaymentServiceDefinition{}
	var type_ string = "payment-service-definition"
	this.Type = &type_
	return &this
}

// NewPaymentServiceDefinitionWithDefaults instantiates a new PaymentServiceDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentServiceDefinitionWithDefaults() *PaymentServiceDefinition {
	this := PaymentServiceDefinition{}
	var type_ string = "payment-service-definition"
	this.Type = &type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaymentServiceDefinition) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceDefinition) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaymentServiceDefinition) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PaymentServiceDefinition) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentServiceDefinition) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceDefinition) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentServiceDefinition) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentServiceDefinition) SetType(v string) {
	o.Type = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PaymentServiceDefinition) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceDefinition) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PaymentServiceDefinition) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PaymentServiceDefinition) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *PaymentServiceDefinition) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceDefinition) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *PaymentServiceDefinition) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *PaymentServiceDefinition) SetMethod(v string) {
	o.Method = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *PaymentServiceDefinition) GetFields() []PaymentServiceDefinitionFields {
	if o == nil || o.Fields == nil {
		var ret []PaymentServiceDefinitionFields
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceDefinition) GetFieldsOk() (*[]PaymentServiceDefinitionFields, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *PaymentServiceDefinition) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []PaymentServiceDefinitionFields and assigns it to the Fields field.
func (o *PaymentServiceDefinition) SetFields(v []PaymentServiceDefinitionFields) {
	o.Fields = &v
}

// GetSupportedCurrencies returns the SupportedCurrencies field value if set, zero value otherwise.
func (o *PaymentServiceDefinition) GetSupportedCurrencies() []string {
	if o == nil || o.SupportedCurrencies == nil {
		var ret []string
		return ret
	}
	return *o.SupportedCurrencies
}

// GetSupportedCurrenciesOk returns a tuple with the SupportedCurrencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceDefinition) GetSupportedCurrenciesOk() (*[]string, bool) {
	if o == nil || o.SupportedCurrencies == nil {
		return nil, false
	}
	return o.SupportedCurrencies, true
}

// HasSupportedCurrencies returns a boolean if a field has been set.
func (o *PaymentServiceDefinition) HasSupportedCurrencies() bool {
	if o != nil && o.SupportedCurrencies != nil {
		return true
	}

	return false
}

// SetSupportedCurrencies gets a reference to the given []string and assigns it to the SupportedCurrencies field.
func (o *PaymentServiceDefinition) SetSupportedCurrencies(v []string) {
	o.SupportedCurrencies = &v
}

// GetSupportedCountries returns the SupportedCountries field value if set, zero value otherwise.
func (o *PaymentServiceDefinition) GetSupportedCountries() []string {
	if o == nil || o.SupportedCountries == nil {
		var ret []string
		return ret
	}
	return *o.SupportedCountries
}

// GetSupportedCountriesOk returns a tuple with the SupportedCountries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceDefinition) GetSupportedCountriesOk() (*[]string, bool) {
	if o == nil || o.SupportedCountries == nil {
		return nil, false
	}
	return o.SupportedCountries, true
}

// HasSupportedCountries returns a boolean if a field has been set.
func (o *PaymentServiceDefinition) HasSupportedCountries() bool {
	if o != nil && o.SupportedCountries != nil {
		return true
	}

	return false
}

// SetSupportedCountries gets a reference to the given []string and assigns it to the SupportedCountries field.
func (o *PaymentServiceDefinition) SetSupportedCountries(v []string) {
	o.SupportedCountries = &v
}

func (o PaymentServiceDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.SupportedCurrencies != nil {
		toSerialize["supported_currencies"] = o.SupportedCurrencies
	}
	if o.SupportedCountries != nil {
		toSerialize["supported_countries"] = o.SupportedCountries
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentServiceDefinition struct {
	value *PaymentServiceDefinition
	isSet bool
}

func (v NullablePaymentServiceDefinition) Get() *PaymentServiceDefinition {
	return v.value
}

func (v *NullablePaymentServiceDefinition) Set(val *PaymentServiceDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentServiceDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentServiceDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentServiceDefinition(val *PaymentServiceDefinition) *NullablePaymentServiceDefinition {
	return &NullablePaymentServiceDefinition{value: val, isSet: true}
}

func (v NullablePaymentServiceDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentServiceDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


