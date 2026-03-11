# InstrumentType

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.InstrumentTypePan

// Open enum: custom values can be created with a direct type cast
custom := components.InstrumentType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `InstrumentTypePan`           | pan                           |
| `InstrumentTypeCardToken`     | card_token                    |
| `InstrumentTypeRedirect`      | redirect                      |
| `InstrumentTypeRedirectToken` | redirect_token                |
| `InstrumentTypeGooglepay`     | googlepay                     |
| `InstrumentTypeApplepay`      | applepay                      |
| `InstrumentTypeNetworkToken`  | network_token                 |
| `InstrumentTypePlaid`         | plaid                         |
| `InstrumentTypeBankDetails`   | bank_details                  |