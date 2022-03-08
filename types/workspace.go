package types

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/google/uuid"
)

// All the possible notification types
type NotificationType int

const (
	Slack NotificationType = iota + 1
)

// Unmarshaler for json
func (n *NotificationType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch strings.ToLower(s) {
	case "slack":
		*n = Slack
	}

	return nil
}

// Marshaler for json
func (n NotificationType) MarshalJSON() ([]byte, error) {
	var s string
	switch n {
	case Slack:
		s = "slack"
	}

	return json.Marshal(s)
}

// Configuration options for slack notification
type SlackConfiguration struct {
	Webhook string `json:"webhook"`
}

// Options for notification
type Notification struct {
	NotificationType   NotificationType    `json:"notificationType,omitempty"`
	SendOnSuccess      bool                `json:"sendOnSuccess,omitempty"`
	SendOnFailure      bool                `json:"sendOnFailure,omitempty"`
	SlackConfiguration *SlackConfiguration `json:"slackConfiguration,omitempty"`
}

// A struct containing workspace related resources
type Workspace struct {
	Name                    string         `json:"name,omitempty"`
	WorkspaceId             *uuid.UUID     `json:"workspaceId,omitempty"`
	CustomerId              *uuid.UUID     `json:"customerId,omitempty"`
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

// WorkspaceFromJSON reads json data from a Reader and returns a workspace
func WorkspaceFromJSON(r io.Reader) (*Workspace, error) {
	workspace := new(Workspace)
	err := json.NewDecoder(r).Decode(workspace)

	return workspace, err
}

// WorkspacesFromJSON reads json data from a Reader and returns a slice of workspaces
func WorkspacesFromJSON(r io.Reader) ([]Workspace, error) {
	var workspaces struct {
		Workspaces []Workspace `json:"workspaces"`
	}

	// Decode JSON
	err := json.NewDecoder(r).Decode(&workspaces)
	return workspaces.Workspaces, err
}
