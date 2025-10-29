# Transactions
(*Transactions*)

## Overview

### Available Operations

* [List](#list) - List transactions
* [Create](#create) - Create transaction
* [Get](#get) - Get transaction
* [Update](#update) - Manually update a transaction
* [Capture](#capture) - Capture transaction
* [Void](#void) - Void transaction
* [Cancel](#cancel) - Cancel transaction
* [Sync](#sync) - Sync transaction

## List

Returns a paginated list of transactions for the merchant account, sorted by most recently updated. You can filter, sort, and search transactions using query parameters.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_transactions" method="get" path="/transactions" -->
```go
package main

import(
	"context"
	"os"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/types"
	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := gr4vygo.New(
        gr4vygo.WithMerchantAccountID("default"),
        gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
    )

    res, err := s.Transactions.List(ctx, operations.ListTransactionsRequest{
        Cursor: gr4vygo.Pointer("ZXhhbXBsZTE"),
        CreatedAtLte: types.MustNewTimeFromString("2022-01-01T12:00:00+08:00"),
        CreatedAtGte: types.MustNewTimeFromString("2022-01-01T12:00:00+08:00"),
        UpdatedAtLte: types.MustNewTimeFromString("2022-01-01T12:00:00+08:00"),
        UpdatedAtGte: types.MustNewTimeFromString("2022-01-01T12:00:00+08:00"),
        Search: gr4vygo.Pointer("transaction-12345"),
        BuyerExternalIdentifier: gr4vygo.Pointer("buyer-12345"),
        BuyerID: gr4vygo.Pointer("fe26475d-ec3e-4884-9553-f7356683f7f9"),
        BuyerEmailAddress: gr4vygo.Pointer("john@example.com"),
        IPAddress: gr4vygo.Pointer("8.214.133.47"),
        Status: []components.TransactionStatus{
            components.TransactionStatusAuthorizationSucceeded,
        },
        ID: gr4vygo.Pointer("7099948d-7286-47e4-aad8-b68f7eb44591"),
        PaymentServiceTransactionID: gr4vygo.Pointer("tx-12345"),
        ExternalIdentifier: gr4vygo.Pointer("transaction-12345"),
        Metadata: []string{
            "{\"first_key\":\"first_value\",\"second_key\":\"second_value\"}",
        },
        AmountEq: gr4vygo.Pointer[int64](1299),
        AmountLte: gr4vygo.Pointer[int64](1299),
        AmountGte: gr4vygo.Pointer[int64](1299),
        Currency: []string{
            "USD",
        },
        Country: []string{
            "US",
        },
        PaymentServiceID: []string{
            "fffd152a-9532-4087-9a4f-de58754210f0",
        },
        PaymentMethodID: gr4vygo.Pointer("ef9496d8-53a5-4aad-8ca2-00eb68334389"),
        PaymentMethodLabel: gr4vygo.Pointer("1234"),
        PaymentMethodScheme: []string{
            "[",
            "\"",
            "v",
            "i",
            "s",
            "a",
            "\"",
            "]",
        },
        PaymentMethodCountry: gr4vygo.Pointer("[\"US\"]"),
        PaymentMethodFingerprint: gr4vygo.Pointer("a50b85c200ee0795d6fd33a5c66f37a4564f554355c5b46a756aac485dd168a4"),
        Method: []components.Method{
            components.MethodCard,
        },
        ErrorCode: []string{
            "insufficient_funds",
        },
        HasRefunds: gr4vygo.Pointer(true),
        PendingReview: gr4vygo.Pointer(true),
        CheckoutSessionID: gr4vygo.Pointer("4137b1cf-39ac-42a8-bad6-1c680d5dab6b"),
        ReconciliationID: gr4vygo.Pointer("7jZXl4gBUNl0CnaLEnfXbt"),
        HasGiftCardRedemptions: gr4vygo.Pointer(true),
        GiftCardID: gr4vygo.Pointer("356d56e5-fe16-42ae-97ee-8d55d846ae2e"),
        GiftCardLast4: gr4vygo.Pointer("7890"),
        HasSettlements: gr4vygo.Pointer(true),
        PaymentMethodBin: gr4vygo.Pointer("411111"),
        PaymentSource: []components.TransactionPaymentSource{
            components.TransactionPaymentSourceRecurring,
        },
        IsSubsequentPayment: gr4vygo.Pointer(true),
        MerchantInitiated: gr4vygo.Pointer(true),
        Used3ds: gr4vygo.Pointer(true),
        BuyerSearch: []string{
            "J",
            "o",
            "h",
            "n",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListTransactionsRequest](../../models/operations/listtransactionsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListTransactionsResponse](../../models/operations/listtransactionsresponse.md), error**

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

Create a new transaction using a supported payment method. If additional buyer authorization is required, an approval URL will be returned. Duplicated gift card numbers are not supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_transaction" method="post" path="/transactions" -->
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

    res, err := s.Transactions.Create(ctx, components.TransactionCreate{
        Amount: 1299,
        Currency: "EUR",
        Store: gr4vygo.Pointer(true),
        IsSubsequentPayment: gr4vygo.Pointer(true),
        MerchantInitiated: gr4vygo.Pointer(true),
        AsyncCapture: gr4vygo.Pointer(true),
        AccountFundingTransaction: gr4vygo.Pointer(true),
    }, gr4vygo.Pointer("request-12345"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                             | Type                                                                                                                                                                                                  | Required                                                                                                                                                                                              | Description                                                                                                                                                                                           | Example                                                                                                                                                                                               |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                    | The context to use for the request.                                                                                                                                                                   |                                                                                                                                                                                                       |
| `transactionCreate`                                                                                                                                                                                   | [components.TransactionCreate](../../models/components/transactioncreate.md)                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                    | N/A                                                                                                                                                                                                   |                                                                                                                                                                                                       |
| `merchantAccountID`                                                                                                                                                                                   | **string*                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                    | The ID of the merchant account to use for this request.                                                                                                                                               |                                                                                                                                                                                                       |
| `idempotencyKey`                                                                                                                                                                                      | **string*                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                    | A unique key that identifies this request. Providing this header will make this an idempotent request. We recommend using V4 UUIDs, or another random string with enough entropy to avoid collisions. | request-12345                                                                                                                                                                                         |
| `xForwardedFor`                                                                                                                                                                                       | **string*                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                    | The IP address to forward from the customer. Use this when calling<br/>our API from the server side to ensure the customer's address is<br/>passed to downstream services, rather than your server IP. | 192.168.0.2                                                                                                                                                                                           |
| `opts`                                                                                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                    | The options for this request.                                                                                                                                                                         |                                                                                                                                                                                                       |

### Response

**[*components.TransactionOutput](../../models/components/transactionoutput.md), error**

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

Retrieve the details of a transaction by its unique identifier.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_transaction" method="get" path="/transactions/{transaction_id}" -->
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
        gr4vygo.WithMerchantAccountID("default"),
        gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
    )

    res, err := s.Transactions.Get(ctx, "7099948d-7286-47e4-aad8-b68f7eb44591")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `transactionID`                                          | *string*                                                 | :heavy_check_mark:                                       | The ID of the transaction                                | 7099948d-7286-47e4-aad8-b68f7eb44591                     |
| `merchantAccountID`                                      | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.TransactionOutput](../../models/components/transactionoutput.md), error**

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

## Update

Manually updates a transaction.

### Example Usage

<!-- UsageSnippet language="go" operationID="update_transaction" method="put" path="/transactions/{transaction_id}" -->
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
        gr4vygo.WithMerchantAccountID("<id>"),
        gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
    )

    res, err := s.Transactions.Update(ctx, "7099948d-7286-47e4-aad8-b68f7eb44591", components.TransactionUpdate{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |                                                                              |
| `transactionID`                                                              | *string*                                                                     | :heavy_check_mark:                                                           | The ID of the transaction                                                    | 7099948d-7286-47e4-aad8-b68f7eb44591                                         |
| `transactionUpdate`                                                          | [components.TransactionUpdate](../../models/components/transactionupdate.md) | :heavy_check_mark:                                                           | N/A                                                                          |                                                                              |
| `merchantAccountID`                                                          | **string*                                                                    | :heavy_minus_sign:                                                           | The ID of the merchant account to use for this request.                      |                                                                              |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |                                                                              |

### Response

**[*components.TransactionOutput](../../models/components/transactionoutput.md), error**

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

## Capture

Captures a previously authorized transaction. You can capture the full or a partial amount, as long as it does not exceed the authorized amount (unless over-capture is enabled).

### Example Usage

<!-- UsageSnippet language="go" operationID="capture_transaction" method="post" path="/transactions/{transaction_id}/capture" -->
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

    res, err := s.Transactions.Capture(ctx, "7099948d-7286-47e4-aad8-b68f7eb44591", components.TransactionCaptureCreate{}, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `transactionID`                                                                            | *string*                                                                                   | :heavy_check_mark:                                                                         | The ID of the transaction                                                                  | 7099948d-7286-47e4-aad8-b68f7eb44591                                                       |
| `transactionCaptureCreate`                                                                 | [components.TransactionCaptureCreate](../../models/components/transactioncapturecreate.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |
| `prefer`                                                                                   | []*string*                                                                                 | :heavy_minus_sign:                                                                         | The preferred resource type in the response.                                               | resource=transaction                                                                       |
| `merchantAccountID`                                                                        | **string*                                                                                  | :heavy_minus_sign:                                                                         | The ID of the merchant account to use for this request.                                    |                                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |

### Response

**[*operations.ResponseCaptureTransaction](../../models/operations/responsecapturetransaction.md), error**

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

## Void

Voids a previously authorized transaction. If the transaction was not yet successfully authorized, or was already captured, the void will not be processed. This operation releases the hold on the buyer's funds. Captured transactions can be refunded instead.

### Example Usage

<!-- UsageSnippet language="go" operationID="void_transaction" method="post" path="/transactions/{transaction_id}/void" -->
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
        gr4vygo.WithMerchantAccountID("default"),
        gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
    )

    res, err := s.Transactions.Void(ctx, "7099948d-7286-47e4-aad8-b68f7eb44591", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `transactionID`                                          | *string*                                                 | :heavy_check_mark:                                       | The ID of the transaction                                | 7099948d-7286-47e4-aad8-b68f7eb44591                     |
| `prefer`                                                 | []*string*                                               | :heavy_minus_sign:                                       | The preferred resource type in the response.             | resource=transaction                                     |
| `merchantAccountID`                                      | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ResponseVoidTransaction](../../models/operations/responsevoidtransaction.md), error**

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

## Cancel

Cancels a pending transaction. If the transaction was successfully authorized, or was already captured, the cancel will not be processed.

### Example Usage

<!-- UsageSnippet language="go" operationID="cancel_transaction" method="post" path="/transactions/{transaction_id}/cancel" -->
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
        gr4vygo.WithMerchantAccountID("<id>"),
        gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
    )

    res, err := s.Transactions.Cancel(ctx, "7099948d-7286-47e4-aad8-b68f7eb44591")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `transactionID`                                          | *string*                                                 | :heavy_check_mark:                                       | The ID of the transaction                                | 7099948d-7286-47e4-aad8-b68f7eb44591                     |
| `merchantAccountID`                                      | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.TransactionCancel](../../models/components/transactioncancel.md), error**

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

## Sync

Synchronizes the status of a transaction with the underlying payment service provider. This is useful for transactions in a pending state to check if they've been completed or failed. Only available for some payment service providers.

### Example Usage

<!-- UsageSnippet language="go" operationID="sync_transaction" method="post" path="/transactions/{transaction_id}/sync" -->
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
        gr4vygo.WithMerchantAccountID("default"),
        gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
    )

    res, err := s.Transactions.Sync(ctx, "2ee546e0-3b11-478e-afec-fdb362611e22")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `transactionID`                                          | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `merchantAccountID`                                      | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.TransactionOutput](../../models/components/transactionoutput.md), error**

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