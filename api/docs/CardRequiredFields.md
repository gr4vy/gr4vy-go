# CardRequiredFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **bool** | The first (given) name of the buyer. | [optional] [readonly] 
**LastName** | Pointer to **bool** | The last (family) name of the buyer. | [optional] [readonly] 
**EmailAddress** | Pointer to **bool** | The email address of the buyer. | [optional] [readonly] 
**PhoneNumber** | Pointer to **bool** | The phone number of the buyer. | [optional] [readonly] 
**Address** | Pointer to [**CardRequiredFieldsAddress**](CardRequiredFieldsAddress.md) |  | [optional] 
**TaxId** | Pointer to **bool** | The tax id code associated with the billing details. | [optional] [readonly] 

## Methods

### NewCardRequiredFields

`func NewCardRequiredFields() *CardRequiredFields`

NewCardRequiredFields instantiates a new CardRequiredFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardRequiredFieldsWithDefaults

`func NewCardRequiredFieldsWithDefaults() *CardRequiredFields`

NewCardRequiredFieldsWithDefaults instantiates a new CardRequiredFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *CardRequiredFields) GetFirstName() bool`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CardRequiredFields) GetFirstNameOk() (*bool, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CardRequiredFields) SetFirstName(v bool)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CardRequiredFields) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *CardRequiredFields) GetLastName() bool`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CardRequiredFields) GetLastNameOk() (*bool, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CardRequiredFields) SetLastName(v bool)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CardRequiredFields) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *CardRequiredFields) GetEmailAddress() bool`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *CardRequiredFields) GetEmailAddressOk() (*bool, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *CardRequiredFields) SetEmailAddress(v bool)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *CardRequiredFields) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *CardRequiredFields) GetPhoneNumber() bool`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CardRequiredFields) GetPhoneNumberOk() (*bool, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CardRequiredFields) SetPhoneNumber(v bool)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CardRequiredFields) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetAddress

`func (o *CardRequiredFields) GetAddress() CardRequiredFieldsAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CardRequiredFields) GetAddressOk() (*CardRequiredFieldsAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CardRequiredFields) SetAddress(v CardRequiredFieldsAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CardRequiredFields) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetTaxId

`func (o *CardRequiredFields) GetTaxId() bool`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *CardRequiredFields) GetTaxIdOk() (*bool, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *CardRequiredFields) SetTaxId(v bool)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *CardRequiredFields) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


