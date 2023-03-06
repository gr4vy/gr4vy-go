# PaymentServiceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | A custom name for the payment service. This will be shown in the Admin UI. | [optional] 
**Fields** | Pointer to [**[]PaymentServiceRequestFields**](PaymentServiceRequestFields.md) | A list of fields, each containing a key-value pair for each field defined by the definition for this payment service e.g. for stripe-card &#x60;secret_key&#x60; is required and so must be sent within this field. | [optional] 
**AcceptedCountries** | Pointer to **[]string** | A list of countries that this payment service needs to support in ISO two-letter code format. | [optional] 
**AcceptedCurrencies** | Pointer to **[]string** | A list of currencies that this payment service needs to support in ISO 4217 three-letter code format. | [optional] 
**ThreeDSecureEnabled** | Pointer to **bool** | Defines if 3-D Secure is enabled for the service (can only be enabled if the payment service definition supports the &#x60;three_d_secure_hosted&#x60; feature). This does not affect pass through 3-D Secure data. | [optional] [default to false]
**AcquirerBinVisa** | Pointer to **NullableString** |  | [optional] 
**AcquirerBinMastercard** | Pointer to **NullableString** |  | [optional] 
**AcquirerBinAmex** | Pointer to **NullableString** |  | [optional] 
**AcquirerBinDiscover** | Pointer to **NullableString** |  | [optional] 
**AcquirerMerchantId** | Pointer to **NullableString** |  | [optional] 
**MerchantName** | Pointer to **NullableString** |  | [optional] 
**MerchantCountryCode** | Pointer to **NullableString** |  | [optional] 
**MerchantCategoryCode** | Pointer to **NullableString** |  | [optional] 
**MerchantProfile** | Pointer to [**NullableMerchantProfile**](MerchantProfile.md) | Configuration for each supported card scheme. When updating a Payment Service, a key not being present will indicate no updates to be done on that scheme, whereas an object being sent as Null for a key will empty the configuration for that scheme. | [optional] 
**MerchantUrl** | Pointer to **NullableString** |  | [optional] 
**Active** | Pointer to **bool** | Defines if this service is currently active or not. | [optional] [default to true]
**Position** | Pointer to **float32** | The numeric rank of a payment service. Payment services with a lower position value are processed first. When a payment services is inserted at a position, any payment services with the the same value or higher are shifted down a position accordingly. When left out, the payment service is inserted at the end of the list. | [optional] 
**OpenLoop** | Pointer to **NullableBool** | Defines if the service works as an open-loop service. This feature can only be enabled if the PSP is set up to accept previous scheme transaction IDs.  If this value is set to &#x60;null&#x60;, it will be set to the value of &#x60;open_loop&#x60; in the payment service definition.  If &#x60;open_loop_toggle&#x60; is &#x60;false&#x60; in the payment service definition, &#x60;open_loop&#x60; should either not be provided or set to &#x60;null&#x60;, or it will fail with a validation error. | [optional] 
**PaymentMethodTokenizationEnabled** | Pointer to **bool** | Defines if tokenization is enabled for the service. This feature can only be enabled if the payment service is NOT set as &#x60;open_loop&#x60; and the PSP is set up to tokenize. | [optional] [default to false]
**NetworkTokensEnabled** | Pointer to **NullableBool** | Defines if network tokens are enabled for the service. This feature can only be enabled if the payment service is set as &#x60;open_loop&#x60; and the PSP is set up to accept network tokens.  If this value is set to &#x60;null&#x60;, it will be set to the value of &#x60;network_tokens_default&#x60; in the payment service definition.  If &#x60;network_tokens_toggle&#x60; is &#x60;false&#x60; in the payment service definition, &#x60;network_tokens_enabled&#x60; should either not be provided or set to &#x60;null&#x60;, or it will fail with a validation error. | [optional] 

## Methods

### NewPaymentServiceUpdate

`func NewPaymentServiceUpdate() *PaymentServiceUpdate`

NewPaymentServiceUpdate instantiates a new PaymentServiceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceUpdateWithDefaults

`func NewPaymentServiceUpdateWithDefaults() *PaymentServiceUpdate`

NewPaymentServiceUpdateWithDefaults instantiates a new PaymentServiceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *PaymentServiceUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PaymentServiceUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PaymentServiceUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PaymentServiceUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFields

`func (o *PaymentServiceUpdate) GetFields() []PaymentServiceRequestFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PaymentServiceUpdate) GetFieldsOk() (*[]PaymentServiceRequestFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PaymentServiceUpdate) SetFields(v []PaymentServiceRequestFields)`

SetFields sets Fields field to given value.

### HasFields

`func (o *PaymentServiceUpdate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetAcceptedCountries

`func (o *PaymentServiceUpdate) GetAcceptedCountries() []string`

GetAcceptedCountries returns the AcceptedCountries field if non-nil, zero value otherwise.

### GetAcceptedCountriesOk

`func (o *PaymentServiceUpdate) GetAcceptedCountriesOk() (*[]string, bool)`

GetAcceptedCountriesOk returns a tuple with the AcceptedCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCountries

`func (o *PaymentServiceUpdate) SetAcceptedCountries(v []string)`

SetAcceptedCountries sets AcceptedCountries field to given value.

### HasAcceptedCountries

`func (o *PaymentServiceUpdate) HasAcceptedCountries() bool`

HasAcceptedCountries returns a boolean if a field has been set.

### GetAcceptedCurrencies

`func (o *PaymentServiceUpdate) GetAcceptedCurrencies() []string`

GetAcceptedCurrencies returns the AcceptedCurrencies field if non-nil, zero value otherwise.

### GetAcceptedCurrenciesOk

`func (o *PaymentServiceUpdate) GetAcceptedCurrenciesOk() (*[]string, bool)`

GetAcceptedCurrenciesOk returns a tuple with the AcceptedCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCurrencies

`func (o *PaymentServiceUpdate) SetAcceptedCurrencies(v []string)`

SetAcceptedCurrencies sets AcceptedCurrencies field to given value.

### HasAcceptedCurrencies

`func (o *PaymentServiceUpdate) HasAcceptedCurrencies() bool`

HasAcceptedCurrencies returns a boolean if a field has been set.

### GetThreeDSecureEnabled

`func (o *PaymentServiceUpdate) GetThreeDSecureEnabled() bool`

GetThreeDSecureEnabled returns the ThreeDSecureEnabled field if non-nil, zero value otherwise.

### GetThreeDSecureEnabledOk

`func (o *PaymentServiceUpdate) GetThreeDSecureEnabledOk() (*bool, bool)`

GetThreeDSecureEnabledOk returns a tuple with the ThreeDSecureEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSecureEnabled

`func (o *PaymentServiceUpdate) SetThreeDSecureEnabled(v bool)`

SetThreeDSecureEnabled sets ThreeDSecureEnabled field to given value.

### HasThreeDSecureEnabled

`func (o *PaymentServiceUpdate) HasThreeDSecureEnabled() bool`

HasThreeDSecureEnabled returns a boolean if a field has been set.

### GetAcquirerBinVisa

`func (o *PaymentServiceUpdate) GetAcquirerBinVisa() string`

GetAcquirerBinVisa returns the AcquirerBinVisa field if non-nil, zero value otherwise.

### GetAcquirerBinVisaOk

`func (o *PaymentServiceUpdate) GetAcquirerBinVisaOk() (*string, bool)`

GetAcquirerBinVisaOk returns a tuple with the AcquirerBinVisa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerBinVisa

`func (o *PaymentServiceUpdate) SetAcquirerBinVisa(v string)`

SetAcquirerBinVisa sets AcquirerBinVisa field to given value.

### HasAcquirerBinVisa

`func (o *PaymentServiceUpdate) HasAcquirerBinVisa() bool`

HasAcquirerBinVisa returns a boolean if a field has been set.

### SetAcquirerBinVisaNil

`func (o *PaymentServiceUpdate) SetAcquirerBinVisaNil(b bool)`

 SetAcquirerBinVisaNil sets the value for AcquirerBinVisa to be an explicit nil

### UnsetAcquirerBinVisa
`func (o *PaymentServiceUpdate) UnsetAcquirerBinVisa()`

UnsetAcquirerBinVisa ensures that no value is present for AcquirerBinVisa, not even an explicit nil
### GetAcquirerBinMastercard

`func (o *PaymentServiceUpdate) GetAcquirerBinMastercard() string`

GetAcquirerBinMastercard returns the AcquirerBinMastercard field if non-nil, zero value otherwise.

### GetAcquirerBinMastercardOk

`func (o *PaymentServiceUpdate) GetAcquirerBinMastercardOk() (*string, bool)`

GetAcquirerBinMastercardOk returns a tuple with the AcquirerBinMastercard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerBinMastercard

`func (o *PaymentServiceUpdate) SetAcquirerBinMastercard(v string)`

SetAcquirerBinMastercard sets AcquirerBinMastercard field to given value.

### HasAcquirerBinMastercard

`func (o *PaymentServiceUpdate) HasAcquirerBinMastercard() bool`

HasAcquirerBinMastercard returns a boolean if a field has been set.

### SetAcquirerBinMastercardNil

`func (o *PaymentServiceUpdate) SetAcquirerBinMastercardNil(b bool)`

 SetAcquirerBinMastercardNil sets the value for AcquirerBinMastercard to be an explicit nil

### UnsetAcquirerBinMastercard
`func (o *PaymentServiceUpdate) UnsetAcquirerBinMastercard()`

UnsetAcquirerBinMastercard ensures that no value is present for AcquirerBinMastercard, not even an explicit nil
### GetAcquirerBinAmex

`func (o *PaymentServiceUpdate) GetAcquirerBinAmex() string`

GetAcquirerBinAmex returns the AcquirerBinAmex field if non-nil, zero value otherwise.

### GetAcquirerBinAmexOk

`func (o *PaymentServiceUpdate) GetAcquirerBinAmexOk() (*string, bool)`

GetAcquirerBinAmexOk returns a tuple with the AcquirerBinAmex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerBinAmex

`func (o *PaymentServiceUpdate) SetAcquirerBinAmex(v string)`

SetAcquirerBinAmex sets AcquirerBinAmex field to given value.

### HasAcquirerBinAmex

`func (o *PaymentServiceUpdate) HasAcquirerBinAmex() bool`

HasAcquirerBinAmex returns a boolean if a field has been set.

### SetAcquirerBinAmexNil

`func (o *PaymentServiceUpdate) SetAcquirerBinAmexNil(b bool)`

 SetAcquirerBinAmexNil sets the value for AcquirerBinAmex to be an explicit nil

### UnsetAcquirerBinAmex
`func (o *PaymentServiceUpdate) UnsetAcquirerBinAmex()`

UnsetAcquirerBinAmex ensures that no value is present for AcquirerBinAmex, not even an explicit nil
### GetAcquirerBinDiscover

`func (o *PaymentServiceUpdate) GetAcquirerBinDiscover() string`

GetAcquirerBinDiscover returns the AcquirerBinDiscover field if non-nil, zero value otherwise.

### GetAcquirerBinDiscoverOk

`func (o *PaymentServiceUpdate) GetAcquirerBinDiscoverOk() (*string, bool)`

GetAcquirerBinDiscoverOk returns a tuple with the AcquirerBinDiscover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerBinDiscover

`func (o *PaymentServiceUpdate) SetAcquirerBinDiscover(v string)`

SetAcquirerBinDiscover sets AcquirerBinDiscover field to given value.

### HasAcquirerBinDiscover

`func (o *PaymentServiceUpdate) HasAcquirerBinDiscover() bool`

HasAcquirerBinDiscover returns a boolean if a field has been set.

### SetAcquirerBinDiscoverNil

`func (o *PaymentServiceUpdate) SetAcquirerBinDiscoverNil(b bool)`

 SetAcquirerBinDiscoverNil sets the value for AcquirerBinDiscover to be an explicit nil

### UnsetAcquirerBinDiscover
`func (o *PaymentServiceUpdate) UnsetAcquirerBinDiscover()`

UnsetAcquirerBinDiscover ensures that no value is present for AcquirerBinDiscover, not even an explicit nil
### GetAcquirerMerchantId

`func (o *PaymentServiceUpdate) GetAcquirerMerchantId() string`

GetAcquirerMerchantId returns the AcquirerMerchantId field if non-nil, zero value otherwise.

### GetAcquirerMerchantIdOk

`func (o *PaymentServiceUpdate) GetAcquirerMerchantIdOk() (*string, bool)`

GetAcquirerMerchantIdOk returns a tuple with the AcquirerMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerMerchantId

`func (o *PaymentServiceUpdate) SetAcquirerMerchantId(v string)`

SetAcquirerMerchantId sets AcquirerMerchantId field to given value.

### HasAcquirerMerchantId

`func (o *PaymentServiceUpdate) HasAcquirerMerchantId() bool`

HasAcquirerMerchantId returns a boolean if a field has been set.

### SetAcquirerMerchantIdNil

`func (o *PaymentServiceUpdate) SetAcquirerMerchantIdNil(b bool)`

 SetAcquirerMerchantIdNil sets the value for AcquirerMerchantId to be an explicit nil

### UnsetAcquirerMerchantId
`func (o *PaymentServiceUpdate) UnsetAcquirerMerchantId()`

UnsetAcquirerMerchantId ensures that no value is present for AcquirerMerchantId, not even an explicit nil
### GetMerchantName

`func (o *PaymentServiceUpdate) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *PaymentServiceUpdate) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *PaymentServiceUpdate) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *PaymentServiceUpdate) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### SetMerchantNameNil

