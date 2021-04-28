# PaymentService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of this payment service. | [optional] 
**Type** | Pointer to **string** | The type of this resource. | [optional] 
**PaymentServiceDefinitionId** | Pointer to **string** | The ID of the payment service definition used to create this service.  | [optional] 
**Method** | Pointer to **string** | Defines the ID of the payment method that this service handles. | [optional] 
**DisplayName** | Pointer to **string** | The custom name set for this service. | [optional] 
**Status** | Pointer to **string** | The current status of this service. This will start off as pending, move to created, and might eventually move to an error status if and when the credentials are no longer valid.  | [optional] 
**AcceptedCurrencies** | Pointer to **[]string** | A list of currencies for which this service is enabled, in ISO 4217 three-letter code format. | [optional] 
**AcceptedCountries** | Pointer to **[]string** | A list of countries for which this service is enabled, in ISO two-letter code format. | [optional] 
**CredentialsMode** | Pointer to **string** | Defines if the credentials are intended for the service&#39;s live API or sandbox/test API. | [optional] [default to "live"]
**Active** | Pointer to **bool** | Defines if this service is currently active or not. | [optional] [default to true]
**Environments** | Pointer to **[]string** | Determines the Gr4vy environments in which this service should be available. This can be used in combination with the &#x60;environment&#x60; parameters in the payment method and transaction APIs to route transactions through this service. | [optional] [default to ["production"]]
**Position** | Pointer to **float32** | The numeric rank of a payment service. Payment services with a lower position value are processed first. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this service was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when this service was last updated. | [optional] 

## Methods

### NewPaymentService

`func NewPaymentService() *PaymentService`

NewPaymentService instantiates a new PaymentService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceWithDefaults

`func NewPaymentServiceWithDefaults() *PaymentService`

NewPaymentServiceWithDefaults instantiates a new PaymentService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PaymentService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPaymentServiceDefinitionId

`func (o *PaymentService) GetPaymentServiceDefinitionId() string`

GetPaymentServiceDefinitionId returns the PaymentServiceDefinitionId field if non-nil, zero value otherwise.

### GetPaymentServiceDefinitionIdOk

`func (o *PaymentService) GetPaymentServiceDefinitionIdOk() (*string, bool)`

GetPaymentServiceDefinitionIdOk returns a tuple with the PaymentServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceDefinitionId

`func (o *PaymentService) SetPaymentServiceDefinitionId(v string)`

SetPaymentServiceDefinitionId sets PaymentServiceDefinitionId field to given value.

### HasPaymentServiceDefinitionId

`func (o *PaymentService) HasPaymentServiceDefinitionId() bool`

HasPaymentServiceDefinitionId returns a boolean if a field has been set.

### GetMethod

`func (o *PaymentService) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PaymentService) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PaymentService) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PaymentService) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetDisplayName

`func (o *PaymentService) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PaymentService) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PaymentService) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PaymentService) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentService) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentService) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentService) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentService) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAcceptedCurrencies

`func (o *PaymentService) GetAcceptedCurrencies() []string`

GetAcceptedCurrencies returns the AcceptedCurrencies field if non-nil, zero value otherwise.

### GetAcceptedCurrenciesOk

`func (o *PaymentService) GetAcceptedCurrenciesOk() (*[]string, bool)`

GetAcceptedCurrenciesOk returns a tuple with the AcceptedCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCurrencies

`func (o *PaymentService) SetAcceptedCurrencies(v []string)`

SetAcceptedCurrencies sets AcceptedCurrencies field to given value.

### HasAcceptedCurrencies

`func (o *PaymentService) HasAcceptedCurrencies() bool`

HasAcceptedCurrencies returns a boolean if a field has been set.

### GetAcceptedCountries

`func (o *PaymentService) GetAcceptedCountries() []string`

GetAcceptedCountries returns the AcceptedCountries field if non-nil, zero value otherwise.

### GetAcceptedCountriesOk

`func (o *PaymentService) GetAcceptedCountriesOk() (*[]string, bool)`

GetAcceptedCountriesOk returns a tuple with the AcceptedCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCountries

`func (o *PaymentService) SetAcceptedCountries(v []string)`

SetAcceptedCountries sets AcceptedCountries field to given value.

### HasAcceptedCountries

`func (o *PaymentService) HasAcceptedCountries() bool`

HasAcceptedCountries returns a boolean if a field has been set.

### GetCredentialsMode

`func (o *PaymentService) GetCredentialsMode() string`

GetCredentialsMode returns the CredentialsMode field if non-nil, zero value otherwise.

### GetCredentialsModeOk

`func (o *PaymentService) GetCredentialsModeOk() (*string, bool)`

GetCredentialsModeOk returns a tuple with the CredentialsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsMode

`func (o *PaymentService) SetCredentialsMode(v string)`

SetCredentialsMode sets CredentialsMode field to given value.

### HasCredentialsMode

`func (o *PaymentService) HasCredentialsMode() bool`

HasCredentialsMode returns a boolean if a field has been set.

### GetActive

`func (o *PaymentService) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PaymentService) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PaymentService) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PaymentService) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEnvironments

`func (o *PaymentService) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *PaymentService) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *PaymentService) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *PaymentService) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetPosition

`func (o *PaymentService) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PaymentService) GetPositionOk() (*float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PaymentService) SetPosition(v float32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PaymentService) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaymentService) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentService) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentService) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaymentService) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PaymentService) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentService) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentService) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PaymentService) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


