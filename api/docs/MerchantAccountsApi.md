# \MerchantAccountsApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMerchantAccuont**](MerchantAccountsApi.md#DeleteMerchantAccuont) | **Delete** /merchant-accounts/{merchant_account_id} | Delete merchant account
[**GetMerchantAccount**](MerchantAccountsApi.md#GetMerchantAccount) | **Get** /merchant-accounts/{merchant_account_id} | Get merchant account
[**ListMerchantAccounts**](MerchantAccountsApi.md#ListMerchantAccounts) | **Get** /merchant-accounts | List merchant accounts
[**NewMerchantAccount**](MerchantAccountsApi.md#NewMerchantAccount) | **Post** /merchant-accounts | New merchant account
[**UpdateMerchantAccount**](MerchantAccountsApi.md#UpdateMerchantAccount) | **Put** /merchant-accounts/{merchant_account_id} | Update merchant account



## DeleteMerchantAccuont

> DeleteMerchantAccuont(ctx, merchantAccountId).Execute()

Delete merchant account



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
    merchantAccountId := "plantly-uk" // string | The unique ID for a merchant account.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MerchantAccountsApi.DeleteMerchantAccuont(context.Background(), merchantAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantAccountsApi.DeleteMerchantAccuont``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantAccountId** | **string** | The unique ID for a merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMerchantAccuontRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMerchantAccount

> MerchantAccount GetMerchantAccount(ctx, merchantAccountId).Execute()

Get merchant account



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
    merchantAccountId := "plantly-uk" // string | The unique ID for a merchant account.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MerchantAccountsApi.GetMerchantAccount(context.Background(), merchantAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantAccountsApi.GetMerchantAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMerchantAccount`: MerchantAccount
    fmt.Fprintf(os.Stdout, "Response from `MerchantAccountsApi.GetMerchantAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantAccountId** | **string** | The unique ID for a merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMerchantAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MerchantAccount**](MerchantAccount.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMerchantAccounts

> MerchantAccounts ListMerchantAccounts(ctx).Execute()

List merchant accounts



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
    resp, r, err := api_client.MerchantAccountsApi.ListMerchantAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantAccountsApi.ListMerchantAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMerchantAccounts`: MerchantAccounts
    fmt.Fprintf(os.Stdout, "Response from `MerchantAccountsApi.ListMerchantAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMerchantAccountsRequest struct via the builder pattern


### Return type

[**MerchantAccounts**](MerchantAccounts.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewMerchantAccount

> MerchantAccount NewMerchantAccount(ctx).MerchantAccountCreate(merchantAccountCreate).Execute()

New merchant account



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
    merchantAccountCreate := *openapiclient.NewMerchantAccountCreate() // MerchantAccountCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MerchantAccountsApi.NewMerchantAccount(context.Background()).MerchantAccountCreate(merchantAccountCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantAccountsApi.NewMerchantAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewMerchantAccount`: MerchantAccount
    fmt.Fprintf(os.Stdout, "Response from `MerchantAccountsApi.NewMerchantAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewMerchantAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantAccountCreate** | [**MerchantAccountCreate**](MerchantAccountCreate.md) |  | 

### Return type

[**MerchantAccount**](MerchantAccount.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMerchantAccount

> MerchantAccount UpdateMerchantAccount(ctx, merchantAccountId).MerchantAccountUpdate(merchantAccountUpdate).Execute()

Update merchant account



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
    merchantAccountId := "plantly-uk" // string | The unique ID for a merchant account.
    merchantAccountUpdate := *openapiclient.NewMerchantAccountUpdate() // MerchantAccountUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MerchantAccountsApi.UpdateMerchantAccount(context.Background(), merchantAccountId).MerchantAccountUpdate(merchantAccountUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantAccountsApi.UpdateMerchantAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMerchantAccount`: MerchantAccount
    fmt.Fprintf(os.Stdout, "Response from `MerchantAccountsApi.UpdateMerchantAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantAccountId** | **string** | The unique ID for a merchant account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMerchantAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **merchantAccountUpdate** | [**MerchantAccountUpdate**](MerchantAccountUpdate.md) |  | 

### Return type

[**MerchantAccount**](MerchantAccount.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

