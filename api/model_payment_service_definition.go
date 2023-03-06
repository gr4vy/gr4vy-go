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
	// The mode of this payment service.
	Mode *string `json:"mode,omitempty"`
	SupportedFeatures *PaymentServiceDefinitionSupportedFeatures `json:"supported_features,omitempty"`
	// An icon to display for the payment service.
	IconUrl NullableString `json:"icon_url,omitempty"`
	Configuration *PaymentServiceDefinitionConfiguration `json:"configuration,omitempty"`
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

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *PaymentServiceDefinition) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceDefinition) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *PaymentServiceDefinition) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *PaymentServiceDefinition) SetMode(v string) {
	o.Mode = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *PaymentServiceDefinition) GetSupportedFeatures() PaymentServiceDefinitionSupportedFeatures {
	if o == nil || o.SupportedFeatures == nil {
		var ret PaymentServiceDefinitionSupportedFeatures
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceDefinition) GetSupportedFeaturesOk() (*PaymentServiceDefinitionSupportedFeatures, bool) {
	if o == nil || o.SupportedFeatures == nil {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *PaymentServiceDefinition) HasSupportedFeatures() bool {
	if o != nil && o.SupportedFeatures != nil {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given PaymentServiceDefinitionSupportedFeatures and assigns it to the SupportedFeatures field.
func (o *PaymentServiceDefinition) SetSupportedFeatures(v PaymentServiceDefinitionSupportedFeatures) {
	o.SupportedFeatures = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServiceDefinition) GetIconUrl() string {
	if o == nil || o.IconUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServiceDefinition) GetIconUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *PaymentServiceDefinition) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *PaymentServiceDefinition) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}
// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *PaymentServiceDefinition) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *PaymentServiceDefinition) UnsetIconUrl() {
	o.IconUrl.Unset()
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *PaymentServiceDefinition) GetConfiguration() PaymentServiceDefinitionConfiguration {
	if o == nil || o.Configuration == nil {
		var ret PaymentServiceDefinitionConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceDefinition) GetConfigurationOk() (*PaymentServiceDefinitionConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *PaymentServiceDefinition) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given PaymentServiceDefinitionConfiguration and assigns it to the Configuration field.
func (o *PaymentServiceDefinition) SetConfiguration(v PaymentServiceDefinitionConfiguration) {
	o.Configuration = &v
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
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.SupportedFeatures != nil {
		toSerialize["supported_features"] = o.SupportedFeatures
	}
	if o.IconUrl.IsSet() {
		toSerialize["icon_url"] = o.IconUrl.Get()
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
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


