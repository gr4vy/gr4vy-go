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
	"time"
)

// PayPal A stored PayPal account.
type PayPal struct {
	// `payment-method`.
	Type *string `json:"type,omitempty"`
	// The unique ID of the payment method.
	Id *string `json:"id,omitempty"`
	// The state of the account tokenization. After the first call this will be set to `buyer_approval_pending` and the response will include an `approval_url`. The buyer needs to be redirected to this URL to authorize the future payments.
	Status *string `json:"status,omitempty"`
	// `paypal`.
	Method *string `json:"method,omitempty"`
	// The date and time when this payment method was first created in our system.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The date and time when this payment method was last updated in our system.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// An external identifier that can be used to match the payment method against your own records.
	ExternalIdentifier NullableString `json:"external_identifier,omitempty"`
	// The optional buyer for which this payment method has been stored.
	Buyer NullableBuyer `json:"buyer,omitempty"`
	Details *PayPalDetails `json:"details,omitempty"`
	// The environment this payment method has been stored for. This will be null of the payment method was not stored.
	Environment NullableString `json:"environment,omitempty"`
}

// NewPayPal instantiates a new PayPal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayPal() *PayPal {
	this := PayPal{}
	var environment string = "production"
	this.Environment = *NewNullableString(&environment)
	return &this
}

// NewPayPalWithDefaults instantiates a new PayPal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayPalWithDefaults() *PayPal {
	this := PayPal{}
	var environment string = "production"
	this.Environment = *NewNullableString(&environment)
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PayPal) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayPal) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PayPal) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PayPal) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PayPal) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayPal) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PayPal) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PayPal) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PayPal) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayPal) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PayPal) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PayPal) SetStatus(v string) {
	o.Status = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *PayPal) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayPal) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *PayPal) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *PayPal) SetMethod(v string) {
	o.Method = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PayPal) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayPal) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PayPal) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PayPal) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PayPal) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayPal) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PayPal) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PayPal) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetExternalIdentifier returns the ExternalIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayPal) GetExternalIdentifier() string {
	if o == nil || o.ExternalIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalIdentifier.Get()
}

// GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayPal) GetExternalIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalIdentifier.Get(), o.ExternalIdentifier.IsSet()
}

// HasExternalIdentifier returns a boolean if a field has been set.
func (o *PayPal) HasExternalIdentifier() bool {
	if o != nil && o.ExternalIdentifier.IsSet() {
		return true
	}

	return false
}

// SetExternalIdentifier gets a reference to the given NullableString and assigns it to the ExternalIdentifier field.
func (o *PayPal) SetExternalIdentifier(v string) {
	o.ExternalIdentifier.Set(&v)
}
// SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil
func (o *PayPal) SetExternalIdentifierNil() {
	o.ExternalIdentifier.Set(nil)
}

// UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
func (o *PayPal) UnsetExternalIdentifier() {
	o.ExternalIdentifier.Unset()
}

// GetBuyer returns the Buyer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayPal) GetBuyer() Buyer {
	if o == nil || o.Buyer.Get() == nil {
		var ret Buyer
		return ret
	}
	return *o.Buyer.Get()
}

// GetBuyerOk returns a tuple with the Buyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayPal) GetBuyerOk() (*Buyer, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Buyer.Get(), o.Buyer.IsSet()
}

// HasBuyer returns a boolean if a field has been set.
func (o *PayPal) HasBuyer() bool {
	if o != nil && o.Buyer.IsSet() {
		return true
	}

	return false
}

// SetBuyer gets a reference to the given NullableBuyer and assigns it to the Buyer field.
func (o *PayPal) SetBuyer(v Buyer) {
	o.Buyer.Set(&v)
}
// SetBuyerNil sets the value for Buyer to be an explicit nil
func (o *PayPal) SetBuyerNil() {
	o.Buyer.Set(nil)
}

// UnsetBuyer ensures that no value is present for Buyer, not even an explicit nil
func (o *PayPal) UnsetBuyer() {
	o.Buyer.Unset()
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *PayPal) GetDetails() PayPalDetails {
	if o == nil || o.Details == nil {
		var ret PayPalDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayPal) GetDetailsOk() (*PayPalDetails, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *PayPal) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given PayPalDetails and assigns it to the Details field.
func (o *PayPal) SetDetails(v PayPalDetails) {
	o.Details = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayPal) GetEnvironment() string {
	if o == nil || o.Environment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayPal) GetEnvironmentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *PayPal) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableString and assigns it to the Environment field.
func (o *PayPal) SetEnvironment(v string) {
	o.Environment.Set(&v)
}
// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *PayPal) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *PayPal) UnsetEnvironment() {
	o.Environment.Unset()
}

func (o PayPal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.ExternalIdentifier.IsSet() {
		toSerialize["external_identifier"] = o.ExternalIdentifier.Get()
	}
	if o.Buyer.IsSet() {
		toSerialize["buyer"] = o.Buyer.Get()
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Environment.IsSet() {
		toSerialize["environment"] = o.Environment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePayPal struct {
	value *PayPal
	isSet bool
}

func (v NullablePayPal) Get() *PayPal {
	return v.value
}

func (v *NullablePayPal) Set(val *PayPal) {
	v.value = val
	v.isSet = true
}

func (v NullablePayPal) IsSet() bool {
	return v.isSet
}

func (v *NullablePayPal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayPal(val *PayPal) *NullablePayPal {
	return &NullablePayPal{value: val, isSet: true}
}

func (v NullablePayPal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayPal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


