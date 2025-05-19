# ShippingDetails
(*Buyers.ShippingDetails*)

## Overview

### Available Operations

* [Create](#create) - Add buyer shipping details
* [List](#list) - List a buyer's shipping details
* [Get](#get) - Get buyer shipping details
* [Update](#update) - Update a buyer's shipping details
* [Delete](#delete) - Delete a buyer's shipping details

## Create

Associate shipping details to a buyer.

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

    res, err := s.Buyers.ShippingDetails.Create(ctx, "fe26475d-ec3e-4884-9553-f7356683f7f9", components.ShippingDetailsCreate{
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
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ShippingDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          | Example                                                                              |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |                                                                                      |
| `buyerID`                                                                            | *string*                                                                             | :heavy_check_mark:                                                                   | The ID of the buyer to add shipping details to.                                      | fe26475d-ec3e-4884-9553-f7356683f7f9                                                 |
| `shippingDetailsCreate`                                                              | [components.ShippingDetailsCreate](../../models/components/shippingdetailscreate.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |                                                                                      |
| `timeoutInSeconds`                                                                   | **float64*                                                                           | :heavy_minus_sign:                                                                   | N/A                                                                                  |                                                                                      |
| `merchantAccountID`                                                                  | **string*                                                                            | :heavy_minus_sign:                                                                   | The ID of the merchant account to use for this request.                              |                                                                                      |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |                                                                                      |

### Response

**[*operations.AddBuyerShippingDetailsResponse](../../models/operations/addbuyershippingdetailsresponse.md), error**

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

List all the shipping details associated to a specific buyer.

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

    res, err := s.Buyers.ShippingDetails.List(ctx, "fe26475d-ec3e-4884-9553-f7356683f7f9", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CollectionNoCursorShippingDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `buyerID`                                                | *string*                                                 | :heavy_check_mark:                                       | The ID of the buyer to retrieve shipping details for.    | fe26475d-ec3e-4884-9553-f7356683f7f9                     |
| `merchantAccountID`                                      | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListBuyerShippingDetailsResponse](../../models/operations/listbuyershippingdetailsresponse.md), error**

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

Get a buyer's shipping details.

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

    res, err := s.Buyers.ShippingDetails.Get(ctx, "fe26475d-ec3e-4884-9553-f7356683f7f9", "bf8c36ad-02d9-4904-b0f9-a230b149e341", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ShippingDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `buyerID`                                                | *string*                                                 | :heavy_check_mark:                                       | The ID of the buyer to retrieve shipping details for.    | fe26475d-ec3e-4884-9553-f7356683f7f9                     |
| `shippingDetailsID`                                      | *string*                                                 | :heavy_check_mark:                                       | The ID of the shipping details to retrieve.              | bf8c36ad-02d9-4904-b0f9-a230b149e341                     |
| `merchantAccountID`                                      | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetBuyerShippingDetailsResponse](../../models/operations/getbuyershippingdetailsresponse.md), error**

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

Update the shipping details associated to a specific buyer.

### Example Usage

```go
package main

import(
	"context"
	"os"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := gr4vygo.New(
        gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
    )

    res, err := s.Buyers.ShippingDetails.Update(ctx, operations.UpdateBuyerShippingDetailsRequest{
        BuyerID: "fe26475d-ec3e-4884-9553-f7356683f7f9",
        ShippingDetailsID: "bf8c36ad-02d9-4904-b0f9-a230b149e341",
        ShippingDetailsUpdate: components.ShippingDetailsUpdate{
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
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ShippingDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.UpdateBuyerShippingDetailsRequest](../../models/operations/updatebuyershippingdetailsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.UpdateBuyerShippingDetailsResponse](../../models/operations/updatebuyershippingdetailsresponse.md), error**

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

Delete the shipping details associated to a specific buyer.

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

    res, err := s.Buyers.ShippingDetails.Delete(ctx, "fe26475d-ec3e-4884-9553-f7356683f7f9", "bf8c36ad-02d9-4904-b0f9-a230b149e341", nil, nil)
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
| `buyerID`                                                | *string*                                                 | :heavy_check_mark:                                       | The ID of the buyer to delete shipping details for.      | fe26475d-ec3e-4884-9553-f7356683f7f9                     |
| `shippingDetailsID`                                      | *string*                                                 | :heavy_check_mark:                                       | The ID of the shipping details to delete.                | bf8c36ad-02d9-4904-b0f9-a230b149e341                     |
| `timeoutInSeconds`                                       | **float64*                                               | :heavy_minus_sign:                                       | N/A                                                      |                                                          |
| `merchantAccountID`                                      | **string*                                                | :heavy_minus_sign:                                       | The ID of the merchant account to use for this request.  |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteBuyerShippingDetailsResponse](../../models/operations/deletebuyershippingdetailsresponse.md), error**

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