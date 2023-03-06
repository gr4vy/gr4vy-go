# GooglePayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | &#x60;googlepay&#x60;. | 
**Token** | **map[string]interface{}** | The encrypted (opaque) token returned by the Google Pay API that represents a payment method. | 
**AssuranceDetails** | Pointer to [**NullableGooglePayRequestAssuranceDetails**](GooglePayRequestAssuranceDetails.md) |  | [optional] 
**CardHolderName** | Pointer to **NullableString** | Name of the card holder. | [optional] 
**RedirectUrl** | Pointer to **NullableString** | The redirect URL to redirect a buyer to after they have authorized their transaction or payment method. This only applies to payment methods that require buyer approval. | [optional] 

## Methods

### NewGooglePayRequest

`func NewGooglePayRequest(method string, token map[string]interface{}, ) *GooglePayRequest`

NewGooglePayRequest instantiates a new GooglePayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGooglePayRequestWithDefaults

`func NewGooglePayRequestWithDefaults() *GooglePayRequest`

NewGooglePayRequestWithDefaults instantiates a new GooglePayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *GooglePayRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *GooglePayRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *GooglePayRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetToken

`func (o *GooglePayRequest) GetToken() map[string]interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GooglePayRequest) GetTokenOk() (*map[string]interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GooglePayRequest) SetToken(v map[string]interface{})`

SetToken sets Token field to given value.


### GetAssuranceDetails

`func (o *GooglePayRequest) GetAssuranceDetails() GooglePayRequestAssuranceDetails`

GetAssuranceDetails returns the AssuranceDetails field if non-nil, zero value otherwise.

### GetAssuranceDetailsOk

`func (o *GooglePayRequest) GetAssuranceDetailsOk() (*GooglePayRequestAssuranceDetails, bool)`

GetAssuranceDetailsOk returns a tuple with the AssuranceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceDetails

`func (o *GooglePayRequest) SetAssuranceDetails(v GooglePayRequestAssuranceDetails)`

SetAssuranceDetails sets AssuranceDetails field to given value.

### HasAssuranceDetails

`func (o *GooglePayRequest) HasAssuranceDetails() bool`

HasAssuranceDetails returns a boolean if a field has been set.

### SetAssuranceDetailsNil

`func (o *GooglePayRequest) SetAssuranceDetailsNil(b bool)`

 SetAssuranceDetailsNil sets the value for AssuranceDetails to be an explicit nil

### UnsetAssuranceDetails
`func (o *GooglePayRequest) UnsetAssuranceDetails()`

UnsetAssuranceDetails ensures that no value is present for AssuranceDetails, not even an explicit nil
### GetCardHolderName

`func (o *GooglePayRequest) GetCardHolderName() string`

GetCardHolderName returns the CardHolderName field if non-nil, zero value otherwise.

### GetCardHolderNameOk

`func (o *GooglePayRequest) GetCardHolderNameOk() (*string, bool)`

GetCardHolderNameOk returns a tuple with the CardHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHolderName

`func (o *GooglePayRequest) SetCardHolderName(v string)`

SetCardHolderName sets CardHolderName field to given value.

### HasCardHolderName

`func (o *GooglePayRequest) HasCardHolderName() bool`

HasCardHolderName returns a boolean if a field has been set.

### SetCardHolderNameNil

`func (o *GooglePayRequest) SetCardHolderNameNil(b bool)`

 SetCardHolderNameNil sets the value for CardHolderName to be an explicit nil

### UnsetCardHolderName
`func (o *GooglePayRequest) UnsetCardHolderName()`

UnsetCardHolderName ensures that no value is present for CardHolderName, not even an explicit nil
### GetRedirectUrl

`func (o *GooglePayRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *GooglePayRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *GooglePayRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *GooglePayRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *GooglePayRequest) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *GooglePayRequest) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


