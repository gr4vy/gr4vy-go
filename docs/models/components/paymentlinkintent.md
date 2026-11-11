# PaymentLinkIntent

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.PaymentLinkIntentAuthorize

// Open enum: custom values can be created with a direct type cast
custom := components.PaymentLinkIntent("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `PaymentLinkIntentAuthorize`       | authorize                          |
| `PaymentLinkIntentCapture`         | capture                            |
| `PaymentLinkIntentPreferAuthorize` | prefer_authorize                   |