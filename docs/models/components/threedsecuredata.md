# ThreeDSecureData

Pass through 3-D Secure data to support external 3-D Secure authorisation. If using an external 3-D Secure provider, you should not pass a `redirect_url` in the `payment_method` object for a transaction.


## Supported Types

### ThreeDSecureDataV1

```go
threeDSecureData := components.CreateThreeDSecureDataThreeDSecureDataV1(components.ThreeDSecureDataV1{/* values here */})
```

### ThreeDSecureDataV2

```go
threeDSecureData := components.CreateThreeDSecureDataThreeDSecureDataV2(components.ThreeDSecureDataV2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch threeDSecureData.Type {
	case components.ThreeDSecureDataTypeThreeDSecureDataV1:
		// threeDSecureData.ThreeDSecureDataV1 is populated
	case components.ThreeDSecureDataTypeThreeDSecureDataV2:
		// threeDSecureData.ThreeDSecureDataV2 is populated
}
```
