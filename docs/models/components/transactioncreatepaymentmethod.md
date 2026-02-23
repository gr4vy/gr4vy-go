# TransactionCreatePaymentMethod

The optional payment method to use for this transaction. This field is required if no `gift_cards` have been added.


## Supported Types

### CardWithURLPaymentMethodCreate

```go
transactionCreatePaymentMethod := components.CreateTransactionCreatePaymentMethodCardWithURLPaymentMethodCreate(components.CardWithURLPaymentMethodCreate{/* values here */})
```

### RedirectPaymentMethodCreate

```go
transactionCreatePaymentMethod := components.CreateTransactionCreatePaymentMethodRedirectPaymentMethodCreate(components.RedirectPaymentMethodCreate{/* values here */})
```

### TokenPaymentMethodCreate

```go
transactionCreatePaymentMethod := components.CreateTransactionCreatePaymentMethodTokenPaymentMethodCreate(components.TokenPaymentMethodCreate{/* values here */})
```

### ApplePayPaymentMethodCreate

```go
transactionCreatePaymentMethod := components.CreateTransactionCreatePaymentMethodApplePayPaymentMethodCreate(components.ApplePayPaymentMethodCreate{/* values here */})
```

### ClickToPayPaymentMethodCreate

```go
transactionCreatePaymentMethod := components.CreateTransactionCreatePaymentMethodClickToPayPaymentMethodCreate(components.ClickToPayPaymentMethodCreate{/* values here */})
```

### ClickToPayFPANPaymentMethodCreate

```go
transactionCreatePaymentMethod := components.CreateTransactionCreatePaymentMethodClickToPayFPANPaymentMethodCreate(components.ClickToPayFPANPaymentMethodCreate{/* values here */})
```

### GooglePayPaymentMethodCreate

```go
transactionCreatePaymentMethod := components.CreateTransactionCreatePaymentMethodGooglePayPaymentMethodCreate(components.GooglePayPaymentMethodCreate{/* values here */})
```

### GooglePayFPANPaymentMethodCreate

```go
transactionCreatePaymentMethod := components.CreateTransactionCreatePaymentMethodGooglePayFPANPaymentMethodCreate(components.GooglePayFPANPaymentMethodCreate{/* values here */})
```

### NetworkTokenPaymentMethodCreate

```go
transactionCreatePaymentMethod := components.CreateTransactionCreatePaymentMethodNetworkTokenPaymentMethodCreate(components.NetworkTokenPaymentMethodCreate{/* values here */})
```

### PlaidPaymentMethodCreate

```go
transactionCreatePaymentMethod := components.CreateTransactionCreatePaymentMethodPlaidPaymentMethodCreate(components.PlaidPaymentMethodCreate{/* values here */})
```

### CheckoutSessionWithURLPaymentMethodCreate

```go
transactionCreatePaymentMethod := components.CreateTransactionCreatePaymentMethodCheckoutSessionWithURLPaymentMethodCreate(components.CheckoutSessionWithURLPaymentMethodCreate{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch transactionCreatePaymentMethod.Type {
	case components.TransactionCreatePaymentMethodTypeCardWithURLPaymentMethodCreate:
		// transactionCreatePaymentMethod.CardWithURLPaymentMethodCreate is populated
	case components.TransactionCreatePaymentMethodTypeRedirectPaymentMethodCreate:
		// transactionCreatePaymentMethod.RedirectPaymentMethodCreate is populated
	case components.TransactionCreatePaymentMethodTypeTokenPaymentMethodCreate:
		// transactionCreatePaymentMethod.TokenPaymentMethodCreate is populated
	case components.TransactionCreatePaymentMethodTypeApplePayPaymentMethodCreate:
		// transactionCreatePaymentMethod.ApplePayPaymentMethodCreate is populated
	case components.TransactionCreatePaymentMethodTypeClickToPayPaymentMethodCreate:
		// transactionCreatePaymentMethod.ClickToPayPaymentMethodCreate is populated
	case components.TransactionCreatePaymentMethodTypeClickToPayFPANPaymentMethodCreate:
		// transactionCreatePaymentMethod.ClickToPayFPANPaymentMethodCreate is populated
	case components.TransactionCreatePaymentMethodTypeGooglePayPaymentMethodCreate:
		// transactionCreatePaymentMethod.GooglePayPaymentMethodCreate is populated
	case components.TransactionCreatePaymentMethodTypeGooglePayFPANPaymentMethodCreate:
		// transactionCreatePaymentMethod.GooglePayFPANPaymentMethodCreate is populated
	case components.TransactionCreatePaymentMethodTypeNetworkTokenPaymentMethodCreate:
		// transactionCreatePaymentMethod.NetworkTokenPaymentMethodCreate is populated
	case components.TransactionCreatePaymentMethodTypePlaidPaymentMethodCreate:
		// transactionCreatePaymentMethod.PlaidPaymentMethodCreate is populated
	case components.TransactionCreatePaymentMethodTypeCheckoutSessionWithURLPaymentMethodCreate:
		// transactionCreatePaymentMethod.CheckoutSessionWithURLPaymentMethodCreate is populated
}
```
