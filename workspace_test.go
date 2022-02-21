package airbytesdk

import "testing"

func TestCreateWorkSpace(t *testing.T) {
	airbyte, err := New("http://localhost:8000/api")
	if err != nil {
		t.Fatalf("could not create instance: %v", err)
	}

	workspace := &Workspace{
		Name:                    "test",
		Email:                   "cptevris@gmail.com",
		AnonymousDataCollection: false,
	}

	new, err := airbyte.CreateWorkspace(workspace)
	if err != nil {
		t.Fatalf("could not create workspace: %v", err)
	}

	if new.WorkspaceId.String() == "" {
		t.Fatal("empty new workspace ID")
	}
}
