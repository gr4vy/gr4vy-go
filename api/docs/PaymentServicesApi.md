# \PaymentServicesApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPaymentService**](PaymentServicesApi.md#AddPaymentService) | **Post** /payment-services | New payment service
[**DeletePaymentService**](PaymentServicesApi.md#DeletePaymentService) | **Delete** /payment-services/{payment_service_id} | Delete payment service
[**GetPaymentService**](PaymentServicesApi.md#GetPaymentService) | **Get** /payment-services/{payment_service_id} | Get payment service
[**ListPaymentServices**](PaymentServicesApi.md#ListPaymentServices) | **Get** /payment-services | List payment services
[**UpdatePaymentService**](PaymentServicesApi.md#UpdatePaymentService) | **Put** /payment-services/{payment_service_id} | Update payment service



## AddPaymentService

> PaymentService AddPaymentService(ctx).PaymentServiceRequest(paymentServiceRequest).Execute()

New payment service



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
    paymentServiceRequest := *openapiclient.NewPaymentServiceRequest("stripe-card", "Stripe (Main)", []openapiclient.PaymentServiceRequestFields{*openapiclient.NewPaymentServiceRequestFields("private_key", "sk_test_26PHem9AhJZvU623DfE1x4sd")}, []string{"AcceptedCountries_example"}, []string{"AcceptedCurrencies_example"}) // PaymentServiceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentServicesApi.AddPaymentService(context.Background()).PaymentServiceRequest(paymentServiceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentServicesApi.AddPaymentService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPaymentService`: PaymentService
    fmt.Fprintf(os.Stdout, "Response from `PaymentServicesApi.AddPaymentService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPaymentServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentServiceRequest** | [**PaymentServiceRequest**](PaymentServiceRequest.md) |  | 

### Return type

[**PaymentService**](PaymentService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePaymentService

> DeletePaymentService(ctx, paymentServiceId).Execute()

Delete payment service



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
    paymentServiceId := "46973e9d-88a7-44a6-abfe-be4ff0134ff4" // string | The ID of the payment service.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentServicesApi.DeletePaymentService(context.Background(), paymentServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentServicesApi.DeletePaymentService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentServiceId** | **string** | The ID of the payment service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePaymentServiceRequest struct via the builder pattern


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


## GetPaymentService

> PaymentService GetPaymentService(ctx, paymentServiceId).Execute()

Get payment service



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
    paymentServiceId := "46973e9d-88a7-44a6-abfe-be4ff0134ff4" // string | The ID of the payment service.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentServicesApi.GetPaymentService(context.Background(), paymentServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentServicesApi.GetPaymentService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentService`: PaymentService
    fmt.Fprintf(os.Stdout, "Response from `PaymentServicesApi.GetPaymentService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentServiceId** | **string** | The ID of the payment service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentService**](PaymentService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaymentServices

> PaymentServices ListPaymentServices(ctx).Limit(limit).Cursor(cursor).Method(method).Deleted(deleted).Execute()

List payment services



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
    limit := int32(1) // int32 | Defines the maximum number of items to return for this request. (optional) (default to 20)
    cursor := "ZXhhbXBsZTE" // string | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the `next_cursor` field. Similarly the `previous_cursor` can be used to reverse backwards in the list. (optional)
    method := "card" // string | Filters the results to only the items for which the `method` has been set to this value. (optional)
    deleted := true // bool | Filters the results to only show items which have been deleted. By default, deleted items will not be returned. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentServicesApi.ListPaymentServices(context.Background()).Limit(limit).Cursor(cursor).Method(method).Deleted(deleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentServicesApi.ListPaymentServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPaymentServices`: PaymentServices
    fmt.Fprintf(os.Stdout, "Response from `PaymentServicesApi.ListPaymentServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Defines the maximum number of items to return for this request. | [default to 20]
 **cursor** | **string** | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the &#x60;next_cursor&#x60; field. Similarly the &#x60;previous_cursor&#x60; can be used to reverse backwards in the list. | 
 **method** | **string** | Filters the results to only the items for which the &#x60;method&#x60; has been set to this value. | 
 **deleted** | **bool** | Filters the results to only show items which have been deleted. By default, deleted items will not be returned. | [default to false]

### Return type

[**PaymentServices**](PaymentServices.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePaymentService

> PaymentService UpdatePaymentService(ctx, paymentServiceId).PaymentServiceUpdate(paymentServiceUpdate).Execute()

Update payment service



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
    paymentServiceId := "46973e9d-88a7-44a6-abfe-be4ff0134ff4" // string | The ID of the payment service.
    paymentServiceUpdate := *openapiclient.NewPaymentServiceUpdate() // PaymentServiceUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentServicesApi.UpdatePaymentService(context.Background(), paymentServiceId).PaymentServiceUpdate(paymentServiceUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentServicesApi.UpdatePaymentService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePaymentService`: PaymentService
    fmt.Fprintf(os.Stdout, "Response from `PaymentServicesApi.UpdatePaymentService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentServiceId** | **string** | The ID of the payment service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaymentServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentServiceUpdate** | [**PaymentServiceUpdate**](PaymentServiceUpdate.md) |  | 

### Return type

[**PaymentService**](PaymentService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

