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

// AirlinePassenger Information of the travelling passenger.
type AirlinePassenger struct {
	// Title of the passenger.
	Title NullableString `json:"title,omitempty"`
	// The first name(s) or given name of the passenger.
	FirstName NullableString `json:"first_name,omitempty"`
	// The last name, or family name, of the passenger.
	LastName NullableString `json:"last_name,omitempty"`
	// The email address of the passenger.
	EmailAddress NullableString `json:"email_address,omitempty"`
	// The phone number of the passenger. This number is formatted according to the [E164 number standard](https://www.twilio.com/docs/glossary/what-e164).
	PhoneNumber NullableString `json:"phone_number,omitempty"`
	// The passenger's unique passport number.
	PassportNumber NullableString `json:"passport_number,omitempty"`
	// The ticket number for a flight.
	TicketNumber NullableString `json:"ticket_number,omitempty"`
	// The passenger's frequent flyer number.
	FrequentFlyerNumber NullableString `json:"frequent_flyer_number,omitempty"`
	// The passenger's date of birth.
	DateOfBirth NullableString `json:"date_of_birth,omitempty"`
	AgeGroup NullableString `json:"age_group,omitempty"`
}

// NewAirlinePassenger instantiates a new AirlinePassenger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAirlinePassenger() *AirlinePassenger {
	this := AirlinePassenger{}
	return &this
}

// NewAirlinePassengerWithDefaults instantiates a new AirlinePassenger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAirlinePassengerWithDefaults() *AirlinePassenger {
	this := AirlinePassenger{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AirlinePassenger) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AirlinePassenger) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *AirlinePassenger) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *AirlinePassenger) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *AirlinePassenger) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *AirlinePassenger) UnsetTitle() {
	o.Title.Unset()
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AirlinePassenger) GetFirstName() string {
	if o == nil || o.FirstName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AirlinePassenger) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *AirlinePassenger) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *AirlinePassenger) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *AirlinePassenger) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *AirlinePassenger) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AirlinePassenger) GetLastName() string {
	if o == nil || o.LastName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AirlinePassenger) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *AirlinePassenger) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *AirlinePassenger) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *AirlinePassenger) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *AirlinePassenger) UnsetLastName() {
	o.LastName.Unset()
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AirlinePassenger) GetEmailAddress() string {
	if o == nil || o.EmailAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress.Get()
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AirlinePassenger) GetEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmailAddress.Get(), o.EmailAddress.IsSet()
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *AirlinePassenger) HasEmailAddress() bool {
	if o != nil && o.EmailAddress.IsSet() {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given NullableString and assigns it to the EmailAddress field.
func (o *AirlinePassenger) SetEmailAddress(v string) {
	o.EmailAddress.Set(&v)
}
// SetEmailAddressNil sets the value for EmailAddress to be an explicit nil
func (o *AirlinePassenger) SetEmailAddressNil() {
	o.EmailAddress.Set(nil)
}

// UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
func (o *AirlinePassenger) UnsetEmailAddress() {
	o.EmailAddress.Unset()
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AirlinePassenger) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AirlinePassenger) GetPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *AirlinePassenger) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *AirlinePassenger) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *AirlinePassenger) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *AirlinePassenger) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetPassportNumber returns the PassportNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AirlinePassenger) GetPassportNumber() string {
	if o == nil || o.PassportNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.PassportNumber.Get()
}

// GetPassportNumberOk returns a tuple with the PassportNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AirlinePassenger) GetPassportNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PassportNumber.Get(), o.PassportNumber.IsSet()
}

// HasPassportNumber returns a boolean if a field has been set.
func (o *AirlinePassenger) HasPassportNumber() bool {
	if o != nil && o.PassportNumber.IsSet() {
		return true
	}

	return false
}

// SetPassportNumber gets a reference to the given NullableString and assigns it to the PassportNumber field.
func (o *AirlinePassenger) SetPassportNumber(v string) {
	o.PassportNumber.Set(&v)
}
// SetPassportNumberNil sets the value for PassportNumber to be an explicit nil
func (o *AirlinePassenger) SetPassportNumberNil() {
	o.PassportNumber.Set(nil)
}

// UnsetPassportNumber ensures that no value is present for PassportNumber, not even an explicit nil
func (o *AirlinePassenger) UnsetPassportNumber() {
	o.PassportNumber.Unset()
}

// GetTicketNumber returns the TicketNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AirlinePassenger) GetTicketNumber() string {
	if o == nil || o.TicketNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.TicketNumber.Get()
}

// GetTicketNumberOk returns a tuple with the TicketNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AirlinePassenger) GetTicketNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TicketNumber.Get(), o.TicketNumber.IsSet()
}

// HasTicketNumber returns a boolean if a field has been set.
func (o *AirlinePassenger) HasTicketNumber() bool {
	if o != nil && o.TicketNumber.IsSet() {
		return true
	}

	return false
}

// SetTicketNumber gets a reference to the given NullableString and assigns it to the TicketNumber field.
func (o *AirlinePassenger) SetTicketNumber(v string) {
	o.TicketNumber.Set(&v)
}
// SetTicketNumberNil sets the value for TicketNumber to be an explicit nil
func (o *AirlinePassenger) SetTicketNumberNil() {
	o.TicketNumber.Set(nil)
}

// UnsetTicketNumber ensures that no value is present for TicketNumber, not even an explicit nil
func (o *AirlinePassenger) UnsetTicketNumber() {
	o.TicketNumber.Unset()
}

// GetFrequentFlyerNumber returns the FrequentFlyerNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AirlinePassenger) GetFrequentFlyerNumber() string {
	if o == nil || o.FrequentFlyerNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.FrequentFlyerNumber.Get()
}

// GetFrequentFlyerNumberOk returns a tuple with the FrequentFlyerNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AirlinePassenger) GetFrequentFlyerNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FrequentFlyerNumber.Get(), o.FrequentFlyerNumber.IsSet()
}

// HasFrequentFlyerNumber returns a boolean if a field has been set.
func (o *AirlinePassenger) HasFrequentFlyerNumber() bool {
	if o != nil && o.FrequentFlyerNumber.IsSet() {
		return true
	}

	return false
}

// SetFrequentFlyerNumber gets a reference to the given NullableString and assigns it to the FrequentFlyerNumber field.
func (o *AirlinePassenger) SetFrequentFlyerNumber(v string) {
	o.FrequentFlyerNumber.Set(&v)
}
// SetFrequentFlyerNumberNil sets the value for FrequentFlyerNumber to be an explicit nil
func (o *AirlinePassenger) SetFrequentFlyerNumberNil() {
	o.FrequentFlyerNumber.Set(nil)
}

// UnsetFrequentFlyerNumber ensures that no value is present for FrequentFlyerNumber, not even an explicit nil
func (o *AirlinePassenger) UnsetFrequentFlyerNumber() {
	o.FrequentFlyerNumber.Unset()
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AirlinePassenger) GetDateOfBirth() string {
	if o == nil || o.DateOfBirth.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateOfBirth.Get()
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AirlinePassenger) GetDateOfBirthOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DateOfBirth.Get(), o.DateOfBirth.IsSet()
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *AirlinePassenger) HasDateOfBirth() bool {
	if o != nil && o.DateOfBirth.IsSet() {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given NullableString and assigns it to the DateOfBirth field.
func (o *AirlinePassenger) SetDateOfBirth(v string) {
	o.DateOfBirth.Set(&v)
}
// SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil
func (o *AirlinePassenger) SetDateOfBirthNil() {
	o.DateOfBirth.Set(nil)
}

// UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
func (o *AirlinePassenger) UnsetDateOfBirth() {
	o.DateOfBirth.Unset()
}

// GetAgeGroup returns the AgeGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AirlinePassenger) GetAgeGroup() string {
	if o == nil || o.AgeGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.AgeGroup.Get()
}

// GetAgeGroupOk returns a tuple with the AgeGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AirlinePassenger) GetAgeGroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AgeGroup.Get(), o.AgeGroup.IsSet()
}

// HasAgeGroup returns a boolean if a field has been set.
func (o *AirlinePassenger) HasAgeGroup() bool {
	if o != nil && o.AgeGroup.IsSet() {
		return true
	}

	return false
}

// SetAgeGroup gets a reference to the given NullableString and assigns it to the AgeGroup field.
func (o *AirlinePassenger) SetAgeGroup(v string) {
	o.AgeGroup.Set(&v)
}
// SetAgeGroupNil sets the value for AgeGroup to be an explicit nil
func (o *AirlinePassenger) SetAgeGroupNil() {
	o.AgeGroup.Set(nil)
}

// UnsetAgeGroup ensures that no value is present for AgeGroup, not even an explicit nil
func (o *AirlinePassenger) UnsetAgeGroup() {
	o.AgeGroup.Unset()
}

func (o AirlinePassenger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
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
	if o.PassportNumber.IsSet() {
		toSerialize["passport_number"] = o.PassportNumber.Get()
	}
	if o.TicketNumber.IsSet() {
		toSerialize["ticket_number"] = o.TicketNumber.Get()
	}
	if o.FrequentFlyerNumber.IsSet() {
		toSerialize["frequent_flyer_number"] = o.FrequentFlyerNumber.Get()
	}
	if o.DateOfBirth.IsSet() {
		toSerialize["date_of_birth"] = o.DateOfBirth.Get()
	}
	if o.AgeGroup.IsSet() {
		toSerialize["age_group"] = o.AgeGroup.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAirlinePassenger struct {
	value *AirlinePassenger
	isSet bool
}

func (v NullableAirlinePassenger) Get() *AirlinePassenger {
	return v.value
}

func (v *NullableAirlinePassenger) Set(val *AirlinePassenger) {
	v.value = val
	v.isSet = true
}

func (v NullableAirlinePassenger) IsSet() bool {
	return v.isSet
}

func (v *NullableAirlinePassenger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAirlinePassenger(val *AirlinePassenger) *NullableAirlinePassenger {
	return &NullableAirlinePassenger{value: val, isSet: true}
}

func (v NullableAirlinePassenger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAirlinePassenger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


