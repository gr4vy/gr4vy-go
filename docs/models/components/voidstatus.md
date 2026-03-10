# VoidStatus

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.VoidStatusSucceeded

// Open enum: custom values can be created with a direct type cast
custom := components.VoidStatus("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `VoidStatusSucceeded` | succeeded             |
| `VoidStatusPending`   | pending               |
| `VoidStatusDeclined`  | declined              |
| `VoidStatusFailed`    | failed                |