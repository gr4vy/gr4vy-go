# \DigitalWalletsApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDigitalWallet**](DigitalWalletsApi.md#DeleteDigitalWallet) | **Delete** /digital-wallets/{digital_wallet_id} | De-register digital wallet
[**GetDigitalWallet**](DigitalWalletsApi.md#GetDigitalWallet) | **Get** /digital-wallets/{digital_wallet_id} | Get digital wallet
[**ListDigitalWallets**](DigitalWalletsApi.md#ListDigitalWallets) | **Get** /digital-wallets | List digital wallets
[**NewDigitalWallet**](DigitalWalletsApi.md#NewDigitalWallet) | **Post** /digital-wallets | Register digital wallet
[**UpdateDigitalWallet**](DigitalWalletsApi.md#UpdateDigitalWallet) | **Put** /digital-wallets/{digital_wallet_id} | Update digital wallet



## DeleteDigitalWallet

> DeleteDigitalWallet(ctx, digitalWalletId).Execute()

De-register digital wallet



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
    digitalWalletId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID of the registered digital wallet.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DigitalWalletsApi.DeleteDigitalWallet(context.Background(), digitalWalletId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsApi.DeleteDigitalWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**digitalWalletId** | **string** | The ID of the registered digital wallet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDigitalWalletRequest struct via the builder pattern


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


## GetDigitalWallet

> DigitalWallet GetDigitalWallet(ctx, digitalWalletId).Execute()

Get digital wallet



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
    digitalWalletId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID of the registered digital wallet.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DigitalWalletsApi.GetDigitalWallet(context.Background(), digitalWalletId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsApi.GetDigitalWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDigitalWallet`: DigitalWallet
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsApi.GetDigitalWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**digitalWalletId** | **string** | The ID of the registered digital wallet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDigitalWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DigitalWallet**](DigitalWallet.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDigitalWallets

> DigitalWallets ListDigitalWallets(ctx).Execute()

List digital wallets



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
    resp, r, err := api_client.DigitalWalletsApi.ListDigitalWallets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsApi.ListDigitalWallets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDigitalWallets`: DigitalWallets
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsApi.ListDigitalWallets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDigitalWalletsRequest struct via the builder pattern


### Return type

[**DigitalWallets**](DigitalWallets.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewDigitalWallet

> DigitalWallet NewDigitalWallet(ctx).DigitalWalletRequest(digitalWalletRequest).Execute()

Register digital wallet



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
    digitalWalletRequest := *openapiclient.NewDigitalWalletRequest("apple", "Gr4vy", []string{"DomainNames_example"}, true) // DigitalWalletRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DigitalWalletsApi.NewDigitalWallet(context.Background()).DigitalWalletRequest(digitalWalletRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsApi.NewDigitalWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewDigitalWallet`: DigitalWallet
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsApi.NewDigitalWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewDigitalWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **digitalWalletRequest** | [**DigitalWalletRequest**](DigitalWalletRequest.md) |  | 

### Return type

[**DigitalWallet**](DigitalWallet.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDigitalWallet

> DigitalWallet UpdateDigitalWallet(ctx, digitalWalletId).DigitalWalletUpdate(digitalWalletUpdate).Execute()

Update digital wallet



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
    digitalWalletId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID of the registered digital wallet.
    digitalWalletUpdate := *openapiclient.NewDigitalWalletUpdate() // DigitalWalletUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DigitalWalletsApi.UpdateDigitalWallet(context.Background(), digitalWalletId).DigitalWalletUpdate(digitalWalletUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsApi.UpdateDigitalWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDigitalWallet`: DigitalWallet
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsApi.UpdateDigitalWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**digitalWalletId** | **string** | The ID of the registered digital wallet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDigitalWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **digitalWalletUpdate** | [**DigitalWalletUpdate**](DigitalWalletUpdate.md) |  | 

### Return type

[**DigitalWallet**](DigitalWallet.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

