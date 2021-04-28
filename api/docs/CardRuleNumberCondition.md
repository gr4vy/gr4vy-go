# CardRuleNumberCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Match** | **string** | &#x60;number&#x60;. | 
**Key** | **string** | The transaction field to filter by. | 
**Operator** | **string** | The comparison to make to &#x60;value&#x60; property. | 
**Value** | **float32** | The values to compare the &#x60;key&#x60; to. | 

## Methods

### NewCardRuleNumberCondition

`func NewCardRuleNumberCondition(match string, key string, operator string, value float32, ) *CardRuleNumberCondition`

NewCardRuleNumberCondition instantiates a new CardRuleNumberCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardRuleNumberConditionWithDefaults

`func NewCardRuleNumberConditionWithDefaults() *CardRuleNumberCondition`

NewCardRuleNumberConditionWithDefaults instantiates a new CardRuleNumberCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatch

`func (o *CardRuleNumberCondition) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *CardRuleNumberCondition) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *CardRuleNumberCondition) SetMatch(v string)`

SetMatch sets Match field to given value.


### GetKey

`func (o *CardRuleNumberCondition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CardRuleNumberCondition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CardRuleNumberCondition) SetKey(v string)`

SetKey sets Key field to given value.


### GetOperator

`func (o *CardRuleNumberCondition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CardRuleNumberCondition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CardRuleNumberCondition) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *CardRuleNumberCondition) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CardRuleNumberCondition) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CardRuleNumberCondition) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


