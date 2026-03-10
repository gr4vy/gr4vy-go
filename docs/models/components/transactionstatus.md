# TransactionStatus

## Example Usage

```go
import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

value := components.TransactionStatusProcessing

// Open enum: custom values can be created with a direct type cast
custom := components.TransactionStatus("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `TransactionStatusProcessing`               | processing                                  |
| `TransactionStatusAuthorizationSucceeded`   | authorization_succeeded                     |
| `TransactionStatusAuthorizationDeclined`    | authorization_declined                      |
| `TransactionStatusAuthorizationFailed`      | authorization_failed                        |
| `TransactionStatusAuthorizationVoided`      | authorization_voided                        |
| `TransactionStatusAuthorizationVoidPending` | authorization_void_pending                  |
| `TransactionStatusCaptureSucceeded`         | capture_succeeded                           |
| `TransactionStatusCapturePending`           | capture_pending                             |
| `TransactionStatusBuyerApprovalPending`     | buyer_approval_pending                      |