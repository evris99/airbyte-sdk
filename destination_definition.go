package airbytesdk

import (
	"context"

	"github.com/evris99/airbyte-sdk/types"
	"github.com/google/uuid"
)

// CreateDestinationDefinition creates and returns a new destination definition
func (c *Client) CreateDestinationDefinition(ctx context.Context, definition *types.DestinationDefinition) (*types.DestinationDefinition, error) {
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

// UpdateDestinationDefinitionDockerImage updates a destination definition.
// Currently, the only allowed attribute to update is the default docker image version.
func (c *Client) UpdateDestinationDefinitionDockerImage(ctx context.Context, id *uuid.UUID, dockerImageTag string) (*types.DestinationDefinition, error) {
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

// ListDestinationDefinitions returns all the destination definitions the current Airbyte deployment is configured to use
func (c *Client) ListDestinationDefinitions(ctx context.Context) ([]types.DestinationDefinition, error) {
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

// ListLatestDestinationDefinitions returns the latest destination definitions the current Airbyte deployment is configured to use
func (c *Client) ListLatestDestinationDefinitions(ctx context.Context) ([]types.DestinationDefinition, error) {
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

// GetDestinationDefinition returns the destination definition with the given ID
func (c *Client) GetDestinationDefinition(ctx context.Context, id *uuid.UUID) (*types.DestinationDefinition, error) {
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

// DeleteDestinationDefinition deletes the destination definition with the given ID
func (c *Client) DeleteDestinationDefinition(ctx context.Context, id *uuid.UUID) error {
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

// GetDestinationDefinitionSpecification returns the source definition specification
func (c *Client) GetDestinationDefinitionSpecification(ctx context.Context, id *uuid.UUID) (*types.DestinationDefinitionSpecification, error) {
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
