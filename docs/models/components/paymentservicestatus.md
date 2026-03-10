# PaymentServiceStatus

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.PaymentServiceStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.PaymentServiceStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `PaymentServiceStatusPending` | pending                       |
| `PaymentServiceStatusCreated` | created                       |
| `PaymentServiceStatusFailed`  | failed                        |