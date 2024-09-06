package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

func Register(router *chi.Mux) huma.API {
	humaAPI := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))

	// Quick and Advanced version of registering an endpoint:
	// huma.Get(humaAPI, "/greeting/{name}", Greeting)
	huma.Register(humaAPI, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/greeting/{name}",
		Tags:        []string{"Greetings"},
		Summary:     "Say hello to someone",
		Description: "Respond with a greeting",
		OperationID: "greeting",
	}, GetGreeting)

	// huma.Post(humaAPI, "/reviews", PostReview)
	huma.Register(humaAPI, huma.Operation{
		Method:        http.MethodPost,
		DefaultStatus: http.StatusCreated,
		Path:          "/reviews",
		Summary:       "Post a review",
		Tags:          []string{"Reviews"},
		OperationID:   "post-review",
	}, PostReview)

	huma.Get(humaAPI, "/error", GetError)
	huma.Register(humaAPI, huma.Operation{
		Method:        http.MethodPost,
		DefaultStatus: http.StatusCreated,
		Path:          "/reviews",
		Summary:       "Post a review",
		Tags:          []string{"Reviews"},
		OperationID:   "post-review",
	}, PostReview)

	return humaAPI
}

type GetGreetingInput struct {
	Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
}

// GetGreetingOutput represents the greeting operation response.
type GetGreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func GetGreeting(ctx context.Context, input *GetGreetingInput) (*GetGreetingOutput, error) {
	resp := &GetGreetingOutput{}
	resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
	return resp, nil
}

// PostReviewInput represents the review operation request.
type PostReviewInput struct {
	Body struct {
		Author  string `json:"author" maxLength:"10" doc:"Author of the review"`
		Rating  int    `json:"rating" minimum:"1" maximum:"5" doc:"Rating from 1 to 5"`
		Message string `json:"message,omitempty" maxLength:"100" doc:"Review message"`
	}
}

func PostReview(ctx context.Context, input *PostReviewInput) (*struct{}, error) {
	slog.Info("review", "body", input.Body)
	return nil, nil
}

func GetError(ctx context.Context, input *struct{}) (*struct{}, error) {
	err := &huma.ErrorDetail{
		Message:  "root cause error",
		Location: "on the error page",
		Value:    224.92,
	}
	return nil, huma.Error400BadRequest("This is an example error", err)
}
