# Item


## Supported Types

### GiftCardRequest

```go
item := components.CreateItemGiftCardRequest(components.GiftCardRequest{/* values here */})
```

### GiftCardStoredRequest

```go
item := components.CreateItemGiftCardStoredRequest(components.GiftCardStoredRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch item.Type {
	case components.ItemTypeGiftCardRequest:
		// item.GiftCardRequest is populated
	case components.ItemTypeGiftCardStoredRequest:
		// item.GiftCardStoredRequest is populated
}
```
