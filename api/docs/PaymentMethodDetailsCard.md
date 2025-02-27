# PaymentMethodDetailsCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bin** | Pointer to **string** | The first 6 digits of the full card number (the BIN). | [optional] 
**CardType** | Pointer to **string** | The type of card, one of &#x60;credit&#x60;, &#x60;debit&#x60; or &#x60;prepaid&#x60;. | [optional] 
**CardIssuerName** | Pointer to **string** | The name of the card issuer. | [optional] 

## Methods

### NewPaymentMethodDetailsCard

`func NewPaymentMethodDetailsCard() *PaymentMethodDetailsCard`

NewPaymentMethodDetailsCard instantiates a new PaymentMethodDetailsCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodDetailsCardWithDefaults

`func NewPaymentMethodDetailsCardWithDefaults() *PaymentMethodDetailsCard`

NewPaymentMethodDetailsCardWithDefaults instantiates a new PaymentMethodDetailsCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBin

`func (o *PaymentMethodDetailsCard) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *PaymentMethodDetailsCard) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *PaymentMethodDetailsCard) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *PaymentMethodDetailsCard) HasBin() bool`

HasBin returns a boolean if a field has been set.

### GetCardType

`func (o *PaymentMethodDetailsCard) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *PaymentMethodDetailsCard) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *PaymentMethodDetailsCard) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *PaymentMethodDetailsCard) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetCardIssuerName

`func (o *PaymentMethodDetailsCard) GetCardIssuerName() string`

GetCardIssuerName returns the CardIssuerName field if non-nil, zero value otherwise.

### GetCardIssuerNameOk

`func (o *PaymentMethodDetailsCard) GetCardIssuerNameOk() (*string, bool)`

GetCardIssuerNameOk returns a tuple with the CardIssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIssuerName

`func (o *PaymentMethodDetailsCard) SetCardIssuerName(v string)`

SetCardIssuerName sets CardIssuerName field to given value.

### HasCardIssuerName

`func (o *PaymentMethodDetailsCard) HasCardIssuerName() bool`

HasCardIssuerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


