# MerchantAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;merchant-account&#x60;. | [optional] 
**Id** | Pointer to **string** | The ID for this merchant account. | [optional] 
**DisplayName** | Pointer to **string** | The display name of this merchant account. | [optional] 
**OutboundWebhookUrl** | Pointer to **NullableString** | The optional URL where webhooks will be received. | [optional] 
**OutboundWebhookUsername** | Pointer to **NullableString** | The optional username to use when &#x60;outbound_webhook_url&#x60; is configured and requires basic authentication. | [optional] 
**VisaNetworkTokensRequestorId** | Pointer to **NullableString** | Requestor ID provided for Visa after onboarding to use Network Tokens. | [optional] 
**VisaNetworkTokensAppId** | Pointer to **NullableString** | Application ID provided for Visa after onboarding to use Network Tokens. | [optional] 
**AmexNetworkTokensRequestorId** | Pointer to **NullableString** | Requestor ID provided for Amex after onboarding to use Network Tokens. | [optional] 
**AmexNetworkTokensAppId** | Pointer to **NullableString** | Application ID provided for Amex after onboarding to use Network Tokens. | [optional] 
**MastercardNetworkTokensRequestorId** | Pointer to **NullableString** | Requestor ID provided for Mastercard after onboarding to use Network Tokens. | [optional] 
**MastercardNetworkTokensAppId** | Pointer to **NullableString** | Application ID provided for Mastercard after onboarding to use Network Tokens. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this merchant account was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when this merchant account was updated. | [optional] 

## Methods

### NewMerchantAccount

`func NewMerchantAccount() *MerchantAccount`

NewMerchantAccount instantiates a new MerchantAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantAccountWithDefaults

`func NewMerchantAccountWithDefaults() *MerchantAccount`

NewMerchantAccountWithDefaults instantiates a new MerchantAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MerchantAccount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MerchantAccount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MerchantAccount) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MerchantAccount) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *MerchantAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MerchantAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MerchantAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MerchantAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *MerchantAccount) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MerchantAccount) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MerchantAccount) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MerchantAccount) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetOutboundWebhookUrl

`func (o *MerchantAccount) GetOutboundWebhookUrl() string`

GetOutboundWebhookUrl returns the OutboundWebhookUrl field if non-nil, zero value otherwise.

### GetOutboundWebhookUrlOk

`func (o *MerchantAccount) GetOutboundWebhookUrlOk() (*string, bool)`

GetOutboundWebhookUrlOk returns a tuple with the OutboundWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundWebhookUrl

`func (o *MerchantAccount) SetOutboundWebhookUrl(v string)`

SetOutboundWebhookUrl sets OutboundWebhookUrl field to given value.

### HasOutboundWebhookUrl

`func (o *MerchantAccount) HasOutboundWebhookUrl() bool`

HasOutboundWebhookUrl returns a boolean if a field has been set.

### SetOutboundWebhookUrlNil

`func (o *MerchantAccount) SetOutboundWebhookUrlNil(b bool)`

 SetOutboundWebhookUrlNil sets the value for OutboundWebhookUrl to be an explicit nil

### UnsetOutboundWebhookUrl
`func (o *MerchantAccount) UnsetOutboundWebhookUrl()`

UnsetOutboundWebhookUrl ensures that no value is present for OutboundWebhookUrl, not even an explicit nil
### GetOutboundWebhookUsername

`func (o *MerchantAccount) GetOutboundWebhookUsername() string`

GetOutboundWebhookUsername returns the OutboundWebhookUsername field if non-nil, zero value otherwise.

### GetOutboundWebhookUsernameOk

`func (o *MerchantAccount) GetOutboundWebhookUsernameOk() (*string, bool)`

GetOutboundWebhookUsernameOk returns a tuple with the OutboundWebhookUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundWebhookUsername

`func (o *MerchantAccount) SetOutboundWebhookUsername(v string)`

SetOutboundWebhookUsername sets OutboundWebhookUsername field to given value.

### HasOutboundWebhookUsername

`func (o *MerchantAccount) HasOutboundWebhookUsername() bool`

HasOutboundWebhookUsername returns a boolean if a field has been set.

### SetOutboundWebhookUsernameNil

`func (o *MerchantAccount) SetOutboundWebhookUsernameNil(b bool)`

 SetOutboundWebhookUsernameNil sets the value for OutboundWebhookUsername to be an explicit nil

### UnsetOutboundWebhookUsername
`func (o *MerchantAccount) UnsetOutboundWebhookUsername()`

UnsetOutboundWebhookUsername ensures that no value is present for OutboundWebhookUsername, not even an explicit nil
### GetVisaNetworkTokensRequestorId

