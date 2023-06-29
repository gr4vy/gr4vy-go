# ThreeDSecureSuccessTransactionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;transaction-event&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this event. | [optional] 
**Name** | Pointer to **string** | The name of this resource. Is always &#x60;three-d-secure-success&#x60;. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this transaction was created in our system. | [optional] 
**Context** | Pointer to [**ThreeDSecureSuccessTransactionEventContext**](ThreeDSecureSuccessTransactionEventContext.md) |  | [optional] 

## Methods

### NewThreeDSecureSuccessTransactionEvent

`func NewThreeDSecureSuccessTransactionEvent() *ThreeDSecureSuccessTransactionEvent`

NewThreeDSecureSuccessTransactionEvent instantiates a new ThreeDSecureSuccessTransactionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSecureSuccessTransactionEventWithDefaults

`func NewThreeDSecureSuccessTransactionEventWithDefaults() *ThreeDSecureSuccessTransactionEvent`

NewThreeDSecureSuccessTransactionEventWithDefaults instantiates a new ThreeDSecureSuccessTransactionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ThreeDSecureSuccessTransactionEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThreeDSecureSuccessTransactionEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThreeDSecureSuccessTransactionEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ThreeDSecureSuccessTransactionEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *ThreeDSecureSuccessTransactionEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThreeDSecureSuccessTransactionEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThreeDSecureSuccessTransactionEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ThreeDSecureSuccessTransactionEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ThreeDSecureSuccessTransactionEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThreeDSecureSuccessTransactionEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThreeDSecureSuccessTransactionEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ThreeDSecureSuccessTransactionEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ThreeDSecureSuccessTransactionEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ThreeDSecureSuccessTransactionEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ThreeDSecureSuccessTransactionEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ThreeDSecureSuccessTransactionEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetContext

`func (o *ThreeDSecureSuccessTransactionEvent) GetContext() ThreeDSecureSuccessTransactionEventContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ThreeDSecureSuccessTransactionEvent) GetContextOk() (*ThreeDSecureSuccessTransactionEventContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ThreeDSecureSuccessTransactionEvent) SetContext(v ThreeDSecureSuccessTransactionEventContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *ThreeDSecureSuccessTransactionEvent) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


