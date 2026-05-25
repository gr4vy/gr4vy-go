# TransactionType

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.TransactionTypePurchase

// Open enum: custom values can be created with a direct type cast
custom := components.TransactionType("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `TransactionTypePurchase`   | PURCHASE                    |
| `TransactionTypeCardOnFile` | CARD_ON_FILE                |
| `TransactionTypeBoth`       | BOTH                        |