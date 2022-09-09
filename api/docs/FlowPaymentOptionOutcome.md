# FlowPaymentOptionOutcome

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this resource. Is always &#x60;action&#x60;. | [optional] 
**Id** | Pointer to **string** | Payment option identifier. | [optional] 
**Label** | Pointer to **string** | Verbose payment option name. | [optional] 
**Active** | Pointer to **bool** | The status of the payment option, true if at least one underlying connection is active, otherwise false. | [optional] 
**Group** | Pointer to **NullableString** | Optional group label for a given payment option, e.g. &#x60;Bank&#x60;. | [optional] 
**IconUrl** | Pointer to **string** | Payment option icon URL. | [optional] 

## Methods

### NewFlowPaymentOptionOutcome

`func NewFlowPaymentOptionOutcome() *FlowPaymentOptionOutcome`

NewFlowPaymentOptionOutcome instantiates a new FlowPaymentOptionOutcome object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowPaymentOptionOutcomeWithDefaults

`func NewFlowPaymentOptionOutcomeWithDefaults() *FlowPaymentOptionOutcome`

NewFlowPaymentOptionOutcomeWithDefaults instantiates a new FlowPaymentOptionOutcome object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FlowPaymentOptionOutcome) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlowPaymentOptionOutcome) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlowPaymentOptionOutcome) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FlowPaymentOptionOutcome) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *FlowPaymentOptionOutcome) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowPaymentOptionOutcome) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowPaymentOptionOutcome) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlowPaymentOptionOutcome) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *FlowPaymentOptionOutcome) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FlowPaymentOptionOutcome) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FlowPaymentOptionOutcome) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FlowPaymentOptionOutcome) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetActive

`func (o *FlowPaymentOptionOutcome) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FlowPaymentOptionOutcome) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FlowPaymentOptionOutcome) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *FlowPaymentOptionOutcome) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetGroup

`func (o *FlowPaymentOptionOutcome) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *FlowPaymentOptionOutcome) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *FlowPaymentOptionOutcome) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *FlowPaymentOptionOutcome) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *FlowPaymentOptionOutcome) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *FlowPaymentOptionOutcome) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetIconUrl

`func (o *FlowPaymentOptionOutcome) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *FlowPaymentOptionOutcome) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *FlowPaymentOptionOutcome) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *FlowPaymentOptionOutcome) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


