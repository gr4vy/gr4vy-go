# TokenizedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | &#x60;id&#x60;. | 
**Id** | **string** | A ID that represents a previously tokenized payment method. This token can represent any type of payment method. | 

## Methods

### NewTokenizedRequest

`func NewTokenizedRequest(method string, id string, ) *TokenizedRequest`

NewTokenizedRequest instantiates a new TokenizedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizedRequestWithDefaults

`func NewTokenizedRequestWithDefaults() *TokenizedRequest`

NewTokenizedRequestWithDefaults instantiates a new TokenizedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *TokenizedRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *TokenizedRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *TokenizedRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetId

`func (o *TokenizedRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenizedRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenizedRequest) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


