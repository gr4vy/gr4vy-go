# AccountType

Specify whether this is a `checking` or `savings` account

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.AccountTypeChecking

// Open enum: custom values can be created with a direct type cast
custom := components.AccountType("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `AccountTypeChecking` | checking              |
| `AccountTypeSavings`  | savings               |