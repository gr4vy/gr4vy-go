# Transactions
(*Transactions*)

## Overview

### Available Operations

* [List](#list) - List transactions
* [Create](#create) - Create transaction
* [Get](#get) - Get transaction
* [Capture](#capture) - Capture transaction
* [Void](#void) - Void transaction
* [Summary](#summary) - Get transaction summary
* [Sync](#sync) - Sync transaction

## List

List all transactions for a specific merchant account sorted by most recently created.

### Example Usage

```go
package main

import(
	"context"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"os"
	"github.com/gr4vy/gr4vy-go/types"
	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := gr4vygo.New(
        "https://api.example.com",
        gr4vygo.WithSecurity(os.Getenv("GR4VY_O_AUTH2_PASSWORD_BEARER")),
    )

    res, err := s.Transactions.List(ctx, operations.ListTransactionsRequest{
        Cursor: gr4vygo.String("ZXhhbXBsZTE"),
        CreatedAtLte: types.MustNewTimeFromString("2022-01-01T12:00:00+08:00"),
        CreatedAtGte: types.MustNewTimeFromString("2022-01-01T12:00:00+08:00"),
        UpdatedAtLte: types.MustNewTimeFromString("2022-01-01T12:00:00+08:00"),
        UpdatedAtGte: types.MustNewTimeFromString("2022-01-01T12:00:00+08:00"),
        Search: gr4vygo.String("transaction-12345"),
        BuyerExternalIdentifier: gr4vygo.String("buyer-12345"),
        BuyerID: gr4vygo.String("fe26475d-ec3e-4884-9553-f7356683f7f9"),
        BuyerEmailAddress: gr4vygo.String("john@example.com"),
        Status: []components.TransactionStatus{
            components.TransactionStatusAuthorizationSucceeded,
        },
        ID: gr4vygo.String("7099948d-7286-47e4-aad8-b68f7eb44591"),
        PaymentServiceTransactionID: gr4vygo.String("tx-12345"),
        ExternalIdentifier: gr4vygo.String("transaction-12345"),
        Metadata: []string{
            "{\"first_key\":\"first_value\",\"second_key\":\"second_value\"}",
        },
        AmountEq: gr4vygo.Int64(1299),
        AmountLte: gr4vygo.Int64(1299),
        AmountGte: gr4vygo.Int64(1299),
        Currency: []string{
            "EUR",
            "GBP",
            "USD",
        },
        PaymentServiceID: []string{
            "fffd152a-9532-4087-9a4f-de58754210f0",
        },
        PaymentMethodID: gr4vygo.String("ef9496d8-53a5-4aad-8ca2-00eb68334389"),
        PaymentMethodLabel: gr4vygo.String("1234"),
        PaymentMethodFingerprint: gr4vygo.String("a50b85c200ee0795d6fd33a5c66f37a4564f554355c5b46a756aac485dd168a4"),
        Method: []components.Method{
            components.MethodCard,
        },
        ErrorCode: []string{
            "insufficient_funds",
        },
        HasRefunds: gr4vygo.Bool(true),
        PendingReview: gr4vygo.Bool(true),
        CheckoutSessionID: gr4vygo.String("4137b1cf-39ac-42a8-bad6-1c680d5dab6b"),
        ReconciliationID: gr4vygo.String("7jZXl4gBUNl0CnaLEnfXbt"),
        HasGiftCardRedemptions: gr4vygo.Bool(true),
        GiftCardID: gr4vygo.String("356d56e5-fe16-42ae-97ee-8d55d846ae2e"),
        GiftCardLast4: gr4vygo.String("7890"),
        HasSettlements: gr4vygo.Bool(true),
        PaymentMethodBin: gr4vygo.String("411111"),
        PaymentSource: []components.TransactionPaymentSource{
            components.TransactionPaymentSourceRecurring,
        },
        IsSubsequentPayment: gr4vygo.Bool(true),
        MerchantInitiated: gr4vygo.Bool(true),
        XGr4vyMerchantAccountID: gr4vygo.String("default"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CollectionTransactionSummary != nil {
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

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| apierrors.Error400                    | 400                                   | application/json                      |
| apierrors.Error401                    | 401                                   | application/json                      |
| apierrors.Response403ListTransactions | 403                                   | application/json                      |
| apierrors.Error404                    | 404                                   | application/json                      |
| apierrors.Error405                    | 405                                   | application/json                      |
| apierrors.Error409                    | 409                                   | application/json                      |
| apierrors.HTTPValidationError         | 422                                   | application/json                      |
| apierrors.Error425                    | 425                                   | application/json                      |
| apierrors.Error429                    | 429                                   | application/json                      |
| apierrors.Error500                    | 500                                   | application/json                      |
| apierrors.Error502                    | 502                                   | application/json                      |
| apierrors.Error504                    | 504                                   | application/json                      |
| apierrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## Create

Create a transaction.

### Example Usage

```go
package main

import(
	"context"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"os"
	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := gr4vygo.New(
        "https://api.example.com",
        gr4vygo.WithSecurity(os.Getenv("GR4VY_O_AUTH2_PASSWORD_BEARER")),
    )

    res, err := s.Transactions.Create(ctx, components.TransactionCreate{
        Amount: 1299,
        Currency: "EUR",
        Country: gr4vygo.String("US"),
        PaymentMethod: gr4vygo.Pointer(components.CreateTransactionCreatePaymentMethodCardWithURLPaymentMethodCreate(
            components.CardWithURLPaymentMethodCreate{
                ExpirationDate: "12/30",
                Number: "4111111111111111",
                BuyerExternalIdentifier: gr4vygo.String("buyer-12345"),
                BuyerID: gr4vygo.String("fe26475d-ec3e-4884-9553-f7356683f7f9"),
                ExternalIdentifier: gr4vygo.String("payment-method-12345"),
                CardType: gr4vygo.String("credit"),
                SecurityCode: gr4vygo.String("123"),
            },
        )),
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
                    Kind: components.TaxIDKindBrCnpj,
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
        BuyerID: gr4vygo.String("fe26475d-ec3e-4884-9553-f7356683f7f9"),
        BuyerExternalIdentifier: gr4vygo.String("buyer-12345"),
        GiftCards: []components.GiftCardUnion{
            components.CreateGiftCardUnionGiftCardTokenTransactionCreate(
                components.GiftCardTokenTransactionCreate{
                    ID: "356d56e5-fe16-42ae-97ee-8d55d846ae2e",
                    Amount: 1299,
                },
            ),
            components.CreateGiftCardUnionGiftCardTokenTransactionCreate(
                components.GiftCardTokenTransactionCreate{
                    ID: "356d56e5-fe16-42ae-97ee-8d55d846ae2e",
                    Amount: 1299,
                },
            ),
            components.CreateGiftCardUnionGiftCardTransactionCreate(
                components.GiftCardTransactionCreate{
                    Number: "4123455541234561234",
                    Pin: "1234",
                    Amount: 1299,
                },
            ),
        },
        ExternalIdentifier: gr4vygo.String("transaction-12345"),
        ThreeDSecureData: gr4vygo.Pointer(components.CreateThreeDSecureDataThreeDSecureDataV1(
            components.ThreeDSecureDataV1{
                Cavv: "3q2+78r+ur7erb7vyv66vv8=",
                Eci: "05",
                Version: "2.1.0",
                DirectoryResponse: "C",
                Scheme: components.CardSchemeVisa.ToPointer(),
                AuthenticationResponse: "Y",
                CavvAlgorithm: "A",
                Xid: "12345",
            },
        )),
        Metadata: map[string]string{
            "cohort": "cohort-12345",
            "order": "order-12345",
        },
        Airline: &components.Airline{
            BookingCode: gr4vygo.String("X36Q9C"),
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
        StatementDescriptor: &components.StatementDescriptor{
            Name: gr4vygo.String("ACME"),
            Description: gr4vygo.String("ACME San Jose Electronics"),
            City: gr4vygo.String("San Jose"),
            Country: gr4vygo.String("US"),
            PhoneNumber: gr4vygo.String("+1234567890"),
            URL: gr4vygo.String("www.example.com"),
        },
        PreviousSchemeTransactionID: gr4vygo.String("123456789012345"),
        BrowserInfo: &components.BrowserInfo{
            JavascriptEnabled: false,
            JavaEnabled: false,
            Language: "<value>",
            ColorDepth: 242405,
            ScreenHeight: 557747,
            ScreenWidth: 288219,
            TimeZoneOffset: 538274,
            UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
            UserDevice: components.UserDeviceDesktop,
            AcceptHeader: gr4vygo.String("*/*"),
        },
        ShippingDetailsID: gr4vygo.String("bf8c36ad-02d9-4904-b0f9-a230b149e341"),
        AntiFraudFingerprint: gr4vygo.String("yGeBAFYgFmM="),
        PaymentServiceID: gr4vygo.String("fffd152a-9532-4087-9a4f-de58754210f0"),
        Recipient: &components.Recipient{
            FirstName: "",
            LastName: "",
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
            AccountNumber: gr4vygo.String("act12345"),
            DateOfBirth: types.MustNewDateFromString("1995-12-23"),
        },
    }, nil, gr4vygo.String("default"), gr4vygo.String("request-12345"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Transaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                             | Type                                                                                                                                                                                                  | Required                                                                                                                                                                                              | Description                                                                                                                                                                                           | Example                                                                                                                                                                                               |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                    | The context to use for the request.                                                                                                                                                                   |                                                                                                                                                                                                       |
| `transactionCreate`                                                                                                                                                                                   | [components.TransactionCreate](../../models/components/transactioncreate.md)                                                                                                                          | :heavy_check_mark:                                                                                                                                                                                    | N/A                                                                                                                                                                                                   |                                                                                                                                                                                                       |
| `timeoutInSeconds`                                                                                                                                                                                    | **float64*                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                    | N/A                                                                                                                                                                                                   |                                                                                                                                                                                                       |
| `xGr4vyMerchantAccountID`                                                                                                                                                                             | **string*                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                    | The ID of the merchant account to use for this request.                                                                                                                                               | default                                                                                                                                                                                               |
| `idempotencyKey`                                                                                                                                                                                      | **string*                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                    | A unique key that identifies this request. Providing this header will make this an idempotent request. We recommend using V4 UUIDs, or another random string with enough entropy to avoid collisions. | request-12345                                                                                                                                                                                         |
| `opts`                                                                                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                    | The options for this request.                                                                                                                                                                         |                                                                                                                                                                                                       |

### Response

**[*operations.CreateTransactionResponse](../../models/operations/createtransactionresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.Error400                     | 400                                    | application/json                       |
| apierrors.Error401                     | 401                                    | application/json                       |
| apierrors.Response403CreateTransaction | 403                                    | application/json                       |
| apierrors.Error404                     | 404                                    | application/json                       |
| apierrors.Error405                     | 405                                    | application/json                       |
| apierrors.Error409                     | 409                                    | application/json                       |
| apierrors.HTTPValidationError          | 422                                    | application/json                       |
| apierrors.Error425                     | 425                                    | application/json                       |
| apierrors.Error429                     | 429                                    | application/json                       |
| apierrors.Error500                     | 500                                    | application/json                       |
| apierrors.Error502                     | 502                                    | application/json                       |
| apierrors.Error504                     | 504                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Get

Fetch a single transaction by its ID.

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

    res, err := s.Transactions.Get(ctx, "bde12786-dce8-4654-b031-196961d1ddcc", gr4vygo.String("default"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Transaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `transactionID`                                          | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `xGr4vyMerchantAccountID`                                | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  | default                                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetTransactionResponse](../../models/operations/gettransactionresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| apierrors.Error400                  | 400                                 | application/json                    |
| apierrors.Error401                  | 401                                 | application/json                    |
| apierrors.Response403GetTransaction | 403                                 | application/json                    |
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

## Capture

Capture a previously authorized transaction.

### Example Usage

```go
package main

import(
	"context"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"os"
	"github.com/gr4vy/gr4vy-go/types"
	"github.com/gr4vy/gr4vy-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := gr4vygo.New(
        "https://api.example.com",
        gr4vygo.WithSecurity(os.Getenv("GR4VY_O_AUTH2_PASSWORD_BEARER")),
    )

    res, err := s.Transactions.Capture(ctx, "9a049029-b958-45dd-86d7-d7f9fc92c411", components.TransactionCapture{
        Amount: gr4vygo.Int64(1299),
        Airline: &components.Airline{
            BookingCode: gr4vygo.String("X36Q9C"),
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
    }, nil, gr4vygo.String("default"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Transaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `transactionID`                                                                | *string*                                                                       | :heavy_check_mark:                                                             | N/A                                                                            |                                                                                |
| `transactionCapture`                                                           | [components.TransactionCapture](../../models/components/transactioncapture.md) | :heavy_check_mark:                                                             | N/A                                                                            |                                                                                |
| `timeoutInSeconds`                                                             | **float64*                                                                     | :heavy_minus_sign:                                                             | N/A                                                                            |                                                                                |
| `xGr4vyMerchantAccountID`                                                      | **string*                                                                      | :heavy_minus_sign:                                                             | The ID of the merchant account to use for this request.                        | default                                                                        |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.CaptureTransactionResponse](../../models/operations/capturetransactionresponse.md), error**

### Errors

| Error Type                              | Status Code                             | Content Type                            |
| --------------------------------------- | --------------------------------------- | --------------------------------------- |
| apierrors.Error400                      | 400                                     | application/json                        |
| apierrors.Error401                      | 401                                     | application/json                        |
| apierrors.Response403CaptureTransaction | 403                                     | application/json                        |
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

## Void

Void a previously authorized transaction.

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

    res, err := s.Transactions.Void(ctx, "7dbc44c9-1ea3-4853-87be-9923dd281b0d", nil, gr4vygo.String("default"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Transaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `transactionID`                                          | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `timeoutInSeconds`                                       | **float64*                                               | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `xGr4vyMerchantAccountID`                                | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  | default                                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.VoidTransactionResponse](../../models/operations/voidtransactionresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.Error400                   | 400                                  | application/json                     |
| apierrors.Error401                   | 401                                  | application/json                     |
| apierrors.Response403VoidTransaction | 403                                  | application/json                     |
| apierrors.Error404                   | 404                                  | application/json                     |
| apierrors.Error405                   | 405                                  | application/json                     |
| apierrors.Error409                   | 409                                  | application/json                     |
| apierrors.HTTPValidationError        | 422                                  | application/json                     |
| apierrors.Error425                   | 425                                  | application/json                     |
| apierrors.Error429                   | 429                                  | application/json                     |
| apierrors.Error500                   | 500                                  | application/json                     |
| apierrors.Error502                   | 502                                  | application/json                     |
| apierrors.Error504                   | 504                                  | application/json                     |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

## Summary

Fetch a summary for a transaction.

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

    res, err := s.Transactions.Summary(ctx, "7099948d-7286-47e4-aad8-b68f7eb44591", gr4vygo.String("default"))
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionStatusSummary != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `transactionID`                                          | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      | 7099948d-7286-47e4-aad8-b68f7eb44591                     |
| `xGr4vyMerchantAccountID`                                | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  | default                                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetTransactionSummaryResponse](../../models/operations/gettransactionsummaryresponse.md), error**

### Errors

| Error Type                                 | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| apierrors.Error400                         | 400                                        | application/json                           |
| apierrors.Error401                         | 401                                        | application/json                           |
| apierrors.Response403GetTransactionSummary | 403                                        | application/json                           |
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

## Sync

Fetch the latest status for a transaction.

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

    res, err := s.Transactions.Sync(ctx, "2ee546e0-3b11-478e-afec-fdb362611e22", nil, gr4vygo.String("default"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Transaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `transactionID`                                          | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `timeoutInSeconds`                                       | **float64*                                               | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `xGr4vyMerchantAccountID`                                | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  | default                                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.SyncTransactionResponse](../../models/operations/synctransactionresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.Error400                   | 400                                  | application/json                     |
| apierrors.Error401                   | 401                                  | application/json                     |
| apierrors.Response403SyncTransaction | 403                                  | application/json                     |
| apierrors.Error404                   | 404                                  | application/json                     |
| apierrors.Error405                   | 405                                  | application/json                     |
| apierrors.Error409                   | 409                                  | application/json                     |
| apierrors.HTTPValidationError        | 422                                  | application/json                     |
| apierrors.Error425                   | 425                                  | application/json                     |
| apierrors.Error429                   | 429                                  | application/json                     |
| apierrors.Error500                   | 500                                  | application/json                     |
| apierrors.Error502                   | 502                                  | application/json                     |
| apierrors.Error504                   | 504                                  | application/json                     |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |