# PaymentMethodTokenized

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;payment-method&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique ID of the payment method. | [optional] 
**Method** | Pointer to **string** | The type of this payment method. | [optional] 
**Label** | Pointer to **string** | A label for the payment method. For a &#x60;card&#x60; payment method this is the last 4 digits on the card. For others it would be the email address. | [optional] 
**Scheme** | Pointer to **NullableString** | The type of the card, if the payment method is a card. | [optional] 
**ExpirationDate** | Pointer to **NullableString** | The expiration date for the payment method. | [optional] 

## Methods

### NewPaymentMethodTokenized

`func NewPaymentMethodTokenized() *PaymentMethodTokenized`

NewPaymentMethodTokenized instantiates a new PaymentMethodTokenized object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodTokenizedWithDefaults

`func NewPaymentMethodTokenizedWithDefaults() *PaymentMethodTokenized`

NewPaymentMethodTokenizedWithDefaults instantiates a new PaymentMethodTokenized object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodTokenized) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodTokenized) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodTokenized) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentMethodTokenized) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *PaymentMethodTokenized) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodTokenized) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodTokenized) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentMethodTokenized) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethod

`func (o *PaymentMethodTokenized) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PaymentMethodTokenized) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PaymentMethodTokenized) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PaymentMethodTokenized) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetLabel

`func (o *PaymentMethodTokenized) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PaymentMethodTokenized) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PaymentMethodTokenized) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PaymentMethodTokenized) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetScheme

`func (o *PaymentMethodTokenized) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *PaymentMethodTokenized) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *PaymentMethodTokenized) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *PaymentMethodTokenized) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *PaymentMethodTokenized) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *PaymentMethodTokenized) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetExpirationDate

`func (o *PaymentMethodTokenized) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *PaymentMethodTokenized) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *PaymentMethodTokenized) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *PaymentMethodTokenized) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *PaymentMethodTokenized) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *PaymentMethodTokenized) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

