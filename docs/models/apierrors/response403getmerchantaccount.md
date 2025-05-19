# Response403GetMerchantAccount

The credentials were invalid or the caller did not have permission to act on the resource.


## Supported Types

### Error403

```go
response403GetMerchantAccount := apierrors.CreateResponse403GetMerchantAccountError403(components.Error403{/* values here */})
```

### Error403Forbidden

```go
response403GetMerchantAccount := apierrors.CreateResponse403GetMerchantAccountError403Forbidden(components.Error403Forbidden{/* values here */})
```

### Error403Active

```go
response403GetMerchantAccount := apierrors.CreateResponse403GetMerchantAccountError403Active(components.Error403Active{/* values here */})
```

