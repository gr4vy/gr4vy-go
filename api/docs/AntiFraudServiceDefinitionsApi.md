# \AntiFraudServiceDefinitionsApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAntiFraudServiceDefinition**](AntiFraudServiceDefinitionsApi.md#GetAntiFraudServiceDefinition) | **Get** /anti-fraud-service-definitions/{anti_fraud_service_definition_id} | Get anti fraud service definition



## GetAntiFraudServiceDefinition

> AntiFraudServiceDefinition GetAntiFraudServiceDefinition(ctx, antiFraudServiceDefinitionId).Execute()

Get anti fraud service definition



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
    antiFraudServiceDefinitionId := "sif-ati-fraud" // string | The unique ID for an anti-fraud service definition.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AntiFraudServiceDefinitionsApi.GetAntiFraudServiceDefinition(context.Background(), antiFraudServiceDefinitionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntiFraudServiceDefinitionsApi.GetAntiFraudServiceDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAntiFraudServiceDefinition`: AntiFraudServiceDefinition
    fmt.Fprintf(os.Stdout, "Response from `AntiFraudServiceDefinitionsApi.GetAntiFraudServiceDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**antiFraudServiceDefinitionId** | **string** | The unique ID for an anti-fraud service definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAntiFraudServiceDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AntiFraudServiceDefinition**](AntiFraudServiceDefinition.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

