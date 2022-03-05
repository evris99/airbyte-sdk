package airbytesdk

import (
	"context"

	"github.com/evris99/airbyte-sdk/types"
	"github.com/google/uuid"
)

// Creates new source definition using the given context
func (c *Client) CreateSourceDefinitionWithContext(ctx context.Context, definition *types.SourceDefinition) (*types.SourceDefinition, error) {
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

// Creates new source definition.
// Equivalent with calling CreateSourceDefinitionWithContext with background as context
func (c *Client) CreateSourceDefinition(definition *types.SourceDefinition) (*types.SourceDefinition, error) {
	return c.CreateSourceDefinitionWithContext(context.Background(), definition)
}

// Updates a source definition. Currently, the only allowed attribute to update is the default docker image version.
func (c *Client) UpdateSourceDefinitionDockerImageWithContext(ctx context.Context, id *uuid.UUID, dockerImageTag string) (*types.SourceDefinition, error) {
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

// Updates a source definition. Currently, the only allowed attribute to update is the default docker image version.
// Equivalent with calling UpdateSourceDefinitionDockerImageWithContext with background as context
func (c *Client) UpdateSourceDefinitionDockerImage(id *uuid.UUID, dockerImageTag string) (*types.SourceDefinition, error) {
	return c.UpdateSourceDefinitionDockerImageWithContext(context.Background(), id, dockerImageTag)
}

// Returns all the source definitions the current Airbyte deployment is configured to use using the given context
func (c *Client) ListSourceDefinitionsWithContext(ctx context.Context) ([]types.SourceDefinition, error) {
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

// Returns all the source definitions the current Airbyte deployment is configured to use.
// Equivalent with calling ListSourceDefinitionsWithContext with background as context
func (c *Client) ListSourceDefinitions() ([]types.SourceDefinition, error) {
	return c.ListSourceDefinitionsWithContext(context.Background())
}

// Returns the latest source definitions the current Airbyte deployment is configured to use using the given context
func (c *Client) ListLatestSourceDefinitionsWithContext(ctx context.Context) ([]types.SourceDefinition, error) {
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

// Returns the latest source definitions the current Airbyte deployment is configured to use.
// Equivalent with calling ListLatestSourceDefinitionsWithContext with background as context
func (c *Client) ListLatestSourceDefinitions() ([]types.SourceDefinition, error) {
	return c.ListLatestSourceDefinitionsWithContext(context.Background())
}

// Returns the source definition with the given ID using the given context
func (c *Client) GetSourceDefinitionWithContext(ctx context.Context, id *uuid.UUID) (*types.SourceDefinition, error) {
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

// Returns the source definition with the given ID.
// Equivalent with calling GetSourceDefinitionsWithContext with background as context
func (c *Client) GetSourceDefinition(id *uuid.UUID) (*types.SourceDefinition, error) {
	return c.GetSourceDefinitionWithContext(context.Background(), id)
}

// Deletes the source definition with the given ID using the given context
func (c *Client) DeleteSourceDefinitionWithContext(ctx context.Context, id *uuid.UUID) error {
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

// Deletes the source definition with the given ID.
// Equivalent with calling DeleteSourceDefinitionsWithContext with background as context
func (c *Client) DeleteSourceDefinition(id *uuid.UUID) error {
	return c.DeleteSourceDefinitionWithContext(context.Background(), id)
}

// Returns the source definition specification using the given context
func (c *Client) GetSourceDefinitionSpecificationWithContext(ctx context.Context, id *uuid.UUID) (*types.SourceDefinitionSpecification, error) {
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

// Returns the source definition specification.
// Equivalent with calling GetSourceDefinitionSpecificationWithContext with background as context
func (c *Client) GetSourceDefinitionSpecification(id *uuid.UUID) (*types.SourceDefinitionSpecification, error) {
	return c.GetSourceDefinitionSpecificationWithContext(context.Background(), id)
}
