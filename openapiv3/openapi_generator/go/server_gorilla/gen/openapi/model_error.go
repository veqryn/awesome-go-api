// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Awesome GO API
 *
 * Actual example use cases for a curated list of golang api generator libraries
 *
 * API version: 1.0.0
 */

package openapi




// Error - An API error
type Error struct {

	// A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
	Title string `json:"title,omitempty"`

	// A human-readable explanation specific to this occurrence of the problem.
	Details string `json:"details,omitempty"`

	// Optional map of properties
	Properties *map[string]interface{} `json:"properties,omitempty"`
}

// AssertErrorRequired checks if the required fields are not zero-ed
func AssertErrorRequired(obj Error) error {
	return nil
}

// AssertErrorConstraints checks if the values respects the defined constraints
func AssertErrorConstraints(obj Error) error {
	return nil
}