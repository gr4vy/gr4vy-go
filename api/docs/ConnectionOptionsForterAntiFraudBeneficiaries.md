# ConnectionOptionsForterAntiFraudBeneficiaries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersonalDetails** | [**ConnectionOptionsForterAntiFraudPersonalDetails**](ConnectionOptionsForterAntiFraudPersonalDetails.md) |  | 
**Address** | Pointer to [**NullableConnectionOptionsForterAntiFraudAddress**](ConnectionOptionsForterAntiFraudAddress.md) |  | [optional] 
**Phone** | Pointer to [**[]ConnectionOptionsForterAntiFraudPhone**](ConnectionOptionsForterAntiFraudPhone.md) | List of all phone numbers for the beneficiary. | [optional] 
**Comments** | Pointer to [**NullableConnectionOptionsForterAntiFraudComments**](ConnectionOptionsForterAntiFraudComments.md) |  | [optional] 

## Methods

### NewConnectionOptionsForterAntiFraudBeneficiaries

`func NewConnectionOptionsForterAntiFraudBeneficiaries(personalDetails ConnectionOptionsForterAntiFraudPersonalDetails, ) *ConnectionOptionsForterAntiFraudBeneficiaries`

NewConnectionOptionsForterAntiFraudBeneficiaries instantiates a new ConnectionOptionsForterAntiFraudBeneficiaries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionOptionsForterAntiFraudBeneficiariesWithDefaults

`func NewConnectionOptionsForterAntiFraudBeneficiariesWithDefaults() *ConnectionOptionsForterAntiFraudBeneficiaries`

NewConnectionOptionsForterAntiFraudBeneficiariesWithDefaults instantiates a new ConnectionOptionsForterAntiFraudBeneficiaries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPersonalDetails

`func (o *ConnectionOptionsForterAntiFraudBeneficiaries) GetPersonalDetails() ConnectionOptionsForterAntiFraudPersonalDetails`

GetPersonalDetails returns the PersonalDetails field if non-nil, zero value otherwise.

### GetPersonalDetailsOk

`func (o *ConnectionOptionsForterAntiFraudBeneficiaries) GetPersonalDetailsOk() (*ConnectionOptionsForterAntiFraudPersonalDetails, bool)`

GetPersonalDetailsOk returns a tuple with the PersonalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalDetails

`func (o *ConnectionOptionsForterAntiFraudBeneficiaries) SetPersonalDetails(v ConnectionOptionsForterAntiFraudPersonalDetails)`

SetPersonalDetails sets PersonalDetails field to given value.


### GetAddress

`func (o *ConnectionOptionsForterAntiFraudBeneficiaries) GetAddress() ConnectionOptionsForterAntiFraudAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ConnectionOptionsForterAntiFraudBeneficiaries) GetAddressOk() (*ConnectionOptionsForterAntiFraudAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ConnectionOptionsForterAntiFraudBeneficiaries) SetAddress(v ConnectionOptionsForterAntiFraudAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ConnectionOptionsForterAntiFraudBeneficiaries) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *ConnectionOptionsForterAntiFraudBeneficiaries) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *ConnectionOptionsForterAntiFraudBeneficiaries) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetPhone

`func (o *ConnectionOptionsForterAntiFraudBeneficiaries) GetPhone() []ConnectionOptionsForterAntiFraudPhone`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ConnectionOptionsForterAntiFraudBeneficiaries) GetPhoneOk() (*[]ConnectionOptionsForterAntiFraudPhone, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ConnectionOptionsForterAntiFraudBeneficiaries) SetPhone(v []ConnectionOptionsForterAntiFraudPhone)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ConnectionOptionsForterAntiFraudBeneficiaries) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetComments

`func (o *ConnectionOptionsForterAntiFraudBeneficiaries) GetComments() ConnectionOptionsForterAntiFraudComments`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ConnectionOptionsForterAntiFraudBeneficiaries) GetCommentsOk() (*ConnectionOptionsForterAntiFraudComments, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ConnectionOptionsForterAntiFraudBeneficiaries) SetComments(v ConnectionOptionsForterAntiFraudComments)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ConnectionOptionsForterAntiFraudBeneficiaries) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetCommentsNil

`func (o *ConnectionOptionsForterAntiFraudBeneficiaries) SetCommentsNil(b bool)`

 SetCommentsNil sets the value for Comments to be an explicit nil

### UnsetComments
`func (o *ConnectionOptionsForterAntiFraudBeneficiaries) UnsetComments()`

UnsetComments ensures that no value is present for Comments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


