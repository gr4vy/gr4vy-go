# PaymentLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | The amount to request payment for. | 
**Currency** | **string** | The ISO-4217 currency code for the payment. | 
**Country** | **string** | The 2-letter ISO code of the country of the transaction. This is used to filter the payment services that is used to process the transaction.  | 
**Buyer** | Pointer to [**TransactionBuyerRequest**](TransactionBuyerRequest.md) |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** | The date and time when this payment link expires. Defaults to 24 hours from creation. | [optional] 
**ConnectionOptions** | Pointer to [**NullableConnectionOptions**](ConnectionOptions.md) | Allows for passing optional configuration per connection to take advantage of connection specific features. When provided, the data is only passed to the target connection type to prevent sharing configuration across connections.  Please note that each of the keys this object are in kebab-case, for example &#x60;cybersource-anti-fraud&#x60; as they represent the ID of the connector. All the other keys will be snake case, for example &#x60;merchant_defined_data&#x60; or camel case to match an external API that the connector uses. | [optional] 
**ExternalIdentifier** | Pointer to **NullableString** | A value that can be used to match the payment link against your own records. This will also be applied to any transaction. | [optional] 
**StatementDescriptor** | Pointer to [**NullableStatementDescriptor**](StatementDescriptor.md) |  | [optional] 
**Locale** | Pointer to **NullableString** | The locale used to translate text within the payment link. | [optional] 
**MerchantName** | Pointer to **NullableString** | The name of the merchant to display on the payment link. | [optional] 
**MerchantUrl** | Pointer to **NullableString** | The URL of the merchant to display on the payment link. | [optional] 
**MerchantBannerUrl** | Pointer to **NullableString** | The URL of the merchant banner to display on the payment link. | [optional] 
**MerchantColor** | Pointer to **NullableString** | The color code of the merchant to display on the payment link. | [optional] 
**MerchantMessage** | Pointer to **NullableString** | The message to display on the payment link. | [optional] 
**MerchantTermsAndConditionsUrl** | Pointer to **NullableString** | The URL of the merchant terms and conditions to display on the payment link. | [optional] 
**MerchantFaviconUrl** | Pointer to **NullableString** | The URL of the merchant favicon icon. | [optional] 
**Intent** | Pointer to **NullableString** | The intent of the payment link. | [optional] [default to "authorize"]
**ReturnUrl** | Pointer to **NullableString** | The URL to redirect the buyer to after payment. | [optional] 
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) | An array of cart items that represents the line items of a payment link. | [optional] 
**Metadata** | Pointer to **map[string]string** | Any additional information about the payment link that you would like to store as key-value pairs. This data is passed to payment service providers that support it. | [optional] 
**PaymentSource** | Pointer to **NullableString** | The source of the payment link. | [optional] 

## Methods

### NewPaymentLinkRequest

`func NewPaymentLinkRequest(amount float32, currency string, country string, ) *PaymentLinkRequest`

NewPaymentLinkRequest instantiates a new PaymentLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentLinkRequestWithDefaults

`func NewPaymentLinkRequestWithDefaults() *PaymentLinkRequest`

NewPaymentLinkRequestWithDefaults instantiates a new PaymentLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PaymentLinkRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentLinkRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentLinkRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *PaymentLinkRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentLinkRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentLinkRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCountry

`func (o *PaymentLinkRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaymentLinkRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaymentLinkRequest) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetBuyer

`func (o *PaymentLinkRequest) GetBuyer() TransactionBuyerRequest`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *PaymentLinkRequest) GetBuyerOk() (*TransactionBuyerRequest, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *PaymentLinkRequest) SetBuyer(v TransactionBuyerRequest)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *PaymentLinkRequest) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PaymentLinkRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PaymentLinkRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PaymentLinkRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PaymentLinkRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *PaymentLinkRequest) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *PaymentLinkRequest) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetConnectionOptions

`func (o *PaymentLinkRequest) GetConnectionOptions() ConnectionOptions`

GetConnectionOptions returns the ConnectionOptions field if non-nil, zero value otherwise.

### GetConnectionOptionsOk

`func (o *PaymentLinkRequest) GetConnectionOptionsOk() (*ConnectionOptions, bool)`

GetConnectionOptionsOk returns a tuple with the ConnectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionOptions

`func (o *PaymentLinkRequest) SetConnectionOptions(v ConnectionOptions)`

SetConnectionOptions sets ConnectionOptions field to given value.

### HasConnectionOptions

`func (o *PaymentLinkRequest) HasConnectionOptions() bool`

HasConnectionOptions returns a boolean if a field has been set.

### SetConnectionOptionsNil

`func (o *PaymentLinkRequest) SetConnectionOptionsNil(b bool)`

 SetConnectionOptionsNil sets the value for ConnectionOptions to be an explicit nil

### UnsetConnectionOptions
`func (o *PaymentLinkRequest) UnsetConnectionOptions()`

UnsetConnectionOptions ensures that no value is present for ConnectionOptions, not even an explicit nil
### GetExternalIdentifier

`func (o *PaymentLinkRequest) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *PaymentLinkRequest) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *PaymentLinkRequest) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *PaymentLinkRequest) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### SetExternalIdentifierNil

