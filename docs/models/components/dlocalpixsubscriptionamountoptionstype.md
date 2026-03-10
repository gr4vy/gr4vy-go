# DlocalPIXSubscriptionAmountOptionsType

Indicates the amount type unit for the subscription. Allowed values are: `FIXED`, `VARIABLE`.

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.DlocalPIXSubscriptionAmountOptionsTypeFixed

// Open enum: custom values can be created with a direct type cast
custom := components.DlocalPIXSubscriptionAmountOptionsType("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `DlocalPIXSubscriptionAmountOptionsTypeFixed`    | FIXED                                            |
| `DlocalPIXSubscriptionAmountOptionsTypeVariable` | VARIABLE                                         |