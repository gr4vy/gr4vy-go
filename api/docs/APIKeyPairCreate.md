# APIKeyPairCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | A name for this key-pair which is used in the Gr4vy admin panel to give the key-pair a human readable name. | [optional] 

## Methods

### NewAPIKeyPairCreate

`func NewAPIKeyPairCreate() *APIKeyPairCreate`

NewAPIKeyPairCreate instantiates a new APIKeyPairCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIKeyPairCreateWithDefaults

`func NewAPIKeyPairCreateWithDefaults() *APIKeyPairCreate`

NewAPIKeyPairCreateWithDefaults instantiates a new APIKeyPairCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *APIKeyPairCreate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *APIKeyPairCreate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *APIKeyPairCreate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *APIKeyPairCreate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


