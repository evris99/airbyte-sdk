package airbytesdk

import (
	"context"

	"github.com/evris99/airbyte-sdk/types"
	"github.com/google/uuid"
)

// Create a new destination using the given context
func (c *Client) CreateDestinationWithContext(ctx context.Context, dest *types.Destination) (*types.Destination, error) {
	u, err := appendToURL(c.endpoint, "/v1/destinations/create")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, dest)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.DestinationFromJSON(res.Body)
}

// Create a new destination.
// Equivalent with calling CreateDestinationWithContext with background as context
func (c *Client) CreateDestination(dest *types.Destination) (*types.Destination, error) {
	return c.CreateDestinationWithContext(context.Background(), dest)
}

// Update a destination using the given context
func (c *Client) UpdateDestinationWithContext(ctx context.Context, dest *types.Destination) (*types.Destination, error) {
	u, err := appendToURL(c.endpoint, "/v1/destinations/update")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, dest)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.DestinationFromJSON(res.Body)
}

// Update a destination.
// Equivalent with calling updateDestinationWithContext with background as context
func (c *Client) UpdateDestination(dest *types.Destination) (*types.Destination, error) {
	return c.UpdateDestinationWithContext(context.Background(), dest)
}

// Returns all the destinations in the workspace with the give ID using the given context
func (c *Client) ListWorkspaceDestinationsWithContext(ctx context.Context, workspaceID *uuid.UUID) ([]types.Destination, error) {
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

	return types.DestinationsFromJSON(res.Body)
}

// Returns all the destinations in the workspace with the give ID.
// Equivalent with calling ListWorkspaceDestinationsWithContext with background as context
func (c *Client) ListWorkspaceDestinations(workspaceID *uuid.UUID) ([]types.Destination, error) {
	return c.ListWorkspaceDestinationsWithContext(context.Background(), workspaceID)
}

// Returns a destination with the given ID using the given context
func (c *Client) GetDestinationWithContext(ctx context.Context, id *uuid.UUID) (*types.Destination, error) {
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

	return types.DestinationFromJSON(res.Body)
}

// Returns a destination with the given ID.
// Equivalent with calling GetDestinationWithContext with background as context
func (c *Client) GetDestination(id *uuid.UUID) (*types.Destination, error) {
	return c.GetDestinationWithContext(context.Background(), id)
}

// Searches for the given destination using the given context
func (c *Client) SearchDestinationWithContext(ctx context.Context, dest *types.Destination) (*types.Destination, error) {
	u, err := appendToURL(c.endpoint, "/v1/destinations/search")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, dest)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.DestinationFromJSON(res.Body)
}

// Searches for the given destination.
// Equivalent with calling SearchDestinationWithContext with background as context
func (c *Client) SearchDestination(dest *types.Destination) (*types.Destination, error) {
	return c.SearchDestinationWithContext(context.Background(), dest)
}

// Makes a copy of the destination with the given ID using the given context
func (c *Client) CloneDestinationWithContext(ctx context.Context, id *uuid.UUID) (*types.Destination, error) {
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

	return types.DestinationFromJSON(res.Body)
}

// Makes a copy of the destination with the given ID.
// Equivalent with calling CloneDestinationWithContext with background as context
func (c *Client) CloneDestination(id *uuid.UUID) (*types.Destination, error) {
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
func (c *Client) CheckDestinationConnectionWithContext(ctx context.Context, id *uuid.UUID) (*types.ConnectionCheck, error) {
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

	return types.ConnectionCheckFromJSON(res.Body)
}

// Checks the connection to the destination with the given ID.
// Equivalent with calling CheckDestinationConnectionWithContext with background as context
func (c *Client) CheckDestinationConnection(id *uuid.UUID) (*types.ConnectionCheck, error) {
	return c.CheckDestinationConnectionWithContext(context.Background(), id)
}

// Checks the connection to the destination with the given ID for updates using the given context
func (c *Client) CheckDestinationConnectionUpdateWithContext(ctx context.Context, dest *types.Destination) (*types.ConnectionCheck, error) {
	u, err := appendToURL(c.endpoint, "/v1/destinations/check_connection_for_update")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, dest)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.ConnectionCheckFromJSON(res.Body)
}

// Checks the connection to the destination with the given ID for updates.
// Equivalent with calling CheckDestinationConnectionUpdateWithContext with background as context
func (c *Client) CheckDestinationConnectionUpdate(dest *types.Destination) (*types.ConnectionCheck, error) {
	return c.CheckDestinationConnectionUpdateWithContext(context.Background(), dest)
}
