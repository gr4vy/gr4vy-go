# TransactionIntent

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.TransactionIntentAuthorize

// Open enum: custom values can be created with a direct type cast
custom := components.TransactionIntent("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `TransactionIntentAuthorize` | authorize                    |
| `TransactionIntentCapture`   | capture                      |