package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	server "github.com/veqryn/awesome-go-api/huma/openapi/oapi_codegen/server_stdlib/gen"
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

		fmt.Println(string(b))
	}
	w.WriteHeader(http.StatusCreated)
}

func main() {
	app := &App{}
	handler := server.HandlerWithOptions(app, server.StdHTTPServerOptions{})
	http.ListenAndServe(":8080", handler)
}
