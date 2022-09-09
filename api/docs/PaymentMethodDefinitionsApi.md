# \PaymentMethodDefinitionsApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPaymentMethodDefinitions**](PaymentMethodDefinitionsApi.md#ListPaymentMethodDefinitions) | **Get** /payment-method-definitions | List payment method definitions



## ListPaymentMethodDefinitions

> PaymentMethodDefinitions ListPaymentMethodDefinitions(ctx).Execute()

List payment method definitions



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
    resp, r, err := api_client.PaymentMethodDefinitionsApi.ListPaymentMethodDefinitions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodDefinitionsApi.ListPaymentMethodDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPaymentMethodDefinitions`: PaymentMethodDefinitions
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodDefinitionsApi.ListPaymentMethodDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentMethodDefinitionsRequest struct via the builder pattern


### Return type

[**PaymentMethodDefinitions**](PaymentMethodDefinitions.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

