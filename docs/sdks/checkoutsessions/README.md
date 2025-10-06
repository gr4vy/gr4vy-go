# CheckoutSessions
(*CheckoutSessions*)

## Overview

### Available Operations

* [Create](#create) - Create checkout session
* [Update](#update) - Update checkout session
* [Get](#get) - Get checkout session
* [Delete](#delete) - Delete checkout session

## Create

Create a new checkout session.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_checkout_session" method="post" path="/checkout/sessions" -->
```go
package main

import(
	"context"
	"os"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := gr4vygo.New(
        gr4vygo.WithMerchantAccountID("default"),
        gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
    )

    res, err := s.CheckoutSessions.Create(ctx, &components.CheckoutSessionCreate{
        CartItems: []components.CartItem{
            components.CartItem{
                Name: "GoPro HD",
                Quantity: 2,
                UnitAmount: 1299,
                DiscountAmount: gr4vygo.Pointer[int64](0),
                TaxAmount: gr4vygo.Pointer[int64](0),
                ExternalIdentifier: gr4vygo.Pointer("goprohd"),
                Sku: gr4vygo.Pointer("GPHD1078"),
                ProductURL: gr4vygo.Pointer("https://example.com/catalog/go-pro-hd"),
                ImageURL: gr4vygo.Pointer("https://example.com/images/go-pro-hd.jpg"),
                Categories: []string{
                    "camera",
                    "travel",
                    "gear",
                },
                ProductType: components.ProductTypePhysical.ToPointer(),
                SellerCountry: gr4vygo.Pointer("US"),
            },
            components.CartItem{
                Name: "GoPro HD",
                Quantity: 2,
                UnitAmount: 1299,
                DiscountAmount: gr4vygo.Pointer[int64](0),
                TaxAmount: gr4vygo.Pointer[int64](0),
                ExternalIdentifier: gr4vygo.Pointer("goprohd"),
                Sku: gr4vygo.Pointer("GPHD1078"),
                ProductURL: gr4vygo.Pointer("https://example.com/catalog/go-pro-hd"),
                ImageURL: gr4vygo.Pointer("https://example.com/images/go-pro-hd.jpg"),
                Categories: []string{
                    "camera",
                    "travel",
                    "gear",
                },
                ProductType: components.ProductTypePhysical.ToPointer(),
                SellerCountry: gr4vygo.Pointer("US"),
            },
        },
        Metadata: map[string]string{
            "cohort": "cohort-a",
            "order_id": "order-12345",
        },
        Buyer: nil,
        Airline: &components.Airline{
            BookingCode: gr4vygo.Pointer("X36Q9C"),
            IsCardholderTraveling: gr4vygo.Pointer(true),
            IssuedAddress: gr4vygo.Pointer("123 Broadway, New York"),
            IssuedAt: types.MustNewTimeFromString("2013-07-16T19:23:00.000+00:00"),
            IssuingCarrierCode: gr4vygo.Pointer("649"),
            IssuingCarrierName: gr4vygo.Pointer("Air Transat A.T. Inc"),
            IssuingIataDesignator: gr4vygo.Pointer("TS"),
            IssuingIcaoCode: gr4vygo.Pointer("TSC"),
            Legs: []components.AirlineLeg{
                components.AirlineLeg{
                    ArrivalAirport: gr4vygo.Pointer("LAX"),
                    ArrivalAt: types.MustNewTimeFromString("2013-07-16T19:23:00.000+00:00"),
                    ArrivalCity: gr4vygo.Pointer("Los Angeles"),
                    ArrivalCountry: gr4vygo.Pointer("US"),
                    CarrierCode: gr4vygo.Pointer("649"),
                    CarrierName: gr4vygo.Pointer("Air Transat A.T. Inc"),
                    IataDesignator: gr4vygo.Pointer("TS"),
                    IcaoCode: gr4vygo.Pointer("TSC"),
                    CouponNumber: gr4vygo.Pointer("15885566"),
                    DepartureAirport: gr4vygo.Pointer("LHR"),
                    DepartureAt: types.MustNewTimeFromString("2013-07-16T19:23:00.000+00:00"),
                    DepartureCity: gr4vygo.Pointer("London"),
                    DepartureCountry: gr4vygo.Pointer("GB"),
                    DepartureTaxAmount: gr4vygo.Pointer[int64](1200),
                    FareAmount: gr4vygo.Pointer[int64](129900),
                    FareBasisCode: gr4vygo.Pointer("FY"),
                    FeeAmount: gr4vygo.Pointer[int64](1200),
                    FlightClass: gr4vygo.Pointer("E"),
                    FlightNumber: gr4vygo.Pointer("101"),
                    RouteType: components.RouteTypeRoundTrip.ToPointer(),
                    SeatClass: gr4vygo.Pointer("F"),
                    StopOver: gr4vygo.Pointer(false),
                    TaxAmount: gr4vygo.Pointer[int64](1200),
                },
            },
            PassengerNameRecord: gr4vygo.Pointer("JOHN L"),
            Passengers: []components.AirlinePassenger{
                components.AirlinePassenger{
                    AgeGroup: components.AgeGroupAdult.ToPointer(),
                    DateOfBirth: types.MustNewDateFromString("2013-07-16"),
                    EmailAddress: gr4vygo.Pointer("john@example.com"),
                    FirstName: gr4vygo.Pointer("John"),
                    FrequentFlyerNumber: gr4vygo.Pointer("15885566"),
                    LastName: gr4vygo.Pointer("Luhn"),
                    PassportNumber: gr4vygo.Pointer("11117700225"),
                    PhoneNumber: gr4vygo.Pointer("+1234567890"),
                    TicketNumber: gr4vygo.Pointer("BA1236699999"),
                    Title: gr4vygo.Pointer("Mr."),
                    CountryCode: gr4vygo.Pointer("US"),
                },
                components.AirlinePassenger{
                    AgeGroup: components.AgeGroupAdult.ToPointer(),
                    DateOfBirth: types.MustNewDateFromString("2013-07-16"),
                    EmailAddress: gr4vygo.Pointer("john@example.com"),
                    FirstName: gr4vygo.Pointer("John"),
                    FrequentFlyerNumber: gr4vygo.Pointer("15885566"),
                    LastName: gr4vygo.Pointer("Luhn"),
                    PassportNumber: gr4vygo.Pointer("11117700225"),
                    PhoneNumber: gr4vygo.Pointer("+1234567890"),
                    TicketNumber: gr4vygo.Pointer("BA1236699999"),
                    Title: gr4vygo.Pointer("Mr."),
                    CountryCode: gr4vygo.Pointer("US"),
                },
            },
            ReservationSystem: gr4vygo.Pointer("Amadeus"),
            RestrictedTicket: gr4vygo.Pointer(false),
            TicketDeliveryMethod: components.TicketDeliveryMethodElectronic.ToPointer(),
            TicketNumber: gr4vygo.Pointer("123-1234-151555"),
            TravelAgencyCode: gr4vygo.Pointer("12345"),
            TravelAgencyInvoiceNumber: gr4vygo.Pointer("EG15555155"),
            TravelAgencyName: gr4vygo.Pointer("ACME Agency"),
            TravelAgencyPlanName: gr4vygo.Pointer("B733"),
        },
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

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |
| `merchantAccountID`                                                                   | **string*                                                                             | :heavy_minus_sign:                                                                    | The ID of the merchant account to use for this request.                               |
| `checkoutSessionCreate`                                                               | [*components.CheckoutSessionCreate](../../models/components/checkoutsessioncreate.md) | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `opts`                                                                                | [][operations.Option](../../models/operations/option.md)                              | :heavy_minus_sign:                                                                    | The options for this request.                                                         |

### Response

**[*components.CheckoutSession](../../models/components/checkoutsession.md), error**

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

Update the information stored on a checkout session.

### Example Usage

<!-- UsageSnippet language="go" operationID="update_checkout_session" method="put" path="/checkout/sessions/{session_id}" -->
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

    res, err := s.CheckoutSessions.Update(ctx, "4137b1cf-39ac-42a8-bad6-1c680d5dab6b", components.CheckoutSessionCreate{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          | Example                                                                              |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |                                                                                      |
| `sessionID`                                                                          | *string*                                                                             | :heavy_check_mark:                                                                   | The ID of the checkout session.                                                      | 4137b1cf-39ac-42a8-bad6-1c680d5dab6b                                                 |
| `checkoutSessionCreate`                                                              | [components.CheckoutSessionCreate](../../models/components/checkoutsessioncreate.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |                                                                                      |
| `merchantAccountID`                                                                  | **string*                                                                            | :heavy_minus_sign:                                                                   | The ID of the merchant account to use for this request.                              |                                                                                      |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |                                                                                      |

### Response

**[*components.CheckoutSession](../../models/components/checkoutsession.md), error**

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

Retrieve the information stored on a checkout session.

### Example Usage

<!-- UsageSnippet language="go" operationID="get_checkout_session" method="get" path="/checkout/sessions/{session_id}" -->
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

    res, err := s.CheckoutSessions.Get(ctx, "4137b1cf-39ac-42a8-bad6-1c680d5dab6b")
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
| `sessionID`                                              | *string*                                                 | :heavy_check_mark:                                       | The ID of the checkout session.                          | 4137b1cf-39ac-42a8-bad6-1c680d5dab6b                     |
| `merchantAccountID`                                      | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*components.CheckoutSession](../../models/components/checkoutsession.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error400 | 400                | application/json   |
| apierrors.Error401 | 401                | application/json   |
| apierrors.Error403 | 403                | application/json   |
| apierrors.Error404 | 404                | application/json   |
| apierrors.Error405 | 405                | application/json   |
| apierrors.Error409 | 409                | application/json   |
| apierrors.Error425 | 425                | application/json   |
| apierrors.Error429 | 429                | application/json   |
| apierrors.Error500 | 500                | application/json   |
| apierrors.Error502 | 502                | application/json   |
| apierrors.Error504 | 504                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a checkout session and all of its (PCI) data.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_checkout_session" method="delete" path="/checkout/sessions/{session_id}" -->
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

    err := s.CheckoutSessions.Delete(ctx, "4137b1cf-39ac-42a8-bad6-1c680d5dab6b")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `sessionID`                                              | *string*                                                 | :heavy_check_mark:                                       | The ID of the checkout session.                          | 4137b1cf-39ac-42a8-bad6-1c680d5dab6b                     |
| `merchantAccountID`                                      | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**error**

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