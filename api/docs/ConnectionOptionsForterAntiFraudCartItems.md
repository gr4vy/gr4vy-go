# ConnectionOptionsForterAntiFraudCartItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicItemData** | Pointer to [**ConnectionOptionsForterAntiFraudBasicItemData**](ConnectionOptionsForterAntiFraudBasicItemData.md) |  | [optional] 
**DeliveryDetails** | Pointer to [**ConnectionOptionsForterAntiFraudDeliveryDetails**](ConnectionOptionsForterAntiFraudDeliveryDetails.md) |  | [optional] 
**Beneficiaries** | Pointer to [**[]ConnectionOptionsForterAntiFraudBeneficiaries**](ConnectionOptionsForterAntiFraudBeneficiaries.md) | List of all entities receiving or using the purchased cart item. | [optional] 

## Methods

### NewConnectionOptionsForterAntiFraudCartItems

`func NewConnectionOptionsForterAntiFraudCartItems() *ConnectionOptionsForterAntiFraudCartItems`

NewConnectionOptionsForterAntiFraudCartItems instantiates a new ConnectionOptionsForterAntiFraudCartItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionOptionsForterAntiFraudCartItemsWithDefaults

`func NewConnectionOptionsForterAntiFraudCartItemsWithDefaults() *ConnectionOptionsForterAntiFraudCartItems`

NewConnectionOptionsForterAntiFraudCartItemsWithDefaults instantiates a new ConnectionOptionsForterAntiFraudCartItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicItemData

`func (o *ConnectionOptionsForterAntiFraudCartItems) GetBasicItemData() ConnectionOptionsForterAntiFraudBasicItemData`

GetBasicItemData returns the BasicItemData field if non-nil, zero value otherwise.

### GetBasicItemDataOk

`func (o *ConnectionOptionsForterAntiFraudCartItems) GetBasicItemDataOk() (*ConnectionOptionsForterAntiFraudBasicItemData, bool)`

GetBasicItemDataOk returns a tuple with the BasicItemData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicItemData

`func (o *ConnectionOptionsForterAntiFraudCartItems) SetBasicItemData(v ConnectionOptionsForterAntiFraudBasicItemData)`

SetBasicItemData sets BasicItemData field to given value.

### HasBasicItemData

`func (o *ConnectionOptionsForterAntiFraudCartItems) HasBasicItemData() bool`

HasBasicItemData returns a boolean if a field has been set.

### GetDeliveryDetails

`func (o *ConnectionOptionsForterAntiFraudCartItems) GetDeliveryDetails() ConnectionOptionsForterAntiFraudDeliveryDetails`

GetDeliveryDetails returns the DeliveryDetails field if non-nil, zero value otherwise.

### GetDeliveryDetailsOk

`func (o *ConnectionOptionsForterAntiFraudCartItems) GetDeliveryDetailsOk() (*ConnectionOptionsForterAntiFraudDeliveryDetails, bool)`

GetDeliveryDetailsOk returns a tuple with the DeliveryDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDetails

`func (o *ConnectionOptionsForterAntiFraudCartItems) SetDeliveryDetails(v ConnectionOptionsForterAntiFraudDeliveryDetails)`

SetDeliveryDetails sets DeliveryDetails field to given value.

### HasDeliveryDetails

`func (o *ConnectionOptionsForterAntiFraudCartItems) HasDeliveryDetails() bool`

HasDeliveryDetails returns a boolean if a field has been set.

### GetBeneficiaries

`func (o *ConnectionOptionsForterAntiFraudCartItems) GetBeneficiaries() []ConnectionOptionsForterAntiFraudBeneficiaries`

GetBeneficiaries returns the Beneficiaries field if non-nil, zero value otherwise.

### GetBeneficiariesOk

`func (o *ConnectionOptionsForterAntiFraudCartItems) GetBeneficiariesOk() (*[]ConnectionOptionsForterAntiFraudBeneficiaries, bool)`

GetBeneficiariesOk returns a tuple with the Beneficiaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaries

`func (o *ConnectionOptionsForterAntiFraudCartItems) SetBeneficiaries(v []ConnectionOptionsForterAntiFraudBeneficiaries)`

SetBeneficiaries sets Beneficiaries field to given value.

### HasBeneficiaries

`func (o *ConnectionOptionsForterAntiFraudCartItems) HasBeneficiaries() bool`

HasBeneficiaries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


