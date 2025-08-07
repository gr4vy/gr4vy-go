# Gr4vy SDK

## Overview

Gr4vy: The Gr4vy API.

### Available Operations

* [BrowsePaymentMethodDefinitionsGet](#browsepaymentmethoddefinitionsget) - Browse

## BrowsePaymentMethodDefinitionsGet

Browse

### Example Usage

<!-- UsageSnippet language="go" operationID="browse_payment_method_definitions_get" method="get" path="/payment-method-definitions" -->
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
        gr4vygo.WithMerchantAccountID("<id>"),
        gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
    )

    res, err := s.BrowsePaymentMethodDefinitionsGet(ctx)
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
| `merchantAccountID`                                      | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[any](../../.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |