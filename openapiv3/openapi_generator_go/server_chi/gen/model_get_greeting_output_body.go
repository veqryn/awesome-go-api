// Code generated by openapi-generator-go DO NOT EDIT.
//
// Source:
//
//	Title: Awesome GO API
//	Version: 1.0.0
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// GetGreetingOutputBody is an object.
type GetGreetingOutputBody struct {
	// Message: Greeting message
	Message string `json:"message" mapstructure:"message"`
}

// Validate implements basic validation for this model
func (m GetGreetingOutputBody) Validate() error {
	return validation.Errors{}.Filter()
}

// GetMessage returns the Message property
func (m GetGreetingOutputBody) GetMessage() string {
	return m.Message
}

// SetMessage sets the Message property
func (m *GetGreetingOutputBody) SetMessage(val string) {
	m.Message = val
}
