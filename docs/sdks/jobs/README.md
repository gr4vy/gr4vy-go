# Jobs
(*AccountUpdater.Jobs*)

## Overview

### Available Operations

* [Create](#create) - Create account updater job

## Create

Schedule one or more stored cards for an account update.

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

    res, err := s.AccountUpdater.Jobs.Create(ctx, components.AccountUpdaterJobCreate{
        PaymentMethodIds: []string{
            "ef9496d8-53a5-4aad-8ca2-00eb68334389",
            "f29e886e-93cc-4714-b4a3-12b7a718e595",
        },
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountUpdaterJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `accountUpdaterJobCreate`                                                                | [components.AccountUpdaterJobCreate](../../models/components/accountupdaterjobcreate.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `timeoutInSeconds`                                                                       | **float64*                                                                               | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `merchantAccountID`                                                                      | **string*                                                                                | :heavy_minus_sign:                                                                       | The ID of the merchant account to use for this request.                                  |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CreateAccountUpdaterJobResponse](../../models/operations/createaccountupdaterjobresponse.md), error**

### Errors

| Error Type                                   | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| apierrors.Error400                           | 400                                          | application/json                             |
| apierrors.Error401                           | 401                                          | application/json                             |
| apierrors.Response403CreateAccountUpdaterJob | 403                                          | application/json                             |
| apierrors.Error404                           | 404                                          | application/json                             |
| apierrors.Error405                           | 405                                          | application/json                             |
| apierrors.Error409                           | 409                                          | application/json                             |
| apierrors.HTTPValidationError                | 422                                          | application/json                             |
| apierrors.Error425                           | 425                                          | application/json                             |
| apierrors.Error429                           | 429                                          | application/json                             |
| apierrors.Error500                           | 500                                          | application/json                             |
| apierrors.Error502                           | 502                                          | application/json                             |
| apierrors.Error504                           | 504                                          | application/json                             |
| apierrors.APIError                           | 4XX, 5XX                                     | \*/\*                                        |