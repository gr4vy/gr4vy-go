# AntiFraudDecision

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.AntiFraudDecisionAccept

// Open enum: custom values can be created with a direct type cast
custom := components.AntiFraudDecision("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `AntiFraudDecisionAccept`    | accept                       |
| `AntiFraudDecisionError`     | error                        |
| `AntiFraudDecisionException` | exception                    |
| `AntiFraudDecisionReject`    | reject                       |
| `AntiFraudDecisionReview`    | review                       |
| `AntiFraudDecisionSkipped`   | skipped                      |
| `AntiFraudDecisionPending`   | pending                      |