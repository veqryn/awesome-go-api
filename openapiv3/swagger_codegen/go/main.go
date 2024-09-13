package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	openapi "github.com/veqryn/awesome-go-api/openapiv3/swagger_codegen/go/gen"
)


// Unfortunately this code does not compile due to a model not being generated
func main() {
	ctx := context.Background()

	// Create the client
	// There is probably a way to set the default host, but I haven't found it
	cfg := openapi.NewConfiguration()
	cfg.Scheme = "http"
	cfg.Host = "localhost:8080"
	c := openapi.NewAPIClient(cfg)

	// Create a request then send it. URL Parameters have become arguments.
	// There is no need to close the response, because the generated code has
	// read the full body, closed it, and replaced the response body with a
	// NoOpCloser of the Buffer.
	fmt.Println("--- Getting a Greeting:")
	greetingOutput, resp, err := c.DefaultApi.Greeting(ctx, "Bob")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Code: %d\n", resp.StatusCode)
	spew.Dump(greetingOutput)

	fmt.Println("--- Sending a Review:")
	resp, err = c.DefaultApi.PostReview(ctx, openapi.PostReviewInputBody{
		Author:  "Bob",
		Message: "foobar",
		Rating:  4,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Code: %d\n", resp.StatusCode)

	fmt.Println("--- Getting an Error:")
	resp, err = c.DefaultApi.GetError(ctx)
	if err != nil {
		// The actual api returned error message is included in "err",
		// so you have to parse it to figure out if it was a network issue
		// or if there is a message.

		// One way is to just always try to parse the response body
		if resp != nil {
			// Unfortunately you need to parse this yourself
			var errorDetails openapi.ModelError
			jsonErr := json.NewDecoder(resp.Body).Decode(&errorDetails)
			if jsonErr != nil {
				panic(jsonErr)
			}
			spew.Dump(errorDetails)
		}

		// Another way, probably better, is to see if it is GenericOpenAPIError,
		// and if so, it will have the error in the Model() function.
		var parsedError *openapi.GenericSwaggerError
		if errors.As(err, &parsedError) {
			spew.Dump(parsedError)
			spew.Dump(parsedError.Error())
			spew.Dump(parsedError.Model())
			spew.Dump(string(parsedError.Body()))

			if errorDetails, ok := parsedError.Model().(openapi.ModelError); ok {
				spew.Dump(errorDetails)
			}
		} else {
			spew.Dump(resp)
			spew.Dump(err)
			panic(err)
		}

	} else {
		fmt.Printf("Status: %s\n", resp.Status)
		fmt.Printf("Code: %d\n", resp.StatusCode)
		spew.Dump(resp)
	}
}
