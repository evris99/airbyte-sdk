package airbytesdk

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type Destination struct {
	DestinationId           *uuid.UUID  `json:"destinationId"`
	DestinationDefinitionId *uuid.UUID  `json:"destinationDefinitionId"`
	WorkspaceId             *uuid.UUID  `json:"workspaceId"`
	ConnectionConfiguration interface{} `json:"connectionConfiguration"`
	Name                    string      `json:"name"`
	DestinationName         string      `json:"destinationName"`
}

// Create a new destination using the given context
func (c *Client) CreateDestinationWithContext(ctx context.Context, dest *Destination) (*Destination, error) {
	u, err := appendToURL(c.endpoint, "/v1/destinations/create")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, dest)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode JSON
	newDest := new(Destination)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(newDest); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return newDest, nil
}

// Create a new destination.
// Equivalent with calling CreateDestinationWithContext with background as context
func (c *Client) CreateDestination(dest *Destination) (*Destination, error) {
	return c.CreateDestinationWithContext(context.Background(), dest)
}

// Update a destination using the given context
func (c *Client) UpdateDestinationWithContext(ctx context.Context, dest *Destination) (*Destination, error) {
	u, err := appendToURL(c.endpoint, "/v1/destinations/update")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, dest)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode JSON
	updatedDestination := new(Destination)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(updatedDestination); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return updatedDestination, nil
}

// Update a destination.
// Equivalent with calling updateDestinationWithContext with background as context
func (c *Client) UpdateDestination(dest *Destination) (*Destination, error) {
	return c.UpdateDestinationWithContext(context.Background(), dest)
}

// Returns all the destinations in the workspace with the give ID using the given context
func (c *Client) ListWorkspaceDestinationsWithContext(ctx context.Context, workspaceID *uuid.UUID) ([]Destination, error) {
	u, err := appendToURL(c.endpoint, "/v1/destinations/list")
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

	var destinations struct {
		Destinations []Destination `json:"destinations"`
	}

	// Decode JSON
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&destinations); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return destinations.Destinations, nil
}

// Returns all the destinations in the workspace with the give ID.
// Equivalent with calling ListWorkspaceDestinationsWithContext with background as context
func (c *Client) ListWorkspaceDestinations(workspaceID *uuid.UUID) ([]Destination, error) {
	return c.ListWorkspaceDestinationsWithContext(context.Background(), workspaceID)
}

// Returns a destination with the given ID using the given context
func (c *Client) GetDestinationWithContext(ctx context.Context, id *uuid.UUID) (*Destination, error) {
	u, err := appendToURL(c.endpoint, "/v1/destinations/get")
	if err != nil {
		return nil, err
	}

	data := make(map[string]*uuid.UUID)
	data["destinationId"] = id

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode JSON
	dest := new(Destination)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(dest); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return dest, nil
}

// Returns a destination with the given ID.
// Equivalent with calling GetDestinationWithContext with background as context
func (c *Client) GetDestination(id *uuid.UUID) (*Destination, error) {
	return c.GetDestinationWithContext(context.Background(), id)
}

// Searches for the given destination using the given context
func (c *Client) SearchDestinationWithContext(ctx context.Context, dest *Destination) (*Destination, error) {
	u, err := appendToURL(c.endpoint, "/v1/destinations/search")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, dest)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	//Decode JSON
	foundDest := new(Destination)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(foundDest); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return foundDest, nil
}

// Searches for the given destination.
// Equivalent with calling SearchDestinationWithContext with background as context
func (c *Client) SearchDestination(dest *Destination) (*Destination, error) {
	return c.SearchDestinationWithContext(context.Background(), dest)
}

// Makes a copy of the destination with the given ID using the given context
func (c *Client) CloneDestinationWithContext(ctx context.Context, id *uuid.UUID) (*Destination, error) {
	u, err := appendToURL(c.endpoint, "/v1/destinations/clone")
	if err != nil {
		return nil, err
	}

	data := make(map[string]*uuid.UUID)
	data["destinationId"] = id

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode JSON
	dest := new(Destination)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(dest); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return dest, nil
}

// Makes a copy of the destination with the given ID.
// Equivalent with calling CloneDestinationWithContext with background as context
func (c *Client) CloneDestination(id *uuid.UUID) (*Destination, error) {
	return c.CloneDestinationWithContext(context.Background(), id)
}

// Deletes a destination with the given ID using the given context
func (c *Client) DeleteDestinationWithContext(ctx context.Context, id *uuid.UUID) error {
	u, err := appendToURL(c.endpoint, "/v1/destinations/delete")
	if err != nil {
		return err
	}

	data := make(map[string]*uuid.UUID)
	data["destinationId"] = id

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// Deletes a destination with the given ID.
// Equivalent with calling DeleteDestinationWithContext with background as context
func (c *Client) DeleteDestination(id *uuid.UUID) error {
	return c.DeleteSourceWithContext(context.Background(), id)
}

// Checks the connection to the destination with the given ID using the given context
func (c *Client) CheckDestinationConnectionWithContext(ctx context.Context, id *uuid.UUID) (*Connection, error) {
	u, err := appendToURL(c.endpoint, "/v1/destinations/check_connection")
	if err != nil {
		return nil, err
	}

	data := make(map[string]*uuid.UUID)
	data["destinationId"] = id

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode JSON
	connection := new(Connection)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(connection); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return connection, nil
}

// Checks the connection to the destination with the given ID.
// Equivalent with calling CheckDestinationConnectionWithContext with background as context
func (c *Client) CheckDestinationConnection(id *uuid.UUID) (*Connection, error) {
	return c.CheckDestinationConnectionWithContext(context.Background(), id)
}

// Checks the connection to the destination with the given ID for updates using the given context
func (c *Client) CheckDestinationConnectionUpdateWithContext(ctx context.Context, dest *Destination) (*Connection, error) {
	u, err := appendToURL(c.endpoint, "/v1/destinations/check_connection_for_update")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, dest)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Decode JSON
	connection := new(Connection)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(connection); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return connection, nil
}

// Checks the connection to the destination with the given ID for updates.
// Equivalent with calling CheckDestinationConnectionUpdateWithContext with background as context
func (c *Client) CheckDestinationConnectionUpdate(dest *Destination) (*Connection, error) {
	return c.CheckDestinationConnectionUpdateWithContext(context.Background(), dest)
}
