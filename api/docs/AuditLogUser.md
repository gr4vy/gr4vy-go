# AuditLogUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the user. | [optional] 
**Name** | Pointer to **string** | The name of the user. | [optional] 
**Staff** | Pointer to **bool** | Whether the user is Gr4vy staff. | [optional] 

## Methods

### NewAuditLogUser

`func NewAuditLogUser() *AuditLogUser`

NewAuditLogUser instantiates a new AuditLogUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogUserWithDefaults

`func NewAuditLogUserWithDefaults() *AuditLogUser`

NewAuditLogUserWithDefaults instantiates a new AuditLogUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditLogUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditLogUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditLogUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuditLogUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AuditLogUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditLogUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditLogUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuditLogUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStaff

`func (o *AuditLogUser) GetStaff() bool`

GetStaff returns the Staff field if non-nil, zero value otherwise.

### GetStaffOk

`func (o *AuditLogUser) GetStaffOk() (*bool, bool)`

GetStaffOk returns a tuple with the Staff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaff

`func (o *AuditLogUser) SetStaff(v bool)`

SetStaff sets Staff field to given value.

### HasStaff

`func (o *AuditLogUser) HasStaff() bool`

HasStaff returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


