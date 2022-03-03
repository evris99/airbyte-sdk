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
	"strings"

	"github.com/google/uuid"
)

type ConfigTypeEnum int

const (
	CheckConnectionSource ConfigTypeEnum = iota
	CheckConnectionDestination
	DiscoverSchema
	GetSpec
	Sync
	ResetConnection
)

// Unmarshaler for json
func (a *ConfigTypeEnum) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch strings.ToLower(s) {
	case "check_connection_source":
		*a = CheckConnectionSource
	case "check_connection_destination":
		*a = CheckConnectionDestination
	case "discover_schema":
		*a = DiscoverSchema
	case "get_spec":
		*a = GetSpec
	case "sync":
		*a = Sync
	case "reset_connection":
		*a = ResetConnection
	default:
		return fmt.Errorf("unknown config type")
	}

	return nil
}

// Marshaler for json
func (a ConfigTypeEnum) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	case CheckConnectionSource:
		s = "check_connection_source"
	case CheckConnectionDestination:
		s = "check_connection_destination"
	case DiscoverSchema:
		s = "discover_schema"
	case GetSpec:
		s = "get_spec"
	case Sync:
		s = "sync"
	case ResetConnection:
		s = "reset_connection"
	default:
		return nil, fmt.Errorf("unknown config type")
	}

	return json.Marshal(s)
}

type LogsType struct {
	LogLines []string `json:"logLines"`
}

type JobInfoType struct {
	ID         *uuid.UUID     `json:"id,omitempty"`
	ConfigType ConfigTypeEnum `json:"configType,omitempty"`
	ConfigId   string         `json:"configId,omitempty"`
	CreatedAt  int            `json:"createdAt,omitempty"`
	EndedAt    int            `json:"endedAt,omitempty"`
	Succeeded  bool           `json:"succeeded,omitempty"`
	Logs       *LogsType      `json:"logLines,omitempty"`
}

var (
	ErrInvalidEndpoint = errors.New("invalid api endpoint")
	ErrServer          = errors.New("airbyte server error")
	ErrInvalidStatus   = errors.New("invalid server response status code")
)

// A client to interact with the airbyte API using HTTP
type Client struct {
	// The underlying HTTP Client
	HttpClient *http.Client
	endpoint   *url.URL
}

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

// Makes an HTTP API request with the give data as body
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

	// If response code is not 2XX return error
	if res.StatusCode >= 300 || res.StatusCode < 200 {
		return nil, getErrorResponse(res)
	}

	return res, nil
}

// Receives an HTTP response with a non 2XX status code
// And returns the according error
func getErrorResponse(res *http.Response) error {
	if res.StatusCode >= 400 && res.StatusCode < 600 {
		decoder := json.NewDecoder(res.Body)
		responseError := new(ResponseError)
		if err := decoder.Decode(responseError); err != nil {
			return fmt.Errorf("could not decode error json: %w", err)
		}

		return responseError
	}

	return ErrInvalidStatus
}

func appendToURL(u *url.URL, path string) (*url.URL, error) {
	return u.Parse(fmt.Sprintf("%s%s", u.Path, path))
}
