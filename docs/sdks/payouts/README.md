# Payouts
(*Payouts*)

## Overview

### Available Operations

* [List](#list) - List payouts created.
* [Create](#create) - Create a payout.
* [Get](#get) - Get a payout.

## List

Returns a list of payouts made.

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

    res, err := s.Payouts.List(ctx, gr4vygo.String("ZXhhbXBsZTE"), nil, nil)
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `cursor`                                                 | **string*                                                | :heavy_minus_sign:                                       | A pointer to the page of results to return.              | ZXhhbXBsZTE                                              |
| `limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | The maximum number of items that are at returned.        | 20                                                       |
| `merchantAccountID`                                      | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListPayoutsResponse](../../models/operations/listpayoutsresponse.md), error**

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

Creates a new payout.

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

    res, err := s.Payouts.Create(ctx, components.PayoutCreate{
        Amount: 1299,
        Currency: "EUR",
        PaymentServiceID: "ed8bd87d-85ad-40cf-8e8f-007e21e55aad",
        PaymentMethod: components.CreatePayoutCreatePaymentMethodPaymentMethodStoredCard(
            components.PaymentMethodStoredCard{
                ID: "852b951c-d7ea-4c98-b09e-4a1c9e97c077",
            },
        ),
        Category: components.PayoutCategoryOnlineGambling.ToPointer(),
        ExternalIdentifier: gr4vygo.String("payout-12345"),
        BuyerID: gr4vygo.String("fe26475d-ec3e-4884-9553-f7356683f7f9"),
        Buyer: &components.GuestBuyerInput{
            DisplayName: gr4vygo.String("John Doe"),
            ExternalIdentifier: gr4vygo.String("buyer-12345"),
            BillingDetails: &components.BillingDetailsInput{
                FirstName: gr4vygo.String("John"),
                LastName: gr4vygo.String("Doe"),
                EmailAddress: gr4vygo.String("john@example.com"),
                PhoneNumber: gr4vygo.String("+1234567890"),
                Address: &components.Address{
                    City: gr4vygo.String("San Jose"),
                    Country: gr4vygo.String("US"),
                    PostalCode: gr4vygo.String("94560"),
                    State: gr4vygo.String("California"),
                    StateCode: gr4vygo.String("US-CA"),
                    HouseNumberOrName: gr4vygo.String("10"),
                    Line1: gr4vygo.String("Stafford Appartments"),
                    Line2: gr4vygo.String("29th Street"),
                    Organization: gr4vygo.String("Gr4vy"),
                },
                TaxID: &components.TaxID{
                    Value: "12345678931",
                    Kind: components.TaxIDKindNoVat,
                },
            },
            ShippingDetails: &components.ShippingDetailsCreate{
                FirstName: gr4vygo.String("John"),
                LastName: gr4vygo.String("Doe"),
                EmailAddress: gr4vygo.String("john@example.com"),
                PhoneNumber: gr4vygo.String("+1234567890"),
                Address: &components.Address{
                    City: gr4vygo.String("San Jose"),
                    Country: gr4vygo.String("US"),
                    PostalCode: gr4vygo.String("94560"),
                    State: gr4vygo.String("California"),
                    StateCode: gr4vygo.String("US-CA"),
                    HouseNumberOrName: gr4vygo.String("10"),
                    Line1: gr4vygo.String("Stafford Appartments"),
                    Line2: gr4vygo.String("29th Street"),
                    Organization: gr4vygo.String("Gr4vy"),
                },
            },
        },
        BuyerExternalIdentifier: gr4vygo.String("buyer-12345"),
        Merchant: &components.PayoutMerchant{
            Name: "Acme Inc",
            IdentificationNumber: "12345",
            PhoneNumber: "+14155552671",
            URL: "https://example.com",
            StatementDescriptor: "Winnings",
            MerchantCategoryCode: "123456",
            Address: &components.Address{
                City: gr4vygo.String("San Jose"),
                Country: gr4vygo.String("US"),
                PostalCode: gr4vygo.String("94560"),
                State: gr4vygo.String("California"),
                StateCode: gr4vygo.String("US-CA"),
                HouseNumberOrName: gr4vygo.String("10"),
                Line1: gr4vygo.String("Stafford Appartments"),
                Line2: gr4vygo.String("29th Street"),
                Organization: gr4vygo.String("Gr4vy"),
            },
        },
        ConnectionOptions: &components.ConnectionOptions{
            CheckoutCard: &components.CheckoutCardConnectionOptions{
                ProcessingChannelID: "channel-1234",
                SourceID: "acct-1234",
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

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `payoutCreate`                                                     | [components.PayoutCreate](../../models/components/payoutcreate.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `merchantAccountID`                                                | **string*                                                          | :heavy_minus_sign:                                                 | The ID of the merchant account to use for this request.            |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*components.PayoutSummary](../../models/components/payoutsummary.md), error**

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

Retreives a payout.

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

    res, err := s.Payouts.Get(ctx, "4344fef2-bc2f-49a6-924f-343e62f67224", nil)
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
| `payoutID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `merchantAccountID`                                      | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*components.PayoutSummary](../../models/components/payoutsummary.md), error**

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