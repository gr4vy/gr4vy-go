# \TransactionsApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizeNewTransaction**](TransactionsApi.md#AuthorizeNewTransaction) | **Post** /transactions | New transaction
[**AuthorizeTransaction**](TransactionsApi.md#AuthorizeTransaction) | **Post** /transactions/{transaction_id}/authorize | Authorize approved transaction
[**CaptureTransaction**](TransactionsApi.md#CaptureTransaction) | **Post** /transactions/{transaction_id}/capture | Capture transaction
[**GetTransaction**](TransactionsApi.md#GetTransaction) | **Get** /transactions/{transaction_id} | Get transaction
[**ListTransactions**](TransactionsApi.md#ListTransactions) | **Get** /transactions | List transactions
[**RefundTransaction**](TransactionsApi.md#RefundTransaction) | **Post** /transactions/{transaction_id}/refund | Refund or void transaction



## AuthorizeNewTransaction

> Transaction AuthorizeNewTransaction(ctx).TransactionRequest(transactionRequest).Execute()

New transaction



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
    transactionRequest := *openapiclient.NewTransactionRequest(float32(1299), "USD", *openapiclient.NewTransactionPaymentMethodRequest("card")) // TransactionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.AuthorizeNewTransaction(context.Background()).TransactionRequest(transactionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.AuthorizeNewTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizeNewTransaction`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.AuthorizeNewTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeNewTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionRequest** | [**TransactionRequest**](TransactionRequest.md) |  | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizeTransaction

> Transaction AuthorizeTransaction(ctx, transactionId).Execute()

Authorize approved transaction



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
    transactionId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID for the transaction to get the information for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.AuthorizeTransaction(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.AuthorizeTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizeTransaction`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.AuthorizeTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The ID for the transaction to get the information for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CaptureTransaction

> Transaction CaptureTransaction(ctx, transactionId).TransactionCaptureRequest(transactionCaptureRequest).Execute()

Capture transaction



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
    transactionId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID for the transaction to get the information for.
    transactionCaptureRequest := *openapiclient.NewTransactionCaptureRequest(float32(1299), "USD") // TransactionCaptureRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.CaptureTransaction(context.Background(), transactionId).TransactionCaptureRequest(transactionCaptureRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.CaptureTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CaptureTransaction`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.CaptureTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The ID for the transaction to get the information for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCaptureTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionCaptureRequest** | [**TransactionCaptureRequest**](TransactionCaptureRequest.md) |  | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransaction

> Transaction GetTransaction(ctx, transactionId).Execute()

Get transaction



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
    transactionId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID for the transaction to get the information for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.GetTransaction(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransaction`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The ID for the transaction to get the information for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactions

> Transactions ListTransactions(ctx).Search(search).TransactionStatus(transactionStatus).BeforeCreatedAt(beforeCreatedAt).AfterCreatedAt(afterCreatedAt).BeforeUpdatedAt(beforeUpdatedAt).AfterUpdatedAt(afterUpdatedAt).Limit(limit).Cursor(cursor).Execute()

List transactions



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
    search := "be828248-56de-481e-a580-44b6e1d4df81" // string | Filters the transactions to only the items for which the `id` or `external_identifier` matches this value. This field allows for a partial match, matching any transaction for which either of the fields partially or completely matches. (optional)
    transactionStatus := "captured" // string | Filters the results to only the transactions for which the `status` matches this value. (optional)
    beforeCreatedAt := "2012-12-12T10:53:43+00:00" // string | Filters the results to only transactions created before this ISO date-time string. (optional)
    afterCreatedAt := "2012-12-12T10:53:43+00:00" // string | Filters the results to only transactions created after this ISO date-time string. (optional)
    beforeUpdatedAt := "2012-12-12T10:53:43+00:00" // string | Filters the results to only transactions last updated before this ISO date-time string. (optional)
    afterUpdatedAt := "2012-12-12T10:53:43+00:00" // string | Filters the results to only transactions last updated after this ISO date-time string. (optional)
    limit := int32(1) // int32 | Defines the maximum number of items to return for this request. (optional) (default to 20)
    cursor := "ZXhhbXBsZTE" // string | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the `next_cursor` field. Similarly the `previous_cursor` can be used to reverse backwards in the list. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.ListTransactions(context.Background()).Search(search).TransactionStatus(transactionStatus).BeforeCreatedAt(beforeCreatedAt).AfterCreatedAt(afterCreatedAt).BeforeUpdatedAt(beforeUpdatedAt).AfterUpdatedAt(afterUpdatedAt).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.ListTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactions`: Transactions
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.ListTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Filters the transactions to only the items for which the &#x60;id&#x60; or &#x60;external_identifier&#x60; matches this value. This field allows for a partial match, matching any transaction for which either of the fields partially or completely matches. | 
 **transactionStatus** | **string** | Filters the results to only the transactions for which the &#x60;status&#x60; matches this value. | 
 **beforeCreatedAt** | **string** | Filters the results to only transactions created before this ISO date-time string. | 
 **afterCreatedAt** | **string** | Filters the results to only transactions created after this ISO date-time string. | 
 **beforeUpdatedAt** | **string** | Filters the results to only transactions last updated before this ISO date-time string. | 
 **afterUpdatedAt** | **string** | Filters the results to only transactions last updated after this ISO date-time string. | 
 **limit** | **int32** | Defines the maximum number of items to return for this request. | [default to 20]
 **cursor** | **string** | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the &#x60;next_cursor&#x60; field. Similarly the &#x60;previous_cursor&#x60; can be used to reverse backwards in the list. | 

### Return type

[**Transactions**](Transactions.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundTransaction

> Transaction RefundTransaction(ctx, transactionId).Execute()

Refund or void transaction



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
    transactionId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID for the transaction to get the information for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.RefundTransaction(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.RefundTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefundTransaction`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.RefundTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The ID for the transaction to get the information for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefundTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
