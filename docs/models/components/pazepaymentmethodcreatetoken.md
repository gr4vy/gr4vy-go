# PazePaymentMethodCreateToken

The opaque token as received from the Paze complete response.


## Supported Types

### 

```go
pazePaymentMethodCreateToken := components.CreatePazePaymentMethodCreateTokenStr(string{/* values here */})
```

### 

```go
pazePaymentMethodCreateToken := components.CreatePazePaymentMethodCreateTokenMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch pazePaymentMethodCreateToken.Type {
	case components.PazePaymentMethodCreateTokenTypeStr:
		// pazePaymentMethodCreateToken.Str is populated
	case components.PazePaymentMethodCreateTokenTypeMapOfAny:
		// pazePaymentMethodCreateToken.MapOfAny is populated
}
```
