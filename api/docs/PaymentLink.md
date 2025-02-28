# PaymentLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of a payment link. | [optional] 
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;payment_link&#x60;. | [optional] 
**Amount** | Pointer to **int32** | The monetary amount for this payment link, in the smallest currency unit for the given currency, for example &#x60;1299&#x60; cents to create an authorization for &#x60;$12.99&#x60;. | [optional] 
**Currency** | Pointer to **string** | A supported ISO-4217 currency code. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this payment link was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when this payment link was created. | [optional] 
**ExpiresAt** | Pointer to **time.Time** | The date and time when this payment link expires. | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ExternalIdentifier** | Pointer to **NullableString** | A value that can be used to match the payment link against your own records. | [optional] 
**StatementDescriptor** | Pointer to [**NullableStatementDescriptor**](StatementDescriptor.md) |  | [optional] 
**Locale** | Pointer to **NullableString** | The locale used to translate text within the payment link. | [optional] 
**MerchantName** | Pointer to **NullableString** | The name of the merchant to display on the payment link. | [optional] 
**MerchantUrl** | Pointer to **NullableString** | The URL of the merchant to display on the payment link. | [optional] 
**MerchantBannerUrl** | Pointer to **NullableString** | The URL of the merchant banner to display on the payment link. | [optional] 
**MerchantColor** | Pointer to **NullableString** | The color code of the merchant to display on the payment link. | [optional] 
**MerchantMessage** | Pointer to **NullableString** | The message to display on the payment link. | [optional] 
**MerchantTermsAndConditionsUrl** | Pointer to **NullableString** | The URL of the merchant terms and conditions to display on the payment link. | [optional] 
**MerchantFaviconUrl** | Pointer to **NullableString** | The URL of the merchant favicon icon. | [optional] 
**Country** | Pointer to **NullableString** | The 2-letter ISO code of the country of the transaction. This is used to filter the payment services that is used to process the transaction.  | [optional] 
**Intent** | Pointer to **NullableString** | The intent of the payment link. | [optional] 
**ReturnUrl** | Pointer to **NullableString** | The URL to redirect the buyer to after payment. | [optional] 
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) | An array of cart items that represents the line items of a payment link. | [optional] 
**Metadata** | Pointer to **map[string]string** | Any additional information about the payment link that you would like to store as key-value pairs. This data is passed to payment service providers that support it. | [optional] 
**PaymentSource** | Pointer to **NullableString** | The source of the payment link. Defaults to &#x60;ecommerce&#x60;. | [optional] 
**Buyer** | Pointer to [**NullableBuyerSnapshot**](Buyer--Snapshot.md) | The buyer used for this transaction. | [optional] 
**ShippingDetails** | Pointer to [**NullableShippingDetail**](ShippingDetail.md) | Shipping details for the payment link. | [optional] 

## Methods

### NewPaymentLink

`func NewPaymentLink() *PaymentLink`

NewPaymentLink instantiates a new PaymentLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentLinkWithDefaults

`func NewPaymentLinkWithDefaults() *PaymentLink`

NewPaymentLinkWithDefaults instantiates a new PaymentLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentLink) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentLink) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentLink) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentLink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PaymentLink) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentLink) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentLink) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentLink) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentLink) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentLink) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentLink) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentLink) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentLink) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentLink) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentLink) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentLink) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaymentLink) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentLink) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentLink) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaymentLink) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PaymentLink) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentLink) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentLink) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PaymentLink) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PaymentLink) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PaymentLink) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PaymentLink) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PaymentLink) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentLink) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentLink) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentLink) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentLink) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExternalIdentifier

