# AntiFraudService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;anti-fraud-service&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique Gr4vy ID for this anti-fraud service. | [optional] 
**AntiFraudServiceDefinitionId** | Pointer to **string** | The name of the Anti-Fraud service provider. During update request, this value is used for validation only but the underlying service can not be changed for an existing service. | [optional] 
**DisplayName** | Pointer to **NullableString** | A unique name for this anti-fraud service which is used in the Gr4vy admin panel to give a anti-fraud service a human readable name. | [optional] 
**Active** | Pointer to **bool** | Defines if this service is currently active or not. | [optional] [default to true]
**Fields** | Pointer to [**[]AntiFraudServiceFields**](AntiFraudServiceFields.md) | A list of fields, each containing a key-value pair for anti-fraud service decision mapping e.g. for sift &#x60;approve_decision&#x60; will be in the response. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this anti-fraud service was created in our system. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when this anti-fraud service was last updated in our system. | [optional] 

## Methods

### NewAntiFraudService

`func NewAntiFraudService() *AntiFraudService`

NewAntiFraudService instantiates a new AntiFraudService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntiFraudServiceWithDefaults

`func NewAntiFraudServiceWithDefaults() *AntiFraudService`

NewAntiFraudServiceWithDefaults instantiates a new AntiFraudService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AntiFraudService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AntiFraudService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AntiFraudService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AntiFraudService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *AntiFraudService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AntiFraudService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AntiFraudService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AntiFraudService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAntiFraudServiceDefinitionId

`func (o *AntiFraudService) GetAntiFraudServiceDefinitionId() string`

GetAntiFraudServiceDefinitionId returns the AntiFraudServiceDefinitionId field if non-nil, zero value otherwise.

### GetAntiFraudServiceDefinitionIdOk

`func (o *AntiFraudService) GetAntiFraudServiceDefinitionIdOk() (*string, bool)`

GetAntiFraudServiceDefinitionIdOk returns a tuple with the AntiFraudServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceDefinitionId

`func (o *AntiFraudService) SetAntiFraudServiceDefinitionId(v string)`

SetAntiFraudServiceDefinitionId sets AntiFraudServiceDefinitionId field to given value.

### HasAntiFraudServiceDefinitionId

`func (o *AntiFraudService) HasAntiFraudServiceDefinitionId() bool`

HasAntiFraudServiceDefinitionId returns a boolean if a field has been set.

### GetDisplayName

`func (o *AntiFraudService) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AntiFraudService) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AntiFraudService) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AntiFraudService) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AntiFraudService) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AntiFraudService) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetActive

`func (o *AntiFraudService) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AntiFraudService) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AntiFraudService) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AntiFraudService) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetFields

`func (o *AntiFraudService) GetFields() []AntiFraudServiceFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AntiFraudService) GetFieldsOk() (*[]AntiFraudServiceFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AntiFraudService) SetFields(v []AntiFraudServiceFields)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AntiFraudService) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AntiFraudService) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AntiFraudService) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AntiFraudService) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AntiFraudService) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AntiFraudService) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AntiFraudService) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AntiFraudService) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AntiFraudService) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


