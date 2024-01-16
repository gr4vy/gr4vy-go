# ConnectionOptionsForterAntiFraudAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | **string** | Country, two-letter ISO 3166-1 alpha 2 country code. | 
**Address1** | Pointer to **NullableString** | Street-level address. Required when full address details are available. | [optional] 
**Address2** | Pointer to **NullableString** | Unit-level address. | [optional] 
**Zip** | Pointer to **NullableString** | Zipcode. | [optional] 
**Region** | Pointer to **NullableString** | Top-level administrative subdivision - state/province/department/etc. Can be either abbreviated format or full name (NY/New York). | [optional] 
**Company** | Pointer to **NullableString** | Company name. | [optional] 
**City** | Pointer to **NullableString** | City. Required when full address details are available. | [optional] 

## Methods

### NewConnectionOptionsForterAntiFraudAddress

`func NewConnectionOptionsForterAntiFraudAddress(country string, ) *ConnectionOptionsForterAntiFraudAddress`

NewConnectionOptionsForterAntiFraudAddress instantiates a new ConnectionOptionsForterAntiFraudAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionOptionsForterAntiFraudAddressWithDefaults

`func NewConnectionOptionsForterAntiFraudAddressWithDefaults() *ConnectionOptionsForterAntiFraudAddress`

NewConnectionOptionsForterAntiFraudAddressWithDefaults instantiates a new ConnectionOptionsForterAntiFraudAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *ConnectionOptionsForterAntiFraudAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ConnectionOptionsForterAntiFraudAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ConnectionOptionsForterAntiFraudAddress) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetAddress1

`func (o *ConnectionOptionsForterAntiFraudAddress) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *ConnectionOptionsForterAntiFraudAddress) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *ConnectionOptionsForterAntiFraudAddress) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *ConnectionOptionsForterAntiFraudAddress) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### SetAddress1Nil

`func (o *ConnectionOptionsForterAntiFraudAddress) SetAddress1Nil(b bool)`

 SetAddress1Nil sets the value for Address1 to be an explicit nil

### UnsetAddress1
`func (o *ConnectionOptionsForterAntiFraudAddress) UnsetAddress1()`

UnsetAddress1 ensures that no value is present for Address1, not even an explicit nil
### GetAddress2

`func (o *ConnectionOptionsForterAntiFraudAddress) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *ConnectionOptionsForterAntiFraudAddress) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *ConnectionOptionsForterAntiFraudAddress) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *ConnectionOptionsForterAntiFraudAddress) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### SetAddress2Nil

`func (o *ConnectionOptionsForterAntiFraudAddress) SetAddress2Nil(b bool)`

 SetAddress2Nil sets the value for Address2 to be an explicit nil

### UnsetAddress2
`func (o *ConnectionOptionsForterAntiFraudAddress) UnsetAddress2()`

UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
### GetZip

`func (o *ConnectionOptionsForterAntiFraudAddress) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *ConnectionOptionsForterAntiFraudAddress) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *ConnectionOptionsForterAntiFraudAddress) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *ConnectionOptionsForterAntiFraudAddress) HasZip() bool`

HasZip returns a boolean if a field has been set.

### SetZipNil

`func (o *ConnectionOptionsForterAntiFraudAddress) SetZipNil(b bool)`

 SetZipNil sets the value for Zip to be an explicit nil

### UnsetZip
`func (o *ConnectionOptionsForterAntiFraudAddress) UnsetZip()`

UnsetZip ensures that no value is present for Zip, not even an explicit nil
### GetRegion

`func (o *ConnectionOptionsForterAntiFraudAddress) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ConnectionOptionsForterAntiFraudAddress) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ConnectionOptionsForterAntiFraudAddress) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ConnectionOptionsForterAntiFraudAddress) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *ConnectionOptionsForterAntiFraudAddress) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *ConnectionOptionsForterAntiFraudAddress) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetCompany

`func (o *ConnectionOptionsForterAntiFraudAddress) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ConnectionOptionsForterAntiFraudAddress) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ConnectionOptionsForterAntiFraudAddress) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ConnectionOptionsForterAntiFraudAddress) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *ConnectionOptionsForterAntiFraudAddress) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *ConnectionOptionsForterAntiFraudAddress) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetCity

`func (o *ConnectionOptionsForterAntiFraudAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ConnectionOptionsForterAntiFraudAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ConnectionOptionsForterAntiFraudAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ConnectionOptionsForterAntiFraudAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *ConnectionOptionsForterAntiFraudAddress) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *ConnectionOptionsForterAntiFraudAddress) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


