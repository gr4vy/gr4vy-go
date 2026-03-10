# PaymentMethodStatus

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.PaymentMethodStatusProcessing

// Open enum: custom values can be created with a direct type cast
custom := components.PaymentMethodStatus("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `PaymentMethodStatusProcessing`            | processing                                 |
| `PaymentMethodStatusBuyerApprovalRequired` | buyer_approval_required                    |
| `PaymentMethodStatusSucceeded`             | succeeded                                  |
| `PaymentMethodStatusFailed`                | failed                                     |
| `PaymentMethodStatusPaused`                | paused                                     |