# PaymentServiceDefinitionSupportedFeatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethodTokenization** | Pointer to **bool** | Supports storing a payment method via tokenization. | [optional] 
**ThreeDSecureHosted** | Pointer to **bool** | Supports hosted 3-D Secure with a redirect. | [optional] 
**ThreeDSecurePassThrough** | Pointer to **bool** | Supports passing 3-D Secure data to the underlying processor. | [optional] 
**ApplePay** | Pointer to **bool** | Supports passing decrypted apple pay token to the underlying processor. | [optional] 

## Methods

### NewPaymentServiceDefinitionSupportedFeatures

`func NewPaymentServiceDefinitionSupportedFeatures() *PaymentServiceDefinitionSupportedFeatures`

NewPaymentServiceDefinitionSupportedFeatures instantiates a new PaymentServiceDefinitionSupportedFeatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceDefinitionSupportedFeaturesWithDefaults

`func NewPaymentServiceDefinitionSupportedFeaturesWithDefaults() *PaymentServiceDefinitionSupportedFeatures`

NewPaymentServiceDefinitionSupportedFeaturesWithDefaults instantiates a new PaymentServiceDefinitionSupportedFeatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethodTokenization

`func (o *PaymentServiceDefinitionSupportedFeatures) GetPaymentMethodTokenization() bool`

GetPaymentMethodTokenization returns the PaymentMethodTokenization field if non-nil, zero value otherwise.

### GetPaymentMethodTokenizationOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetPaymentMethodTokenizationOk() (*bool, bool)`

GetPaymentMethodTokenizationOk returns a tuple with the PaymentMethodTokenization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTokenization

`func (o *PaymentServiceDefinitionSupportedFeatures) SetPaymentMethodTokenization(v bool)`

SetPaymentMethodTokenization sets PaymentMethodTokenization field to given value.

### HasPaymentMethodTokenization

`func (o *PaymentServiceDefinitionSupportedFeatures) HasPaymentMethodTokenization() bool`

HasPaymentMethodTokenization returns a boolean if a field has been set.

### GetThreeDSecureHosted

`func (o *PaymentServiceDefinitionSupportedFeatures) GetThreeDSecureHosted() bool`

GetThreeDSecureHosted returns the ThreeDSecureHosted field if non-nil, zero value otherwise.

### GetThreeDSecureHostedOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetThreeDSecureHostedOk() (*bool, bool)`

GetThreeDSecureHostedOk returns a tuple with the ThreeDSecureHosted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSecureHosted

`func (o *PaymentServiceDefinitionSupportedFeatures) SetThreeDSecureHosted(v bool)`

SetThreeDSecureHosted sets ThreeDSecureHosted field to given value.

### HasThreeDSecureHosted

`func (o *PaymentServiceDefinitionSupportedFeatures) HasThreeDSecureHosted() bool`

HasThreeDSecureHosted returns a boolean if a field has been set.

### GetThreeDSecurePassThrough

`func (o *PaymentServiceDefinitionSupportedFeatures) GetThreeDSecurePassThrough() bool`

GetThreeDSecurePassThrough returns the ThreeDSecurePassThrough field if non-nil, zero value otherwise.

### GetThreeDSecurePassThroughOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetThreeDSecurePassThroughOk() (*bool, bool)`

GetThreeDSecurePassThroughOk returns a tuple with the ThreeDSecurePassThrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSecurePassThrough

`func (o *PaymentServiceDefinitionSupportedFeatures) SetThreeDSecurePassThrough(v bool)`

SetThreeDSecurePassThrough sets ThreeDSecurePassThrough field to given value.

### HasThreeDSecurePassThrough

`func (o *PaymentServiceDefinitionSupportedFeatures) HasThreeDSecurePassThrough() bool`

HasThreeDSecurePassThrough returns a boolean if a field has been set.

### GetApplePay

`func (o *PaymentServiceDefinitionSupportedFeatures) GetApplePay() bool`

GetApplePay returns the ApplePay field if non-nil, zero value otherwise.

### GetApplePayOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetApplePayOk() (*bool, bool)`

GetApplePayOk returns a tuple with the ApplePay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplePay

`func (o *PaymentServiceDefinitionSupportedFeatures) SetApplePay(v bool)`

SetApplePay sets ApplePay field to given value.

### HasApplePay

`func (o *PaymentServiceDefinitionSupportedFeatures) HasApplePay() bool`

HasApplePay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


