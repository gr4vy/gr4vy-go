# ThreeDSecureDataV2AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationResponse** | Pointer to **NullableString** | The transaction status after a the 3DS challenge. This will be null in case of a frictionless 3DS flow. | [optional] 
**DirectoryResponse** | Pointer to **string** | The transaction status received as part of the authentication request. | [optional] 
**DirectoryTransactionId** | **string** | The transaction identifier. | 

## Methods

### NewThreeDSecureDataV2AllOf

`func NewThreeDSecureDataV2AllOf(directoryTransactionId string, ) *ThreeDSecureDataV2AllOf`

NewThreeDSecureDataV2AllOf instantiates a new ThreeDSecureDataV2AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSecureDataV2AllOfWithDefaults

`func NewThreeDSecureDataV2AllOfWithDefaults() *ThreeDSecureDataV2AllOf`

NewThreeDSecureDataV2AllOfWithDefaults instantiates a new ThreeDSecureDataV2AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationResponse

`func (o *ThreeDSecureDataV2AllOf) GetAuthenticationResponse() string`

GetAuthenticationResponse returns the AuthenticationResponse field if non-nil, zero value otherwise.

### GetAuthenticationResponseOk

`func (o *ThreeDSecureDataV2AllOf) GetAuthenticationResponseOk() (*string, bool)`

GetAuthenticationResponseOk returns a tuple with the AuthenticationResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationResponse

`func (o *ThreeDSecureDataV2AllOf) SetAuthenticationResponse(v string)`

SetAuthenticationResponse sets AuthenticationResponse field to given value.

### HasAuthenticationResponse

`func (o *ThreeDSecureDataV2AllOf) HasAuthenticationResponse() bool`

HasAuthenticationResponse returns a boolean if a field has been set.

### SetAuthenticationResponseNil

`func (o *ThreeDSecureDataV2AllOf) SetAuthenticationResponseNil(b bool)`

 SetAuthenticationResponseNil sets the value for AuthenticationResponse to be an explicit nil

### UnsetAuthenticationResponse
`func (o *ThreeDSecureDataV2AllOf) UnsetAuthenticationResponse()`

UnsetAuthenticationResponse ensures that no value is present for AuthenticationResponse, not even an explicit nil
### GetDirectoryResponse

`func (o *ThreeDSecureDataV2AllOf) GetDirectoryResponse() string`

GetDirectoryResponse returns the DirectoryResponse field if non-nil, zero value otherwise.

### GetDirectoryResponseOk

`func (o *ThreeDSecureDataV2AllOf) GetDirectoryResponseOk() (*string, bool)`

GetDirectoryResponseOk returns a tuple with the DirectoryResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryResponse

`func (o *ThreeDSecureDataV2AllOf) SetDirectoryResponse(v string)`

SetDirectoryResponse sets DirectoryResponse field to given value.

### HasDirectoryResponse

`func (o *ThreeDSecureDataV2AllOf) HasDirectoryResponse() bool`

HasDirectoryResponse returns a boolean if a field has been set.

### GetDirectoryTransactionId

`func (o *ThreeDSecureDataV2AllOf) GetDirectoryTransactionId() string`

GetDirectoryTransactionId returns the DirectoryTransactionId field if non-nil, zero value otherwise.

### GetDirectoryTransactionIdOk

`func (o *ThreeDSecureDataV2AllOf) GetDirectoryTransactionIdOk() (*string, bool)`

GetDirectoryTransactionIdOk returns a tuple with the DirectoryTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryTransactionId

`func (o *ThreeDSecureDataV2AllOf) SetDirectoryTransactionId(v string)`

SetDirectoryTransactionId sets DirectoryTransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


