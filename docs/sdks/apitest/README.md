# APITest SDK


## Overview

### Available Operations

* [PatchPets](#patchpets)

## PatchPets

### Example Usage

```go
package main

import(
	blakesapigo "github.com/speakeasy-sdks/blakes-api-go"
	"context"
	"log"
)

func main() {
    s := blakesapigo.New()

    ctx := context.Background()
    res, err := s.PatchPets(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.PatchPetsResponse](../../pkg/models/operations/patchpetsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
