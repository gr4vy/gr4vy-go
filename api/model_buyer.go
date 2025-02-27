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

// Buyer struct for Buyer
type Buyer struct {
	// The type of this resource. Is always `buyer`.
	Type *string `json:"type,omitempty"`
	// The unique Gr4vy ID for this buyer.
	Id *string `json:"id,omitempty"`
	// The billing details associated with a buyer.
	BillingDetails NullableBillingDetails `json:"billing_details,omitempty"`
	// The date and time when this buyer was created in our system.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// A unique name for this buyer which is used in the Gr4vy admin panel to give a buyer a human readable name.
	DisplayName NullableString `json:"display_name,omitempty"`
	// An external identifier that can be used to match the buyer against your own records.
	ExternalIdentifier NullableString `json:"external_identifier,omitempty"`
	// The unique ID for a merchant account.
	MerchantAccountId *string `json:"merchant_account_id,omitempty"`
	// The date and time when this buyer was last updated in our system.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The source account number to perform an account funding transaction.
	AccountNumber NullableString `json:"account_number,omitempty"`
}

// NewBuyer instantiates a new Buyer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyer() *Buyer {
	this := Buyer{}
	return &this
}

// NewBuyerWithDefaults instantiates a new Buyer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyerWithDefaults() *Buyer {
	this := Buyer{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Buyer) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Buyer) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Buyer) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Buyer) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Buyer) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Buyer) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Buyer) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Buyer) SetId(v string) {
	o.Id = &v
}

// GetBillingDetails returns the BillingDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Buyer) GetBillingDetails() BillingDetails {
	if o == nil || o.BillingDetails.Get() == nil {
		var ret BillingDetails
		return ret
	}
	return *o.BillingDetails.Get()
}

// GetBillingDetailsOk returns a tuple with the BillingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Buyer) GetBillingDetailsOk() (*BillingDetails, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BillingDetails.Get(), o.BillingDetails.IsSet()
}

// HasBillingDetails returns a boolean if a field has been set.
func (o *Buyer) HasBillingDetails() bool {
	if o != nil && o.BillingDetails.IsSet() {
		return true
	}

	return false
}

// SetBillingDetails gets a reference to the given NullableBillingDetails and assigns it to the BillingDetails field.
func (o *Buyer) SetBillingDetails(v BillingDetails) {
	o.BillingDetails.Set(&v)
}
// SetBillingDetailsNil sets the value for BillingDetails to be an explicit nil
func (o *Buyer) SetBillingDetailsNil() {
	o.BillingDetails.Set(nil)
}

// UnsetBillingDetails ensures that no value is present for BillingDetails, not even an explicit nil
func (o *Buyer) UnsetBillingDetails() {
	o.BillingDetails.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Buyer) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Buyer) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Buyer) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Buyer) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Buyer) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Buyer) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Buyer) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *Buyer) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *Buyer) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *Buyer) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetExternalIdentifier returns the ExternalIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Buyer) GetExternalIdentifier() string {
	if o == nil || o.ExternalIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalIdentifier.Get()
}

// GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Buyer) GetExternalIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalIdentifier.Get(), o.ExternalIdentifier.IsSet()
}

// HasExternalIdentifier returns a boolean if a field has been set.
func (o *Buyer) HasExternalIdentifier() bool {
	if o != nil && o.ExternalIdentifier.IsSet() {
		return true
	}

	return false
}

// SetExternalIdentifier gets a reference to the given NullableString and assigns it to the ExternalIdentifier field.
func (o *Buyer) SetExternalIdentifier(v string) {
	o.ExternalIdentifier.Set(&v)
}
// SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil
func (o *Buyer) SetExternalIdentifierNil() {
	o.ExternalIdentifier.Set(nil)
}

// UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
func (o *Buyer) UnsetExternalIdentifier() {
	o.ExternalIdentifier.Unset()
}

// GetMerchantAccountId returns the MerchantAccountId field value if set, zero value otherwise.
func (o *Buyer) GetMerchantAccountId() string {
	if o == nil || o.MerchantAccountId == nil {
		var ret string
		return ret
	}
	return *o.MerchantAccountId
}

// GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Buyer) GetMerchantAccountIdOk() (*string, bool) {
	if o == nil || o.MerchantAccountId == nil {
		return nil, false
	}
	return o.MerchantAccountId, true
}

// HasMerchantAccountId returns a boolean if a field has been set.
func (o *Buyer) HasMerchantAccountId() bool {
	if o != nil && o.MerchantAccountId != nil {
		return true
	}

	return false
}

// SetMerchantAccountId gets a reference to the given string and assigns it to the MerchantAccountId field.
func (o *Buyer) SetMerchantAccountId(v string) {
	o.MerchantAccountId = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Buyer) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Buyer) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Buyer) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Buyer) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Buyer) GetAccountNumber() string {
	if o == nil || o.AccountNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber.Get()
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Buyer) GetAccountNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountNumber.Get(), o.AccountNumber.IsSet()
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *Buyer) HasAccountNumber() bool {
	if o != nil && o.AccountNumber.IsSet() {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given NullableString and assigns it to the AccountNumber field.
func (o *Buyer) SetAccountNumber(v string) {
	o.AccountNumber.Set(&v)
}
// SetAccountNumberNil sets the value for AccountNumber to be an explicit nil
func (o *Buyer) SetAccountNumberNil() {
	o.AccountNumber.Set(nil)
}

// UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
func (o *Buyer) UnsetAccountNumber() {
	o.AccountNumber.Unset()
}

func (o Buyer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.BillingDetails.IsSet() {
		toSerialize["billing_details"] = o.BillingDetails.Get()
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DisplayName.IsSet() {
		toSerialize["display_name"] = o.DisplayName.Get()
	}
	if o.ExternalIdentifier.IsSet() {
		toSerialize["external_identifier"] = o.ExternalIdentifier.Get()
	}
	if o.MerchantAccountId != nil {
		toSerialize["merchant_account_id"] = o.MerchantAccountId
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.AccountNumber.IsSet() {
		toSerialize["account_number"] = o.AccountNumber.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableBuyer struct {
	value *Buyer
	isSet bool
}

func (v NullableBuyer) Get() *Buyer {
	return v.value
}

func (v *NullableBuyer) Set(val *Buyer) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyer) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyer(val *Buyer) *NullableBuyer {
	return &NullableBuyer{value: val, isSet: true}
}

func (v NullableBuyer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


