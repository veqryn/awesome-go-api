package main

import (
	"context"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	api "github.com/veqryn/awesome-go-api/protobufv3/openapiv3/ogen/gen"
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
	greeting, err := c.DefaultGreeting(ctx, api.DefaultGreetingParams{Name: "Bob"})
	if err != nil {
		panic(err)
	}
	spew.Dump(greeting) // Interestingly this is a struct, unlike the pure openapiv3 version where it is an interface

	// Send a review. Body became an argument.
	fmt.Println("--- Sending a Review:")
	err = c.DefaultReview(ctx, &api.AwesomeReviewReq{
		Author:  "Bob",
		Message: api.NewOptString("foobar"),
		Rating:  4,
	})
	if err != nil {
		panic(err)
	}
	// Interestingly, we only get an error back here, no review resp. Ogen must know the response object is empty.

	// Get an error.
	fmt.Println("--- Getting an Error:")
	err = c.DefaultError(ctx)
	spew.Dump(err)
}
