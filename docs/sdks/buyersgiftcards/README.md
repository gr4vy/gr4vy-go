# BuyersGiftCards
(*Buyers.GiftCards*)

## Overview

### Available Operations

* [List](#list) - List gift cards for a buyer

## List

List all the stored gift cards for a specific buyer.

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

    res, err := s.Buyers.GiftCards.List(ctx, nil, nil, nil, gr4vygo.String("default"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CollectionNoCursorGiftCardSummary != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `buyerExternalIdentifier`                                | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `buyerID`                                                | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `timeoutInSeconds`                                       | **float64*                                               | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `xGr4vyMerchantAccountID`                                | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  | default                                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListBuyerGiftCardsResponse](../../models/operations/listbuyergiftcardsresponse.md), error**

### Errors

| Error Type                              | Status Code                             | Content Type                            |
| --------------------------------------- | --------------------------------------- | --------------------------------------- |
| apierrors.Error400                      | 400                                     | application/json                        |
| apierrors.Error401                      | 401                                     | application/json                        |
| apierrors.Response403ListBuyerGiftCards | 403                                     | application/json                        |
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