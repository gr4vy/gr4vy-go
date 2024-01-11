# \AccountUpdaterApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NewAccountUpdaterJob**](AccountUpdaterApi.md#NewAccountUpdaterJob) | **Post** /account-updater/jobs | Create Account Updater job



## NewAccountUpdaterJob

> AccountUpdaterJob NewAccountUpdaterJob(ctx).AccountUpdaterJobCreate(accountUpdaterJobCreate).Execute()

Create Account Updater job



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
    accountUpdaterJobCreate := *openapiclient.NewAccountUpdaterJobCreate([]string{"PaymentMethodIds_example"}) // AccountUpdaterJobCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountUpdaterApi.NewAccountUpdaterJob(context.Background()).AccountUpdaterJobCreate(accountUpdaterJobCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountUpdaterApi.NewAccountUpdaterJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewAccountUpdaterJob`: AccountUpdaterJob
    fmt.Fprintf(os.Stdout, "Response from `AccountUpdaterApi.NewAccountUpdaterJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewAccountUpdaterJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountUpdaterJobCreate** | [**AccountUpdaterJobCreate**](AccountUpdaterJobCreate.md) |  | 

### Return type

[**AccountUpdaterJob**](AccountUpdaterJob.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

