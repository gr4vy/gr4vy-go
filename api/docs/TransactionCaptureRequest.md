# TransactionCaptureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | The (partial) amount to capture.  When left blank, this will capture the entire amount. | 
**Currency** | **string** | A supported ISO-4217 currency code.  | 
**ExternalIdentifier** | Pointer to **string** | An external identifier that can be used to match the transaction against your own records. | [optional] 

## Methods

### NewTransactionCaptureRequest

`func NewTransactionCaptureRequest(amount float32, currency string, ) *TransactionCaptureRequest`

NewTransactionCaptureRequest instantiates a new TransactionCaptureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionCaptureRequestWithDefaults

`func NewTransactionCaptureRequestWithDefaults() *TransactionCaptureRequest`

NewTransactionCaptureRequestWithDefaults instantiates a new TransactionCaptureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransactionCaptureRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionCaptureRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionCaptureRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *TransactionCaptureRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionCaptureRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionCaptureRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetExternalIdentifier

`func (o *TransactionCaptureRequest) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *TransactionCaptureRequest) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *TransactionCaptureRequest) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *TransactionCaptureRequest) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


