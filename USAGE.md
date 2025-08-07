<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	gr4vygo "github.com/gr4vy/gr4vy-go"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := gr4vygo.New(
		gr4vygo.WithMerchantAccountID("<id>"),
		gr4vygo.WithSecurity(os.Getenv("GR4VY_BEARER_AUTH")),
	)

	res, err := s.BrowsePaymentMethodDefinitionsGet(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->