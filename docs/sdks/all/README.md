# All
(*Transactions.Refunds.All*)

## Overview

### Available Operations

* [Create](#create) - Create batch transaction refund

## Create

Create a refund for all instruments on a transaction.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_full_transaction_refund" method="post" path="/transactions/{transaction_id}/refunds/all" -->
```go
package main

import(
	"context"
	"os"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := gr4vygo.New(
        gr4vygo.WithMerchantAccountID("default"),
        gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
    )

    res, err := s.Transactions.Refunds.All.Create(ctx, "7099948d-7286-47e4-aad8-b68f7eb44591", &components.TransactionRefundAllCreate{
        Reason: gr4vygo.Pointer("Refund due to user request."),
        ExternalIdentifier: gr4vygo.Pointer("refund-12345"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                       | Type                                                                                            | Required                                                                                        | Description                                                                                     | Example                                                                                         |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `ctx`                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                           | :heavy_check_mark:                                                                              | The context to use for the request.                                                             |                                                                                                 |
| `transactionID`                                                                                 | *string*                                                                                        | :heavy_check_mark:                                                                              | The ID of the transaction                                                                       | 7099948d-7286-47e4-aad8-b68f7eb44591                                                            |
| `merchantAccountID`                                                                             | **string*                                                                                       | :heavy_minus_sign:                                                                              | The ID of the merchant account to use for this request.                                         |                                                                                                 |
| `transactionRefundAllCreate`                                                                    | [*components.TransactionRefundAllCreate](../../models/components/transactionrefundallcreate.md) | :heavy_minus_sign:                                                                              | N/A                                                                                             |                                                                                                 |
| `opts`                                                                                          | [][operations.Option](../../models/operations/option.md)                                        | :heavy_minus_sign:                                                                              | The options for this request.                                                                   |                                                                                                 |

### Response

**[*components.Refunds](../../models/components/refunds.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.Error400            | 400                           | application/json              |
| apierrors.Error401            | 401                           | application/json              |
| apierrors.Error403            | 403                           | application/json              |
| apierrors.Error404            | 404                           | application/json              |
| apierrors.Error405            | 405                           | application/json              |
| apierrors.Error409            | 409                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.Error425            | 425                           | application/json              |
| apierrors.Error429            | 429                           | application/json              |
| apierrors.Error500            | 500                           | application/json              |
| apierrors.Error502            | 502                           | application/json              |
| apierrors.Error504            | 504                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |