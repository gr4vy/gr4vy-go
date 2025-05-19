# Gr4vy Go SDK (Beta)

Developer-friendly & type-safe Go SDK specifically catered to leverage the **Gr4vy** API.

<div align="left">
	<img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/gr4vy/gr4vy-go?style=for-the-badge">
    <a href="https://www.speakeasy.com/?utm_source=github-com/gr4vy/gr4vy-go&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
</div>


<br /><br />
> [!IMPORTANT]
> This is a Beta release of our latest SDK. Please refer to the [legacy Go SDK](https://github.com/gr4vy/gr4vy-go/tree/legacy) for the latest stable build.

<!-- Start Summary [summary] -->
## Summary

Gr4vy: The Gr4vy API.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/gr4vy/gr4vy-go](#githubcomgr4vygr4vy-go)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Pagination](#pagination)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Special Types](#special-types)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/gr4vy/gr4vy-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/components"
	"log"
	"os"
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
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type | Scheme      | Environment Variable |
| ------------ | ---- | ----------- | -------------------- |
| `BearerAuth` | http | HTTP Bearer | `GR4VY_BEARER_AUTH`  |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/components"
	"log"
	"os"
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
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [AccountUpdater](docs/sdks/accountupdater/README.md)


#### [AccountUpdater.Jobs](docs/sdks/jobs/README.md)

* [Create](docs/sdks/jobs/README.md#create) - Create account updater job

### [AuditLogs](docs/sdks/auditlogs/README.md)

* [List](docs/sdks/auditlogs/README.md#list) - List audit log entries

### [Buyers](docs/sdks/buyers/README.md)

* [List](docs/sdks/buyers/README.md#list) - List all buyers
* [Create](docs/sdks/buyers/README.md#create) - Add a buyer
* [Get](docs/sdks/buyers/README.md#get) - Get a buyer
* [Update](docs/sdks/buyers/README.md#update) - Update a buyer
* [Delete](docs/sdks/buyers/README.md#delete) - Delete a buyer

#### [Buyers.GiftCards](docs/sdks/buyersgiftcards/README.md)

* [List](docs/sdks/buyersgiftcards/README.md#list) - List gift cards for a buyer

#### [Buyers.PaymentMethods](docs/sdks/buyerspaymentmethods/README.md)

* [List](docs/sdks/buyerspaymentmethods/README.md#list) - List payment methods for a buyer

#### [Buyers.ShippingDetails](docs/sdks/shippingdetails/README.md)

* [Create](docs/sdks/shippingdetails/README.md#create) - Add buyer shipping details
* [List](docs/sdks/shippingdetails/README.md#list) - List a buyer's shipping details
* [Get](docs/sdks/shippingdetails/README.md#get) - Get buyer shipping details
* [Update](docs/sdks/shippingdetails/README.md#update) - Update a buyer's shipping details
* [Delete](docs/sdks/shippingdetails/README.md#delete) - Delete a buyer's shipping details

### [CardSchemeDefinitions](docs/sdks/cardschemedefinitions/README.md)

* [List](docs/sdks/cardschemedefinitions/README.md#list) - List card scheme definitions

### [CheckoutSessions](docs/sdks/checkoutsessions/README.md)

* [Create](docs/sdks/checkoutsessions/README.md#create) - Create checkout session
* [Update](docs/sdks/checkoutsessions/README.md#update) - Update checkout session
* [Get](docs/sdks/checkoutsessions/README.md#get) - Get checkout session
* [Delete](docs/sdks/checkoutsessions/README.md#delete) - Delete checkout session

### [DigitalWallets](docs/sdks/digitalwallets/README.md)

* [Create](docs/sdks/digitalwallets/README.md#create) - Register digital wallet
* [List](docs/sdks/digitalwallets/README.md#list) - List digital wallets
* [Get](docs/sdks/digitalwallets/README.md#get) - Get digital wallet
* [Delete](docs/sdks/digitalwallets/README.md#delete) - Delete digital wallet
* [Update](docs/sdks/digitalwallets/README.md#update) - Update digital wallet

#### [DigitalWallets.Domains](docs/sdks/domains/README.md)

* [Create](docs/sdks/domains/README.md#create) - Register a digital wallet domain
* [Delete](docs/sdks/domains/README.md#delete) - Remove a digital wallet domain

#### [DigitalWallets.Sessions](docs/sdks/sessions/README.md)

* [GooglePay](docs/sdks/sessions/README.md#googlepay) - Create a Google Pay session
* [ApplePay](docs/sdks/sessions/README.md#applepay) - Create a Apple Pay session
* [ClickToPay](docs/sdks/sessions/README.md#clicktopay) - Create a Click to Pay session

### [GiftCards](docs/sdks/giftcards/README.md)

* [Get](docs/sdks/giftcards/README.md#get) - Get gift card
* [Delete](docs/sdks/giftcards/README.md#delete) - Delete a gift card
* [Create](docs/sdks/giftcards/README.md#create) - Create gift card
* [List](docs/sdks/giftcards/README.md#list) - List gift cards

#### [GiftCards.Balances](docs/sdks/balances/README.md)

* [List](docs/sdks/balances/README.md#list) - List gift card balances


### [MerchantAccounts](docs/sdks/merchantaccounts/README.md)

* [List](docs/sdks/merchantaccounts/README.md#list) - List all merchant accounts
* [Create](docs/sdks/merchantaccounts/README.md#create) - Create a merchant account
* [Get](docs/sdks/merchantaccounts/README.md#get) - Get a merchant account
* [Update](docs/sdks/merchantaccounts/README.md#update) - Update a merchant account

### [PaymentMethods](docs/sdks/paymentmethods/README.md)

* [List](docs/sdks/paymentmethods/README.md#list) - List all payment methods
* [Create](docs/sdks/paymentmethods/README.md#create) - Create payment method
* [Get](docs/sdks/paymentmethods/README.md#get) - Get payment method
* [Delete](docs/sdks/paymentmethods/README.md#delete) - Delete payment method

#### [PaymentMethods.NetworkTokens](docs/sdks/networktokens/README.md)

* [List](docs/sdks/networktokens/README.md#list) - List network tokens
* [Create](docs/sdks/networktokens/README.md#create) - Provision network token
* [Suspend](docs/sdks/networktokens/README.md#suspend) - Suspend network token
* [Resume](docs/sdks/networktokens/README.md#resume) - Resume network token
* [Delete](docs/sdks/networktokens/README.md#delete) - Delete network token

#### [PaymentMethods.NetworkTokens.Cryptogram](docs/sdks/cryptogram/README.md)

* [Create](docs/sdks/cryptogram/README.md#create) - Provision network token cryptogram

#### [PaymentMethods.PaymentServiceTokens](docs/sdks/paymentservicetokens/README.md)

* [List](docs/sdks/paymentservicetokens/README.md#list) - List payment service tokens
* [Create](docs/sdks/paymentservicetokens/README.md#create) - Create payment service token
* [Delete](docs/sdks/paymentservicetokens/README.md#delete) - Delete payment service token

### [PaymentOptions](docs/sdks/paymentoptions/README.md)

* [List](docs/sdks/paymentoptions/README.md#list) - List payment options

### [PaymentServiceDefinitions](docs/sdks/paymentservicedefinitions/README.md)

* [List](docs/sdks/paymentservicedefinitions/README.md#list) - List payment service definitions
* [Get](docs/sdks/paymentservicedefinitions/README.md#get) - Get a payment service definition
* [Session](docs/sdks/paymentservicedefinitions/README.md#session) - Create a session for apayment service definition

### [PaymentServices](docs/sdks/paymentservices/README.md)

* [List](docs/sdks/paymentservices/README.md#list) - List payment services
* [Create](docs/sdks/paymentservices/README.md#create) - Update a configured payment service
* [Get](docs/sdks/paymentservices/README.md#get) - Get payment service
* [Update](docs/sdks/paymentservices/README.md#update) - Configure a payment service
* [Delete](docs/sdks/paymentservices/README.md#delete) - Delete a configured payment service
* [Verify](docs/sdks/paymentservices/README.md#verify) - Verify payment service credentials
* [Session](docs/sdks/paymentservices/README.md#session) - Create a session for apayment service definition

### [Payouts](docs/sdks/payouts/README.md)

* [List](docs/sdks/payouts/README.md#list) - List payouts created.
* [Create](docs/sdks/payouts/README.md#create) - Create a payout.
* [Get](docs/sdks/payouts/README.md#get) - Get a payout.

### [Refunds](docs/sdks/refunds/README.md)

* [Get](docs/sdks/refunds/README.md#get) - Get refund

### [Transactions](docs/sdks/transactions/README.md)

* [List](docs/sdks/transactions/README.md#list) - List transactions
* [Create](docs/sdks/transactions/README.md#create) - Create transaction
* [Get](docs/sdks/transactions/README.md#get) - Get transaction
* [Capture](docs/sdks/transactions/README.md#capture) - Capture transaction
* [Void](docs/sdks/transactions/README.md#void) - Void transaction
* [Summary](docs/sdks/transactions/README.md#summary) - Get transaction summary
* [Sync](docs/sdks/transactions/README.md#sync) - Sync transaction

#### [Transactions.Refunds](docs/sdks/transactionsrefunds/README.md)

* [List](docs/sdks/transactionsrefunds/README.md#list) - List transaction refunds
* [Create](docs/sdks/transactionsrefunds/README.md#create) - Create transaction refund
* [Get](docs/sdks/transactionsrefunds/README.md#get) - Get transaction refund

#### [Transactions.Refunds.All](docs/sdks/all/README.md)

* [Create](docs/sdks/all/README.md#create) - Create batch transaction refund

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Pagination [pagination] -->
## Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
```go
package main

import (
	"context"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := gr4vygo.New(
		gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
	)

	res, err := s.Buyers.List(ctx, operations.ListBuyersRequest{
		Cursor:             gr4vygo.String("ZXhhbXBsZTE"),
		Search:             gr4vygo.String("John"),
		ExternalIdentifier: gr4vygo.String("buyer-12345"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.CollectionBuyer != nil {
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
<!-- End Pagination [pagination] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/retry"
	"log"
	"models/operations"
	"os"
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
	}, nil, nil, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.AccountUpdaterJob != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/components"
	"github.com/gr4vy/gr4vy-go/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := gr4vygo.New(
		gr4vygo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
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
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Create` function may return the following errors:

| Error Type                                   | Status Code | Content Type     |
| -------------------------------------------- | ----------- | ---------------- |
| apierrors.Error400                           | 400         | application/json |
| apierrors.Error401                           | 401         | application/json |
| apierrors.Response403CreateAccountUpdaterJob | 403         | application/json |
| apierrors.Error404                           | 404         | application/json |
| apierrors.Error405                           | 405         | application/json |
| apierrors.Error409                           | 409         | application/json |
| apierrors.HTTPValidationError                | 422         | application/json |
| apierrors.Error425                           | 425         | application/json |
| apierrors.Error429                           | 429         | application/json |
| apierrors.Error500                           | 500         | application/json |
| apierrors.Error502                           | 502         | application/json |
| apierrors.Error504                           | 504         | application/json |
| apierrors.APIError                           | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/apierrors"
	"github.com/gr4vy/gr4vy-go/models/components"
	"log"
	"os"
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

		var e *apierrors.Error400
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.Error401
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.Response403CreateAccountUpdaterJob
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.Error404
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.Error405
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.Error409
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.HTTPValidationError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.Error425
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.Error429
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.Error500
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.Error502
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.Error504
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Name

You can override the default server globally using the `WithServer(server string)` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the names associated with the available servers:

| Name         | Server                               | Variables | Description |
| ------------ | ------------------------------------ | --------- | ----------- |
| `production` | `https://api.{id}.gr4vy.app`         | `id`      |             |
| `sandbox`    | `https://api.sandbox.{id}.gr4vy.app` | `id`      |             |

If the selected server has variables, you may override its default values using the associated option(s):

| Variable | Option              | Default     | Description                            |
| -------- | ------------------- | ----------- | -------------------------------------- |
| `id`     | `WithID(id string)` | `"example"` | The subdomain for your Gr4vy instance. |

#### Example

```go
package main

import (
	"context"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := gr4vygo.New(
		gr4vygo.WithServer("sandbox"),
		gr4vygo.WithID("<id>"),
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

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := gr4vygo.New(
		gr4vygo.WithServerURL("https://api.example.gr4vy.app"),
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Special Types [types] -->
## Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

### Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

#### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/gr4vy/gr4vy-go&utm_campaign=go)
