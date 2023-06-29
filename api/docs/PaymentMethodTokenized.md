# PaymentMethodTokenized

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;payment-method&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique ID of the payment method. | [optional] 
**MerchantAccountId** | Pointer to **string** | The unique ID for a merchant account. | [optional] 
**Method** | Pointer to **string** | The type of this payment method. | [optional] 
**Label** | Pointer to **string** | A label for the payment method. For a &#x60;card&#x60; payment method this is the last 4 digits on the card. For others it would be the email address. | [optional] 
**Scheme** | Pointer to **NullableString** | The type of the card, if the payment method is a card. | [optional] 
**ExpirationDate** | Pointer to **NullableString** | The expiration date for the payment method. | [optional] 
**ApprovalTarget** | Pointer to **NullableString** | The browser target that an approval URL must be opened in. If &#x60;any&#x60; or &#x60;null&#x60;, then there is no specific requirement. | [optional] 
**ApprovalUrl** | Pointer to **NullableString** | The optional URL that the buyer needs to be redirected to to further authorize their payment. | [optional] 
**Currency** | Pointer to **NullableString** | The ISO-4217 currency code that this payment method can be used for. If this value is &#x60;null&#x60; the payment method may be used for multiple currencies. | [optional] 
**Country** | Pointer to **NullableString** | The 2-letter ISO code of the country this payment method can be used for. If this value is &#x60;null&#x60; the payment method may be used in multiple countries. | [optional] 

## Methods

### NewPaymentMethodTokenized

`func NewPaymentMethodTokenized() *PaymentMethodTokenized`

NewPaymentMethodTokenized instantiates a new PaymentMethodTokenized object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodTokenizedWithDefaults

`func NewPaymentMethodTokenizedWithDefaults() *PaymentMethodTokenized`

NewPaymentMethodTokenizedWithDefaults instantiates a new PaymentMethodTokenized object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodTokenized) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodTokenized) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodTokenized) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentMethodTokenized) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *PaymentMethodTokenized) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodTokenized) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodTokenized) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentMethodTokenized) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantAccountId

`func (o *PaymentMethodTokenized) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *PaymentMethodTokenized) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *PaymentMethodTokenized) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.

### HasMerchantAccountId

`func (o *PaymentMethodTokenized) HasMerchantAccountId() bool`

HasMerchantAccountId returns a boolean if a field has been set.

### GetMethod

`func (o *PaymentMethodTokenized) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PaymentMethodTokenized) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PaymentMethodTokenized) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PaymentMethodTokenized) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetLabel

`func (o *PaymentMethodTokenized) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PaymentMethodTokenized) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PaymentMethodTokenized) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PaymentMethodTokenized) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetScheme

`func (o *PaymentMethodTokenized) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *PaymentMethodTokenized) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *PaymentMethodTokenized) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *PaymentMethodTokenized) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *PaymentMethodTokenized) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *PaymentMethodTokenized) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetExpirationDate

`func (o *PaymentMethodTokenized) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *PaymentMethodTokenized) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *PaymentMethodTokenized) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *PaymentMethodTokenized) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *PaymentMethodTokenized) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *PaymentMethodTokenized) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetApprovalTarget

`func (o *PaymentMethodTokenized) GetApprovalTarget() string`

GetApprovalTarget returns the ApprovalTarget field if non-nil, zero value otherwise.

### GetApprovalTargetOk

`func (o *PaymentMethodTokenized) GetApprovalTargetOk() (*string, bool)`

GetApprovalTargetOk returns a tuple with the ApprovalTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalTarget

`func (o *PaymentMethodTokenized) SetApprovalTarget(v string)`

SetApprovalTarget sets ApprovalTarget field to given value.

### HasApprovalTarget

`func (o *PaymentMethodTokenized) HasApprovalTarget() bool`

HasApprovalTarget returns a boolean if a field has been set.

### SetApprovalTargetNil

`func (o *PaymentMethodTokenized) SetApprovalTargetNil(b bool)`

 SetApprovalTargetNil sets the value for ApprovalTarget to be an explicit nil

### UnsetApprovalTarget
`func (o *PaymentMethodTokenized) UnsetApprovalTarget()`

UnsetApprovalTarget ensures that no value is present for ApprovalTarget, not even an explicit nil
### GetApprovalUrl

`func (o *PaymentMethodTokenized) GetApprovalUrl() string`

GetApprovalUrl returns the ApprovalUrl field if non-nil, zero value otherwise.

### GetApprovalUrlOk

`func (o *PaymentMethodTokenized) GetApprovalUrlOk() (*string, bool)`

GetApprovalUrlOk returns a tuple with the ApprovalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalUrl

`func (o *PaymentMethodTokenized) SetApprovalUrl(v string)`

SetApprovalUrl sets ApprovalUrl field to given value.

### HasApprovalUrl

`func (o *PaymentMethodTokenized) HasApprovalUrl() bool`

HasApprovalUrl returns a boolean if a field has been set.

### SetApprovalUrlNil

`func (o *PaymentMethodTokenized) SetApprovalUrlNil(b bool)`

 SetApprovalUrlNil sets the value for ApprovalUrl to be an explicit nil

### UnsetApprovalUrl
`func (o *PaymentMethodTokenized) UnsetApprovalUrl()`

UnsetApprovalUrl ensures that no value is present for ApprovalUrl, not even an explicit nil
### GetCurrency

`func (o *PaymentMethodTokenized) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentMethodTokenized) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentMethodTokenized) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentMethodTokenized) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *PaymentMethodTokenized) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *PaymentMethodTokenized) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetCountry

`func (o *PaymentMethodTokenized) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaymentMethodTokenized) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaymentMethodTokenized) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PaymentMethodTokenized) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *PaymentMethodTokenized) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *PaymentMethodTokenized) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


