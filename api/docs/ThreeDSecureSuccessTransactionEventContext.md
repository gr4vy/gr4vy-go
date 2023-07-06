# ThreeDSecureSuccessTransactionEventContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eci** | Pointer to **string** | The electronic commerce indicator for the 3DS transaction. | [optional] 
**Cavv** | Pointer to **string** | The cardholder authentication value or AAV. | [optional] 
**Version** | Pointer to **string** | The version of 3-D Secure that was used. | [optional] 
**DirectoryResponse** | Pointer to **string** | For 3-D Secure version 1, the enrolment response. For 3-D Secure version , the transaction status from the &#x60;ARes&#x60;. | [optional] 
**AuthenticationResponse** | Pointer to **NullableString** | The transaction status from the challenge result (not required for frictionless). | [optional] 
**DirectoryTransactionId** | Pointer to **NullableString** | The transaction identifier. | [optional] 
**CavvAlgorithm** | Pointer to **NullableString** | The CAVV Algorithm used. | [optional] 
**Method** | Pointer to **NullableString** | The method used for 3DS authentication for this transaction. | [optional] 

## Methods

### NewThreeDSecureSuccessTransactionEventContext

`func NewThreeDSecureSuccessTransactionEventContext() *ThreeDSecureSuccessTransactionEventContext`

NewThreeDSecureSuccessTransactionEventContext instantiates a new ThreeDSecureSuccessTransactionEventContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSecureSuccessTransactionEventContextWithDefaults

`func NewThreeDSecureSuccessTransactionEventContextWithDefaults() *ThreeDSecureSuccessTransactionEventContext`

NewThreeDSecureSuccessTransactionEventContextWithDefaults instantiates a new ThreeDSecureSuccessTransactionEventContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEci

`func (o *ThreeDSecureSuccessTransactionEventContext) GetEci() string`

GetEci returns the Eci field if non-nil, zero value otherwise.

### GetEciOk

`func (o *ThreeDSecureSuccessTransactionEventContext) GetEciOk() (*string, bool)`

GetEciOk returns a tuple with the Eci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEci

`func (o *ThreeDSecureSuccessTransactionEventContext) SetEci(v string)`

SetEci sets Eci field to given value.

### HasEci

`func (o *ThreeDSecureSuccessTransactionEventContext) HasEci() bool`

HasEci returns a boolean if a field has been set.

### GetCavv

`func (o *ThreeDSecureSuccessTransactionEventContext) GetCavv() string`

GetCavv returns the Cavv field if non-nil, zero value otherwise.

### GetCavvOk

`func (o *ThreeDSecureSuccessTransactionEventContext) GetCavvOk() (*string, bool)`

GetCavvOk returns a tuple with the Cavv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCavv

`func (o *ThreeDSecureSuccessTransactionEventContext) SetCavv(v string)`

SetCavv sets Cavv field to given value.

### HasCavv

`func (o *ThreeDSecureSuccessTransactionEventContext) HasCavv() bool`

HasCavv returns a boolean if a field has been set.

### GetVersion

