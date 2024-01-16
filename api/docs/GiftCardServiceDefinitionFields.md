# GiftCardServiceDefinitionFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The key of a field that needs to be submitted. | [optional] 
**DisplayName** | Pointer to **string** | The name to display for a field in the dashboard. | [optional] 
**Required** | Pointer to **bool** | Defines if this field is required when the service is created. | [optional] 
**Format** | Pointer to **string** | Defines the type of input that needs to be rendered for this field. | [optional] 
**Secret** | Pointer to **bool** | Defines if this field is secret. When &#x60;true&#x60; the field is not returned when querying the payment service. | [optional] 

## Methods

### NewGiftCardServiceDefinitionFields

`func NewGiftCardServiceDefinitionFields() *GiftCardServiceDefinitionFields`

NewGiftCardServiceDefinitionFields instantiates a new GiftCardServiceDefinitionFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardServiceDefinitionFieldsWithDefaults

`func NewGiftCardServiceDefinitionFieldsWithDefaults() *GiftCardServiceDefinitionFields`

NewGiftCardServiceDefinitionFieldsWithDefaults instantiates a new GiftCardServiceDefinitionFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *GiftCardServiceDefinitionFields) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GiftCardServiceDefinitionFields) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GiftCardServiceDefinitionFields) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GiftCardServiceDefinitionFields) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetDisplayName

`func (o *GiftCardServiceDefinitionFields) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GiftCardServiceDefinitionFields) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GiftCardServiceDefinitionFields) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GiftCardServiceDefinitionFields) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetRequired

`func (o *GiftCardServiceDefinitionFields) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *GiftCardServiceDefinitionFields) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *GiftCardServiceDefinitionFields) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *GiftCardServiceDefinitionFields) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetFormat

`func (o *GiftCardServiceDefinitionFields) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *GiftCardServiceDefinitionFields) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *GiftCardServiceDefinitionFields) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *GiftCardServiceDefinitionFields) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetSecret

`func (o *GiftCardServiceDefinitionFields) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *GiftCardServiceDefinitionFields) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *GiftCardServiceDefinitionFields) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *GiftCardServiceDefinitionFields) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


