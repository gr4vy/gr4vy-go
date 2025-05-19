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

### CheckoutSessionWithURLPaymentMethodCreate

```go
transactionCreatePaymentMethod := components.CreateTransactionCreatePaymentMethodCheckoutSessionWithURLPaymentMethodCreate(components.CheckoutSessionWithURLPaymentMethodCreate{/* values here */})
```