`func (o *PaymentLinkRequest) SetExternalIdentifierNil(b bool)`

 SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil

### UnsetExternalIdentifier
`func (o *PaymentLinkRequest) UnsetExternalIdentifier()`

UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
### GetStatementDescriptor

`func (o *PaymentLinkRequest) GetStatementDescriptor() StatementDescriptor`

GetStatementDescriptor returns the StatementDescriptor field if non-nil, zero value otherwise.

### GetStatementDescriptorOk

`func (o *PaymentLinkRequest) GetStatementDescriptorOk() (*StatementDescriptor, bool)`

GetStatementDescriptorOk returns a tuple with the StatementDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescriptor

`func (o *PaymentLinkRequest) SetStatementDescriptor(v StatementDescriptor)`

SetStatementDescriptor sets StatementDescriptor field to given value.

### HasStatementDescriptor

`func (o *PaymentLinkRequest) HasStatementDescriptor() bool`

HasStatementDescriptor returns a boolean if a field has been set.

### SetStatementDescriptorNil

`func (o *PaymentLinkRequest) SetStatementDescriptorNil(b bool)`

 SetStatementDescriptorNil sets the value for StatementDescriptor to be an explicit nil

### UnsetStatementDescriptor
`func (o *PaymentLinkRequest) UnsetStatementDescriptor()`

UnsetStatementDescriptor ensures that no value is present for StatementDescriptor, not even an explicit nil
### GetLocale

`func (o *PaymentLinkRequest) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *PaymentLinkRequest) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *PaymentLinkRequest) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *PaymentLinkRequest) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *PaymentLinkRequest) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *PaymentLinkRequest) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetMerchantName

`func (o *PaymentLinkRequest) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *PaymentLinkRequest) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *PaymentLinkRequest) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *PaymentLinkRequest) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### SetMerchantNameNil

`func (o *PaymentLinkRequest) SetMerchantNameNil(b bool)`

 SetMerchantNameNil sets the value for MerchantName to be an explicit nil

### UnsetMerchantName
`func (o *PaymentLinkRequest) UnsetMerchantName()`

UnsetMerchantName ensures that no value is present for MerchantName, not even an explicit nil
### GetMerchantUrl

`func (o *PaymentLinkRequest) GetMerchantUrl() string`

GetMerchantUrl returns the MerchantUrl field if non-nil, zero value otherwise.

### GetMerchantUrlOk

`func (o *PaymentLinkRequest) GetMerchantUrlOk() (*string, bool)`

GetMerchantUrlOk returns a tuple with the MerchantUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantUrl

`func (o *PaymentLinkRequest) SetMerchantUrl(v string)`

SetMerchantUrl sets MerchantUrl field to given value.

### HasMerchantUrl

`func (o *PaymentLinkRequest) HasMerchantUrl() bool`

HasMerchantUrl returns a boolean if a field has been set.

### SetMerchantUrlNil

`func (o *PaymentLinkRequest) SetMerchantUrlNil(b bool)`

 SetMerchantUrlNil sets the value for MerchantUrl to be an explicit nil

### UnsetMerchantUrl
`func (o *PaymentLinkRequest) UnsetMerchantUrl()`

UnsetMerchantUrl ensures that no value is present for MerchantUrl, not even an explicit nil
### GetMerchantBannerUrl

`func (o *PaymentLinkRequest) GetMerchantBannerUrl() string`

GetMerchantBannerUrl returns the MerchantBannerUrl field if non-nil, zero value otherwise.

### GetMerchantBannerUrlOk

