# PaymentMethodToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;payment-method-token&#x60;. | [optional] 
**Id** | Pointer to **string** | The external ID of this payment method as it has been registered with the payment service, which can be used directly in combination with the &#x60;token&#x60; without the need to go through Gr4vy for a transaction.  In some cases this is a different value to the &#x60;token&#x60; while in others this value is identical. Please see the documentation for the payment service for more details. | [optional] 
**Token** | Pointer to **string** | The token of this payment method as it has been registered with the payment service, which can be used directly in combination with the &#x60;id&#x60; without the need to go through Gr4vy for a transaction.  In some cases this is a different value to the &#x60;id&#x60; while in others this value is identical. Please see the documentation for the payment service for more details. | [optional] 
**PaymentService** | Pointer to [**PaymentServiceSnapshot**](PaymentService--Snapshot.md) |  | [optional] 

## Methods

### NewPaymentMethodToken

`func NewPaymentMethodToken() *PaymentMethodToken`

NewPaymentMethodToken instantiates a new PaymentMethodToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodTokenWithDefaults

`func NewPaymentMethodTokenWithDefaults() *PaymentMethodToken`

NewPaymentMethodTokenWithDefaults instantiates a new PaymentMethodToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodToken) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodToken) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodToken) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentMethodToken) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *PaymentMethodToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentMethodToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetToken

`func (o *PaymentMethodToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentMethodToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentMethodToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PaymentMethodToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetPaymentService

`func (o *PaymentMethodToken) GetPaymentService() PaymentServiceSnapshot`

GetPaymentService returns the PaymentService field if non-nil, zero value otherwise.

### GetPaymentServiceOk

`func (o *PaymentMethodToken) GetPaymentServiceOk() (*PaymentServiceSnapshot, bool)`

GetPaymentServiceOk returns a tuple with the PaymentService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentService

`func (o *PaymentMethodToken) SetPaymentService(v PaymentServiceSnapshot)`

SetPaymentService sets PaymentService field to given value.

### HasPaymentService

`func (o *PaymentMethodToken) HasPaymentService() bool`

HasPaymentService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


