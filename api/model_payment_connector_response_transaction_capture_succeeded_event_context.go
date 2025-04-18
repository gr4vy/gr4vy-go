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

// PaymentConnectorResponseTransactionCaptureSucceededEventContext Additional context for this event.
type PaymentConnectorResponseTransactionCaptureSucceededEventContext struct {
	// The unique ID of the payment service used.
	PaymentServiceId *string `json:"payment_service_id,omitempty"`
	// The display name of the payment service used.
	PaymentServiceDisplayName *string `json:"payment_service_display_name,omitempty"`
	// The payment service definition used.
	PaymentServiceDefinitionId *string `json:"payment_service_definition_id,omitempty"`
	// The external ID of the transaction as set by the payment service.
	PaymentServiceTransactionId NullableString `json:"payment_service_transaction_id,omitempty"`
	// The new status code for the transaction. This is always set to `capture_succeeded`.
	Status *string `json:"status,omitempty"`
	// The type of instrument used for this transaction.
	InstrumentType NullableString `json:"instrument_type,omitempty"`
	// This will always be `null` because the transaction succeeded.
	RetryRule NullableString `json:"retry_rule,omitempty"`
	// This is the response code received from the payment service. This can be set to any value and is not standardized across different payment services.
	RawResponseCode NullableString `json:"raw_response_code,omitempty"`
	// This is the response description received from the payment service. This can be set to any value and is not standardized across different payment services.
	RawResponseDescription NullableString `json:"raw_response_description,omitempty"`
	// The response code received from the payment service for the Address Verification Check (AVS). This code is mapped to a standardized Gr4vy AVS response code.  - `no_match` - neither address or postal code match - `match` - both address and postal code match - `partial_match_address` - address matches but postal code does not - `partial_match_postcode` - postal code matches but address does not - `unavailable ` - AVS is unavailable for card/country  The value of this field can be `null` if the payment service did not provide a response.
	AvsResponseCode NullableString `json:"avs_response_code,omitempty"`
	// The response code received from the payment service for the Card Verification Value (CVV). This code is mapped to a standardized Gr4vy CVV response code.  - `no_match` - the CVV does not match the expected value - `match` - the CVV matches the expected value - `unavailable ` - CVV check unavailable for card our country - `not_provided ` - CVV not provided  The value of this field can be `null` if the payment service did not provide a response.
	CvvResponseCode NullableString `json:"cvv_response_code,omitempty"`
	// The card scheme sent to the connector.
	PaymentMethodScheme NullableString `json:"payment_method_scheme,omitempty"`
}

// NewPaymentConnectorResponseTransactionCaptureSucceededEventContext instantiates a new PaymentConnectorResponseTransactionCaptureSucceededEventContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentConnectorResponseTransactionCaptureSucceededEventContext() *PaymentConnectorResponseTransactionCaptureSucceededEventContext {
	this := PaymentConnectorResponseTransactionCaptureSucceededEventContext{}
	return &this
}

// NewPaymentConnectorResponseTransactionCaptureSucceededEventContextWithDefaults instantiates a new PaymentConnectorResponseTransactionCaptureSucceededEventContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentConnectorResponseTransactionCaptureSucceededEventContextWithDefaults() *PaymentConnectorResponseTransactionCaptureSucceededEventContext {
	this := PaymentConnectorResponseTransactionCaptureSucceededEventContext{}
	return &this
}

// GetPaymentServiceId returns the PaymentServiceId field value if set, zero value otherwise.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetPaymentServiceId() string {
	if o == nil || o.PaymentServiceId == nil {
		var ret string
		return ret
	}
	return *o.PaymentServiceId
}

// GetPaymentServiceIdOk returns a tuple with the PaymentServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetPaymentServiceIdOk() (*string, bool) {
	if o == nil || o.PaymentServiceId == nil {
		return nil, false
	}
	return o.PaymentServiceId, true
}

// HasPaymentServiceId returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasPaymentServiceId() bool {
	if o != nil && o.PaymentServiceId != nil {
		return true
	}

	return false
}

// SetPaymentServiceId gets a reference to the given string and assigns it to the PaymentServiceId field.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetPaymentServiceId(v string) {
	o.PaymentServiceId = &v
}

