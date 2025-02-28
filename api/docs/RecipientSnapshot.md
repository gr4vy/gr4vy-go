# RecipientSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **NullableString** | The first name(s) or given name for the recipient. | [optional] 
**LastName** | Pointer to **NullableString** | The last name, or family name, of the recipient. | [optional] 
**Address** | Pointer to [**NullableAddress**](Address.md) | The billing address for the recipient. | [optional] 
**AccountNumber** | Pointer to **NullableString** | The destination account number to receive an account funding transaction. | [optional] 
**DateOfBirth** | Pointer to **NullableString** | The recipient&#39;s date of birth. | [optional] 

## Methods

### NewRecipientSnapshot

`func NewRecipientSnapshot() *RecipientSnapshot`

NewRecipientSnapshot instantiates a new RecipientSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipientSnapshotWithDefaults

`func NewRecipientSnapshotWithDefaults() *RecipientSnapshot`

NewRecipientSnapshotWithDefaults instantiates a new RecipientSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *RecipientSnapshot) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *RecipientSnapshot) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *RecipientSnapshot) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *RecipientSnapshot) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *RecipientSnapshot) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *RecipientSnapshot) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *RecipientSnapshot) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *RecipientSnapshot) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *RecipientSnapshot) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *RecipientSnapshot) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *RecipientSnapshot) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *RecipientSnapshot) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetAddress

`func (o *RecipientSnapshot) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RecipientSnapshot) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RecipientSnapshot) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *RecipientSnapshot) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *RecipientSnapshot) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *RecipientSnapshot) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAccountNumber

`func (o *RecipientSnapshot) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *RecipientSnapshot) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *RecipientSnapshot) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *RecipientSnapshot) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### SetAccountNumberNil

`func (o *RecipientSnapshot) SetAccountNumberNil(b bool)`

 SetAccountNumberNil sets the value for AccountNumber to be an explicit nil

### UnsetAccountNumber
`func (o *RecipientSnapshot) UnsetAccountNumber()`

UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
### GetDateOfBirth

`func (o *RecipientSnapshot) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *RecipientSnapshot) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *RecipientSnapshot) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *RecipientSnapshot) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### SetDateOfBirthNil

`func (o *RecipientSnapshot) SetDateOfBirthNil(b bool)`

 SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil

### UnsetDateOfBirth
`func (o *RecipientSnapshot) UnsetDateOfBirth()`

UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


