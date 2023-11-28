<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	ccgv1go "github.com/speakeasy-sdks/ccg-v1-go"
	"github.com/speakeasy-sdks/ccg-v1-go/pkg/models/shared"
	"log"
)

func main() {
	s := ccgv1go.New(
		ccgv1go.WithSecurity(shared.Security{
			HTTPCCG: "",
		}),
	)

	ctx := context.Background()
	res, err := s.Service.Getstatus(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.ServiceStatus != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->