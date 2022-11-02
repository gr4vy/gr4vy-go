# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;transaction&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique identifier for this transaction. | [optional] 
**Status** | Pointer to **string** | The status of the transaction. The status may change over time as asynchronous processing events occur. | [optional] 
**Intent** | Pointer to **string** | The original &#x60;intent&#x60; used when the transaction was [created](#operation/authorize-new-transaction). | [optional] 
**Amount** | Pointer to **int32** | The authorized amount for this transaction. This can be more than the actual captured amount and part of this amount may be refunded. | [optional] 
**CapturedAmount** | Pointer to **int32** | The captured amount for this transaction. This can be the total or a portion of the authorized amount. | [optional] 
**RefundedAmount** | Pointer to **int32** | The refunded amount for this transaction. This can be the total or a portion of the captured amount. | [optional] 
**Currency** | Pointer to **string** | The currency code for this transaction. | [optional] 
**Country** | Pointer to **NullableString** | The 2-letter ISO code of the country of the transaction. This is used to filter the payment services that is used to process the transaction.  | [optional] 
**PaymentMethod** | Pointer to [**PaymentMethodSnapshot**](PaymentMethod--Snapshot.md) |  | [optional] 
**Buyer** | Pointer to [**BuyerSnapshot**](Buyer--Snapshot.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this transaction was created in our system. | [optional] 
**ExternalIdentifier** | Pointer to **NullableString** | An external identifier that can be used to match the transaction against your own records. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Defines when the transaction was last updated. | [optional] 
**PaymentService** | Pointer to [**PaymentServiceSnapshot**](PaymentService--Snapshot.md) |  | [optional] 
**MerchantInitiated** | Pointer to **bool** | Indicates whether the transaction was initiated by the merchant (true) or customer (false). | [optional] [default to false]
**PaymentSource** | Pointer to **string** | The source of the transaction. Defaults to &#x60;ecommerce&#x60;. | [optional] 
**IsSubsequentPayment** | Pointer to **bool** | Indicates whether the transaction represents a subsequent payment coming from a setup recurring payment. Please note there are some restrictions on how this flag may be used.  The flag can only be &#x60;false&#x60; (or not set) when the transaction meets one of the following criteria:  * It is not &#x60;merchant_initiated&#x60;. * &#x60;payment_source&#x60; is set to &#x60;card_on_file&#x60;.  The flag can only be set to &#x60;true&#x60; when the transaction meets one of the following criteria:  * It is not &#x60;merchant_initiated&#x60;. * &#x60;payment_source&#x60; is set to &#x60;recurring&#x60; or &#x60;installment&#x60; and &#x60;merchant_initiated&#x60; is set to &#x60;true&#x60;. * &#x60;payment_source&#x60; is set to &#x60;card_on_file&#x60;. | [optional] [default to false]
**StatementDescriptor** | Pointer to [**StatementDescriptor**](StatementDescriptor.md) |  | [optional] 
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) | An array of cart items that represents the line items of a transaction. | [optional] 
**SchemeTransactionId** | Pointer to **NullableString** | An identifier for the transaction used by the scheme itself, when available.  e.g. the Visa Transaction Identifier, or Mastercard Trace ID. | [optional] [default to "null"]
**RawResponseCode** | Pointer to **NullableString** | This is the response code received from the payment service. This can be set to any value and is not standardized across different payment services. | [optional] 
**RawResponseDescription** | Pointer to **NullableString** | This is the response description received from the payment service. This can be set to any value and is not standardized across different payment services. | [optional] 
**AvsResponseCode** | Pointer to **NullableString** | The response code received from the payment service for the Address Verification Check (AVS). This code is mapped to a standardized Gr4vy AVS response code.  - &#x60;no_match&#x60; - neither address or postal code match - &#x60;match&#x60; - both address and postal code match - &#x60;partial_match_address&#x60; - address matches but postal code does not - &#x60;partial_match_postcode&#x60; - postal code matches but address does not - &#x60;unavailable &#x60; - AVS is unavailable for card/country  The value of this field can be &#x60;null&#x60; if the payment service did not provide a response. | [optional] 
**CvvResponseCode** | Pointer to **NullableString** | The response code received from the payment service for the Card Verification Value (CVV). This code is mapped to a standardized Gr4vy CVV response code.  - &#x60;no_match&#x60; - the CVV does not match the expected value - &#x60;match&#x60; - the CVV matches the expected value - &#x60;unavailable &#x60; - CVV check unavailable for card our country - &#x60;not_provided &#x60; - CVV not provided  The value of this field can be &#x60;null&#x60; if the payment service did not provide a response. | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**PaymentServiceTransactionId** | Pointer to **string** | The payment service&#39;s unique ID for the transaction. | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional information about the transaction stored as key-value pairs. | [optional] 
**ShippingDetails** | Pointer to [**ShippingAddress**](ShippingAddress.md) |  | [optional] 
**ThreeDSecure** | Pointer to [**ThreeDSecureSummary**](ThreeDSecureSummary.md) |  | [optional] 
**AuthorizedAt** | Pointer to **NullableTime** | The date and time when this transaction was authorized in the payment service.  Don&#39;t use this field to determine whether the transaction was authorized. A &#x60;null&#x60; value doesn&#39;t necessarily imply that the transaction wasn&#39;t authorized, it can mean that the payment service doesn&#39;t provide this value, that it didn&#39;t provide it at the time the transaction was authorized or that the transaction was authorized before the introduction of this field. | [optional] 
**CapturedAt** | Pointer to **NullableTime** | The date and time when this transaction was captured in the payment service.  Don&#39;t use this field to determine whether the transaction was captured. A &#x60;null&#x60; value doesn&#39;t necessarily imply that the transaction wasn&#39;t captured, it can mean that the payment service doesn&#39;t provide this value, that it didn&#39;t provide it at the time the transaction was captured or that the transaction was captured before the introduction of this field. | [optional] 
**VoidedAt** | Pointer to **NullableTime** | The date and time when this transaction was voided in the payment service.  Don&#39;t use this field to determine whether the transaction was voided. A &#x60;null&#x60; value doesn&#39;t necessarily imply that the transaction wasn&#39;t voided, it can mean that the payment service doesn&#39;t provide this value, that it didn&#39;t provide it at the time the transaction was voided or that the transaction was voided before the introduction of this field. | [optional] 

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

