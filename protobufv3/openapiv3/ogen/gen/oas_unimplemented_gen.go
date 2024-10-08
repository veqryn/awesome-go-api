// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// DefaultError implements Default_Error operation.
//
// Responds with an error.
//
// GET /error
func (UnimplementedHandler) DefaultError(ctx context.Context) error {
	return ht.ErrNotImplemented
}

// DefaultGreeting implements Default_Greeting operation.
//
// Responds with a greeting.
//
// GET /greeting/{name}
func (UnimplementedHandler) DefaultGreeting(ctx context.Context, params DefaultGreetingParams) (r *AwesomeGreetingResp, _ error) {
	return r, ht.ErrNotImplemented
}

// DefaultReview implements Default_Review operation.
//
// Post a review to be saved.
//
// POST /reviews
func (UnimplementedHandler) DefaultReview(ctx context.Context, req *AwesomeReviewReq) error {
	return ht.ErrNotImplemented
}

// NewError creates *RpcStatusStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *RpcStatusStatusCode) {
	r = new(RpcStatusStatusCode)
	return r
}
