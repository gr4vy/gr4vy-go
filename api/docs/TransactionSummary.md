# TransactionSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;transaction&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this transaction. | [optional] 
**Status** | Pointer to **string** | The status of the transaction. The status may change over time as asynchronous  processing events occur. | [optional] 
**Amount** | Pointer to **int32** | The authorized amount for this transaction. This can be different than the actual captured amount and part of this amount may be refunded. | [optional] 
**CapturedAmount** | Pointer to **int32** | The captured amount for this transaction. This can be a part and in some cases even more than the authorized amount. | [optional] 
**RefundedAmount** | Pointer to **int32** | The refunded amount for this transaction. This can be a part or all of the captured amount. | [optional] 
**Currency** | Pointer to **string** | The currency code for this transaction. | [optional] 
**PaymentMethod** | Pointer to [**PaymentMethodSnapshot**](PaymentMethod--Snapshot.md) |  | [optional] 
**Buyer** | Pointer to [**BuyerSnapshot**](Buyer--Snapshot.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this transaction was created in our system. | [optional] 
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the transaction against your own records. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Defines when the transaction was last updated. | [optional] 
**PaymentService** | Pointer to [**PaymentServiceSnapshot**](PaymentService--Snapshot.md) |  | [optional] 

## Methods

### NewTransactionSummary

`func NewTransactionSummary() *TransactionSummary`

NewTransactionSummary instantiates a new TransactionSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSummaryWithDefaults

`func NewTransactionSummaryWithDefaults() *TransactionSummary`

NewTransactionSummaryWithDefaults instantiates a new TransactionSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransactionSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *TransactionSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *TransactionSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransactionSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionSummary) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionSummary) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionSummary) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionSummary) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCapturedAmount

`func (o *TransactionSummary) GetCapturedAmount() int32`

GetCapturedAmount returns the CapturedAmount field if non-nil, zero value otherwise.

### GetCapturedAmountOk

`func (o *TransactionSummary) GetCapturedAmountOk() (*int32, bool)`

GetCapturedAmountOk returns a tuple with the CapturedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedAmount

`func (o *TransactionSummary) SetCapturedAmount(v int32)`

SetCapturedAmount sets CapturedAmount field to given value.

### HasCapturedAmount

`func (o *TransactionSummary) HasCapturedAmount() bool`

HasCapturedAmount returns a boolean if a field has been set.

### GetRefundedAmount

`func (o *TransactionSummary) GetRefundedAmount() int32`

GetRefundedAmount returns the RefundedAmount field if non-nil, zero value otherwise.

### GetRefundedAmountOk

`func (o *TransactionSummary) GetRefundedAmountOk() (*int32, bool)`

GetRefundedAmountOk returns a tuple with the RefundedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAmount

`func (o *TransactionSummary) SetRefundedAmount(v int32)`

SetRefundedAmount sets RefundedAmount field to given value.

### HasRefundedAmount

`func (o *TransactionSummary) HasRefundedAmount() bool`

HasRefundedAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *TransactionSummary) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionSummary) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionSummary) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *TransactionSummary) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *TransactionSummary) GetPaymentMethod() PaymentMethodSnapshot`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *TransactionSummary) GetPaymentMethodOk() (*PaymentMethodSnapshot, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *TransactionSummary) SetPaymentMethod(v PaymentMethodSnapshot)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *TransactionSummary) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetBuyer

`func (o *TransactionSummary) GetBuyer() BuyerSnapshot`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *TransactionSummary) GetBuyerOk() (*BuyerSnapshot, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *TransactionSummary) SetBuyer(v BuyerSnapshot)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *TransactionSummary) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TransactionSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransactionSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransactionSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TransactionSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExternalIdentifier

`func (o *TransactionSummary) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *TransactionSummary) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *TransactionSummary) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *TransactionSummary) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *TransactionSummary) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *TransactionSummary) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetUpdatedAt

`func (o *TransactionSummary) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TransactionSummary) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TransactionSummary) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TransactionSummary) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPaymentService

`func (o *TransactionSummary) GetPaymentService() PaymentServiceSnapshot`

GetPaymentService returns the PaymentService field if non-nil, zero value otherwise.

### GetPaymentServiceOk

`func (o *TransactionSummary) GetPaymentServiceOk() (*PaymentServiceSnapshot, bool)`

GetPaymentServiceOk returns a tuple with the PaymentService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentService

`func (o *TransactionSummary) SetPaymentService(v PaymentServiceSnapshot)`

SetPaymentService sets PaymentService field to given value.

### HasPaymentService

`func (o *TransactionSummary) HasPaymentService() bool`

HasPaymentService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


