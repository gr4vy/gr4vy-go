# PaymentLinkStatus

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.PaymentLinkStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.PaymentLinkStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `PaymentLinkStatusActive`     | active                        |
| `PaymentLinkStatusCompleted`  | completed                     |
| `PaymentLinkStatusExpired`    | expired                       |
| `PaymentLinkStatusProcessing` | processing                    |