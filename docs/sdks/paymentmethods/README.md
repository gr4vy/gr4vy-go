# PaymentMethods
(*PaymentMethods*)

## Overview

### Available Operations

* [List](#list) - List all payment methods
* [Create](#create) - Create payment method
* [Get](#get) - Get payment method
* [Delete](#delete) - Delete payment method

## List

List all stored payment method.

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

    res, err := s.PaymentMethods.List(ctx, operations.ListPaymentMethodsRequest{
        Cursor: gr4vygo.String("ZXhhbXBsZTE"),
        BuyerID: gr4vygo.String("fe26475d-ec3e-4884-9553-f7356683f7f9"),
        BuyerExternalIdentifier: gr4vygo.String("buyer-12345"),
        ExternalIdentifier: gr4vygo.String("payment-method-12345"),
        XGr4vyMerchantAccountID: gr4vygo.String("default"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CollectionPaymentMethod != nil {
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListPaymentMethodsRequest](../../models/operations/listpaymentmethodsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListPaymentMethodsResponse](../../models/operations/listpaymentmethodsresponse.md), error**

### Errors

| Error Type                              | Status Code                             | Content Type                            |
| --------------------------------------- | --------------------------------------- | --------------------------------------- |
| apierrors.Error400                      | 400                                     | application/json                        |
| apierrors.Error401                      | 401                                     | application/json                        |
| apierrors.Response403ListPaymentMethods | 403                                     | application/json                        |
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

## Create

Store a new payment method.

### Example Usage

```go
package main

import(
	"context"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"os"
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

    res, err := s.PaymentMethods.Create(ctx, operations.CreateCreatePaymentMethodBodyCheckoutSessionPaymentMethodCreate(
        components.CheckoutSessionPaymentMethodCreate{
            ID: "4137b1cf-39ac-42a8-bad6-1c680d5dab6b",
            ExternalIdentifier: gr4vygo.String("card-12345"),
            BuyerID: gr4vygo.String("fe26475d-ec3e-4884-9553-f7356683f7f9"),
            BuyerExternalIdentifier: gr4vygo.String("buyer-12345"),
        },
    ), nil, gr4vygo.String("default"))
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentMethod != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `requestBody`                                                                            | [operations.CreatePaymentMethodBody](../../models/operations/createpaymentmethodbody.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |                                                                                          |
| `timeoutInSeconds`                                                                       | **float64*                                                                               | :heavy_minus_sign:                                                                       | N/A                                                                                      |                                                                                          |
| `xGr4vyMerchantAccountID`                                                                | **string*                                                                                | :heavy_minus_sign:                                                                       | The ID of the merchant account to use for this request.                                  | default                                                                                  |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |                                                                                          |

### Response

**[*operations.CreatePaymentMethodResponse](../../models/operations/createpaymentmethodresponse.md), error**

### Errors

| Error Type                               | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| apierrors.Error400                       | 400                                      | application/json                         |
| apierrors.Error401                       | 401                                      | application/json                         |
| apierrors.Response403CreatePaymentMethod | 403                                      | application/json                         |
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

## Get

Retrieve a payment method.

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

    res, err := s.PaymentMethods.Get(ctx, "ef9496d8-53a5-4aad-8ca2-00eb68334389", gr4vygo.String("default"))
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentMethod != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `paymentMethodID`                                        | *string*                                                 | :heavy_check_mark:                                       | The ID of the payment method                             | ef9496d8-53a5-4aad-8ca2-00eb68334389                     |
| `xGr4vyMerchantAccountID`                                | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  | default                                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPaymentMethodResponse](../../models/operations/getpaymentmethodresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| apierrors.Error400                    | 400                                   | application/json                      |
| apierrors.Error401                    | 401                                   | application/json                      |
| apierrors.Response403GetPaymentMethod | 403                                   | application/json                      |
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

## Delete

Delete a payment method.

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

    res, err := s.PaymentMethods.Delete(ctx, "ef9496d8-53a5-4aad-8ca2-00eb68334389", gr4vygo.String("default"))
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
| `paymentMethodID`                                        | *string*                                                 | :heavy_check_mark:                                       | The ID of the payment method                             | ef9496d8-53a5-4aad-8ca2-00eb68334389                     |
| `xGr4vyMerchantAccountID`                                | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  | default                                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeletePaymentMethodResponse](../../models/operations/deletepaymentmethodresponse.md), error**

### Errors

| Error Type                               | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| apierrors.Error400                       | 400                                      | application/json                         |
| apierrors.Error401                       | 401                                      | application/json                         |
| apierrors.Response403DeletePaymentMethod | 403                                      | application/json                         |
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