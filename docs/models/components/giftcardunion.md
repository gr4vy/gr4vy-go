# GiftCardUnion


## Supported Types

### GiftCardTransactionCreate

```go
giftCardUnion := components.CreateGiftCardUnionGiftCardTransactionCreate(components.GiftCardTransactionCreate{/* values here */})
```

### GiftCardTokenTransactionCreate

```go
giftCardUnion := components.CreateGiftCardUnionGiftCardTokenTransactionCreate(components.GiftCardTokenTransactionCreate{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch giftCardUnion.Type {
	case components.GiftCardUnionTypeGiftCardTransactionCreate:
		// giftCardUnion.GiftCardTransactionCreate is populated
	case components.GiftCardUnionTypeGiftCardTokenTransactionCreate:
		// giftCardUnion.GiftCardTokenTransactionCreate is populated
}
```
