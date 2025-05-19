<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"github.com/gr4vy/gr4vy-go/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := gr4vygo.New(
		"https://api.example.com",
		gr4vygo.WithSecurity(os.Getenv("GR4VY_O_AUTH2_PASSWORD_BEARER")),
	)

	res, err := s.AccountUpdater.Jobs.Create(ctx, components.AccountUpdaterJobCreate{
		PaymentMethodIds: []string{
			"ef9496d8-53a5-4aad-8ca2-00eb68334389",
			"f29e886e-93cc-4714-b4a3-12b7a718e595",
		},
	}, nil, gr4vygo.String("default"))
	if err != nil {
		log.Fatal(err)
	}
	if res.AccountUpdaterJob != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->