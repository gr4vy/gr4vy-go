# Interval

The cadence unit for the subscription plan.

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.IntervalDay

// Open enum: custom values can be created with a direct type cast
custom := components.Interval("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `IntervalDay`   | DAY             |
| `IntervalWeek`  | WEEK            |
| `IntervalMonth` | MONTH           |
| `IntervalYear`  | YEAR            |