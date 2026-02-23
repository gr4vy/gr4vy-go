# Response200VoidTransaction

Successful Response


## Supported Types

### Transaction

```go
response200VoidTransaction := operations.CreateResponse200VoidTransactionTransaction(components.Transaction{/* values here */})
```

### TransactionVoid

```go
response200VoidTransaction := operations.CreateResponse200VoidTransactionTransactionVoid(components.TransactionVoid{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch response200VoidTransaction.Type {
	case operations.Response200VoidTransactionTypeTransaction:
		// response200VoidTransaction.Transaction is populated
	case operations.Response200VoidTransactionTypeTransactionVoid:
		// response200VoidTransaction.TransactionVoid is populated
}
```
