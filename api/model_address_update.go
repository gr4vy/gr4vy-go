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

// AddressUpdate struct for AddressUpdate
type AddressUpdate struct {
	// The city for the billing address.
	City *string `json:"city,omitempty"`
	// The country for the billing address.
	Country *string `json:"country,omitempty"`
	// The postal code or zip code for the billing address.
	PostalCode *string `json:"postal_code,omitempty"`
	// The state, county, or province for the billing address.
	State *string `json:"state,omitempty"`
	// The code of state, county, or province for the billing address in ISO 3166-2 format.
	StateCode NullableString `json:"state_code,omitempty"`
	// The house number or name for the billing address. Not all payment services use this field but some do.
	HouseNumberOrName NullableString `json:"house_number_or_name,omitempty"`
	// The first line of the billing address.
	Line1 *string `json:"line1,omitempty"`
	// The second line of the billing address.
	Line2 NullableString `json:"line2,omitempty"`
	// The optional name of the company or organisation to add to the billing address.
	Organization NullableString `json:"organization,omitempty"`
}

// NewAddressUpdate instantiates a new AddressUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressUpdate() *AddressUpdate {
	this := AddressUpdate{}
	return &this
}

// NewAddressUpdateWithDefaults instantiates a new AddressUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressUpdateWithDefaults() *AddressUpdate {
	this := AddressUpdate{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *AddressUpdate) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressUpdate) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *AddressUpdate) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *AddressUpdate) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *AddressUpdate) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressUpdate) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *AddressUpdate) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *AddressUpdate) SetCountry(v string) {
	o.Country = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *AddressUpdate) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressUpdate) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *AddressUpdate) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *AddressUpdate) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AddressUpdate) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressUpdate) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AddressUpdate) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AddressUpdate) SetState(v string) {
	o.State = &v
}

// GetStateCode returns the StateCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressUpdate) GetStateCode() string {
	if o == nil || o.StateCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.StateCode.Get()
}

// GetStateCodeOk returns a tuple with the StateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressUpdate) GetStateCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StateCode.Get(), o.StateCode.IsSet()
}

// HasStateCode returns a boolean if a field has been set.
func (o *AddressUpdate) HasStateCode() bool {
	if o != nil && o.StateCode.IsSet() {
		return true
	}

	return false
}

// SetStateCode gets a reference to the given NullableString and assigns it to the StateCode field.
func (o *AddressUpdate) SetStateCode(v string) {
	o.StateCode.Set(&v)
}
// SetStateCodeNil sets the value for StateCode to be an explicit nil
func (o *AddressUpdate) SetStateCodeNil() {
	o.StateCode.Set(nil)
}

// UnsetStateCode ensures that no value is present for StateCode, not even an explicit nil
func (o *AddressUpdate) UnsetStateCode() {
	o.StateCode.Unset()
}

// GetHouseNumberOrName returns the HouseNumberOrName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressUpdate) GetHouseNumberOrName() string {
	if o == nil || o.HouseNumberOrName.Get() == nil {
		var ret string
		return ret
	}
	return *o.HouseNumberOrName.Get()
}

// GetHouseNumberOrNameOk returns a tuple with the HouseNumberOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressUpdate) GetHouseNumberOrNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HouseNumberOrName.Get(), o.HouseNumberOrName.IsSet()
}

// HasHouseNumberOrName returns a boolean if a field has been set.
func (o *AddressUpdate) HasHouseNumberOrName() bool {
	if o != nil && o.HouseNumberOrName.IsSet() {
		return true
	}

	return false
}

// SetHouseNumberOrName gets a reference to the given NullableString and assigns it to the HouseNumberOrName field.
func (o *AddressUpdate) SetHouseNumberOrName(v string) {
	o.HouseNumberOrName.Set(&v)
}
// SetHouseNumberOrNameNil sets the value for HouseNumberOrName to be an explicit nil
func (o *AddressUpdate) SetHouseNumberOrNameNil() {
	o.HouseNumberOrName.Set(nil)
}

// UnsetHouseNumberOrName ensures that no value is present for HouseNumberOrName, not even an explicit nil
func (o *AddressUpdate) UnsetHouseNumberOrName() {
	o.HouseNumberOrName.Unset()
}

// GetLine1 returns the Line1 field value if set, zero value otherwise.
func (o *AddressUpdate) GetLine1() string {
	if o == nil || o.Line1 == nil {
		var ret string
		return ret
	}
	return *o.Line1
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressUpdate) GetLine1Ok() (*string, bool) {
	if o == nil || o.Line1 == nil {
		return nil, false
	}
	return o.Line1, true
}

// HasLine1 returns a boolean if a field has been set.
func (o *AddressUpdate) HasLine1() bool {
	if o != nil && o.Line1 != nil {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given string and assigns it to the Line1 field.
func (o *AddressUpdate) SetLine1(v string) {
	o.Line1 = &v
}

// GetLine2 returns the Line2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressUpdate) GetLine2() string {
	if o == nil || o.Line2.Get() == nil {
		var ret string
		return ret
	}
	return *o.Line2.Get()
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressUpdate) GetLine2Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Line2.Get(), o.Line2.IsSet()
}

// HasLine2 returns a boolean if a field has been set.
func (o *AddressUpdate) HasLine2() bool {
	if o != nil && o.Line2.IsSet() {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given NullableString and assigns it to the Line2 field.
func (o *AddressUpdate) SetLine2(v string) {
	o.Line2.Set(&v)
}
// SetLine2Nil sets the value for Line2 to be an explicit nil
func (o *AddressUpdate) SetLine2Nil() {
	o.Line2.Set(nil)
}

// UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
func (o *AddressUpdate) UnsetLine2() {
	o.Line2.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressUpdate) GetOrganization() string {
	if o == nil || o.Organization.Get() == nil {
		var ret string
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressUpdate) GetOrganizationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *AddressUpdate) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableString and assigns it to the Organization field.
func (o *AddressUpdate) SetOrganization(v string) {
	o.Organization.Set(&v)
}
// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *AddressUpdate) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *AddressUpdate) UnsetOrganization() {
	o.Organization.Unset()
}

func (o AddressUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.PostalCode != nil {
		toSerialize["postal_code"] = o.PostalCode
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.StateCode.IsSet() {
		toSerialize["state_code"] = o.StateCode.Get()
	}
	if o.HouseNumberOrName.IsSet() {
		toSerialize["house_number_or_name"] = o.HouseNumberOrName.Get()
	}
	if o.Line1 != nil {
		toSerialize["line1"] = o.Line1
	}
	if o.Line2.IsSet() {
		toSerialize["line2"] = o.Line2.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["organization"] = o.Organization.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAddressUpdate struct {
	value *AddressUpdate
	isSet bool
}

func (v NullableAddressUpdate) Get() *AddressUpdate {
	return v.value
}

func (v *NullableAddressUpdate) Set(val *AddressUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressUpdate(val *AddressUpdate) *NullableAddressUpdate {
	return &NullableAddressUpdate{value: val, isSet: true}
}

func (v NullableAddressUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


