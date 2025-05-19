# GiftCards
(*GiftCards*)

## Overview

### Available Operations

* [Get](#get) - Get gift card
* [Delete](#delete) - Delete a gift card
* [Create](#create) - Create gift card
* [List](#list) - List gift cards

## Get

Fetch details about a gift card.

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

    res, err := s.GiftCards.Get(ctx, "356d56e5-fe16-42ae-97ee-8d55d846ae2e", gr4vygo.String("default"))
    if err != nil {
        log.Fatal(err)
    }
    if res.GiftCard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `giftCardID`                                             | *string*                                                 | :heavy_check_mark:                                       | The ID of the gift card.                                 | 356d56e5-fe16-42ae-97ee-8d55d846ae2e                     |
| `xGr4vyMerchantAccountID`                                | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  | default                                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetGiftCardResponse](../../models/operations/getgiftcardresponse.md), error**

### Errors

| Error Type                       | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| apierrors.Error400               | 400                              | application/json                 |
| apierrors.Error401               | 401                              | application/json                 |
| apierrors.Response403GetGiftCard | 403                              | application/json                 |
| apierrors.Error404               | 404                              | application/json                 |
| apierrors.Error405               | 405                              | application/json                 |
| apierrors.Error409               | 409                              | application/json                 |
| apierrors.HTTPValidationError    | 422                              | application/json                 |
| apierrors.Error425               | 425                              | application/json                 |
| apierrors.Error429               | 429                              | application/json                 |
| apierrors.Error500               | 500                              | application/json                 |
| apierrors.Error502               | 502                              | application/json                 |
| apierrors.Error504               | 504                              | application/json                 |
| apierrors.APIError               | 4XX, 5XX                         | \*/\*                            |

## Delete

Removes a gift card from our system.

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

    res, err := s.GiftCards.Delete(ctx, "356d56e5-fe16-42ae-97ee-8d55d846ae2e", nil, gr4vygo.String("default"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Any != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `giftCardID`                                             | *string*                                                 | :heavy_check_mark:                                       | The ID of the gift card.                                 | 356d56e5-fe16-42ae-97ee-8d55d846ae2e                     |
| `timeoutInSeconds`                                       | **float64*                                               | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `xGr4vyMerchantAccountID`                                | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  | default                                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteGiftCardResponse](../../models/operations/deletegiftcardresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| apierrors.Error400                  | 400                                 | application/json                    |
| apierrors.Error401                  | 401                                 | application/json                    |
| apierrors.Response403DeleteGiftCard | 403                                 | application/json                    |
| apierrors.Error404                  | 404                                 | application/json                    |
| apierrors.Error405                  | 405                                 | application/json                    |
| apierrors.Error409                  | 409                                 | application/json                    |
| apierrors.HTTPValidationError       | 422                                 | application/json                    |
| apierrors.Error425                  | 425                                 | application/json                    |
| apierrors.Error429                  | 429                                 | application/json                    |
| apierrors.Error500                  | 500                                 | application/json                    |
| apierrors.Error502                  | 502                                 | application/json                    |
| apierrors.Error504                  | 504                                 | application/json                    |
| apierrors.APIError                  | 4XX, 5XX                            | \*/\*                               |

## Create

Store a new gift card in the vault.

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

    res, err := s.GiftCards.Create(ctx, components.GiftCardCreate{
        Number: "4123455541234561234",
        Pin: "1234",
        BuyerID: gr4vygo.String("fe26475d-ec3e-4884-9553-f7356683f7f9"),
        BuyerExternalIdentifier: gr4vygo.String("buyer-12345"),
    }, nil, gr4vygo.String("default"))
    if err != nil {
        log.Fatal(err)
    }
    if res.GiftCard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |                                                                        |
| `giftCardCreate`                                                       | [components.GiftCardCreate](../../models/components/giftcardcreate.md) | :heavy_check_mark:                                                     | N/A                                                                    |                                                                        |
| `timeoutInSeconds`                                                     | **float64*                                                             | :heavy_minus_sign:                                                     | N/A                                                                    |                                                                        |
| `xGr4vyMerchantAccountID`                                              | **string*                                                              | :heavy_minus_sign:                                                     | The ID of the merchant account to use for this request.                | default                                                                |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |                                                                        |

### Response

**[*operations.CreateGiftCardResponse](../../models/operations/creategiftcardresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| apierrors.Error400                  | 400                                 | application/json                    |
| apierrors.Error401                  | 401                                 | application/json                    |
| apierrors.Response403CreateGiftCard | 403                                 | application/json                    |
| apierrors.Error404                  | 404                                 | application/json                    |
| apierrors.Error405                  | 405                                 | application/json                    |
| apierrors.Error409                  | 409                                 | application/json                    |
| apierrors.HTTPValidationError       | 422                                 | application/json                    |
| apierrors.Error425                  | 425                                 | application/json                    |
| apierrors.Error429                  | 429                                 | application/json                    |
| apierrors.Error500                  | 500                                 | application/json                    |
| apierrors.Error502                  | 502                                 | application/json                    |
| apierrors.Error504                  | 504                                 | application/json                    |
| apierrors.APIError                  | 4XX, 5XX                            | \*/\*                               |

## List

Browser all gift cards.

### Example Usage

```go
package main

import(
	"context"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"os"
	"github.com/gr4vy/gr4vy-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := gr4vygo.New(
        "https://api.example.com",
        gr4vygo.WithSecurity(os.Getenv("GR4VY_O_AUTH2_PASSWORD_BEARER")),
    )

    res, err := s.GiftCards.List(ctx, operations.ListGiftCardsRequest{
        XGr4vyMerchantAccountID: gr4vygo.String("default"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CollectionGiftCard != nil {
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListGiftCardsRequest](../../models/operations/listgiftcardsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ListGiftCardsResponse](../../models/operations/listgiftcardsresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.Error400                 | 400                                | application/json                   |
| apierrors.Error401                 | 401                                | application/json                   |
| apierrors.Response403ListGiftCards | 403                                | application/json                   |
| apierrors.Error404                 | 404                                | application/json                   |
| apierrors.Error405                 | 405                                | application/json                   |
| apierrors.Error409                 | 409                                | application/json                   |
| apierrors.HTTPValidationError      | 422                                | application/json                   |
| apierrors.Error425                 | 425                                | application/json                   |
| apierrors.Error429                 | 429                                | application/json                   |
| apierrors.Error500                 | 500                                | application/json                   |
| apierrors.Error502                 | 502                                | application/json                   |
| apierrors.Error504                 | 504                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |