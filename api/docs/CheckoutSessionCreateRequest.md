# CheckoutSessionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) | An array of cart items that represents the line items of a transaction. | [optional] 
**Metadata** | Pointer to **map[string]string** | Any additional information about the transaction that you would like to store as key-value pairs. This data is passed to payment service providers that support it. | [optional] 

## Methods

### NewCheckoutSessionCreateRequest

`func NewCheckoutSessionCreateRequest() *CheckoutSessionCreateRequest`

NewCheckoutSessionCreateRequest instantiates a new CheckoutSessionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutSessionCreateRequestWithDefaults

`func NewCheckoutSessionCreateRequestWithDefaults() *CheckoutSessionCreateRequest`

NewCheckoutSessionCreateRequestWithDefaults instantiates a new CheckoutSessionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCartItems

`func (o *CheckoutSessionCreateRequest) GetCartItems() []CartItem`

GetCartItems returns the CartItems field if non-nil, zero value otherwise.

### GetCartItemsOk

`func (o *CheckoutSessionCreateRequest) GetCartItemsOk() (*[]CartItem, bool)`

GetCartItemsOk returns a tuple with the CartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItems

`func (o *CheckoutSessionCreateRequest) SetCartItems(v []CartItem)`

SetCartItems sets CartItems field to given value.

### HasCartItems

`func (o *CheckoutSessionCreateRequest) HasCartItems() bool`

HasCartItems returns a boolean if a field has been set.

### SetCartItemsNil

`func (o *CheckoutSessionCreateRequest) SetCartItemsNil(b bool)`

 SetCartItemsNil sets the value for CartItems to be an explicit nil

### UnsetCartItems
`func (o *CheckoutSessionCreateRequest) UnsetCartItems()`

UnsetCartItems ensures that no value is present for CartItems, not even an explicit nil
### GetMetadata

`func (o *CheckoutSessionCreateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CheckoutSessionCreateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CheckoutSessionCreateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CheckoutSessionCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CheckoutSessionCreateRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CheckoutSessionCreateRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


