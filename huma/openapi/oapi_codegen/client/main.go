package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	client "github.com/veqryn/awesome-go-api/huma/openapi/oapi_codegen/client/gen"
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
	greeting, err := c.GreetingWithResponse(ctx, "Bob")
	if err != nil {
		panic(err)
	}

	// It would be nice if this was done automatically, so that you don't have
	// to guess which type goes with which function/endpoint.
	var greetingOutput client.GetGreetingOutputBody
	err = json.Unmarshal(greeting.Body, &greetingOutput)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status: %s\n", greeting.Status())
	fmt.Printf("Code: %d\n", greeting.StatusCode())
	spew.Dump(greetingOutput)

	// Send a review. Body became an argument.
	fmt.Println("--- Sending a Review:")
	message := "foobar"
	review, err := c.PostReviewWithResponse(ctx, client.PostReviewInputBody{
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
	errorResp, err := c.GetErrorWithResponse(ctx)
	if err != nil {
		// Bad responses with error messages don't end up here.
		// This is for things like an error making the connection.
		panic(err)
	}

	// I guess if the status code >= 400, we can always unmarshal to an error?
	var errorDetails client.ErrorModel
	err = json.Unmarshal(errorResp.Body, &errorDetails)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status: %s\n", errorResp.Status())
	fmt.Printf("Code: %d\n", errorResp.StatusCode())
	spew.Dump(errorDetails)
}
