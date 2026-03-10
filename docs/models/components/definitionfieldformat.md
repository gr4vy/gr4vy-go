# DefinitionFieldFormat

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.DefinitionFieldFormatText

// Open enum: custom values can be created with a direct type cast
custom := components.DefinitionFieldFormat("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `DefinitionFieldFormatText`      | text                             |
| `DefinitionFieldFormatMultiline` | multiline                        |
| `DefinitionFieldFormatNumber`    | number                           |
| `DefinitionFieldFormatTimezone`  | timezone                         |
| `DefinitionFieldFormatBoolean`   | boolean                          |