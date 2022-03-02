package airbytesdk

import (
	"encoding/json"
	"fmt"
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

type Connection struct {
	Status  StatusType   `json:"status,omitempty"`
	Message string       `json:"message,omitempty"`
	JobInfo *JobInfoType `json:"jobInfo,omitempty"`
}
