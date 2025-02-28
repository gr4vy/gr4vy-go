# \TokensApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNetworkToken**](TokensApi.md#DeleteNetworkToken) | **Delete** /payment-methods/{payment_method_id}/network-tokens/{network_token_id} | Delete network token
[**DeletePaymentServiceToken**](TokensApi.md#DeletePaymentServiceToken) | **Delete** /payment-methods/{payment_method_id}/payment-service-tokens/{payment_service_token_id} | Delete payment service token
[**GetNetworkTokens**](TokensApi.md#GetNetworkTokens) | **Get** /payment-methods/{payment_method_id}/network-tokens | Get network tokens
[**GetPaymentServiceTokens**](TokensApi.md#GetPaymentServiceTokens) | **Get** /payment-methods/{payment_method_id}/payment-service-tokens | Get payment service tokens
[**IssueCryptogram**](TokensApi.md#IssueCryptogram) | **Post** /payment-methods/{payment_method_id}/network-tokens/{network_token_id}/cryptogram | Issue cryptogram
[**ProvisionNetworkToken**](TokensApi.md#ProvisionNetworkToken) | **Post** /payment-methods/{payment_method_id}/network-tokens | Provision network token
[**ProvisionPaymentServiceToken**](TokensApi.md#ProvisionPaymentServiceToken) | **Post** /payment-methods/{payment_method_id}/payment-service-tokens | Provision payment service token
[**ResumeNetworkToken**](TokensApi.md#ResumeNetworkToken) | **Post** /payment-methods/{payment_method_id}/network-tokens/{network_token_id}/resume | Resume network token
[**SuspendNetworkToken**](TokensApi.md#SuspendNetworkToken) | **Post** /payment-methods/{payment_method_id}/network-tokens/{network_token_id}/suspend | Suspend network token



## DeleteNetworkToken

> DeleteNetworkToken(ctx, paymentMethodId, networkTokenId).Execute()

Delete network token



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
    paymentMethodId := TODO // string | The ID of the payment method.
    networkTokenId := TODO // string | The ID of the network token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensApi.DeleteNetworkToken(context.Background(), paymentMethodId, networkTokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.DeleteNetworkToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | [**string**](.md) | The ID of the payment method. | 
**networkTokenId** | [**string**](.md) | The ID of the network token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkTokenRequest struct via the builder pattern


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


## DeletePaymentServiceToken

> DeletePaymentServiceToken(ctx, paymentMethodId, paymentServiceTokenId).Execute()

Delete payment service token



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
    paymentMethodId := TODO // string | The ID of the payment method.
    paymentServiceTokenId := TODO // string | The ID of the payment service token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensApi.DeletePaymentServiceToken(context.Background(), paymentMethodId, paymentServiceTokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.DeletePaymentServiceToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | [**string**](.md) | The ID of the payment method. | 
**paymentServiceTokenId** | [**string**](.md) | The ID of the payment service token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePaymentServiceTokenRequest struct via the builder pattern


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


## GetNetworkTokens

> NetworkTokens GetNetworkTokens(ctx, paymentMethodId).PaymentMethodId2(paymentMethodId2).Execute()

Get network tokens



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
    paymentMethodId := TODO // string | The ID of the payment method.
    paymentMethodId2 := TODO // string | Filters for transactions that have a payment method with an ID that matches exactly with the provided value. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensApi.GetNetworkTokens(context.Background(), paymentMethodId).PaymentMethodId2(paymentMethodId2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.GetNetworkTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkTokens`: NetworkTokens
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.GetNetworkTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | [**string**](.md) | The ID of the payment method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentMethodId2** | [**string**](string.md) | Filters for transactions that have a payment method with an ID that matches exactly with the provided value. | 

### Return type

[**NetworkTokens**](NetworkTokens.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentServiceTokens

> PaymentServiceTokens GetPaymentServiceTokens(ctx, paymentMethodId).PaymentMethodId2(paymentMethodId2).Execute()

Get payment service tokens



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
    paymentMethodId := TODO // string | The ID of the payment method.
    paymentMethodId2 := TODO // string | Filters for transactions that have a payment method with an ID that matches exactly with the provided value. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensApi.GetPaymentServiceTokens(context.Background(), paymentMethodId).PaymentMethodId2(paymentMethodId2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.GetPaymentServiceTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentServiceTokens`: PaymentServiceTokens
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.GetPaymentServiceTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | [**string**](.md) | The ID of the payment method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentServiceTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentMethodId2** | [**string**](string.md) | Filters for transactions that have a payment method with an ID that matches exactly with the provided value. | 

### Return type

[**PaymentServiceTokens**](PaymentServiceTokens.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueCryptogram

> Cryptogram IssueCryptogram(ctx, paymentMethodId, networkTokenId).IssueCryptogramRequest(issueCryptogramRequest).Execute()

Issue cryptogram



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
    paymentMethodId := TODO // string | The ID of the payment method.
    networkTokenId := TODO // string | The ID of the network token.
    issueCryptogramRequest := *openapiclient.NewIssueCryptogramRequest(false) // IssueCryptogramRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensApi.IssueCryptogram(context.Background(), paymentMethodId, networkTokenId).IssueCryptogramRequest(issueCryptogramRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.IssueCryptogram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssueCryptogram`: Cryptogram
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.IssueCryptogram`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | [**string**](.md) | The ID of the payment method. | 
**networkTokenId** | [**string**](.md) | The ID of the network token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueCryptogramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **issueCryptogramRequest** | [**IssueCryptogramRequest**](IssueCryptogramRequest.md) |  | 

### Return type

[**Cryptogram**](Cryptogram.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisionNetworkToken

> NetworkToken ProvisionNetworkToken(ctx, paymentMethodId).NetworkTokenRequest(networkTokenRequest).Execute()

Provision network token



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
    paymentMethodId := TODO // string | The ID of the payment method.
    networkTokenRequest := *openapiclient.NewNetworkTokenRequest(false, true) // NetworkTokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensApi.ProvisionNetworkToken(context.Background(), paymentMethodId).NetworkTokenRequest(networkTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.ProvisionNetworkToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisionNetworkToken`: NetworkToken
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.ProvisionNetworkToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | [**string**](.md) | The ID of the payment method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisionNetworkTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkTokenRequest** | [**NetworkTokenRequest**](NetworkTokenRequest.md) |  | 

### Return type

[**NetworkToken**](NetworkToken.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisionPaymentServiceToken

> PaymentServiceToken ProvisionPaymentServiceToken(ctx, paymentMethodId).PaymentServiceTokenRequest(paymentServiceTokenRequest).Execute()

Provision payment service token



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
    paymentMethodId := TODO // string | The ID of the payment method.
    paymentServiceTokenRequest := *openapiclient.NewPaymentServiceTokenRequest("a7d6b829-aea5-407d-ab7f-138784b5ad2c", "https://example.com/callback") // PaymentServiceTokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensApi.ProvisionPaymentServiceToken(context.Background(), paymentMethodId).PaymentServiceTokenRequest(paymentServiceTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.ProvisionPaymentServiceToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisionPaymentServiceToken`: PaymentServiceToken
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.ProvisionPaymentServiceToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | [**string**](.md) | The ID of the payment method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisionPaymentServiceTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentServiceTokenRequest** | [**PaymentServiceTokenRequest**](PaymentServiceTokenRequest.md) |  | 

### Return type

[**PaymentServiceToken**](PaymentServiceToken.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeNetworkToken

> NetworkToken ResumeNetworkToken(ctx, paymentMethodId, networkTokenId).Execute()

Resume network token



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
    paymentMethodId := TODO // string | The ID of the payment method.
    networkTokenId := TODO // string | The ID of the network token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensApi.ResumeNetworkToken(context.Background(), paymentMethodId, networkTokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.ResumeNetworkToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeNetworkToken`: NetworkToken
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.ResumeNetworkToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | [**string**](.md) | The ID of the payment method. | 
**networkTokenId** | [**string**](.md) | The ID of the network token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeNetworkTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkToken**](NetworkToken.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendNetworkToken

> NetworkToken SuspendNetworkToken(ctx, paymentMethodId, networkTokenId).Execute()

Suspend network token



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
    paymentMethodId := TODO // string | The ID of the payment method.
    networkTokenId := TODO // string | The ID of the network token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensApi.SuspendNetworkToken(context.Background(), paymentMethodId, networkTokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.SuspendNetworkToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuspendNetworkToken`: NetworkToken
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.SuspendNetworkToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | [**string**](.md) | The ID of the payment method. | 
**networkTokenId** | [**string**](.md) | The ID of the network token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendNetworkTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkToken**](NetworkToken.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

