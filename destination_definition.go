package airbytesdk

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type DestinationDefinition struct {
	Definition
	DestinationDefinitionId *uuid.UUID `json:"destinationDefinitionId,omitempty"`
}

type SupportedDestinationSyncModesType int

const (
	Append SupportedDestinationSyncModesType = iota
	Overwrite
	AppendDedup
)

// Unmarshaler for json
func (sup *SupportedDestinationSyncModesType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch strings.ToLower(s) {
	case "append":
		*sup = Append
	case "overwrite":
		*sup = Overwrite
	case "append_dedup":
		*sup = AppendDedup
	default:
		return fmt.Errorf("unknown auth flow type")
	}

	return nil
}

// Marshaler for json
func (sup SupportedDestinationSyncModesType) MarshalJSON() ([]byte, error) {
	var s string
	switch sup {
	case Append:
		s = "append"
	case Overwrite:
		s = "overwrite"
	case AppendDedup:
		s = "append_dedup"
	default:
		return nil, fmt.Errorf("unknown auth flow type")
	}

	return json.Marshal(s)
}

type DestinationDefinitionSpecification struct {
	DefinitionSpecification
	DestinationDefinitionId       *uuid.UUID                        `json:"destinationDefinitionId"`
	SupportedDestinationSyncModes SupportedDestinationSyncModesType `json:"supportedDestinationSyncModes"`
	SupportsDbt                   bool                              `json:"supportsDbt"`
	SupportsNormalization         bool                              `json:"supportsNormalization"`
}

// Creates and returns a new destination definition using the given context
func (c *Client) CreateDestinationDefinitionWithContext(ctx context.Context, definition *DestinationDefinition) (*DestinationDefinition, error) {
	u, err := appendToURL(c.endpoint, "/v1/destination_definitions/create")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, definition)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode JSON
	newDefinition := new(DestinationDefinition)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(newDefinition); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return newDefinition, nil
}

// Creates and returns a new destination definition.
// Equivalent with calling CreateDestinationDefinitionWithContext with background as context
func (c *Client) CreateDestinationDefinition(definition *DestinationDefinition) (*DestinationDefinition, error) {
	return c.CreateDestinationDefinitionWithContext(context.Background(), definition)
}

// Updates a destination definition. Currently, the only allowed attribute to update is the default docker image version.
func (c *Client) UpdateDestinationDefinitionDockerImageWithContext(ctx context.Context, id *uuid.UUID, dockerImageTag string) (*DestinationDefinition, error) {
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

	// Decode JSON
	updatedDefinition := new(DestinationDefinition)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(updatedDefinition); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return updatedDefinition, nil
}

// Updates a destination definition. Currently, the only allowed attribute to update is the default docker image version.
// Equivalent with calling UpdateDestinationDefinitionDockerImageWithContext with background as context
func (c *Client) UpdateDestinationDefinitionDockerImage(id *uuid.UUID, dockerImageTag string) (*DestinationDefinition, error) {
	return c.UpdateDestinationDefinitionDockerImageWithContext(context.Background(), id, dockerImageTag)
}

// Returns all the destination definitions the current Airbyte deployment is configured to use using the given context
func (c *Client) ListDestinationDefinitionsWithContext(ctx context.Context) ([]DestinationDefinition, error) {
	u, err := appendToURL(c.endpoint, "/v1/destination_definitions/list")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// This is needed because the response list is contained in a destinationDefinitions object
	var destinationDefinitions struct {
		DestinationDefinitions []DestinationDefinition `json:"destinationDefinitions"`
	}

	// Decode JSON
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&destinationDefinitions); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return destinationDefinitions.DestinationDefinitions, nil
}

// Returns all the destination definitions the current Airbyte deployment is configured to use.
// Equivalent with calling ListDestinationDefinitionsWithContext with background as context
func (c *Client) ListDestinationDefinitions() ([]DestinationDefinition, error) {
	return c.ListDestinationDefinitionsWithContext(context.Background())
}

// Returns the latest destination definitions the current Airbyte deployment is configured to use using the given context
func (c *Client) ListLatestDestinationDefinitionsWithContext(ctx context.Context) ([]DestinationDefinition, error) {
	u, err := appendToURL(c.endpoint, "/v1/destination_definitions/list_latest")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// This is needed because the response list is contained in a sourceDefinitions object
	var destinationDefinitions struct {
		DestinationDefinitions []DestinationDefinition `json:"destinationDefinitions"`
	}

	// Decode JSON
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&destinationDefinitions); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return destinationDefinitions.DestinationDefinitions, nil
}

// Returns the latest destination definitions the current Airbyte deployment is configured to use.
// Equivalent with calling ListLatestDestinationDefinitionsWithContext with background as context
func (c *Client) ListLatestDestinationDefinitions() ([]DestinationDefinition, error) {
	return c.ListLatestDestinationDefinitionsWithContext(context.Background())
}

// Returns the destination definition with the given ID using the given context
func (c *Client) GetDestinationDefinitionWithContext(ctx context.Context, id *uuid.UUID) (*DestinationDefinition, error) {
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

	// Decode JSON
	definition := new(DestinationDefinition)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(definition); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return definition, nil
}

// Returns the destination definition with the given ID.
// Equivalent with calling GetDestinationDefinitionsWithContext with background as context
func (c *Client) GetDestinationDefinition(id *uuid.UUID) (*DestinationDefinition, error) {
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
func (c *Client) GetDestinationDefinitionSpecificationWithContext(ctx context.Context, id *uuid.UUID) (*DestinationDefinitionSpecification, error) {
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

	// Decode JSON
	definitionSpecification := new(DestinationDefinitionSpecification)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(definitionSpecification); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return definitionSpecification, nil
}

// Returns the destination definition specification.
// Equivalent with calling GetDestinationDefinitionSpecificationWithContext with background as context
func (c *Client) GetDestinationDefinitionSpecification(id *uuid.UUID) (*DestinationDefinitionSpecification, error) {
	return c.GetDestinationDefinitionSpecificationWithContext(context.Background(), id)
}
