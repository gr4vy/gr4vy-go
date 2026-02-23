# Response200CaptureTransaction

Successful Response


## Supported Types

### Transaction

```go
response200CaptureTransaction := operations.CreateResponse200CaptureTransactionTransaction(components.Transaction{/* values here */})
```

### TransactionCapture

```go
response200CaptureTransaction := operations.CreateResponse200CaptureTransactionTransactionCapture(components.TransactionCapture{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch response200CaptureTransaction.Type {
	case operations.Response200CaptureTransactionTypeTransaction:
		// response200CaptureTransaction.Transaction is populated
	case operations.Response200CaptureTransactionTypeTransactionCapture:
		// response200CaptureTransaction.TransactionCapture is populated
}
```
