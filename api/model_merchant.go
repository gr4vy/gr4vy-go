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

// Merchant struct for Merchant
type Merchant struct {
	// The type of this resource. Is always `merchant`.
	Type *string `json:"type,omitempty"`
	// The name of the merchant.
	Name *string `json:"name,omitempty"`
	// Unique value which identifies a merchant for processing transactions, also known as a MID.
	IdentificationNumber *string `json:"identification_number,omitempty"`
	// The phone number for the merchant which should be formatted according to the [E164 number standard](https://www.twilio.com/docs/glossary/what-e164).
	PhoneNumber *string `json:"phone_number,omitempty"`
	// Merchant website URL.
	Url *string `json:"url,omitempty"`
	// Value to explain charges or payments on bank statements.
	StatementDescriptor *string `json:"statement_descriptor,omitempty"`
	// Merchant classification for the type of goods or services it provides.
	MerchantCategoryCode *string `json:"merchant_category_code,omitempty"`
	// The address for the merchant.
	Address NullableAddress `json:"address,omitempty"`
}

// NewMerchant instantiates a new Merchant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchant() *Merchant {
	this := Merchant{}
	return &this
}

// NewMerchantWithDefaults instantiates a new Merchant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantWithDefaults() *Merchant {
	this := Merchant{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Merchant) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Merchant) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Merchant) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Merchant) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Merchant) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Merchant) SetName(v string) {
	o.Name = &v
}

// GetIdentificationNumber returns the IdentificationNumber field value if set, zero value otherwise.
func (o *Merchant) GetIdentificationNumber() string {
	if o == nil || o.IdentificationNumber == nil {
		var ret string
		return ret
	}
	return *o.IdentificationNumber
}

// GetIdentificationNumberOk returns a tuple with the IdentificationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetIdentificationNumberOk() (*string, bool) {
	if o == nil || o.IdentificationNumber == nil {
		return nil, false
	}
	return o.IdentificationNumber, true
}

// HasIdentificationNumber returns a boolean if a field has been set.
func (o *Merchant) HasIdentificationNumber() bool {
	if o != nil && o.IdentificationNumber != nil {
		return true
	}

	return false
}

// SetIdentificationNumber gets a reference to the given string and assigns it to the IdentificationNumber field.
func (o *Merchant) SetIdentificationNumber(v string) {
	o.IdentificationNumber = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *Merchant) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *Merchant) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *Merchant) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Merchant) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Merchant) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Merchant) SetUrl(v string) {
	o.Url = &v
}

// GetStatementDescriptor returns the StatementDescriptor field value if set, zero value otherwise.
func (o *Merchant) GetStatementDescriptor() string {
	if o == nil || o.StatementDescriptor == nil {
		var ret string
		return ret
	}
	return *o.StatementDescriptor
}

// GetStatementDescriptorOk returns a tuple with the StatementDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetStatementDescriptorOk() (*string, bool) {
	if o == nil || o.StatementDescriptor == nil {
		return nil, false
	}
	return o.StatementDescriptor, true
}

// HasStatementDescriptor returns a boolean if a field has been set.
func (o *Merchant) HasStatementDescriptor() bool {
	if o != nil && o.StatementDescriptor != nil {
		return true
	}

	return false
}

// SetStatementDescriptor gets a reference to the given string and assigns it to the StatementDescriptor field.
func (o *Merchant) SetStatementDescriptor(v string) {
	o.StatementDescriptor = &v
}

// GetMerchantCategoryCode returns the MerchantCategoryCode field value if set, zero value otherwise.
func (o *Merchant) GetMerchantCategoryCode() string {
	if o == nil || o.MerchantCategoryCode == nil {
		var ret string
		return ret
	}
	return *o.MerchantCategoryCode
}

// GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchant) GetMerchantCategoryCodeOk() (*string, bool) {
	if o == nil || o.MerchantCategoryCode == nil {
		return nil, false
	}
	return o.MerchantCategoryCode, true
}

// HasMerchantCategoryCode returns a boolean if a field has been set.
func (o *Merchant) HasMerchantCategoryCode() bool {
	if o != nil && o.MerchantCategoryCode != nil {
		return true
	}

	return false
}

// SetMerchantCategoryCode gets a reference to the given string and assigns it to the MerchantCategoryCode field.
func (o *Merchant) SetMerchantCategoryCode(v string) {
	o.MerchantCategoryCode = &v
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Merchant) GetAddress() Address {
	if o == nil || o.Address.Get() == nil {
		var ret Address
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Merchant) GetAddressOk() (*Address, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *Merchant) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableAddress and assigns it to the Address field.
func (o *Merchant) SetAddress(v Address) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *Merchant) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *Merchant) UnsetAddress() {
	o.Address.Unset()
}

func (o Merchant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.IdentificationNumber != nil {
		toSerialize["identification_number"] = o.IdentificationNumber
	}
	if o.PhoneNumber != nil {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.StatementDescriptor != nil {
		toSerialize["statement_descriptor"] = o.StatementDescriptor
	}
	if o.MerchantCategoryCode != nil {
		toSerialize["merchant_category_code"] = o.MerchantCategoryCode
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMerchant struct {
	value *Merchant
	isSet bool
}

func (v NullableMerchant) Get() *Merchant {
	return v.value
}

func (v *NullableMerchant) Set(val *Merchant) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchant) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchant(val *Merchant) *NullableMerchant {
	return &NullableMerchant{value: val, isSet: true}
}

func (v NullableMerchant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


