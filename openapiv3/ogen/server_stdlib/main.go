package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-faster/jx"
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
		// Unfortunately Ogen generates free-form objects as a map of string->[]byte, which is annoying
		Properties: api.NewOptNilErrorProperties(api.ErrorProperties{
			"cause": MustMarshal(224.92),
		}),
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

func MustMarshal(v any) jx.Raw {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
