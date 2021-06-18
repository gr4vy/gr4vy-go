# CardRuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Match** | **string** | The type of match made for this rule. | 
**Key** | **string** | The transaction field to filter by. | 
**Operator** | **string** | The comparison to make to &#x60;value&#x60; property. | 
**Value** | [**string**](string.md) |  | 

## Methods

### NewCardRuleCondition

`func NewCardRuleCondition(match string, key string, operator string, value string, ) *CardRuleCondition`

NewCardRuleCondition instantiates a new CardRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardRuleConditionWithDefaults

`func NewCardRuleConditionWithDefaults() *CardRuleCondition`

NewCardRuleConditionWithDefaults instantiates a new CardRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatch

`func (o *CardRuleCondition) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *CardRuleCondition) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *CardRuleCondition) SetMatch(v string)`

SetMatch sets Match field to given value.


### GetKey

`func (o *CardRuleCondition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CardRuleCondition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CardRuleCondition) SetKey(v string)`

SetKey sets Key field to given value.


### GetOperator

`func (o *CardRuleCondition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CardRuleCondition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CardRuleCondition) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *CardRuleCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CardRuleCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CardRuleCondition) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


