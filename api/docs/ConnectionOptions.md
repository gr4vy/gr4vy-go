# ConnectionOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CybersourceAntiFraud** | Pointer to [**NullableConnectionOptionsCybersourceAntiFraud**](ConnectionOptionsCybersourceAntiFraud.md) |  | [optional] 
**AdyenCard** | Pointer to [**NullableConnectionOptionsAdyenCard**](ConnectionOptionsAdyenCard.md) |  | [optional] 

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


