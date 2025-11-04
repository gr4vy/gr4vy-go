# Gr4vy Go SDK

Developer-friendly & type-safe Go SDK specifically catered to leverage the **Gr4vy** API.

<div align="left">
    <img alt="GitHub Tag" src="https://img.shields.io/github/v/tag/gr4vy/gr4vy-go?style=for-the-badge&label=Version&color=blue">
    <a href="https://www.speakeasy.com/?utm_source=github-com/gr4vy/gr4vy-go&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
</div>

## Summary

Gr4vy Go SDK

The official Gr4vy SDK for Go provides a convenient way to interact with the Gr4vy API from your server-side application. This SDK allows you to seamlessly integrate Gr4vy's powerful payment orchestration capabilities, including:

* Creating Transactions: Initiate and process payments with various payment methods and services.
* Managing Buyers: Store and manage buyer information securely.
* Storing Payment Methods: Securely store and tokenize payment methods for future use.
* Handling Webhooks: Easily process and respond to webhook events from Gr4vy.
* And much more: Access the full suite of Gr4vy API payment features.

This SDK is designed to simplify development, reduce boilerplate code, and help you get up and running with Gr4vy quickly and efficiently. It handles authentication, request signing, and provides easy-to-use methods for most API endpoints.

<!-- No Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [Gr4vy Go SDK](#gr4vy-go-sdk)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Bearer token generation](#bearer-token-generation)
  * [Embed token generation](#embed-token-generation)
  * [Merchant account ID selection](#merchant-account-id-selection)
  * [Webhooks verification](#webhooks-verification)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Pagination](#pagination)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Timeouts](#timeouts)
  * [Custom HTTP Client](#custom-http-client)
  * [Special Types](#special-types)
* [Development](#development)
  * [Testing](#testing)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/gr4vy/gr4vy-go
```
<!-- End SDK Installation [installation] -->

## SDK Example Usage

### Example

```go
package main

import (
	"context"
	gr4vy "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	privateKey := "...." // Private key loaded from disk or env var
	withToken := gr4vy.WithToken(privateKey, []JWTScope{ReadAll, WriteAll}, 60)

	s := gr4vy.New(
		gr4vy.WithID("example"),
		gr4vy.WithServer(gr4vy.ServerSandbox),
		gr4vy.WithSecuritySource(withToken),
		gr4vy.WithMerchantAccountID("default"),
	)

	res, err := s.Transactions.List(ctx, operations.ListTransactionsRequest{}, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```

<br /><br />
> [!IMPORTANT]
> Please use `WithToken` where the documentation mentions `os.Getenv("GR4VY_BEARER_AUTH")`.

<!-- No SDK Example Usage [usage] -->

## Bearer token generation

Alternatively, you can create a token for use with the SDK or with your own client library.

```go
import (
	gr4vy "github.com/gr4vy/gr4vy-go"
	"log"
	"os"
)

privateKey := "...." // Private key loaded from disk or env var

token, err := GetToken(privateKey, []JWTScope{ReadAll, WriteAll}, 5)
if err != nil {
	log.Fatal(err)
}
```

> **Note:** This will only create a token once. Use `WithToken` to dynamically generate a token
> for every request.


## Embed token generation

Alternatively, you can create a token for use with Embed as follows.

```go
import (
	gr4vy "github.com/gr4vy/gr4vy-go"
	"log"
	"os"
)

privateKey := "...." // Private key loaded from disk or env var

token, err := GetEmbedToken(privateKey, nil, "")
if err != nil {
	log.Fatal(err)
}
```

> **Note:** This will only create a token once. Use `withToken` to dynamically generate a token
> for every request.

## Merchant account ID selection

Depending on the key used, you might need to explicitly define a merchant account ID to use. In our API, 
this uses the `X-GR4VY-MERCHANT-ACCOUNT-ID` header. When using the SDK, you can set the `merchantAccountId`
on every request.

```js
s := gr4vy.New(
	gr4vy.WithID("example"),
	gr4vy.WithServer(gr4vy.ServerSandbox),
	gr4vy.WithSecuritySource(withToken),
	gr4vy.WithMerchantAccountID("my-merchant-id"),
)
```

## Webhooks verification

The SDK provides a `VerifyWebhook` method to validate incoming webhook requests from Gr4vy. This ensures that the webhook payload is authentic and has not been tampered with.

```go
import (
  "log"
  gr4vy "github.com/gr4vy/gr4vy-go"
)

func main() {
  secret := "your_webhook_secret"
  payload := "webhook_payload"
  signatureHeader := "signature_from_header"
  timestampHeader := "timestamp_from_header"
  timestampTolerance := 300 // Optional: Tolerance in seconds for timestamp validation

  err := gr4vy.VerifyWebhook(secret, payload, &signatureHeader, &timestampHeader, timestampTolerance)
  if err != nil {
    log.Fatal(err)
  }
}
```

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
		gr4vygo.WithMerchantAccountID("default"),
	)

	res, err := s.AccountUpdater.Jobs.Create(ctx, components.AccountUpdaterJobCreate{
		PaymentMethodIds: []string{
			"ef9496d8-53a5-4aad-8ca2-00eb68334389",
			"f29e886e-93cc-4714-b4a3-12b7a718e595",
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
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

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

### [PaymentLinks](docs/sdks/paymentlinks/README.md)

* [Create](docs/sdks/paymentlinks/README.md#create) - Add a payment link
* [List](docs/sdks/paymentlinks/README.md#list) - List all payment links
* [Expire](docs/sdks/paymentlinks/README.md#expire) - Expire a payment link
* [Get](docs/sdks/paymentlinks/README.md#get) - Get payment link

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
* [Session](docs/sdks/paymentservicedefinitions/README.md#session) - Create a session for a payment service definition

### [PaymentServices](docs/sdks/paymentservices/README.md)

* [List](docs/sdks/paymentservices/README.md#list) - List payment services
* [Create](docs/sdks/paymentservices/README.md#create) - Update a configured payment service
* [Get](docs/sdks/paymentservices/README.md#get) - Get payment service
* [Update](docs/sdks/paymentservices/README.md#update) - Configure a payment service
* [Delete](docs/sdks/paymentservices/README.md#delete) - Delete a configured payment service
* [Verify](docs/sdks/paymentservices/README.md#verify) - Verify payment service credentials
* [Session](docs/sdks/paymentservices/README.md#session) - Create a session for a payment service definition

### [Payouts](docs/sdks/payouts/README.md)

* [List](docs/sdks/payouts/README.md#list) - List payouts created
* [Create](docs/sdks/payouts/README.md#create) - Create a payout
* [Get](docs/sdks/payouts/README.md#get) - Get a payout

### [Refunds](docs/sdks/refunds/README.md)

* [Get](docs/sdks/refunds/README.md#get) - Get refund

### [ReportExecutions](docs/sdks/reportexecutions/README.md)

* [List](docs/sdks/reportexecutions/README.md#list) - List executed reports

### [Reports](docs/sdks/reports/README.md)

* [List](docs/sdks/reports/README.md#list) - List configured reports
* [Create](docs/sdks/reports/README.md#create) - Add a report
* [Get](docs/sdks/reports/README.md#get) - Get a report
* [Put](docs/sdks/reports/README.md#put) - Update a report

#### [Reports.Executions](docs/sdks/executions/README.md)

* [List](docs/sdks/executions/README.md#list) - List executions for report
* [URL](docs/sdks/executions/README.md#url) - Create URL for executed report
* [Get](docs/sdks/executions/README.md#get) - Get executed report

### [Transactions](docs/sdks/transactions/README.md)

* [List](docs/sdks/transactions/README.md#list) - List transactions
* [Create](docs/sdks/transactions/README.md#create) - Create transaction
* [Get](docs/sdks/transactions/README.md#get) - Get transaction
* [Update](docs/sdks/transactions/README.md#update) - Manually update a transaction
* [Capture](docs/sdks/transactions/README.md#capture) - Capture transaction
* [Void](docs/sdks/transactions/README.md#void) - Void transaction
* [Cancel](docs/sdks/transactions/README.md#cancel) - Cancel transaction
* [Sync](docs/sdks/transactions/README.md#sync) - Sync transaction

#### [Transactions.Events](docs/sdks/events/README.md)

* [List](docs/sdks/events/README.md#list) - List transaction events

#### [Transactions.Refunds](docs/sdks/transactionsrefunds/README.md)

* [List](docs/sdks/transactionsrefunds/README.md#list) - List transaction refunds
* [Create](docs/sdks/transactionsrefunds/README.md#create) - Create transaction refund
* [Get](docs/sdks/transactionsrefunds/README.md#get) - Get transaction refund

#### [Transactions.Refunds.All](docs/sdks/all/README.md)

* [Create](docs/sdks/all/README.md#create) - Create batch transaction refund

#### [Transactions.Settlements](docs/sdks/settlements/README.md)

* [Get](docs/sdks/settlements/README.md#get) - Get transaction settlement
* [List](docs/sdks/settlements/README.md#list) - List transaction settlements

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
		gr4vygo.WithMerchantAccountID("default"),
		gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
	)

	res, err := s.Buyers.List(ctx, operations.ListBuyersRequest{
		Cursor:             gr4vygo.Pointer("ZXhhbXBsZTE"),
		Search:             gr4vygo.Pointer("John"),
		ExternalIdentifier: gr4vygo.Pointer("buyer-12345"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
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
		gr4vygo.WithMerchantAccountID("default"),
		gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
	)

	res, err := s.AccountUpdater.Jobs.Create(ctx, components.AccountUpdaterJobCreate{
		PaymentMethodIds: []string{
			"ef9496d8-53a5-4aad-8ca2-00eb68334389",
			"f29e886e-93cc-4714-b4a3-12b7a718e595",
		},
	}, operations.WithRetries(
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
	if res != nil {
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
		gr4vygo.WithMerchantAccountID("default"),
		gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
	)

	res, err := s.AccountUpdater.Jobs.Create(ctx, components.AccountUpdaterJobCreate{
		PaymentMethodIds: []string{
			"ef9496d8-53a5-4aad-8ca2-00eb68334389",
			"f29e886e-93cc-4714-b4a3-12b7a718e595",
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
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Create` function may return the following errors:

| Error Type                    | Status Code | Content Type     |
| ----------------------------- | ----------- | ---------------- |
| apierrors.Error400            | 400         | application/json |
| apierrors.Error401            | 401         | application/json |
| apierrors.Error403            | 403         | application/json |
| apierrors.Error404            | 404         | application/json |
| apierrors.Error405            | 405         | application/json |
| apierrors.Error409            | 409         | application/json |
| apierrors.HTTPValidationError | 422         | application/json |
| apierrors.Error425            | 425         | application/json |
| apierrors.Error429            | 429         | application/json |
| apierrors.Error500            | 500         | application/json |
| apierrors.Error502            | 502         | application/json |
| apierrors.Error504            | 504         | application/json |
| apierrors.APIError            | 4XX, 5XX    | \*/\*            |

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
		gr4vygo.WithMerchantAccountID("default"),
		gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
	)

	res, err := s.AccountUpdater.Jobs.Create(ctx, components.AccountUpdaterJobCreate{
		PaymentMethodIds: []string{
			"ef9496d8-53a5-4aad-8ca2-00eb68334389",
			"f29e886e-93cc-4714-b4a3-12b7a718e595",
		},
	})
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

		var e *apierrors.Error403
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
| `sandbox`    | `https://api.sandbox.{id}.gr4vy.app` | `id`      |             |
| `production` | `https://api.{id}.gr4vy.app`         | `id`      |             |

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
		gr4vygo.WithID("example"),
		gr4vygo.WithMerchantAccountID("default"),
		gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
	)

	res, err := s.AccountUpdater.Jobs.Create(ctx, components.AccountUpdaterJobCreate{
		PaymentMethodIds: []string{
			"ef9496d8-53a5-4aad-8ca2-00eb68334389",
			"f29e886e-93cc-4714-b4a3-12b7a718e595",
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
		gr4vygo.WithServerURL("https://api.sandbox.example.gr4vy.app"),
		gr4vygo.WithMerchantAccountID("default"),
		gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
	)

	res, err := s.AccountUpdater.Jobs.Create(ctx, components.AccountUpdaterJobCreate{
		PaymentMethodIds: []string{
			"ef9496d8-53a5-4aad-8ca2-00eb68334389",
			"f29e886e-93cc-4714-b4a3-12b7a718e595",
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
<!-- End Server Selection [server] -->

## Timeouts

The GO SDK supports custom timeouts for the API requests. This timeout can be set globally or per request. To set it globally, pass the `WithTimeout` function to the SDK constructor.

```go
s := gr4vy.New(
	gr4vy.WithID("example"),
	gr4vy.WithServer(gr4vy.ServerSandbox),
	gr4vy.WithSecuritySource(withToken),
	gr4vy.WithMerchantAccountID("default"),
	// 5 second timeout
	gr4vy.WithTimeout(time.Duration(5*time.Second))
)
```

Alternatively, the timeout can be set for every API request.

```go
_, err = merchantClient.PaymentServices.Create(ctx, paymentServiceCreate, &merchantAccountID,
	operations.WithOperationTimeout(time.Duration(5*time.Second)),
)
```

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

	"github.com/gr4vy/gr4vy-go"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = gr4vygo.New(gr4vygo.WithClient(httpClient))
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

## Testing

To run the tests, install Go and run the following.

```sh
go install
go test -v
```

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/gr4vy/gr4vy-go&utm_campaign=go)
