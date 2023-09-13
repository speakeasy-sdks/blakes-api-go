# APITest SDK

## Overview

### Available Operations

* [PatchPets](#patchpets)

## PatchPets

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/blakes-api-go"
)

func main() {
    s := apitest.New()

    ctx := context.Background()
    res, err := s.APITest.PatchPets(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.PatchPets200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.PatchPetsResponse](../../models/operations/patchpetsresponse.md), error**

