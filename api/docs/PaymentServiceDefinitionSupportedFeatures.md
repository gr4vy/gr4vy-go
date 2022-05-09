# PaymentServiceDefinitionSupportedFeatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethodTokenization** | Pointer to **bool** | Supports storing a payment method via tokenization. | [optional] 
**PaymentMethodTokenizationToggle** | Pointer to **bool** | Supports toggling tokenization for a payment method on or off from the dashboard. | [optional] 
**ThreeDSecureHosted** | Pointer to **bool** | Supports hosted 3-D Secure with a redirect. | [optional] 
**ThreeDSecurePassThrough** | Pointer to **bool** | Supports passing 3-D Secure data to the underlying processor. | [optional] 
**NetworkTokens** | Pointer to **bool** | Supports passing decrypted digital wallet (e.g. Apple Pay) tokens to the underlying processor. | [optional] 
**VerifyCredentials** | Pointer to **bool** | Supports verifying the credentials entered while setting up the underlying processor. This is for internal use only. | [optional] 
**Void** | Pointer to **bool** | Supports [voiding](#operation/void-transaction) authorized transactions. | [optional] 
**Refunds** | Pointer to **bool** | Supports [refunding](#operation/refund-transaction) captured transactions. | [optional] 
**PartialRefunds** | Pointer to **bool** | Supports [partially refunding](#operation/refund-transaction) captured transactions. | [optional] 

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

### GetPaymentMethodTokenizationToggle

`func (o *PaymentServiceDefinitionSupportedFeatures) GetPaymentMethodTokenizationToggle() bool`

GetPaymentMethodTokenizationToggle returns the PaymentMethodTokenizationToggle field if non-nil, zero value otherwise.

### GetPaymentMethodTokenizationToggleOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetPaymentMethodTokenizationToggleOk() (*bool, bool)`

GetPaymentMethodTokenizationToggleOk returns a tuple with the PaymentMethodTokenizationToggle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTokenizationToggle

`func (o *PaymentServiceDefinitionSupportedFeatures) SetPaymentMethodTokenizationToggle(v bool)`

SetPaymentMethodTokenizationToggle sets PaymentMethodTokenizationToggle field to given value.

### HasPaymentMethodTokenizationToggle

`func (o *PaymentServiceDefinitionSupportedFeatures) HasPaymentMethodTokenizationToggle() bool`

HasPaymentMethodTokenizationToggle returns a boolean if a field has been set.

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

### GetNetworkTokens

`func (o *PaymentServiceDefinitionSupportedFeatures) GetNetworkTokens() bool`

GetNetworkTokens returns the NetworkTokens field if non-nil, zero value otherwise.

### GetNetworkTokensOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetNetworkTokensOk() (*bool, bool)`

GetNetworkTokensOk returns a tuple with the NetworkTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTokens

`func (o *PaymentServiceDefinitionSupportedFeatures) SetNetworkTokens(v bool)`

SetNetworkTokens sets NetworkTokens field to given value.

### HasNetworkTokens

`func (o *PaymentServiceDefinitionSupportedFeatures) HasNetworkTokens() bool`

HasNetworkTokens returns a boolean if a field has been set.

### GetVerifyCredentials

`func (o *PaymentServiceDefinitionSupportedFeatures) GetVerifyCredentials() bool`

GetVerifyCredentials returns the VerifyCredentials field if non-nil, zero value otherwise.

### GetVerifyCredentialsOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetVerifyCredentialsOk() (*bool, bool)`

GetVerifyCredentialsOk returns a tuple with the VerifyCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCredentials

`func (o *PaymentServiceDefinitionSupportedFeatures) SetVerifyCredentials(v bool)`

SetVerifyCredentials sets VerifyCredentials field to given value.

### HasVerifyCredentials

`func (o *PaymentServiceDefinitionSupportedFeatures) HasVerifyCredentials() bool`

HasVerifyCredentials returns a boolean if a field has been set.

### GetVoid

`func (o *PaymentServiceDefinitionSupportedFeatures) GetVoid() bool`

GetVoid returns the Void field if non-nil, zero value otherwise.

### GetVoidOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetVoidOk() (*bool, bool)`

GetVoidOk returns a tuple with the Void field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoid

`func (o *PaymentServiceDefinitionSupportedFeatures) SetVoid(v bool)`

SetVoid sets Void field to given value.

### HasVoid

`func (o *PaymentServiceDefinitionSupportedFeatures) HasVoid() bool`

HasVoid returns a boolean if a field has been set.

### GetRefunds

`func (o *PaymentServiceDefinitionSupportedFeatures) GetRefunds() bool`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetRefundsOk() (*bool, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *PaymentServiceDefinitionSupportedFeatures) SetRefunds(v bool)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *PaymentServiceDefinitionSupportedFeatures) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.

### GetPartialRefunds

`func (o *PaymentServiceDefinitionSupportedFeatures) GetPartialRefunds() bool`

GetPartialRefunds returns the PartialRefunds field if non-nil, zero value otherwise.

### GetPartialRefundsOk

`func (o *PaymentServiceDefinitionSupportedFeatures) GetPartialRefundsOk() (*bool, bool)`

GetPartialRefundsOk returns a tuple with the PartialRefunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialRefunds

`func (o *PaymentServiceDefinitionSupportedFeatures) SetPartialRefunds(v bool)`

SetPartialRefunds sets PartialRefunds field to given value.

### HasPartialRefunds

`func (o *PaymentServiceDefinitionSupportedFeatures) HasPartialRefunds() bool`

HasPartialRefunds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


