# WebhookSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Defines if this subscription is currently active or not. When deactivated, webhook events are not sent to the endpoint configured for this subscription. | [optional] 
**Url** | Pointer to **string** | The URL of the endpoint to deliver webhook events to. | [optional] 
**Authentication** | Pointer to [**NullableWebhookSubscriptionAuthentication**](WebhookSubscriptionAuthentication.md) |  | [optional] 

## Methods

### NewWebhookSubscriptionRequest

`func NewWebhookSubscriptionRequest() *WebhookSubscriptionRequest`

NewWebhookSubscriptionRequest instantiates a new WebhookSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSubscriptionRequestWithDefaults

`func NewWebhookSubscriptionRequestWithDefaults() *WebhookSubscriptionRequest`

NewWebhookSubscriptionRequestWithDefaults instantiates a new WebhookSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *WebhookSubscriptionRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WebhookSubscriptionRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WebhookSubscriptionRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WebhookSubscriptionRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookSubscriptionRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookSubscriptionRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookSubscriptionRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WebhookSubscriptionRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAuthentication

`func (o *WebhookSubscriptionRequest) GetAuthentication() WebhookSubscriptionAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *WebhookSubscriptionRequest) GetAuthenticationOk() (*WebhookSubscriptionAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *WebhookSubscriptionRequest) SetAuthentication(v WebhookSubscriptionAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *WebhookSubscriptionRequest) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### SetAuthenticationNil

`func (o *WebhookSubscriptionRequest) SetAuthenticationNil(b bool)`

 SetAuthenticationNil sets the value for Authentication to be an explicit nil

### UnsetAuthentication
`func (o *WebhookSubscriptionRequest) UnsetAuthentication()`

UnsetAuthentication ensures that no value is present for Authentication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


