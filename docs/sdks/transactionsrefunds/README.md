# TransactionsRefunds
(*Transactions.Refunds*)

## Overview

### Available Operations

* [List](#list) - List transaction refunds
* [Create](#create) - Create transaction refund
* [Get](#get) - Get transaction refund

## List

List refunds for a transaction.

### Example Usage

```go
package main

import(
	"context"
	"os"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := gr4vygo.New(
        gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
    )

    res, err := s.Transactions.Refunds.List(ctx, "7099948d-7286-47e4-aad8-b68f7eb44591", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CollectionRefund != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `transactionID`                                          | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 7099948d-7286-47e4-aad8-b68f7eb44591                     |
| `merchantAccountID`                                      | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListTransactionRefundsResponse](../../models/operations/listtransactionrefundsresponse.md), error**

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

## Create

Create a refund for a transaction.

### Example Usage

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
        gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
    )

    res, err := s.Transactions.Refunds.Create(ctx, "7099948d-7286-47e4-aad8-b68f7eb44591", components.TransactionRefundCreate{
        Amount: gr4vygo.Int64(1299),
        TargetID: gr4vygo.String("7a6c366d-9205-45ab-8021-0d9ee37f20f2"),
        Reason: gr4vygo.String("Refund due to user request."),
        ExternalIdentifier: gr4vygo.String("refund-12345"),
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Refund != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `transactionID`                                                                          | *string*                                                                                 | :heavy_check_mark:                                                                       | N/A                                                                                      | 7099948d-7286-47e4-aad8-b68f7eb44591                                                     |
| `transactionRefundCreate`                                                                | [components.TransactionRefundCreate](../../models/components/transactionrefundcreate.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |                                                                                          |
| `timeoutInSeconds`                                                                       | **float64*                                                                               | :heavy_minus_sign:                                                                       | N/A                                                                                      |                                                                                          |
| `merchantAccountID`                                                                      | **string*                                                                                | :heavy_minus_sign:                                                                       | The ID of the merchant account to use for this request.                                  |                                                                                          |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |                                                                                          |

### Response

**[*operations.CreateTransactionRefundResponse](../../models/operations/createtransactionrefundresponse.md), error**

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

## Get

Fetch refund for a transaction.

### Example Usage

```go
package main

import(
	"context"
	"os"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := gr4vygo.New(
        gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
    )

    res, err := s.Transactions.Refunds.Get(ctx, "7099948d-7286-47e4-aad8-b68f7eb44591", "6a1d4e46-14ed-4fe1-a45f-eff4e025d211", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Refund != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `transactionID`                                          | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 7099948d-7286-47e4-aad8-b68f7eb44591                     |
| `refundID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 6a1d4e46-14ed-4fe1-a45f-eff4e025d211                     |
| `merchantAccountID`                                      | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetTransactionRefundResponse](../../models/operations/gettransactionrefundresponse.md), error**

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