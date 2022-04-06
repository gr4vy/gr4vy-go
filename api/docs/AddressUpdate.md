# AddressUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | The city for the billing address. | [optional] 
**Country** | Pointer to **string** | The country for the billing address. | [optional] 
**PostalCode** | Pointer to **string** | The postal code or zip code for the billing address. | [optional] 
**State** | Pointer to **string** | The state, county, or province for the billing address. | [optional] 
**StateCode** | Pointer to **NullableString** | The code of state, county, or province for the billing address in ISO 3166-2 format. | [optional] 
**HouseNumberOrName** | Pointer to **NullableString** | The house number or name for the billing address. Not all payment services use this field but some do. | [optional] 
**Line1** | Pointer to **string** | The first line of the billing address. | [optional] 
**Line2** | Pointer to **NullableString** | The second line of the billing address. | [optional] 
**Organization** | Pointer to **NullableString** | The optional name of the company or organisation to add to the billing address. | [optional] 

## Methods

### NewAddressUpdate

`func NewAddressUpdate() *AddressUpdate`

NewAddressUpdate instantiates a new AddressUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressUpdateWithDefaults

`func NewAddressUpdateWithDefaults() *AddressUpdate`

NewAddressUpdateWithDefaults instantiates a new AddressUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *AddressUpdate) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressUpdate) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressUpdate) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AddressUpdate) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *AddressUpdate) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AddressUpdate) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AddressUpdate) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AddressUpdate) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPostalCode

`func (o *AddressUpdate) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AddressUpdate) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AddressUpdate) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AddressUpdate) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *AddressUpdate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AddressUpdate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AddressUpdate) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AddressUpdate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateCode

`func (o *AddressUpdate) GetStateCode() string`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *AddressUpdate) GetStateCodeOk() (*string, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *AddressUpdate) SetStateCode(v string)`

SetStateCode sets StateCode field to given value.

### HasStateCode

`func (o *AddressUpdate) HasStateCode() bool`

HasStateCode returns a boolean if a field has been set.

### SetStateCodeNil

`func (o *AddressUpdate) SetStateCodeNil(b bool)`

 SetStateCodeNil sets the value for StateCode to be an explicit nil

### UnsetStateCode
`func (o *AddressUpdate) UnsetStateCode()`

UnsetStateCode ensures that no value is present for StateCode, not even an explicit nil
### GetHouseNumberOrName

`func (o *AddressUpdate) GetHouseNumberOrName() string`

GetHouseNumberOrName returns the HouseNumberOrName field if non-nil, zero value otherwise.

### GetHouseNumberOrNameOk

`func (o *AddressUpdate) GetHouseNumberOrNameOk() (*string, bool)`

GetHouseNumberOrNameOk returns a tuple with the HouseNumberOrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumberOrName

`func (o *AddressUpdate) SetHouseNumberOrName(v string)`

SetHouseNumberOrName sets HouseNumberOrName field to given value.

### HasHouseNumberOrName

`func (o *AddressUpdate) HasHouseNumberOrName() bool`

HasHouseNumberOrName returns a boolean if a field has been set.

### SetHouseNumberOrNameNil

`func (o *AddressUpdate) SetHouseNumberOrNameNil(b bool)`

 SetHouseNumberOrNameNil sets the value for HouseNumberOrName to be an explicit nil

### UnsetHouseNumberOrName
`func (o *AddressUpdate) UnsetHouseNumberOrName()`

UnsetHouseNumberOrName ensures that no value is present for HouseNumberOrName, not even an explicit nil
### GetLine1

`func (o *AddressUpdate) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *AddressUpdate) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *AddressUpdate) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *AddressUpdate) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### GetLine2

`func (o *AddressUpdate) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *AddressUpdate) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *AddressUpdate) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *AddressUpdate) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### SetLine2Nil

`func (o *AddressUpdate) SetLine2Nil(b bool)`

 SetLine2Nil sets the value for Line2 to be an explicit nil

### UnsetLine2
`func (o *AddressUpdate) UnsetLine2()`

UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
### GetOrganization

`func (o *AddressUpdate) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AddressUpdate) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AddressUpdate) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *AddressUpdate) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *AddressUpdate) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *AddressUpdate) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


