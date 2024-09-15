/*
ogen generates both clients and servers from an OpenAPI v3 spec.
The client turns path params and bodies into arguments on the client method calls.
The response is the expected parsed object model from the spec.
For routes that don't return anything, no response object is returned.
When a route returns an error, the method call puts that in the returned error.
*/
package main

import (
	"context"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	api "github.com/veqryn/awesome-go-api/openapiv3/ogen/gen"
)

func main() {
	ctx := context.Background()

	// Create client
	c, err := api.NewClient("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	// It seems as if this library's client completely hides the response status
	// codes, and instead you are expected to be dealing with a very complete
	// openapi spec where there are different components returned for different
	// status codes, with everything documented.

	// Get a greeting. URL Params became arguments.
	fmt.Println("--- Getting a Greeting:")
	greeting, err := c.Greeting(ctx, api.GreetingParams{Name: "Bob"})
	if err != nil {
		panic(err)
	}
	spew.Dump(greeting)

	// Send a review. Body became an argument.
	fmt.Println("--- Sending a Review:")
	message := "foobar"
	err = c.PostReview(ctx, &api.PostReviewInputBody{
		Author:  "Bob",
		Message: api.NewOptString(message),
		Rating:  4,
	})
	if err != nil {
		panic(err)
	}

	// Get an error.
	fmt.Println("--- Getting an Error:")
	err = c.GetError(ctx)
	spew.Dump(err)
}
