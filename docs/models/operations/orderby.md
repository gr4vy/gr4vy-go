# OrderBy

The direction to sort the payment methods in.

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/operations"
)

value := operations.OrderByAsc

// Open enum: custom values can be created with a direct type cast
custom := operations.OrderBy("custom_value")
```


## Values

| Name          | Value         |
| ------------- | ------------- |
| `OrderByAsc`  | asc           |
| `OrderByDesc` | desc          |