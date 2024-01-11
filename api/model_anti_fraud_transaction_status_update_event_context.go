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

// AntiFraudTransactionStatusUpdateEventContext Additional context for this event.
type AntiFraudTransactionStatusUpdateEventContext struct {
	// The unique ID of the anti-fraud service used.
	AntiFraudServiceId *string `json:"anti_fraud_service_id,omitempty"`
	// The name of the anti-fraud service used.
	AntiFraudServiceName *string `json:"anti_fraud_service_name,omitempty"`
	// The anti-fraud service definition used.
	AntiFraudServiceDefinitionId *string `json:"anti_fraud_service_definition_id,omitempty"`
	// The HTTP body sent to fetch a decision.
	Request *string `json:"request,omitempty"`
	// The HTTP body received from the anti-fraud provider.
	Response *string `json:"response,omitempty"`
	// The HTTP response status code from the anti-fraud provider.
	ResponseStatusCode *float32 `json:"response_status_code,omitempty"`
}

// NewAntiFraudTransactionStatusUpdateEventContext instantiates a new AntiFraudTransactionStatusUpdateEventContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAntiFraudTransactionStatusUpdateEventContext() *AntiFraudTransactionStatusUpdateEventContext {
	this := AntiFraudTransactionStatusUpdateEventContext{}
	return &this
}

// NewAntiFraudTransactionStatusUpdateEventContextWithDefaults instantiates a new AntiFraudTransactionStatusUpdateEventContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAntiFraudTransactionStatusUpdateEventContextWithDefaults() *AntiFraudTransactionStatusUpdateEventContext {
	this := AntiFraudTransactionStatusUpdateEventContext{}
	return &this
}

// GetAntiFraudServiceId returns the AntiFraudServiceId field value if set, zero value otherwise.
func (o *AntiFraudTransactionStatusUpdateEventContext) GetAntiFraudServiceId() string {
	if o == nil || o.AntiFraudServiceId == nil {
		var ret string
		return ret
	}
	return *o.AntiFraudServiceId
}

// GetAntiFraudServiceIdOk returns a tuple with the AntiFraudServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudTransactionStatusUpdateEventContext) GetAntiFraudServiceIdOk() (*string, bool) {
	if o == nil || o.AntiFraudServiceId == nil {
		return nil, false
	}
	return o.AntiFraudServiceId, true
}

// HasAntiFraudServiceId returns a boolean if a field has been set.
func (o *AntiFraudTransactionStatusUpdateEventContext) HasAntiFraudServiceId() bool {
	if o != nil && o.AntiFraudServiceId != nil {
		return true
	}

	return false
}

// SetAntiFraudServiceId gets a reference to the given string and assigns it to the AntiFraudServiceId field.
func (o *AntiFraudTransactionStatusUpdateEventContext) SetAntiFraudServiceId(v string) {
	o.AntiFraudServiceId = &v
}

// GetAntiFraudServiceName returns the AntiFraudServiceName field value if set, zero value otherwise.
func (o *AntiFraudTransactionStatusUpdateEventContext) GetAntiFraudServiceName() string {
	if o == nil || o.AntiFraudServiceName == nil {
		var ret string
		return ret
	}
	return *o.AntiFraudServiceName
}

// GetAntiFraudServiceNameOk returns a tuple with the AntiFraudServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudTransactionStatusUpdateEventContext) GetAntiFraudServiceNameOk() (*string, bool) {
	if o == nil || o.AntiFraudServiceName == nil {
		return nil, false
	}
	return o.AntiFraudServiceName, true
}

// HasAntiFraudServiceName returns a boolean if a field has been set.
func (o *AntiFraudTransactionStatusUpdateEventContext) HasAntiFraudServiceName() bool {
	if o != nil && o.AntiFraudServiceName != nil {
		return true
	}

	return false
}

// SetAntiFraudServiceName gets a reference to the given string and assigns it to the AntiFraudServiceName field.
func (o *AntiFraudTransactionStatusUpdateEventContext) SetAntiFraudServiceName(v string) {
	o.AntiFraudServiceName = &v
}

// GetAntiFraudServiceDefinitionId returns the AntiFraudServiceDefinitionId field value if set, zero value otherwise.
func (o *AntiFraudTransactionStatusUpdateEventContext) GetAntiFraudServiceDefinitionId() string {
	if o == nil || o.AntiFraudServiceDefinitionId == nil {
		var ret string
		return ret
	}
	return *o.AntiFraudServiceDefinitionId
}

