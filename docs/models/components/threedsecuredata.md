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

