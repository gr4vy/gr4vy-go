# Sessions
(*DigitalWallets.Sessions*)

## Overview

### Available Operations

* [GooglePay](#googlepay) - Create a Google Pay session
* [ApplePay](#applepay) - Create a Apple Pay session
* [ClickToPay](#clicktopay) - Create a Click to Pay session

## GooglePay

Create a session for use with Google Pay.

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

    res, err := s.DigitalWallets.Sessions.GooglePay(ctx, components.GooglePaySessionRequest{
        OriginDomain: "example.com",
    }, gr4vygo.String("default"))
    if err != nil {
        log.Fatal(err)
    }
    if res.GooglePaySession != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `googlePaySessionRequest`                                                                | [components.GooglePaySessionRequest](../../models/components/googlepaysessionrequest.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |                                                                                          |
| `xGr4vyMerchantAccountID`                                                                | **string*                                                                                | :heavy_minus_sign:                                                                       | The ID of the merchant account to use for this request.                                  | default                                                                                  |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |                                                                                          |

### Response

**[*operations.CreateGooglePayDigitalWalletSessionResponse](../../models/operations/creategooglepaydigitalwalletsessionresponse.md), error**

### Errors

| Error Type                                               | Status Code                                              | Content Type                                             |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| apierrors.Error400                                       | 400                                                      | application/json                                         |
| apierrors.Error401                                       | 401                                                      | application/json                                         |
| apierrors.Response403CreateGooglePayDigitalWalletSession | 403                                                      | application/json                                         |
| apierrors.Error404                                       | 404                                                      | application/json                                         |
| apierrors.Error405                                       | 405                                                      | application/json                                         |
| apierrors.Error409                                       | 409                                                      | application/json                                         |
| apierrors.HTTPValidationError                            | 422                                                      | application/json                                         |
| apierrors.Error425                                       | 425                                                      | application/json                                         |
| apierrors.Error429                                       | 429                                                      | application/json                                         |
| apierrors.Error500                                       | 500                                                      | application/json                                         |
| apierrors.Error502                                       | 502                                                      | application/json                                         |
| apierrors.Error504                                       | 504                                                      | application/json                                         |
| apierrors.APIError                                       | 4XX, 5XX                                                 | \*/\*                                                    |

## ApplePay

Create a session for use with Apple Pay.

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

    res, err := s.DigitalWallets.Sessions.ApplePay(ctx, components.ApplePaySessionRequest{
        ValidationURL: "https://apple-pay-gateway-cert.apple.com",
        DomainName: "example.com",
    }, gr4vygo.String("default"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ApplePaySession != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `applePaySessionRequest`                                                               | [components.ApplePaySessionRequest](../../models/components/applepaysessionrequest.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `xGr4vyMerchantAccountID`                                                              | **string*                                                                              | :heavy_minus_sign:                                                                     | The ID of the merchant account to use for this request.                                | default                                                                                |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

### Response

**[*operations.CreateApplePayDigitalWalletSessionResponse](../../models/operations/createapplepaydigitalwalletsessionresponse.md), error**

### Errors

| Error Type                                              | Status Code                                             | Content Type                                            |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| apierrors.Error400                                      | 400                                                     | application/json                                        |
| apierrors.Error401                                      | 401                                                     | application/json                                        |
| apierrors.Response403CreateApplePayDigitalWalletSession | 403                                                     | application/json                                        |
| apierrors.Error404                                      | 404                                                     | application/json                                        |
| apierrors.Error405                                      | 405                                                     | application/json                                        |
| apierrors.Error409                                      | 409                                                     | application/json                                        |
| apierrors.HTTPValidationError                           | 422                                                     | application/json                                        |
| apierrors.Error425                                      | 425                                                     | application/json                                        |
| apierrors.Error429                                      | 429                                                     | application/json                                        |
| apierrors.Error500                                      | 500                                                     | application/json                                        |
| apierrors.Error502                                      | 502                                                     | application/json                                        |
| apierrors.Error504                                      | 504                                                     | application/json                                        |
| apierrors.APIError                                      | 4XX, 5XX                                                | \*/\*                                                   |

## ClickToPay

Create a session for use with Click to Pay.

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

    res, err := s.DigitalWallets.Sessions.ClickToPay(ctx, components.ClickToPaySessionRequest{
        CheckoutSessionID: "4137b1cf-39ac-42a8-bad6-1c680d5dab6b",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ClickToPaySession != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.ClickToPaySessionRequest](../../models/components/clicktopaysessionrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateClickToPayDigitalWalletSessionResponse](../../models/operations/createclicktopaydigitalwalletsessionresponse.md), error**

### Errors

| Error Type                                                | Status Code                                               | Content Type                                              |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| apierrors.Error400                                        | 400                                                       | application/json                                          |
| apierrors.Error401                                        | 401                                                       | application/json                                          |
| apierrors.Response403CreateClickToPayDigitalWalletSession | 403                                                       | application/json                                          |
| apierrors.Error404                                        | 404                                                       | application/json                                          |
| apierrors.Error405                                        | 405                                                       | application/json                                          |
| apierrors.Error409                                        | 409                                                       | application/json                                          |
| apierrors.HTTPValidationError                             | 422                                                       | application/json                                          |
| apierrors.Error425                                        | 425                                                       | application/json                                          |
| apierrors.Error429                                        | 429                                                       | application/json                                          |
| apierrors.Error500                                        | 500                                                       | application/json                                          |
| apierrors.Error502                                        | 502                                                       | application/json                                          |
| apierrors.Error504                                        | 504                                                       | application/json                                          |
| apierrors.APIError                                        | 4XX, 5XX                                                  | \*/\*                                                     |