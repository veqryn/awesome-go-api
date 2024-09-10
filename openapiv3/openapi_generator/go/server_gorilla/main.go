package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/veqryn/awesome-go-api/openapiv3/openapi_generator/go/server_gorilla/gen/openapi"
)

type Service struct{}

func (s *Service) Greeting(ctx context.Context, name string) (openapi.ImplResponse, error) {
	// There is no type checking that the response is a GetGreetingOutputBody.
	// Content-Type headers are added automatically.
	// There is a generator flag to allow setting other headers.
	out := openapi.GetGreetingOutputBody{
		Message: fmt.Sprintf("Hello %s!", name),
	}
	return openapi.Response(http.StatusOK, out), nil
}

func (s *Service) PostReview(ctx context.Context, postReviewInputBody openapi.PostReviewInputBody) (openapi.ImplResponse, error) {
	fmt.Printf("%+v\n", postReviewInputBody)
	return openapi.Response(http.StatusCreated, nil), nil
}

func (s *Service) GetError(ctx context.Context) (openapi.ImplResponse, error) {
	respErr := openapi.Error{
		Details: "This is an example error",
		Properties: &map[string]interface{}{
			"cause": 224.92,
		},
		Title: "Bad Request",
	}

	return openapi.Response(http.StatusBadRequest, respErr), nil
}

func main() {
	log.Printf("Server started")

	// Services appear to be grouped by tags.
	// The interface does not match http.Handler, but is instead specific to each endpoint as
	// defined by the OpenAPI yaml
	svc := &Service{}
	defaultAPIController := openapi.NewDefaultAPIController(svc)

	// This uses gorilla/mux under the hood
	router := openapi.NewRouter(defaultAPIController)
	log.Fatal(http.ListenAndServe(":8080", router))
}
