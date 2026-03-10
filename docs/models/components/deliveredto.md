# DeliveredTo

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.DeliveredToShippingAddress

// Open enum: custom values can be created with a direct type cast
custom := components.DeliveredTo("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `DeliveredToShippingAddress` | shipping_address             |
| `DeliveredToStorePickup`     | store_pickup                 |