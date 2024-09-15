/*
oapi-codegen generates both clients and servers from an OpenAPI v3 spec.
The client turns path params and bodies into arguments on the client method calls.
The response is a wrapper tha contains the http.Response, status code, and the expected parsed object model from the spec.
For routes that don't return anything, a wrapper is still returned from the call.
When a route returns an error, the method call puts it in the response wrapper,
not in the returned error.
The response object contains each possible object model described by the spec,
and you must pick which one based on the status code.
*/
package main

import (
	"context"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	client "github.com/veqryn/awesome-go-api/openapiv3/oapi_codegen/client/gen"
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

	fmt.Printf("Status: %s\n", greeting.Status())
	fmt.Printf("Code: %d\n", greeting.StatusCode())
	spew.Dump(greeting.JSON200)
	spew.Dump(greeting.JSONDefault)

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

	fmt.Printf("Status: %s\n", errorResp.Status())
	fmt.Printf("Code: %d\n", errorResp.StatusCode())
	spew.Dump(errorResp.JSONDefault)
}
