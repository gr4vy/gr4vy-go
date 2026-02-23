# Spec

The report specification.


## Supported Types

### AccountsReceivablesReportSpec

```go
spec := components.CreateSpecAccountsReceivables(components.AccountsReceivablesReportSpec{/* values here */})
```

### DetailedSettlementReportSpec

```go
spec := components.CreateSpecDetailedSettlement(components.DetailedSettlementReportSpec{/* values here */})
```

### TransactionRetriesReportSpec

```go
spec := components.CreateSpecTransactionRetries(components.TransactionRetriesReportSpec{/* values here */})
```

### TransactionsReportSpec

```go
spec := components.CreateSpecTransactions(components.TransactionsReportSpec{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch spec.Type {
	case components.SpecTypeAccountsReceivables:
		// spec.AccountsReceivablesReportSpec is populated
	case components.SpecTypeDetailedSettlement:
		// spec.DetailedSettlementReportSpec is populated
	case components.SpecTypeTransactionRetries:
		// spec.TransactionRetriesReportSpec is populated
	case components.SpecTypeTransactions:
		// spec.TransactionsReportSpec is populated
}
```
