# GiftCards.Activations

## Overview

### Available Operations

* [Create](#create) - Activate a gift card

## Create

Activate a physical gift card through the primary gift card service. Set `store` to `true` to also store the activated gift card.

### Example Usage

<!-- UsageSnippet language="go" operationID="activate_gift_card" method="post" path="/gift-cards/activations" -->
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

    res, err := s.GiftCards.Activations.Create(ctx, components.GiftCardActivationCreate{
        Number: "4123455541234561234",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                       | Type                                                                                                                                                                                                                                            | Required                                                                                                                                                                                                                                        | Description                                                                                                                                                                                                                                     |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                                              | The context to use for the request.                                                                                                                                                                                                             |
| `giftCardActivationCreate`                                                                                                                                                                                                                      | [components.GiftCardActivationCreate](../../models/components/giftcardactivationcreate.md)                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                              | N/A                                                                                                                                                                                                                                             |
| `idempotencyKey`                                                                                                                                                                                                                                | `*string`                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                              | A unique key that identifies this request. If supported by the gift card service, the value will be forwarded to make the activation idempotent. We recommend using V4 UUIDs, or another random string with enough entropy to avoid collisions. |
| `merchantAccountID`                                                                                                                                                                                                                             | `*string`                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                              | The ID of the merchant account to use for this request.                                                                                                                                                                                         |
| `opts`                                                                                                                                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                              | The options for this request.                                                                                                                                                                                                                   |

### Response

**[*components.GiftCard](../../models/components/giftcard.md), error**

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