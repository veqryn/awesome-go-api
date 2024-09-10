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

	switch t := greeting.(type) {
	case *api.GetGreetingOutputBody:
		spew.Dump(t)
	default:
		panic("unknown type")
	}

	// Send a review. Body became an argument.
	fmt.Println("--- Sending a Review:")
	message := "foobar"
	review, err := c.PostReview(ctx, &api.PostReviewInputBody{
		Author:  "Bob",
		Message: api.NewOptString(message),
		Rating:  4,
	})
	if err != nil {
		panic(err)
	}

	switch t := review.(type) {
	case *api.PostReviewCreated:
		spew.Dump(t)
	default:
		panic("unknown type")
	}

	// Get an error.
	fmt.Println("--- Getting an Error:")
	errorResp, err := c.GetError(ctx)
	if err != nil {
		panic(err)
	}

	spew.Dump(errorResp)
}