`func (o *PaymentLink) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *PaymentLink) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *PaymentLink) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *PaymentLink) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *PaymentLink) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *PaymentLink) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetStatementDescriptor

`func (o *PaymentLink) GetStatementDescriptor() StatementDescriptor`

GetStatementDescriptor returns the StatementDescriptor field if non-nil, zero value otherwise.

### GetStatementDescriptorOk

`func (o *PaymentLink) GetStatementDescriptorOk() (*StatementDescriptor, bool)`

GetStatementDescriptorOk returns a tuple with the StatementDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescriptor

`func (o *PaymentLink) SetStatementDescriptor(v StatementDescriptor)`

SetStatementDescriptor sets StatementDescriptor field to given value.

### HasStatementDescriptor

`func (o *PaymentLink) HasStatementDescriptor() bool`

HasStatementDescriptor returns a boolean if a field has been set.

### SetStatementDescriptorNil

`func (o *PaymentLink) SetStatementDescriptorNil(b bool)`

 SetStatementDescriptorNil sets the value for StatementDescriptor to be an explicit nil

### UnsetStatementDescriptor
`func (o *PaymentLink) UnsetStatementDescriptor()`

UnsetStatementDescriptor ensures that no value is present for StatementDescriptor, not even an explicit nil
### GetLocale

`func (o *PaymentLink) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *PaymentLink) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *PaymentLink) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *PaymentLink) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *PaymentLink) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *PaymentLink) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetMerchantName

