# GiftCardRedemptionStatus

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.GiftCardRedemptionStatusCreated

// Open enum: custom values can be created with a direct type cast
custom := components.GiftCardRedemptionStatus("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `GiftCardRedemptionStatusCreated`   | created                             |
| `GiftCardRedemptionStatusSucceeded` | succeeded                           |
| `GiftCardRedemptionStatusFailed`    | failed                              |
| `GiftCardRedemptionStatusSkipped`   | skipped                             |