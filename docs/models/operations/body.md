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

### ACHBankPaymentMethodCreate

```go
body := operations.CreateBodyACHBankPaymentMethodCreate(components.ACHBankPaymentMethodCreate{/* values here */})
```

### BACSBankPaymentMethodCreate

```go
body := operations.CreateBodyBACSBankPaymentMethodCreate(components.BACSBankPaymentMethodCreate{/* values here */})
```

### SEPABankPaymentMethodCreate

```go
body := operations.CreateBodySEPABankPaymentMethodCreate(components.SEPABankPaymentMethodCreate{/* values here */})
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
	case operations.BodyTypeACHBankPaymentMethodCreate:
		// body.ACHBankPaymentMethodCreate is populated
	case operations.BodyTypeBACSBankPaymentMethodCreate:
		// body.BACSBankPaymentMethodCreate is populated
	case operations.BodyTypeSEPABankPaymentMethodCreate:
		// body.SEPABankPaymentMethodCreate is populated
}
```
