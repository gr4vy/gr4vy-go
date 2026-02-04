# MerchantAccounts.ThreeDsConfiguration

## Overview

### Available Operations

* [Create](#create) - Create 3DS configuration for merchant
* [List](#list) - List 3DS configurations for merchant
* [Update](#update) - Edit 3DS configuration
* [Delete](#delete) - Delete 3DS configuration for a merchant

## Create

Create a new 3DS configuration for a merchant account.

### Example Usage

<!-- UsageSnippet language="go" operationID="create_three_ds_configuration" method="post" path="/merchant-accounts/{merchant_account_id}/three-ds-configurations" -->
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

    res, err := s.MerchantAccounts.ThreeDsConfiguration.Create(ctx, "merchant-12345", components.MerchantAccountThreeDSConfigurationCreate{
        MerchantAcquirerBin: "516327",
        MerchantAcquirerID: "123456789012345",
        MerchantName: "Acme Inc.",
        MerchantCountryCode: "840",
        MerchantCategoryCode: "1234",
        MerchantURL: "https://example.com",
        Scheme: components.CardSchemeVisa,
        Metadata: map[string]string{
            "key": "<value>",
            "key1": "<value>",
            "key2": "<value>",
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

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  | Example                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |                                                                                                                              |
| `merchantAccountID`                                                                                                          | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The ID of the merchant account.                                                                                              | merchant-12345                                                                                                               |
| `merchantAccountThreeDSConfigurationCreate`                                                                                  | [components.MerchantAccountThreeDSConfigurationCreate](../../models/components/merchantaccountthreedsconfigurationcreate.md) | :heavy_check_mark:                                                                                                           | N/A                                                                                                                          |                                                                                                                              |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |                                                                                                                              |

### Response

**[*components.MerchantAccountThreeDSConfiguration](../../models/components/merchantaccountthreedsconfiguration.md), error**

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

## List

List all 3DS configurations for a merchant account.

### Example Usage

<!-- UsageSnippet language="go" operationID="list_three_ds_configurations" method="get" path="/merchant-accounts/{merchant_account_id}/three-ds-configurations" -->
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

    res, err := s.MerchantAccounts.ThreeDsConfiguration.List(ctx, "merchant-12345", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ctx`                                                               | [context.Context](https://pkg.go.dev/context#Context)               | :heavy_check_mark:                                                  | The context to use for the request.                                 |                                                                     |
| `merchantAccountID`                                                 | *string*                                                            | :heavy_check_mark:                                                  | The ID of the merchant account.                                     | merchant-12345                                                      |
| `currency`                                                          | **string*                                                           | :heavy_minus_sign:                                                  | ISO 4217 currency code (3 characters) to filter 3DS configurations. | USD                                                                 |
| `opts`                                                              | [][operations.Option](../../models/operations/option.md)            | :heavy_minus_sign:                                                  | The options for this request.                                       |                                                                     |

### Response

**[*components.MerchantAccountThreeDSConfigurations](../../models/components/merchantaccountthreedsconfigurations.md), error**

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

Update the 3DS configuration for a merchant account.

### Example Usage

<!-- UsageSnippet language="go" operationID="edit_three_ds_configuration" method="put" path="/merchant-accounts/{merchant_account_id}/three-ds-configurations/{three_ds_configuration_id}" -->
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

    res, err := s.MerchantAccounts.ThreeDsConfiguration.Update(ctx, "merchant-12345", "1808f5e6-b49c-4db9-94fa-22371ea352f5", components.MerchantAccountThreeDSConfigurationUpdate{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  | Example                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |                                                                                                                              |
| `merchantAccountID`                                                                                                          | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The ID of the merchant account.                                                                                              | merchant-12345                                                                                                               |
| `threeDsConfigurationID`                                                                                                     | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The ID of the 3DS configuration for a merchant account.                                                                      | 1808f5e6-b49c-4db9-94fa-22371ea352f5                                                                                         |
| `merchantAccountThreeDSConfigurationUpdate`                                                                                  | [components.MerchantAccountThreeDSConfigurationUpdate](../../models/components/merchantaccountthreedsconfigurationupdate.md) | :heavy_check_mark:                                                                                                           | N/A                                                                                                                          |                                                                                                                              |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |                                                                                                                              |

### Response

**[*components.MerchantAccountThreeDSConfiguration](../../models/components/merchantaccountthreedsconfiguration.md), error**

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

Delete a 3DS configuration for a merchant account.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete_three_ds_configuration" method="delete" path="/merchant-accounts/{merchant_account_id}/three-ds-configurations/{three_ds_configuration_id}" -->
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

    err := s.MerchantAccounts.ThreeDsConfiguration.Delete(ctx, "merchant-12345", "1808f5e6-b49c-4db9-94fa-22371ea352f5")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `merchantAccountID`                                      | *string*                                                 | :heavy_check_mark:                                       | The ID of the merchant account.                          | merchant-12345                                           |
| `threeDsConfigurationID`                                 | *string*                                                 | :heavy_check_mark:                                       | The ID of the 3DS configuration for a merchant account.  | 1808f5e6-b49c-4db9-94fa-22371ea352f5                     |
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