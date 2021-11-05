# ApplePayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | &#x60;applepay&#x60;. | 
**Token** | **map[string]interface{}** | The encrypted (opaque) token that was passed to the &#x60;onpaymentauthorized&#x60; callback by the Apple Pay integration. | 

## Methods

### NewApplePayRequest

`func NewApplePayRequest(method string, token map[string]interface{}, ) *ApplePayRequest`

NewApplePayRequest instantiates a new ApplePayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplePayRequestWithDefaults

`func NewApplePayRequestWithDefaults() *ApplePayRequest`

NewApplePayRequestWithDefaults instantiates a new ApplePayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *ApplePayRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ApplePayRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ApplePayRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetToken

`func (o *ApplePayRequest) GetToken() map[string]interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ApplePayRequest) GetTokenOk() (*map[string]interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ApplePayRequest) SetToken(v map[string]interface{})`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


