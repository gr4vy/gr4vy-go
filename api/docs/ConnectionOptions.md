# ConnectionOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CybersourceCard** | Pointer to [**NullableConnectionOptionsCybersourceCard**](ConnectionOptionsCybersourceCard.md) |  | [optional] 
**CybersourceAntiFraud** | Pointer to [**NullableConnectionOptionsCybersourceAntiFraud**](ConnectionOptionsCybersourceAntiFraud.md) |  | [optional] 
**ForterAntiFraud** | Pointer to [**NullableConnectionOptionsForterAntiFraud**](ConnectionOptionsForterAntiFraud.md) |  | [optional] 
**AdyenCard** | Pointer to [**NullableConnectionOptionsAdyenCard**](ConnectionOptionsAdyenCard.md) |  | [optional] 
**PaypalPaypal** | Pointer to [**NullableConnectionOptionsPaypalPaypal**](ConnectionOptionsPaypalPaypal.md) |  | [optional] 
**PaypalPaypalpaylater** | Pointer to [**NullableConnectionOptionsPaypalPaypal**](ConnectionOptionsPaypalPaypal.md) |  | [optional] 
**StripeCard** | Pointer to [**NullableConnectionOptionsStripeCard**](ConnectionOptionsStripeCard.md) |  | [optional] 

## Methods

### NewConnectionOptions

`func NewConnectionOptions() *ConnectionOptions`

NewConnectionOptions instantiates a new ConnectionOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionOptionsWithDefaults

`func NewConnectionOptionsWithDefaults() *ConnectionOptions`

NewConnectionOptionsWithDefaults instantiates a new ConnectionOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCybersourceCard

`func (o *ConnectionOptions) GetCybersourceCard() ConnectionOptionsCybersourceCard`

GetCybersourceCard returns the CybersourceCard field if non-nil, zero value otherwise.

### GetCybersourceCardOk

`func (o *ConnectionOptions) GetCybersourceCardOk() (*ConnectionOptionsCybersourceCard, bool)`

GetCybersourceCardOk returns a tuple with the CybersourceCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCybersourceCard

`func (o *ConnectionOptions) SetCybersourceCard(v ConnectionOptionsCybersourceCard)`

SetCybersourceCard sets CybersourceCard field to given value.

### HasCybersourceCard

`func (o *ConnectionOptions) HasCybersourceCard() bool`

HasCybersourceCard returns a boolean if a field has been set.

### SetCybersourceCardNil

`func (o *ConnectionOptions) SetCybersourceCardNil(b bool)`

 SetCybersourceCardNil sets the value for CybersourceCard to be an explicit nil

### UnsetCybersourceCard
`func (o *ConnectionOptions) UnsetCybersourceCard()`

UnsetCybersourceCard ensures that no value is present for CybersourceCard, not even an explicit nil
### GetCybersourceAntiFraud

`func (o *ConnectionOptions) GetCybersourceAntiFraud() ConnectionOptionsCybersourceAntiFraud`

GetCybersourceAntiFraud returns the CybersourceAntiFraud field if non-nil, zero value otherwise.

### GetCybersourceAntiFraudOk

`func (o *ConnectionOptions) GetCybersourceAntiFraudOk() (*ConnectionOptionsCybersourceAntiFraud, bool)`

GetCybersourceAntiFraudOk returns a tuple with the CybersourceAntiFraud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCybersourceAntiFraud

`func (o *ConnectionOptions) SetCybersourceAntiFraud(v ConnectionOptionsCybersourceAntiFraud)`

SetCybersourceAntiFraud sets CybersourceAntiFraud field to given value.

### HasCybersourceAntiFraud

`func (o *ConnectionOptions) HasCybersourceAntiFraud() bool`

HasCybersourceAntiFraud returns a boolean if a field has been set.

### SetCybersourceAntiFraudNil

`func (o *ConnectionOptions) SetCybersourceAntiFraudNil(b bool)`

 SetCybersourceAntiFraudNil sets the value for CybersourceAntiFraud to be an explicit nil

### UnsetCybersourceAntiFraud
`func (o *ConnectionOptions) UnsetCybersourceAntiFraud()`

UnsetCybersourceAntiFraud ensures that no value is present for CybersourceAntiFraud, not even an explicit nil
### GetForterAntiFraud

`func (o *ConnectionOptions) GetForterAntiFraud() ConnectionOptionsForterAntiFraud`

GetForterAntiFraud returns the ForterAntiFraud field if non-nil, zero value otherwise.

### GetForterAntiFraudOk

`func (o *ConnectionOptions) GetForterAntiFraudOk() (*ConnectionOptionsForterAntiFraud, bool)`

GetForterAntiFraudOk returns a tuple with the ForterAntiFraud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForterAntiFraud

`func (o *ConnectionOptions) SetForterAntiFraud(v ConnectionOptionsForterAntiFraud)`

SetForterAntiFraud sets ForterAntiFraud field to given value.

### HasForterAntiFraud

`func (o *ConnectionOptions) HasForterAntiFraud() bool`

HasForterAntiFraud returns a boolean if a field has been set.

### SetForterAntiFraudNil

