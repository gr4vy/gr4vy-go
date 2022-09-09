# FlowRuleMethodOutcome

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of action outcome for the given rule. | 
**Result** | **[]string** | Results for a given flow action. | 

## Methods

### NewFlowRuleMethodOutcome

`func NewFlowRuleMethodOutcome(type_ string, result []string, ) *FlowRuleMethodOutcome`

NewFlowRuleMethodOutcome instantiates a new FlowRuleMethodOutcome object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowRuleMethodOutcomeWithDefaults

`func NewFlowRuleMethodOutcomeWithDefaults() *FlowRuleMethodOutcome`

NewFlowRuleMethodOutcomeWithDefaults instantiates a new FlowRuleMethodOutcome object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FlowRuleMethodOutcome) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlowRuleMethodOutcome) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlowRuleMethodOutcome) SetType(v string)`

SetType sets Type field to given value.


### GetResult

`func (o *FlowRuleMethodOutcome) GetResult() []string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *FlowRuleMethodOutcome) GetResultOk() (*[]string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *FlowRuleMethodOutcome) SetResult(v []string)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


