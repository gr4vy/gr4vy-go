# \CheckoutSessionsApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCheckoutSession**](CheckoutSessionsApi.md#DeleteCheckoutSession) | **Delete** /checkout/sessions/{checkout_session_id} | Delete checkout session
[**GetCheckoutSession**](CheckoutSessionsApi.md#GetCheckoutSession) | **Get** /checkout/sessions/{checkout_session_id} | Get checkout session
[**NewCheckoutSession**](CheckoutSessionsApi.md#NewCheckoutSession) | **Post** /checkout/sessions | New checkout session
[**UpdateCheckoutSession**](CheckoutSessionsApi.md#UpdateCheckoutSession) | **Put** /checkout/sessions/{checkout_session_id} | Update checkout session
[**UpdateCheckoutSessionFields**](CheckoutSessionsApi.md#UpdateCheckoutSessionFields) | **Put** /checkout/sessions/{checkout_session_id}/fields | Update fields for checkout session



## DeleteCheckoutSession

> DeleteCheckoutSession(ctx, checkoutSessionId).Execute()

Delete checkout session



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
    checkoutSessionId := TODO // string | The unique ID for a Checkout Session.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CheckoutSessionsApi.DeleteCheckoutSession(context.Background(), checkoutSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSessionsApi.DeleteCheckoutSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutSessionId** | [**string**](.md) | The unique ID for a Checkout Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCheckoutSessionRequest struct via the builder pattern


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


## GetCheckoutSession

> CheckoutSession GetCheckoutSession(ctx, checkoutSessionId).Execute()

Get checkout session



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
    checkoutSessionId := TODO // string | The unique ID for a Checkout Session.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CheckoutSessionsApi.GetCheckoutSession(context.Background(), checkoutSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSessionsApi.GetCheckoutSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCheckoutSession`: CheckoutSession
    fmt.Fprintf(os.Stdout, "Response from `CheckoutSessionsApi.GetCheckoutSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutSessionId** | [**string**](.md) | The unique ID for a Checkout Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckoutSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CheckoutSession**](CheckoutSession.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewCheckoutSession

> CheckoutSession NewCheckoutSession(ctx).CheckoutSessionCreateRequest(checkoutSessionCreateRequest).Execute()

New checkout session



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
    checkoutSessionCreateRequest := *openapiclient.NewCheckoutSessionCreateRequest() // CheckoutSessionCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CheckoutSessionsApi.NewCheckoutSession(context.Background()).CheckoutSessionCreateRequest(checkoutSessionCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSessionsApi.NewCheckoutSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewCheckoutSession`: CheckoutSession
    fmt.Fprintf(os.Stdout, "Response from `CheckoutSessionsApi.NewCheckoutSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewCheckoutSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkoutSessionCreateRequest** | [**CheckoutSessionCreateRequest**](CheckoutSessionCreateRequest.md) |  | 

### Return type

[**CheckoutSession**](CheckoutSession.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCheckoutSession

> CheckoutSession UpdateCheckoutSession(ctx, checkoutSessionId).CheckoutSessionUpdateRequest(checkoutSessionUpdateRequest).Execute()

Update checkout session



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
    checkoutSessionId := TODO // string | The unique ID for a Checkout Session.
    checkoutSessionUpdateRequest := *openapiclient.NewCheckoutSessionUpdateRequest() // CheckoutSessionUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CheckoutSessionsApi.UpdateCheckoutSession(context.Background(), checkoutSessionId).CheckoutSessionUpdateRequest(checkoutSessionUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSessionsApi.UpdateCheckoutSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCheckoutSession`: CheckoutSession
    fmt.Fprintf(os.Stdout, "Response from `CheckoutSessionsApi.UpdateCheckoutSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutSessionId** | [**string**](.md) | The unique ID for a Checkout Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCheckoutSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkoutSessionUpdateRequest** | [**CheckoutSessionUpdateRequest**](CheckoutSessionUpdateRequest.md) |  | 

### Return type

[**CheckoutSession**](CheckoutSession.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCheckoutSessionFields

> UpdateCheckoutSessionFields(ctx, checkoutSessionId).CheckoutSessionSecureFieldsUpdate(checkoutSessionSecureFieldsUpdate).Execute()

Update fields for checkout session



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
    checkoutSessionId := TODO // string | The unique ID for a Checkout Session.
    checkoutSessionSecureFieldsUpdate := *openapiclient.NewCheckoutSessionSecureFieldsUpdate() // CheckoutSessionSecureFieldsUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CheckoutSessionsApi.UpdateCheckoutSessionFields(context.Background(), checkoutSessionId).CheckoutSessionSecureFieldsUpdate(checkoutSessionSecureFieldsUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSessionsApi.UpdateCheckoutSessionFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutSessionId** | [**string**](.md) | The unique ID for a Checkout Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCheckoutSessionFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkoutSessionSecureFieldsUpdate** | [**CheckoutSessionSecureFieldsUpdate**](CheckoutSessionSecureFieldsUpdate.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

