# ResponseData

The 3DS data sent to the payment service for this transaction. This will only be populated if external 3DS data was passed in directly as part of the transaction API call, or if our 3DS server returned a status code of `Y` or `A`. In case of a failure to authenticate (status `N`, `R`, or `U`) this field will not be populated. To see full details about the 3DS calls please use our transaction events API.


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
