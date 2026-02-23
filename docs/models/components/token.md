# Token

The opaque token as received from the Google Pay JS library. This format may change between JS library versions.


## Supported Types

### 

```go
token := components.CreateTokenStr(string{/* values here */})
```

### 

```go
token := components.CreateTokenMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch token.Type {
	case components.TokenTypeStr:
		// token.Str is populated
	case components.TokenTypeMapOfAny:
		// token.MapOfAny is populated
}
```
