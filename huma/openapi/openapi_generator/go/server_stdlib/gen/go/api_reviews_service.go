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
	"context"
	"net/http"
	"errors"
)

// ReviewsAPIService is a service that implements the logic for the ReviewsAPIServicer
// This service should implement the business logic for every endpoint for the ReviewsAPI API.
// Include any external packages or services that will be required by this service.
type ReviewsAPIService struct {
}

// NewReviewsAPIService creates a default api service
func NewReviewsAPIService() *ReviewsAPIService {
	return &ReviewsAPIService{}
}

// PostReview - Post a review
func (s *ReviewsAPIService) PostReview(ctx context.Context, postReviewInputBody PostReviewInputBody) (ImplResponse, error) {
	// TODO - update PostReview with the required logic for this service method.
	// Add api_reviews_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, {}) or use other options such as http.Ok ...
	// return Response(201, nil),nil

	// TODO: Uncomment the next line to return response Response(0, ErrorModel{}) or use other options such as http.Ok ...
	// return Response(0, ErrorModel{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PostReview method not implemented")
}