`func (o *MerchantAccount) GetVisaNetworkTokensRequestorId() string`

GetVisaNetworkTokensRequestorId returns the VisaNetworkTokensRequestorId field if non-nil, zero value otherwise.

### GetVisaNetworkTokensRequestorIdOk

`func (o *MerchantAccount) GetVisaNetworkTokensRequestorIdOk() (*string, bool)`

GetVisaNetworkTokensRequestorIdOk returns a tuple with the VisaNetworkTokensRequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisaNetworkTokensRequestorId

`func (o *MerchantAccount) SetVisaNetworkTokensRequestorId(v string)`

SetVisaNetworkTokensRequestorId sets VisaNetworkTokensRequestorId field to given value.

### HasVisaNetworkTokensRequestorId

`func (o *MerchantAccount) HasVisaNetworkTokensRequestorId() bool`

HasVisaNetworkTokensRequestorId returns a boolean if a field has been set.

### SetVisaNetworkTokensRequestorIdNil

`func (o *MerchantAccount) SetVisaNetworkTokensRequestorIdNil(b bool)`

 SetVisaNetworkTokensRequestorIdNil sets the value for VisaNetworkTokensRequestorId to be an explicit nil

### UnsetVisaNetworkTokensRequestorId
`func (o *MerchantAccount) UnsetVisaNetworkTokensRequestorId()`

UnsetVisaNetworkTokensRequestorId ensures that no value is present for VisaNetworkTokensRequestorId, not even an explicit nil
### GetVisaNetworkTokensAppId

`func (o *MerchantAccount) GetVisaNetworkTokensAppId() string`

GetVisaNetworkTokensAppId returns the VisaNetworkTokensAppId field if non-nil, zero value otherwise.

### GetVisaNetworkTokensAppIdOk

`func (o *MerchantAccount) GetVisaNetworkTokensAppIdOk() (*string, bool)`

GetVisaNetworkTokensAppIdOk returns a tuple with the VisaNetworkTokensAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisaNetworkTokensAppId

`func (o *MerchantAccount) SetVisaNetworkTokensAppId(v string)`

SetVisaNetworkTokensAppId sets VisaNetworkTokensAppId field to given value.

### HasVisaNetworkTokensAppId

`func (o *MerchantAccount) HasVisaNetworkTokensAppId() bool`

HasVisaNetworkTokensAppId returns a boolean if a field has been set.

### SetVisaNetworkTokensAppIdNil

`func (o *MerchantAccount) SetVisaNetworkTokensAppIdNil(b bool)`

 SetVisaNetworkTokensAppIdNil sets the value for VisaNetworkTokensAppId to be an explicit nil

### UnsetVisaNetworkTokensAppId
`func (o *MerchantAccount) UnsetVisaNetworkTokensAppId()`

UnsetVisaNetworkTokensAppId ensures that no value is present for VisaNetworkTokensAppId, not even an explicit nil
### GetAmexNetworkTokensRequestorId

`func (o *MerchantAccount) GetAmexNetworkTokensRequestorId() string`

GetAmexNetworkTokensRequestorId returns the AmexNetworkTokensRequestorId field if non-nil, zero value otherwise.

### GetAmexNetworkTokensRequestorIdOk

`func (o *MerchantAccount) GetAmexNetworkTokensRequestorIdOk() (*string, bool)`

GetAmexNetworkTokensRequestorIdOk returns a tuple with the AmexNetworkTokensRequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmexNetworkTokensRequestorId

`func (o *MerchantAccount) SetAmexNetworkTokensRequestorId(v string)`

SetAmexNetworkTokensRequestorId sets AmexNetworkTokensRequestorId field to given value.

### HasAmexNetworkTokensRequestorId

`func (o *MerchantAccount) HasAmexNetworkTokensRequestorId() bool`

HasAmexNetworkTokensRequestorId returns a boolean if a field has been set.

### SetAmexNetworkTokensRequestorIdNil

`func (o *MerchantAccount) SetAmexNetworkTokensRequestorIdNil(b bool)`

 SetAmexNetworkTokensRequestorIdNil sets the value for AmexNetworkTokensRequestorId to be an explicit nil

### UnsetAmexNetworkTokensRequestorId
`func (o *MerchantAccount) UnsetAmexNetworkTokensRequestorId()`

UnsetAmexNetworkTokensRequestorId ensures that no value is present for AmexNetworkTokensRequestorId, not even an explicit nil
### GetAmexNetworkTokensAppId

`func (o *MerchantAccount) GetAmexNetworkTokensAppId() string`

