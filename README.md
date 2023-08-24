# github.com/speakeasy-sdks/ccg-v1-go

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/ccg-v1-go
```
<!-- End SDK Installation -->

## SDK Example Usage
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Service](docs/sdks/service/README.md)

* [Getstatus](docs/sdks/service/README.md#getstatus) - get status

### [User](docs/sdks/user/README.md)

* [Getuser](docs/sdks/user/README.md#getuser) - get user
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
