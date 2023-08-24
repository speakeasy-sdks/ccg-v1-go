<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/ccg-v1-go"
	"github.com/speakeasy-sdks/ccg-v1-go/pkg/models/shared"
)

func main() {
    s := ccgauth.New(
        ccgauth.WithSecurity(shared.Security{
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
<!-- End SDK Example Usage -->