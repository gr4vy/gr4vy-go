# CardRuleTextCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Match** | **string** | &#x60;text&#x60;. | 
**Key** | **string** | The transaction field to filter by. | 
**Operator** | **string** | The comparison to make to &#x60;value&#x60; property. | 
**Values** | **[]string** | The values to compare the &#x60;key&#x60; to. | 

## Methods

### NewCardRuleTextCondition

`func NewCardRuleTextCondition(match string, key string, operator string, values []string, ) *CardRuleTextCondition`

NewCardRuleTextCondition instantiates a new CardRuleTextCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardRuleTextConditionWithDefaults

`func NewCardRuleTextConditionWithDefaults() *CardRuleTextCondition`

NewCardRuleTextConditionWithDefaults instantiates a new CardRuleTextCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatch

`func (o *CardRuleTextCondition) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *CardRuleTextCondition) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *CardRuleTextCondition) SetMatch(v string)`

SetMatch sets Match field to given value.


### GetKey

`func (o *CardRuleTextCondition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CardRuleTextCondition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CardRuleTextCondition) SetKey(v string)`

SetKey sets Key field to given value.


### GetOperator

`func (o *CardRuleTextCondition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CardRuleTextCondition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CardRuleTextCondition) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValues

`func (o *CardRuleTextCondition) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CardRuleTextCondition) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CardRuleTextCondition) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


