# Frequency

Indicates the frequency unit for the subscription. Allowed values are: `WEEKLY`, `MONTHLY`, `QUARTERLY`, `SEMI_ANNUAL`, `ANNUAL`.

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.FrequencyWeekly

// Open enum: custom values can be created with a direct type cast
custom := components.Frequency("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `FrequencyWeekly`     | WEEKLY                |
| `FrequencyMonthly`    | MONTHLY               |
| `FrequencyQuarterly`  | QUARTERLY             |
| `FrequencySemiAnnual` | SEMI_ANNUAL           |
| `FrequencyAnnual`     | ANNUAL                |