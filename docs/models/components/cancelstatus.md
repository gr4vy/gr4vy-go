# CancelStatus

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.CancelStatusSucceeded

// Open enum: custom values can be created with a direct type cast
custom := components.CancelStatus("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `CancelStatusSucceeded` | succeeded               |
| `CancelStatusPending`   | pending                 |
| `CancelStatusFailed`    | failed                  |