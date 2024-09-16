/*
openapi-generator-go only generates servers from an OpenAPI v3 spec.
It provides an interface that you implement.
The individual routes are all compatible with the http.Handler interface.
Other than creating the interface and the models, this library and its generated code do nothing extra.
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	api "github.com/veqryn/awesome-go-api/openapiv3/openapi_generator_go/server_chi/gen"
)

var _ api.DefaultHandlerHandler = &App{}

type App struct {
	// server.Unimplemented // Can be used to future proof against new api endpoints
}

func (s *App) Greeting(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	out := api.GetGreetingOutputBody{
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

		var input api.PostReviewInputBody
		err = json.Unmarshal(b, &input)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#+v\n", input)
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *App) GetError(w http.ResponseWriter, r *http.Request) {
	respErr := api.Error{
		Details: "This is an example error",
		Properties: map[string]any{
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
	router := api.NewRouter(app)
	http.ListenAndServe(":8080", router)
}
