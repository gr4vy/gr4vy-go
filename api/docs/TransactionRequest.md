# TransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The monetary amount to create an authorization for, in the smallest currency unit for the given currency, for example &#x60;1299&#x60; cents to create an authorization for &#x60;$12.99&#x60;. | 
**Currency** | **string** | A supported ISO-4217 currency code. | 
**PaymentMethod** | [**OneOfCardRequestRedirectRequestTokenizedRequest**](oneOf&lt;CardRequest,RedirectRequest,TokenizedRequest&gt;.md) | The optional payment method details to create an authorization for. This field is required for processing a card. | 
**Store** | Pointer to **bool** | Whether or not to also try and store the payment method with us so that it can be used again for future use. This is only supported for payment methods that support this feature. | [optional] [default to false]
**Intent** | Pointer to **string** | Defines the intent of this API call. This determines the desired initial state of the transaction.  * &#x60;authorize&#x60; - (Default) Optionally approves and then authorizes a transaction but does not capture the funds. * &#x60;capture&#x60; - Optionally approves and then authorizes and captures the funds of the transaction. | [optional] 
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the transaction against your own records. | [optional] 
**Environment** | Pointer to **string** | Defines the environment to create this transaction in. Setting this to anything other than &#x60;production&#x60; will force Gr4vy to use the payment a service configured for that environment. | [optional] 
**ThreeDSecureData** | Pointer to [**OneOfThreeDSecureDataV1ThreeDSecureDataV2**](oneOf&lt;ThreeDSecureDataV1,ThreeDSecureDataV2&gt;.md) | Pass through 3-D Secure data to support external 3-D Secure authorisation. If using an external 3-D Secure provider, you should not pass a &#x60;redirect_url&#x60; in the &#x60;payment_method&#x60; object for a transaction. | [optional] 

## Methods

### NewTransactionRequest

`func NewTransactionRequest(amount int32, currency string, paymentMethod OneOfCardRequestRedirectRequestTokenizedRequest, ) *TransactionRequest`

NewTransactionRequest instantiates a new TransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionRequestWithDefaults

`func NewTransactionRequestWithDefaults() *TransactionRequest`

NewTransactionRequestWithDefaults instantiates a new TransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransactionRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *TransactionRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPaymentMethod

`func (o *TransactionRequest) GetPaymentMethod() OneOfCardRequestRedirectRequestTokenizedRequest`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *TransactionRequest) GetPaymentMethodOk() (*OneOfCardRequestRedirectRequestTokenizedRequest, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *TransactionRequest) SetPaymentMethod(v OneOfCardRequestRedirectRequestTokenizedRequest)`

SetPaymentMethod sets PaymentMethod field to given value.


### SetPaymentMethodNil

`func (o *TransactionRequest) SetPaymentMethodNil(b bool)`

 SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil

### UnsetPaymentMethod
`func (o *TransactionRequest) UnsetPaymentMethod()`

UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil
### GetStore

`func (o *TransactionRequest) GetStore() bool`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *TransactionRequest) GetStoreOk() (*bool, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *TransactionRequest) SetStore(v bool)`

SetStore sets Store field to given value.

### HasStore

`func (o *TransactionRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetIntent

`func (o *TransactionRequest) GetIntent() string`

GetIntent returns the Intent field if non-nil, zero value otherwise.

### GetIntentOk

`func (o *TransactionRequest) GetIntentOk() (*string, bool)`

GetIntentOk returns a tuple with the Intent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntent

`func (o *TransactionRequest) SetIntent(v string)`

SetIntent sets Intent field to given value.

### HasIntent

`func (o *TransactionRequest) HasIntent() bool`

HasIntent returns a boolean if a field has been set.

### GetExternalIdentifier

`func (o *TransactionRequest) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *TransactionRequest) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *TransactionRequest) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *TransactionRequest) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *TransactionRequest) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *TransactionRequest) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetEnvironment

`func (o *TransactionRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *TransactionRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *TransactionRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *TransactionRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetThreeDSecureData

`func (o *TransactionRequest) GetThreeDSecureData() OneOfThreeDSecureDataV1ThreeDSecureDataV2`

GetThreeDSecureData returns the ThreeDSecureData field if non-nil, zero value otherwise.

### GetThreeDSecureDataOk

`func (o *TransactionRequest) GetThreeDSecureDataOk() (*OneOfThreeDSecureDataV1ThreeDSecureDataV2, bool)`

GetThreeDSecureDataOk returns a tuple with the ThreeDSecureData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSecureData

`func (o *TransactionRequest) SetThreeDSecureData(v OneOfThreeDSecureDataV1ThreeDSecureDataV2)`

SetThreeDSecureData sets ThreeDSecureData field to given value.

### HasThreeDSecureData

`func (o *TransactionRequest) HasThreeDSecureData() bool`

HasThreeDSecureData returns a boolean if a field has been set.

### SetThreeDSecureDataNil

`func (o *TransactionRequest) SetThreeDSecureDataNil(b bool)`

 SetThreeDSecureDataNil sets the value for ThreeDSecureData to be an explicit nil

### UnsetThreeDSecureData
`func (o *TransactionRequest) UnsetThreeDSecureData()`

UnsetThreeDSecureData ensures that no value is present for ThreeDSecureData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


