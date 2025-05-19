# MerchantAccounts
(*MerchantAccounts*)

## Overview

### Available Operations

* [List](#list) - List all merchant accounts
* [Create](#create) - Create a merchant account
* [Get](#get) - Get a merchant account
* [Update](#update) - Update a merchant account

## List

List all merchant accounts in an instance.

### Example Usage

```go
package main

import(
	"context"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"os"
	"log"
)

func main() {
    ctx := context.Background()

    s := gr4vygo.New(
        "https://api.example.com",
        gr4vygo.WithSecurity(os.Getenv("GR4VY_O_AUTH2_PASSWORD_BEARER")),
    )

    res, err := s.MerchantAccounts.List(ctx, gr4vygo.String("ZXhhbXBsZTE"), nil, gr4vygo.String("merchant-12345"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CollectionMerchantAccount != nil {
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `cursor`                                                 | **string*                                                | :heavy_minus_sign:                                       | A pointer to the page of results to return.              | ZXhhbXBsZTE                                              |
| `limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | The maximum number of items that are at returned.        | 20                                                       |
| `search`                                                 | **string*                                                | :heavy_minus_sign:                                       | The search term to filter merchant accounts by.          | merchant-12345                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListMerchantAccountsResponse](../../models/operations/listmerchantaccountsresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.Error400                        | 400                                       | application/json                          |
| apierrors.Error401                        | 401                                       | application/json                          |
| apierrors.Response403ListMerchantAccounts | 403                                       | application/json                          |
| apierrors.Error404                        | 404                                       | application/json                          |
| apierrors.Error405                        | 405                                       | application/json                          |
| apierrors.Error409                        | 409                                       | application/json                          |
| apierrors.HTTPValidationError             | 422                                       | application/json                          |
| apierrors.Error425                        | 425                                       | application/json                          |
| apierrors.Error429                        | 429                                       | application/json                          |
| apierrors.Error500                        | 500                                       | application/json                          |
| apierrors.Error502                        | 502                                       | application/json                          |
| apierrors.Error504                        | 504                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## Create

Create a new merchant account in an instance.

### Example Usage

```go
package main

import(
	"context"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"os"
	"github.com/gr4vy/gr4vy-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := gr4vygo.New(
        "https://api.example.com",
        gr4vygo.WithSecurity(os.Getenv("GR4VY_O_AUTH2_PASSWORD_BEARER")),
    )

    res, err := s.MerchantAccounts.Create(ctx, components.MerchantAccountCreate{
        AccountUpdaterRequestEncryptionKey: gr4vygo.String("key-1234"),
        AccountUpdaterRequestEncryptionKeyID: gr4vygo.String("key-id-1234"),
        AccountUpdaterResponseDecryptionKey: gr4vygo.String("key-1234"),
        AccountUpdaterResponseDecryptionKeyID: gr4vygo.String("key-id-1234"),
        OverCaptureAmount: gr4vygo.Int64(1299),
        OverCapturePercentage: gr4vygo.Int64(25),
        LoonClientKey: gr4vygo.String("client-key-1234"),
        LoonSecretKey: gr4vygo.String("key-12345"),
        LoonAcceptedSchemes: []components.CardScheme{
            components.CardSchemeVisa,
        },
        VisaNetworkTokensRequestorID: gr4vygo.String("id-12345"),
        VisaNetworkTokensAppID: gr4vygo.String("id-12345"),
        AmexNetworkTokensRequestorID: gr4vygo.String("id-12345"),
        AmexNetworkTokensAppID: gr4vygo.String("id-12345"),
        MastercardNetworkTokensRequestorID: gr4vygo.String("id-12345"),
        MastercardNetworkTokensAppID: gr4vygo.String("id-12345"),
        OutboundWebhookURL: gr4vygo.String("https://example.com/callback"),
        OutboundWebhookUsername: gr4vygo.String("user-12345"),
        OutboundWebhookPassword: gr4vygo.String("password-12345"),
        ID: "merchant-12345",
        DisplayName: "Example",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.MerchantAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `merchantAccountCreate`                                                              | [components.MerchantAccountCreate](../../models/components/merchantaccountcreate.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `timeoutInSeconds`                                                                   | **float64*                                                                           | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateMerchantAccountResponse](../../models/operations/createmerchantaccountresponse.md), error**

### Errors

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| apierrors.Error400                         | 400                                        | application/json                           |
| apierrors.Error401                         | 401                                        | application/json                           |
| apierrors.Response403CreateMerchantAccount | 403                                        | application/json                           |
| apierrors.Error404                         | 404                                        | application/json                           |
| apierrors.Error405                         | 405                                        | application/json                           |
| apierrors.Error409                         | 409                                        | application/json                           |
| apierrors.HTTPValidationError              | 422                                        | application/json                           |
| apierrors.Error425                         | 425                                        | application/json                           |
| apierrors.Error429                         | 429                                        | application/json                           |
| apierrors.Error500                         | 500                                        | application/json                           |
| apierrors.Error502                         | 502                                        | application/json                           |
| apierrors.Error504                         | 504                                        | application/json                           |
| apierrors.APIError                         | 4XX, 5XX                                   | \*/\*                                      |

## Get

Get info about a merchant account in an instance.

### Example Usage

```go
package main

import(
	"context"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"os"
	"log"
)

func main() {
    ctx := context.Background()

    s := gr4vygo.New(
        "https://api.example.com",
        gr4vygo.WithSecurity(os.Getenv("GR4VY_O_AUTH2_PASSWORD_BEARER")),
    )

    res, err := s.MerchantAccounts.Get(ctx, "merchant-12345")
    if err != nil {
        log.Fatal(err)
    }
    if res.MerchantAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `merchantAccountID`                                      | *string*                                                 | :heavy_check_mark:                                       | The ID of the merchant account                           | merchant-12345                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetMerchantAccountResponse](../../models/operations/getmerchantaccountresponse.md), error**

### Errors

| Error Type                              | Status Code                             | Content Type                            |
| --------------------------------------- | --------------------------------------- | --------------------------------------- |
| apierrors.Error400                      | 400                                     | application/json                        |
| apierrors.Error401                      | 401                                     | application/json                        |
| apierrors.Response403GetMerchantAccount | 403                                     | application/json                        |
| apierrors.Error404                      | 404                                     | application/json                        |
| apierrors.Error405                      | 405                                     | application/json                        |
| apierrors.Error409                      | 409                                     | application/json                        |
| apierrors.HTTPValidationError           | 422                                     | application/json                        |
| apierrors.Error425                      | 425                                     | application/json                        |
| apierrors.Error429                      | 429                                     | application/json                        |
| apierrors.Error500                      | 500                                     | application/json                        |
| apierrors.Error502                      | 502                                     | application/json                        |
| apierrors.Error504                      | 504                                     | application/json                        |
| apierrors.APIError                      | 4XX, 5XX                                | \*/\*                                   |

## Update

Update info for a merchant account in an instance.

### Example Usage

```go
package main

import(
	"context"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"os"
	"github.com/gr4vy/gr4vy-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := gr4vygo.New(
        "https://api.example.com",
        gr4vygo.WithSecurity(os.Getenv("GR4VY_O_AUTH2_PASSWORD_BEARER")),
    )

    res, err := s.MerchantAccounts.Update(ctx, "merchant-12345", components.MerchantAccountUpdate{
        AccountUpdaterRequestEncryptionKey: gr4vygo.String("key-1234"),
        AccountUpdaterRequestEncryptionKeyID: gr4vygo.String("key-id-1234"),
        AccountUpdaterResponseDecryptionKey: gr4vygo.String("key-1234"),
        AccountUpdaterResponseDecryptionKeyID: gr4vygo.String("key-id-1234"),
        OverCaptureAmount: gr4vygo.Int64(1299),
        OverCapturePercentage: gr4vygo.Int64(25),
        LoonClientKey: gr4vygo.String("client-key-1234"),
        LoonSecretKey: gr4vygo.String("key-12345"),
        LoonAcceptedSchemes: []components.CardScheme{
            components.CardSchemeVisa,
        },
        VisaNetworkTokensRequestorID: gr4vygo.String("id-12345"),
        VisaNetworkTokensAppID: gr4vygo.String("id-12345"),
        AmexNetworkTokensRequestorID: gr4vygo.String("id-12345"),
        AmexNetworkTokensAppID: gr4vygo.String("id-12345"),
        MastercardNetworkTokensRequestorID: gr4vygo.String("id-12345"),
        MastercardNetworkTokensAppID: gr4vygo.String("id-12345"),
        DisplayName: gr4vygo.String("Example"),
        OutboundWebhookURL: gr4vygo.String("https://example.com/callback"),
        OutboundWebhookUsername: gr4vygo.String("user-12345"),
        OutboundWebhookPassword: gr4vygo.String("password-12345"),
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.MerchantAccount != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          | Example                                                                              |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |                                                                                      |
| `merchantAccountID`                                                                  | *string*                                                                             | :heavy_check_mark:                                                                   | The ID of the merchant account                                                       | merchant-12345                                                                       |
| `merchantAccountUpdate`                                                              | [components.MerchantAccountUpdate](../../models/components/merchantaccountupdate.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |                                                                                      |
| `timeoutInSeconds`                                                                   | **float64*                                                                           | :heavy_minus_sign:                                                                   | N/A                                                                                  |                                                                                      |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |                                                                                      |

### Response

**[*operations.UpdateMerchantAccountResponse](../../models/operations/updatemerchantaccountresponse.md), error**

### Errors

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| apierrors.Error400                         | 400                                        | application/json                           |
| apierrors.Error401                         | 401                                        | application/json                           |
| apierrors.Response403UpdateMerchantAccount | 403                                        | application/json                           |
| apierrors.Error404                         | 404                                        | application/json                           |
| apierrors.Error405                         | 405                                        | application/json                           |
| apierrors.Error409                         | 409                                        | application/json                           |
| apierrors.HTTPValidationError              | 422                                        | application/json                           |
| apierrors.Error425                         | 425                                        | application/json                           |
| apierrors.Error429                         | 429                                        | application/json                           |
| apierrors.Error500                         | 500                                        | application/json                           |
| apierrors.Error502                         | 502                                        | application/json                           |
| apierrors.Error504                         | 504                                        | application/json                           |
| apierrors.APIError                         | 4XX, 5XX                                   | \*/\*                                      |