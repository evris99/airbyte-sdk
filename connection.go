package airbytesdk

import (
	"context"

	"github.com/evris99/airbyte-sdk/types"
	"github.com/google/uuid"
)

// CreateConnection creates a connection between a source and a destination
func (c *Client) CreateConnection(ctx context.Context, conn *types.Connection) (*types.Connection, error) {
	u, err := appendToURL(c.endpoint, "/v1/connections/create")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, conn)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.ConnectionFromJSON(res.Body)
}

// UpdateConnection updates a connection between a source and a destination
func (c *Client) UpdateConnection(ctx context.Context, conn *types.Connection) (*types.Connection, error) {
	u, err := appendToURL(c.endpoint, "/v1/connections/update")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, conn)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.ConnectionFromJSON(res.Body)
}

// ListWorkspaceConnections lists connections for workspace.
// Does not return deleted connections
func (c *Client) ListWorkspaceConnections(ctx context.Context, workspaceID *uuid.UUID) ([]types.Connection, error) {
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

	return types.ConnectionsFromJSON(res.Body)
}

// ListAllWorkspaceConnections lists all connections for workspace, including deleted connections.
func (c *Client) ListAllWorkspaceConnections(ctx context.Context, workspaceID *uuid.UUID) ([]types.Connection, error) {
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

	return types.ConnectionsFromJSON(res.Body)
}

// GetConnection returns the connection with the given ID
func (c *Client) GetConnection(ctx context.Context, id *uuid.UUID) (*types.Connection, error) {
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

	return types.ConnectionFromJSON(res.Body)
}

// SearchConnection searches for the given connection
func (c *Client) SearchConnection(ctx context.Context, conn *types.Connection) (*types.Connection, error) {
	u, err := appendToURL(c.endpoint, "/v1/connections/search")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, conn)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.ConnectionFromJSON(res.Body)
}

// DeleteConnection deletes the connection with the given ID
func (c *Client) DeleteConnection(ctx context.Context, id *uuid.UUID) error {
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
