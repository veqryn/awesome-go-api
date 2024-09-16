package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	server "github.com/veqryn/awesome-go-api/openapiv3/oapi_codegen/server_stdlib/gen"
)

var _ server.ServerInterface = &App{}

type App struct {
	// server.Unimplemented // Can be used to future proof against new api endpoints
}

func (s *App) Greeting(w http.ResponseWriter, r *http.Request, name string) {
	out := server.GetGreetingOutputBody{
		Message: fmt.Sprintf("Hello %s!", name),
	}

	b, err := json.Marshal(out)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(b)
	if err != nil {
		panic(err)
	}
}

func (s *App) PostReview(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
		b, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		var input server.PostReviewInputBody
		err = json.Unmarshal(b, &input)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#+v\n", input)
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *App) GetError(w http.ResponseWriter, r *http.Request) {
	respErr := server.Error{
		Details: Ptr("This is an example error"),
		Properties: &map[string]any{
			"cause": 224.92,
		},
		Title: "Bad Request",
	}
	b, err := json.Marshal(respErr)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_, err = w.Write(b)
	if err != nil {
		panic(err)
	}
}

func main() {
	app := &App{}
	handler := server.HandlerWithOptions(app, server.StdHTTPServerOptions{})
	http.ListenAndServe(":8080", handler)
}

func Ptr[T any](t T) *T {
	return &t
}
