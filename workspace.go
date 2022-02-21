package airbytesdk

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

// All the possible notification types
type NotificationTypeEnum int

const (
	Slack NotificationTypeEnum = iota
)

// Unmarshaler for json
func (n *NotificationTypeEnum) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch strings.ToLower(s) {
	case "slack":
		*n = Slack
	default:
		return fmt.Errorf("unknown notification type")
	}

	return nil
}

// Marshaler for json
func (n NotificationTypeEnum) MarshalJSON() ([]byte, error) {
	var s string
	switch n {
	case Slack:
		s = "slack"
	default:
		return nil, fmt.Errorf("unknown notification type")
	}

	return json.Marshal(s)
}

// Configuration options for slack notification
type SlackConfiguration struct {
	Webhook string `json:"webhook"`
}

// Options for notification
type Notification struct {
	NotificationType   NotificationTypeEnum `json:"notificationType"`
	SendOnSuccess      bool                 `json:"sendOnSuccess"`
	SendOnFailure      bool                 `json:"sendOnFailure"`
	SlackConfiguration SlackConfiguration   `json:"slackConfiguration,omitempty"`
}

// A struct containing workspace related resources
type Workspace struct {
	Name                    string         `json:"name,omitempty"`
	WorkspaceId             uuid.UUID      `json:"workspaceId,omitempty"`
	CustomerId              uuid.UUID      `json:"customerId,omitempty"`
	Email                   string         `json:"email,omitempty"`
	Slug                    string         `json:"slug,omitempty"`
	AnonymousDataCollection bool           `json:"anonymousDataCollection,omitempty"`
	News                    bool           `json:"news,omitempty"`
	SecurityUpdates         bool           `json:"securityUpdates,omitempty"`
	Notifications           []Notification `json:"notifications,omitempty"`
	DisplaySetupWizard      bool           `json:"displaySetupWizard,omitempty"`
	InitialSetupComplete    bool           `json:"initialSetupComplete,omitempty"`
	FirstCompletedSync      bool           `json:"firstCompletedSync,omitempty"`
	FeedbackDone            bool           `json:"feedbackDone,omitempty"`
}

// Creates and returns a new workspace using the given context
func (c *Client) CreateWorkspaceWithContext(ctx context.Context, workspace *Workspace) (*Workspace, error) {
	// TODO: Check if name exists
	u, err := c.endpoint.Parse(fmt.Sprintf("%s/v1/workspaces/create", c.endpoint.Path))
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, workspace)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	newWorkspace := new(Workspace)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(newWorkspace); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return newWorkspace, nil
}

// Creates and returns a new workspace.
// Equivalent with calling CreateWorkspaceWithContext with background as context
func (c *Client) CreateWorkspace(workspace *Workspace) (*Workspace, error) {
	return c.CreateWorkspaceWithContext(context.Background(), workspace)
}

// Deletes the workspace with the given UUID using the given context
func (c *Client) DeleteWorkspaceWithContext(ctx context.Context, id uuid.UUID) error {
	u, err := c.endpoint.Parse("/v1/workspaces/delete")
	if err != nil {
		return err
	}

	data := make(map[string]uuid.UUID)
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
func (c *Client) DeleteWorkspace(id uuid.UUID) error {
	return c.DeleteWorkspaceWithContext(context.Background(), id)
}

// Returns all the workspaces using the given context
func (c *Client) ListWorkspacesWithContext(ctx context.Context) ([]Workspace, error) {
	u, err := c.endpoint.Parse(fmt.Sprintf("%s/v1/workspaces/list", c.endpoint.Path))
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var workspaces []Workspace
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&workspaces); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return workspaces, nil
}

// Returns all the workspaces.
// Equivalent with calling ListWorkspacesWithContext with background as context
func (c *Client) ListWorkspaces() ([]Workspace, error) {
	return c.ListWorkspacesWithContext(context.Background())
}

