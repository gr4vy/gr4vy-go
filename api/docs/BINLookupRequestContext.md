# BINLookupRequestContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to **NullableString** | The response body received from the &#x60;url&#x60;. | [optional] 
**ResponseStatusCode** | Pointer to **int32** | The response status code received from the &#x60;url&#x60;. | [optional] 
**Success** | Pointer to **bool** | Whether the BIN lookup was successful or not. | [optional] 
**Bin** | Pointer to **NullableString** | The value used to lookup BIN details. | [optional] 
**Type** | Pointer to **NullableString** | The type of card, i.e. credit or debit, from the lookup response. | [optional] 
**Scheme** | Pointer to **NullableString** | The card scheme result from the lookup response. | [optional] 
**AdditionalSchemes** | Pointer to **[]string** | The card additional schemes from the lookup response. | [optional] 
**CountryCode** | Pointer to **NullableString** | The card country code result from the lookup response. | [optional] 
**AccountUpdater** | Pointer to **NullableBool** | Whether Account Updater is enabled for this card. | [optional] 
**IssuerTokenization** | Pointer to **NullableBool** | Whether the issuing bank supports network tokenization for this card. | [optional] 

## Methods

### NewBINLookupRequestContext

`func NewBINLookupRequestContext() *BINLookupRequestContext`

NewBINLookupRequestContext instantiates a new BINLookupRequestContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBINLookupRequestContextWithDefaults

`func NewBINLookupRequestContextWithDefaults() *BINLookupRequestContext`

NewBINLookupRequestContextWithDefaults instantiates a new BINLookupRequestContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *BINLookupRequestContext) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *BINLookupRequestContext) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *BINLookupRequestContext) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *BINLookupRequestContext) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *BINLookupRequestContext) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *BINLookupRequestContext) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetResponseStatusCode

`func (o *BINLookupRequestContext) GetResponseStatusCode() int32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *BINLookupRequestContext) GetResponseStatusCodeOk() (*int32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *BINLookupRequestContext) SetResponseStatusCode(v int32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *BINLookupRequestContext) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### GetSuccess

`func (o *BINLookupRequestContext) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *BINLookupRequestContext) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *BINLookupRequestContext) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *BINLookupRequestContext) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetBin

`func (o *BINLookupRequestContext) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *BINLookupRequestContext) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *BINLookupRequestContext) SetBin(v string)`

SetBin sets Bin field to given value.

### HasBin

`func (o *BINLookupRequestContext) HasBin() bool`

HasBin returns a boolean if a field has been set.

### SetBinNil

`func (o *BINLookupRequestContext) SetBinNil(b bool)`

 SetBinNil sets the value for Bin to be an explicit nil

### UnsetBin
`func (o *BINLookupRequestContext) UnsetBin()`

UnsetBin ensures that no value is present for Bin, not even an explicit nil
### GetType

`func (o *BINLookupRequestContext) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BINLookupRequestContext) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BINLookupRequestContext) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BINLookupRequestContext) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BINLookupRequestContext) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BINLookupRequestContext) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetScheme

`func (o *BINLookupRequestContext) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *BINLookupRequestContext) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *BINLookupRequestContext) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *BINLookupRequestContext) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *BINLookupRequestContext) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *BINLookupRequestContext) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetAdditionalSchemes

`func (o *BINLookupRequestContext) GetAdditionalSchemes() []string`

GetAdditionalSchemes returns the AdditionalSchemes field if non-nil, zero value otherwise.

### GetAdditionalSchemesOk

`func (o *BINLookupRequestContext) GetAdditionalSchemesOk() (*[]string, bool)`

GetAdditionalSchemesOk returns a tuple with the AdditionalSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSchemes

`func (o *BINLookupRequestContext) SetAdditionalSchemes(v []string)`

SetAdditionalSchemes sets AdditionalSchemes field to given value.

### HasAdditionalSchemes

`func (o *BINLookupRequestContext) HasAdditionalSchemes() bool`

HasAdditionalSchemes returns a boolean if a field has been set.

### SetAdditionalSchemesNil

`func (o *BINLookupRequestContext) SetAdditionalSchemesNil(b bool)`

 SetAdditionalSchemesNil sets the value for AdditionalSchemes to be an explicit nil

### UnsetAdditionalSchemes
`func (o *BINLookupRequestContext) UnsetAdditionalSchemes()`

UnsetAdditionalSchemes ensures that no value is present for AdditionalSchemes, not even an explicit nil
### GetCountryCode

`func (o *BINLookupRequestContext) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *BINLookupRequestContext) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *BINLookupRequestContext) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *BINLookupRequestContext) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *BINLookupRequestContext) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *BINLookupRequestContext) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetAccountUpdater

`func (o *BINLookupRequestContext) GetAccountUpdater() bool`

GetAccountUpdater returns the AccountUpdater field if non-nil, zero value otherwise.

### GetAccountUpdaterOk

`func (o *BINLookupRequestContext) GetAccountUpdaterOk() (*bool, bool)`

GetAccountUpdaterOk returns a tuple with the AccountUpdater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUpdater

`func (o *BINLookupRequestContext) SetAccountUpdater(v bool)`

SetAccountUpdater sets AccountUpdater field to given value.

### HasAccountUpdater

`func (o *BINLookupRequestContext) HasAccountUpdater() bool`

HasAccountUpdater returns a boolean if a field has been set.

### SetAccountUpdaterNil

`func (o *BINLookupRequestContext) SetAccountUpdaterNil(b bool)`

 SetAccountUpdaterNil sets the value for AccountUpdater to be an explicit nil

### UnsetAccountUpdater
`func (o *BINLookupRequestContext) UnsetAccountUpdater()`

UnsetAccountUpdater ensures that no value is present for AccountUpdater, not even an explicit nil
### GetIssuerTokenization

`func (o *BINLookupRequestContext) GetIssuerTokenization() bool`

GetIssuerTokenization returns the IssuerTokenization field if non-nil, zero value otherwise.

### GetIssuerTokenizationOk

`func (o *BINLookupRequestContext) GetIssuerTokenizationOk() (*bool, bool)`

GetIssuerTokenizationOk returns a tuple with the IssuerTokenization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerTokenization

`func (o *BINLookupRequestContext) SetIssuerTokenization(v bool)`

SetIssuerTokenization sets IssuerTokenization field to given value.

### HasIssuerTokenization

`func (o *BINLookupRequestContext) HasIssuerTokenization() bool`

HasIssuerTokenization returns a boolean if a field has been set.

### SetIssuerTokenizationNil

`func (o *BINLookupRequestContext) SetIssuerTokenizationNil(b bool)`

 SetIssuerTokenizationNil sets the value for IssuerTokenization to be an explicit nil

### UnsetIssuerTokenization
`func (o *BINLookupRequestContext) UnsetIssuerTokenization()`

UnsetIssuerTokenization ensures that no value is present for IssuerTokenization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


