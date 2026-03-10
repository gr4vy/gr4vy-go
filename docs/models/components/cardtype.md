# CardType

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.CardTypeCredit

// Open enum: custom values can be created with a direct type cast
custom := components.CardType("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `CardTypeCredit`  | credit            |
| `CardTypeDebit`   | debit             |
| `CardTypePrepaid` | prepaid           |