GetAmexNetworkTokensAppId returns the AmexNetworkTokensAppId field if non-nil, zero value otherwise.

### GetAmexNetworkTokensAppIdOk

`func (o *MerchantAccount) GetAmexNetworkTokensAppIdOk() (*string, bool)`

GetAmexNetworkTokensAppIdOk returns a tuple with the AmexNetworkTokensAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmexNetworkTokensAppId

`func (o *MerchantAccount) SetAmexNetworkTokensAppId(v string)`

SetAmexNetworkTokensAppId sets AmexNetworkTokensAppId field to given value.

### HasAmexNetworkTokensAppId

`func (o *MerchantAccount) HasAmexNetworkTokensAppId() bool`

HasAmexNetworkTokensAppId returns a boolean if a field has been set.

### SetAmexNetworkTokensAppIdNil

`func (o *MerchantAccount) SetAmexNetworkTokensAppIdNil(b bool)`

 SetAmexNetworkTokensAppIdNil sets the value for AmexNetworkTokensAppId to be an explicit nil

### UnsetAmexNetworkTokensAppId
`func (o *MerchantAccount) UnsetAmexNetworkTokensAppId()`

UnsetAmexNetworkTokensAppId ensures that no value is present for AmexNetworkTokensAppId, not even an explicit nil
### GetMastercardNetworkTokensRequestorId

`func (o *MerchantAccount) GetMastercardNetworkTokensRequestorId() string`

GetMastercardNetworkTokensRequestorId returns the MastercardNetworkTokensRequestorId field if non-nil, zero value otherwise.

### GetMastercardNetworkTokensRequestorIdOk

`func (o *MerchantAccount) GetMastercardNetworkTokensRequestorIdOk() (*string, bool)`

GetMastercardNetworkTokensRequestorIdOk returns a tuple with the MastercardNetworkTokensRequestorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMastercardNetworkTokensRequestorId

`func (o *MerchantAccount) SetMastercardNetworkTokensRequestorId(v string)`

SetMastercardNetworkTokensRequestorId sets MastercardNetworkTokensRequestorId field to given value.

### HasMastercardNetworkTokensRequestorId

`func (o *MerchantAccount) HasMastercardNetworkTokensRequestorId() bool`

HasMastercardNetworkTokensRequestorId returns a boolean if a field has been set.

### SetMastercardNetworkTokensRequestorIdNil

`func (o *MerchantAccount) SetMastercardNetworkTokensRequestorIdNil(b bool)`

 SetMastercardNetworkTokensRequestorIdNil sets the value for MastercardNetworkTokensRequestorId to be an explicit nil

### UnsetMastercardNetworkTokensRequestorId
`func (o *MerchantAccount) UnsetMastercardNetworkTokensRequestorId()`

UnsetMastercardNetworkTokensRequestorId ensures that no value is present for MastercardNetworkTokensRequestorId, not even an explicit nil
### GetMastercardNetworkTokensAppId

`func (o *MerchantAccount) GetMastercardNetworkTokensAppId() string`

GetMastercardNetworkTokensAppId returns the MastercardNetworkTokensAppId field if non-nil, zero value otherwise.

### GetMastercardNetworkTokensAppIdOk

`func (o *MerchantAccount) GetMastercardNetworkTokensAppIdOk() (*string, bool)`

GetMastercardNetworkTokensAppIdOk returns a tuple with the MastercardNetworkTokensAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMastercardNetworkTokensAppId

`func (o *MerchantAccount) SetMastercardNetworkTokensAppId(v string)`

SetMastercardNetworkTokensAppId sets MastercardNetworkTokensAppId field to given value.

### HasMastercardNetworkTokensAppId

`func (o *MerchantAccount) HasMastercardNetworkTokensAppId() bool`

HasMastercardNetworkTokensAppId returns a boolean if a field has been set.

### SetMastercardNetworkTokensAppIdNil

`func (o *MerchantAccount) SetMastercardNetworkTokensAppIdNil(b bool)`

 SetMastercardNetworkTokensAppIdNil sets the value for MastercardNetworkTokensAppId to be an explicit nil

### UnsetMastercardNetworkTokensAppId
`func (o *MerchantAccount) UnsetMastercardNetworkTokensAppId()`

UnsetMastercardNetworkTokensAppId ensures that no value is present for MastercardNetworkTokensAppId, not even an explicit nil
### GetCreatedAt

`func (o *MerchantAccount) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MerchantAccount) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MerchantAccount) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MerchantAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MerchantAccount) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MerchantAccount) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MerchantAccount) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MerchantAccount) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


