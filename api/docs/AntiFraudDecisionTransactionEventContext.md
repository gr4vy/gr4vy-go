# AntiFraudDecisionTransactionEventContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to **string** | The HTTP body sent to fetch a decision. | [optional] 
**Response** | Pointer to **string** | The HTTP body received from the anti-fraud provider. | [optional] 
**ResponseStatusCode** | Pointer to **float32** | The HTTP response status code from the anti-fraud provider. | [optional] 
**Decision** | Pointer to **string** | The parsed decision response from the anti-fraud provider response. | [optional] 
**AntiFraudServiceId** | Pointer to **string** | The unique ID of the anti-fraud service used. | [optional] 
**AntiFraudServiceName** | Pointer to **string** | The name of the anti-fraud service used. | [optional] 
**AntiFraudServiceDefinitionId** | Pointer to **string** | The anti-fraud service definition used. | [optional] 

## Methods

### NewAntiFraudDecisionTransactionEventContext

`func NewAntiFraudDecisionTransactionEventContext() *AntiFraudDecisionTransactionEventContext`

NewAntiFraudDecisionTransactionEventContext instantiates a new AntiFraudDecisionTransactionEventContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntiFraudDecisionTransactionEventContextWithDefaults

`func NewAntiFraudDecisionTransactionEventContextWithDefaults() *AntiFraudDecisionTransactionEventContext`

NewAntiFraudDecisionTransactionEventContextWithDefaults instantiates a new AntiFraudDecisionTransactionEventContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *AntiFraudDecisionTransactionEventContext) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *AntiFraudDecisionTransactionEventContext) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *AntiFraudDecisionTransactionEventContext) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *AntiFraudDecisionTransactionEventContext) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResponse

`func (o *AntiFraudDecisionTransactionEventContext) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *AntiFraudDecisionTransactionEventContext) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *AntiFraudDecisionTransactionEventContext) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *AntiFraudDecisionTransactionEventContext) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetResponseStatusCode

`func (o *AntiFraudDecisionTransactionEventContext) GetResponseStatusCode() float32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *AntiFraudDecisionTransactionEventContext) GetResponseStatusCodeOk() (*float32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *AntiFraudDecisionTransactionEventContext) SetResponseStatusCode(v float32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *AntiFraudDecisionTransactionEventContext) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### GetDecision

`func (o *AntiFraudDecisionTransactionEventContext) GetDecision() string`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *AntiFraudDecisionTransactionEventContext) GetDecisionOk() (*string, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *AntiFraudDecisionTransactionEventContext) SetDecision(v string)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *AntiFraudDecisionTransactionEventContext) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetAntiFraudServiceId

`func (o *AntiFraudDecisionTransactionEventContext) GetAntiFraudServiceId() string`

GetAntiFraudServiceId returns the AntiFraudServiceId field if non-nil, zero value otherwise.

### GetAntiFraudServiceIdOk

`func (o *AntiFraudDecisionTransactionEventContext) GetAntiFraudServiceIdOk() (*string, bool)`

GetAntiFraudServiceIdOk returns a tuple with the AntiFraudServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceId

`func (o *AntiFraudDecisionTransactionEventContext) SetAntiFraudServiceId(v string)`

SetAntiFraudServiceId sets AntiFraudServiceId field to given value.

### HasAntiFraudServiceId

`func (o *AntiFraudDecisionTransactionEventContext) HasAntiFraudServiceId() bool`

HasAntiFraudServiceId returns a boolean if a field has been set.

### GetAntiFraudServiceName

`func (o *AntiFraudDecisionTransactionEventContext) GetAntiFraudServiceName() string`

GetAntiFraudServiceName returns the AntiFraudServiceName field if non-nil, zero value otherwise.

### GetAntiFraudServiceNameOk

`func (o *AntiFraudDecisionTransactionEventContext) GetAntiFraudServiceNameOk() (*string, bool)`

GetAntiFraudServiceNameOk returns a tuple with the AntiFraudServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceName

`func (o *AntiFraudDecisionTransactionEventContext) SetAntiFraudServiceName(v string)`

SetAntiFraudServiceName sets AntiFraudServiceName field to given value.

### HasAntiFraudServiceName

`func (o *AntiFraudDecisionTransactionEventContext) HasAntiFraudServiceName() bool`

HasAntiFraudServiceName returns a boolean if a field has been set.

### GetAntiFraudServiceDefinitionId

`func (o *AntiFraudDecisionTransactionEventContext) GetAntiFraudServiceDefinitionId() string`

GetAntiFraudServiceDefinitionId returns the AntiFraudServiceDefinitionId field if non-nil, zero value otherwise.

### GetAntiFraudServiceDefinitionIdOk

`func (o *AntiFraudDecisionTransactionEventContext) GetAntiFraudServiceDefinitionIdOk() (*string, bool)`

GetAntiFraudServiceDefinitionIdOk returns a tuple with the AntiFraudServiceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiFraudServiceDefinitionId

`func (o *AntiFraudDecisionTransactionEventContext) SetAntiFraudServiceDefinitionId(v string)`

SetAntiFraudServiceDefinitionId sets AntiFraudServiceDefinitionId field to given value.

### HasAntiFraudServiceDefinitionId

`func (o *AntiFraudDecisionTransactionEventContext) HasAntiFraudServiceDefinitionId() bool`

HasAntiFraudServiceDefinitionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


