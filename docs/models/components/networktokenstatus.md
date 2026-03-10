# NetworkTokenStatus

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.NetworkTokenStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.NetworkTokenStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `NetworkTokenStatusActive`    | active                        |
| `NetworkTokenStatusInactive`  | inactive                      |
| `NetworkTokenStatusSuspended` | suspended                     |
| `NetworkTokenStatusDeleted`   | deleted                       |