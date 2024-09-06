// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * My API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"
)

// ReviewsAPIController binds http requests to an api service and writes the service results to the http response
type ReviewsAPIController struct {
	service ReviewsAPIServicer
	errorHandler ErrorHandler
}

// ReviewsAPIOption for how the controller is set up.
type ReviewsAPIOption func(*ReviewsAPIController)

// WithReviewsAPIErrorHandler inject ErrorHandler into controller
func WithReviewsAPIErrorHandler(h ErrorHandler) ReviewsAPIOption {
	return func(c *ReviewsAPIController) {
		c.errorHandler = h
	}
}

// NewReviewsAPIController creates a default api controller
func NewReviewsAPIController(s ReviewsAPIServicer, opts ...ReviewsAPIOption) *ReviewsAPIController {
	controller := &ReviewsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ReviewsAPIController
func (c *ReviewsAPIController) Routes() Routes {
	return Routes{
		"PostReview": Route{
			strings.ToUpper("Post"),
			"/reviews",
			c.PostReview,
		},
	}
}

// PostReview - Post a review
func (c *ReviewsAPIController) PostReview(w http.ResponseWriter, r *http.Request) {
	postReviewInputBodyParam := PostReviewInputBody{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&postReviewInputBodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPostReviewInputBodyRequired(postReviewInputBodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertPostReviewInputBodyConstraints(postReviewInputBodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PostReview(r.Context(), postReviewInputBodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}