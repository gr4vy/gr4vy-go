# \AntiFraudServicesApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAntiFraudService**](AntiFraudServicesApi.md#AddAntiFraudService) | **Post** /anti-fraud-services | New anti-fraud service
[**DeleteAntiFraudService**](AntiFraudServicesApi.md#DeleteAntiFraudService) | **Delete** /anti-fraud-services/{anti_fraud_service_id} | Delete anti-fraud service
[**GetAntiFraudService**](AntiFraudServicesApi.md#GetAntiFraudService) | **Get** /anti-fraud-services/{anti_fraud_service_id} | Get anti-fraud service
[**UpdateAntiFraudService**](AntiFraudServicesApi.md#UpdateAntiFraudService) | **Put** /anti-fraud-services/{anti_fraud_service_id} | Update anti-fraud service



## AddAntiFraudService

> AntiFraudService AddAntiFraudService(ctx).AntiFraudServiceCreate(antiFraudServiceCreate).Execute()

New anti-fraud service



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
    antiFraudServiceCreate := *openapiclient.NewAntiFraudServiceCreate("sift", "Sift Anti-Fraud Service.", []openapiclient.AntiFraudServiceUpdateFields{*openapiclient.NewAntiFraudServiceUpdateFields("api_key", "sk_test_26PHem9AhJZvU623DfE1x4sd")}) // AntiFraudServiceCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AntiFraudServicesApi.AddAntiFraudService(context.Background()).AntiFraudServiceCreate(antiFraudServiceCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntiFraudServicesApi.AddAntiFraudService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAntiFraudService`: AntiFraudService
    fmt.Fprintf(os.Stdout, "Response from `AntiFraudServicesApi.AddAntiFraudService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAntiFraudServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **antiFraudServiceCreate** | [**AntiFraudServiceCreate**](AntiFraudServiceCreate.md) |  | 

### Return type

[**AntiFraudService**](AntiFraudService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAntiFraudService

> DeleteAntiFraudService(ctx, antiFraudServiceId).Execute()

Delete anti-fraud service



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
    antiFraudServiceId := TODO // string | The unique ID for an anti-fraud service.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AntiFraudServicesApi.DeleteAntiFraudService(context.Background(), antiFraudServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntiFraudServicesApi.DeleteAntiFraudService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**antiFraudServiceId** | [**string**](.md) | The unique ID for an anti-fraud service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAntiFraudServiceRequest struct via the builder pattern


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


## GetAntiFraudService

> AntiFraudService GetAntiFraudService(ctx, antiFraudServiceId).Execute()

Get anti-fraud service



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
    antiFraudServiceId := TODO // string | The unique ID for an anti-fraud service.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AntiFraudServicesApi.GetAntiFraudService(context.Background(), antiFraudServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntiFraudServicesApi.GetAntiFraudService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAntiFraudService`: AntiFraudService
    fmt.Fprintf(os.Stdout, "Response from `AntiFraudServicesApi.GetAntiFraudService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**antiFraudServiceId** | [**string**](.md) | The unique ID for an anti-fraud service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAntiFraudServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AntiFraudService**](AntiFraudService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAntiFraudService

> AntiFraudService UpdateAntiFraudService(ctx, antiFraudServiceId).AntiFraudServiceUpdate(antiFraudServiceUpdate).Execute()

Update anti-fraud service



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
    antiFraudServiceId := TODO // string | The unique ID for an anti-fraud service.
    antiFraudServiceUpdate := *openapiclient.NewAntiFraudServiceUpdate("sift") // AntiFraudServiceUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AntiFraudServicesApi.UpdateAntiFraudService(context.Background(), antiFraudServiceId).AntiFraudServiceUpdate(antiFraudServiceUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntiFraudServicesApi.UpdateAntiFraudService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAntiFraudService`: AntiFraudService
    fmt.Fprintf(os.Stdout, "Response from `AntiFraudServicesApi.UpdateAntiFraudService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**antiFraudServiceId** | [**string**](.md) | The unique ID for an anti-fraud service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAntiFraudServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **antiFraudServiceUpdate** | [**AntiFraudServiceUpdate**](AntiFraudServiceUpdate.md) |  | 

### Return type

[**AntiFraudService**](AntiFraudService.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

