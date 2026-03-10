# ProductType

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.ProductTypePhysical

// Open enum: custom values can be created with a direct type cast
custom := components.ProductType("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `ProductTypePhysical`    | physical                 |
| `ProductTypeDiscount`    | discount                 |
| `ProductTypeShippingFee` | shipping_fee             |
| `ProductTypeSalesTax`    | sales_tax                |
| `ProductTypeDigital`     | digital                  |
| `ProductTypeGiftCard`    | gift_card                |
| `ProductTypeStoreCredit` | store_credit             |
| `ProductTypeSurcharge`   | surcharge                |