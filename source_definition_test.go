package airbytesdk

import (
	"testing"

	"github.com/evris99/airbyte-sdk/types"
)

func TestListSourceDefinitions(t *testing.T) {
	airbyte, err := New("http://localhost:8000/api")
	if err != nil {
		t.Fatalf("could not create instance: %v", err)
	}
	sourceDefinition := &types.SourceDefinition{
		Definition: types.Definition{
			Name:             "test",
			DockerRepository: "test",
			DockerImageTag:   "test",
			DocumentationURL: "https://test.com",
		},
	}

	new, err := airbyte.CreateSourceDefinition(sourceDefinition)
	if err != nil {
		t.Fatalf("could not create source definition: %v", err)
	}

	allSourceDefinitions, err := airbyte.ListSourceDefinitions()
	if err != nil {
		t.Fatalf("could not get all source definitions")
	}

	found := false
	for _, sd := range allSourceDefinitions {
		if sd.SourceDefinitionId.String() == new.SourceDefinitionId.String() {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("could not find newly created source definition in list")
	}

	latestSourceDefinitions, err := airbyte.ListLatestSourceDefinitions()
	if err != nil {
		t.Fatalf("could not get latest source definitions")
	}

	found = false
	for _, sd := range latestSourceDefinitions {
		if sd.SourceDefinitionId.String() == new.SourceDefinitionId.String() {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("could not find newly created source definition in latest list")
	}

	if err := airbyte.DeleteSourceDefinition(new.SourceDefinitionId); err != nil {
		t.Fatalf("could not delete source definition: %v", err)
	}
}
