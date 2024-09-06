package main

import (
	"context"
	"fmt"

	client "github.com/veqryn/awesome-go-api/huma/openapi/oapi_codegen/client/gen"
)

func main() {
	ctx := context.Background()

	c, err := client.NewClientWithResponses("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	greeting, err := c.GreetingWithResponse(ctx, "Bob")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Status: %s\n", greeting.Status())
	fmt.Printf("Code: %d\n", greeting.StatusCode())
	fmt.Printf("BODY: %s\n", greeting.Body)

	review, err := c.PostReviewWithResponse(ctx, client.PostReviewInputBody{
		Author:  "Bob",
		Message: nil,
		Rating:  4,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Status: %s\n", review.Status())
	fmt.Printf("Code: %d\n", review.StatusCode())
	fmt.Printf("BODY: %s\n", review.Body)
}
