package airbytesdk

import (
	"context"
	"fmt"

	"github.com/evris99/airbyte-sdk/types"
	"github.com/google/uuid"
)

// CreateWorkspace creates and returns a new workspace
func (c *Client) CreateWorkspace(ctx context.Context, workspace *types.Workspace) (*types.Workspace, error) {
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

// DeleteWorkspace deletes the workspace with the given UUID
func (c *Client) DeleteWorkspace(ctx context.Context, id *uuid.UUID) error {
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

// ListWorkspaces returns all the workspaces
func (c *Client) ListWorkspaces(ctx context.Context) ([]types.Workspace, error) {
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

// FindWorkspaceByID returns the workspace with the given ID
func (c *Client) FindWorkspaceByID(ctx context.Context, id *uuid.UUID) (*types.Workspace, error) {
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

// FindWorkspaceBySlug returns the workspace with the given slug
func (c *Client) FindWorkspaceBySlug(ctx context.Context, slug string) (*types.Workspace, error) {
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

// UpdateWorkspaceState updates the workspace. The WorkspaceId field must be included and the Name field must be empty.
// The whole object must be passed in, even the fields that did not change
func (c *Client) UpdateWorkspaceState(ctx context.Context, workspace types.Workspace) (*types.Workspace, error) {
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

// UpdateWorkspaceName updates the name of workspace with the given id
func (c *Client) UpdateWorkspaceName(ctx context.Context, id *uuid.UUID, name string) (*types.Workspace, error) {
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

// UpdateWorkspaceFeedbackState tags the feedback status of the workspace as done
func (c *Client) UpdateWorkspaceFeedbackState(ctx context.Context, id *uuid.UUID) error {
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
