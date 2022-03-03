package types

import (
	"encoding/json"
	"io"
)

type ValidationError struct {
	PropertyPath string `json:"propertyPath,omitempty"`
	InvalidValue string `json:"invalidValue,omitempty"`
	Message      string `json:"message,omitempty"`
}

// The server's response in case of an error
// It implements the error interface
type ResponseError struct {
	ID                          string            `json:"id,omitempty"`
	Message                     string            `json:"message,omitempty"`
	ExceptionClassName          string            `json:"exceptionClassName,omitempty"`
	ExceptionStack              []string          `json:"exceptionStack,omitempty"`
	ValidationErrors            []ValidationError `json:"validationErrors,omitempty"`
	RootCauseExceptionClassName string            `json:"rootCauseExceptionClassName,omitempty"`
	RootCauseExceptionStack     []string          `json:"rootCauseExceptionStack,omitempty"`
}

// The implementation of the error interface for ResponseError
func (e *ResponseError) Error() string {
	return e.Message
}

// ResponseErrorFromJSON reads JSON data from a Reader and returns a response error struct
func ResponseErrorFromJSON(r io.Reader) (*ResponseError, error) {
	responseErr := new(ResponseError)
	err := json.NewDecoder(r).Decode(responseErr)

	return responseErr, err
}
