# Flow

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.FlowCheckout

// Open enum: custom values can be created with a direct type cast
custom := components.Flow("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `FlowCheckout`            | checkout                  |
| `FlowCardTransaction`     | card-transaction          |
| `FlowNonCardTransaction`  | non-card-transaction      |
| `FlowRedirectTransaction` | redirect-transaction      |