`func (o *PaymentLinkRequest) GetMerchantBannerUrlOk() (*string, bool)`

GetMerchantBannerUrlOk returns a tuple with the MerchantBannerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantBannerUrl

`func (o *PaymentLinkRequest) SetMerchantBannerUrl(v string)`

SetMerchantBannerUrl sets MerchantBannerUrl field to given value.

### HasMerchantBannerUrl

`func (o *PaymentLinkRequest) HasMerchantBannerUrl() bool`

HasMerchantBannerUrl returns a boolean if a field has been set.

### SetMerchantBannerUrlNil

`func (o *PaymentLinkRequest) SetMerchantBannerUrlNil(b bool)`

 SetMerchantBannerUrlNil sets the value for MerchantBannerUrl to be an explicit nil

### UnsetMerchantBannerUrl
`func (o *PaymentLinkRequest) UnsetMerchantBannerUrl()`

UnsetMerchantBannerUrl ensures that no value is present for MerchantBannerUrl, not even an explicit nil
### GetMerchantColor

`func (o *PaymentLinkRequest) GetMerchantColor() string`

GetMerchantColor returns the MerchantColor field if non-nil, zero value otherwise.

### GetMerchantColorOk

`func (o *PaymentLinkRequest) GetMerchantColorOk() (*string, bool)`

GetMerchantColorOk returns a tuple with the MerchantColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantColor

`func (o *PaymentLinkRequest) SetMerchantColor(v string)`

SetMerchantColor sets MerchantColor field to given value.

### HasMerchantColor

`func (o *PaymentLinkRequest) HasMerchantColor() bool`

HasMerchantColor returns a boolean if a field has been set.

### SetMerchantColorNil

`func (o *PaymentLinkRequest) SetMerchantColorNil(b bool)`

 SetMerchantColorNil sets the value for MerchantColor to be an explicit nil

### UnsetMerchantColor
`func (o *PaymentLinkRequest) UnsetMerchantColor()`

UnsetMerchantColor ensures that no value is present for MerchantColor, not even an explicit nil
### GetMerchantMessage

`func (o *PaymentLinkRequest) GetMerchantMessage() string`

GetMerchantMessage returns the MerchantMessage field if non-nil, zero value otherwise.

### GetMerchantMessageOk

`func (o *PaymentLinkRequest) GetMerchantMessageOk() (*string, bool)`

GetMerchantMessageOk returns a tuple with the MerchantMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMessage

`func (o *PaymentLinkRequest) SetMerchantMessage(v string)`

SetMerchantMessage sets MerchantMessage field to given value.

### HasMerchantMessage

`func (o *PaymentLinkRequest) HasMerchantMessage() bool`

HasMerchantMessage returns a boolean if a field has been set.

### SetMerchantMessageNil

`func (o *PaymentLinkRequest) SetMerchantMessageNil(b bool)`

 SetMerchantMessageNil sets the value for MerchantMessage to be an explicit nil

### UnsetMerchantMessage
`func (o *PaymentLinkRequest) UnsetMerchantMessage()`

UnsetMerchantMessage ensures that no value is present for MerchantMessage, not even an explicit nil
### GetMerchantTermsAndConditionsUrl

`func (o *PaymentLinkRequest) GetMerchantTermsAndConditionsUrl() string`

GetMerchantTermsAndConditionsUrl returns the MerchantTermsAndConditionsUrl field if non-nil, zero value otherwise.

### GetMerchantTermsAndConditionsUrlOk

`func (o *PaymentLinkRequest) GetMerchantTermsAndConditionsUrlOk() (*string, bool)`

GetMerchantTermsAndConditionsUrlOk returns a tuple with the MerchantTermsAndConditionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTermsAndConditionsUrl

`func (o *PaymentLinkRequest) SetMerchantTermsAndConditionsUrl(v string)`

SetMerchantTermsAndConditionsUrl sets MerchantTermsAndConditionsUrl field to given value.

### HasMerchantTermsAndConditionsUrl

`func (o *PaymentLinkRequest) HasMerchantTermsAndConditionsUrl() bool`

HasMerchantTermsAndConditionsUrl returns a boolean if a field has been set.

### SetMerchantTermsAndConditionsUrlNil

`func (o *PaymentLinkRequest) SetMerchantTermsAndConditionsUrlNil(b bool)`

 SetMerchantTermsAndConditionsUrlNil sets the value for MerchantTermsAndConditionsUrl to be an explicit nil

