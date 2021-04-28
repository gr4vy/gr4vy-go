# CardRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Whether this rule is currently in use. Rules can be deactivated to allow for them to be kept around and re-activated at a later date. | [optional] 
**Environment** | Pointer to **string** | The environment to use this rule in. This rule will only be used for transactions created in that environment. | [optional] [default to "production"]
**Position** | Pointer to **float32** | The numeric rank of a rule. Rules with a lower position value are processed first. When a rule is inserted at a position, any rules with the the same value or higher are down a position accordingly. When left out, the rule is inserted at the end of the list. | [optional] 
**Conditions** | [**[]CardRule**](CardRule.md) | One or more conditions that apply for this rule. Each condition needs to match for this rule to go into effect. | 
**PaymentServiceIds** | **[]string** | A list of IDs for the payment services to use, in order of priority. The payment services all need to process cards. | 
**UnprocessableFallbackStrategy** | Pointer to **string** | Defines what strategy to use when all of the payment services defined in this rule declined or otherwise were not able to process the card.  * &#x60;use_all_providers&#x60; - Try all payment services enabled for this currency in order of priority, even if they are not listed in this rule. This is the default behaviour for a rule. * &#x60;decline&#x60; - Decline the transaction. | [optional] [default to "use_all_providers"]
**InvalidRuleFallbackStrategy** | Pointer to **string** | Defines what strategy to use when this rule is not valid. This can happen when the rule has triggered for a certain transaction but none of the listed payment services are eligible to process that transaction currency.  * &#x60;use_all_providers&#x60; - Try all payment services enabled for this currency in order of priority, even if they are not listed in this rule. This is the default behaviour for a rule. * &#x60;skip&#x60; - Skip this rule and instead move on to the next highest priority rule. * &#x60;decline&#x60; - Decline the transaction. | [optional] [default to "use_all_providers"]

## Methods

### NewCardRuleRequest

`func NewCardRuleRequest(conditions []CardRule, paymentServiceIds []string, ) *CardRuleRequest`

NewCardRuleRequest instantiates a new CardRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardRuleRequestWithDefaults

`func NewCardRuleRequestWithDefaults() *CardRuleRequest`

NewCardRuleRequestWithDefaults instantiates a new CardRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CardRuleRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CardRuleRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CardRuleRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CardRuleRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEnvironment

`func (o *CardRuleRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CardRuleRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CardRuleRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CardRuleRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetPosition

`func (o *CardRuleRequest) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CardRuleRequest) GetPositionOk() (*float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CardRuleRequest) SetPosition(v float32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CardRuleRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetConditions

`func (o *CardRuleRequest) GetConditions() []CardRule`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CardRuleRequest) GetConditionsOk() (*[]CardRule, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CardRuleRequest) SetConditions(v []CardRule)`

SetConditions sets Conditions field to given value.


### GetPaymentServiceIds

`func (o *CardRuleRequest) GetPaymentServiceIds() []string`

GetPaymentServiceIds returns the PaymentServiceIds field if non-nil, zero value otherwise.

### GetPaymentServiceIdsOk

`func (o *CardRuleRequest) GetPaymentServiceIdsOk() (*[]string, bool)`

GetPaymentServiceIdsOk returns a tuple with the PaymentServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceIds

`func (o *CardRuleRequest) SetPaymentServiceIds(v []string)`

SetPaymentServiceIds sets PaymentServiceIds field to given value.


### GetUnprocessableFallbackStrategy

`func (o *CardRuleRequest) GetUnprocessableFallbackStrategy() string`

GetUnprocessableFallbackStrategy returns the UnprocessableFallbackStrategy field if non-nil, zero value otherwise.

### GetUnprocessableFallbackStrategyOk

`func (o *CardRuleRequest) GetUnprocessableFallbackStrategyOk() (*string, bool)`

GetUnprocessableFallbackStrategyOk returns a tuple with the UnprocessableFallbackStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprocessableFallbackStrategy

`func (o *CardRuleRequest) SetUnprocessableFallbackStrategy(v string)`

SetUnprocessableFallbackStrategy sets UnprocessableFallbackStrategy field to given value.

### HasUnprocessableFallbackStrategy

`func (o *CardRuleRequest) HasUnprocessableFallbackStrategy() bool`

HasUnprocessableFallbackStrategy returns a boolean if a field has been set.

### GetInvalidRuleFallbackStrategy

`func (o *CardRuleRequest) GetInvalidRuleFallbackStrategy() string`

GetInvalidRuleFallbackStrategy returns the InvalidRuleFallbackStrategy field if non-nil, zero value otherwise.

### GetInvalidRuleFallbackStrategyOk

`func (o *CardRuleRequest) GetInvalidRuleFallbackStrategyOk() (*string, bool)`

GetInvalidRuleFallbackStrategyOk returns a tuple with the InvalidRuleFallbackStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidRuleFallbackStrategy

`func (o *CardRuleRequest) SetInvalidRuleFallbackStrategy(v string)`

SetInvalidRuleFallbackStrategy sets InvalidRuleFallbackStrategy field to given value.

### HasInvalidRuleFallbackStrategy

`func (o *CardRuleRequest) HasInvalidRuleFallbackStrategy() bool`

HasInvalidRuleFallbackStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


