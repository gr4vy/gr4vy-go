# ThreeDSecureStatus

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.ThreeDSecureStatusSetupError

// Open enum: custom values can be created with a direct type cast
custom := components.ThreeDSecureStatus("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `ThreeDSecureStatusSetupError` | setup_error                    |
| `ThreeDSecureStatusError`      | error                          |
| `ThreeDSecureStatusDeclined`   | declined                       |
| `ThreeDSecureStatusCancelled`  | cancelled                      |
| `ThreeDSecureStatusComplete`   | complete                       |