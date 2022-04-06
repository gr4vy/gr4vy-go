# Error400InvalidCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;error&#x60;. | [optional] 
**Code** | Pointer to **string** | &#x60;invalid_credentials&#x60;. | [optional] 
**Status** | Pointer to **int32** | &#x60;400&#x60;. | [optional] 
**Message** | Pointer to **string** | The provided credentials are invalid. | [optional] 

## Methods

### NewError400InvalidCredentials

`func NewError400InvalidCredentials() *Error400InvalidCredentials`

NewError400InvalidCredentials instantiates a new Error400InvalidCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewError400InvalidCredentialsWithDefaults

`func NewError400InvalidCredentialsWithDefaults() *Error400InvalidCredentials`

NewError400InvalidCredentialsWithDefaults instantiates a new Error400InvalidCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Error400InvalidCredentials) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Error400InvalidCredentials) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Error400InvalidCredentials) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Error400InvalidCredentials) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *Error400InvalidCredentials) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Error400InvalidCredentials) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Error400InvalidCredentials) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Error400InvalidCredentials) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetStatus

`func (o *Error400InvalidCredentials) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Error400InvalidCredentials) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Error400InvalidCredentials) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Error400InvalidCredentials) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *Error400InvalidCredentials) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Error400InvalidCredentials) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Error400InvalidCredentials) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Error400InvalidCredentials) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


