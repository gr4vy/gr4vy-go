# CardTokenized

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | &#x60;payment-method&#x60;. | [optional] 
**Id** | Pointer to **string** | The unique ID of the payment method. | [optional] 
**Method** | Pointer to **string** | &#x60;card&#x60;. | [optional] 
**Details** | Pointer to [**CardDetails**](CardDetails.md) |  | [optional] 

## Methods

### NewCardTokenized

`func NewCardTokenized() *CardTokenized`

NewCardTokenized instantiates a new CardTokenized object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardTokenizedWithDefaults

`func NewCardTokenizedWithDefaults() *CardTokenized`

NewCardTokenizedWithDefaults instantiates a new CardTokenized object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CardTokenized) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CardTokenized) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CardTokenized) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CardTokenized) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *CardTokenized) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardTokenized) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardTokenized) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CardTokenized) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethod

`func (o *CardTokenized) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CardTokenized) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CardTokenized) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *CardTokenized) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetDetails

`func (o *CardTokenized) GetDetails() CardDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CardTokenized) GetDetailsOk() (*CardDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CardTokenized) SetDetails(v CardDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *CardTokenized) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


