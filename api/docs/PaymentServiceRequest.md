# PaymentServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | A custom name for the payment service. This will be shown in the Admin UI. | 
**Fields** | [**[]PaymentServiceUpdateFields**](PaymentServiceUpdateFields.md) | A list of fields, each containing a key-value pair for each field defined by the definition for this payment service e.g. for stripe-card &#x60;secret_key&#x60; is required and so must be sent within this field. | 
**AcceptedCountries** | **[]string** | A list of countries that this payment service needs to support in ISO two-letter code format. | 
**AcceptedCurrencies** | **[]string** | A list of currencies that this payment service needs to support in ISO 4217 three-letter code format. | 
**ThreeDSecureEnabled** | Pointer to **bool** | Defines if 3-D Secure is enabled for the service (can only be enabled if the payment service definition supports the &#x60;three_d_secure_hosted&#x60; feature). This does not affect pass through 3-D Secure data. | [optional] [default to false]
**AcquirerBinVisa** | Pointer to **NullableString** | Acquiring institution identification code for VISA. | [optional] 
**AcquirerBinMastercard** | Pointer to **NullableString** | Acquiring institution identification code for Mastercard. | [optional] 
**AcquirerBinAmex** | Pointer to **NullableString** | Acquiring institution identification code for Amex. | [optional] 
**AcquirerBinDiscover** | Pointer to **NullableString** | Acquiring institution identification code for Discover. | [optional] 
**AcquirerMerchantId** | Pointer to **NullableString** | Merchant identifier used in authorisation requests (assigned by the acquirer). | [optional] 
**MerchantName** | Pointer to **NullableString** | Merchant name (assigned by the acquirer). | [optional] 
**MerchantCountryCode** | Pointer to **NullableString** | ISO 3166-1 numeric three-digit country code. | [optional] 
**MerchantCategoryCode** | Pointer to **NullableString** | Merchant category code that describes the business. | [optional] 
**MerchantUrl** | Pointer to **NullableString** | Fully qualified URL of 3-D Secure requestor website or customer care site. | [optional] 
**Active** | Pointer to **bool** | Defines if this service is currently active or not. | [optional] [default to true]
**Position** | Pointer to **float32** | The numeric rank of a payment service. Payment services with a lower position value are processed first. When a payment services is inserted at a position, any payment services with the the same value or higher are shifted down a position accordingly. When left out, the payment service is inserted at the end of the list. | [optional] 
**PaymentMethodTokenizationEnabled** | Pointer to **bool** | Defines if tokenization is enabled for the service (can only be enabled if the payment service definition supports it). | [optional] [default to false]
**PaymentServiceDefinitionId** | **string** | The ID of the payment service to use. | 

## Methods

### NewPaymentServiceRequest

`func NewPaymentServiceRequest(displayName string, fields []PaymentServiceUpdateFields, acceptedCountries []string, acceptedCurrencies []string, paymentServiceDefinitionId string, ) *PaymentServiceRequest`

NewPaymentServiceRequest instantiates a new PaymentServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceRequestWithDefaults

`func NewPaymentServiceRequestWithDefaults() *PaymentServiceRequest`

NewPaymentServiceRequestWithDefaults instantiates a new PaymentServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *PaymentServiceRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PaymentServiceRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PaymentServiceRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetFields

`func (o *PaymentServiceRequest) GetFields() []PaymentServiceUpdateFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PaymentServiceRequest) GetFieldsOk() (*[]PaymentServiceUpdateFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PaymentServiceRequest) SetFields(v []PaymentServiceUpdateFields)`

SetFields sets Fields field to given value.


### GetAcceptedCountries

`func (o *PaymentServiceRequest) GetAcceptedCountries() []string`

GetAcceptedCountries returns the AcceptedCountries field if non-nil, zero value otherwise.

### GetAcceptedCountriesOk

`func (o *PaymentServiceRequest) GetAcceptedCountriesOk() (*[]string, bool)`

GetAcceptedCountriesOk returns a tuple with the AcceptedCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCountries

`func (o *PaymentServiceRequest) SetAcceptedCountries(v []string)`

SetAcceptedCountries sets AcceptedCountries field to given value.


### GetAcceptedCurrencies

`func (o *PaymentServiceRequest) GetAcceptedCurrencies() []string`

GetAcceptedCurrencies returns the AcceptedCurrencies field if non-nil, zero value otherwise.

### GetAcceptedCurrenciesOk

`func (o *PaymentServiceRequest) GetAcceptedCurrenciesOk() (*[]string, bool)`

GetAcceptedCurrenciesOk returns a tuple with the AcceptedCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCurrencies

`func (o *PaymentServiceRequest) SetAcceptedCurrencies(v []string)`

SetAcceptedCurrencies sets AcceptedCurrencies field to given value.


### GetThreeDSecureEnabled

`func (o *PaymentServiceRequest) GetThreeDSecureEnabled() bool`

GetThreeDSecureEnabled returns the ThreeDSecureEnabled field if non-nil, zero value otherwise.

### GetThreeDSecureEnabledOk

`func (o *PaymentServiceRequest) GetThreeDSecureEnabledOk() (*bool, bool)`

GetThreeDSecureEnabledOk returns a tuple with the ThreeDSecureEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSecureEnabled

`func (o *PaymentServiceRequest) SetThreeDSecureEnabled(v bool)`

SetThreeDSecureEnabled sets ThreeDSecureEnabled field to given value.

### HasThreeDSecureEnabled

`func (o *PaymentServiceRequest) HasThreeDSecureEnabled() bool`

HasThreeDSecureEnabled returns a boolean if a field has been set.

### GetAcquirerBinVisa

`func (o *PaymentServiceRequest) GetAcquirerBinVisa() string`

GetAcquirerBinVisa returns the AcquirerBinVisa field if non-nil, zero value otherwise.

### GetAcquirerBinVisaOk

`func (o *PaymentServiceRequest) GetAcquirerBinVisaOk() (*string, bool)`

GetAcquirerBinVisaOk returns a tuple with the AcquirerBinVisa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerBinVisa

`func (o *PaymentServiceRequest) SetAcquirerBinVisa(v string)`

SetAcquirerBinVisa sets AcquirerBinVisa field to given value.

### HasAcquirerBinVisa

`func (o *PaymentServiceRequest) HasAcquirerBinVisa() bool`

HasAcquirerBinVisa returns a boolean if a field has been set.

### SetAcquirerBinVisaNil

`func (o *PaymentServiceRequest) SetAcquirerBinVisaNil(b bool)`

 SetAcquirerBinVisaNil sets the value for AcquirerBinVisa to be an explicit nil

### UnsetAcquirerBinVisa
`func (o *PaymentServiceRequest) UnsetAcquirerBinVisa()`

UnsetAcquirerBinVisa ensures that no value is present for AcquirerBinVisa, not even an explicit nil
### GetAcquirerBinMastercard

`func (o *PaymentServiceRequest) GetAcquirerBinMastercard() string`

GetAcquirerBinMastercard returns the AcquirerBinMastercard field if non-nil, zero value otherwise.

### GetAcquirerBinMastercardOk

`func (o *PaymentServiceRequest) GetAcquirerBinMastercardOk() (*string, bool)`

GetAcquirerBinMastercardOk returns a tuple with the AcquirerBinMastercard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerBinMastercard

`func (o *PaymentServiceRequest) SetAcquirerBinMastercard(v string)`

SetAcquirerBinMastercard sets AcquirerBinMastercard field to given value.

### HasAcquirerBinMastercard

`func (o *PaymentServiceRequest) HasAcquirerBinMastercard() bool`

HasAcquirerBinMastercard returns a boolean if a field has been set.

### SetAcquirerBinMastercardNil

`func (o *PaymentServiceRequest) SetAcquirerBinMastercardNil(b bool)`

 SetAcquirerBinMastercardNil sets the value for AcquirerBinMastercard to be an explicit nil

### UnsetAcquirerBinMastercard
`func (o *PaymentServiceRequest) UnsetAcquirerBinMastercard()`

UnsetAcquirerBinMastercard ensures that no value is present for AcquirerBinMastercard, not even an explicit nil
### GetAcquirerBinAmex

`func (o *PaymentServiceRequest) GetAcquirerBinAmex() string`

GetAcquirerBinAmex returns the AcquirerBinAmex field if non-nil, zero value otherwise.

### GetAcquirerBinAmexOk

`func (o *PaymentServiceRequest) GetAcquirerBinAmexOk() (*string, bool)`

GetAcquirerBinAmexOk returns a tuple with the AcquirerBinAmex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerBinAmex

`func (o *PaymentServiceRequest) SetAcquirerBinAmex(v string)`

SetAcquirerBinAmex sets AcquirerBinAmex field to given value.

### HasAcquirerBinAmex

`func (o *PaymentServiceRequest) HasAcquirerBinAmex() bool`

HasAcquirerBinAmex returns a boolean if a field has been set.

### SetAcquirerBinAmexNil

`func (o *PaymentServiceRequest) SetAcquirerBinAmexNil(b bool)`

 SetAcquirerBinAmexNil sets the value for AcquirerBinAmex to be an explicit nil

### UnsetAcquirerBinAmex
`func (o *PaymentServiceRequest) UnsetAcquirerBinAmex()`

UnsetAcquirerBinAmex ensures that no value is present for AcquirerBinAmex, not even an explicit nil
### GetAcquirerBinDiscover

`func (o *PaymentServiceRequest) GetAcquirerBinDiscover() string`

GetAcquirerBinDiscover returns the AcquirerBinDiscover field if non-nil, zero value otherwise.

### GetAcquirerBinDiscoverOk

`func (o *PaymentServiceRequest) GetAcquirerBinDiscoverOk() (*string, bool)`

GetAcquirerBinDiscoverOk returns a tuple with the AcquirerBinDiscover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerBinDiscover

`func (o *PaymentServiceRequest) SetAcquirerBinDiscover(v string)`

SetAcquirerBinDiscover sets AcquirerBinDiscover field to given value.

### HasAcquirerBinDiscover

`func (o *PaymentServiceRequest) HasAcquirerBinDiscover() bool`

HasAcquirerBinDiscover returns a boolean if a field has been set.

### SetAcquirerBinDiscoverNil

`func (o *PaymentServiceRequest) SetAcquirerBinDiscoverNil(b bool)`

 SetAcquirerBinDiscoverNil sets the value for AcquirerBinDiscover to be an explicit nil

### UnsetAcquirerBinDiscover
`func (o *PaymentServiceRequest) UnsetAcquirerBinDiscover()`

UnsetAcquirerBinDiscover ensures that no value is present for AcquirerBinDiscover, not even an explicit nil
### GetAcquirerMerchantId

`func (o *PaymentServiceRequest) GetAcquirerMerchantId() string`

GetAcquirerMerchantId returns the AcquirerMerchantId field if non-nil, zero value otherwise.

### GetAcquirerMerchantIdOk

`func (o *PaymentServiceRequest) GetAcquirerMerchantIdOk() (*string, bool)`

GetAcquirerMerchantIdOk returns a tuple with the AcquirerMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerMerchantId

`func (o *PaymentServiceRequest) SetAcquirerMerchantId(v string)`

SetAcquirerMerchantId sets AcquirerMerchantId field to given value.

### HasAcquirerMerchantId

`func (o *PaymentServiceRequest) HasAcquirerMerchantId() bool`

HasAcquirerMerchantId returns a boolean if a field has been set.

### SetAcquirerMerchantIdNil

`func (o *PaymentServiceRequest) SetAcquirerMerchantIdNil(b bool)`

 SetAcquirerMerchantIdNil sets the value for AcquirerMerchantId to be an explicit nil

### UnsetAcquirerMerchantId
`func (o *PaymentServiceRequest) UnsetAcquirerMerchantId()`

UnsetAcquirerMerchantId ensures that no value is present for AcquirerMerchantId, not even an explicit nil
### GetMerchantName

`func (o *PaymentServiceRequest) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *PaymentServiceRequest) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *PaymentServiceRequest) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *PaymentServiceRequest) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### SetMerchantNameNil

`func (o *PaymentServiceRequest) SetMerchantNameNil(b bool)`

 SetMerchantNameNil sets the value for MerchantName to be an explicit nil

### UnsetMerchantName
`func (o *PaymentServiceRequest) UnsetMerchantName()`

UnsetMerchantName ensures that no value is present for MerchantName, not even an explicit nil
### GetMerchantCountryCode

`func (o *PaymentServiceRequest) GetMerchantCountryCode() string`

GetMerchantCountryCode returns the MerchantCountryCode field if non-nil, zero value otherwise.

### GetMerchantCountryCodeOk

`func (o *PaymentServiceRequest) GetMerchantCountryCodeOk() (*string, bool)`

GetMerchantCountryCodeOk returns a tuple with the MerchantCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCountryCode

`func (o *PaymentServiceRequest) SetMerchantCountryCode(v string)`

SetMerchantCountryCode sets MerchantCountryCode field to given value.

### HasMerchantCountryCode

`func (o *PaymentServiceRequest) HasMerchantCountryCode() bool`

HasMerchantCountryCode returns a boolean if a field has been set.

### SetMerchantCountryCodeNil

`func (o *PaymentServiceRequest) SetMerchantCountryCodeNil(b bool)`

 SetMerchantCountryCodeNil sets the value for MerchantCountryCode to be an explicit nil

### UnsetMerchantCountryCode
`func (o *PaymentServiceRequest) UnsetMerchantCountryCode()`

UnsetMerchantCountryCode ensures that no value is present for MerchantCountryCode, not even an explicit nil
### GetMerchantCategoryCode

`func (o *PaymentServiceRequest) GetMerchantCategoryCode() string`

GetMerchantCategoryCode returns the MerchantCategoryCode field if non-nil, zero value otherwise.

### GetMerchantCategoryCodeOk

`func (o *PaymentServiceRequest) GetMerchantCategoryCodeOk() (*string, bool)`

GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCode

`func (o *PaymentServiceRequest) SetMerchantCategoryCode(v string)`

SetMerchantCategoryCode sets MerchantCategoryCode field to given value.

### HasMerchantCategoryCode

`func (o *PaymentServiceRequest) HasMerchantCategoryCode() bool`

HasMerchantCategoryCode returns a boolean if a field has been set.

### SetMerchantCategoryCodeNil

`func (o *PaymentServiceRequest) SetMerchantCategoryCodeNil(b bool)`

 SetMerchantCategoryCodeNil sets the value for MerchantCategoryCode to be an explicit nil

### UnsetMerchantCategoryCode
`func (o *PaymentServiceRequest) UnsetMerchantCategoryCode()`

UnsetMerchantCategoryCode ensures that no value is present for MerchantCategoryCode, not even an explicit nil
### GetMerchantUrl

`func (o *PaymentServiceRequest) GetMerchantUrl() string`

GetMerchantUrl returns the MerchantUrl field if non-nil, zero value otherwise.

### GetMerchantUrlOk

`func (o *PaymentServiceRequest) GetMerchantUrlOk() (*string, bool)`

GetMerchantUrlOk returns a tuple with the MerchantUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantUrl

`func (o *PaymentServiceRequest) SetMerchantUrl(v string)`

SetMerchantUrl sets MerchantUrl field to given value.

### HasMerchantUrl

`func (o *PaymentServiceRequest) HasMerchantUrl() bool`

HasMerchantUrl returns a boolean if a field has been set.

### SetMerchantUrlNil

`func (o *PaymentServiceRequest) SetMerchantUrlNil(b bool)`

 SetMerchantUrlNil sets the value for MerchantUrl to be an explicit nil

### UnsetMerchantUrl
`func (o *PaymentServiceRequest) UnsetMerchantUrl()`

UnsetMerchantUrl ensures that no value is present for MerchantUrl, not even an explicit nil
### GetActive

`func (o *PaymentServiceRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PaymentServiceRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PaymentServiceRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PaymentServiceRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPosition

`func (o *PaymentServiceRequest) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PaymentServiceRequest) GetPositionOk() (*float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PaymentServiceRequest) SetPosition(v float32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PaymentServiceRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPaymentMethodTokenizationEnabled

`func (o *PaymentServiceRequest) GetPaymentMethodTokenizationEnabled() bool`

GetPaymentMethodTokenizationEnabled returns the PaymentMethodTokenizationEnabled field if non-nil, zero value otherwise.

### GetPaymentMethodTokenizationEnabledOk

`func (o *PaymentServiceRequest) GetPaymentMethodTokenizationEnabledOk() (*bool, bool)`

GetPaymentMethodTokenizationEnabledOk returns a tuple with the PaymentMethodTokenizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTokenizationEnabled

`func (o *PaymentServiceRequest) SetPaymentMethodTokenizationEnabled(v bool)`

SetPaymentMethodTokenizationEnabled sets PaymentMethodTokenizationEnabled field to given value.

### HasPaymentMethodTokenizationEnabled

`func (o *PaymentServiceRequest) HasPaymentMethodTokenizationEnabled() bool`

HasPaymentMethodTokenizationEnabled returns a boolean if a field has been set.

### GetPaymentServiceDefinitionId

`func (o *PaymentServiceRequest) GetPaymentServiceDefinitionId() string`

GetPaymentServiceDefinitionId returns the PaymentServiceDefinitionId field if non-nil, zero value otherwise.

### GetPaymentServiceDefinitionIdOk

`func (o *PaymentServiceRequest) GetPaymentServiceDefinitionIdOk() (*string, bool)`

GetPaymentServiceDefinitionIdOk returns a tuple with the PaymentServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceDefinitionId

`func (o *PaymentServiceRequest) SetPaymentServiceDefinitionId(v string)`

SetPaymentServiceDefinitionId sets PaymentServiceDefinitionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


