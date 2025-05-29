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
        gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
    )

    res, err := s.CheckoutSessions.Create(ctx, nil, &components.CheckoutSessionCreate{
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
        Metadata: map[string]string{
            "cohort": "cohort-a",
            "order_id": "order-12345",
        },
        Buyer: nil,
        Airline: &components.Airline{
            BookingCode: gr4vygo.String("X36Q9C"),
            IsCardholderTraveling: gr4vygo.Bool(true),
            IssuedAddress: gr4vygo.String("123 Broadway, New York"),
            IssuedAt: types.MustNewTimeFromString("2013-07-16T19:23:00.000+00:00"),
            IssuingCarrierCode: gr4vygo.String("649"),
            IssuingCarrierName: gr4vygo.String("Air Transat A.T. Inc"),
            IssuingIataDesignator: gr4vygo.String("TS"),
            IssuingIcaoCode: gr4vygo.String("TSC"),
            Legs: []components.AirlineLeg{
                components.AirlineLeg{
                    ArrivalAirport: gr4vygo.String("LAX"),
                    ArrivalAt: types.MustNewTimeFromString("2013-07-16T19:23:00.000+00:00"),
                    ArrivalCity: gr4vygo.String("Los Angeles"),
                    ArrivalCountry: gr4vygo.String("US"),
                    CarrierCode: gr4vygo.String("649"),
                    CarrierName: gr4vygo.String("Air Transat A.T. Inc"),
                    IataDesignator: gr4vygo.String("TS"),
                    IcaoCode: gr4vygo.String("TSC"),
                    CouponNumber: gr4vygo.String("15885566"),
                    DepartureAirport: gr4vygo.String("LHR"),
                    DepartureAt: types.MustNewTimeFromString("2013-07-16T19:23:00.000+00:00"),
                    DepartureCity: gr4vygo.String("London"),
                    DepartureCountry: gr4vygo.String("GB"),
                    DepartureTaxAmount: gr4vygo.Int64(1200),
                    FareAmount: gr4vygo.Int64(129900),
                    FareBasisCode: gr4vygo.String("FY"),
                    FeeAmount: gr4vygo.Int64(1200),
                    FlightClass: gr4vygo.String("E"),
                    FlightNumber: gr4vygo.String("101"),
                    RouteType: components.RouteTypeRoundTrip.ToPointer(),
                    SeatClass: gr4vygo.String("F"),
                    StopOver: gr4vygo.Bool(false),
                    TaxAmount: gr4vygo.Int64(1200),
                },
            },
            PassengerNameRecord: gr4vygo.String("JOHN L"),
            Passengers: []components.AirlinePassenger{
                components.AirlinePassenger{
                    AgeGroup: components.AgeGroupAdult.ToPointer(),
                    DateOfBirth: types.MustNewDateFromString("2013-07-16"),
                    EmailAddress: gr4vygo.String("john@example.com"),
                    FirstName: gr4vygo.String("John"),
                    FrequentFlyerNumber: gr4vygo.String("15885566"),
                    LastName: gr4vygo.String("Luhn"),
                    PassportNumber: gr4vygo.String("11117700225"),
                    PhoneNumber: gr4vygo.String("+1234567890"),
                    TicketNumber: gr4vygo.String("BA1236699999"),
                    Title: gr4vygo.String("Mr."),
                    CountryCode: gr4vygo.String("US"),
                },
                components.AirlinePassenger{
                    AgeGroup: components.AgeGroupAdult.ToPointer(),
                    DateOfBirth: types.MustNewDateFromString("2013-07-16"),
                    EmailAddress: gr4vygo.String("john@example.com"),
                    FirstName: gr4vygo.String("John"),
                    FrequentFlyerNumber: gr4vygo.String("15885566"),
                    LastName: gr4vygo.String("Luhn"),
                    PassportNumber: gr4vygo.String("11117700225"),
                    PhoneNumber: gr4vygo.String("+1234567890"),
                    TicketNumber: gr4vygo.String("BA1236699999"),
                    Title: gr4vygo.String("Mr."),
                    CountryCode: gr4vygo.String("US"),
                },
            },
            ReservationSystem: gr4vygo.String("Amadeus"),
            RestrictedTicket: gr4vygo.Bool(false),
            TicketDeliveryMethod: components.TicketDeliveryMethodElectronic.ToPointer(),
            TicketNumber: gr4vygo.String("123-1234-151555"),
            TravelAgencyCode: gr4vygo.String("12345"),
            TravelAgencyInvoiceNumber: gr4vygo.String("EG15555155"),
            TravelAgencyName: gr4vygo.String("ACME Agency"),
            TravelAgencyPlanName: gr4vygo.String("B733"),
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
        gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
    )

    res, err := s.CheckoutSessions.Update(ctx, "4137b1cf-39ac-42a8-bad6-1c680d5dab6b", components.CheckoutSessionCreate{
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
                SellerCountry: gr4vygo.String("GB"),
            },
        },
        Metadata: map[string]string{
            "cohort": "cohort-a",
            "order_id": "order-12345",
        },
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
                    Kind: components.TaxIDKindMySst,
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
        Airline: &components.Airline{
            BookingCode: gr4vygo.String("X36Q9C"),
            IsCardholderTraveling: gr4vygo.Bool(true),
            IssuedAddress: gr4vygo.String("123 Broadway, New York"),
            IssuedAt: types.MustNewTimeFromString("2013-07-16T19:23:00.000+00:00"),
            IssuingCarrierCode: gr4vygo.String("649"),
            IssuingCarrierName: gr4vygo.String("Air Transat A.T. Inc"),
            IssuingIataDesignator: gr4vygo.String("TS"),
            IssuingIcaoCode: gr4vygo.String("TSC"),
            Legs: []components.AirlineLeg{
                components.AirlineLeg{
                    ArrivalAirport: gr4vygo.String("LAX"),
                    ArrivalAt: types.MustNewTimeFromString("2013-07-16T19:23:00.000+00:00"),
                    ArrivalCity: gr4vygo.String("Los Angeles"),
                    ArrivalCountry: gr4vygo.String("US"),
                    CarrierCode: gr4vygo.String("649"),
                    CarrierName: gr4vygo.String("Air Transat A.T. Inc"),
                    IataDesignator: gr4vygo.String("TS"),
                    IcaoCode: gr4vygo.String("TSC"),
                    CouponNumber: gr4vygo.String("15885566"),
                    DepartureAirport: gr4vygo.String("LHR"),
                    DepartureAt: types.MustNewTimeFromString("2013-07-16T19:23:00.000+00:00"),
                    DepartureCity: gr4vygo.String("London"),
                    DepartureCountry: gr4vygo.String("GB"),
                    DepartureTaxAmount: gr4vygo.Int64(1200),
                    FareAmount: gr4vygo.Int64(129900),
                    FareBasisCode: gr4vygo.String("FY"),
                    FeeAmount: gr4vygo.Int64(1200),
                    FlightClass: gr4vygo.String("E"),
                    FlightNumber: gr4vygo.String("101"),
                    RouteType: components.RouteTypeRoundTrip.ToPointer(),
                    SeatClass: gr4vygo.String("F"),
                    StopOver: gr4vygo.Bool(false),
                    TaxAmount: gr4vygo.Int64(1200),
                },
                components.AirlineLeg{
                    ArrivalAirport: gr4vygo.String("LAX"),
                    ArrivalAt: types.MustNewTimeFromString("2013-07-16T19:23:00.000+00:00"),
                    ArrivalCity: gr4vygo.String("Los Angeles"),
                    ArrivalCountry: gr4vygo.String("US"),
                    CarrierCode: gr4vygo.String("649"),
                    CarrierName: gr4vygo.String("Air Transat A.T. Inc"),
                    IataDesignator: gr4vygo.String("TS"),
                    IcaoCode: gr4vygo.String("TSC"),
                    CouponNumber: gr4vygo.String("15885566"),
                    DepartureAirport: gr4vygo.String("LHR"),
                    DepartureAt: types.MustNewTimeFromString("2013-07-16T19:23:00.000+00:00"),
                    DepartureCity: gr4vygo.String("London"),
                    DepartureCountry: gr4vygo.String("GB"),
                    DepartureTaxAmount: gr4vygo.Int64(1200),
                    FareAmount: gr4vygo.Int64(129900),
                    FareBasisCode: gr4vygo.String("FY"),
                    FeeAmount: gr4vygo.Int64(1200),
                    FlightClass: gr4vygo.String("E"),
                    FlightNumber: gr4vygo.String("101"),
                    RouteType: components.RouteTypeRoundTrip.ToPointer(),
                    SeatClass: gr4vygo.String("F"),
                    StopOver: gr4vygo.Bool(false),
                    TaxAmount: gr4vygo.Int64(1200),
                },
            },
            PassengerNameRecord: gr4vygo.String("JOHN L"),
            Passengers: []components.AirlinePassenger{
                components.AirlinePassenger{
                    AgeGroup: components.AgeGroupAdult.ToPointer(),
                    DateOfBirth: types.MustNewDateFromString("2013-07-16"),
                    EmailAddress: gr4vygo.String("john@example.com"),
                    FirstName: gr4vygo.String("John"),
                    FrequentFlyerNumber: gr4vygo.String("15885566"),
                    LastName: gr4vygo.String("Luhn"),
                    PassportNumber: gr4vygo.String("11117700225"),
                    PhoneNumber: gr4vygo.String("+1234567890"),
                    TicketNumber: gr4vygo.String("BA1236699999"),
                    Title: gr4vygo.String("Mr."),
                    CountryCode: gr4vygo.String("US"),
                },
                components.AirlinePassenger{
                    AgeGroup: components.AgeGroupAdult.ToPointer(),
                    DateOfBirth: types.MustNewDateFromString("2013-07-16"),
                    EmailAddress: gr4vygo.String("john@example.com"),
                    FirstName: gr4vygo.String("John"),
                    FrequentFlyerNumber: gr4vygo.String("15885566"),
                    LastName: gr4vygo.String("Luhn"),
                    PassportNumber: gr4vygo.String("11117700225"),
                    PhoneNumber: gr4vygo.String("+1234567890"),
                    TicketNumber: gr4vygo.String("BA1236699999"),
                    Title: gr4vygo.String("Mr."),
                    CountryCode: gr4vygo.String("US"),
                },
                components.AirlinePassenger{
                    AgeGroup: components.AgeGroupAdult.ToPointer(),
                    DateOfBirth: types.MustNewDateFromString("2013-07-16"),
                    EmailAddress: gr4vygo.String("john@example.com"),
                    FirstName: gr4vygo.String("John"),
                    FrequentFlyerNumber: gr4vygo.String("15885566"),
                    LastName: gr4vygo.String("Luhn"),
                    PassportNumber: gr4vygo.String("11117700225"),
                    PhoneNumber: gr4vygo.String("+1234567890"),
                    TicketNumber: gr4vygo.String("BA1236699999"),
                    Title: gr4vygo.String("Mr."),
                    CountryCode: gr4vygo.String("US"),
                },
            },
            ReservationSystem: gr4vygo.String("Amadeus"),
            RestrictedTicket: gr4vygo.Bool(false),
            TicketDeliveryMethod: components.TicketDeliveryMethodElectronic.ToPointer(),
            TicketNumber: gr4vygo.String("123-1234-151555"),
            TravelAgencyCode: gr4vygo.String("12345"),
            TravelAgencyInvoiceNumber: gr4vygo.String("EG15555155"),
            TravelAgencyName: gr4vygo.String("ACME Agency"),
            TravelAgencyPlanName: gr4vygo.String("B733"),
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

    res, err := s.CheckoutSessions.Get(ctx, "4137b1cf-39ac-42a8-bad6-1c680d5dab6b", nil)
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

Deleta a checkout session and all of its (PCI) data.

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

    err := s.CheckoutSessions.Delete(ctx, "4137b1cf-39ac-42a8-bad6-1c680d5dab6b", nil)
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