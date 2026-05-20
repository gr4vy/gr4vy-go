# Source

The platform that the Paze session is being created for. Defaults to `web`.

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.SourceWeb

// Open enum: custom values can be created with a direct type cast
custom := components.Source("custom_value")
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `SourceWeb`    | web            |
| `SourceMobile` | mobile         |