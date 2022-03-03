package airbytesdk

import (
	"context"

	"github.com/evris99/airbyte-sdk/types"
	"github.com/google/uuid"
)

// Create a new source using the given context
func (c *Client) CreateSourceWithContext(ctx context.Context, source *types.Source) (*types.Source, error) {
	u, err := appendToURL(c.endpoint, "/v1/sources/create")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, source)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.SourceFromJSON(res.Body)
}

// Create a new source.
// Equivalent with calling CreateSourceWithContext with background as context
func (c *Client) CreateSource(source *types.Source) (*types.Source, error) {
	return c.CreateSourceWithContext(context.Background(), source)
}

// Update a source using the given context
func (c *Client) UpdateSourceWithContext(ctx context.Context, source *types.Source) (*types.Source, error) {
	u, err := appendToURL(c.endpoint, "/v1/sources/update")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, source)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.SourceFromJSON(res.Body)
}

// Update a source.
// Equivalent with calling UpdateSourceWithContext with background as context
func (c *Client) UpdateSource(source *types.Source) (*types.Source, error) {
	return c.UpdateSourceWithContext(context.Background(), source)
}

// Returns all the source in the workspace with the give ID using the given context
func (c *Client) ListWorkspaceSourcesWithContext(ctx context.Context, workspaceID *uuid.UUID) ([]types.Source, error) {
	u, err := appendToURL(c.endpoint, "/v1/sources/list")
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

	return types.SourcesFromJSON(res.Body)
}

// Returns all the source in the workspace with the give ID.
// Equivalent with calling CreateSourceWithContext with background as context
func (c *Client) ListWorkspaceSources(workspaceID *uuid.UUID) ([]types.Source, error) {
	return c.ListWorkspaceSourcesWithContext(context.Background(), workspaceID)
}

// Returns a source with the given ID using the given context
func (c *Client) GetSourceWithContext(ctx context.Context, id *uuid.UUID) (*types.Source, error) {
	u, err := appendToURL(c.endpoint, "/v1/sources/get")
	if err != nil {
		return nil, err
	}

	data := make(map[string]*uuid.UUID)
	data["sourceId"] = id

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.SourceFromJSON(res.Body)
}

// Returns a source with the given ID using the given context.
// Equivalent with calling GetSourceWithContext with background as context
func (c *Client) GetSource(id *uuid.UUID) (*types.Source, error) {
	return c.GetSourceWithContext(context.Background(), id)
}

// Searches for the given source using the given context
func (c *Client) SearchSourceWithContext(ctx context.Context, source *types.Source) (*types.Source, error) {
	u, err := appendToURL(c.endpoint, "/v1/sources/search")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, source)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.SourceFromJSON(res.Body)
}

// Searches for the given source.
// Equivalent with calling SearchSourceWithContext with background as context
func (c *Client) SearchSource(source *types.Source) (*types.Source, error) {
	return c.SearchSourceWithContext(context.Background(), source)
}

// Makes a copy of the source with the given ID using the given context
func (c *Client) CloneSourceWithContext(ctx context.Context, id *uuid.UUID) (*types.Source, error) {
	u, err := appendToURL(c.endpoint, "/v1/sources/clone")
	if err != nil {
		return nil, err
	}

	data := make(map[string]*uuid.UUID)
	data["sourceId"] = id

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.SourceFromJSON(res.Body)
}

// Makes a copy of the source with the given ID.
// Equivalent with calling CloneSourceWithContext with background as context
func (c *Client) CloneSource(id *uuid.UUID) (*types.Source, error) {
	return c.CloneSourceWithContext(context.Background(), id)
}

// Deletes a source with the given ID using the given context
func (c *Client) DeleteSourceWithContext(ctx context.Context, id *uuid.UUID) error {
	u, err := appendToURL(c.endpoint, "/v1/sources/delete")
	if err != nil {
		return err
	}

	data := make(map[string]*uuid.UUID)
	data["sourceId"] = id

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// Deletes a source with the given ID.
// Equivalent with calling DeleteSourceWithContext with background as context
func (c *Client) DeleteSource(id *uuid.UUID) error {
	return c.DeleteSourceWithContext(context.Background(), id)
}

// Checks the connection to the source with the given ID using the given context
func (c *Client) CheckSourceConnectionWithContext(ctx context.Context, id *uuid.UUID) (*types.ConnectionCheck, error) {
	u, err := appendToURL(c.endpoint, "/v1/sources/check_connection")
	if err != nil {
		return nil, err
	}

	data := make(map[string]*uuid.UUID)
	data["sourceId"] = id

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.ConnectionCheckFromJSON(res.Body)
}

// Checks the connection to the source with the given ID.
// Equivalent with calling CheckSourceConnectionWithContext with background as context
func (c *Client) CheckSourceConnection(id *uuid.UUID) (*types.ConnectionCheck, error) {
	return c.CheckSourceConnectionWithContext(context.Background(), id)
}

// Checks the connection to the source with the given ID for updates using the given context
func (c *Client) CheckSourceConnectionUpdateWithContext(ctx context.Context, source *types.Source) (*types.ConnectionCheck, error) {
	u, err := appendToURL(c.endpoint, "/v1/sources/check_connection_for_update")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, source)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.ConnectionCheckFromJSON(res.Body)
}

// Checks the connection to the source with the given ID for updates.
// Equivalent with calling CheckSourceConnectionUpdateWithContext with background as context
func (c *Client) CheckSourceConnectionUpdate(source *types.Source) (*types.ConnectionCheck, error) {
	return c.CheckSourceConnectionUpdateWithContext(context.Background(), source)
}
