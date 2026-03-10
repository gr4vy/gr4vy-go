# PayoutStatus

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.PayoutStatusDeclined

// Open enum: custom values can be created with a direct type cast
custom := components.PayoutStatus("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `PayoutStatusDeclined`  | declined                |
| `PayoutStatusFailed`    | failed                  |
| `PayoutStatusPending`   | pending                 |
| `PayoutStatusSucceeded` | succeeded               |