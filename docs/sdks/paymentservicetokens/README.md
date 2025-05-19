# PaymentServiceTokens
(*PaymentMethods.PaymentServiceTokens*)

## Overview

### Available Operations

* [List](#list) - List payment service tokens
* [Create](#create) - Create payment service token
* [Delete](#delete) - Delete payment service token

## List

List all gateway tokens stored for a payment method.

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

    res, err := s.PaymentMethods.PaymentServiceTokens.List(ctx, "ef9496d8-53a5-4aad-8ca2-00eb68334389", gr4vygo.String("fffd152a-9532-4087-9a4f-de58754210f0"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CollectionNoCursorPaymentServiceToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `paymentMethodID`                                        | *string*                                                 | :heavy_check_mark:                                       | The ID of the payment method                             | ef9496d8-53a5-4aad-8ca2-00eb68334389                     |
| `paymentServiceID`                                       | **string*                                                | :heavy_minus_sign:                                       | The ID of the payment service                            | fffd152a-9532-4087-9a4f-de58754210f0                     |
| `merchantAccountID`                                      | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListPaymentMethodPaymentServiceTokensResponse](../../models/operations/listpaymentmethodpaymentservicetokensresponse.md), error**

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

Create a gateway tokens for a payment method.

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

    res, err := s.PaymentMethods.PaymentServiceTokens.Create(ctx, "ef9496d8-53a5-4aad-8ca2-00eb68334389", components.PaymentServiceTokenCreate{
        SecurityCode: gr4vygo.String("123"),
        PaymentServiceID: "fffd152a-9532-4087-9a4f-de58754210f0",
        RedirectURL: "https://dual-futon.biz",
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentServiceToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `paymentMethodID`                                                                            | *string*                                                                                     | :heavy_check_mark:                                                                           | The ID of the payment method                                                                 | ef9496d8-53a5-4aad-8ca2-00eb68334389                                                         |
| `paymentServiceTokenCreate`                                                                  | [components.PaymentServiceTokenCreate](../../models/components/paymentservicetokencreate.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |
| `timeoutInSeconds`                                                                           | **float64*                                                                                   | :heavy_minus_sign:                                                                           | N/A                                                                                          |                                                                                              |
| `merchantAccountID`                                                                          | **string*                                                                                    | :heavy_minus_sign:                                                                           | The ID of the merchant account to use for this request.                                      |                                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |                                                                                              |

### Response

**[*operations.CreatePaymentMethodPaymentServiceTokenResponse](../../models/operations/createpaymentmethodpaymentservicetokenresponse.md), error**

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

## Delete

Delete a gateway tokens for a payment method.

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

    res, err := s.PaymentMethods.PaymentServiceTokens.Delete(ctx, "ef9496d8-53a5-4aad-8ca2-00eb68334389", "703f2d99-3fd1-44bc-9cbd-a25a2d597886", nil, nil)
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
| `paymentServiceTokenID`                                  | *string*                                                 | :heavy_check_mark:                                       | The ID of the payment service token                      | 703f2d99-3fd1-44bc-9cbd-a25a2d597886                     |
| `timeoutInSeconds`                                       | **float64*                                               | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `merchantAccountID`                                      | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeletePaymentMethodPaymentServiceTokenResponse](../../models/operations/deletepaymentmethodpaymentservicetokenresponse.md), error**

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