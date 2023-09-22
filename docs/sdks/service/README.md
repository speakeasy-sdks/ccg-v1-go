# Service

### Available Operations

* [Getstatus](#getstatus) - get status

## Getstatus

get status

### Example Usage

```go
package main

import(
	"context"
	"log"
	ccgv1go "github.com/speakeasy-sdks/ccg-v1-go"
	"github.com/speakeasy-sdks/ccg-v1-go/pkg/models/shared"
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetstatusResponse](../../models/operations/getstatusresponse.md), error**