### UnsetMerchantTermsAndConditionsUrl
`func (o *PaymentLinkRequest) UnsetMerchantTermsAndConditionsUrl()`

UnsetMerchantTermsAndConditionsUrl ensures that no value is present for MerchantTermsAndConditionsUrl, not even an explicit nil
### GetMerchantFaviconUrl

`func (o *PaymentLinkRequest) GetMerchantFaviconUrl() string`

GetMerchantFaviconUrl returns the MerchantFaviconUrl field if non-nil, zero value otherwise.

### GetMerchantFaviconUrlOk

`func (o *PaymentLinkRequest) GetMerchantFaviconUrlOk() (*string, bool)`

GetMerchantFaviconUrlOk returns a tuple with the MerchantFaviconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantFaviconUrl

`func (o *PaymentLinkRequest) SetMerchantFaviconUrl(v string)`

SetMerchantFaviconUrl sets MerchantFaviconUrl field to given value.

### HasMerchantFaviconUrl

`func (o *PaymentLinkRequest) HasMerchantFaviconUrl() bool`

HasMerchantFaviconUrl returns a boolean if a field has been set.

### SetMerchantFaviconUrlNil

`func (o *PaymentLinkRequest) SetMerchantFaviconUrlNil(b bool)`

 SetMerchantFaviconUrlNil sets the value for MerchantFaviconUrl to be an explicit nil

### UnsetMerchantFaviconUrl
`func (o *PaymentLinkRequest) UnsetMerchantFaviconUrl()`

UnsetMerchantFaviconUrl ensures that no value is present for MerchantFaviconUrl, not even an explicit nil
### GetIntent

`func (o *PaymentLinkRequest) GetIntent() string`

GetIntent returns the Intent field if non-nil, zero value otherwise.

### GetIntentOk

`func (o *PaymentLinkRequest) GetIntentOk() (*string, bool)`

GetIntentOk returns a tuple with the Intent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntent

`func (o *PaymentLinkRequest) SetIntent(v string)`

SetIntent sets Intent field to given value.

### HasIntent

`func (o *PaymentLinkRequest) HasIntent() bool`

HasIntent returns a boolean if a field has been set.

### SetIntentNil

`func (o *PaymentLinkRequest) SetIntentNil(b bool)`

 SetIntentNil sets the value for Intent to be an explicit nil

### UnsetIntent
`func (o *PaymentLinkRequest) UnsetIntent()`

UnsetIntent ensures that no value is present for Intent, not even an explicit nil
### GetReturnUrl

`func (o *PaymentLinkRequest) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *PaymentLinkRequest) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *PaymentLinkRequest) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *PaymentLinkRequest) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### SetReturnUrlNil

`func (o *PaymentLinkRequest) SetReturnUrlNil(b bool)`

 SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil

### UnsetReturnUrl
`func (o *PaymentLinkRequest) UnsetReturnUrl()`

UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
### GetCartItems

`func (o *PaymentLinkRequest) GetCartItems() []CartItem`

GetCartItems returns the CartItems field if non-nil, zero value otherwise.

### GetCartItemsOk

`func (o *PaymentLinkRequest) GetCartItemsOk() (*[]CartItem, bool)`

GetCartItemsOk returns a tuple with the CartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItems

`func (o *PaymentLinkRequest) SetCartItems(v []CartItem)`

SetCartItems sets CartItems field to given value.

### HasCartItems

`func (o *PaymentLinkRequest) HasCartItems() bool`

HasCartItems returns a boolean if a field has been set.

### GetMetadata

`func (o *PaymentLinkRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentLinkRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentLinkRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaymentLinkRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PaymentLinkRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PaymentLinkRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetPaymentSource

`func (o *PaymentLinkRequest) GetPaymentSource() string`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *PaymentLinkRequest) GetPaymentSourceOk() (*string, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *PaymentLinkRequest) SetPaymentSource(v string)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *PaymentLinkRequest) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.

### SetPaymentSourceNil

`func (o *PaymentLinkRequest) SetPaymentSourceNil(b bool)`

 SetPaymentSourceNil sets the value for PaymentSource to be an explicit nil

### UnsetPaymentSource
`func (o *PaymentLinkRequest) UnsetPaymentSource()`

UnsetPaymentSource ensures that no value is present for PaymentSource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


