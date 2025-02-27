# WebhookSubscriptionUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **NullableBool** | Defines if this subscription is currently active or not. When deactivated, webhook events are not sent to the endpoint configured for this subscription. | [optional] 
**Url** | Pointer to **NullableString** | The URL of the endpoint to deliver webhook events to. | [optional] 
**Authentication** | Pointer to [**NullableWebhookSubscriptionAuthentication**](WebhookSubscriptionAuthentication.md) |  | [optional] 

## Methods

### NewWebhookSubscriptionUpdateRequest

`func NewWebhookSubscriptionUpdateRequest() *WebhookSubscriptionUpdateRequest`

NewWebhookSubscriptionUpdateRequest instantiates a new WebhookSubscriptionUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSubscriptionUpdateRequestWithDefaults

`func NewWebhookSubscriptionUpdateRequestWithDefaults() *WebhookSubscriptionUpdateRequest`

NewWebhookSubscriptionUpdateRequestWithDefaults instantiates a new WebhookSubscriptionUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *WebhookSubscriptionUpdateRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WebhookSubscriptionUpdateRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WebhookSubscriptionUpdateRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WebhookSubscriptionUpdateRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *WebhookSubscriptionUpdateRequest) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *WebhookSubscriptionUpdateRequest) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetUrl

`func (o *WebhookSubscriptionUpdateRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookSubscriptionUpdateRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookSubscriptionUpdateRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WebhookSubscriptionUpdateRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *WebhookSubscriptionUpdateRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *WebhookSubscriptionUpdateRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetAuthentication

`func (o *WebhookSubscriptionUpdateRequest) GetAuthentication() WebhookSubscriptionAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *WebhookSubscriptionUpdateRequest) GetAuthenticationOk() (*WebhookSubscriptionAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *WebhookSubscriptionUpdateRequest) SetAuthentication(v WebhookSubscriptionAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *WebhookSubscriptionUpdateRequest) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### SetAuthenticationNil

`func (o *WebhookSubscriptionUpdateRequest) SetAuthenticationNil(b bool)`

 SetAuthenticationNil sets the value for Authentication to be an explicit nil

### UnsetAuthentication
`func (o *WebhookSubscriptionUpdateRequest) UnsetAuthentication()`

UnsetAuthentication ensures that no value is present for Authentication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