`func (o *PaymentServiceUpdate) SetMerchantNameNil(b bool)`

 SetMerchantNameNil sets the value for MerchantName to be an explicit nil

### UnsetMerchantName
`func (o *PaymentServiceUpdate) UnsetMerchantName()`

UnsetMerchantName ensures that no value is present for MerchantName, not even an explicit nil
### GetMerchantCountryCode

`func (o *PaymentServiceUpdate) GetMerchantCountryCode() string`

GetMerchantCountryCode returns the MerchantCountryCode field if non-nil, zero value otherwise.

### GetMerchantCountryCodeOk

`func (o *PaymentServiceUpdate) GetMerchantCountryCodeOk() (*string, bool)`

GetMerchantCountryCodeOk returns a tuple with the MerchantCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCountryCode

`func (o *PaymentServiceUpdate) SetMerchantCountryCode(v string)`

SetMerchantCountryCode sets MerchantCountryCode field to given value.

### HasMerchantCountryCode

`func (o *PaymentServiceUpdate) HasMerchantCountryCode() bool`

HasMerchantCountryCode returns a boolean if a field has been set.

### SetMerchantCountryCodeNil

`func (o *PaymentServiceUpdate) SetMerchantCountryCodeNil(b bool)`

 SetMerchantCountryCodeNil sets the value for MerchantCountryCode to be an explicit nil