`func (o *PaymentLink) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *PaymentLink) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *PaymentLink) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *PaymentLink) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### SetMerchantNameNil

`func (o *PaymentLink) SetMerchantNameNil(b bool)`

 SetMerchantNameNil sets the value for MerchantName to be an explicit nil

### UnsetMerchantName
`func (o *PaymentLink) UnsetMerchantName()`

UnsetMerchantName ensures that no value is present for MerchantName, not even an explicit nil
### GetMerchantUrl

`func (o *PaymentLink) GetMerchantUrl() string`

GetMerchantUrl returns the MerchantUrl field if non-nil, zero value otherwise.

### GetMerchantUrlOk

`func (o *PaymentLink) GetMerchantUrlOk() (*string, bool)`

GetMerchantUrlOk returns a tuple with the MerchantUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantUrl

`func (o *PaymentLink) SetMerchantUrl(v string)`

SetMerchantUrl sets MerchantUrl field to given value.

### HasMerchantUrl

`func (o *PaymentLink) HasMerchantUrl() bool`

HasMerchantUrl returns a boolean if a field has been set.

### SetMerchantUrlNil

`func (o *PaymentLink) SetMerchantUrlNil(b bool)`

 SetMerchantUrlNil sets the value for MerchantUrl to be an explicit nil

### UnsetMerchantUrl
`func (o *PaymentLink) UnsetMerchantUrl()`

UnsetMerchantUrl ensures that no value is present for MerchantUrl, not even an explicit nil
### GetMerchantBannerUrl

`func (o *PaymentLink) GetMerchantBannerUrl() string`

GetMerchantBannerUrl returns the MerchantBannerUrl field if non-nil, zero value otherwise.

### GetMerchantBannerUrlOk

`func (o *PaymentLink) GetMerchantBannerUrlOk() (*string, bool)`

GetMerchantBannerUrlOk returns a tuple with the MerchantBannerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantBannerUrl

`func (o *PaymentLink) SetMerchantBannerUrl(v string)`

SetMerchantBannerUrl sets MerchantBannerUrl field to given value.

### HasMerchantBannerUrl

`func (o *PaymentLink) HasMerchantBannerUrl() bool`

HasMerchantBannerUrl returns a boolean if a field has been set.

### SetMerchantBannerUrlNil

`func (o *PaymentLink) SetMerchantBannerUrlNil(b bool)`

 SetMerchantBannerUrlNil sets the value for MerchantBannerUrl to be an explicit nil

### UnsetMerchantBannerUrl
`func (o *PaymentLink) UnsetMerchantBannerUrl()`

UnsetMerchantBannerUrl ensures that no value is present for MerchantBannerUrl, not even an explicit nil
### GetMerchantColor

`func (o *PaymentLink) GetMerchantColor() string`

GetMerchantColor returns the MerchantColor field if non-nil, zero value otherwise.

### GetMerchantColorOk

`func (o *PaymentLink) GetMerchantColorOk() (*string, bool)`

GetMerchantColorOk returns a tuple with the MerchantColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantColor

`func (o *PaymentLink) SetMerchantColor(v string)`

SetMerchantColor sets MerchantColor field to given value.

### HasMerchantColor

`func (o *PaymentLink) HasMerchantColor() bool`

HasMerchantColor returns a boolean if a field has been set.

### SetMerchantColorNil

`func (o *PaymentLink) SetMerchantColorNil(b bool)`

 SetMerchantColorNil sets the value for MerchantColor to be an explicit nil

### UnsetMerchantColor
`func (o *PaymentLink) UnsetMerchantColor()`

UnsetMerchantColor ensures that no value is present for MerchantColor, not even an explicit nil
### GetMerchantMessage

`func (o *PaymentLink) GetMerchantMessage() string`

GetMerchantMessage returns the MerchantMessage field if non-nil, zero value otherwise.

### GetMerchantMessageOk

`func (o *PaymentLink) GetMerchantMessageOk() (*string, bool)`

GetMerchantMessageOk returns a tuple with the MerchantMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMessage

`func (o *PaymentLink) SetMerchantMessage(v string)`

SetMerchantMessage sets MerchantMessage field to given value.

### HasMerchantMessage

`func (o *PaymentLink) HasMerchantMessage() bool`

HasMerchantMessage returns a boolean if a field has been set.

### SetMerchantMessageNil

`func (o *PaymentLink) SetMerchantMessageNil(b bool)`

 SetMerchantMessageNil sets the value for MerchantMessage to be an explicit nil

### UnsetMerchantMessage
`func (o *PaymentLink) UnsetMerchantMessage()`

UnsetMerchantMessage ensures that no value is present for MerchantMessage, not even an explicit nil
### GetMerchantTermsAndConditionsUrl

`func (o *PaymentLink) GetMerchantTermsAndConditionsUrl() string`

GetMerchantTermsAndConditionsUrl returns the MerchantTermsAndConditionsUrl field if non-nil, zero value otherwise.

### GetMerchantTermsAndConditionsUrlOk

`func (o *PaymentLink) GetMerchantTermsAndConditionsUrlOk() (*string, bool)`

GetMerchantTermsAndConditionsUrlOk returns a tuple with the MerchantTermsAndConditionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTermsAndConditionsUrl

`func (o *PaymentLink) SetMerchantTermsAndConditionsUrl(v string)`

SetMerchantTermsAndConditionsUrl sets MerchantTermsAndConditionsUrl field to given value.

### HasMerchantTermsAndConditionsUrl

`func (o *PaymentLink) HasMerchantTermsAndConditionsUrl() bool`

HasMerchantTermsAndConditionsUrl returns a boolean if a field has been set.

### SetMerchantTermsAndConditionsUrlNil

`func (o *PaymentLink) SetMerchantTermsAndConditionsUrlNil(b bool)`

 SetMerchantTermsAndConditionsUrlNil sets the value for MerchantTermsAndConditionsUrl to be an explicit nil

### UnsetMerchantTermsAndConditionsUrl
`func (o *PaymentLink) UnsetMerchantTermsAndConditionsUrl()`

UnsetMerchantTermsAndConditionsUrl ensures that no value is present for MerchantTermsAndConditionsUrl, not even an explicit nil
### GetMerchantFaviconUrl

`func (o *PaymentLink) GetMerchantFaviconUrl() string`

GetMerchantFaviconUrl returns the MerchantFaviconUrl field if non-nil, zero value otherwise.

### GetMerchantFaviconUrlOk

`func (o *PaymentLink) GetMerchantFaviconUrlOk() (*string, bool)`

GetMerchantFaviconUrlOk returns a tuple with the MerchantFaviconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantFaviconUrl

`func (o *PaymentLink) SetMerchantFaviconUrl(v string)`

SetMerchantFaviconUrl sets MerchantFaviconUrl field to given value.

### HasMerchantFaviconUrl

`func (o *PaymentLink) HasMerchantFaviconUrl() bool`

HasMerchantFaviconUrl returns a boolean if a field has been set.

### SetMerchantFaviconUrlNil

`func (o *PaymentLink) SetMerchantFaviconUrlNil(b bool)`

 SetMerchantFaviconUrlNil sets the value for MerchantFaviconUrl to be an explicit nil

### UnsetMerchantFaviconUrl
`func (o *PaymentLink) UnsetMerchantFaviconUrl()`

UnsetMerchantFaviconUrl ensures that no value is present for MerchantFaviconUrl, not even an explicit nil
### GetCountry

`func (o *PaymentLink) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaymentLink) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaymentLink) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PaymentLink) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *PaymentLink) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *PaymentLink) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetIntent

