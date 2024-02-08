# ConnectionOptionsStripeCardStripeConnect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StripeAccount** | Pointer to **NullableString** | The ID of the connected Stripe account to process for. | [optional] 
**ApplicationFeeAmount** | Pointer to **NullableFloat32** | The application fee to charge when processing for a connected account. | [optional] 

## Methods

### NewConnectionOptionsStripeCardStripeConnect

`func NewConnectionOptionsStripeCardStripeConnect() *ConnectionOptionsStripeCardStripeConnect`

NewConnectionOptionsStripeCardStripeConnect instantiates a new ConnectionOptionsStripeCardStripeConnect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionOptionsStripeCardStripeConnectWithDefaults

`func NewConnectionOptionsStripeCardStripeConnectWithDefaults() *ConnectionOptionsStripeCardStripeConnect`

NewConnectionOptionsStripeCardStripeConnectWithDefaults instantiates a new ConnectionOptionsStripeCardStripeConnect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStripeAccount

`func (o *ConnectionOptionsStripeCardStripeConnect) GetStripeAccount() string`

GetStripeAccount returns the StripeAccount field if non-nil, zero value otherwise.

### GetStripeAccountOk

`func (o *ConnectionOptionsStripeCardStripeConnect) GetStripeAccountOk() (*string, bool)`

GetStripeAccountOk returns a tuple with the StripeAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeAccount

`func (o *ConnectionOptionsStripeCardStripeConnect) SetStripeAccount(v string)`

SetStripeAccount sets StripeAccount field to given value.

### HasStripeAccount

`func (o *ConnectionOptionsStripeCardStripeConnect) HasStripeAccount() bool`

HasStripeAccount returns a boolean if a field has been set.

### SetStripeAccountNil

`func (o *ConnectionOptionsStripeCardStripeConnect) SetStripeAccountNil(b bool)`

 SetStripeAccountNil sets the value for StripeAccount to be an explicit nil

### UnsetStripeAccount
`func (o *ConnectionOptionsStripeCardStripeConnect) UnsetStripeAccount()`

UnsetStripeAccount ensures that no value is present for StripeAccount, not even an explicit nil
### GetApplicationFeeAmount

`func (o *ConnectionOptionsStripeCardStripeConnect) GetApplicationFeeAmount() float32`

GetApplicationFeeAmount returns the ApplicationFeeAmount field if non-nil, zero value otherwise.

### GetApplicationFeeAmountOk

`func (o *ConnectionOptionsStripeCardStripeConnect) GetApplicationFeeAmountOk() (*float32, bool)`

GetApplicationFeeAmountOk returns a tuple with the ApplicationFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFeeAmount

`func (o *ConnectionOptionsStripeCardStripeConnect) SetApplicationFeeAmount(v float32)`

SetApplicationFeeAmount sets ApplicationFeeAmount field to given value.

### HasApplicationFeeAmount

`func (o *ConnectionOptionsStripeCardStripeConnect) HasApplicationFeeAmount() bool`

HasApplicationFeeAmount returns a boolean if a field has been set.

### SetApplicationFeeAmountNil

`func (o *ConnectionOptionsStripeCardStripeConnect) SetApplicationFeeAmountNil(b bool)`

 SetApplicationFeeAmountNil sets the value for ApplicationFeeAmount to be an explicit nil

### UnsetApplicationFeeAmount
`func (o *ConnectionOptionsStripeCardStripeConnect) UnsetApplicationFeeAmount()`

UnsetApplicationFeeAmount ensures that no value is present for ApplicationFeeAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


