# DigitalWalletUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantName** | Pointer to **string** | The name of the merchant. This is used to update the value initially used to register with a digital wallet provider and this name is not displayed to the buyer. | [optional] 
**DomainNames** | Pointer to **[]string** | The list of domain names that a digital wallet can be used on. To use a digital wallet on a website, the domain of the site is required to be in this list. | [optional] 

## Methods

### NewDigitalWalletUpdate

`func NewDigitalWalletUpdate() *DigitalWalletUpdate`

NewDigitalWalletUpdate instantiates a new DigitalWalletUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletUpdateWithDefaults

`func NewDigitalWalletUpdateWithDefaults() *DigitalWalletUpdate`

NewDigitalWalletUpdateWithDefaults instantiates a new DigitalWalletUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantName

`func (o *DigitalWalletUpdate) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *DigitalWalletUpdate) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *DigitalWalletUpdate) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *DigitalWalletUpdate) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### GetDomainNames

`func (o *DigitalWalletUpdate) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *DigitalWalletUpdate) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *DigitalWalletUpdate) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *DigitalWalletUpdate) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


