package airbytesdk

import (
	"context"
	"testing"

	"github.com/evris99/airbyte-sdk/types"
)

func TestListWorkspace(t *testing.T) {
	airbyte, err := New("http://localhost:8000/api")
	if err != nil {
		t.Fatalf("could not create instance: %v", err)
	}

	// Create new workspace
	workspace := &types.Workspace{
		Name:                    "test",
		Email:                   "test@gmail.com",
		AnonymousDataCollection: false,
	}

	new, err := airbyte.CreateWorkspace(context.Background(), workspace)
	if err != nil {
		t.Fatalf("could not create workspace: %v", err)
	}

	// Get all workspaces
	workspaces, err := airbyte.ListWorkspaces(context.Background())
	if err != nil {
		t.Fatalf("could not list workspaces: %v", err)
	}

	//Find if workspace is in list
	found := false
	for _, w := range workspaces {
		if w.WorkspaceId.String() == new.WorkspaceId.String() {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("could not find newly created workspace in workspace list")
	}

	// Delete workspace
	if err := airbyte.DeleteWorkspace(context.Background(), new.WorkspaceId); err != nil {
		t.Fatalf("could not delete workspace: %v", err)
	}
}

func TestFindWorkspace(t *testing.T) {
	airbyte, err := New("http://localhost:8000/api")
	if err != nil {
		t.Fatalf("could not create instance: %v", err)
	}

	workspace := &types.Workspace{
		Name:                    "test",
		Email:                   "test@gmail.com",
		AnonymousDataCollection: false,
	}

	// Create new workspace
	new, err := airbyte.CreateWorkspace(context.Background(), workspace)
	if err != nil {
		t.Fatalf("could not create workspace: %v", err)
	}

	// Find new workspace by ID
	idWorkspace, err := airbyte.FindWorkspaceByID(context.Background(), new.WorkspaceId)
	if err != nil {
		t.Fatalf("could not find workspace by ID: %v", err)
	}

	if idWorkspace.Name != new.Name {
		t.Fatal("found workspace name does not match created workspace")
	}

	// Find new workspace by slug
	slugWorkspace, err := airbyte.FindWorkspaceBySlug(context.Background(), new.Slug)
	if err != nil {
		t.Fatalf("could not find workspace by slug: %v", err)
	}

	if slugWorkspace.Name != new.Name {
		t.Fatal("found workspace name does not match created workspace")
	}

	// Delete workspace
	if err := airbyte.DeleteWorkspace(context.Background(), slugWorkspace.WorkspaceId); err != nil {
		t.Fatalf("could not delete workspace: %v", err)
	}
}

func TestUpdateWorkspace(t *testing.T) {
	airbyte, err := New("http://localhost:8000/api")
	if err != nil {
		t.Fatalf("could not create instance: %v", err)
	}

	// Create new workspace
	workspace := &types.Workspace{
		Name:                    "test",
		Email:                   "test@gmail.com",
		AnonymousDataCollection: false,
	}

	new, err := airbyte.CreateWorkspace(context.Background(), workspace)
	if err != nil {
		t.Fatalf("could not create workspace: %v", err)
	}

	// Update workspace email
	update := types.Workspace{
		WorkspaceId:             new.WorkspaceId,
		InitialSetupComplete:    true,
		AnonymousDataCollection: true,
		News:                    true,
		SecurityUpdates:         true,
	}

	updatedWorkspace, err := airbyte.UpdateWorkspaceState(context.Background(), update)

	if err != nil {
		t.Fatalf("could not update workspace state: %v", err)
	}

	if !updatedWorkspace.News {
		t.Fatal("incorrect news setting")
	}

	// Update workspace name
	updatedNameWorkspace, err := airbyte.UpdateWorkspaceName(context.Background(), updatedWorkspace.WorkspaceId, "changed")
	if err != nil {
		t.Fatalf("could not update workspace name: %v", err)
	}

	if updatedNameWorkspace.Name != "changed" {
		t.Fatal("incorrect updated name")
	}

	// Update workspace feedback state
	if err := airbyte.UpdateWorkspaceFeedbackState(context.Background(), updatedNameWorkspace.WorkspaceId); err != nil {
		t.Fatalf("could not update workspace feedback state: %v", err)
	}

	// Delete workspace
	if err := airbyte.DeleteWorkspace(context.Background(), updatedNameWorkspace.WorkspaceId); err != nil {
		t.Fatalf("could not delete workspace: %v", err)
	}
}
