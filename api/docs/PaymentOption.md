# PaymentOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;payment-option&#x60;. | [optional] 
**Method** | Pointer to **string** | The type of payment method that is available. | [optional] 

## Methods

### NewPaymentOption

`func NewPaymentOption() *PaymentOption`

NewPaymentOption instantiates a new PaymentOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentOptionWithDefaults

`func NewPaymentOptionWithDefaults() *PaymentOption`

NewPaymentOptionWithDefaults instantiates a new PaymentOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentOption) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentOption) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMethod

`func (o *PaymentOption) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PaymentOption) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PaymentOption) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PaymentOption) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


