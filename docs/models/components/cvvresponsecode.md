# CVVResponseCode

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.CVVResponseCodeMatch

// Open enum: custom values can be created with a direct type cast
custom := components.CVVResponseCode("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `CVVResponseCodeMatch`       | match                        |
| `CVVResponseCodeNoMatch`     | no_match                     |
| `CVVResponseCodeUnavailable` | unavailable                  |
| `CVVResponseCodeNotProvided` | not_provided                 |