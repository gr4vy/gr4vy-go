# PazeSessionCompleteRequestTransactiontype

The type of transaction being completed. PURCHASE for a one-off checkout, CARD_ON_FILE to retain the card for future use, or BOTH.

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.PazeSessionCompleteRequestTransactiontypePurchase

// Open enum: custom values can be created with a direct type cast
custom := components.PazeSessionCompleteRequestTransactiontype("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `PazeSessionCompleteRequestTransactiontypePurchase`   | PURCHASE                                              |
| `PazeSessionCompleteRequestTransactiontypeCardOnFile` | CARD_ON_FILE                                          |
| `PazeSessionCompleteRequestTransactiontypeBoth`       | BOTH                                                  |