package airbytesdk

import (
	"context"

	"github.com/evris99/airbyte-sdk/types"
	"github.com/google/uuid"
)

// Creates a connection between a source and a destination using the given context
func (c *Client) CreateConnectionWithContext(ctx context.Context, conn *types.Connection) (*types.Connection, error) {
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

// Creates a connection between a source and a destination.
// Equivalent with calling CreateConnectionWithContext with background as context
func (c *Client) CreateConnection(conn *types.Connection) (*types.Connection, error) {
	return c.CreateConnectionWithContext(context.Background(), conn)
}

// Upadates a connection between a source and a destination using the given context
func (c *Client) UpdateConnectionWithContext(ctx context.Context, conn *types.Connection) (*types.Connection, error) {
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

// Upadates a connection between a source and a destination.
// Equivalent with calling UpdateConnectionWithContext with background as context
func (c *Client) UpdateConnection(conn *types.Connection) (*types.Connection, error) {
	return c.UpdateConnectionWithContext(context.Background(), conn)
}

// Lists connections for workspace using the given context.
// Does not return deleted connections
func (c *Client) ListWorkspaceConnectionsWithContext(ctx context.Context, workspaceID *uuid.UUID) ([]types.Connection, error) {
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

// Returns connections for workspace.
// Does not return deleted connections.
// Equivalent with calling ListWorkspaceConnetionsWithContext with background as context
func (c *Client) ListWorkspaceConnections(workspaceID *uuid.UUID) ([]types.Connection, error) {
	return c.ListWorkspaceConnectionsWithContext(context.Background(), workspaceID)
}

// Lists connections for workspace using the given context, including deleted connections.
func (c *Client) ListAllWorkspaceConnectionsWithContext(ctx context.Context, workspaceID *uuid.UUID) ([]types.Connection, error) {
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

// Lists connections for workspace, including deleted connections.
// Equivalent with calling ListAllWorkspaceConnetionsWithContext with background as context
func (c *Client) ListAllWorkspaceConnections(workspaceID *uuid.UUID) ([]types.Connection, error) {
	return c.ListAllWorkspaceConnectionsWithContext(context.Background(), workspaceID)
}

// Returns the connection with the given ID using the given context
func (c *Client) GetConnectionWithContext(ctx context.Context, id *uuid.UUID) (*types.Connection, error) {
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

// Returns the connection with the given ID.
// Equivalent with calling GetConnectionWithContext with background as context
func (c *Client) GetConnection(id *uuid.UUID) (*types.Connection, error) {
	return c.GetConnectionWithContext(context.Background(), id)
}

// Searches for the connection using the given context
func (c *Client) SearchConnectionWithContext(ctx context.Context, conn *types.Connection) (*types.Connection, error) {
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

// Searches for the connection.
// Equivalent with calling SearchConnectionWithContext with background as context
func (c *Client) SearchConnection(conn *types.Connection) (*types.Connection, error) {
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
