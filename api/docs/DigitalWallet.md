# DigitalWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;digital-wallet&#x60;. | [optional] 
**Provider** | Pointer to **string** | The name of the digital wallet provider. | [optional] 
**Id** | Pointer to **string** | The ID of the registered digital wallet. | [optional] 
**MerchantName** | Pointer to **string** | The name of the merchant the digital wallet is registered to. | [optional] 
**MerchantUrl** | Pointer to **NullableString** | The main URL of the merchant. | [optional] [default to "null"]
**DomainNames** | Pointer to **[]string** | The list of fully qualified domain names that a digital wallet provider processes payments for. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this digital wallet was registered. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when this digital wallet was last updated. | [optional] 
**Environments** | Pointer to **[]string** | The Gr4vy environments in which this digital wallet is available. | [optional] [default to ["production"]]

## Methods

### NewDigitalWallet

`func NewDigitalWallet() *DigitalWallet`

NewDigitalWallet instantiates a new DigitalWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletWithDefaults

`func NewDigitalWalletWithDefaults() *DigitalWallet`

NewDigitalWalletWithDefaults instantiates a new DigitalWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DigitalWallet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DigitalWallet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DigitalWallet) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DigitalWallet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProvider

`func (o *DigitalWallet) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DigitalWallet) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DigitalWallet) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *DigitalWallet) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetId

`func (o *DigitalWallet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DigitalWallet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DigitalWallet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DigitalWallet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantName

`func (o *DigitalWallet) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *DigitalWallet) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *DigitalWallet) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *DigitalWallet) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### GetMerchantUrl

`func (o *DigitalWallet) GetMerchantUrl() string`

GetMerchantUrl returns the MerchantUrl field if non-nil, zero value otherwise.

### GetMerchantUrlOk

`func (o *DigitalWallet) GetMerchantUrlOk() (*string, bool)`

GetMerchantUrlOk returns a tuple with the MerchantUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantUrl

`func (o *DigitalWallet) SetMerchantUrl(v string)`

SetMerchantUrl sets MerchantUrl field to given value.

### HasMerchantUrl

`func (o *DigitalWallet) HasMerchantUrl() bool`

HasMerchantUrl returns a boolean if a field has been set.

### SetMerchantUrlNil

`func (o *DigitalWallet) SetMerchantUrlNil(b bool)`

 SetMerchantUrlNil sets the value for MerchantUrl to be an explicit nil

### UnsetMerchantUrl
`func (o *DigitalWallet) UnsetMerchantUrl()`

UnsetMerchantUrl ensures that no value is present for MerchantUrl, not even an explicit nil
### GetDomainNames

`func (o *DigitalWallet) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *DigitalWallet) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *DigitalWallet) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *DigitalWallet) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DigitalWallet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DigitalWallet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DigitalWallet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DigitalWallet) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DigitalWallet) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DigitalWallet) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DigitalWallet) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DigitalWallet) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEnvironments

`func (o *DigitalWallet) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *DigitalWallet) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *DigitalWallet) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *DigitalWallet) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


