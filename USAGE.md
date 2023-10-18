<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	blakesapigo "github.com/speakeasy-sdks/blakes-api-go"
	"log"
)

func main() {
	s := blakesapigo.New()

	ctx := context.Background()
	res, err := s.APITest.PatchPets(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.PatchPets200ApplicationJSONOneOf != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->