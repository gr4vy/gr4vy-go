# UserDevice

The platform that is being used to access the website.

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.UserDeviceDesktop

// Open enum: custom values can be created with a direct type cast
custom := components.UserDevice("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `UserDeviceDesktop` | desktop             |
| `UserDeviceMobile`  | mobile              |