# \PaymentMethodTokensApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPaymentMethodTokens**](PaymentMethodTokensApi.md#ListPaymentMethodTokens) | **Get** /payment-methods/{payment_method_id}/tokens | List payment method tokens



## ListPaymentMethodTokens

> PaymentMethodTokens ListPaymentMethodTokens(ctx, paymentMethodId).Execute()

List payment method tokens



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
    paymentMethodId := TODO // string | The ID of the payment method.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentMethodTokensApi.ListPaymentMethodTokens(context.Background(), paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodTokensApi.ListPaymentMethodTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPaymentMethodTokens`: PaymentMethodTokens
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodTokensApi.ListPaymentMethodTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | [**string**](.md) | The ID of the payment method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentMethodTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentMethodTokens**](PaymentMethodTokens.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

