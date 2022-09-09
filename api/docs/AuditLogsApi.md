# \AuditLogsApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAuditLogs**](AuditLogsApi.md#ListAuditLogs) | **Get** /audit-logs | List Audit Logs



## ListAuditLogs

> AuditLogs ListAuditLogs(ctx).Limit(limit).Cursor(cursor).Execute()

List Audit Logs



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogsApi.ListAuditLogs(context.Background()).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsApi.ListAuditLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuditLogs`: AuditLogs
    fmt.Fprintf(os.Stdout, "Response from `AuditLogsApi.ListAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Defines the maximum number of items to return for this request. | [default to 20]
 **cursor** | **string** | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the &#x60;next_cursor&#x60; field. Similarly the &#x60;previous_cursor&#x60; can be used to reverse backwards in the list. | 

### Return type

[**AuditLogs**](AuditLogs.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

