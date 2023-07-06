# AntiFraudDecisionTransactionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;transaction-event&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this event. | [optional] 
**Name** | Pointer to **string** | The name of this resource. Is always &#x60;anti-fraud-decision &#x60;. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this transaction was created in our system. | [optional] 
**Context** | Pointer to [**AntiFraudDecisionTransactionEventContext**](AntiFraudDecisionTransactionEventContext.md) |  | [optional] 

## Methods

### NewAntiFraudDecisionTransactionEvent

`func NewAntiFraudDecisionTransactionEvent() *AntiFraudDecisionTransactionEvent`

NewAntiFraudDecisionTransactionEvent instantiates a new AntiFraudDecisionTransactionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntiFraudDecisionTransactionEventWithDefaults

`func NewAntiFraudDecisionTransactionEventWithDefaults() *AntiFraudDecisionTransactionEvent`

NewAntiFraudDecisionTransactionEventWithDefaults instantiates a new AntiFraudDecisionTransactionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AntiFraudDecisionTransactionEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AntiFraudDecisionTransactionEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AntiFraudDecisionTransactionEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AntiFraudDecisionTransactionEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *AntiFraudDecisionTransactionEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AntiFraudDecisionTransactionEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AntiFraudDecisionTransactionEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AntiFraudDecisionTransactionEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AntiFraudDecisionTransactionEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AntiFraudDecisionTransactionEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AntiFraudDecisionTransactionEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AntiFraudDecisionTransactionEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AntiFraudDecisionTransactionEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AntiFraudDecisionTransactionEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AntiFraudDecisionTransactionEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AntiFraudDecisionTransactionEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetContext

`func (o *AntiFraudDecisionTransactionEvent) GetContext() AntiFraudDecisionTransactionEventContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AntiFraudDecisionTransactionEvent) GetContextOk() (*AntiFraudDecisionTransactionEventContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AntiFraudDecisionTransactionEvent) SetContext(v AntiFraudDecisionTransactionEventContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *AntiFraudDecisionTransactionEvent) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


