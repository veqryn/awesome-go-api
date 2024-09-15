package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	openapi "github.com/veqryn/awesome-go-api/huma/openapi/openapi_generator/go/gen"
)

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
	greetingReq := c.GreetingsAPI.Greeting(ctx, "Bob")
	greetingOutput, resp, err := greetingReq.Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Code: %d\n", resp.StatusCode)
	spew.Dump(greetingOutput)

	// Create the body of the request. Note: the creation only has parameters
	// for the required fields. Option fields have to be set manually after.
	// Not sure why the 'New' method returns a pointer, when method to set it on
	// the request wants a value.
	fmt.Println("--- Sending a Review:")
	message := "foobar"
	review := openapi.NewPostReviewInputBody("bob", 4)
	review.Message = &message

	// Create a request, then set the body, before sending
	reviewReq := c.ReviewsAPI.PostReview(ctx)
	reviewReq = reviewReq.PostReviewInputBody(*review) // This returns a copy with the body
	resp, err = reviewReq.Execute()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Code: %d\n", resp.StatusCode)

	fmt.Println("--- Getting an Error:")
	errReq := c.DefaultAPI.GetError(ctx)
	resp, err = errReq.Execute()
	if err != nil {
		// The actual api returned error message is included in "err",
		// so you have to parse it to figure out if it was a network issue
		// or if there is a message.

		// One way is to just always try to parse the response body
		if resp != nil {
			// Unfortunately you need to parse this yourself
			var errorDetails openapi.ErrorModel
			jsonErr := json.NewDecoder(resp.Body).Decode(&errorDetails)
			if jsonErr != nil {
				panic(jsonErr)
			}
			spew.Dump(errorDetails)
		}

		// Another way, probably better, is to see if it is GenericOpenAPIError,
		// and if so, it will have the error in the Model() function.
		var parsedError *openapi.GenericOpenAPIError
		if errors.As(err, &parsedError) {
			spew.Dump(parsedError)
			spew.Dump(parsedError.Error())
			spew.Dump(parsedError.Model())
			spew.Dump(string(parsedError.Body()))

			if errorDetails, ok := parsedError.Model().(openapi.ErrorModel); ok {
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
