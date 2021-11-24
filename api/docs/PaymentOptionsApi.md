# \PaymentOptionsApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPaymentOptions**](PaymentOptionsApi.md#ListPaymentOptions) | **Get** /payment-options | List payment options



## ListPaymentOptions

> PaymentOptions ListPaymentOptions(ctx).Country(country).Currency(currency).Locale(locale).Execute()

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
    locale := "en-US" // string | An ISO 639-1 Language Code and optional ISO 3166 Country Code. This locale determines the language for the labels returned for every payment option. (optional) (default to "en-US")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentOptionsApi.ListPaymentOptions(context.Background()).Country(country).Currency(currency).Locale(locale).Execute()
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
 **locale** | **string** | An ISO 639-1 Language Code and optional ISO 3166 Country Code. This locale determines the language for the labels returned for every payment option. | [default to &quot;en-US&quot;]

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

