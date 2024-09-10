package main

import (
	"context"
	"fmt"
	"net/http"

	api "github.com/veqryn/awesome-go-api/openapiv3/ogen/gen"
)

var _ api.Handler = &App{}

type App struct {
	// server.Unimplemented // Can be used to future proof against new api endpoints
}

func (s *App) Greeting(ctx context.Context, params api.GreetingParams) (api.GreetingRes, error) {
	out := &api.GetGreetingOutputBody{
		Message: fmt.Sprintf("Hello %s!", params.Name),
	}
	return out, nil
}

func (s *App) PostReview(ctx context.Context, req *api.PostReviewInputBody) (api.PostReviewRes, error) {
	fmt.Printf("%+v\n", req)
	return &api.PostReviewCreated{}, nil
}

func (s *App) GetError(ctx context.Context) (*api.Error, error) {
	respErr := &api.Error{
		Title:   api.NewOptString("Bad Request"),
		Details: api.NewOptString("This is an example error"),
		// Looks like this didn't get generated as a map, and therefore we can't set arbitrary fields on it.
		// Also, the generated client can't handle empty structs, and instead wants them to be optional->nil.
		Properties: api.OptErrorProperties{},
	}
	return respErr, nil
}

func main() {
	app := &App{}
	server, err := api.NewServer(app)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", server)
}
