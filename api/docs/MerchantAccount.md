# MerchantAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;merchant-account&#x60;. | [optional] 
**Id** | Pointer to **string** | The ID for this merchant account. | [optional] 
**DisplayName** | Pointer to **string** | The display name of this merchant account. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this merchant account was created in our system. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when this merchant account was updated in our system. | [optional] 

## Methods

### NewMerchantAccount

`func NewMerchantAccount() *MerchantAccount`

NewMerchantAccount instantiates a new MerchantAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantAccountWithDefaults

`func NewMerchantAccountWithDefaults() *MerchantAccount`

NewMerchantAccountWithDefaults instantiates a new MerchantAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MerchantAccount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MerchantAccount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MerchantAccount) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MerchantAccount) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *MerchantAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MerchantAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MerchantAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MerchantAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *MerchantAccount) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MerchantAccount) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MerchantAccount) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MerchantAccount) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MerchantAccount) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MerchantAccount) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MerchantAccount) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MerchantAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MerchantAccount) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MerchantAccount) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MerchantAccount) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MerchantAccount) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


