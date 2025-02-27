# \WebhookSubscriptionsApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebhookSubscription**](WebhookSubscriptionsApi.md#DeleteWebhookSubscription) | **Delete** /webhook-subscriptions/{webhook_subscription_id} | Delete a webhook subscription
[**GetWebhookSubscription**](WebhookSubscriptionsApi.md#GetWebhookSubscription) | **Get** /webhook-subscriptions/{webhook_subscription_id} | Get webhook subscription
[**ListWebhookSubscriptions**](WebhookSubscriptionsApi.md#ListWebhookSubscriptions) | **Get** /webhook-subscriptions | List webhook subscriptions
[**NewWebhookSubscription**](WebhookSubscriptionsApi.md#NewWebhookSubscription) | **Post** /webhook-subscriptions | New webhook subscription
[**UpdateWebhookSubscription**](WebhookSubscriptionsApi.md#UpdateWebhookSubscription) | **Put** /webhook-subscriptions/{webhook_subscription_id} | Update webhook subscription



## DeleteWebhookSubscription

> DeleteWebhookSubscription(ctx, webhookSubscriptionId).Execute()

Delete a webhook subscription



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
    webhookSubscriptionId := "8fd77b13-a5e3-43de-bd54-26a8a714ac18" // string | The ID for the webhook subscription.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhookSubscriptionsApi.DeleteWebhookSubscription(context.Background(), webhookSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionsApi.DeleteWebhookSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookSubscriptionId** | **string** | The ID for the webhook subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookSubscriptionRequest struct via the builder pattern


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


## GetWebhookSubscription

> WebhookSubscription GetWebhookSubscription(ctx, webhookSubscriptionId).Execute()

Get webhook subscription



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
    webhookSubscriptionId := "8fd77b13-a5e3-43de-bd54-26a8a714ac18" // string | The ID for the webhook subscription.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhookSubscriptionsApi.GetWebhookSubscription(context.Background(), webhookSubscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionsApi.GetWebhookSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookSubscription`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `WebhookSubscriptionsApi.GetWebhookSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookSubscriptionId** | **string** | The ID for the webhook subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookSubscription**](WebhookSubscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhookSubscriptions

> WebhookSubscriptions ListWebhookSubscriptions(ctx).Execute()

List webhook subscriptions



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
    resp, r, err := api_client.WebhookSubscriptionsApi.ListWebhookSubscriptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionsApi.ListWebhookSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebhookSubscriptions`: WebhookSubscriptions
    fmt.Fprintf(os.Stdout, "Response from `WebhookSubscriptionsApi.ListWebhookSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWebhookSubscriptionsRequest struct via the builder pattern


### Return type

[**WebhookSubscriptions**](WebhookSubscriptions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewWebhookSubscription

> WebhookSubscription NewWebhookSubscription(ctx).WebhookSubscriptionRequest(webhookSubscriptionRequest).Execute()

New webhook subscription



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
    webhookSubscriptionRequest := *openapiclient.NewWebhookSubscriptionRequest() // WebhookSubscriptionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhookSubscriptionsApi.NewWebhookSubscription(context.Background()).WebhookSubscriptionRequest(webhookSubscriptionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionsApi.NewWebhookSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewWebhookSubscription`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `WebhookSubscriptionsApi.NewWebhookSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewWebhookSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookSubscriptionRequest** | [**WebhookSubscriptionRequest**](WebhookSubscriptionRequest.md) |  | 

### Return type

[**WebhookSubscription**](WebhookSubscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhookSubscription

> WebhookSubscription UpdateWebhookSubscription(ctx, webhookSubscriptionId).WebhookSubscriptionUpdateRequest(webhookSubscriptionUpdateRequest).Execute()

Update webhook subscription



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
    webhookSubscriptionId := "8fd77b13-a5e3-43de-bd54-26a8a714ac18" // string | The ID for the webhook subscription.
    webhookSubscriptionUpdateRequest := *openapiclient.NewWebhookSubscriptionUpdateRequest() // WebhookSubscriptionUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhookSubscriptionsApi.UpdateWebhookSubscription(context.Background(), webhookSubscriptionId).WebhookSubscriptionUpdateRequest(webhookSubscriptionUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookSubscriptionsApi.UpdateWebhookSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebhookSubscription`: WebhookSubscription
    fmt.Fprintf(os.Stdout, "Response from `WebhookSubscriptionsApi.UpdateWebhookSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookSubscriptionId** | **string** | The ID for the webhook subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhookSubscriptionUpdateRequest** | [**WebhookSubscriptionUpdateRequest**](WebhookSubscriptionUpdateRequest.md) |  | 

### Return type

[**WebhookSubscription**](WebhookSubscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

