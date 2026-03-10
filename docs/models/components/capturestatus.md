# CaptureStatus

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.CaptureStatusSucceeded

// Open enum: custom values can be created with a direct type cast
custom := components.CaptureStatus("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `CaptureStatusSucceeded` | succeeded                |
| `CaptureStatusPending`   | pending                  |
| `CaptureStatusDeclined`  | declined                 |
| `CaptureStatusFailed`    | failed                   |