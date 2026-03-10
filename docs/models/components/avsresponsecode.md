# AVSResponseCode

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.AVSResponseCodeMatch

// Open enum: custom values can be created with a direct type cast
custom := components.AVSResponseCode("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `AVSResponseCodeMatch`                | match                                 |
| `AVSResponseCodeNoMatch`              | no_match                              |
| `AVSResponseCodePartialMatchAddress`  | partial_match_address                 |
| `AVSResponseCodePartialMatchPostcode` | partial_match_postcode                |
| `AVSResponseCodePartialMatchName`     | partial_match_name                    |
| `AVSResponseCodeUnavailable`          | unavailable                           |