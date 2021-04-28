# \BuyersApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBuyer**](BuyersApi.md#AddBuyer) | **Post** /buyers | New buyer
[**DeleteBuyer**](BuyersApi.md#DeleteBuyer) | **Delete** /buyers/{buyer_id} | Delete buyer
[**GetBuyer**](BuyersApi.md#GetBuyer) | **Get** /buyers/{buyer_id} | Get buyer
[**ListBuyers**](BuyersApi.md#ListBuyers) | **Get** /buyers | List buyers
[**UpdateBuyer**](BuyersApi.md#UpdateBuyer) | **Put** /buyers/{buyer_id} | Update buyer



## AddBuyer

> Buyer AddBuyer(ctx).BuyerRequest(buyerRequest).Execute()

New buyer



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
    buyerRequest := *openapiclient.NewBuyerRequest() // BuyerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuyersApi.AddBuyer(context.Background()).BuyerRequest(buyerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyersApi.AddBuyer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddBuyer`: Buyer
    fmt.Fprintf(os.Stdout, "Response from `BuyersApi.AddBuyer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddBuyerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buyerRequest** | [**BuyerRequest**](BuyerRequest.md) |  | 

### Return type

[**Buyer**](Buyer.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBuyer

> DeleteBuyer(ctx, buyerId).Execute()

Delete buyer



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
    buyerId := TODO // string | The unique ID for a buyer.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuyersApi.DeleteBuyer(context.Background(), buyerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyersApi.DeleteBuyer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buyerId** | [**string**](.md) | The unique ID for a buyer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuyer

> Buyer GetBuyer(ctx, buyerId).Execute()

Get buyer



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
    buyerId := TODO // string | The unique ID for a buyer.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuyersApi.GetBuyer(context.Background(), buyerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyersApi.GetBuyer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyer`: Buyer
    fmt.Fprintf(os.Stdout, "Response from `BuyersApi.GetBuyer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buyerId** | [**string**](.md) | The unique ID for a buyer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Buyer**](Buyer.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuyers

> Buyers ListBuyers(ctx).Search(search).Limit(limit).Cursor(cursor).Execute()

List buyers



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
    search := "John" // string | Filters the results to only the buyers for which the `display_name` or `external_identifier` matches this value. This field allows for a partial match, matching any buyer for which either of the fields partially or completely matches. (optional)
    limit := int32(1) // int32 | Defines the maximum number of items to return for this request. (optional) (default to 20)
    cursor := "ZXhhbXBsZTE" // string | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the `next_cursor` field. Similarly the `previous_cursor` can be used to reverse backwards in the list. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuyersApi.ListBuyers(context.Background()).Search(search).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyersApi.ListBuyers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyers`: Buyers
    fmt.Fprintf(os.Stdout, "Response from `BuyersApi.ListBuyers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBuyersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Filters the results to only the buyers for which the &#x60;display_name&#x60; or &#x60;external_identifier&#x60; matches this value. This field allows for a partial match, matching any buyer for which either of the fields partially or completely matches. | 
 **limit** | **int32** | Defines the maximum number of items to return for this request. | [default to 20]
 **cursor** | **string** | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the &#x60;next_cursor&#x60; field. Similarly the &#x60;previous_cursor&#x60; can be used to reverse backwards in the list. | 

### Return type

[**Buyers**](Buyers.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBuyer

> Buyer UpdateBuyer(ctx, buyerId).BuyerUpdate(buyerUpdate).Execute()

Update buyer



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
    buyerId := TODO // string | The unique ID for a buyer.
    buyerUpdate := *openapiclient.NewBuyerUpdate() // BuyerUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BuyersApi.UpdateBuyer(context.Background(), buyerId).BuyerUpdate(buyerUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyersApi.UpdateBuyer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBuyer`: Buyer
    fmt.Fprintf(os.Stdout, "Response from `BuyersApi.UpdateBuyer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buyerId** | [**string**](.md) | The unique ID for a buyer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBuyerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buyerUpdate** | [**BuyerUpdate**](BuyerUpdate.md) |  | 

### Return type

[**Buyer**](Buyer.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

