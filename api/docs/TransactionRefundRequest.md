# TransactionRefundRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | The amount requested to refund.  If omitted, a full refund will be requested for the main payment method.  When set, the amount must be lower than or equal to the remaining balance in the associated transaction. Negative and zero-amount refunds are not supported. | [optional] 
**TargetType** | Pointer to **NullableString** | The target type to refund for. This can be used to target a gift card to refund to instead of the main payment method. | [optional] [default to "payment-method"]
**TargetId** | Pointer to **NullableString** | The optional ID of the instrument to refund for. This is only required when the &#x60;target_type&#x60; is set to &#x60;gift-card-redemption&#x60;. | [optional] 

## Methods

### NewTransactionRefundRequest

`func NewTransactionRefundRequest() *TransactionRefundRequest`

NewTransactionRefundRequest instantiates a new TransactionRefundRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRefundRequestWithDefaults

`func NewTransactionRefundRequestWithDefaults() *TransactionRefundRequest`

NewTransactionRefundRequestWithDefaults instantiates a new TransactionRefundRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransactionRefundRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionRefundRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionRefundRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionRefundRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTargetType

`func (o *TransactionRefundRequest) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *TransactionRefundRequest) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *TransactionRefundRequest) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *TransactionRefundRequest) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *TransactionRefundRequest) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *TransactionRefundRequest) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTargetId

`func (o *TransactionRefundRequest) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *TransactionRefundRequest) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *TransactionRefundRequest) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *TransactionRefundRequest) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *TransactionRefundRequest) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *TransactionRefundRequest) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