### UnsetMerchantCountryCode
`func (o *PaymentServiceUpdate) UnsetMerchantCountryCode()`

UnsetMerchantCountryCode ensures that no value is present for MerchantCountryCode, not even an explicit nil
### GetMerchantCategoryCode

`func (o *PaymentServiceUpdate) GetMerchantCategoryCode() string`

GetMerchantCategoryCode returns the MerchantCategoryCode field if non-nil, zero value otherwise.

### GetMerchantCategoryCodeOk

`func (o *PaymentServiceUpdate) GetMerchantCategoryCodeOk() (*string, bool)`

GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCode

`func (o *PaymentServiceUpdate) SetMerchantCategoryCode(v string)`

SetMerchantCategoryCode sets MerchantCategoryCode field to given value.

### HasMerchantCategoryCode

`func (o *PaymentServiceUpdate) HasMerchantCategoryCode() bool`

HasMerchantCategoryCode returns a boolean if a field has been set.

### SetMerchantCategoryCodeNil

`func (o *PaymentServiceUpdate) SetMerchantCategoryCodeNil(b bool)`

 SetMerchantCategoryCodeNil sets the value for MerchantCategoryCode to be an explicit nil

### UnsetMerchantCategoryCode
`func (o *PaymentServiceUpdate) UnsetMerchantCategoryCode()`

