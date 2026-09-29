# ResponseData

The 3DS data used for this transaction. To see full details about the 3DS calls please use our transaction events API.


## Supported Types

### ThreeDSecureDataV1

```go
responseData := components.CreateResponseDataThreeDSecureDataV1(components.ThreeDSecureDataV1{/* values here */})
```

### ThreeDSecureV2

```go
responseData := components.CreateResponseDataThreeDSecureV2(components.ThreeDSecureV2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch responseData.Type {
	case components.ResponseDataTypeThreeDSecureDataV1:
		// responseData.ThreeDSecureDataV1 is populated
	case components.ResponseDataTypeThreeDSecureV2:
		// responseData.ThreeDSecureV2 is populated
}
```
