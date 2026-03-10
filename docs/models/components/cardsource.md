# CardSource

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.CardSourceApplePay

// Open enum: custom values can be created with a direct type cast
custom := components.CardSource("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `CardSourceApplePay`  | apple-pay             |
| `CardSourceGooglePay` | google-pay            |