# RiskifiedAccountType

The customer account types recommended by Riskified.

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.RiskifiedAccountTypeGuest

// Open enum: custom values can be created with a direct type cast
custom := components.RiskifiedAccountType("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `RiskifiedAccountTypeGuest`      | guest                            |
| `RiskifiedAccountTypeRegistered` | registered                       |
| `RiskifiedAccountTypePremium`    | premium                          |
| `RiskifiedAccountTypeBusiness`   | business                         |