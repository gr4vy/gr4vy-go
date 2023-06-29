# CheckoutSessionFieldsPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | &#x60;card&#x60;. | 
**Number** | Pointer to **string** | The 13-19 digit number for this card as it can be found on the front of the card. | [optional] 
**ExpirationDate** | Pointer to **string** | The expiration date of the card, formatted &#x60;MM/YY&#x60;. | [optional] 
**SecurityCode** | Pointer to **string** | The 3 or 4 digit security code often found on the card. This often referred to as the CVV or CVD.  The security code can only be set if the stored payment method represents a card. | [optional] 

## Methods

### NewCheckoutSessionFieldsPaymentMethod

`func NewCheckoutSessionFieldsPaymentMethod(method string, ) *CheckoutSessionFieldsPaymentMethod`

NewCheckoutSessionFieldsPaymentMethod instantiates a new CheckoutSessionFieldsPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutSessionFieldsPaymentMethodWithDefaults

`func NewCheckoutSessionFieldsPaymentMethodWithDefaults() *CheckoutSessionFieldsPaymentMethod`

NewCheckoutSessionFieldsPaymentMethodWithDefaults instantiates a new CheckoutSessionFieldsPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *CheckoutSessionFieldsPaymentMethod) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CheckoutSessionFieldsPaymentMethod) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CheckoutSessionFieldsPaymentMethod) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetNumber

`func (o *CheckoutSessionFieldsPaymentMethod) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CheckoutSessionFieldsPaymentMethod) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CheckoutSessionFieldsPaymentMethod) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CheckoutSessionFieldsPaymentMethod) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CheckoutSessionFieldsPaymentMethod) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CheckoutSessionFieldsPaymentMethod) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CheckoutSessionFieldsPaymentMethod) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CheckoutSessionFieldsPaymentMethod) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetSecurityCode

`func (o *CheckoutSessionFieldsPaymentMethod) GetSecurityCode() string`

GetSecurityCode returns the SecurityCode field if non-nil, zero value otherwise.

### GetSecurityCodeOk

`func (o *CheckoutSessionFieldsPaymentMethod) GetSecurityCodeOk() (*string, bool)`

GetSecurityCodeOk returns a tuple with the SecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityCode

`func (o *CheckoutSessionFieldsPaymentMethod) SetSecurityCode(v string)`

SetSecurityCode sets SecurityCode field to given value.

### HasSecurityCode

`func (o *CheckoutSessionFieldsPaymentMethod) HasSecurityCode() bool`

HasSecurityCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


