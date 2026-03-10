# AuditLogAction

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.AuditLogActionCreated

// Open enum: custom values can be created with a direct type cast
custom := components.AuditLogAction("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `AuditLogActionCreated`  | created                  |
| `AuditLogActionUpdated`  | updated                  |
| `AuditLogActionDeleted`  | deleted                  |
| `AuditLogActionVoided`   | voided                   |
| `AuditLogActionCanceled` | canceled                 |
| `AuditLogActionCaptured` | captured                 |