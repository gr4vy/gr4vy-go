# ErrorGeneric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this object. This is always &#x60;error&#x60;. | [optional] [default to "error"]
**Code** | Pointer to **string** | A custom code to further describe the type of error being returned. This code provides further specification within the HTTP &#x60;status&#x60; code and can be used by a program to define logic based on the error. | [optional] 
**Status** | Pointer to **int32** | The HTTP status code of this error. | [optional] 
**Message** | Pointer to **string** | A human readable message that describes the error. The content of this field should not be used to determine any business logic.  | [optional] 
**Details** | Pointer to [**[]ErrorDetail**](ErrorDetail.md) | A list of detail objects that further clarify the reason for the error. Not every error supports more detail. | [optional] 

## Methods

### NewErrorGeneric

`func NewErrorGeneric() *ErrorGeneric`

NewErrorGeneric instantiates a new ErrorGeneric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorGenericWithDefaults

`func NewErrorGenericWithDefaults() *ErrorGeneric`

NewErrorGenericWithDefaults instantiates a new ErrorGeneric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorGeneric) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorGeneric) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorGeneric) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ErrorGeneric) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *ErrorGeneric) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorGeneric) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorGeneric) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorGeneric) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetStatus

`func (o *ErrorGeneric) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorGeneric) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorGeneric) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ErrorGeneric) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorGeneric) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorGeneric) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorGeneric) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorGeneric) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetails

`func (o *ErrorGeneric) GetDetails() []ErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ErrorGeneric) GetDetailsOk() (*[]ErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ErrorGeneric) SetDetails(v []ErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ErrorGeneric) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


