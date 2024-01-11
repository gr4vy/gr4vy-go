# \GiftCardServicesApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteGiftCardService**](GiftCardServicesApi.md#DeleteGiftCardService) | **Delete** /gift-card-services/{gift_card_service_id} | Delete gift card service
[**GetGiftCardService**](GiftCardServicesApi.md#GetGiftCardService) | **Get** /gift-card-services/{gift_card_service_id} | Get gift card service
[**NewGiftCardService**](GiftCardServicesApi.md#NewGiftCardService) | **Post** /gift-card-services | New gift card service
[**UpdateGiftCardService**](GiftCardServicesApi.md#UpdateGiftCardService) | **Put** /gift-card-services/{gift_card_service_id} | Update gift card service
[**VerifyGiftCardService**](GiftCardServicesApi.md#VerifyGiftCardService) | **Post** /gift-card-services/verify | Verify gift card service credentials



## DeleteGiftCardService

> DeleteGiftCardService(ctx, giftCardServiceId).Execute()

Delete gift card service



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
    giftCardServiceId := TODO // string | The unique ID of the gift card service.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GiftCardServicesApi.DeleteGiftCardService(context.Background(), giftCardServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardServicesApi.DeleteGiftCardService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardServiceId** | [**string**](.md) | The unique ID of the gift card service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGiftCardServiceRequest struct via the builder pattern


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


## GetGiftCardService

> GiftCardService GetGiftCardService(ctx, giftCardServiceId).Execute()

Get gift card service



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
    giftCardServiceId := TODO // string | The unique ID of the gift card service.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GiftCardServicesApi.GetGiftCardService(context.Background(), giftCardServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardServicesApi.GetGiftCardService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGiftCardService`: GiftCardService
    fmt.Fprintf(os.Stdout, "Response from `GiftCardServicesApi.GetGiftCardService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardServiceId** | [**string**](.md) | The unique ID of the gift card service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGiftCardServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GiftCardService**](GiftCardService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewGiftCardService

> GiftCardService NewGiftCardService(ctx).GiftCardServiceCreateRequest(giftCardServiceCreateRequest).Execute()

New gift card service



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
    giftCardServiceCreateRequest := *openapiclient.NewGiftCardServiceCreateRequest("qwikcilver-gift-card", "Qwikcilver UK", []openapiclient.GiftCardServiceCreateRequestFields{*openapiclient.NewGiftCardServiceCreateRequestFields("private_key", "pk_26PHem9AhJZvU623DfE1x4sd")}) // GiftCardServiceCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GiftCardServicesApi.NewGiftCardService(context.Background()).GiftCardServiceCreateRequest(giftCardServiceCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardServicesApi.NewGiftCardService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewGiftCardService`: GiftCardService
    fmt.Fprintf(os.Stdout, "Response from `GiftCardServicesApi.NewGiftCardService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewGiftCardServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **giftCardServiceCreateRequest** | [**GiftCardServiceCreateRequest**](GiftCardServiceCreateRequest.md) |  | 

### Return type

[**GiftCardService**](GiftCardService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGiftCardService

> GiftCardService UpdateGiftCardService(ctx, giftCardServiceId).GiftCardServiceUpdateRequest(giftCardServiceUpdateRequest).Execute()

Update gift card service



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
    giftCardServiceId := TODO // string | The unique ID of the gift card service.
    giftCardServiceUpdateRequest := *openapiclient.NewGiftCardServiceUpdateRequest() // GiftCardServiceUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GiftCardServicesApi.UpdateGiftCardService(context.Background(), giftCardServiceId).GiftCardServiceUpdateRequest(giftCardServiceUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardServicesApi.UpdateGiftCardService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGiftCardService`: GiftCardService
    fmt.Fprintf(os.Stdout, "Response from `GiftCardServicesApi.UpdateGiftCardService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardServiceId** | [**string**](.md) | The unique ID of the gift card service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGiftCardServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **giftCardServiceUpdateRequest** | [**GiftCardServiceUpdateRequest**](GiftCardServiceUpdateRequest.md) |  | 

### Return type

[**GiftCardService**](GiftCardService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyGiftCardService

> VerifyGiftCardService(ctx).GiftCardServiceVerifyRequest(giftCardServiceVerifyRequest).Execute()

Verify gift card service credentials



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
    giftCardServiceVerifyRequest := *openapiclient.NewGiftCardServiceVerifyRequest("qwikcilver-gift-card", []openapiclient.GiftCardServiceVerifyRequestFields{*openapiclient.NewGiftCardServiceVerifyRequestFields("private_key", "pk_26PHem9AhJZvU623DfE1x4sd")}) // GiftCardServiceVerifyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GiftCardServicesApi.VerifyGiftCardService(context.Background()).GiftCardServiceVerifyRequest(giftCardServiceVerifyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GiftCardServicesApi.VerifyGiftCardService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyGiftCardServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **giftCardServiceVerifyRequest** | [**GiftCardServiceVerifyRequest**](GiftCardServiceVerifyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