// GetPaymentServiceDisplayName returns the PaymentServiceDisplayName field value if set, zero value otherwise.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetPaymentServiceDisplayName() string {
	if o == nil || o.PaymentServiceDisplayName == nil {
		var ret string
		return ret
	}
	return *o.PaymentServiceDisplayName
}

// GetPaymentServiceDisplayNameOk returns a tuple with the PaymentServiceDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetPaymentServiceDisplayNameOk() (*string, bool) {
	if o == nil || o.PaymentServiceDisplayName == nil {
		return nil, false
	}
	return o.PaymentServiceDisplayName, true
}

// HasPaymentServiceDisplayName returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasPaymentServiceDisplayName() bool {
	if o != nil && o.PaymentServiceDisplayName != nil {
		return true
	}

	return false
}

// SetPaymentServiceDisplayName gets a reference to the given string and assigns it to the PaymentServiceDisplayName field.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetPaymentServiceDisplayName(v string) {
	o.PaymentServiceDisplayName = &v
}

// GetPaymentServiceDefinitionId returns the PaymentServiceDefinitionId field value if set, zero value otherwise.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetPaymentServiceDefinitionId() string {
	if o == nil || o.PaymentServiceDefinitionId == nil {
		var ret string
		return ret
	}
	return *o.PaymentServiceDefinitionId
}

// GetPaymentServiceDefinitionIdOk returns a tuple with the PaymentServiceDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetPaymentServiceDefinitionIdOk() (*string, bool) {
	if o == nil || o.PaymentServiceDefinitionId == nil {
		return nil, false
	}
	return o.PaymentServiceDefinitionId, true
}

// HasPaymentServiceDefinitionId returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasPaymentServiceDefinitionId() bool {
	if o != nil && o.PaymentServiceDefinitionId != nil {
		return true
	}

	return false
}

// SetPaymentServiceDefinitionId gets a reference to the given string and assigns it to the PaymentServiceDefinitionId field.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetPaymentServiceDefinitionId(v string) {
	o.PaymentServiceDefinitionId = &v
}

// GetPaymentServiceTransactionId returns the PaymentServiceTransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetPaymentServiceTransactionId() string {
	if o == nil || o.PaymentServiceTransactionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PaymentServiceTransactionId.Get()
}

// GetPaymentServiceTransactionIdOk returns a tuple with the PaymentServiceTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetPaymentServiceTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentServiceTransactionId.Get(), o.PaymentServiceTransactionId.IsSet()
}

// HasPaymentServiceTransactionId returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasPaymentServiceTransactionId() bool {
	if o != nil && o.PaymentServiceTransactionId.IsSet() {
		return true
	}

	return false
}

// SetPaymentServiceTransactionId gets a reference to the given NullableString and assigns it to the PaymentServiceTransactionId field.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetPaymentServiceTransactionId(v string) {
	o.PaymentServiceTransactionId.Set(&v)
}
// SetPaymentServiceTransactionIdNil sets the value for PaymentServiceTransactionId to be an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetPaymentServiceTransactionIdNil() {
	o.PaymentServiceTransactionId.Set(nil)
}

// UnsetPaymentServiceTransactionId ensures that no value is present for PaymentServiceTransactionId, not even an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) UnsetPaymentServiceTransactionId() {
	o.PaymentServiceTransactionId.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetStatus(v string) {
	o.Status = &v
}

// GetInstrumentType returns the InstrumentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetInstrumentType() string {
	if o == nil || o.InstrumentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.InstrumentType.Get()
}

// GetInstrumentTypeOk returns a tuple with the InstrumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetInstrumentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InstrumentType.Get(), o.InstrumentType.IsSet()
}

// HasInstrumentType returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasInstrumentType() bool {
	if o != nil && o.InstrumentType.IsSet() {
		return true
	}

	return false
}

// SetInstrumentType gets a reference to the given NullableString and assigns it to the InstrumentType field.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetInstrumentType(v string) {
	o.InstrumentType.Set(&v)
}
// SetInstrumentTypeNil sets the value for InstrumentType to be an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetInstrumentTypeNil() {
	o.InstrumentType.Set(nil)
}

// UnsetInstrumentType ensures that no value is present for InstrumentType, not even an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) UnsetInstrumentType() {
	o.InstrumentType.Unset()
}

// GetRetryRule returns the RetryRule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetRetryRule() string {
	if o == nil || o.RetryRule.Get() == nil {
		var ret string
		return ret
	}
	return *o.RetryRule.Get()
}