UnsetMerchantCategoryCode ensures that no value is present for MerchantCategoryCode, not even an explicit nil
### GetMerchantProfile

`func (o *PaymentServiceUpdate) GetMerchantProfile() MerchantProfile`

GetMerchantProfile returns the MerchantProfile field if non-nil, zero value otherwise.

### GetMerchantProfileOk

`func (o *PaymentServiceUpdate) GetMerchantProfileOk() (*MerchantProfile, bool)`

GetMerchantProfileOk returns a tuple with the MerchantProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantProfile

`func (o *PaymentServiceUpdate) SetMerchantProfile(v MerchantProfile)`

SetMerchantProfile sets MerchantProfile field to given value.

### HasMerchantProfile

`func (o *PaymentServiceUpdate) HasMerchantProfile() bool`

HasMerchantProfile returns a boolean if a field has been set.

### SetMerchantProfileNil

`func (o *PaymentServiceUpdate) SetMerchantProfileNil(b bool)`

 SetMerchantProfileNil sets the value for MerchantProfile to be an explicit nil

### UnsetMerchantProfile
`func (o *PaymentServiceUpdate) UnsetMerchantProfile()`

UnsetMerchantProfile ensures that no value is present for MerchantProfile, not even an explicit nil
### GetMerchantUrl

`func (o *PaymentServiceUpdate) GetMerchantUrl() string`

GetMerchantUrl returns the MerchantUrl field if non-nil, zero value otherwise.

### GetMerchantUrlOk

`func (o *PaymentServiceUpdate) GetMerchantUrlOk() (*string, bool)`

