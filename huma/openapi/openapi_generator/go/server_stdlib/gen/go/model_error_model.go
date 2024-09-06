// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * My API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package openapi




type ErrorModel struct {

	// A URL to the JSON Schema for this object.
	Schema string `json:"$schema,omitempty"`

	// A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail,omitempty"`

	// Optional list of individual error details
	Errors *[]ErrorDetail `json:"errors,omitempty"`

	// A URI reference that identifies the specific occurrence of the problem.
	Instance string `json:"instance,omitempty"`

	// HTTP status code
	Status int64 `json:"status,omitempty"`

	// A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
	Title string `json:"title,omitempty"`

	// A URI reference to human-readable documentation for the error.
	Type string `json:"type,omitempty"`
}

// AssertErrorModelRequired checks if the required fields are not zero-ed
func AssertErrorModelRequired(obj ErrorModel) error {
	if obj.Errors != nil {
		for _, el := range *obj.Errors {
			if err := AssertErrorDetailRequired(el); err != nil {
				return err
			}
		}
	}
	return nil
}

// AssertErrorModelConstraints checks if the values respects the defined constraints
func AssertErrorModelConstraints(obj ErrorModel) error {
    if obj.Errors != nil {
     	for _, el := range *obj.Errors {
     		if err := AssertErrorDetailConstraints(el); err != nil {
     			return err
     		}
     	}
    }
	return nil
}