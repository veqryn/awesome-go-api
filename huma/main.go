package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/veqryn/awesome-go-api/huma/api"
)

// curl -vv -k -L -X POST localhost:8080/reviews
// curl -vv -k -L -X POST localhost:8080/reviews -H 'Content-Type: application/json' --data '{"author": "bob", "rating": 3}'
// http://localhost:8888/greeting/bob
// http://localhost:8888/schemas/GreetingOutputBody.json
// http://localhost:8888/docs#/
// http://localhost:8888/openapi.json
// http://localhost:8888/openapi-3.0.yaml
func main() {
	router := chi.NewMux()
	api.Register(router)
	http.ListenAndServe(":8080", router)
}