// GetRetryRuleOk returns a tuple with the RetryRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetRetryRuleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RetryRule.Get(), o.RetryRule.IsSet()
}

// HasRetryRule returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasRetryRule() bool {
	if o != nil && o.RetryRule.IsSet() {
		return true
	}

	return false
}

// SetRetryRule gets a reference to the given NullableString and assigns it to the RetryRule field.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetRetryRule(v string) {
	o.RetryRule.Set(&v)
}
// SetRetryRuleNil sets the value for RetryRule to be an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetRetryRuleNil() {
	o.RetryRule.Set(nil)
}

// UnsetRetryRule ensures that no value is present for RetryRule, not even an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) UnsetRetryRule() {
	o.RetryRule.Unset()
}

// GetRawResponseCode returns the RawResponseCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetRawResponseCode() string {
	if o == nil || o.RawResponseCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.RawResponseCode.Get()
}

// GetRawResponseCodeOk returns a tuple with the RawResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetRawResponseCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RawResponseCode.Get(), o.RawResponseCode.IsSet()
}

// HasRawResponseCode returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasRawResponseCode() bool {
	if o != nil && o.RawResponseCode.IsSet() {
		return true
	}

	return false
}

// SetRawResponseCode gets a reference to the given NullableString and assigns it to the RawResponseCode field.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetRawResponseCode(v string) {
	o.RawResponseCode.Set(&v)
}
// SetRawResponseCodeNil sets the value for RawResponseCode to be an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetRawResponseCodeNil() {
	o.RawResponseCode.Set(nil)
}

// UnsetRawResponseCode ensures that no value is present for RawResponseCode, not even an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) UnsetRawResponseCode() {
	o.RawResponseCode.Unset()
}

// GetRawResponseDescription returns the RawResponseDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetRawResponseDescription() string {
	if o == nil || o.RawResponseDescription.Get() == nil {
		var ret string
		return ret
	}
	return *o.RawResponseDescription.Get()
}

// GetRawResponseDescriptionOk returns a tuple with the RawResponseDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetRawResponseDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RawResponseDescription.Get(), o.RawResponseDescription.IsSet()
}

// HasRawResponseDescription returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasRawResponseDescription() bool {
	if o != nil && o.RawResponseDescription.IsSet() {
		return true
	}

	return false
}

// SetRawResponseDescription gets a reference to the given NullableString and assigns it to the RawResponseDescription field.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetRawResponseDescription(v string) {
	o.RawResponseDescription.Set(&v)
}
// SetRawResponseDescriptionNil sets the value for RawResponseDescription to be an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetRawResponseDescriptionNil() {
	o.RawResponseDescription.Set(nil)
}

// UnsetRawResponseDescription ensures that no value is present for RawResponseDescription, not even an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) UnsetRawResponseDescription() {
	o.RawResponseDescription.Unset()
}

// GetAvsResponseCode returns the AvsResponseCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetAvsResponseCode() string {
	if o == nil || o.AvsResponseCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.AvsResponseCode.Get()
}

// GetAvsResponseCodeOk returns a tuple with the AvsResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetAvsResponseCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AvsResponseCode.Get(), o.AvsResponseCode.IsSet()
}

// HasAvsResponseCode returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasAvsResponseCode() bool {
	if o != nil && o.AvsResponseCode.IsSet() {
		return true
	}

	return false
}

// SetAvsResponseCode gets a reference to the given NullableString and assigns it to the AvsResponseCode field.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetAvsResponseCode(v string) {
	o.AvsResponseCode.Set(&v)
}
// SetAvsResponseCodeNil sets the value for AvsResponseCode to be an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetAvsResponseCodeNil() {
	o.AvsResponseCode.Set(nil)
}

// UnsetAvsResponseCode ensures that no value is present for AvsResponseCode, not even an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) UnsetAvsResponseCode() {
	o.AvsResponseCode.Unset()
}

// GetCvvResponseCode returns the CvvResponseCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetCvvResponseCode() string {
	if o == nil || o.CvvResponseCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.CvvResponseCode.Get()
}

// GetCvvResponseCodeOk returns a tuple with the CvvResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetCvvResponseCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CvvResponseCode.Get(), o.CvvResponseCode.IsSet()
}

