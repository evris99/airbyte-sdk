package airbytesdk

import (
	"context"

	"github.com/evris99/airbyte-sdk/types"
	"github.com/google/uuid"
)

// CreateSourceDefinition creates a new source definition and returns it
func (c *Client) CreateSourceDefinition(ctx context.Context, definition *types.SourceDefinition) (*types.SourceDefinition, error) {
	u, err := appendToURL(c.endpoint, "/v1/source_definitions/create")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, definition)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.SourceDefinitionFromJSON(res.Body)
}

// UpdateSourceDefinitionDockerImage updates a source definition and returns it.
// Currently, the only allowed attribute to update is the default docker image version.
func (c *Client) UpdateSourceDefinitionDockerImage(ctx context.Context, id *uuid.UUID, dockerImageTag string) (*types.SourceDefinition, error) {
	u, err := appendToURL(c.endpoint, "/v1/source_definitions/update")
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	data["sourceDefinitionId"] = id
	data["dockerImageTag"] = dockerImageTag

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.SourceDefinitionFromJSON(res.Body)
}

// ListSourceDefinitions returns all the source definitions the current Airbyte deployment is configured to use
func (c *Client) ListSourceDefinitions(ctx context.Context) ([]types.SourceDefinition, error) {
	u, err := appendToURL(c.endpoint, "/v1/source_definitions/list")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.SourceDefinitionsFromJSON(res.Body)
}

// ListLatestSourceDefinitions returns the latest source definitions the current Airbyte deployment is configured to use
func (c *Client) ListLatestSourceDefinitions(ctx context.Context) ([]types.SourceDefinition, error) {
	u, err := appendToURL(c.endpoint, "/v1/source_definitions/list_latest")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.SourceDefinitionsFromJSON(res.Body)
}

// GetSourceDefinition returns the source definition with the given ID
func (c *Client) GetSourceDefinition(ctx context.Context, id *uuid.UUID) (*types.SourceDefinition, error) {
	u, err := appendToURL(c.endpoint, "/v1/source_definitions/get")
	if err != nil {
		return nil, err
	}

	data := make(map[string]*uuid.UUID)
	data["sourceDefinitionId"] = id

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.SourceDefinitionFromJSON(res.Body)
}

// DeleteSourceDefinition deletes the source definition with the given ID
func (c *Client) DeleteSourceDefinition(ctx context.Context, id *uuid.UUID) error {
	u, err := appendToURL(c.endpoint, "/v1/source_definitions/delete")
	if err != nil {
		return err
	}

	data := make(map[string]*uuid.UUID)
	data["sourceDefinitionId"] = id

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// GetSourceDefinitionSpecification returns the source definition specification with the given source definition ID
func (c *Client) GetSourceDefinitionSpecification(ctx context.Context, id *uuid.UUID) (*types.SourceDefinitionSpecification, error) {
	u, err := appendToURL(c.endpoint, "/v1/source_definition_specifications/get")
	if err != nil {
		return nil, err
	}

	data := make(map[string]*uuid.UUID)
	data["sourceDefinitionId"] = id

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.SourceDefinitionSpecificationFromJSON(res.Body)
}
