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

// Address Billing details associated to a buyer.
type Address struct {
	// The city for the billing address.
	City string `json:"city"`
	// The country for the billing address.
	Country string `json:"country"`
	// The postal code or zip code for the billing address.
	PostalCode string `json:"postal_code"`
	// The state, county, or province for the billing address.
	State string `json:"state"`
	// The code of state, county, or province for the billing address in ISO 3166-2 format.
	StateCode NullableString `json:"state_code,omitempty"`
	// The house number or name for the billing address. Not all payment services use this field but some do.
	HouseNumberOrName NullableString `json:"house_number_or_name,omitempty"`
	// The first line of the billing address.
	Line1 string `json:"line1"`
	// The second line of the billing address.
	Line2 NullableString `json:"line2,omitempty"`
	// The optional name of the company or organisation to add to the billing address.
	Organization NullableString `json:"organization,omitempty"`
}

// NewAddress instantiates a new Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress(city string, country string, postalCode string, state string, line1 string) *Address {
	this := Address{}
	this.City = city
	this.Country = country
	this.PostalCode = postalCode
	this.State = state
	this.Line1 = line1
	return &this
}

// NewAddressWithDefaults instantiates a new Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressWithDefaults() *Address {
	this := Address{}
	return &this
}

// GetCity returns the City field value
func (o *Address) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *Address) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *Address) SetCity(v string) {
	o.City = v
}

// GetCountry returns the Country field value
func (o *Address) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *Address) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *Address) SetCountry(v string) {
	o.Country = v
}

// GetPostalCode returns the PostalCode field value
func (o *Address) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *Address) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *Address) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetState returns the State field value
func (o *Address) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Address) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Address) SetState(v string) {
	o.State = v
}

// GetStateCode returns the StateCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetStateCode() string {
	if o == nil || o.StateCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.StateCode.Get()
}

// GetStateCodeOk returns a tuple with the StateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetStateCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StateCode.Get(), o.StateCode.IsSet()
}

// HasStateCode returns a boolean if a field has been set.
func (o *Address) HasStateCode() bool {
	if o != nil && o.StateCode.IsSet() {
		return true
	}

	return false
}

// SetStateCode gets a reference to the given NullableString and assigns it to the StateCode field.
func (o *Address) SetStateCode(v string) {
	o.StateCode.Set(&v)
}
// SetStateCodeNil sets the value for StateCode to be an explicit nil
func (o *Address) SetStateCodeNil() {
	o.StateCode.Set(nil)
}

// UnsetStateCode ensures that no value is present for StateCode, not even an explicit nil
func (o *Address) UnsetStateCode() {
	o.StateCode.Unset()
}

// GetHouseNumberOrName returns the HouseNumberOrName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetHouseNumberOrName() string {
	if o == nil || o.HouseNumberOrName.Get() == nil {
		var ret string
		return ret
	}
	return *o.HouseNumberOrName.Get()
}

// GetHouseNumberOrNameOk returns a tuple with the HouseNumberOrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetHouseNumberOrNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HouseNumberOrName.Get(), o.HouseNumberOrName.IsSet()
}

// HasHouseNumberOrName returns a boolean if a field has been set.
func (o *Address) HasHouseNumberOrName() bool {
	if o != nil && o.HouseNumberOrName.IsSet() {
		return true
	}

	return false
}

// SetHouseNumberOrName gets a reference to the given NullableString and assigns it to the HouseNumberOrName field.
func (o *Address) SetHouseNumberOrName(v string) {
	o.HouseNumberOrName.Set(&v)
}
// SetHouseNumberOrNameNil sets the value for HouseNumberOrName to be an explicit nil
func (o *Address) SetHouseNumberOrNameNil() {
	o.HouseNumberOrName.Set(nil)
}

// UnsetHouseNumberOrName ensures that no value is present for HouseNumberOrName, not even an explicit nil
func (o *Address) UnsetHouseNumberOrName() {
	o.HouseNumberOrName.Unset()
}

// GetLine1 returns the Line1 field value
func (o *Address) GetLine1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Line1
}

// GetLine1Ok returns a tuple with the Line1 field value
// and a boolean to check if the value has been set.
func (o *Address) GetLine1Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Line1, true
}

// SetLine1 sets field value
func (o *Address) SetLine1(v string) {
	o.Line1 = v
}

// GetLine2 returns the Line2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetLine2() string {
	if o == nil || o.Line2.Get() == nil {
		var ret string
		return ret
	}
	return *o.Line2.Get()
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetLine2Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Line2.Get(), o.Line2.IsSet()
}

// HasLine2 returns a boolean if a field has been set.
func (o *Address) HasLine2() bool {
	if o != nil && o.Line2.IsSet() {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given NullableString and assigns it to the Line2 field.
func (o *Address) SetLine2(v string) {
	o.Line2.Set(&v)
}
// SetLine2Nil sets the value for Line2 to be an explicit nil
func (o *Address) SetLine2Nil() {
	o.Line2.Set(nil)
}

// UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
func (o *Address) UnsetLine2() {
	o.Line2.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetOrganization() string {
	if o == nil || o.Organization.Get() == nil {
		var ret string
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetOrganizationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *Address) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableString and assigns it to the Organization field.
func (o *Address) SetOrganization(v string) {
	o.Organization.Set(&v)
}
// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *Address) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *Address) UnsetOrganization() {
	o.Organization.Unset()
}

func (o Address) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["city"] = o.City
	}
	if true {
		toSerialize["country"] = o.Country
	}
	if true {
		toSerialize["postal_code"] = o.PostalCode
	}
	if true {
		toSerialize["state"] = o.State
	}
	if o.StateCode.IsSet() {
		toSerialize["state_code"] = o.StateCode.Get()
	}
	if o.HouseNumberOrName.IsSet() {
		toSerialize["house_number_or_name"] = o.HouseNumberOrName.Get()
	}
	if true {
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

type NullableAddress struct {
	value *Address
	isSet bool
}

func (v NullableAddress) Get() *Address {
	return v.value
}

func (v *NullableAddress) Set(val *Address) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress(val *Address) *NullableAddress {
	return &NullableAddress{value: val, isSet: true}
}

func (v NullableAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


