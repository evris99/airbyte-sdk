package types

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type StatusType int

const (
	Succeeded StatusType = iota
	Failed
)

// Unmarshaler for json
func (st *StatusType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch strings.ToLower(s) {
	case "succeeded":
		*st = Succeeded
	case "failed":
		*st = Failed
	default:
		return fmt.Errorf("unknown status")
	}

	return nil
}

// Marshaler for json
func (st StatusType) MarshalJSON() ([]byte, error) {
	var s string
	switch st {
	case Succeeded:
		s = "succeeded"
	case Failed:
		s = "failed"
	default:
		return nil, fmt.Errorf("unknown status")
	}

	return json.Marshal(s)
}

type ConnectionCheck struct {
	Status  StatusType `json:"status,omitempty"`
	Message string     `json:"message,omitempty"`
	JobInfo *JobInfo   `json:"jobInfo,omitempty"`
}

// ConnectionCheckFromJSON reads json data from a Reader and returns a workspace
func ConnectionCheckFromJSON(r io.Reader) (*ConnectionCheck, error) {
	connCheck := new(ConnectionCheck)
	err := json.NewDecoder(r).Decode(connCheck)

	return connCheck, fmt.Errorf("could not decode JSON: %w", err)
}
