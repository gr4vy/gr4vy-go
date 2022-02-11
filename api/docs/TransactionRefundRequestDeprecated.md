# TransactionRefundRequestDeprecated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | The (partial) amount to refund.  When omitted blank, this will refund the entire amount. | [optional] 

## Methods

### NewTransactionRefundRequestDeprecated

`func NewTransactionRefundRequestDeprecated() *TransactionRefundRequestDeprecated`

NewTransactionRefundRequestDeprecated instantiates a new TransactionRefundRequestDeprecated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRefundRequestDeprecatedWithDefaults

`func NewTransactionRefundRequestDeprecatedWithDefaults() *TransactionRefundRequestDeprecated`

NewTransactionRefundRequestDeprecatedWithDefaults instantiates a new TransactionRefundRequestDeprecated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransactionRefundRequestDeprecated) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionRefundRequestDeprecated) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionRefundRequestDeprecated) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionRefundRequestDeprecated) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


