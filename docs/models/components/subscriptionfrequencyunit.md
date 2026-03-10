# SubscriptionFrequencyUnit

Indicates the frequency unit for the subscription. Allowed values are: `DAY`, `WEEK`, `MONTH`, `BI_MONTHLY`, `QUARTER`, `SEMI_ANNUALLY`, `YEAR`, `ONDEMAND`.

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.SubscriptionFrequencyUnitMonth

// Open enum: custom values can be created with a direct type cast
custom := components.SubscriptionFrequencyUnit("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `SubscriptionFrequencyUnitMonth`        | MONTH                                   |
| `SubscriptionFrequencyUnitWeek`         | WEEK                                    |
| `SubscriptionFrequencyUnitBiMonthly`    | BI_MONTHLY                              |
| `SubscriptionFrequencyUnitOndemand`     | ONDEMAND                                |
| `SubscriptionFrequencyUnitQuarter`      | QUARTER                                 |
| `SubscriptionFrequencyUnitYear`         | YEAR                                    |
| `SubscriptionFrequencyUnitSemiAnnually` | SEMI_ANNUALLY                           |
| `SubscriptionFrequencyUnitDay`          | DAY                                     |