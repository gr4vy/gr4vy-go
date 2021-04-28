# PaymentServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | A custom name for the payment service. This will be shown in the Admin UI. | 
**Fields** | [**[]PaymentServiceUpdateFields**](PaymentServiceUpdateFields.md) | A list of fields, each containing a key-value pair for each field defined by the definition for this payment service e.g. for stripe-card &#x60;secret_key&#x60; is required and so must be sent with in this field. | 
**AcceptedCountries** | **[]string** | A list of countries that this payment service needs to support in ISO two-letter code format. | 
**AcceptedCurrencies** | **[]string** | A list of currencies that this payment service needs to support in ISO 4217 three-letter code format. | 
**CredentialsMode** | Pointer to **string** | Defines if the credentials are intended for the service&#39;s live API or sandbox/test API. | [optional] [default to "live"]
**Active** | Pointer to **bool** | Defines if this service is currently active or not. | [optional] [default to true]
**Environments** | Pointer to **[]string** | Determines the Gr4vy environments in which this service should be available. This can be used in combination with the &#x60;environment&#x60; parameters in the payment method and transaction APIs to route transactions through this service. | [optional] [default to ["production"]]
**Position** | Pointer to **float32** | The numeric rank of a payment service. Payment services with a lower position value are processed first. When a payment services is inserted at a position, any payment services with the the same value or higher are shifted down a position accordingly. When left out, the payment service is inserted at the end of the list. | [optional] 
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


### GetCredentialsMode

`func (o *PaymentServiceRequest) GetCredentialsMode() string`

GetCredentialsMode returns the CredentialsMode field if non-nil, zero value otherwise.

### GetCredentialsModeOk

`func (o *PaymentServiceRequest) GetCredentialsModeOk() (*string, bool)`

GetCredentialsModeOk returns a tuple with the CredentialsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsMode

`func (o *PaymentServiceRequest) SetCredentialsMode(v string)`

SetCredentialsMode sets CredentialsMode field to given value.

### HasCredentialsMode

`func (o *PaymentServiceRequest) HasCredentialsMode() bool`

HasCredentialsMode returns a boolean if a field has been set.

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

### GetEnvironments

`func (o *PaymentServiceRequest) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *PaymentServiceRequest) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *PaymentServiceRequest) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *PaymentServiceRequest) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

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