GetMerchantUrlOk returns a tuple with the MerchantUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantUrl

`func (o *PaymentServiceUpdate) SetMerchantUrl(v string)`

SetMerchantUrl sets MerchantUrl field to given value.

### HasMerchantUrl

`func (o *PaymentServiceUpdate) HasMerchantUrl() bool`

HasMerchantUrl returns a boolean if a field has been set.

### SetMerchantUrlNil

`func (o *PaymentServiceUpdate) SetMerchantUrlNil(b bool)`

 SetMerchantUrlNil sets the value for MerchantUrl to be an explicit nil

### UnsetMerchantUrl
`func (o *PaymentServiceUpdate) UnsetMerchantUrl()`

UnsetMerchantUrl ensures that no value is present for MerchantUrl, not even an explicit nil
### GetActive

`func (o *PaymentServiceUpdate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PaymentServiceUpdate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PaymentServiceUpdate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PaymentServiceUpdate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPosition

`func (o *PaymentServiceUpdate) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PaymentServiceUpdate) GetPositionOk() (*float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PaymentServiceUpdate) SetPosition(v float32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PaymentServiceUpdate) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetOpenLoop

`func (o *PaymentServiceUpdate) GetOpenLoop() bool`

GetOpenLoop returns the OpenLoop field if non-nil, zero value otherwise.

### GetOpenLoopOk

`func (o *PaymentServiceUpdate) GetOpenLoopOk() (*bool, bool)`

GetOpenLoopOk returns a tuple with the OpenLoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenLoop

`func (o *PaymentServiceUpdate) SetOpenLoop(v bool)`

SetOpenLoop sets OpenLoop field to given value.

### HasOpenLoop

`func (o *PaymentServiceUpdate) HasOpenLoop() bool`

HasOpenLoop returns a boolean if a field has been set.

### SetOpenLoopNil

`func (o *PaymentServiceUpdate) SetOpenLoopNil(b bool)`

 SetOpenLoopNil sets the value for OpenLoop to be an explicit nil

### UnsetOpenLoop
`func (o *PaymentServiceUpdate) UnsetOpenLoop()`

UnsetOpenLoop ensures that no value is present for OpenLoop, not even an explicit nil
### GetPaymentMethodTokenizationEnabled

`func (o *PaymentServiceUpdate) GetPaymentMethodTokenizationEnabled() bool`

GetPaymentMethodTokenizationEnabled returns the PaymentMethodTokenizationEnabled field if non-nil, zero value otherwise.

### GetPaymentMethodTokenizationEnabledOk

`func (o *PaymentServiceUpdate) GetPaymentMethodTokenizationEnabledOk() (*bool, bool)`

GetPaymentMethodTokenizationEnabledOk returns a tuple with the PaymentMethodTokenizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTokenizationEnabled

`func (o *PaymentServiceUpdate) SetPaymentMethodTokenizationEnabled(v bool)`

SetPaymentMethodTokenizationEnabled sets PaymentMethodTokenizationEnabled field to given value.

### HasPaymentMethodTokenizationEnabled

`func (o *PaymentServiceUpdate) HasPaymentMethodTokenizationEnabled() bool`

HasPaymentMethodTokenizationEnabled returns a boolean if a field has been set.

### GetNetworkTokensEnabled

`func (o *PaymentServiceUpdate) GetNetworkTokensEnabled() bool`

GetNetworkTokensEnabled returns the NetworkTokensEnabled field if non-nil, zero value otherwise.

### GetNetworkTokensEnabledOk

`func (o *PaymentServiceUpdate) GetNetworkTokensEnabledOk() (*bool, bool)`

GetNetworkTokensEnabledOk returns a tuple with the NetworkTokensEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTokensEnabled

`func (o *PaymentServiceUpdate) SetNetworkTokensEnabled(v bool)`

SetNetworkTokensEnabled sets NetworkTokensEnabled field to given value.

### HasNetworkTokensEnabled

`func (o *PaymentServiceUpdate) HasNetworkTokensEnabled() bool`

HasNetworkTokensEnabled returns a boolean if a field has been set.

### SetNetworkTokensEnabledNil

`func (o *PaymentServiceUpdate) SetNetworkTokensEnabledNil(b bool)`

 SetNetworkTokensEnabledNil sets the value for NetworkTokensEnabled to be an explicit nil

### UnsetNetworkTokensEnabled
`func (o *PaymentServiceUpdate) UnsetNetworkTokensEnabled()`

UnsetNetworkTokensEnabled ensures that no value is present for NetworkTokensEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


