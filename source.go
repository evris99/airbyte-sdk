package airbytesdk

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type Source struct {
	SourceId                *uuid.UUID  `json:"sourceId"`
	SourceDefinitionId      *uuid.UUID  `json:"sourceDefinitionId"`
	WorkspaceId             *uuid.UUID  `json:"workspaceId"`
	ConnectionConfiguration interface{} `json:"connectionConfiguration"`
	Name                    string      `json:"name"`
	SourceName              string      `json:"sourceName"`
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
	if err := decoder.Decode(&newSource); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return newSource, nil
}

// Create a new source.
// Equivalent with calling CreateSourceWithContext with background as context
func (c *Client) CreateSource(ctx context.Context, source *Source) (*Source, error) {
	return c.CreateSourceWithContext(ctx, source)
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
	if err := decoder.Decode(&updatedSource); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return updatedSource, nil
}

// Update a source.
// Equivalent with calling CreateSourceWithContext with background as context
func (c *Client) UpdateSource(ctx context.Context, source *Source) (*Source, error) {
	return c.UpdateSourceWithContext(ctx, source)
}

// Create a new source using the given context
func (c *Client) ListSourcesWithContext(ctx context.Context) ([]Source, error) {
	u, err := appendToURL(c.endpoint, "/v1/sources/list")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, nil)
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

// Create a new source.
// Equivalent with calling CreateSourceWithContext with background as context
func (c *Client) ListSources(ctx context.Context, source *Source) ([]Source, error) {
	return c.ListSourcesWithContext(ctx)
}
