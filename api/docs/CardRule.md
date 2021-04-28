# CardRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;card-rule&#x60;. | [optional] 
**Id** | Pointer to **string** | The ID of the rule. | [optional] 
**Active** | Pointer to **bool** | Whether this rule is currently in use. Rules can be deactivated to allow for them to be kept around and re-activated at a later date. | [optional] 
**Environment** | Pointer to **string** | The environment to use this rule in. This rule will only be used for transactions created in that environment. | [optional] [default to "production"]
**Conditions** | Pointer to [**[]CardRuleCondition**](CardRuleCondition.md) | One or more conditions that apply for this rule. Each condition needs to match for this rule to go into effect. | [optional] 
**PaymentServiceIds** | Pointer to **[]string** | A list of IDs for the payment services to use, in order of priority. The payment services all need to process cards. | [optional] 
**Position** | Pointer to **float32** | The numeric rank of a rule. Rules with a lower position value are processed first. | [optional] 
**UnprocessableFallbackStrategy** | Pointer to **string** | Defines what strategy to use when all of the payment services defined in this rule declined or otherwise were not able to process the card.  * &#x60;use_all_providers&#x60; - Try all payment services enabled for this currency in order of priority, even if they are not listed in this rule. This is the default behaviour for a rule. * &#x60;decline&#x60; - Decline the transaction. | [optional] [default to "use_all_providers"]
**InvalidRuleFallbackStrategy** | Pointer to **string** | Defines what strategy to use when this rule is not valid. This can happen when the rule has triggered for a certain transaction but none of the listed payment services are eligible to process that transaction currency.  * &#x60;use_all_providers&#x60; - Try all payment services enabled for this currency in order of priority, even if they are not listed in this rule. This is the default behaviour for a rule. * &#x60;skip&#x60; - Skip this rule and instead move on to the next highest priority rule. * &#x60;decline&#x60; - Decline the transaction. | [optional] [default to "use_all_providers"]
**CreatedAt** | Pointer to **time.Time** | The date and time when this rule was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when this rule was last updated. | [optional] 

## Methods

### NewCardRule

`func NewCardRule() *CardRule`

NewCardRule instantiates a new CardRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardRuleWithDefaults

`func NewCardRuleWithDefaults() *CardRule`

NewCardRuleWithDefaults instantiates a new CardRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CardRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CardRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CardRule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CardRule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *CardRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CardRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActive

`func (o *CardRule) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CardRule) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CardRule) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CardRule) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEnvironment

`func (o *CardRule) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CardRule) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CardRule) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CardRule) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetConditions

`func (o *CardRule) GetConditions() []CardRuleCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CardRule) GetConditionsOk() (*[]CardRuleCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CardRule) SetConditions(v []CardRuleCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *CardRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetPaymentServiceIds

`func (o *CardRule) GetPaymentServiceIds() []string`

GetPaymentServiceIds returns the PaymentServiceIds field if non-nil, zero value otherwise.

### GetPaymentServiceIdsOk

`func (o *CardRule) GetPaymentServiceIdsOk() (*[]string, bool)`

GetPaymentServiceIdsOk returns a tuple with the PaymentServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceIds

`func (o *CardRule) SetPaymentServiceIds(v []string)`

SetPaymentServiceIds sets PaymentServiceIds field to given value.

### HasPaymentServiceIds

`func (o *CardRule) HasPaymentServiceIds() bool`

HasPaymentServiceIds returns a boolean if a field has been set.

### GetPosition

`func (o *CardRule) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CardRule) GetPositionOk() (*float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CardRule) SetPosition(v float32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CardRule) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetUnprocessableFallbackStrategy

`func (o *CardRule) GetUnprocessableFallbackStrategy() string`

GetUnprocessableFallbackStrategy returns the UnprocessableFallbackStrategy field if non-nil, zero value otherwise.

### GetUnprocessableFallbackStrategyOk

`func (o *CardRule) GetUnprocessableFallbackStrategyOk() (*string, bool)`

GetUnprocessableFallbackStrategyOk returns a tuple with the UnprocessableFallbackStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprocessableFallbackStrategy

`func (o *CardRule) SetUnprocessableFallbackStrategy(v string)`

SetUnprocessableFallbackStrategy sets UnprocessableFallbackStrategy field to given value.

### HasUnprocessableFallbackStrategy

`func (o *CardRule) HasUnprocessableFallbackStrategy() bool`

HasUnprocessableFallbackStrategy returns a boolean if a field has been set.

### GetInvalidRuleFallbackStrategy

`func (o *CardRule) GetInvalidRuleFallbackStrategy() string`

GetInvalidRuleFallbackStrategy returns the InvalidRuleFallbackStrategy field if non-nil, zero value otherwise.

### GetInvalidRuleFallbackStrategyOk

`func (o *CardRule) GetInvalidRuleFallbackStrategyOk() (*string, bool)`

GetInvalidRuleFallbackStrategyOk returns a tuple with the InvalidRuleFallbackStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidRuleFallbackStrategy

`func (o *CardRule) SetInvalidRuleFallbackStrategy(v string)`

SetInvalidRuleFallbackStrategy sets InvalidRuleFallbackStrategy field to given value.

### HasInvalidRuleFallbackStrategy

`func (o *CardRule) HasInvalidRuleFallbackStrategy() bool`

HasInvalidRuleFallbackStrategy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CardRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CardRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CardRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CardRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CardRule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CardRule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CardRule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CardRule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


