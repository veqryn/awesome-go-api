//go:build go1.22

// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package server

import (
	"fmt"
	"net/http"

	"github.com/oapi-codegen/runtime"
)

// ErrorDetail defines model for ErrorDetail.
type ErrorDetail struct {
	// Location Where the error occurred, e.g. 'body.items[3].tags' or 'path.thing-id'
	Location *string `json:"location,omitempty"`

	// Message Error message text
	Message *string `json:"message,omitempty"`

	// Value The value at the given location
	Value *interface{} `json:"value,omitempty"`
}

// ErrorModel defines model for ErrorModel.
type ErrorModel struct {
	// Schema A URL to the JSON Schema for this object.
	Schema *string `json:"$schema,omitempty"`

	// Detail A human-readable explanation specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`

	// Errors Optional list of individual error details
	Errors *[]ErrorDetail `json:"errors"`

	// Instance A URI reference that identifies the specific occurrence of the problem.
	Instance *string `json:"instance,omitempty"`

	// Status HTTP status code
	Status *int64 `json:"status,omitempty"`

	// Title A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
	Title *string `json:"title,omitempty"`

	// Type A URI reference to human-readable documentation for the error.
	Type *string `json:"type,omitempty"`
}

// GetGreetingOutputBody defines model for GetGreetingOutputBody.
type GetGreetingOutputBody struct {
	// Schema A URL to the JSON Schema for this object.
	Schema *string `json:"$schema,omitempty"`

	// Message Greeting message
	Message string `json:"message"`
}

// PostReviewInputBody defines model for PostReviewInputBody.
type PostReviewInputBody struct {
	// Schema A URL to the JSON Schema for this object.
	Schema *string `json:"$schema,omitempty"`

	// Author Author of the review
	Author string `json:"author"`

	// Message Review message
	Message *string `json:"message,omitempty"`

	// Rating Rating from 1 to 5
	Rating int64 `json:"rating"`
}

// PostReviewJSONRequestBody defines body for PostReview for application/json ContentType.
type PostReviewJSONRequestBody = PostReviewInputBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get error
	// (GET /error)
	GetError(w http.ResponseWriter, r *http.Request)
	// Say hello to someone
	// (GET /greeting/{name})
	Greeting(w http.ResponseWriter, r *http.Request, name string)
	// Post a review
	// (POST /reviews)
	PostReview(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetError operation middleware
func (siw *ServerInterfaceWrapper) GetError(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetError(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// Greeting operation middleware
func (siw *ServerInterfaceWrapper) Greeting(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameterWithOptions("simple", "name", r.PathValue("name"), &name, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "name", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Greeting(w, r, name)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostReview operation middleware
func (siw *ServerInterfaceWrapper) PostReview(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostReview(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       *http.ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m *http.ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m *http.ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("GET "+options.BaseURL+"/error", wrapper.GetError)
	m.HandleFunc("GET "+options.BaseURL+"/greeting/{name}", wrapper.Greeting)
	m.HandleFunc("POST "+options.BaseURL+"/reviews", wrapper.PostReview)

	return m
}
