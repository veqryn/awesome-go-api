package api

// Code generated by openapi-generator-go DO NOT EDIT., don't modify it manually

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// DefaultHandlerHandler handles the operations of the 'DefaultHandler' handler group.
type DefaultHandlerHandler interface {
	// GetError Responds with an error
	GetError(w http.ResponseWriter, r *http.Request)
	// Greeting Responds with a greeting
	Greeting(w http.ResponseWriter, r *http.Request)
	// PostReview Post a review to be saved
	PostReview(w http.ResponseWriter, r *http.Request)
}

// NewRouter creates a new router for the spec and the given handlers.
// Awesome GO API
//
// # Actual example use cases for a curated list of golang api generator libraries
//
// 1.0.0
func NewRouter(
	defaultHandlerHandler DefaultHandlerHandler,
) http.Handler {

	r := chi.NewRouter()

	// 'DefaultHandler' group

	// '/error'
	r.Options("/error", optionsHandlerFunc(
		http.MethodGet,
	))
	r.Get("/error", defaultHandlerHandler.GetError)

	// '/greeting/{name}'
	r.Options("/greeting/{name}", optionsHandlerFunc(
		http.MethodGet,
	))
	r.Get("/greeting/{name}", defaultHandlerHandler.Greeting)

	// '/reviews'
	r.Options("/reviews", optionsHandlerFunc(
		http.MethodPost,
	))
	r.Post("/reviews", defaultHandlerHandler.PostReview)

	return r
}

func optionsHandlerFunc(allowedMethods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Allow", strings.Join(allowedMethods, ", "))
	}
}