# \PayoutsApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSinglePayout**](PayoutsApi.md#GetSinglePayout) | **Get** /payouts/{payout_id} | Get payout
[**ListPayouts**](PayoutsApi.md#ListPayouts) | **Get** /payouts | List payouts
[**NewPayout**](PayoutsApi.md#NewPayout) | **Post** /payouts | Create payout



## GetSinglePayout

> Payout GetSinglePayout(ctx, payoutId).Execute()

Get payout



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
    payoutId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID for the payout to get the information for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayoutsApi.GetSinglePayout(context.Background(), payoutId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutsApi.GetSinglePayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSinglePayout`: Payout
    fmt.Fprintf(os.Stdout, "Response from `PayoutsApi.GetSinglePayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | **string** | The ID for the payout to get the information for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSinglePayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Payout**](Payout.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPayouts

> Payouts ListPayouts(ctx).Execute()

List payouts



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
    resp, r, err := api_client.PayoutsApi.ListPayouts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutsApi.ListPayouts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPayouts`: Payouts
    fmt.Fprintf(os.Stdout, "Response from `PayoutsApi.ListPayouts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPayoutsRequest struct via the builder pattern


### Return type

[**Payouts**](Payouts.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewPayout

> Payout NewPayout(ctx).PayoutRequest(payoutRequest).Execute()

Create payout



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
    payoutRequest := *openapiclient.NewPayoutRequest(float32(1299), "USD", "a7d6b829-aea5-407d-ab7f-138784b5ad2c", *openapiclient.NewPayoutPaymentMethodRequest("card")) // PayoutRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayoutsApi.NewPayout(context.Background()).PayoutRequest(payoutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutsApi.NewPayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewPayout`: Payout
    fmt.Fprintf(os.Stdout, "Response from `PayoutsApi.NewPayout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewPayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payoutRequest** | [**PayoutRequest**](PayoutRequest.md) |  | 

### Return type

[**Payout**](Payout.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

