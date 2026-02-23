# RequiredFields2


## Supported Types

### 

```go
requiredFields2 := components.CreateRequiredFields2Boolean(bool{/* values here */})
```

### 

```go
requiredFields2 := components.CreateRequiredFields2MapOfRequiredFields1(map[string]components.RequiredFields1{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch requiredFields2.Type {
	case components.RequiredFields2TypeBoolean:
		// requiredFields2.Boolean is populated
	case components.RequiredFields2TypeMapOfRequiredFields1:
		// requiredFields2.MapOfRequiredFields1 is populated
}
```
