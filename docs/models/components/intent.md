# Intent

Primary intent of the checkout session.

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.IntentReviewAndPay

// Open enum: custom values can be created with a direct type cast
custom := components.Intent("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `IntentReviewAndPay`    | REVIEW_AND_PAY          |
| `IntentExpressCheckout` | EXPRESS_CHECKOUT        |
| `IntentAddCard`         | ADD_CARD                |