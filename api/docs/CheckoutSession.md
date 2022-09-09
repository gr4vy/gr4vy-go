# CheckoutSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;checkout-session&#x60;. | [optional] 
**Id** | Pointer to **string** | The ID of the Checkout Session. | [optional] 
**ExpiresAt** | Pointer to **string** | The date and time when the Checkout Session will expire. By default this will be set to 1 hour from the date of creation. | [optional] 

## Methods

### NewCheckoutSession

`func NewCheckoutSession() *CheckoutSession`

NewCheckoutSession instantiates a new CheckoutSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutSessionWithDefaults

`func NewCheckoutSessionWithDefaults() *CheckoutSession`

NewCheckoutSessionWithDefaults instantiates a new CheckoutSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CheckoutSession) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckoutSession) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckoutSession) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CheckoutSession) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *CheckoutSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckoutSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckoutSession) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CheckoutSession) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CheckoutSession) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CheckoutSession) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CheckoutSession) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CheckoutSession) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