### GetIntent

`func (o *Transaction) GetIntent() string`

GetIntent returns the Intent field if non-nil, zero value otherwise.

### GetIntentOk

`func (o *Transaction) GetIntentOk() (*string, bool)`

GetIntentOk returns a tuple with the Intent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntent

`func (o *Transaction) SetIntent(v string)`

SetIntent sets Intent field to given value.

### HasIntent

`func (o *Transaction) HasIntent() bool`

HasIntent returns a boolean if a field has been set.

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

### GetCountry

`func (o *Transaction) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Transaction) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Transaction) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Transaction) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *Transaction) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *Transaction) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
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
### GetMethod

`func (o *Transaction) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Transaction) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Transaction) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *Transaction) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPaymentServiceTransactionId

`func (o *Transaction) GetPaymentServiceTransactionId() string`

GetPaymentServiceTransactionId returns the PaymentServiceTransactionId field if non-nil, zero value otherwise.

### GetPaymentServiceTransactionIdOk

`func (o *Transaction) GetPaymentServiceTransactionIdOk() (*string, bool)`

GetPaymentServiceTransactionIdOk returns a tuple with the PaymentServiceTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceTransactionId

`func (o *Transaction) SetPaymentServiceTransactionId(v string)`

SetPaymentServiceTransactionId sets PaymentServiceTransactionId field to given value.

### HasPaymentServiceTransactionId

`func (o *Transaction) HasPaymentServiceTransactionId() bool`

HasPaymentServiceTransactionId returns a boolean if a field has been set.

### GetMetadata

