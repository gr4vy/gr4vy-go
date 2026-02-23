# PayoutCreatePaymentMethod

The type of payment method to send funds too.


## Supported Types

### PaymentMethodCard

```go
payoutCreatePaymentMethod := components.CreatePayoutCreatePaymentMethodPaymentMethodCard(components.PaymentMethodCard{/* values here */})
```

### PaymentMethodStoredCard

```go
payoutCreatePaymentMethod := components.CreatePayoutCreatePaymentMethodPaymentMethodStoredCard(components.PaymentMethodStoredCard{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch payoutCreatePaymentMethod.Type {
	case components.PayoutCreatePaymentMethodTypePaymentMethodCard:
		// payoutCreatePaymentMethod.PaymentMethodCard is populated
	case components.PayoutCreatePaymentMethodTypePaymentMethodStoredCard:
		// payoutCreatePaymentMethod.PaymentMethodStoredCard is populated
}
```
