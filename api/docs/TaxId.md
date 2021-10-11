# TaxId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The tax ID for the buyer. | [optional] 
**Kind** | **string** | The kind of tax ID. | 

## Methods

### NewTaxId

`func NewTaxId(kind string, ) *TaxId`

NewTaxId instantiates a new TaxId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxIdWithDefaults

`func NewTaxIdWithDefaults() *TaxId`

NewTaxIdWithDefaults instantiates a new TaxId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxId) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxId) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxId) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxId) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *TaxId) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TaxId) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TaxId) SetKind(v string)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


