# Context


## Supported Types

### WalletPaymentOptionContext

```go
context := components.CreateContextWalletPaymentOptionContext(components.WalletPaymentOptionContext{/* values here */})
```

### GooglePayPaymentOptionContext

```go
context := components.CreateContextGooglePayPaymentOptionContext(components.GooglePayPaymentOptionContext{/* values here */})
```

### PaymentOptionContext

```go
context := components.CreateContextPaymentOptionContext(components.PaymentOptionContext{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch context.Type {
	case components.ContextTypeWalletPaymentOptionContext:
		// context.WalletPaymentOptionContext is populated
	case components.ContextTypeGooglePayPaymentOptionContext:
		// context.GooglePayPaymentOptionContext is populated
	case components.ContextTypePaymentOptionContext:
		// context.PaymentOptionContext is populated
}
```
