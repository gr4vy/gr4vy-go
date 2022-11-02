# ShippingAddressUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **NullableString** | The first name(s) or given name for the buyer. | [optional] 
**LastName** | Pointer to **NullableString** | The last name, or family name, of the buyer. | [optional] 
**EmailAddress** | Pointer to **NullableString** | The email address for the buyer. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The phone number for the buyer which should be formatted according to the [E164 number standard](https://www.twilio.com/docs/glossary/what-e164). | [optional] 
**Address** | Pointer to [**NullableAddress**](Address.md) | The physical shipping address associated to this buyer. | [optional] 

## Methods

### NewShippingAddressUpdateRequest

`func NewShippingAddressUpdateRequest() *ShippingAddressUpdateRequest`

NewShippingAddressUpdateRequest instantiates a new ShippingAddressUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingAddressUpdateRequestWithDefaults

`func NewShippingAddressUpdateRequestWithDefaults() *ShippingAddressUpdateRequest`

NewShippingAddressUpdateRequestWithDefaults instantiates a new ShippingAddressUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *ShippingAddressUpdateRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ShippingAddressUpdateRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ShippingAddressUpdateRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ShippingAddressUpdateRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ShippingAddressUpdateRequest) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ShippingAddressUpdateRequest) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *ShippingAddressUpdateRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ShippingAddressUpdateRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ShippingAddressUpdateRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ShippingAddressUpdateRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ShippingAddressUpdateRequest) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ShippingAddressUpdateRequest) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetEmailAddress

`func (o *ShippingAddressUpdateRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *ShippingAddressUpdateRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *ShippingAddressUpdateRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *ShippingAddressUpdateRequest) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *ShippingAddressUpdateRequest) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *ShippingAddressUpdateRequest) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetPhoneNumber

`func (o *ShippingAddressUpdateRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ShippingAddressUpdateRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ShippingAddressUpdateRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ShippingAddressUpdateRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *ShippingAddressUpdateRequest) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *ShippingAddressUpdateRequest) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetAddress

`func (o *ShippingAddressUpdateRequest) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ShippingAddressUpdateRequest) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ShippingAddressUpdateRequest) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ShippingAddressUpdateRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *ShippingAddressUpdateRequest) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *ShippingAddressUpdateRequest) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


