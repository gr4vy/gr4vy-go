# GooglePayPaymentMethodCreateToken

The opaque token as received from the Google Pay JS library. This format may change between JS library versions.


## Supported Types

### 

```go
googlePayPaymentMethodCreateToken := components.CreateGooglePayPaymentMethodCreateTokenStr(string{/* values here */})
```

### 

```go
googlePayPaymentMethodCreateToken := components.CreateGooglePayPaymentMethodCreateTokenMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch googlePayPaymentMethodCreateToken.Type {
	case components.GooglePayPaymentMethodCreateTokenTypeStr:
		// googlePayPaymentMethodCreateToken.Str is populated
	case components.GooglePayPaymentMethodCreateTokenTypeMapOfAny:
		// googlePayPaymentMethodCreateToken.MapOfAny is populated
}
```
