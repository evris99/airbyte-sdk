package airbytesdk

import (
	"context"

	"github.com/evris99/airbyte-sdk/types"
	"github.com/google/uuid"
)

// CreateDestination creates a new destination
func (c *Client) CreateDestination(ctx context.Context, dest *types.Destination) (*types.Destination, error) {
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

// UpdateDestination updates a destination
func (c *Client) UpdateDestination(ctx context.Context, dest *types.Destination) (*types.Destination, error) {
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

// ListWorkspaceDestinations returns all the destinations in the workspace with the given ID
func (c *Client) ListWorkspaceDestinations(ctx context.Context, workspaceID *uuid.UUID) ([]types.Destination, error) {
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

// GetDestination returns the destination with the given ID
func (c *Client) GetDestination(ctx context.Context, id *uuid.UUID) (*types.Destination, error) {
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

// SearchDestination searches for the given destination
func (c *Client) SearchDestination(ctx context.Context, dest *types.Destination) (*types.Destination, error) {
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

// CloneDestination makes a copy of the destination with the given ID
func (c *Client) CloneDestination(ctx context.Context, id *uuid.UUID) (*types.Destination, error) {
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

// DeleteDestination deletes a destination with the given ID
func (c *Client) DeleteDestination(ctx context.Context, id *uuid.UUID) error {
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

// CheckDestinationConnection checks the connection to the destination with the given ID
func (c *Client) CheckDestinationConnection(ctx context.Context, id *uuid.UUID) (*types.ConnectionCheck, error) {
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

// CheckDestinationConnectionUpdate checks the connection to the destination with the given ID for updates
func (c *Client) CheckDestinationConnectionUpdate(ctx context.Context, dest *types.Destination) (*types.ConnectionCheck, error) {
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
