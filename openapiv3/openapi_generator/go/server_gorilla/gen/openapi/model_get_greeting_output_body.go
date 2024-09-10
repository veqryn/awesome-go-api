// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Awesome GO API
 *
 * Actual example use cases for a curated list of golang api generator libraries
 *
 * API version: 1.0.0
 */

package openapi




type GetGreetingOutputBody struct {

	// Greeting message
	Message string `json:"message"`
}

// AssertGetGreetingOutputBodyRequired checks if the required fields are not zero-ed
func AssertGetGreetingOutputBodyRequired(obj GetGreetingOutputBody) error {
	elements := map[string]interface{}{
		"message": obj.Message,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertGetGreetingOutputBodyConstraints checks if the values respects the defined constraints
func AssertGetGreetingOutputBodyConstraints(obj GetGreetingOutputBody) error {
	return nil
}