// GetAntiFraudServiceDefinitionIdOk returns a tuple with the AntiFraudServiceDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudTransactionStatusUpdateEventContext) GetAntiFraudServiceDefinitionIdOk() (*string, bool) {
	if o == nil || o.AntiFraudServiceDefinitionId == nil {
		return nil, false
	}
	return o.AntiFraudServiceDefinitionId, true
}

// HasAntiFraudServiceDefinitionId returns a boolean if a field has been set.
func (o *AntiFraudTransactionStatusUpdateEventContext) HasAntiFraudServiceDefinitionId() bool {
	if o != nil && o.AntiFraudServiceDefinitionId != nil {
		return true
	}

	return false
}

// SetAntiFraudServiceDefinitionId gets a reference to the given string and assigns it to the AntiFraudServiceDefinitionId field.
func (o *AntiFraudTransactionStatusUpdateEventContext) SetAntiFraudServiceDefinitionId(v string) {
	o.AntiFraudServiceDefinitionId = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *AntiFraudTransactionStatusUpdateEventContext) GetRequest() string {
	if o == nil || o.Request == nil {
		var ret string
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudTransactionStatusUpdateEventContext) GetRequestOk() (*string, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *AntiFraudTransactionStatusUpdateEventContext) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given string and assigns it to the Request field.
func (o *AntiFraudTransactionStatusUpdateEventContext) SetRequest(v string) {
	o.Request = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *AntiFraudTransactionStatusUpdateEventContext) GetResponse() string {
	if o == nil || o.Response == nil {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudTransactionStatusUpdateEventContext) GetResponseOk() (*string, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *AntiFraudTransactionStatusUpdateEventContext) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *AntiFraudTransactionStatusUpdateEventContext) SetResponse(v string) {
	o.Response = &v
}

// GetResponseStatusCode returns the ResponseStatusCode field value if set, zero value otherwise.
func (o *AntiFraudTransactionStatusUpdateEventContext) GetResponseStatusCode() float32 {
	if o == nil || o.ResponseStatusCode == nil {
		var ret float32
		return ret
	}
	return *o.ResponseStatusCode
}

// GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AntiFraudTransactionStatusUpdateEventContext) GetResponseStatusCodeOk() (*float32, bool) {
	if o == nil || o.ResponseStatusCode == nil {
		return nil, false
	}
	return o.ResponseStatusCode, true
}

// HasResponseStatusCode returns a boolean if a field has been set.
func (o *AntiFraudTransactionStatusUpdateEventContext) HasResponseStatusCode() bool {
	if o != nil && o.ResponseStatusCode != nil {
		return true
	}

	return false
}

// SetResponseStatusCode gets a reference to the given float32 and assigns it to the ResponseStatusCode field.
func (o *AntiFraudTransactionStatusUpdateEventContext) SetResponseStatusCode(v float32) {
	o.ResponseStatusCode = &v
}

func (o AntiFraudTransactionStatusUpdateEventContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AntiFraudServiceId != nil {
		toSerialize["anti_fraud_service_id"] = o.AntiFraudServiceId
	}
	if o.AntiFraudServiceName != nil {
		toSerialize["anti_fraud_service_name"] = o.AntiFraudServiceName
	}
	if o.AntiFraudServiceDefinitionId != nil {
		toSerialize["anti_fraud_service_definition_id"] = o.AntiFraudServiceDefinitionId
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	if o.ResponseStatusCode != nil {
		toSerialize["response_status_code"] = o.ResponseStatusCode
	}
	return json.Marshal(toSerialize)
}

type NullableAntiFraudTransactionStatusUpdateEventContext struct {
	value *AntiFraudTransactionStatusUpdateEventContext
	isSet bool
}

func (v NullableAntiFraudTransactionStatusUpdateEventContext) Get() *AntiFraudTransactionStatusUpdateEventContext {
	return v.value
}

func (v *NullableAntiFraudTransactionStatusUpdateEventContext) Set(val *AntiFraudTransactionStatusUpdateEventContext) {
	v.value = val
	v.isSet = true
}

func (v NullableAntiFraudTransactionStatusUpdateEventContext) IsSet() bool {
	return v.isSet
}

func (v *NullableAntiFraudTransactionStatusUpdateEventContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntiFraudTransactionStatusUpdateEventContext(val *AntiFraudTransactionStatusUpdateEventContext) *NullableAntiFraudTransactionStatusUpdateEventContext {
	return &NullableAntiFraudTransactionStatusUpdateEventContext{value: val, isSet: true}
}

func (v NullableAntiFraudTransactionStatusUpdateEventContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntiFraudTransactionStatusUpdateEventContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


