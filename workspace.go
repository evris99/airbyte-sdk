package airbytesdk

import (
	"context"
	"fmt"

	"github.com/evris99/airbyte-sdk/types"
	"github.com/google/uuid"
)

// Creates and returns a new workspace using the given context
func (c *Client) CreateWorkspaceWithContext(ctx context.Context, workspace *types.Workspace) (*types.Workspace, error) {
	u, err := appendToURL(c.endpoint, "/v1/workspaces/create")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, workspace)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.WorkspaceFromJSON(res.Body)
}

// Creates and returns a new workspace.
// Equivalent with calling CreateWorkspaceWithContext with background as context
func (c *Client) CreateWorkspace(workspace *types.Workspace) (*types.Workspace, error) {
	return c.CreateWorkspaceWithContext(context.Background(), workspace)
}

// Deletes the workspace with the given UUID using the given context
func (c *Client) DeleteWorkspaceWithContext(ctx context.Context, id *uuid.UUID) error {
	u, err := appendToURL(c.endpoint, "/v1/workspaces/delete")
	if err != nil {
		return err
	}

	data := make(map[string]*uuid.UUID)
	data["workspaceId"] = id

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// Deletes the workspace with the given UUID.
// Equivalent with calling DeleteWorkspaceWithContext with background as context
func (c *Client) DeleteWorkspace(id *uuid.UUID) error {
	return c.DeleteWorkspaceWithContext(context.Background(), id)
}

// Returns all the workspaces using the given context
func (c *Client) ListWorkspacesWithContext(ctx context.Context) ([]types.Workspace, error) {
	u, err := appendToURL(c.endpoint, "/v1/workspaces/list")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.WorkspacesFromJSON(res.Body)
}

// Returns all the workspaces.
// Equivalent with calling ListWorkspacesWithContext with background as context
func (c *Client) ListWorkspaces() ([]types.Workspace, error) {
	return c.ListWorkspacesWithContext(context.Background())
}

// Returns the workspace with the given ID using the given context
func (c *Client) FindWorkspaceByIDWithContext(ctx context.Context, id *uuid.UUID) (*types.Workspace, error) {
	u, err := appendToURL(c.endpoint, "/v1/workspaces/get")
	if err != nil {
		return nil, err
	}

	data := make(map[string]*uuid.UUID)
	data["workspaceId"] = id

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.WorkspaceFromJSON(res.Body)
}

// Returns the workspace with the given ID.
// Equivalent with calling FindWorkspaceByIDWithContext with background as context
func (c *Client) FindWorkspaceByID(id *uuid.UUID) (*types.Workspace, error) {
	return c.FindWorkspaceByIDWithContext(context.Background(), id)
}

// Returns the workspace with the given slug using the given context
func (c *Client) FindWorkspaceBySlugWithContext(ctx context.Context, slug string) (*types.Workspace, error) {
	u, err := appendToURL(c.endpoint, "/v1/workspaces/get_by_slug")
	if err != nil {
		return nil, err
	}

	data := make(map[string]string)
	data["slug"] = slug

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.WorkspaceFromJSON(res.Body)
}

// Returns the workspace with the given slug.
// Equivalent with calling FindWorkspaceBySlugWithContext with background as context
func (c *Client) FindWorkspaceBySlug(slug string) (*types.Workspace, error) {
	return c.FindWorkspaceBySlugWithContext(context.Background(), slug)
}

// Updates the workspace using the given context. The WorkspaceId field must be included and the Name field must be empty.
// The whole object must be passed in, even the fields that did not change
func (c *Client) UpdateWorkspaceStateWithContext(ctx context.Context, workspace types.Workspace) (*types.Workspace, error) {
	if workspace.WorkspaceId.String() == "" {
		return nil, fmt.Errorf("the workspaceId must be set")
	}

	// Omited fields
	workspace.Name = ""
	workspace.Slug = ""
	workspace.CustomerId = nil

	u, err := appendToURL(c.endpoint, "/v1/workspaces/update")
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, workspace)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.WorkspaceFromJSON(res.Body)
}

// Updates the workspace. The WorkspaceId field must be included.
// The whole object must be passed in, even the fields that did not change.
// Equivalent with calling UpdateWorkspaceStateWithContext with background as context
func (c *Client) UpdateWorkspaceState(workspace types.Workspace) (*types.Workspace, error) {
	return c.UpdateWorkspaceStateWithContext(context.Background(), workspace)
}

// Updates the name of workspace with the given id using the given context
func (c *Client) UpdateWorkspaceNameWithContext(ctx context.Context, id *uuid.UUID, name string) (*types.Workspace, error) {
	u, err := appendToURL(c.endpoint, "/v1/workspaces/update_name")
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	data["workspaceId"] = id
	data["name"] = name

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return types.WorkspaceFromJSON(res.Body)
}

// Updates the name of workspace with the given id.
// Equivalent with calling UpdateWorkspaceNameWithContext with background as context
func (c *Client) UpdateWorkspaceName(id *uuid.UUID, name string) (*types.Workspace, error) {
	return c.UpdateWorkspaceNameWithContext(context.Background(), id, name)
}

// Tags the feedback status of the workspace as done using the given context
func (c *Client) UpdateWorkspaceFeedbackStateWithContext(ctx context.Context, id *uuid.UUID) error {
	u, err := appendToURL(c.endpoint, "/v1/workspaces/tag_feedback_status_as_done")
	if err != nil {
		return err
	}

	data := make(map[string]interface{})
	data["workspaceId"] = id

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// Tags the feedback status of the workspace as done.
// Equivalent with calling UpdateWorkspaceFeedbackStateWithContext with background as context
func (c *Client) UpdateWorkspaceFeedbackState(id *uuid.UUID) error {
	return c.UpdateWorkspaceFeedbackStateWithContext(context.Background(), id)
}
