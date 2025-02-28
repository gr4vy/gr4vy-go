# WebhookSubscriptionAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Only &#x60;basic&#x60;&#x60; is currently supported. | [optional] 
**Username** | Pointer to **NullableString** | The basic authentication username to use for authentication. | [optional] 
**Password** | Pointer to **NullableString** | The basic authentication password to use for authentication. | [optional] 

## Methods

### NewWebhookSubscriptionAuthentication

`func NewWebhookSubscriptionAuthentication() *WebhookSubscriptionAuthentication`

NewWebhookSubscriptionAuthentication instantiates a new WebhookSubscriptionAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSubscriptionAuthenticationWithDefaults

`func NewWebhookSubscriptionAuthenticationWithDefaults() *WebhookSubscriptionAuthentication`

NewWebhookSubscriptionAuthenticationWithDefaults instantiates a new WebhookSubscriptionAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *WebhookSubscriptionAuthentication) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WebhookSubscriptionAuthentication) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WebhookSubscriptionAuthentication) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *WebhookSubscriptionAuthentication) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetUsername

`func (o *WebhookSubscriptionAuthentication) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *WebhookSubscriptionAuthentication) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *WebhookSubscriptionAuthentication) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *WebhookSubscriptionAuthentication) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *WebhookSubscriptionAuthentication) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *WebhookSubscriptionAuthentication) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *WebhookSubscriptionAuthentication) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *WebhookSubscriptionAuthentication) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *WebhookSubscriptionAuthentication) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *WebhookSubscriptionAuthentication) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *WebhookSubscriptionAuthentication) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *WebhookSubscriptionAuthentication) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


