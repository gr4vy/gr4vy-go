# Cryptogram
(*PaymentMethods.NetworkTokens.Cryptogram*)

## Overview

### Available Operations

* [Create](#create) - Provision network token cryptogram

## Create

Provision a cryptogram for a network token.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_payment_method_network_token_cryptogram" method="post" path="/payment-methods/{payment_method_id}/network-tokens/{network_token_id}/cryptogram" -->
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
        gr4vygo.WithMerchantAccountID("default"),
        gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
    )

    res, err := s.PaymentMethods.NetworkTokens.Cryptogram.Create(ctx, "ef9496d8-53a5-4aad-8ca2-00eb68334389", "f8dd5cfc-7834-4847-95dc-f75a360e2298", components.CryptogramCreate{
        MerchantInitiated: false,
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                | Example                                                                    |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |                                                                            |
| `paymentMethodID`                                                          | *string*                                                                   | :heavy_check_mark:                                                         | The ID of the payment method                                               | ef9496d8-53a5-4aad-8ca2-00eb68334389                                       |
| `networkTokenID`                                                           | *string*                                                                   | :heavy_check_mark:                                                         | The ID of the network token                                                | f8dd5cfc-7834-4847-95dc-f75a360e2298                                       |
| `cryptogramCreate`                                                         | [components.CryptogramCreate](../../models/components/cryptogramcreate.md) | :heavy_check_mark:                                                         | N/A                                                                        |                                                                            |
| `merchantAccountID`                                                        | **string*                                                                  | :heavy_minus_sign:                                                         | The ID of the merchant account to use for this request.                    |                                                                            |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |                                                                            |

### Response

**[*components.Cryptogram](../../models/components/cryptogram.md), error**

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