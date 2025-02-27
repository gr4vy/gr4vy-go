# \GiftCardsApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckGiftCardBalances**](GiftCardsApi.md#CheckGiftCardBalances) | **Post** /gift-cards/balances | Verify and check gift card balances
[**DeleteGiftCard**](GiftCardsApi.md#DeleteGiftCard) | **Delete** /gift-cards/{gift_card_id} | Delete gift card
[**GetGiftCard**](GiftCardsApi.md#GetGiftCard) | **Get** /gift-cards/{gift_card_id} | Get gift card
[**ListBuyerGiftCards**](GiftCardsApi.md#ListBuyerGiftCards) | **Get** /buyers/gift-cards | List gift cards for buyer
[**ListGiftCards**](GiftCardsApi.md#ListGiftCards) | **Get** /gift-cards | List gift cards
[**StoreGiftCard**](GiftCardsApi.md#StoreGiftCard) | **Post** /gift-cards | Store gift card



## CheckGiftCardBalances

> GiftCardsSummary CheckGiftCardBalances(ctx).GiftCardBalancesRequest(giftCardBalancesRequest).Execute()

Verify and check gift card balances



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    giftCardBalancesRequest := *openapiclient.NewGiftCardBalancesRequest() // GiftCardBalancesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GiftCardsApi.CheckGiftCardBalances(context.Background()).GiftCardBalancesRequest(giftCardBalancesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardsApi.CheckGiftCardBalances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckGiftCardBalances`: GiftCardsSummary
    fmt.Fprintf(os.Stdout, "Response from `GiftCardsApi.CheckGiftCardBalances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckGiftCardBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **giftCardBalancesRequest** | [**GiftCardBalancesRequest**](GiftCardBalancesRequest.md) |  | 

### Return type

[**GiftCardsSummary**](GiftCardsSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGiftCard

> DeleteGiftCard(ctx, giftCardId).Execute()

Delete gift card



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    giftCardId := TODO // string | The unique ID of a stored gift card.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GiftCardsApi.DeleteGiftCard(context.Background(), giftCardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardsApi.DeleteGiftCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardId** | [**string**](.md) | The unique ID of a stored gift card. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGiftCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGiftCard

> GiftCard GetGiftCard(ctx, giftCardId).Execute()

Get gift card



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    giftCardId := TODO // string | The unique ID of a stored gift card.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GiftCardsApi.GetGiftCard(context.Background(), giftCardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardsApi.GetGiftCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGiftCard`: GiftCard
    fmt.Fprintf(os.Stdout, "Response from `GiftCardsApi.GetGiftCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardId** | [**string**](.md) | The unique ID of a stored gift card. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGiftCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GiftCard**](GiftCard.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuyerGiftCards

> GiftCardsSummary ListBuyerGiftCards(ctx).BuyerId(buyerId).BuyerExternalIdentifier(buyerExternalIdentifier).Execute()

List gift cards for buyer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    buyerId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | Filters the results to only the items for which the `buyer` has an `id` that matches this value. (optional)
    buyerExternalIdentifier := "user-12345" // string | Filters the results to only the items for which the `buyer` has an `external_identifier` that matches this value. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GiftCardsApi.ListBuyerGiftCards(context.Background()).BuyerId(buyerId).BuyerExternalIdentifier(buyerExternalIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardsApi.ListBuyerGiftCards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerGiftCards`: GiftCardsSummary
    fmt.Fprintf(os.Stdout, "Response from `GiftCardsApi.ListBuyerGiftCards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerGiftCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buyerId** | **string** | Filters the results to only the items for which the &#x60;buyer&#x60; has an &#x60;id&#x60; that matches this value. | 
 **buyerExternalIdentifier** | **string** | Filters the results to only the items for which the &#x60;buyer&#x60; has an &#x60;external_identifier&#x60; that matches this value. | 

### Return type

[**GiftCardsSummary**](GiftCardsSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGiftCards

> GiftCards ListGiftCards(ctx).BuyerId(buyerId).BuyerExternalIdentifier(buyerExternalIdentifier).Limit(limit).Cursor(cursor).Execute()

List gift cards



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    buyerId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | Filters the results to only the items for which the `buyer` has an `id` that matches this value. (optional)
    buyerExternalIdentifier := "user-12345" // string | Filters the results to only the items for which the `buyer` has an `external_identifier` that matches this value. (optional)
    limit := int32(1) // int32 | Defines the maximum number of items to return for this request. (optional) (default to 20)
    cursor := "ZXhhbXBsZTE" // string | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the `next_cursor` field. Similarly the `previous_cursor` can be used to reverse backwards in the list. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GiftCardsApi.ListGiftCards(context.Background()).BuyerId(buyerId).BuyerExternalIdentifier(buyerExternalIdentifier).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardsApi.ListGiftCards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGiftCards`: GiftCards
    fmt.Fprintf(os.Stdout, "Response from `GiftCardsApi.ListGiftCards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGiftCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buyerId** | **string** | Filters the results to only the items for which the &#x60;buyer&#x60; has an &#x60;id&#x60; that matches this value. | 
 **buyerExternalIdentifier** | **string** | Filters the results to only the items for which the &#x60;buyer&#x60; has an &#x60;external_identifier&#x60; that matches this value. | 
 **limit** | **int32** | Defines the maximum number of items to return for this request. | [default to 20]
 **cursor** | **string** | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the &#x60;next_cursor&#x60; field. Similarly the &#x60;previous_cursor&#x60; can be used to reverse backwards in the list. | 

### Return type

[**GiftCards**](GiftCards.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreGiftCard

> GiftCard StoreGiftCard(ctx).GiftCardStoreRequest(giftCardStoreRequest).Execute()

Store gift card



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    giftCardStoreRequest := *openapiclient.NewGiftCardStoreRequest("4123455541234561234", "1234") // GiftCardStoreRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GiftCardsApi.StoreGiftCard(context.Background()).GiftCardStoreRequest(giftCardStoreRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardsApi.StoreGiftCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreGiftCard`: GiftCard
    fmt.Fprintf(os.Stdout, "Response from `GiftCardsApi.StoreGiftCard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreGiftCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **giftCardStoreRequest** | [**GiftCardStoreRequest**](GiftCardStoreRequest.md) |  | 

### Return type

[**GiftCard**](GiftCard.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