`func (o *Transaction) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Transaction) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Transaction) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Transaction) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetShippingDetails

`func (o *Transaction) GetShippingDetails() ShippingAddress`

GetShippingDetails returns the ShippingDetails field if non-nil, zero value otherwise.

### GetShippingDetailsOk

`func (o *Transaction) GetShippingDetailsOk() (*ShippingAddress, bool)`

GetShippingDetailsOk returns a tuple with the ShippingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingDetails

`func (o *Transaction) SetShippingDetails(v ShippingAddress)`

SetShippingDetails sets ShippingDetails field to given value.

### HasShippingDetails

`func (o *Transaction) HasShippingDetails() bool`

HasShippingDetails returns a boolean if a field has been set.

### GetThreeDSecure

`func (o *Transaction) GetThreeDSecure() ThreeDSecureSummary`

GetThreeDSecure returns the ThreeDSecure field if non-nil, zero value otherwise.

### GetThreeDSecureOk

`func (o *Transaction) GetThreeDSecureOk() (*ThreeDSecureSummary, bool)`

GetThreeDSecureOk returns a tuple with the ThreeDSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSecure

`func (o *Transaction) SetThreeDSecure(v ThreeDSecureSummary)`

SetThreeDSecure sets ThreeDSecure field to given value.

### HasThreeDSecure

`func (o *Transaction) HasThreeDSecure() bool`

HasThreeDSecure returns a boolean if a field has been set.

### GetAuthorizedAt

`func (o *Transaction) GetAuthorizedAt() time.Time`

GetAuthorizedAt returns the AuthorizedAt field if non-nil, zero value otherwise.

### GetAuthorizedAtOk

`func (o *Transaction) GetAuthorizedAtOk() (*time.Time, bool)`

GetAuthorizedAtOk returns a tuple with the AuthorizedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedAt

`func (o *Transaction) SetAuthorizedAt(v time.Time)`

SetAuthorizedAt sets AuthorizedAt field to given value.

### HasAuthorizedAt

`func (o *Transaction) HasAuthorizedAt() bool`

HasAuthorizedAt returns a boolean if a field has been set.

### SetAuthorizedAtNil

`func (o *Transaction) SetAuthorizedAtNil(b bool)`

 SetAuthorizedAtNil sets the value for AuthorizedAt to be an explicit nil

### UnsetAuthorizedAt
`func (o *Transaction) UnsetAuthorizedAt()`

UnsetAuthorizedAt ensures that no value is present for AuthorizedAt, not even an explicit nil
### GetCapturedAt

`func (o *Transaction) GetCapturedAt() time.Time`

GetCapturedAt returns the CapturedAt field if non-nil, zero value otherwise.

### GetCapturedAtOk

`func (o *Transaction) GetCapturedAtOk() (*time.Time, bool)`

GetCapturedAtOk returns a tuple with the CapturedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedAt

`func (o *Transaction) SetCapturedAt(v time.Time)`

SetCapturedAt sets CapturedAt field to given value.

### HasCapturedAt

`func (o *Transaction) HasCapturedAt() bool`

HasCapturedAt returns a boolean if a field has been set.

### SetCapturedAtNil

`func (o *Transaction) SetCapturedAtNil(b bool)`

 SetCapturedAtNil sets the value for CapturedAt to be an explicit nil

### UnsetCapturedAt
`func (o *Transaction) UnsetCapturedAt()`

UnsetCapturedAt ensures that no value is present for CapturedAt, not even an explicit nil
### GetVoidedAt

`func (o *Transaction) GetVoidedAt() time.Time`

GetVoidedAt returns the VoidedAt field if non-nil, zero value otherwise.

### GetVoidedAtOk

`func (o *Transaction) GetVoidedAtOk() (*time.Time, bool)`

GetVoidedAtOk returns a tuple with the VoidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidedAt

`func (o *Transaction) SetVoidedAt(v time.Time)`

SetVoidedAt sets VoidedAt field to given value.

### HasVoidedAt

`func (o *Transaction) HasVoidedAt() bool`

HasVoidedAt returns a boolean if a field has been set.

### SetVoidedAtNil

`func (o *Transaction) SetVoidedAtNil(b bool)`

 SetVoidedAtNil sets the value for VoidedAt to be an explicit nil

### UnsetVoidedAt
`func (o *Transaction) UnsetVoidedAt()`

UnsetVoidedAt ensures that no value is present for VoidedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