`func (o *PaymentLink) GetIntent() string`

GetIntent returns the Intent field if non-nil, zero value otherwise.

### GetIntentOk

`func (o *PaymentLink) GetIntentOk() (*string, bool)`

GetIntentOk returns a tuple with the Intent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntent

`func (o *PaymentLink) SetIntent(v string)`

SetIntent sets Intent field to given value.

### HasIntent

`func (o *PaymentLink) HasIntent() bool`

HasIntent returns a boolean if a field has been set.

### SetIntentNil

`func (o *PaymentLink) SetIntentNil(b bool)`

 SetIntentNil sets the value for Intent to be an explicit nil

### UnsetIntent
`func (o *PaymentLink) UnsetIntent()`

UnsetIntent ensures that no value is present for Intent, not even an explicit nil
### GetReturnUrl

`func (o *PaymentLink) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *PaymentLink) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *PaymentLink) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *PaymentLink) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### SetReturnUrlNil

`func (o *PaymentLink) SetReturnUrlNil(b bool)`

 SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil

### UnsetReturnUrl
`func (o *PaymentLink) UnsetReturnUrl()`

UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
### GetCartItems

`func (o *PaymentLink) GetCartItems() []CartItem`

GetCartItems returns the CartItems field if non-nil, zero value otherwise.

### GetCartItemsOk

`func (o *PaymentLink) GetCartItemsOk() (*[]CartItem, bool)`

GetCartItemsOk returns a tuple with the CartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItems

`func (o *PaymentLink) SetCartItems(v []CartItem)`

SetCartItems sets CartItems field to given value.

### HasCartItems

`func (o *PaymentLink) HasCartItems() bool`

HasCartItems returns a boolean if a field has been set.

### GetMetadata

`func (o *PaymentLink) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentLink) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentLink) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaymentLink) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PaymentLink) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PaymentLink) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetPaymentSource

`func (o *PaymentLink) GetPaymentSource() string`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *PaymentLink) GetPaymentSourceOk() (*string, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *PaymentLink) SetPaymentSource(v string)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *PaymentLink) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.

### SetPaymentSourceNil

`func (o *PaymentLink) SetPaymentSourceNil(b bool)`

 SetPaymentSourceNil sets the value for PaymentSource to be an explicit nil

### UnsetPaymentSource
`func (o *PaymentLink) UnsetPaymentSource()`

UnsetPaymentSource ensures that no value is present for PaymentSource, not even an explicit nil
### GetBuyer

`func (o *PaymentLink) GetBuyer() BuyerSnapshot`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *PaymentLink) GetBuyerOk() (*BuyerSnapshot, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *PaymentLink) SetBuyer(v BuyerSnapshot)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *PaymentLink) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### SetBuyerNil

`func (o *PaymentLink) SetBuyerNil(b bool)`

 SetBuyerNil sets the value for Buyer to be an explicit nil

### UnsetBuyer
`func (o *PaymentLink) UnsetBuyer()`

UnsetBuyer ensures that no value is present for Buyer, not even an explicit nil
### GetShippingDetails

`func (o *PaymentLink) GetShippingDetails() ShippingDetail`

GetShippingDetails returns the ShippingDetails field if non-nil, zero value otherwise.

### GetShippingDetailsOk

`func (o *PaymentLink) GetShippingDetailsOk() (*ShippingDetail, bool)`

GetShippingDetailsOk returns a tuple with the ShippingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingDetails

`func (o *PaymentLink) SetShippingDetails(v ShippingDetail)`

SetShippingDetails sets ShippingDetails field to given value.

### HasShippingDetails

`func (o *PaymentLink) HasShippingDetails() bool`

HasShippingDetails returns a boolean if a field has been set.

### SetShippingDetailsNil

`func (o *PaymentLink) SetShippingDetailsNil(b bool)`

 SetShippingDetailsNil sets the value for ShippingDetails to be an explicit nil

### UnsetShippingDetails
`func (o *PaymentLink) UnsetShippingDetails()`

UnsetShippingDetails ensures that no value is present for ShippingDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


