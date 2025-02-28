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

// PaymentServiceSession A session request for a payment session. This is an opaque object passed to the payment service.
type PaymentServiceSession struct {
	// The type of this resource.
	Type *string `json:"type,omitempty"`
	// The status of the response.  - `succeeded` - The session was successfully generated. - `failed` - The session could not be generated.
	Status *string `json:"status,omitempty"`
	// A generic error code that may be returned when the session could not be generated.
	Code NullableString `json:"code,omitempty"`
	// The HTTP status code received from the payment service.
	StatusCode NullableFloat32 `json:"status_code,omitempty"`
	// The parsed JSON received from the payment service.
	ResponseBody map[string]map[string]interface{} `json:"response_body,omitempty"`
}

// NewPaymentServiceSession instantiates a new PaymentServiceSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentServiceSession() *PaymentServiceSession {
	this := PaymentServiceSession{}
	return &this
}

// NewPaymentServiceSessionWithDefaults instantiates a new PaymentServiceSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentServiceSessionWithDefaults() *PaymentServiceSession {
	this := PaymentServiceSession{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentServiceSession) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceSession) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentServiceSession) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentServiceSession) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PaymentServiceSession) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceSession) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PaymentServiceSession) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PaymentServiceSession) SetStatus(v string) {
	o.Status = &v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServiceSession) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServiceSession) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *PaymentServiceSession) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *PaymentServiceSession) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *PaymentServiceSession) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *PaymentServiceSession) UnsetCode() {
	o.Code.Unset()
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServiceSession) GetStatusCode() float32 {
	if o == nil || o.StatusCode.Get() == nil {
		var ret float32
		return ret
	}
	return *o.StatusCode.Get()
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServiceSession) GetStatusCodeOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatusCode.Get(), o.StatusCode.IsSet()
}

// HasStatusCode returns a boolean if a field has been set.
func (o *PaymentServiceSession) HasStatusCode() bool {
	if o != nil && o.StatusCode.IsSet() {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given NullableFloat32 and assigns it to the StatusCode field.
func (o *PaymentServiceSession) SetStatusCode(v float32) {
	o.StatusCode.Set(&v)
}
// SetStatusCodeNil sets the value for StatusCode to be an explicit nil
func (o *PaymentServiceSession) SetStatusCodeNil() {
	o.StatusCode.Set(nil)
}

// UnsetStatusCode ensures that no value is present for StatusCode, not even an explicit nil
func (o *PaymentServiceSession) UnsetStatusCode() {
	o.StatusCode.Unset()
}

// GetResponseBody returns the ResponseBody field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServiceSession) GetResponseBody() map[string]map[string]interface{} {
	if o == nil  {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.ResponseBody
}

// GetResponseBodyOk returns a tuple with the ResponseBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServiceSession) GetResponseBodyOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.ResponseBody == nil {
		return nil, false
	}
	return &o.ResponseBody, true
}

// HasResponseBody returns a boolean if a field has been set.
func (o *PaymentServiceSession) HasResponseBody() bool {
	if o != nil && o.ResponseBody != nil {
		return true
	}

	return false
}

// SetResponseBody gets a reference to the given map[string]map[string]interface{} and assigns it to the ResponseBody field.
func (o *PaymentServiceSession) SetResponseBody(v map[string]map[string]interface{}) {
	o.ResponseBody = v
}

func (o PaymentServiceSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.StatusCode.IsSet() {
		toSerialize["status_code"] = o.StatusCode.Get()
	}
	if o.ResponseBody != nil {
		toSerialize["response_body"] = o.ResponseBody
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentServiceSession struct {
	value *PaymentServiceSession
	isSet bool
}

func (v NullablePaymentServiceSession) Get() *PaymentServiceSession {
	return v.value
}

func (v *NullablePaymentServiceSession) Set(val *PaymentServiceSession) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentServiceSession) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentServiceSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentServiceSession(val *PaymentServiceSession) *NullablePaymentServiceSession {
	return &NullablePaymentServiceSession{value: val, isSet: true}
}

func (v NullablePaymentServiceSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentServiceSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


