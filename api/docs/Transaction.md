# Transaction

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
**MerchantInitiated** | Pointer to **bool** | Indicates whether the transaction was initiated by the merchant (true) or customer (false). | [optional] [default to false]
**PaymentSource** | Pointer to **string** | The source of the transaction. Defaults to &#x60;ecommerce&#x60;. | [optional] 
**IsSubsequentPayment** | Pointer to **bool** | Indicates whether the transaction represents a subsequent payment coming from a setup recurring payment. Please note this flag is only compatible with &#x60;payment_source&#x60; set to &#x60;recurring&#x60;, &#x60;installment&#x60;, or &#x60;card_on_file&#x60; and will be ignored for other values or if &#x60;payment_source&#x60; is not present. | [optional] [default to false]
**StatementDescriptor** | Pointer to [**StatementDescriptor**](StatementDescriptor.md) |  | [optional] 
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) | An array of cart items that represents the line items of a transaction. | [optional] 
**SchemeTransactionId** | Pointer to **NullableString** | An identifier for the transaction used by the scheme itself, when available.  e.g. the Visa Transaction Identifier, or Mastercard Trace ID. | [optional] [default to "null"]
**RawResponseCode** | Pointer to **NullableString** | This is the response code received from the payment service. This can be set to any value and is not standardized across different payment services. | [optional] 
**RawResponseDescription** | Pointer to **NullableString** | This is the response description received from the payment service. This can be set to any value and is not standardized across different payment services. | [optional] 
**AvsResponseCode** | Pointer to **NullableString** | The response code received from the payment service for the Address Verification Check (AVS). This code is mapped to a standardized Gr4vy AVS response code.  - &#x60;no_match&#x60; - neither address or postal code match - &#x60;match&#x60; - both address and postal code match - &#x60;partial_match_address&#x60; - address matches but postal code does not - &#x60;partial_match_postcode&#x60; - postal code matches but address does not - &#x60;unavailable &#x60; - AVS is unavailable for card/country  The value of this field can be &#x60;null&#x60; if the payment service did not provide a response. | [optional] 
**CvvResponseCode** | Pointer to **NullableString** | The response code received from the payment service for the Card Verification Value (CVV). This code is mapped to a standardized Gr4vy CVV response code.  - &#x60;no_match&#x60; - the CVV does not match the expected value - &#x60;match&#x60; - the CVV matches the expected value - &#x60;unavailable &#x60; - CVV check unavailable for card our country - &#x60;not_provided &#x60; - CVV not provided  The value of this field can be &#x60;null&#x60; if the payment service did not provide a response. | [optional] 

## Methods

### NewTransaction

`func NewTransaction() *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Transaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Transaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *Transaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transaction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Transaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *Transaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transaction) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Transaction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAmount

`func (o *Transaction) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Transaction) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Transaction) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Transaction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCapturedAmount

`func (o *Transaction) GetCapturedAmount() int32`

GetCapturedAmount returns the CapturedAmount field if non-nil, zero value otherwise.

### GetCapturedAmountOk

`func (o *Transaction) GetCapturedAmountOk() (*int32, bool)`

GetCapturedAmountOk returns a tuple with the CapturedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedAmount

`func (o *Transaction) SetCapturedAmount(v int32)`

SetCapturedAmount sets CapturedAmount field to given value.

### HasCapturedAmount

`func (o *Transaction) HasCapturedAmount() bool`

HasCapturedAmount returns a boolean if a field has been set.

### GetRefundedAmount

`func (o *Transaction) GetRefundedAmount() int32`

GetRefundedAmount returns the RefundedAmount field if non-nil, zero value otherwise.

### GetRefundedAmountOk

`func (o *Transaction) GetRefundedAmountOk() (*int32, bool)`

GetRefundedAmountOk returns a tuple with the RefundedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAmount

`func (o *Transaction) SetRefundedAmount(v int32)`

SetRefundedAmount sets RefundedAmount field to given value.

### HasRefundedAmount

`func (o *Transaction) HasRefundedAmount() bool`

HasRefundedAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *Transaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Transaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Transaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Transaction) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *Transaction) GetPaymentMethod() PaymentMethodSnapshot`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *Transaction) GetPaymentMethodOk() (*PaymentMethodSnapshot, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *Transaction) SetPaymentMethod(v PaymentMethodSnapshot)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *Transaction) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetBuyer

`func (o *Transaction) GetBuyer() BuyerSnapshot`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *Transaction) GetBuyerOk() (*BuyerSnapshot, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *Transaction) SetBuyer(v BuyerSnapshot)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *Transaction) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Transaction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Transaction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Transaction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Transaction) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExternalIdentifier

`func (o *Transaction) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *Transaction) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *Transaction) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *Transaction) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *Transaction) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *Transaction) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetUpdatedAt

`func (o *Transaction) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Transaction) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Transaction) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Transaction) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPaymentService

`func (o *Transaction) GetPaymentService() PaymentServiceSnapshot`

GetPaymentService returns the PaymentService field if non-nil, zero value otherwise.

### GetPaymentServiceOk

`func (o *Transaction) GetPaymentServiceOk() (*PaymentServiceSnapshot, bool)`

GetPaymentServiceOk returns a tuple with the PaymentService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentService

`func (o *Transaction) SetPaymentService(v PaymentServiceSnapshot)`

SetPaymentService sets PaymentService field to given value.

### HasPaymentService

`func (o *Transaction) HasPaymentService() bool`

HasPaymentService returns a boolean if a field has been set.

### GetMerchantInitiated

`func (o *Transaction) GetMerchantInitiated() bool`

GetMerchantInitiated returns the MerchantInitiated field if non-nil, zero value otherwise.

### GetMerchantInitiatedOk

`func (o *Transaction) GetMerchantInitiatedOk() (*bool, bool)`

GetMerchantInitiatedOk returns a tuple with the MerchantInitiated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantInitiated

`func (o *Transaction) SetMerchantInitiated(v bool)`

SetMerchantInitiated sets MerchantInitiated field to given value.

### HasMerchantInitiated

`func (o *Transaction) HasMerchantInitiated() bool`

HasMerchantInitiated returns a boolean if a field has been set.

### GetPaymentSource

`func (o *Transaction) GetPaymentSource() string`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *Transaction) GetPaymentSourceOk() (*string, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *Transaction) SetPaymentSource(v string)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *Transaction) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.

### GetIsSubsequentPayment

`func (o *Transaction) GetIsSubsequentPayment() bool`

GetIsSubsequentPayment returns the IsSubsequentPayment field if non-nil, zero value otherwise.

### GetIsSubsequentPaymentOk

`func (o *Transaction) GetIsSubsequentPaymentOk() (*bool, bool)`

GetIsSubsequentPaymentOk returns a tuple with the IsSubsequentPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubsequentPayment

`func (o *Transaction) SetIsSubsequentPayment(v bool)`

SetIsSubsequentPayment sets IsSubsequentPayment field to given value.

### HasIsSubsequentPayment

`func (o *Transaction) HasIsSubsequentPayment() bool`

HasIsSubsequentPayment returns a boolean if a field has been set.

### GetStatementDescriptor

`func (o *Transaction) GetStatementDescriptor() StatementDescriptor`

GetStatementDescriptor returns the StatementDescriptor field if non-nil, zero value otherwise.

### GetStatementDescriptorOk

`func (o *Transaction) GetStatementDescriptorOk() (*StatementDescriptor, bool)`

GetStatementDescriptorOk returns a tuple with the StatementDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescriptor

`func (o *Transaction) SetStatementDescriptor(v StatementDescriptor)`

SetStatementDescriptor sets StatementDescriptor field to given value.

### HasStatementDescriptor

`func (o *Transaction) HasStatementDescriptor() bool`

HasStatementDescriptor returns a boolean if a field has been set.

### GetCartItems

`func (o *Transaction) GetCartItems() []CartItem`

GetCartItems returns the CartItems field if non-nil, zero value otherwise.

### GetCartItemsOk

`func (o *Transaction) GetCartItemsOk() (*[]CartItem, bool)`

GetCartItemsOk returns a tuple with the CartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItems

`func (o *Transaction) SetCartItems(v []CartItem)`

SetCartItems sets CartItems field to given value.

### HasCartItems

`func (o *Transaction) HasCartItems() bool`

HasCartItems returns a boolean if a field has been set.

### GetSchemeTransactionId

`func (o *Transaction) GetSchemeTransactionId() string`

GetSchemeTransactionId returns the SchemeTransactionId field if non-nil, zero value otherwise.

### GetSchemeTransactionIdOk

`func (o *Transaction) GetSchemeTransactionIdOk() (*string, bool)`

GetSchemeTransactionIdOk returns a tuple with the SchemeTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeTransactionId

`func (o *Transaction) SetSchemeTransactionId(v string)`

SetSchemeTransactionId sets SchemeTransactionId field to given value.

### HasSchemeTransactionId

`func (o *Transaction) HasSchemeTransactionId() bool`

HasSchemeTransactionId returns a boolean if a field has been set.

### SetSchemeTransactionIdNil

`func (o *Transaction) SetSchemeTransactionIdNil(b bool)`

 SetSchemeTransactionIdNil sets the value for SchemeTransactionId to be an explicit nil

### UnsetSchemeTransactionId
`func (o *Transaction) UnsetSchemeTransactionId()`

UnsetSchemeTransactionId ensures that no value is present for SchemeTransactionId, not even an explicit nil
### GetRawResponseCode

`func (o *Transaction) GetRawResponseCode() string`

GetRawResponseCode returns the RawResponseCode field if non-nil, zero value otherwise.

### GetRawResponseCodeOk

`func (o *Transaction) GetRawResponseCodeOk() (*string, bool)`

GetRawResponseCodeOk returns a tuple with the RawResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawResponseCode

`func (o *Transaction) SetRawResponseCode(v string)`

SetRawResponseCode sets RawResponseCode field to given value.

### HasRawResponseCode

`func (o *Transaction) HasRawResponseCode() bool`

HasRawResponseCode returns a boolean if a field has been set.

### SetRawResponseCodeNil

`func (o *Transaction) SetRawResponseCodeNil(b bool)`

 SetRawResponseCodeNil sets the value for RawResponseCode to be an explicit nil

### UnsetRawResponseCode
`func (o *Transaction) UnsetRawResponseCode()`

UnsetRawResponseCode ensures that no value is present for RawResponseCode, not even an explicit nil
### GetRawResponseDescription

`func (o *Transaction) GetRawResponseDescription() string`

GetRawResponseDescription returns the RawResponseDescription field if non-nil, zero value otherwise.

### GetRawResponseDescriptionOk

`func (o *Transaction) GetRawResponseDescriptionOk() (*string, bool)`

GetRawResponseDescriptionOk returns a tuple with the RawResponseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawResponseDescription

`func (o *Transaction) SetRawResponseDescription(v string)`

SetRawResponseDescription sets RawResponseDescription field to given value.

### HasRawResponseDescription

`func (o *Transaction) HasRawResponseDescription() bool`

HasRawResponseDescription returns a boolean if a field has been set.

### SetRawResponseDescriptionNil

`func (o *Transaction) SetRawResponseDescriptionNil(b bool)`

 SetRawResponseDescriptionNil sets the value for RawResponseDescription to be an explicit nil

### UnsetRawResponseDescription
`func (o *Transaction) UnsetRawResponseDescription()`

UnsetRawResponseDescription ensures that no value is present for RawResponseDescription, not even an explicit nil
### GetAvsResponseCode

`func (o *Transaction) GetAvsResponseCode() string`

GetAvsResponseCode returns the AvsResponseCode field if non-nil, zero value otherwise.

### GetAvsResponseCodeOk

`func (o *Transaction) GetAvsResponseCodeOk() (*string, bool)`

GetAvsResponseCodeOk returns a tuple with the AvsResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsResponseCode

`func (o *Transaction) SetAvsResponseCode(v string)`

SetAvsResponseCode sets AvsResponseCode field to given value.

### HasAvsResponseCode

`func (o *Transaction) HasAvsResponseCode() bool`

HasAvsResponseCode returns a boolean if a field has been set.

### SetAvsResponseCodeNil

`func (o *Transaction) SetAvsResponseCodeNil(b bool)`

 SetAvsResponseCodeNil sets the value for AvsResponseCode to be an explicit nil

### UnsetAvsResponseCode
`func (o *Transaction) UnsetAvsResponseCode()`

UnsetAvsResponseCode ensures that no value is present for AvsResponseCode, not even an explicit nil
### GetCvvResponseCode

`func (o *Transaction) GetCvvResponseCode() string`

GetCvvResponseCode returns the CvvResponseCode field if non-nil, zero value otherwise.

### GetCvvResponseCodeOk

`func (o *Transaction) GetCvvResponseCodeOk() (*string, bool)`

GetCvvResponseCodeOk returns a tuple with the CvvResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvResponseCode

`func (o *Transaction) SetCvvResponseCode(v string)`

SetCvvResponseCode sets CvvResponseCode field to given value.

### HasCvvResponseCode

`func (o *Transaction) HasCvvResponseCode() bool`

HasCvvResponseCode returns a boolean if a field has been set.

### SetCvvResponseCodeNil

`func (o *Transaction) SetCvvResponseCodeNil(b bool)`

 SetCvvResponseCodeNil sets the value for CvvResponseCode to be an explicit nil

### UnsetCvvResponseCode
`func (o *Transaction) UnsetCvvResponseCode()`

UnsetCvvResponseCode ensures that no value is present for CvvResponseCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


