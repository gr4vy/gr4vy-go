# ReportExecutionStatus

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.ReportExecutionStatusDispatched

// Open enum: custom values can be created with a direct type cast
custom := components.ReportExecutionStatus("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `ReportExecutionStatusDispatched` | dispatched                        |
| `ReportExecutionStatusFailed`     | failed                            |
| `ReportExecutionStatusPending`    | pending                           |
| `ReportExecutionStatusProcessing` | processing                        |
| `ReportExecutionStatusSucceeded`  | succeeded                         |