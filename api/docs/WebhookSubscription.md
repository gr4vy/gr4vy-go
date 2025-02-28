# WebhookSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;webhook-subscription&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique Gr4vy ID for this webhook subscription. | [optional] 
**MerchantAccountId** | Pointer to **NullableString** | The unique ID for a merchant account. | [optional] 
**Active** | Pointer to **bool** | Defines if this subscription is currently active or not. When deactivated, webhook events are not sent to the endpoint configured for this subscription. | [optional] 
**Url** | Pointer to **string** | The URL of the endpoint to deliver webhook events to. | [optional] 
**Authentication** | Pointer to [**NullableWebhookSubscriptionAuthentication**](WebhookSubscriptionAuthentication.md) |  | [optional] 

## Methods

### NewWebhookSubscription

`func NewWebhookSubscription() *WebhookSubscription`

NewWebhookSubscription instantiates a new WebhookSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSubscriptionWithDefaults

`func NewWebhookSubscriptionWithDefaults() *WebhookSubscription`

NewWebhookSubscriptionWithDefaults instantiates a new WebhookSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WebhookSubscription) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookSubscription) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookSubscription) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WebhookSubscription) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *WebhookSubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookSubscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookSubscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantAccountId

`func (o *WebhookSubscription) GetMerchantAccountId() string`

GetMerchantAccountId returns the MerchantAccountId field if non-nil, zero value otherwise.

### GetMerchantAccountIdOk

`func (o *WebhookSubscription) GetMerchantAccountIdOk() (*string, bool)`

GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountId

`func (o *WebhookSubscription) SetMerchantAccountId(v string)`

SetMerchantAccountId sets MerchantAccountId field to given value.

### HasMerchantAccountId

`func (o *WebhookSubscription) HasMerchantAccountId() bool`

HasMerchantAccountId returns a boolean if a field has been set.

### SetMerchantAccountIdNil

`func (o *WebhookSubscription) SetMerchantAccountIdNil(b bool)`

 SetMerchantAccountIdNil sets the value for MerchantAccountId to be an explicit nil

### UnsetMerchantAccountId
`func (o *WebhookSubscription) UnsetMerchantAccountId()`

UnsetMerchantAccountId ensures that no value is present for MerchantAccountId, not even an explicit nil
### GetActive

`func (o *WebhookSubscription) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WebhookSubscription) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WebhookSubscription) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WebhookSubscription) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookSubscription) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookSubscription) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookSubscription) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WebhookSubscription) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAuthentication

`func (o *WebhookSubscription) GetAuthentication() WebhookSubscriptionAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *WebhookSubscription) GetAuthenticationOk() (*WebhookSubscriptionAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *WebhookSubscription) SetAuthentication(v WebhookSubscriptionAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *WebhookSubscription) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### SetAuthenticationNil

`func (o *WebhookSubscription) SetAuthenticationNil(b bool)`

 SetAuthenticationNil sets the value for Authentication to be an explicit nil

### UnsetAuthentication
`func (o *WebhookSubscription) UnsetAuthentication()`

UnsetAuthentication ensures that no value is present for Authentication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


