# RequiredFields1


## Supported Types

### 

```go
requiredFields1 := components.CreateRequiredFields1Boolean(bool{/* values here */})
```

### 

```go
requiredFields1 := components.CreateRequiredFields1Any(any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch requiredFields1.Type {
	case components.RequiredFields1TypeBoolean:
		// requiredFields1.Boolean is populated
	case components.RequiredFields1TypeAny:
		// requiredFields1.Any is populated
}
```
