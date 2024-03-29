# User
(*User*)

### Available Operations

* [Getuser](#getuser) - get user

## Getuser

get user

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/ccg-v1-go/pkg/models/shared"
	ccgv1go "github.com/speakeasy-sdks/ccg-v1-go"
	"context"
	"log"
)

func main() {
    s := ccgv1go.New(
        ccgv1go.WithSecurity(shared.Security{
            HTTPCCG: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
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

**[*operations.GetuserResponse](../../pkg/models/operations/getuserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
