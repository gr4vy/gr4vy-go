# Mode

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.ModeCard

// Open enum: custom values can be created with a direct type cast
custom := components.Mode("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `ModeCard`            | card                  |
| `ModeRedirect`        | redirect              |
| `ModeApplepay`        | applepay              |
| `ModeGooglepay`       | googlepay             |
| `ModeCheckoutSession` | checkout-session      |
| `ModeClickToPay`      | click-to-pay          |
| `ModeGiftCard`        | gift-card             |
| `ModeBank`            | bank                  |