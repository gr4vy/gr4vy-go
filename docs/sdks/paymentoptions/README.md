# PaymentOptions
(*PaymentOptions*)

## Overview

### Available Operations

* [List](#list) - List payment options

## List

List the payment options available at checkout. filtering by country, currency, and additional fields passed to Flow rules.

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

    res, err := s.PaymentOptions.List(ctx, components.PaymentOptionRequest{
        Metadata: map[string]string{
            "cohort": "a",
        },
        Country: gr4vygo.String("US"),
        Currency: gr4vygo.String("USD"),
        Amount: gr4vygo.Int64(1299),
        CartItems: []components.CartItem{
            components.CartItem{
                Name: "GoPro HD",
                Quantity: 2,
                UnitAmount: 1299,
                DiscountAmount: gr4vygo.Int64(0),
                TaxAmount: gr4vygo.Int64(0),
                ExternalIdentifier: gr4vygo.String("goprohd"),
                Sku: gr4vygo.String("GPHD1078"),
                ProductURL: gr4vygo.String("https://example.com/catalog/go-pro-hd"),
                ImageURL: gr4vygo.String("https://example.com/images/go-pro-hd.jpg"),
                Categories: []string{
                    "camera",
                    "travel",
                    "gear",
                },
                ProductType: components.ProductTypePhysical.ToPointer(),
                SellerCountry: gr4vygo.String("US"),
            },
        },
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `paymentOptionRequest`                                                             | [components.PaymentOptionRequest](../../models/components/paymentoptionrequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `merchantAccountID`                                                                | **string*                                                                          | :heavy_minus_sign:                                                                 | The ID of the merchant account to use for this request.                            |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*components.CollectionNoCursorPaymentOption](../../models/components/collectionnocursorpaymentoption.md), error**

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