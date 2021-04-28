# Error404PendingCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;error&#x60;. | [optional] 
**Code** | Pointer to **string** | &#x60;pending_creation&#x60;. | [optional] 
**Status** | Pointer to **int32** | &#x60;404&#x60;. | [optional] 
**Message** | Pointer to **string** | The resource is still pending. | [optional] 
**Details** | Pointer to [**[]ErrorDetail**](ErrorDetail.md) | A list of detail objects that further clarify the reason for the error. Not every error supports more detail. | [optional] 

## Methods

### NewError404PendingCreation

`func NewError404PendingCreation() *Error404PendingCreation`

NewError404PendingCreation instantiates a new Error404PendingCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewError404PendingCreationWithDefaults

`func NewError404PendingCreationWithDefaults() *Error404PendingCreation`

NewError404PendingCreationWithDefaults instantiates a new Error404PendingCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Error404PendingCreation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Error404PendingCreation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Error404PendingCreation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Error404PendingCreation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *Error404PendingCreation) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Error404PendingCreation) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Error404PendingCreation) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Error404PendingCreation) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetStatus

`func (o *Error404PendingCreation) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Error404PendingCreation) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Error404PendingCreation) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Error404PendingCreation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *Error404PendingCreation) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Error404PendingCreation) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Error404PendingCreation) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Error404PendingCreation) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetails

`func (o *Error404PendingCreation) GetDetails() []ErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Error404PendingCreation) GetDetailsOk() (*[]ErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Error404PendingCreation) SetDetails(v []ErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Error404PendingCreation) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


