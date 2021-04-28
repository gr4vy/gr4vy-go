# PaymentServiceDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the payment service. This is the underlying provider followed by a dash followed by the payment method ID. | [optional] 
**Type** | Pointer to **string** | &#x60;payment-service-definition&#x60;. | [optional] [default to "payment-service-definition"]
**DisplayName** | Pointer to **string** | The display name of this service. | [optional] 
**Method** | Pointer to **string** | The ID of the payment method that this services handles. | [optional] 
**Fields** | Pointer to [**[]PaymentServiceDefinitionFields**](PaymentServiceDefinitionFields.md) | A list of fields that need to be submitted when activating the payment. service. | [optional] 
**SupportedCurrencies** | Pointer to **[]string** | A list of three-letter ISO currency codes that this service supports. | [optional] 
**SupportedCountries** | Pointer to **[]string** | A list of two-letter ISO country codes that this service supports. | [optional] 

## Methods

### NewPaymentServiceDefinition

`func NewPaymentServiceDefinition() *PaymentServiceDefinition`

NewPaymentServiceDefinition instantiates a new PaymentServiceDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceDefinitionWithDefaults

`func NewPaymentServiceDefinitionWithDefaults() *PaymentServiceDefinition`

NewPaymentServiceDefinitionWithDefaults instantiates a new PaymentServiceDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentServiceDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentServiceDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentServiceDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentServiceDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PaymentServiceDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentServiceDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentServiceDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentServiceDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayName

`func (o *PaymentServiceDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PaymentServiceDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PaymentServiceDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PaymentServiceDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMethod

`func (o *PaymentServiceDefinition) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PaymentServiceDefinition) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PaymentServiceDefinition) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *PaymentServiceDefinition) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetFields

`func (o *PaymentServiceDefinition) GetFields() []PaymentServiceDefinitionFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PaymentServiceDefinition) GetFieldsOk() (*[]PaymentServiceDefinitionFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PaymentServiceDefinition) SetFields(v []PaymentServiceDefinitionFields)`

SetFields sets Fields field to given value.

### HasFields

`func (o *PaymentServiceDefinition) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetSupportedCurrencies

`func (o *PaymentServiceDefinition) GetSupportedCurrencies() []string`

GetSupportedCurrencies returns the SupportedCurrencies field if non-nil, zero value otherwise.

### GetSupportedCurrenciesOk

`func (o *PaymentServiceDefinition) GetSupportedCurrenciesOk() (*[]string, bool)`

GetSupportedCurrenciesOk returns a tuple with the SupportedCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCurrencies

`func (o *PaymentServiceDefinition) SetSupportedCurrencies(v []string)`

SetSupportedCurrencies sets SupportedCurrencies field to given value.

### HasSupportedCurrencies

`func (o *PaymentServiceDefinition) HasSupportedCurrencies() bool`

HasSupportedCurrencies returns a boolean if a field has been set.

### GetSupportedCountries

`func (o *PaymentServiceDefinition) GetSupportedCountries() []string`

GetSupportedCountries returns the SupportedCountries field if non-nil, zero value otherwise.

### GetSupportedCountriesOk

`func (o *PaymentServiceDefinition) GetSupportedCountriesOk() (*[]string, bool)`

GetSupportedCountriesOk returns a tuple with the SupportedCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCountries

`func (o *PaymentServiceDefinition) SetSupportedCountries(v []string)`

SetSupportedCountries sets SupportedCountries field to given value.

### HasSupportedCountries

`func (o *PaymentServiceDefinition) HasSupportedCountries() bool`

HasSupportedCountries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


