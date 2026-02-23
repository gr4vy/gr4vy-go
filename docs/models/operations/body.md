# Body


## Supported Types

### CardPaymentMethodCreate

```go
body := operations.CreateBodyCardPaymentMethodCreate(components.CardPaymentMethodCreate{/* values here */})
```

### RedirectPaymentMethodCreate

```go
body := operations.CreateBodyRedirectPaymentMethodCreate(components.RedirectPaymentMethodCreate{/* values here */})
```

### CheckoutSessionPaymentMethodCreate

```go
body := operations.CreateBodyCheckoutSessionPaymentMethodCreate(components.CheckoutSessionPaymentMethodCreate{/* values here */})
```

### PlaidPaymentMethodCreate

```go
body := operations.CreateBodyPlaidPaymentMethodCreate(components.PlaidPaymentMethodCreate{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch body.Type {
	case operations.BodyTypeCardPaymentMethodCreate:
		// body.CardPaymentMethodCreate is populated
	case operations.BodyTypeRedirectPaymentMethodCreate:
		// body.RedirectPaymentMethodCreate is populated
	case operations.BodyTypeCheckoutSessionPaymentMethodCreate:
		// body.CheckoutSessionPaymentMethodCreate is populated
	case operations.BodyTypePlaidPaymentMethodCreate:
		// body.PlaidPaymentMethodCreate is populated
}
```
