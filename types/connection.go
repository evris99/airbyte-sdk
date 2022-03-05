package types

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/google/uuid"
)

type SupportedSyncModesEnum int

const (
	FullRefresh SupportedSyncModesEnum = iota
	Incremental
)

// Unmarshaler for json
func (su *SupportedSyncModesEnum) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch strings.ToLower(s) {
	case "full_refresh":
		*su = FullRefresh
	case "incremental":
		*su = Incremental
	}

	return nil
}

// Marshaler for json
func (su SupportedSyncModesEnum) MarshalJSON() ([]byte, error) {
	var s string
	switch su {
	case FullRefresh:
		s = "full_refresh"
	case Incremental:
		s = "incremental"
	}

	return json.Marshal(s)
}

type StreamType struct {
	Name                    string                 `json:"name,omitempty"`
	JsonSchema              map[string]interface{} `json:"jsonSchema,omitempty"`
	SupportedSyncModes      SupportedSyncModesEnum `json:"supportedSyncModes,omitempty"`
	SourceDefinedCursor     bool                   `json:"sourceDefinedCursor,omitempty"`
	DefaultCursorField      []string               `json:"defaultCursorField,omitempty"`
	SourceDefinedPrimaryKey [][]string             `json:"sourceDefinedPrimaryKey,omitempty"`
	Namespace               string                 `json:"namespace,omitempty"`
}

type Config struct {
	SyncMode            SupportedSyncModesEnum            `json:"syncMode,omitempty"`
	CursorField         []string                          `json:"cursorField,omitempty"`
	DestinationSyncMode SupportedDestinationSyncModesType `json:"destinationSyncMode,omitempty"`
	PrimaryKey          [][]string                        `json:"primaryKey,omitempty"`
	AliasName           string                            `json:"aliasName,omitempty"`
	Selected            bool                              `json:"selected,omitempty"`
}

type SyncCatalogType struct {
	Streams []struct {
		Stream *StreamType `json:"stream,omitempty"`
		Config *Config     `json:"config,omitempty"`
	} `json:"streams,omitempty"`
}

type TimeUnit int

const (
	Minutes TimeUnit = iota
	Hours
	Days
	Weeks
	Months
)

// Unmarshaler for json
func (t *TimeUnit) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch strings.ToLower(s) {
	case "minutes":
		*t = Minutes
	case "hours":
		*t = Hours
	case "days":
		*t = Days
	case "weeks":
		*t = Weeks
	case "months":
		*t = Months
	}

	return nil
}

// Marshaler for json
func (t TimeUnit) MarshalJSON() ([]byte, error) {
	var s string
	switch t {
	case Minutes:
		s = "minutes"
	case Hours:
		s = "hours"
	case Days:
		s = "days"
	case Weeks:
		s = "weeks"
	case Months:
		s = "months"
	}

	return json.Marshal(s)
}

type Schedule struct {
	Units    int      `json:"units,omitempty"`
	TimeUnit TimeUnit `json:"timeUnit"`
}

type ConnectionStatus int

const (
	Active ConnectionStatus = iota
	Inactive
	Deprecated
)

// Unmarshaler for json
func (c *ConnectionStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch strings.ToLower(s) {
	case "active":
		*c = Active
	case "inactive":
		*c = Inactive
	case "deprecated":
		*c = Deprecated
	}

	return nil
}

// Marshaler for json
func (c ConnectionStatus) MarshalJSON() ([]byte, error) {
	var s string
	switch c {
	case Active:
		s = "active"
	case Inactive:
		s = "inactive"
	case Deprecated:
		s = "deprecated"
	}

	return json.Marshal(s)
}

type ResourceRequirements struct {
	CpuRequest    string `json:"cpu_request,omitempty"`
	CpuLimit      string `json:"cpu_limit,omitempty"`
	MemoryRequest string `json:"memory_request,omitempty"`
	MemoryLimit   string `json:"memory_limit,omitempty"`
}

type Connection struct {
	ConnectionId         *uuid.UUID            `json:"connectionId,omitempty"`
	Name                 string                `json:"name,omitempty"`
	NamespaceDefinition  string                `json:"namespaceDefinition,omitempty"`
	NamespaceFormat      string                `json:"namespaceFormat,omitempty"`
	Prefix               string                `json:"prefix,omitempty"`
	SourceID             *uuid.UUID            `json:"sourceId,omitempty"`
	DestinationId        *uuid.UUID            `json:"destinationId,omitempty"`
	OperationIds         []uuid.UUID           `json:"operationIds"`
	SyncCatalog          *SyncCatalogType      `json:"syncCatalog,omitempty"`
	Schedule             *Schedule             `json:"schedule,omitempty"`
	Status               StatusType            `json:"status,omitempty"`
	ResourceRequirements *ResourceRequirements `json:"resourceRequirements,omitempty"`
}

// ConnectionToJSON reads json data from a Reader and returns a connection
func ConnectionFromJSON(r io.Reader) (*Connection, error) {
	conn := new(Connection)
	err := json.NewDecoder(r).Decode(conn)
	return conn, err
}

// ConnectionsToJSON reads json data from a Reader and returns a slice of connections
func ConnectionsFromJSON(r io.Reader) ([]Connection, error) {
	var connections struct {
		Connections []Connection `json:"connections"`
	}

	// Decode JSON
	err := json.NewDecoder(r).Decode(&connections)
	return connections.Connections, err
}
