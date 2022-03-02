package airbytesdk

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type Source struct {
	SourceId                *uuid.UUID  `json:"sourceId,omitempty"`
	SourceDefinitionId      *uuid.UUID  `json:"sourceDefinitionId,omitempty"`
	WorkspaceId             *uuid.UUID  `json:"workspaceId,omitempty"`
	ConnectionConfiguration interface{} `json:"connectionConfiguration,omitempty"`
	Name                    string      `json:"name,omitempty"`
	SourceName              string      `json:"sourceName,omitempty"`
}

// Create a new source using the given context
func (c *Client) CreateSourceWithContext(ctx context.Context, source *Source) (*Source, error) {
	u, err := appendToURL(c.endpoint, "/v1/sources/create")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, source)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode JSON
	newSource := new(Source)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(newSource); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return newSource, nil
}

// Create a new source.
// Equivalent with calling CreateSourceWithContext with background as context
func (c *Client) CreateSource(source *Source) (*Source, error) {
	return c.CreateSourceWithContext(context.Background(), source)
}

// Update a source using the given context
func (c *Client) UpdateSourceWithContext(ctx context.Context, source *Source) (*Source, error) {
	u, err := appendToURL(c.endpoint, "/v1/sources/update")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, source)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode JSON
	updatedSource := new(Source)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(updatedSource); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return updatedSource, nil
}

// Update a source.
// Equivalent with calling UpdateSourceWithContext with background as context
func (c *Client) UpdateSource(source *Source) (*Source, error) {
	return c.UpdateSourceWithContext(context.Background(), source)
}

// Returns all the source in the workspace with the give ID using the given context
func (c *Client) ListWorkspaceSourcesWithContext(ctx context.Context, workspaceID *uuid.UUID) ([]Source, error) {
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

	var sources struct {
		Sources []Source `json:"sources"`
	}

	// Decode JSON
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&sources); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return sources.Sources, nil
}

// Returns all the source in the workspace with the give ID.
// Equivalent with calling CreateSourceWithContext with background as context
func (c *Client) ListWorkspaceSources(workspaceID *uuid.UUID) ([]Source, error) {
	return c.ListWorkspaceSourcesWithContext(context.Background(), workspaceID)
}

// Returns a source with the given ID using the given context
func (c *Client) GetSourceWithContext(ctx context.Context, id *uuid.UUID) (*Source, error) {
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

	// Decode JSON
	source := new(Source)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(source); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return source, nil
}

// Returns a source with the given ID using the given context.
// Equivalent with calling GetSourceWithContext with background as context
func (c *Client) GetSource(id *uuid.UUID) (*Source, error) {
	return c.GetSourceWithContext(context.Background(), id)
}

// Searches for the given source using the given context
func (c *Client) SearchSourceWithContext(ctx context.Context, source *Source) (*Source, error) {
	u, err := appendToURL(c.endpoint, "/v1/sources/search")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, source)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	//Decode JSON
	foundSource := new(Source)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(foundSource); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return foundSource, nil
}

// Searches for the given source.
// Equivalent with calling SearchSourceWithContext with background as context
func (c *Client) SearchSource(source *Source) (*Source, error) {
	return c.SearchSourceWithContext(context.Background(), source)
}

// Makes a copy of the source with the given ID using the given context
func (c *Client) CloneSourceWithContext(ctx context.Context, id *uuid.UUID) (*Source, error) {
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

	// Decode JSON
	source := new(Source)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(source); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return source, nil
}

// Makes a copy of the source with the given ID.
// Equivalent with calling CloneSourceWithContext with background as context
func (c *Client) CloneSource(id *uuid.UUID) (*Source, error) {
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
func (c *Client) CheckSourceConnectionWithContext(ctx context.Context, id *uuid.UUID) (*ConnectionCheck, error) {
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

	// Decode JSON
	connection := new(ConnectionCheck)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(connection); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return connection, nil
}

// Checks the connection to the source with the given ID.
// Equivalent with calling CheckSourceConnectionWithContext with background as context
func (c *Client) CheckSourceConnection(id *uuid.UUID) (*ConnectionCheck, error) {
	return c.CheckSourceConnectionWithContext(context.Background(), id)
}

// Checks the connection to the source with the given ID for updates using the given context
func (c *Client) CheckSourceConnectionUpdateWithContext(ctx context.Context, source *Source) (*ConnectionCheck, error) {
	u, err := appendToURL(c.endpoint, "/v1/sources/check_connection_for_update")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, source)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode JSON
	connection := new(ConnectionCheck)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(connection); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return connection, nil
}

// Checks the connection to the source with the given ID for updates.
// Equivalent with calling CheckSourceConnectionUpdateWithContext with background as context
func (c *Client) CheckSourceConnectionUpdate(source *Source) (*ConnectionCheck, error) {
	return c.CheckSourceConnectionUpdateWithContext(context.Background(), source)
}
