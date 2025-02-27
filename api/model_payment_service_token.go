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
	"time"
)

// PaymentServiceToken A payment service token.
type PaymentServiceToken struct {
	// The type of this resource.
	Type *string `json:"type,omitempty"`
	// The unique ID of the token.
	Id *string `json:"id,omitempty"`
	// The unique ID of the payment method.
	PaymentMethodId *string `json:"payment_method_id,omitempty"`
	// The unique ID of the payment service.
	PaymentServiceId *string `json:"payment_service_id,omitempty"`
	// The state of the token.  - `processing` - The payment method is still being stored. - `buyer_approval_required` - Storing the payment method requires   the buyer to provide approval. Follow the `approval_url` for next steps. - `succeeded` - The payment method is approved and stored with all   relevant payment services. - `failed` - Storing the payment method did not succeed.
	Status *string `json:"status,omitempty"`
	// The optional URL that the buyer needs to be redirected to to further authorize their payment.
	ApprovalUrl NullableString `json:"approval_url,omitempty"`
	// The token value. Will be present if succeeded.
	Token NullableString `json:"token,omitempty"`
	// The date and time when this token was first created in our system.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The date and time when this token was last updated in our system.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewPaymentServiceToken instantiates a new PaymentServiceToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentServiceToken() *PaymentServiceToken {
	this := PaymentServiceToken{}
	return &this
}

// NewPaymentServiceTokenWithDefaults instantiates a new PaymentServiceToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentServiceTokenWithDefaults() *PaymentServiceToken {
	this := PaymentServiceToken{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentServiceToken) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceToken) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentServiceToken) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentServiceToken) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaymentServiceToken) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceToken) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaymentServiceToken) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PaymentServiceToken) SetId(v string) {
	o.Id = &v
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *PaymentServiceToken) GetPaymentMethodId() string {
	if o == nil || o.PaymentMethodId == nil {
		var ret string
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceToken) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil || o.PaymentMethodId == nil {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *PaymentServiceToken) HasPaymentMethodId() bool {
	if o != nil && o.PaymentMethodId != nil {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given string and assigns it to the PaymentMethodId field.
func (o *PaymentServiceToken) SetPaymentMethodId(v string) {
	o.PaymentMethodId = &v
}

// GetPaymentServiceId returns the PaymentServiceId field value if set, zero value otherwise.
func (o *PaymentServiceToken) GetPaymentServiceId() string {
	if o == nil || o.PaymentServiceId == nil {
		var ret string
		return ret
	}
	return *o.PaymentServiceId
}

// GetPaymentServiceIdOk returns a tuple with the PaymentServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceToken) GetPaymentServiceIdOk() (*string, bool) {
	if o == nil || o.PaymentServiceId == nil {
		return nil, false
	}
	return o.PaymentServiceId, true
}

// HasPaymentServiceId returns a boolean if a field has been set.
func (o *PaymentServiceToken) HasPaymentServiceId() bool {
	if o != nil && o.PaymentServiceId != nil {
		return true
	}

	return false
}

// SetPaymentServiceId gets a reference to the given string and assigns it to the PaymentServiceId field.
func (o *PaymentServiceToken) SetPaymentServiceId(v string) {
	o.PaymentServiceId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PaymentServiceToken) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceToken) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PaymentServiceToken) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PaymentServiceToken) SetStatus(v string) {
	o.Status = &v
}

// GetApprovalUrl returns the ApprovalUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServiceToken) GetApprovalUrl() string {
	if o == nil || o.ApprovalUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApprovalUrl.Get()
}

// GetApprovalUrlOk returns a tuple with the ApprovalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServiceToken) GetApprovalUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApprovalUrl.Get(), o.ApprovalUrl.IsSet()
}

// HasApprovalUrl returns a boolean if a field has been set.
func (o *PaymentServiceToken) HasApprovalUrl() bool {
	if o != nil && o.ApprovalUrl.IsSet() {
		return true
	}

	return false
}

// SetApprovalUrl gets a reference to the given NullableString and assigns it to the ApprovalUrl field.
func (o *PaymentServiceToken) SetApprovalUrl(v string) {
	o.ApprovalUrl.Set(&v)
}
// SetApprovalUrlNil sets the value for ApprovalUrl to be an explicit nil
func (o *PaymentServiceToken) SetApprovalUrlNil() {
	o.ApprovalUrl.Set(nil)
}

// UnsetApprovalUrl ensures that no value is present for ApprovalUrl, not even an explicit nil
func (o *PaymentServiceToken) UnsetApprovalUrl() {
	o.ApprovalUrl.Unset()
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServiceToken) GetToken() string {
	if o == nil || o.Token.Get() == nil {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServiceToken) GetTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *PaymentServiceToken) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *PaymentServiceToken) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *PaymentServiceToken) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *PaymentServiceToken) UnsetToken() {
	o.Token.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PaymentServiceToken) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceToken) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PaymentServiceToken) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PaymentServiceToken) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PaymentServiceToken) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceToken) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PaymentServiceToken) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PaymentServiceToken) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o PaymentServiceToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PaymentMethodId != nil {
		toSerialize["payment_method_id"] = o.PaymentMethodId
	}
	if o.PaymentServiceId != nil {
		toSerialize["payment_service_id"] = o.PaymentServiceId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ApprovalUrl.IsSet() {
		toSerialize["approval_url"] = o.ApprovalUrl.Get()
	}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentServiceToken struct {
	value *PaymentServiceToken
	isSet bool
}

func (v NullablePaymentServiceToken) Get() *PaymentServiceToken {
	return v.value
}

func (v *NullablePaymentServiceToken) Set(val *PaymentServiceToken) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentServiceToken) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentServiceToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentServiceToken(val *PaymentServiceToken) *NullablePaymentServiceToken {
	return &NullablePaymentServiceToken{value: val, isSet: true}
}

func (v NullablePaymentServiceToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentServiceToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


