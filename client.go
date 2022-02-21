package airbytesdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	ErrInvalidEndpoint = errors.New("invalid api endpoint")
)

// A client to interact with the airbyte API using HTTP
type Client struct {
	// The underlying HTTP Client
	HttpClient *http.Client
	endpoint   *url.URL
}

type ValidationError struct {
	PropertyPath string `json:"propertyPath"`
	InvalidValue string `json:"invalidValue"`
	Message      string `json:"message"`
}

// The server's response in case of an error
// It implements the error interface
type ResponseError struct {
	ID                          string            `json:"id"`
	Message                     string            `json:"message"`
	ExceptionClassName          string            `json:"exceptionClassName"`
	ExceptionStack              []string          `json:"exceptionStack"`
	ValidationErrors            []ValidationError `json:"validationErrors"`
	RootCauseExceptionClassName string            `json:"rootCauseExceptionClassName"`
	RootCauseExceptionStack     string            `json:"rootCauseExceptionStack"`
}

// The implementation of the error interface for ResponseError
func (e *ResponseError) Error() string {
	return e.Message
}

// Creates and returns a new airbyte API client
func New(apiEndpoint string) (*Client, error) {
	_, err := url.ParseRequestURI(apiEndpoint)
	if err != nil {
		return nil, fmt.Errorf("could not parse URL: %w", err)
	}

	endpoint, err := url.Parse(apiEndpoint)
	if err != nil || endpoint.Scheme == "" || endpoint.Host == "" {
		return nil, fmt.Errorf("could not parse URL: %w", err)
	}

	return &Client{
		HttpClient: &http.Client{},
		endpoint:   endpoint,
	}, nil
}

func (c *Client) makeRequest(ctx context.Context, u *url.URL, data interface{}) (*http.Response, error) {

	// If the data exists encode it to json
	var httpBodyReader io.Reader
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("could not encode data: %w", err)
		}
		httpBodyReader = bytes.NewReader(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), httpBodyReader)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not execute request: %w", err)
	}

	// TODO: Handle error

	return res, nil
}
