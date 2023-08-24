# User

### Available Operations

* [Getuser](#getuser) - get user

## Getuser

get user

### Example Usage

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
    res, err := s.User.Getuser(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetuserResponse](../../models/operations/getuserresponse.md), error**

