# \TransactionsApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizeNewTransaction**](TransactionsApi.md#AuthorizeNewTransaction) | **Post** /transactions | New transaction
[**CaptureTransaction**](TransactionsApi.md#CaptureTransaction) | **Post** /transactions/{transaction_id}/capture | Capture transaction
[**GetTransaction**](TransactionsApi.md#GetTransaction) | **Get** /transactions/{transaction_id} | Get transaction
[**GetTransactionRefund**](TransactionsApi.md#GetTransactionRefund) | **Get** /transactions/{transaction_id}/refunds/{refund_id} | Get transaction refund
[**ListTransactionRefunds**](TransactionsApi.md#ListTransactionRefunds) | **Get** /transactions/{transaction_id}/refunds | List transaction refunds
[**ListTransactions**](TransactionsApi.md#ListTransactions) | **Get** /transactions | List transactions
[**RefundTransaction**](TransactionsApi.md#RefundTransaction) | **Post** /transactions/{transaction_id}/refunds | Refund transaction
[**VoidTransaction**](TransactionsApi.md#VoidTransaction) | **Post** /transactions/{transaction_id}/void | Void transaction



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
    transactionRequest := *openapiclient.NewTransactionRequest(int32(1299), "USD", *openapiclient.NewTransactionPaymentMethodRequest(string(123))) // TransactionRequest |  (optional)

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
    transactionCaptureRequest := *openapiclient.NewTransactionCaptureRequest() // TransactionCaptureRequest |  (optional)

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


## GetTransactionRefund

> Refund GetTransactionRefund(ctx, transactionId, refundId).Execute()

Get transaction refund



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
    refundId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID of the refund.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.GetTransactionRefund(context.Background(), transactionId, refundId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetTransactionRefund``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionRefund`: Refund
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetTransactionRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The ID for the transaction to get the information for. | 
**refundId** | **string** | The unique ID of the refund. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Refund**](Refund.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionRefunds

> Refunds ListTransactionRefunds(ctx, transactionId).Limit(limit).Cursor(cursor).Execute()

List transaction refunds



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
    limit := int32(1) // int32 | Defines the maximum number of items to return for this request. (optional) (default to 20)
    cursor := "ZXhhbXBsZTE" // string | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the `next_cursor` field. Similarly the `previous_cursor` can be used to reverse backwards in the list. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.ListTransactionRefunds(context.Background(), transactionId).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.ListTransactionRefunds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionRefunds`: Refunds
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.ListTransactionRefunds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The ID for the transaction to get the information for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionRefundsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Defines the maximum number of items to return for this request. | [default to 20]
 **cursor** | **string** | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the &#x60;next_cursor&#x60; field. Similarly the &#x60;previous_cursor&#x60; can be used to reverse backwards in the list. | 

### Return type

[**Refunds**](Refunds.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactions

> Transactions ListTransactions(ctx).BuyerExternalIdentifier(buyerExternalIdentifier).BuyerId(buyerId).Cursor(cursor).Limit(limit).AmountEq(amountEq).AmountGte(amountGte).AmountLte(amountLte).CreatedAtGte(createdAtGte).CreatedAtLte(createdAtLte).Currency(currency).ExternalIdentifier(externalIdentifier).HasRefunds(hasRefunds).Id(id).Metadata(metadata).Method(method).PaymentMethodId(paymentMethodId).PaymentServiceId(paymentServiceId).PaymentServiceTransactionId(paymentServiceTransactionId).Search(search).Status(status).UpdatedAtGte(updatedAtGte).UpdatedAtLte(updatedAtLte).BeforeCreatedAt(beforeCreatedAt).AfterCreatedAt(afterCreatedAt).BeforeUpdatedAt(beforeUpdatedAt).AfterUpdatedAt(afterUpdatedAt).TransactionStatus(transactionStatus).Execute()

List transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    buyerExternalIdentifier := "user-12345" // string | Filters the results to only the items for which the `buyer` has an `external_identifier` that matches this value. (optional)
    buyerId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | Filters the results to only the items for which the `buyer` has an `id` that matches this value. (optional)
    cursor := "ZXhhbXBsZTE" // string | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the `next_cursor` field. Similarly the `previous_cursor` can be used to reverse backwards in the list. (optional)
    limit := int32(1) // int32 | Defines the maximum number of items to return for this request. (optional) (default to 20)
    amountEq := int32(500) // int32 | Filters for transactions that have an `amount` that is equal to the provided `amount_eq` value. (optional)
    amountGte := int32(500) // int32 | Filters for transactions that have an `amount` that is greater than or equal to the `amount_gte` value. (optional)
    amountLte := int32(500) // int32 | Filters for transactions that have an `amount` that is less than or equal to the `amount_lte` value. (optional)
    createdAtGte := time.Now() // time.Time | Filters the results to only transactions created after this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. `2022-01-01T12:00:00+08:00` must be encoded as `2022-01-01T12%3A00%3A00%2B08%3A00`. (optional)
    createdAtLte := time.Now() // time.Time | Filters the results to only transactions created before this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. `2022-01-01T12:00:00+08:00` must be encoded as `2022-01-01T12%3A00%3A00%2B08%3A00`. (optional)
    currency := []string{"Inner_example"} // []string | Filters for transactions that have matching `currency` values. The `currency` values provided must be formatted as 3-letter ISO currency code. (optional)
    externalIdentifier := "user-12345" // string | Filters the results to only the items for which the `external_identifier` matches this value. (optional)
    hasRefunds := true // bool | When set to `true`, filter for transactions that have at least one completed refund associated with it. When set to `false`, filter for transactions that have no completed refunds. (optional)
    id := TODO // string | Filters for the transaction that has a matching `id` value. (optional)
    metadata := []string{"Inner_example"} // []string | Filters for transactions where their `metadata` values contain all of the provided `metadata` keys. The value sent for `metadata` must be formatted as a JSON string, and all keys and values must be strings. This value should also be URL encoded.  Duplicate keys are not supported. If a key is duplicated, only the last appearing value will be used. (optional)
    method := []string{"card"} // []string | Filters the results to only the items for which the `method` has been set to this value. (optional)
    paymentMethodId := TODO // string | The ID of the payment method. (optional)
    paymentServiceId := []string{"Inner_example"} // []string | Filters for transactions that were processed by the provided `payment_service_id` values. (optional)
    paymentServiceTransactionId := "transaction_123" // string | Filters for transactions that have a matching `payment_service_transaction_id` value. The `payment_service_transaction_id` is the identifier of the transaction given by the payment service. (optional)
    search := "be828248-56de-481e-a580-44b6e1d4df81" // string | Filters for transactions that have one of the following fields match exactly with the provided `search` value: * `buyer_external_identifier` * `buyer_id` * `external_identifier` * `id` * `payment_service_transaction_id` (optional)
    status := []string{"Status_example"} // []string | Filters the results to only the transactions that have a `status` that matches with any of the provided status values. (optional)
    updatedAtGte := time.Now() // time.Time | Filters the results to only transactions last updated after this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. `2022-01-01T12:00:00+08:00` must be encoded as `2022-01-01T12%3A00%3A00%2B08%3A00`. (optional)
    updatedAtLte := time.Now() // time.Time | Filters the results to only transactions last updated before this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. `2022-01-01T12:00:00+08:00` must be encoded as `2022-01-01T12%3A00%3A00%2B08%3A00`. (optional)
    beforeCreatedAt := time.Now() // time.Time | Filters the results to only transactions created before this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. `2022-01-01T12:00:00+08:00` must be encoded as `2022-01-01T12%3A00%3A00%2B08%3A00`.  **WARNING** This filter is deprecated and may be removed eventually, use `created_at_lte` instead. (optional)
    afterCreatedAt := time.Now() // time.Time | Filters the results to only transactions created after this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. `2022-01-01T12:00:00+08:00` must be encoded as `2022-01-01T12%3A00%3A00%2B08%3A00`.  **WARNING** This filter is deprecated and may be removed eventually, use `created_at_gte` instead. (optional)
    beforeUpdatedAt := time.Now() // time.Time | Filters the results to only transactions last updated before this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. `2022-01-01T12:00:00+08:00` must be encoded as `2022-01-01T12%3A00%3A00%2B08%3A00`.  **WARNING** This filter is deprecated and may be removed eventually, use `updated_at_lte` instead. (optional)
    afterUpdatedAt := time.Now() // time.Time | Filters the results to only transactions last updated after this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. `2022-01-01T12:00:00+08:00` must be encoded as `2022-01-01T12%3A00%3A00%2B08%3A00`.  **WARNING** This filter is deprecated and may be removed eventually, use `updated_at_gte` instead. (optional)
    transactionStatus := "capture_succeeded" // string | Filters the results to only the transactions for which the `status` matches this value.  **WARNING** This filter is deprecated and may be removed eventually, use `status` instead. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.ListTransactions(context.Background()).BuyerExternalIdentifier(buyerExternalIdentifier).BuyerId(buyerId).Cursor(cursor).Limit(limit).AmountEq(amountEq).AmountGte(amountGte).AmountLte(amountLte).CreatedAtGte(createdAtGte).CreatedAtLte(createdAtLte).Currency(currency).ExternalIdentifier(externalIdentifier).HasRefunds(hasRefunds).Id(id).Metadata(metadata).Method(method).PaymentMethodId(paymentMethodId).PaymentServiceId(paymentServiceId).PaymentServiceTransactionId(paymentServiceTransactionId).Search(search).Status(status).UpdatedAtGte(updatedAtGte).UpdatedAtLte(updatedAtLte).BeforeCreatedAt(beforeCreatedAt).AfterCreatedAt(afterCreatedAt).BeforeUpdatedAt(beforeUpdatedAt).AfterUpdatedAt(afterUpdatedAt).TransactionStatus(transactionStatus).Execute()
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
 **buyerExternalIdentifier** | **string** | Filters the results to only the items for which the &#x60;buyer&#x60; has an &#x60;external_identifier&#x60; that matches this value. | 
 **buyerId** | **string** | Filters the results to only the items for which the &#x60;buyer&#x60; has an &#x60;id&#x60; that matches this value. | 
 **cursor** | **string** | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the &#x60;next_cursor&#x60; field. Similarly the &#x60;previous_cursor&#x60; can be used to reverse backwards in the list. | 
 **limit** | **int32** | Defines the maximum number of items to return for this request. | [default to 20]
 **amountEq** | **int32** | Filters for transactions that have an &#x60;amount&#x60; that is equal to the provided &#x60;amount_eq&#x60; value. | 
 **amountGte** | **int32** | Filters for transactions that have an &#x60;amount&#x60; that is greater than or equal to the &#x60;amount_gte&#x60; value. | 
 **amountLte** | **int32** | Filters for transactions that have an &#x60;amount&#x60; that is less than or equal to the &#x60;amount_lte&#x60; value. | 
 **createdAtGte** | **time.Time** | Filters the results to only transactions created after this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. &#x60;2022-01-01T12:00:00+08:00&#x60; must be encoded as &#x60;2022-01-01T12%3A00%3A00%2B08%3A00&#x60;. | 
 **createdAtLte** | **time.Time** | Filters the results to only transactions created before this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. &#x60;2022-01-01T12:00:00+08:00&#x60; must be encoded as &#x60;2022-01-01T12%3A00%3A00%2B08%3A00&#x60;. | 
 **currency** | **[]string** | Filters for transactions that have matching &#x60;currency&#x60; values. The &#x60;currency&#x60; values provided must be formatted as 3-letter ISO currency code. | 
 **externalIdentifier** | **string** | Filters the results to only the items for which the &#x60;external_identifier&#x60; matches this value. | 
 **hasRefunds** | **bool** | When set to &#x60;true&#x60;, filter for transactions that have at least one completed refund associated with it. When set to &#x60;false&#x60;, filter for transactions that have no completed refunds. | 
 **id** | [**string**](string.md) | Filters for the transaction that has a matching &#x60;id&#x60; value. | 
 **metadata** | **[]string** | Filters for transactions where their &#x60;metadata&#x60; values contain all of the provided &#x60;metadata&#x60; keys. The value sent for &#x60;metadata&#x60; must be formatted as a JSON string, and all keys and values must be strings. This value should also be URL encoded.  Duplicate keys are not supported. If a key is duplicated, only the last appearing value will be used. | 
 **method** | **[]string** | Filters the results to only the items for which the &#x60;method&#x60; has been set to this value. | 
 **paymentMethodId** | [**string**](string.md) | The ID of the payment method. | 
 **paymentServiceId** | **[]string** | Filters for transactions that were processed by the provided &#x60;payment_service_id&#x60; values. | 
 **paymentServiceTransactionId** | **string** | Filters for transactions that have a matching &#x60;payment_service_transaction_id&#x60; value. The &#x60;payment_service_transaction_id&#x60; is the identifier of the transaction given by the payment service. | 
 **search** | **string** | Filters for transactions that have one of the following fields match exactly with the provided &#x60;search&#x60; value: * &#x60;buyer_external_identifier&#x60; * &#x60;buyer_id&#x60; * &#x60;external_identifier&#x60; * &#x60;id&#x60; * &#x60;payment_service_transaction_id&#x60; | 
 **status** | **[]string** | Filters the results to only the transactions that have a &#x60;status&#x60; that matches with any of the provided status values. | 
 **updatedAtGte** | **time.Time** | Filters the results to only transactions last updated after this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. &#x60;2022-01-01T12:00:00+08:00&#x60; must be encoded as &#x60;2022-01-01T12%3A00%3A00%2B08%3A00&#x60;. | 
 **updatedAtLte** | **time.Time** | Filters the results to only transactions last updated before this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. &#x60;2022-01-01T12:00:00+08:00&#x60; must be encoded as &#x60;2022-01-01T12%3A00%3A00%2B08%3A00&#x60;. | 
 **beforeCreatedAt** | **time.Time** | Filters the results to only transactions created before this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. &#x60;2022-01-01T12:00:00+08:00&#x60; must be encoded as &#x60;2022-01-01T12%3A00%3A00%2B08%3A00&#x60;.  **WARNING** This filter is deprecated and may be removed eventually, use &#x60;created_at_lte&#x60; instead. | 
 **afterCreatedAt** | **time.Time** | Filters the results to only transactions created after this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. &#x60;2022-01-01T12:00:00+08:00&#x60; must be encoded as &#x60;2022-01-01T12%3A00%3A00%2B08%3A00&#x60;.  **WARNING** This filter is deprecated and may be removed eventually, use &#x60;created_at_gte&#x60; instead. | 
 **beforeUpdatedAt** | **time.Time** | Filters the results to only transactions last updated before this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. &#x60;2022-01-01T12:00:00+08:00&#x60; must be encoded as &#x60;2022-01-01T12%3A00%3A00%2B08%3A00&#x60;.  **WARNING** This filter is deprecated and may be removed eventually, use &#x60;updated_at_lte&#x60; instead. | 
 **afterUpdatedAt** | **time.Time** | Filters the results to only transactions last updated after this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. &#x60;2022-01-01T12:00:00+08:00&#x60; must be encoded as &#x60;2022-01-01T12%3A00%3A00%2B08%3A00&#x60;.  **WARNING** This filter is deprecated and may be removed eventually, use &#x60;updated_at_gte&#x60; instead. | 
 **transactionStatus** | **string** | Filters the results to only the transactions for which the &#x60;status&#x60; matches this value.  **WARNING** This filter is deprecated and may be removed eventually, use &#x60;status&#x60; instead. | 

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

> Refund RefundTransaction(ctx, transactionId).TransactionRefundRequest(transactionRefundRequest).Execute()

Refund transaction



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
    transactionRefundRequest := *openapiclient.NewTransactionRefundRequest() // TransactionRefundRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.RefundTransaction(context.Background(), transactionId).TransactionRefundRequest(transactionRefundRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.RefundTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefundTransaction`: Refund
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

 **transactionRefundRequest** | [**TransactionRefundRequest**](TransactionRefundRequest.md) |  | 

### Return type

[**Refund**](Refund.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidTransaction

> Transaction VoidTransaction(ctx, transactionId).Execute()

Void transaction



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
    resp, r, err := api_client.TransactionsApi.VoidTransaction(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.VoidTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VoidTransaction`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.VoidTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The ID for the transaction to get the information for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoidTransactionRequest struct via the builder pattern


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

