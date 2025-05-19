# PaymentServices
(*PaymentServices*)

## Overview

### Available Operations

* [List](#list) - List payment services
* [Create](#create) - Update a configured payment service
* [Get](#get) - Get payment service
* [Update](#update) - Configure a payment service
* [Delete](#delete) - Delete a configured payment service
* [Verify](#verify) - Verify payment service credentials
* [Session](#session) - Create a session for apayment service definition

## List

List the configured payment services.

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

    res, err := s.PaymentServices.List(ctx, operations.ListPaymentServicesRequest{
        Cursor: gr4vygo.String("ZXhhbXBsZTE"),
        Deleted: gr4vygo.Bool(true),
        XGr4vyMerchantAccountID: gr4vygo.String("default"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CollectionPaymentService != nil {
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListPaymentServicesRequest](../../models/operations/listpaymentservicesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListPaymentServicesResponse](../../models/operations/listpaymentservicesresponse.md), error**

### Errors

| Error Type                               | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| apierrors.Error400                       | 400                                      | application/json                         |
| apierrors.Error401                       | 401                                      | application/json                         |
| apierrors.Response403ListPaymentServices | 403                                      | application/json                         |
| apierrors.Error404                       | 404                                      | application/json                         |
| apierrors.Error405                       | 405                                      | application/json                         |
| apierrors.Error409                       | 409                                      | application/json                         |
| apierrors.HTTPValidationError            | 422                                      | application/json                         |
| apierrors.Error425                       | 425                                      | application/json                         |
| apierrors.Error429                       | 429                                      | application/json                         |
| apierrors.Error500                       | 500                                      | application/json                         |
| apierrors.Error502                       | 502                                      | application/json                         |
| apierrors.Error504                       | 504                                      | application/json                         |
| apierrors.APIError                       | 4XX, 5XX                                 | \*/\*                                    |

## Create

Updates the configuration of a payment service.

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

    res, err := s.PaymentServices.Create(ctx, components.PaymentServiceCreate{
        DisplayName: "Stripe",
        PaymentServiceDefinitionID: "stripe-card",
        Fields: []components.Field{
            components.Field{
                Key: "api_key",
                Value: "key-12345",
            },
            components.Field{
                Key: "api_key",
                Value: "key-12345",
            },
        },
        ReportingFields: []components.Field{
            components.Field{
                Key: "api_key",
                Value: "key-12345",
            },
            components.Field{
                Key: "api_key",
                Value: "key-12345",
            },
            components.Field{
                Key: "api_key",
                Value: "key-12345",
            },
        },
        Position: gr4vygo.Int64(1),
        AcceptedCurrencies: []string{
            "EUR",
            "GBP",
            "USD",
        },
        AcceptedCountries: []string{
            "DE",
            "GB",
            "US",
        },
        Active: gr4vygo.Bool(false),
        MerchantProfile: map[string]*components.MerchantProfileScheme{
            "key": &components.MerchantProfileScheme{
                MerchantAcquirerBin: "516327",
                MerchantURL: "https://example.com",
                MerchantAcquirerID: "123456789012345",
                MerchantName: "Acme Inc.",
                MerchantCountryCode: "USD",
                MerchantCategoryCode: "1234",
            },
            "key1": &components.MerchantProfileScheme{
                MerchantAcquirerBin: "516327",
                MerchantURL: "https://example.com",
                MerchantAcquirerID: "123456789012345",
                MerchantName: "Acme Inc.",
                MerchantCountryCode: "USD",
                MerchantCategoryCode: "1234",
            },
        },
        PaymentMethodTokenizationEnabled: gr4vygo.Bool(true),
        NetworkTokensEnabled: gr4vygo.Bool(true),
        OpenLoop: gr4vygo.Bool(true),
    }, gr4vygo.String("default"))
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentService != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `paymentServiceCreate`                                                             | [components.PaymentServiceCreate](../../models/components/paymentservicecreate.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `xGr4vyMerchantAccountID`                                                          | **string*                                                                          | :heavy_minus_sign:                                                                 | The ID of the merchant account to use for this request.                            | default                                                                            |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.UpdatePaymentServiceResponse](../../models/operations/updatepaymentserviceresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.Error400                        | 400                                       | application/json                          |
| apierrors.Error401                        | 401                                       | application/json                          |
| apierrors.Response403UpdatePaymentService | 403                                       | application/json                          |
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

## Get

Get the details of a configured payment service.

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

    res, err := s.PaymentServices.Get(ctx, "fffd152a-9532-4087-9a4f-de58754210f0", gr4vygo.String("default"))
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentService != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `paymentServiceID`                                       | *string*                                                 | :heavy_check_mark:                                       | the ID of the payment service                            | fffd152a-9532-4087-9a4f-de58754210f0                     |
| `xGr4vyMerchantAccountID`                                | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  | default                                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPaymentServiceResponse](../../models/operations/getpaymentserviceresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.Error400                     | 400                                    | application/json                       |
| apierrors.Error401                     | 401                                    | application/json                       |
| apierrors.Response403GetPaymentService | 403                                    | application/json                       |
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

## Update

Configures a new payment service for use by merchants.

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

    res, err := s.PaymentServices.Update(ctx, "fffd152a-9532-4087-9a4f-de58754210f0", components.PaymentServiceUpdate{
        DisplayName: gr4vygo.String("Stripe"),
        Position: gr4vygo.Int64(1),
        AcceptedCurrencies: []string{
            "EUR",
            "GBP",
            "USD",
        },
        AcceptedCountries: []string{
            "DE",
            "GB",
            "US",
        },
        Active: gr4vygo.Bool(false),
        ThreeDSecureEnabled: gr4vygo.Bool(true),
        MerchantProfile: map[string]*components.MerchantProfileScheme{
            "key": &components.MerchantProfileScheme{
                MerchantAcquirerBin: "516327",
                MerchantURL: "https://example.com",
                MerchantAcquirerID: "123456789012345",
                MerchantName: "Acme Inc.",
                MerchantCountryCode: "USD",
                MerchantCategoryCode: "1234",
            },
        },
        PaymentMethodTokenizationEnabled: gr4vygo.Bool(true),
        NetworkTokensEnabled: gr4vygo.Bool(true),
        OpenLoop: gr4vygo.Bool(true),
    }, gr4vygo.String("default"))
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentService != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `paymentServiceID`                                                                 | *string*                                                                           | :heavy_check_mark:                                                                 | the ID of the payment service                                                      | fffd152a-9532-4087-9a4f-de58754210f0                                               |
| `paymentServiceUpdate`                                                             | [components.PaymentServiceUpdate](../../models/components/paymentserviceupdate.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `xGr4vyMerchantAccountID`                                                          | **string*                                                                          | :heavy_minus_sign:                                                                 | The ID of the merchant account to use for this request.                            | default                                                                            |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.CreatePaymentServiceResponse](../../models/operations/createpaymentserviceresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.Error400                        | 400                                       | application/json                          |
| apierrors.Error401                        | 401                                       | application/json                          |
| apierrors.Response403CreatePaymentService | 403                                       | application/json                          |
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

## Delete

Deletes all the configuration of a payment service.

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

    res, err := s.PaymentServices.Delete(ctx, "fffd152a-9532-4087-9a4f-de58754210f0", nil, gr4vygo.String("default"))
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
| `paymentServiceID`                                       | *string*                                                 | :heavy_check_mark:                                       | the ID of the payment service                            | fffd152a-9532-4087-9a4f-de58754210f0                     |
| `timeoutInSeconds`                                       | **float64*                                               | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `xGr4vyMerchantAccountID`                                | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  | default                                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeletePaymentServiceResponse](../../models/operations/deletepaymentserviceresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.Error400                        | 400                                       | application/json                          |
| apierrors.Error401                        | 401                                       | application/json                          |
| apierrors.Response403DeletePaymentService | 403                                       | application/json                          |
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

## Verify

Verify the credentials of a configured payment service

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

    res, err := s.PaymentServices.Verify(ctx, components.VerifyCredentials{
        PaymentServiceDefinitionID: "stripe-card",
        PaymentServiceID: gr4vygo.String("fffd152a-9532-4087-9a4f-de58754210f0"),
        Fields: []components.Field{},
    }, nil, gr4vygo.String("default"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Any != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |                                                                              |
| `verifyCredentials`                                                          | [components.VerifyCredentials](../../models/components/verifycredentials.md) | :heavy_check_mark:                                                           | N/A                                                                          |                                                                              |
| `timeoutInSeconds`                                                           | **float64*                                                                   | :heavy_minus_sign:                                                           | N/A                                                                          |                                                                              |
| `xGr4vyMerchantAccountID`                                                    | **string*                                                                    | :heavy_minus_sign:                                                           | The ID of the merchant account to use for this request.                      | default                                                                      |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |                                                                              |

### Response

**[*operations.VerifyPaymentServiceCredentialsResponse](../../models/operations/verifypaymentservicecredentialsresponse.md), error**

### Errors

| Error Type                                           | Status Code                                          | Content Type                                         |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| apierrors.Error400                                   | 400                                                  | application/json                                     |
| apierrors.Error401                                   | 401                                                  | application/json                                     |
| apierrors.Response403VerifyPaymentServiceCredentials | 403                                                  | application/json                                     |
| apierrors.Error404                                   | 404                                                  | application/json                                     |
| apierrors.Error405                                   | 405                                                  | application/json                                     |
| apierrors.Error409                                   | 409                                                  | application/json                                     |
| apierrors.HTTPValidationError                        | 422                                                  | application/json                                     |
| apierrors.Error425                                   | 425                                                  | application/json                                     |
| apierrors.Error429                                   | 429                                                  | application/json                                     |
| apierrors.Error500                                   | 500                                                  | application/json                                     |
| apierrors.Error502                                   | 502                                                  | application/json                                     |
| apierrors.Error504                                   | 504                                                  | application/json                                     |
| apierrors.APIError                                   | 4XX, 5XX                                             | \*/\*                                                |

## Session

Creates a session for a payment service that supports sessions.

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

    res, err := s.PaymentServices.Session(ctx, "fffd152a-9532-4087-9a4f-de58754210f0", map[string]any{

    }, gr4vygo.String("default"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateSession != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `paymentServiceID`                                       | *string*                                                 | :heavy_check_mark:                                       | the ID of the payment service                            | fffd152a-9532-4087-9a4f-de58754210f0                     |
| `requestBody`                                            | map[string]*any*                                         | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `xGr4vyMerchantAccountID`                                | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  | default                                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.CreatePaymentServiceSessionResponse](../../models/operations/createpaymentservicesessionresponse.md), error**

### Errors

| Error Type                                       | Status Code                                      | Content Type                                     |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| apierrors.Error400                               | 400                                              | application/json                                 |
| apierrors.Error401                               | 401                                              | application/json                                 |
| apierrors.Response403CreatePaymentServiceSession | 403                                              | application/json                                 |
| apierrors.Error404                               | 404                                              | application/json                                 |
| apierrors.Error405                               | 405                                              | application/json                                 |
| apierrors.Error409                               | 409                                              | application/json                                 |
| apierrors.HTTPValidationError                    | 422                                              | application/json                                 |
| apierrors.Error425                               | 425                                              | application/json                                 |
| apierrors.Error429                               | 429                                              | application/json                                 |
| apierrors.Error500                               | 500                                              | application/json                                 |
| apierrors.Error502                               | 502                                              | application/json                                 |
| apierrors.Error504                               | 504                                              | application/json                                 |
| apierrors.APIError                               | 4XX, 5XX                                         | \*/\*                                            |