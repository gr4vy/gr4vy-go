# \PaymentOptionsApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPaymentOptions**](PaymentOptionsApi.md#ListPaymentOptions) | **Get** /payment-options | List payment options



## ListPaymentOptions

> PaymentOptions ListPaymentOptions(ctx).Country(country).Currency(currency).Environment(environment).Execute()

List payment options



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
    country := "US" // string | Filters the results to only the items which support this country code. A country is formatted as 2-letter ISO country code. (optional)
    currency := "USD" // string | Filters the results to only the items which support this currency code. A currency is formatted as 3-letter ISO currency code. (optional)
    environment := "staging" // string | Filters the results to only the items available in this environment. (optional) (default to "production")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentOptionsApi.ListPaymentOptions(context.Background()).Country(country).Currency(currency).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentOptionsApi.ListPaymentOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPaymentOptions`: PaymentOptions
    fmt.Fprintf(os.Stdout, "Response from `PaymentOptionsApi.ListPaymentOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | Filters the results to only the items which support this country code. A country is formatted as 2-letter ISO country code. | 
 **currency** | **string** | Filters the results to only the items which support this currency code. A currency is formatted as 3-letter ISO currency code. | 
 **environment** | **string** | Filters the results to only the items available in this environment. | [default to &quot;production&quot;]

### Return type

[**PaymentOptions**](PaymentOptions.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