// Returns the workspace with the given ID using the given context
func (c *Client) FindWorkspaceByIDWithContext(ctx context.Context, id uuid.UUID) (*Workspace, error) {
	u, err := c.endpoint.Parse(fmt.Sprintf("%s/v1/workspaces/get", c.endpoint.Path))
	if err != nil {
		return nil, err
	}

	data := make(map[string]uuid.UUID)
	data["workspaceId"] = id

	res, err := c.makeRequest(ctx, u, data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	workspace := new(Workspace)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(workspace); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return workspace, nil
}

// Returns the workspace with the given ID.
// Equivalent with calling FindWorkspaceByIDWithContext with background as context
func (c *Client) FindWorkspaceByID(id uuid.UUID) (*Workspace, error) {
	return c.FindWorkspaceByIDWithContext(context.Background(), id)
}

// Returns the workspace with the given slug using the given context
func (c *Client) FindWorkspaceBySlugWithContext(ctx context.Context, slug string) (*Workspace, error) {
	u, err := c.endpoint.Parse(fmt.Sprintf("%s/v1/workspaces/get_by_slug", c.endpoint.Path))
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

	workspace := new(Workspace)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(workspace); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return workspace, nil
}

// Returns the workspace with the given slug.
// Equivalent with calling FindWorkspaceBySlugWithContext with background as context
func (c *Client) FindWorkspaceBySlug(slug string) (*Workspace, error) {
	return c.FindWorkspaceBySlugWithContext(context.Background(), slug)
}

// Updates the workspace using the given context. The WorkspaceId field must be included.
// The whole object must be passed in, even the fields that did not change
func (c *Client) UpdateWorkspaceStateWithContext(ctx context.Context, workspace *Workspace) (*Workspace, error) {
	if workspace.WorkspaceId.String() == "" {
		return nil, fmt.Errorf("the workspaceId must be set")
	}

	u, err := c.endpoint.Parse(fmt.Sprintf("%s/v1/workspaces/update", c.endpoint.Path))
	if err != nil {
		return nil, err
	}

	res, err := c.makeRequest(ctx, u, workspace)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	newWorkspace := new(Workspace)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(newWorkspace); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return newWorkspace, nil
}

// Updates the workspace. The WorkspaceId field must be included.
// The whole object must be passed in, even the fields that did not change.
// Equivalent with calling UpdateWorkspaceStateWithContext with background as context
func (c *Client) UpdateWorkspaceState(workspace *Workspace) (*Workspace, error) {
	return c.UpdateWorkspaceStateWithContext(context.Background(), workspace)
}

// Updates the name of workspace with the given id using the given context
func (c *Client) UpdateWorkspaceNameWithContext(ctx context.Context, id uuid.UUID, name string) (*Workspace, error) {
	u, err := c.endpoint.Parse(fmt.Sprintf("%s/v1/workspaces/update_name", c.endpoint.Path))
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

	newWorkspace := new(Workspace)
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(newWorkspace); err != nil {
		return nil, fmt.Errorf("could not decode response: %w", err)
	}

	return newWorkspace, nil
}

// Updates the name of workspace with the given id.
// Equivalent with calling UpdateWorkspaceNameWithContext with background as context
func (c *Client) UpdateWorkspaceName(id uuid.UUID, name string) (*Workspace, error) {
	return c.UpdateWorkspaceNameWithContext(context.Background(), id, name)
}

// Tags the feedback status of the workspace as done using the given context
func (c *Client) UpdateWorkspaceFeedbackStateWithContext(ctx context.Context, id uuid.UUID) error {
	u, err := c.endpoint.Parse(fmt.Sprintf("%s/v1/workspaces/tag_feedback_status_as_done", c.endpoint.Path))
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
func (c *Client) UpdateWorkspaceFeedbackState(id uuid.UUID) error {
	return c.UpdateWorkspaceFeedbackStateWithContext(context.Background(), id)
}
