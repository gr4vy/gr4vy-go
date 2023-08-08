# ConnectionOptionsForterAntiFraud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryType** | Pointer to **NullableString** | Value to populate the &#x60;deliveryType&#x60; field in &#x60;primaryDeliveryDetails&#x60;.  Represents the type of delivery. This can be set to &#x60;PHYSICAL&#x60; for any type of shipped goods, &#x60;DIGITAL&#x60; for non-shipped goods (services, gift cards etc.), or &#x60;HYBRID&#x60; for others. | [optional] 
**DeliveryMethod** | Pointer to **NullableString** | Value to populate the &#x60;deliveryMethod&#x60; field in &#x60;primaryDeliveryDetails&#x60;.  Represents the delivery method chosen by customer such as postal service, email, in game transfer, etc. | [optional] 

## Methods

### NewConnectionOptionsForterAntiFraud

`func NewConnectionOptionsForterAntiFraud() *ConnectionOptionsForterAntiFraud`

NewConnectionOptionsForterAntiFraud instantiates a new ConnectionOptionsForterAntiFraud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionOptionsForterAntiFraudWithDefaults

`func NewConnectionOptionsForterAntiFraudWithDefaults() *ConnectionOptionsForterAntiFraud`

NewConnectionOptionsForterAntiFraudWithDefaults instantiates a new ConnectionOptionsForterAntiFraud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryType

`func (o *ConnectionOptionsForterAntiFraud) GetDeliveryType() string`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *ConnectionOptionsForterAntiFraud) GetDeliveryTypeOk() (*string, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *ConnectionOptionsForterAntiFraud) SetDeliveryType(v string)`

SetDeliveryType sets DeliveryType field to given value.

### HasDeliveryType

`func (o *ConnectionOptionsForterAntiFraud) HasDeliveryType() bool`

HasDeliveryType returns a boolean if a field has been set.

### SetDeliveryTypeNil

`func (o *ConnectionOptionsForterAntiFraud) SetDeliveryTypeNil(b bool)`

 SetDeliveryTypeNil sets the value for DeliveryType to be an explicit nil

### UnsetDeliveryType
`func (o *ConnectionOptionsForterAntiFraud) UnsetDeliveryType()`

UnsetDeliveryType ensures that no value is present for DeliveryType, not even an explicit nil
### GetDeliveryMethod

`func (o *ConnectionOptionsForterAntiFraud) GetDeliveryMethod() string`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *ConnectionOptionsForterAntiFraud) GetDeliveryMethodOk() (*string, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *ConnectionOptionsForterAntiFraud) SetDeliveryMethod(v string)`

SetDeliveryMethod sets DeliveryMethod field to given value.

### HasDeliveryMethod

`func (o *ConnectionOptionsForterAntiFraud) HasDeliveryMethod() bool`

HasDeliveryMethod returns a boolean if a field has been set.

### SetDeliveryMethodNil

`func (o *ConnectionOptionsForterAntiFraud) SetDeliveryMethodNil(b bool)`

 SetDeliveryMethodNil sets the value for DeliveryMethod to be an explicit nil

### UnsetDeliveryMethod
`func (o *ConnectionOptionsForterAntiFraud) UnsetDeliveryMethod()`

UnsetDeliveryMethod ensures that no value is present for DeliveryMethod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


