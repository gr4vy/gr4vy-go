# CardDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suffix** | Pointer to **string** | The last 4 digits of the card number. | [optional] 
**Scheme** | Pointer to **string** | The type of the card. | [optional] 
**ExpirationDate** | Pointer to **string** | The expiration date for a card. | [optional] 

## Methods

### NewCardDetails

`func NewCardDetails() *CardDetails`

NewCardDetails instantiates a new CardDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardDetailsWithDefaults

`func NewCardDetailsWithDefaults() *CardDetails`

NewCardDetailsWithDefaults instantiates a new CardDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuffix

`func (o *CardDetails) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *CardDetails) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *CardDetails) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *CardDetails) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### GetScheme

`func (o *CardDetails) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *CardDetails) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *CardDetails) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *CardDetails) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CardDetails) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CardDetails) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CardDetails) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CardDetails) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


