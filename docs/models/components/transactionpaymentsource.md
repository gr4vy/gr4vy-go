# TransactionPaymentSource

The way payment method information made it to this transaction.

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.TransactionPaymentSourceEcommerce

// Open enum: custom values can be created with a direct type cast
custom := components.TransactionPaymentSource("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `TransactionPaymentSourceEcommerce`   | ecommerce                             |
| `TransactionPaymentSourceMoto`        | moto                                  |
| `TransactionPaymentSourceRecurring`   | recurring                             |
| `TransactionPaymentSourceInstallment` | installment                           |
| `TransactionPaymentSourceCardOnFile`  | card_on_file                          |