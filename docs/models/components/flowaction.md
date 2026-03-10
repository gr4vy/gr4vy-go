# FlowAction

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.FlowActionSelectPaymentOptions

// Open enum: custom values can be created with a direct type cast
custom := components.FlowAction("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `FlowActionSelectPaymentOptions` | select-payment-options           |
| `FlowActionRouteTransaction`     | route-transaction                |
| `FlowActionDeclineEarly`         | decline-early                    |
| `FlowActionSkip3ds`              | skip-3ds                         |