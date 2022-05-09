# TransactionRefundRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | The amount requested to refund.  If omitted, a full refund will be requested.  Otherwise, the amount must be lower than or equal to the remaining balance in the associated transaction.  Negative and zero-amount refunds are not supported. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


