package types

import (
	"encoding/json"
	"strings"

	"github.com/google/uuid"
)

type Logs struct {
	LogLines []string `json:"logLines"`
}

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
	}

	return json.Marshal(s)
}

type JobInfo struct {
	ID         *uuid.UUID     `json:"id,omitempty"`
	ConfigType ConfigTypeEnum `json:"configType,omitempty"`
	ConfigId   string         `json:"configId,omitempty"`
	CreatedAt  int            `json:"createdAt,omitempty"`
	EndedAt    int            `json:"endedAt,omitempty"`
	Succeeded  bool           `json:"succeeded,omitempty"`
	Logs       *Logs          `json:"logLines,omitempty"`
}
