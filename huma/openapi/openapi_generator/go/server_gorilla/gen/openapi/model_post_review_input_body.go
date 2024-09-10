// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Awesome GO API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package openapi


import (
	"errors"
)



type PostReviewInputBody struct {

	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema,omitempty"`

	// Author of the review
	Author string `json:"author"`

	// Review message
	Message string `json:"message,omitempty"`

	// Rating from 1 to 5
	Rating int64 `json:"rating"`
}

// AssertPostReviewInputBodyRequired checks if the required fields are not zero-ed
func AssertPostReviewInputBodyRequired(obj PostReviewInputBody) error {
	elements := map[string]interface{}{
		"author": obj.Author,
		"rating": obj.Rating,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertPostReviewInputBodyConstraints checks if the values respects the defined constraints
func AssertPostReviewInputBodyConstraints(obj PostReviewInputBody) error {
	if obj.Rating < 1 {
		return &ParsingError{Param: "Rating", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Rating > 5 {
		return &ParsingError{Param: "Rating", Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
