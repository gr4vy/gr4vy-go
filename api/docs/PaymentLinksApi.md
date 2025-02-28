# \PaymentLinksApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExpirePaymentLink**](PaymentLinksApi.md#ExpirePaymentLink) | **Post** /payment-links/{payment_link_id}/expire | Expire payment link
[**GetPaymentLink**](PaymentLinksApi.md#GetPaymentLink) | **Get** /payment-links/{payment_link_id} | Get payment link
[**ListPaymentLinks**](PaymentLinksApi.md#ListPaymentLinks) | **Get** /payment-links | List payment links
[**NewPaymentLink**](PaymentLinksApi.md#NewPaymentLink) | **Post** /payment-links | Create payment link



## ExpirePaymentLink

> ExpirePaymentLink(ctx, paymentLinkId).Execute()

Expire payment link



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
    paymentLinkId := TODO // string | The ID of the payment link to expire.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentLinksApi.ExpirePaymentLink(context.Background(), paymentLinkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksApi.ExpirePaymentLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentLinkId** | [**string**](.md) | The ID of the payment link to expire. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpirePaymentLinkRequest struct via the builder pattern


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


## GetPaymentLink

> PaymentLink GetPaymentLink(ctx, paymentLinkId).Execute()

Get payment link



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
    paymentLinkId := "1b2e1e3d-ec3e-4884-9553-f7356683f7f9" // string | The ID for the payment link to get the information for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentLinksApi.GetPaymentLink(context.Background(), paymentLinkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksApi.GetPaymentLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentLink`: PaymentLink
    fmt.Fprintf(os.Stdout, "Response from `PaymentLinksApi.GetPaymentLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentLinkId** | **string** | The ID for the payment link to get the information for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentLink**](PaymentLink.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaymentLinks

> PaymentLinks ListPaymentLinks(ctx).Execute()

List payment links



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentLinksApi.ListPaymentLinks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksApi.ListPaymentLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPaymentLinks`: PaymentLinks
    fmt.Fprintf(os.Stdout, "Response from `PaymentLinksApi.ListPaymentLinks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentLinksRequest struct via the builder pattern


### Return type

[**PaymentLinks**](PaymentLinks.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewPaymentLink

> PaymentLink NewPaymentLink(ctx).PaymentLinkRequest(paymentLinkRequest).Execute()

Create payment link



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
    paymentLinkRequest := *openapiclient.NewPaymentLinkRequest(float32(1299), "USD", "US") // PaymentLinkRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentLinksApi.NewPaymentLink(context.Background()).PaymentLinkRequest(paymentLinkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksApi.NewPaymentLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewPaymentLink`: PaymentLink
    fmt.Fprintf(os.Stdout, "Response from `PaymentLinksApi.NewPaymentLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewPaymentLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentLinkRequest** | [**PaymentLinkRequest**](PaymentLinkRequest.md) |  | 

### Return type

[**PaymentLink**](PaymentLink.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

