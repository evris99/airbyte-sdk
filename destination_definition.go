package airbytesdk

import (
	"context"

	"github.com/evris99/airbyte-sdk/types"
	"github.com/google/uuid"
)

// Creates and returns a new destination definition using the given context
func (c *Client) CreateDestinationDefinitionWithContext(ctx context.Context, definition *types.DestinationDefinition) (*types.DestinationDefinition, error) {
	u, err := appendToURL(c.endpoint, "/v1/destination_definitions/create")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, definition)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.DestinationDefinitionFromJSON(res.Body)
}

// Creates and returns a new destination definition.
// Equivalent with calling CreateDestinationDefinitionWithContext with background as context
func (c *Client) CreateDestinationDefinition(definition *types.DestinationDefinition) (*types.DestinationDefinition, error) {
	return c.CreateDestinationDefinitionWithContext(context.Background(), definition)
}

// Updates a destination definition. Currently, the only allowed attribute to update is the default docker image version.
func (c *Client) UpdateDestinationDefinitionDockerImageWithContext(ctx context.Context, id *uuid.UUID, dockerImageTag string) (*types.DestinationDefinition, error) {
	u, err := appendToURL(c.endpoint, "/v1/destination_definitions/update")
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	data["destinationDefinitionId"] = id
	data["dockerImageTag"] = dockerImageTag

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.DestinationDefinitionFromJSON(res.Body)
}

// Updates a destination definition. Currently, the only allowed attribute to update is the default docker image version.
// Equivalent with calling UpdateDestinationDefinitionDockerImageWithContext with background as context
func (c *Client) UpdateDestinationDefinitionDockerImage(id *uuid.UUID, dockerImageTag string) (*types.DestinationDefinition, error) {
	return c.UpdateDestinationDefinitionDockerImageWithContext(context.Background(), id, dockerImageTag)
}

// Returns all the destination definitions the current Airbyte deployment is configured to use using the given context
func (c *Client) ListDestinationDefinitionsWithContext(ctx context.Context) ([]types.DestinationDefinition, error) {
	u, err := appendToURL(c.endpoint, "/v1/destination_definitions/list")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.DestinationDefinitionsFromJSON(res.Body)
}

// Returns all the destination definitions the current Airbyte deployment is configured to use.
// Equivalent with calling ListDestinationDefinitionsWithContext with background as context
func (c *Client) ListDestinationDefinitions() ([]types.DestinationDefinition, error) {
	return c.ListDestinationDefinitionsWithContext(context.Background())
}

// Returns the latest destination definitions the current Airbyte deployment is configured to use using the given context
func (c *Client) ListLatestDestinationDefinitionsWithContext(ctx context.Context) ([]types.DestinationDefinition, error) {
	u, err := appendToURL(c.endpoint, "/v1/destination_definitions/list_latest")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.DestinationDefinitionsFromJSON(res.Body)
}

// Returns the latest destination definitions the current Airbyte deployment is configured to use.
// Equivalent with calling ListLatestDestinationDefinitionsWithContext with background as context
func (c *Client) ListLatestDestinationDefinitions() ([]types.DestinationDefinition, error) {
	return c.ListLatestDestinationDefinitionsWithContext(context.Background())
}

// Returns the destination definition with the given ID using the given context
func (c *Client) GetDestinationDefinitionWithContext(ctx context.Context, id *uuid.UUID) (*types.DestinationDefinition, error) {
	u, err := appendToURL(c.endpoint, "/v1/destination_definitions/get")
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

	return types.DestinationDefinitionFromJSON(res.Body)
}

// Returns the destination definition with the given ID.
// Equivalent with calling GetDestinationDefinitionsWithContext with background as context
func (c *Client) GetDestinationDefinition(id *uuid.UUID) (*types.DestinationDefinition, error) {
	return c.GetDestinationDefinitionWithContext(context.Background(), id)
}

// Deletes the destination definition with the given ID using the given context
func (c *Client) DeleteDestinationDefinitionWithContext(ctx context.Context, id *uuid.UUID) error {
	u, err := appendToURL(c.endpoint, "/v1/destination_definitions/delete")
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
// Equivalent with calling DeleteDestinationDefinitionsWithContext with background as context
func (c *Client) DeleteDestinationDefinition(id *uuid.UUID) error {
	return c.DeleteDestinationDefinitionWithContext(context.Background(), id)
}

// Returns the source definition specification using the given context
func (c *Client) GetDestinationDefinitionSpecificationWithContext(ctx context.Context, id *uuid.UUID) (*types.DestinationDefinitionSpecification, error) {
	u, err := appendToURL(c.endpoint, "/v1/source_definition_specifications/get")
	if err != nil {
		return nil, err
	}

	data := make(map[string]*uuid.UUID)
	data["destinationDefinitionId"] = id

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.DestinationDefinitionSpecificationToJSON(res.Body)
}

// Returns the destination definition specification.
// Equivalent with calling GetDestinationDefinitionSpecificationWithContext with background as context
func (c *Client) GetDestinationDefinitionSpecification(id *uuid.UUID) (*types.DestinationDefinitionSpecification, error) {
	return c.GetDestinationDefinitionSpecificationWithContext(context.Background(), id)
}
