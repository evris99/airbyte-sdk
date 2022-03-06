package airbytesdk

import (
	"context"

	"github.com/evris99/airbyte-sdk/types"
	"github.com/google/uuid"
)

// CreateSource creates a new source
func (c *Client) CreateSource(ctx context.Context, source *types.Source) (*types.Source, error) {
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

// UpdateSource update a source
func (c *Client) UpdateSource(ctx context.Context, source *types.Source) (*types.Source, error) {
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

// ListWorkspaceSources returns all the source in the workspace with the given ID
func (c *Client) ListWorkspaceSources(ctx context.Context, workspaceID *uuid.UUID) ([]types.Source, error) {
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

// GetSource returns the source with the given ID
func (c *Client) GetSource(ctx context.Context, id *uuid.UUID) (*types.Source, error) {
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

// SearchSource searches for the given source
func (c *Client) SearchSource(ctx context.Context, source *types.Source) (*types.Source, error) {
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

// CloneSource makes a copy of the source with the given ID
func (c *Client) CloneSource(ctx context.Context, id *uuid.UUID) (*types.Source, error) {
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

// DeleteSource deletes the source with the given ID
func (c *Client) DeleteSource(ctx context.Context, id *uuid.UUID) error {
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

// CheckSourceConnection checks the connection to the source with the given ID
func (c *Client) CheckSourceConnection(ctx context.Context, id *uuid.UUID) (*types.ConnectionCheck, error) {
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

// CheckSourceConnectionUpdate checks the connection to the source with the given ID for updates
func (c *Client) CheckSourceConnectionUpdate(ctx context.Context, source *types.Source) (*types.ConnectionCheck, error) {
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