`func (o *ConnectionOptions) SetForterAntiFraudNil(b bool)`

 SetForterAntiFraudNil sets the value for ForterAntiFraud to be an explicit nil

### UnsetForterAntiFraud
`func (o *ConnectionOptions) UnsetForterAntiFraud()`

UnsetForterAntiFraud ensures that no value is present for ForterAntiFraud, not even an explicit nil
### GetAdyenCard

`func (o *ConnectionOptions) GetAdyenCard() ConnectionOptionsAdyenCard`

GetAdyenCard returns the AdyenCard field if non-nil, zero value otherwise.

### GetAdyenCardOk

`func (o *ConnectionOptions) GetAdyenCardOk() (*ConnectionOptionsAdyenCard, bool)`

GetAdyenCardOk returns a tuple with the AdyenCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdyenCard

`func (o *ConnectionOptions) SetAdyenCard(v ConnectionOptionsAdyenCard)`

SetAdyenCard sets AdyenCard field to given value.

### HasAdyenCard

`func (o *ConnectionOptions) HasAdyenCard() bool`

HasAdyenCard returns a boolean if a field has been set.

### SetAdyenCardNil

`func (o *ConnectionOptions) SetAdyenCardNil(b bool)`

 SetAdyenCardNil sets the value for AdyenCard to be an explicit nil

### UnsetAdyenCard
`func (o *ConnectionOptions) UnsetAdyenCard()`

UnsetAdyenCard ensures that no value is present for AdyenCard, not even an explicit nil
### GetPaypalPaypal

`func (o *ConnectionOptions) GetPaypalPaypal() ConnectionOptionsPaypalPaypal`

GetPaypalPaypal returns the PaypalPaypal field if non-nil, zero value otherwise.

### GetPaypalPaypalOk

`func (o *ConnectionOptions) GetPaypalPaypalOk() (*ConnectionOptionsPaypalPaypal, bool)`

GetPaypalPaypalOk returns a tuple with the PaypalPaypal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalPaypal

`func (o *ConnectionOptions) SetPaypalPaypal(v ConnectionOptionsPaypalPaypal)`

SetPaypalPaypal sets PaypalPaypal field to given value.

### HasPaypalPaypal

`func (o *ConnectionOptions) HasPaypalPaypal() bool`

HasPaypalPaypal returns a boolean if a field has been set.

### SetPaypalPaypalNil

`func (o *ConnectionOptions) SetPaypalPaypalNil(b bool)`

 SetPaypalPaypalNil sets the value for PaypalPaypal to be an explicit nil

### UnsetPaypalPaypal
`func (o *ConnectionOptions) UnsetPaypalPaypal()`

UnsetPaypalPaypal ensures that no value is present for PaypalPaypal, not even an explicit nil
### GetPaypalPaypalpaylater

`func (o *ConnectionOptions) GetPaypalPaypalpaylater() ConnectionOptionsPaypalPaypal`

GetPaypalPaypalpaylater returns the PaypalPaypalpaylater field if non-nil, zero value otherwise.

### GetPaypalPaypalpaylaterOk

`func (o *ConnectionOptions) GetPaypalPaypalpaylaterOk() (*ConnectionOptionsPaypalPaypal, bool)`

GetPaypalPaypalpaylaterOk returns a tuple with the PaypalPaypalpaylater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalPaypalpaylater

`func (o *ConnectionOptions) SetPaypalPaypalpaylater(v ConnectionOptionsPaypalPaypal)`

SetPaypalPaypalpaylater sets PaypalPaypalpaylater field to given value.

### HasPaypalPaypalpaylater

`func (o *ConnectionOptions) HasPaypalPaypalpaylater() bool`

HasPaypalPaypalpaylater returns a boolean if a field has been set.

### SetPaypalPaypalpaylaterNil

`func (o *ConnectionOptions) SetPaypalPaypalpaylaterNil(b bool)`

 SetPaypalPaypalpaylaterNil sets the value for PaypalPaypalpaylater to be an explicit nil

### UnsetPaypalPaypalpaylater
`func (o *ConnectionOptions) UnsetPaypalPaypalpaylater()`

UnsetPaypalPaypalpaylater ensures that no value is present for PaypalPaypalpaylater, not even an explicit nil
### GetStripeCard

`func (o *ConnectionOptions) GetStripeCard() ConnectionOptionsStripeCard`

GetStripeCard returns the StripeCard field if non-nil, zero value otherwise.

### GetStripeCardOk

`func (o *ConnectionOptions) GetStripeCardOk() (*ConnectionOptionsStripeCard, bool)`

GetStripeCardOk returns a tuple with the StripeCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeCard

`func (o *ConnectionOptions) SetStripeCard(v ConnectionOptionsStripeCard)`

SetStripeCard sets StripeCard field to given value.

### HasStripeCard

`func (o *ConnectionOptions) HasStripeCard() bool`

HasStripeCard returns a boolean if a field has been set.

### SetStripeCardNil

`func (o *ConnectionOptions) SetStripeCardNil(b bool)`

 SetStripeCardNil sets the value for StripeCard to be an explicit nil

### UnsetStripeCard
`func (o *ConnectionOptions) UnsetStripeCard()`

UnsetStripeCard ensures that no value is present for StripeCard, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


