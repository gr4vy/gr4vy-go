# UserStatus

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.UserStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.UserStatus("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `UserStatusActive`  | active              |
| `UserStatusPending` | pending             |
| `UserStatusDeleted` | deleted             |