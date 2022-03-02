package airbytesdk

import (
	"context"
	"encoding/json"
	"fmt"
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
	default:
		return fmt.Errorf("unknown sync mode")
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
	default:
		return nil, fmt.Errorf("unknown sync mode")
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
	default:
		return fmt.Errorf("unknown time unit")
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
	default:
		return nil, fmt.Errorf("unknown time unit")
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
	default:
		return fmt.Errorf("unknown connection status")
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
	default:
		return nil, fmt.Errorf("unknown connection status")
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

// Creates a connection between a source and a destination using the given context
func (c *Client) CreateConnectionWithContext(ctx context.Context, conn *Connection) (*Connection, error) {
	u, err := appendToURL(c.endpoint, "/v1/connections/create")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, conn)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode JSON
	newConn := new(Connection)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(newConn); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return newConn, nil
}

// Creates a connection between a source and a destination.
// Equivalent with calling CreateConnectionWithContext with background as context
func (c *Client) CreateConnection(conn *Connection) (*Connection, error) {
	return c.CreateConnectionWithContext(context.Background(), conn)
}

// Upadates a connection between a source and a destination using the given context
func (c *Client) UpdateConnectionWithContext(ctx context.Context, conn *Connection) (*Connection, error) {
	u, err := appendToURL(c.endpoint, "/v1/connections/update")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, conn)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode JSON
	updatedConn := new(Connection)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(updatedConn); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return updatedConn, nil
}

// Upadates a connection between a source and a destination.
// Equivalent with calling UpdateConnectionWithContext with background as context
func (c *Client) UpdateConnection(conn *Connection) (*Connection, error) {
	return c.UpdateConnectionWithContext(context.Background(), conn)
}

// Lists connections for workspace using the given context.
// Does not return deleted connections
func (c *Client) ListWorkspaceConnectionsWithContext(ctx context.Context, workspaceID *uuid.UUID) ([]Connection, error) {
	u, err := appendToURL(c.endpoint, "/v1/connections/list")
	if err != nil {
		return nil, err
	}

	data := make(map[string]*uuid.UUID)
	data["workspaceId"] = workspaceID

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var connections struct {
		Connections []Connection `json:"connections"`
	}

	// Decode JSON
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&connections); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return connections.Connections, nil
}

// Returns connections for workspace.
// Does not return deleted connections.
// Equivalent with calling ListWorkspaceConnetionsWithContext with background as context
func (c *Client) ListWorkspaceConnections(workspaceID *uuid.UUID) ([]Connection, error) {
	return c.ListWorkspaceConnectionsWithContext(context.Background(), workspaceID)
}

// Lists connections for workspace using the given context, including deleted connections.
func (c *Client) ListAllWorkspaceConnectionsWithContext(ctx context.Context, workspaceID *uuid.UUID) ([]Connection, error) {
	u, err := appendToURL(c.endpoint, "/v1/connections/list_all")
	if err != nil {
		return nil, err
	}

	data := make(map[string]*uuid.UUID)
	data["workspaceId"] = workspaceID

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var connections struct {
		Connections []Connection `json:"connections"`
	}

	// Decode JSON
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&connections); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return connections.Connections, nil
}

// Lists connections for workspace, including deleted connections.
// Equivalent with calling ListAllWorkspaceConnetionsWithContext with background as context
func (c *Client) ListAllWorkspaceConnections(workspaceID *uuid.UUID) ([]Connection, error) {
	return c.ListAllWorkspaceConnectionsWithContext(context.Background(), workspaceID)
}

// Returns the connection with the given ID using the given context
func (c *Client) GetConnectionWithContext(ctx context.Context, id *uuid.UUID) (*Connection, error) {
	u, err := appendToURL(c.endpoint, "/v1/connections/get")
	if err != nil {
		return nil, err
	}

	data := make(map[string]*uuid.UUID)
	data["connectionId"] = id

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode JSON
	conn := new(Connection)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(conn); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return conn, nil
}

// Returns the connection with the given ID.
// Equivalent with calling GetConnectionWithContext with background as context
func (c *Client) GetConnection(id *uuid.UUID) (*Connection, error) {
	return c.GetConnectionWithContext(context.Background(), id)
}

// Searches for the connection using the given context
func (c *Client) SearchConnectionWithContext(ctx context.Context, conn *Connection) (*Connection, error) {
	u, err := appendToURL(c.endpoint, "/v1/connections/search")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, conn)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode JSON
	foundConn := new(Connection)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(foundConn); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return foundConn, nil
}

// Searches for the connection.
// Equivalent with calling SearchConnectionWithContext with background as context
func (c *Client) SearchConnection(conn *Connection) (*Connection, error) {
	return c.SearchConnectionWithContext(context.Background(), conn)
}

// Deletes the connection with the given ID using the given context
func (c *Client) DeleteConnectionWithContext(ctx context.Context, id *uuid.UUID) error {
	u, err := appendToURL(c.endpoint, "/v1/connections/delete")
	if err != nil {
		return err
	}

	data := make(map[string]*uuid.UUID)
	data["connectionId"] = id

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// Deletes the connection with the given ID.
// Equivalent with calling DeleteConnectionWithContext with background as context
func (c *Client) DeleteConnection(id *uuid.UUID) error {
	return c.DeleteConnectionWithContext(context.Background(), id)
}
