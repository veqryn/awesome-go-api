package main

import (
	"context"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	client "github.com/veqryn/awesome-go-api/protobufv3/openapiv3/oapi_codegen/gen"
)

func main() {
	ctx := context.Background()

	// Create client
	c, err := client.NewClientWithResponses("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	// Get a greeting. URL Params became arguments.
	fmt.Println("--- Getting a Greeting:")
	greeting, err := c.DefaultGreetingWithResponse(ctx, "Bob")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status: %s\n", greeting.Status())
	fmt.Printf("Code: %d\n", greeting.StatusCode())
	spew.Dump(greeting.JSON200)
	spew.Dump(greeting.JSONDefault)

	// Send a review. Body became an argument.
	fmt.Println("--- Sending a Review:")
	message := "foobar"
	review, err := c.DefaultReviewWithResponse(ctx, client.DefaultReviewJSONRequestBody{
		Author:  "Bob",
		Message: &message,
		Rating:  4,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Status: %s\n", review.Status())
	fmt.Printf("Code: %d\n", review.StatusCode())
	fmt.Printf("BODY: %s\n", review.Body)

	// Get an error.
	fmt.Println("--- Getting an Error:")
	errorResp, err := c.DefaultErrorWithResponse(ctx)
	if err != nil {
		// Bad responses with error messages don't end up here.
		// This is for things like an error making the connection.
		panic(err)
	}
	fmt.Printf("Status: %s\n", errorResp.Status())
	fmt.Printf("Code: %d\n", errorResp.StatusCode())
	spew.Dump(errorResp.JSON200)
	spew.Dump(errorResp.JSONDefault)
}
