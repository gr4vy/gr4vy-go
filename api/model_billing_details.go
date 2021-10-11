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

// BillingDetails Billing details to use associated to a buyer.
type BillingDetails struct {
	// The first name(s) or given name for the buyer.
	FirstName NullableString `json:"first_name,omitempty"`
	// The last name, or family name, of the buyer.
	LastName NullableString `json:"last_name,omitempty"`
	// The email address for the buyer.
	EmailAddress NullableString `json:"email_address,omitempty"`
	// The phone number to use for this request. This expect the number in the [E164 number standard](https://www.twilio.com/docs/glossary/what-e164).
	PhoneNumber NullableString `json:"phone_number,omitempty"`
	// The billing address for the buyer.
	Address NullableAddress `json:"address,omitempty"`
	// The tax information associated with the billing details.
	TaxId NullableTaxId `json:"tax_id,omitempty"`
}

// NewBillingDetails instantiates a new BillingDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingDetails() *BillingDetails {
	this := BillingDetails{}
	return &this
}

// NewBillingDetailsWithDefaults instantiates a new BillingDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingDetailsWithDefaults() *BillingDetails {
	this := BillingDetails{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingDetails) GetFirstName() string {
	if o == nil || o.FirstName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingDetails) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *BillingDetails) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *BillingDetails) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *BillingDetails) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *BillingDetails) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingDetails) GetLastName() string {
	if o == nil || o.LastName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingDetails) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *BillingDetails) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *BillingDetails) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *BillingDetails) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *BillingDetails) UnsetLastName() {
	o.LastName.Unset()
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingDetails) GetEmailAddress() string {
	if o == nil || o.EmailAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress.Get()
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingDetails) GetEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmailAddress.Get(), o.EmailAddress.IsSet()
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *BillingDetails) HasEmailAddress() bool {
	if o != nil && o.EmailAddress.IsSet() {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given NullableString and assigns it to the EmailAddress field.
func (o *BillingDetails) SetEmailAddress(v string) {
	o.EmailAddress.Set(&v)
}
// SetEmailAddressNil sets the value for EmailAddress to be an explicit nil
func (o *BillingDetails) SetEmailAddressNil() {
	o.EmailAddress.Set(nil)
}

// UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
func (o *BillingDetails) UnsetEmailAddress() {
	o.EmailAddress.Unset()
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingDetails) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingDetails) GetPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *BillingDetails) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *BillingDetails) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *BillingDetails) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *BillingDetails) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingDetails) GetAddress() Address {
	if o == nil || o.Address.Get() == nil {
		var ret Address
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingDetails) GetAddressOk() (*Address, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *BillingDetails) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableAddress and assigns it to the Address field.
func (o *BillingDetails) SetAddress(v Address) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *BillingDetails) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *BillingDetails) UnsetAddress() {
	o.Address.Unset()
}

// GetTaxId returns the TaxId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingDetails) GetTaxId() TaxId {
	if o == nil || o.TaxId.Get() == nil {
		var ret TaxId
		return ret
	}
	return *o.TaxId.Get()
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingDetails) GetTaxIdOk() (*TaxId, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TaxId.Get(), o.TaxId.IsSet()
}

// HasTaxId returns a boolean if a field has been set.
func (o *BillingDetails) HasTaxId() bool {
	if o != nil && o.TaxId.IsSet() {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given NullableTaxId and assigns it to the TaxId field.
func (o *BillingDetails) SetTaxId(v TaxId) {
	o.TaxId.Set(&v)
}
// SetTaxIdNil sets the value for TaxId to be an explicit nil
func (o *BillingDetails) SetTaxIdNil() {
	o.TaxId.Set(nil)
}

// UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil
func (o *BillingDetails) UnsetTaxId() {
	o.TaxId.Unset()
}

func (o BillingDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstName.IsSet() {
		toSerialize["first_name"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["last_name"] = o.LastName.Get()
	}
	if o.EmailAddress.IsSet() {
		toSerialize["email_address"] = o.EmailAddress.Get()
	}
	if o.PhoneNumber.IsSet() {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.TaxId.IsSet() {
		toSerialize["tax_id"] = o.TaxId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableBillingDetails struct {
	value *BillingDetails
	isSet bool
}

func (v NullableBillingDetails) Get() *BillingDetails {
	return v.value
}

func (v *NullableBillingDetails) Set(val *BillingDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingDetails(val *BillingDetails) *NullableBillingDetails {
	return &NullableBillingDetails{value: val, isSet: true}
}

func (v NullableBillingDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