`func (o *ThreeDSecureSuccessTransactionEventContext) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThreeDSecureSuccessTransactionEventContext) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThreeDSecureSuccessTransactionEventContext) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ThreeDSecureSuccessTransactionEventContext) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDirectoryResponse

`func (o *ThreeDSecureSuccessTransactionEventContext) GetDirectoryResponse() string`

GetDirectoryResponse returns the DirectoryResponse field if non-nil, zero value otherwise.

### GetDirectoryResponseOk

`func (o *ThreeDSecureSuccessTransactionEventContext) GetDirectoryResponseOk() (*string, bool)`

GetDirectoryResponseOk returns a tuple with the DirectoryResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryResponse

`func (o *ThreeDSecureSuccessTransactionEventContext) SetDirectoryResponse(v string)`

SetDirectoryResponse sets DirectoryResponse field to given value.

### HasDirectoryResponse

`func (o *ThreeDSecureSuccessTransactionEventContext) HasDirectoryResponse() bool`

HasDirectoryResponse returns a boolean if a field has been set.

### GetAuthenticationResponse

`func (o *ThreeDSecureSuccessTransactionEventContext) GetAuthenticationResponse() string`

GetAuthenticationResponse returns the AuthenticationResponse field if non-nil, zero value otherwise.

### GetAuthenticationResponseOk

`func (o *ThreeDSecureSuccessTransactionEventContext) GetAuthenticationResponseOk() (*string, bool)`

GetAuthenticationResponseOk returns a tuple with the AuthenticationResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationResponse

`func (o *ThreeDSecureSuccessTransactionEventContext) SetAuthenticationResponse(v string)`

SetAuthenticationResponse sets AuthenticationResponse field to given value.

### HasAuthenticationResponse

`func (o *ThreeDSecureSuccessTransactionEventContext) HasAuthenticationResponse() bool`

HasAuthenticationResponse returns a boolean if a field has been set.

### SetAuthenticationResponseNil

`func (o *ThreeDSecureSuccessTransactionEventContext) SetAuthenticationResponseNil(b bool)`

 SetAuthenticationResponseNil sets the value for AuthenticationResponse to be an explicit nil

### UnsetAuthenticationResponse
`func (o *ThreeDSecureSuccessTransactionEventContext) UnsetAuthenticationResponse()`

UnsetAuthenticationResponse ensures that no value is present for AuthenticationResponse, not even an explicit nil
### GetDirectoryTransactionId

`func (o *ThreeDSecureSuccessTransactionEventContext) GetDirectoryTransactionId() string`

GetDirectoryTransactionId returns the DirectoryTransactionId field if non-nil, zero value otherwise.

### GetDirectoryTransactionIdOk

`func (o *ThreeDSecureSuccessTransactionEventContext) GetDirectoryTransactionIdOk() (*string, bool)`

GetDirectoryTransactionIdOk returns a tuple with the DirectoryTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryTransactionId

`func (o *ThreeDSecureSuccessTransactionEventContext) SetDirectoryTransactionId(v string)`

SetDirectoryTransactionId sets DirectoryTransactionId field to given value.

### HasDirectoryTransactionId

`func (o *ThreeDSecureSuccessTransactionEventContext) HasDirectoryTransactionId() bool`

HasDirectoryTransactionId returns a boolean if a field has been set.

### SetDirectoryTransactionIdNil

`func (o *ThreeDSecureSuccessTransactionEventContext) SetDirectoryTransactionIdNil(b bool)`

 SetDirectoryTransactionIdNil sets the value for DirectoryTransactionId to be an explicit nil

### UnsetDirectoryTransactionId
`func (o *ThreeDSecureSuccessTransactionEventContext) UnsetDirectoryTransactionId()`

UnsetDirectoryTransactionId ensures that no value is present for DirectoryTransactionId, not even an explicit nil
### GetCavvAlgorithm

`func (o *ThreeDSecureSuccessTransactionEventContext) GetCavvAlgorithm() string`

GetCavvAlgorithm returns the CavvAlgorithm field if non-nil, zero value otherwise.

### GetCavvAlgorithmOk

`func (o *ThreeDSecureSuccessTransactionEventContext) GetCavvAlgorithmOk() (*string, bool)`

GetCavvAlgorithmOk returns a tuple with the CavvAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCavvAlgorithm

`func (o *ThreeDSecureSuccessTransactionEventContext) SetCavvAlgorithm(v string)`

SetCavvAlgorithm sets CavvAlgorithm field to given value.

### HasCavvAlgorithm

`func (o *ThreeDSecureSuccessTransactionEventContext) HasCavvAlgorithm() bool`

HasCavvAlgorithm returns a boolean if a field has been set.

### SetCavvAlgorithmNil

`func (o *ThreeDSecureSuccessTransactionEventContext) SetCavvAlgorithmNil(b bool)`

 SetCavvAlgorithmNil sets the value for CavvAlgorithm to be an explicit nil

### UnsetCavvAlgorithm
`func (o *ThreeDSecureSuccessTransactionEventContext) UnsetCavvAlgorithm()`

UnsetCavvAlgorithm ensures that no value is present for CavvAlgorithm, not even an explicit nil
### GetMethod

`func (o *ThreeDSecureSuccessTransactionEventContext) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ThreeDSecureSuccessTransactionEventContext) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ThreeDSecureSuccessTransactionEventContext) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ThreeDSecureSuccessTransactionEventContext) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *ThreeDSecureSuccessTransactionEventContext) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *ThreeDSecureSuccessTransactionEventContext) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


