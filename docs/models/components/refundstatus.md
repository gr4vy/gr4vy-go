# RefundStatus

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.RefundStatusProcessing

// Open enum: custom values can be created with a direct type cast
custom := components.RefundStatus("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `RefundStatusProcessing` | processing               |
| `RefundStatusSucceeded`  | succeeded                |
| `RefundStatusFailed`     | failed                   |
| `RefundStatusDeclined`   | declined                 |
| `RefundStatusVoided`     | voided                   |