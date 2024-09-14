/*
ogen generates both clients and servers from an OpenAPI v3 spec.
It provides an interface that you implement.
The individual route handlers do not conform to the http.Handler interface.
Path params and input bodies are added as fields in structs, added to the route method signature.
The response bodies are the generated models, put as the return value in the method signature.
Errors are interesting: if all routes implement the same default response, that response is made into
the default error, and its model can be easily wrapped with a status code and returned.
When not returning an error, the status code is set by the OpenAPI spec.
This library forces the use of its json parsing library dependency, jx, instead of letting you choose.
This library has some middleware integration with OpenTelemetry.
Has the ability to add middlewares, though the signature differs from stdlib.
*/
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

func (s *App) Greeting(ctx context.Context, params api.GreetingParams) (*api.GetGreetingOutputBody, error) {
	out := &api.GetGreetingOutputBody{
		Message: fmt.Sprintf("Hello %s!", params.Name),
	}
	return out, nil
}

func (s *App) PostReview(ctx context.Context, req *api.PostReviewInputBody) error {
	fmt.Printf("%+v\n", req)
	return nil
}

func (s *App) GetError(ctx context.Context) error {
	return &api.ErrorStatusCode{
		StatusCode: http.StatusBadRequest,
		Response: api.Error{
			Title:   "Bad Request",
			Details: api.NewOptString("This is an example error"),
			// Unfortunately Ogen generates free-form objects as a map of string->[]byte, which is super annoying
			Properties: api.NewOptNilErrorProperties(api.ErrorProperties{
				"cause": MustMarshal(224.92),
			}),
		},
	}
}

// A little confusing, but this method converts a regular non-*api.ErrorStatusCode error into that.
// It is only called if the error is not a *api.ErrorStatusCode or ErrNotImplemented.
// It contains a status code and our generated error model.
func (s *App) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	return &api.ErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: api.Error{
			Title:   "Internal Server Error",
			Details: api.NewOptString(err.Error()),
		},
	}
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