// HasCvvResponseCode returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasCvvResponseCode() bool {
	if o != nil && o.CvvResponseCode.IsSet() {
		return true
	}

	return false
}

// SetCvvResponseCode gets a reference to the given NullableString and assigns it to the CvvResponseCode field.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetCvvResponseCode(v string) {
	o.CvvResponseCode.Set(&v)
}
// SetCvvResponseCodeNil sets the value for CvvResponseCode to be an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetCvvResponseCodeNil() {
	o.CvvResponseCode.Set(nil)
}

// UnsetCvvResponseCode ensures that no value is present for CvvResponseCode, not even an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) UnsetCvvResponseCode() {
	o.CvvResponseCode.Unset()
}

// GetPaymentMethodScheme returns the PaymentMethodScheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetPaymentMethodScheme() string {
	if o == nil || o.PaymentMethodScheme.Get() == nil {
		var ret string
		return ret
	}
	return *o.PaymentMethodScheme.Get()
}

// GetPaymentMethodSchemeOk returns a tuple with the PaymentMethodScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) GetPaymentMethodSchemeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentMethodScheme.Get(), o.PaymentMethodScheme.IsSet()
}

// HasPaymentMethodScheme returns a boolean if a field has been set.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) HasPaymentMethodScheme() bool {
	if o != nil && o.PaymentMethodScheme.IsSet() {
		return true
	}

	return false
}

// SetPaymentMethodScheme gets a reference to the given NullableString and assigns it to the PaymentMethodScheme field.
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetPaymentMethodScheme(v string) {
	o.PaymentMethodScheme.Set(&v)
}
// SetPaymentMethodSchemeNil sets the value for PaymentMethodScheme to be an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) SetPaymentMethodSchemeNil() {
	o.PaymentMethodScheme.Set(nil)
}

// UnsetPaymentMethodScheme ensures that no value is present for PaymentMethodScheme, not even an explicit nil
func (o *PaymentConnectorResponseTransactionCaptureSucceededEventContext) UnsetPaymentMethodScheme() {
	o.PaymentMethodScheme.Unset()
}

func (o PaymentConnectorResponseTransactionCaptureSucceededEventContext) MarshalJSON() ([]byte, error) {
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
	if o.PaymentServiceTransactionId.IsSet() {
		toSerialize["payment_service_transaction_id"] = o.PaymentServiceTransactionId.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.InstrumentType.IsSet() {
		toSerialize["instrument_type"] = o.InstrumentType.Get()
	}
	if o.RetryRule.IsSet() {
		toSerialize["retry_rule"] = o.RetryRule.Get()
	}
	if o.RawResponseCode.IsSet() {
		toSerialize["raw_response_code"] = o.RawResponseCode.Get()
	}
	if o.RawResponseDescription.IsSet() {
		toSerialize["raw_response_description"] = o.RawResponseDescription.Get()
	}
	if o.AvsResponseCode.IsSet() {
		toSerialize["avs_response_code"] = o.AvsResponseCode.Get()
	}
	if o.CvvResponseCode.IsSet() {
		toSerialize["cvv_response_code"] = o.CvvResponseCode.Get()
	}
	if o.PaymentMethodScheme.IsSet() {
		toSerialize["payment_method_scheme"] = o.PaymentMethodScheme.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentConnectorResponseTransactionCaptureSucceededEventContext struct {
	value *PaymentConnectorResponseTransactionCaptureSucceededEventContext
	isSet bool
}

func (v NullablePaymentConnectorResponseTransactionCaptureSucceededEventContext) Get() *PaymentConnectorResponseTransactionCaptureSucceededEventContext {
	return v.value
}

func (v *NullablePaymentConnectorResponseTransactionCaptureSucceededEventContext) Set(val *PaymentConnectorResponseTransactionCaptureSucceededEventContext) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentConnectorResponseTransactionCaptureSucceededEventContext) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentConnectorResponseTransactionCaptureSucceededEventContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentConnectorResponseTransactionCaptureSucceededEventContext(val *PaymentConnectorResponseTransactionCaptureSucceededEventContext) *NullablePaymentConnectorResponseTransactionCaptureSucceededEventContext {
	return &NullablePaymentConnectorResponseTransactionCaptureSucceededEventContext{value: val, isSet: true}
}

func (v NullablePaymentConnectorResponseTransactionCaptureSucceededEventContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentConnectorResponseTransactionCaptureSucceededEventContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


