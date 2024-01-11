# ConnectionOptionsCybersourceAntiFraud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaKeyMerchantId** | Pointer to **NullableString** | An override for the merchant ID configured for the connector, used in combination with meta keys. | [optional] 
**MerchantDefinedData** | Pointer to **map[string]string** | This is a key-value object for merchant defined data. Each key needs to be a numeric string identifying the MDD field to set. For example, for field 1 set the key to \&quot;1\&quot;. | [optional] 

## Methods

### NewConnectionOptionsCybersourceAntiFraud

`func NewConnectionOptionsCybersourceAntiFraud() *ConnectionOptionsCybersourceAntiFraud`

NewConnectionOptionsCybersourceAntiFraud instantiates a new ConnectionOptionsCybersourceAntiFraud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionOptionsCybersourceAntiFraudWithDefaults

`func NewConnectionOptionsCybersourceAntiFraudWithDefaults() *ConnectionOptionsCybersourceAntiFraud`

NewConnectionOptionsCybersourceAntiFraudWithDefaults instantiates a new ConnectionOptionsCybersourceAntiFraud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaKeyMerchantId

`func (o *ConnectionOptionsCybersourceAntiFraud) GetMetaKeyMerchantId() string`

GetMetaKeyMerchantId returns the MetaKeyMerchantId field if non-nil, zero value otherwise.

### GetMetaKeyMerchantIdOk

`func (o *ConnectionOptionsCybersourceAntiFraud) GetMetaKeyMerchantIdOk() (*string, bool)`

GetMetaKeyMerchantIdOk returns a tuple with the MetaKeyMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaKeyMerchantId

`func (o *ConnectionOptionsCybersourceAntiFraud) SetMetaKeyMerchantId(v string)`

SetMetaKeyMerchantId sets MetaKeyMerchantId field to given value.

### HasMetaKeyMerchantId

`func (o *ConnectionOptionsCybersourceAntiFraud) HasMetaKeyMerchantId() bool`

HasMetaKeyMerchantId returns a boolean if a field has been set.

### SetMetaKeyMerchantIdNil

`func (o *ConnectionOptionsCybersourceAntiFraud) SetMetaKeyMerchantIdNil(b bool)`

 SetMetaKeyMerchantIdNil sets the value for MetaKeyMerchantId to be an explicit nil

### UnsetMetaKeyMerchantId
`func (o *ConnectionOptionsCybersourceAntiFraud) UnsetMetaKeyMerchantId()`

UnsetMetaKeyMerchantId ensures that no value is present for MetaKeyMerchantId, not even an explicit nil
### GetMerchantDefinedData

`func (o *ConnectionOptionsCybersourceAntiFraud) GetMerchantDefinedData() map[string]string`

GetMerchantDefinedData returns the MerchantDefinedData field if non-nil, zero value otherwise.

### GetMerchantDefinedDataOk

`func (o *ConnectionOptionsCybersourceAntiFraud) GetMerchantDefinedDataOk() (*map[string]string, bool)`

GetMerchantDefinedDataOk returns a tuple with the MerchantDefinedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantDefinedData

`func (o *ConnectionOptionsCybersourceAntiFraud) SetMerchantDefinedData(v map[string]string)`

SetMerchantDefinedData sets MerchantDefinedData field to given value.

### HasMerchantDefinedData

`func (o *ConnectionOptionsCybersourceAntiFraud) HasMerchantDefinedData() bool`

HasMerchantDefinedData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


