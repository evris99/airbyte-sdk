package airbytesdk

import (
	"context"
	"testing"

	"github.com/evris99/airbyte-sdk/types"
)

func TestSourceDefinitions(t *testing.T) {
	airbyte, err := New("http://localhost:8000/api")
	if err != nil {
		t.Fatalf("could not create instance: %v", err)
	}

	sourceDefinition := &types.SourceDefinition{
		Definition: types.Definition{
			Name:             "Test",
			DockerRepository: "airbyte/source-facebook-marketing",
			DockerImageTag:   "0.2.37",
			DocumentationURL: "https://test.com",
		},
	}

	new, err := airbyte.CreateSourceDefinition(context.Background(), sourceDefinition)
	if err != nil {
		t.Fatalf("could not create source definition: %v", err)
	}

	_, err = airbyte.ListSourceDefinitions(context.Background())
	if err != nil {
		t.Fatalf("could not get all source definitions: %v", err)
	}

	_, err = airbyte.ListLatestSourceDefinitions(context.Background())
	if err != nil {
		t.Fatalf("could not get latest source definitions: %v", err)
	}

	found, err := airbyte.GetSourceDefinition(context.Background(), new.SourceDefinitionId)
	if err != nil {
		t.Fatalf("could not get newly created source definition: %v", err)
	}

	if found.SourceDefinitionId.String() != new.SourceDefinitionId.String() {
		t.Fatal("got incorrect ID")
	}

	if err := airbyte.DeleteSourceDefinition(context.Background(), found.SourceDefinitionId); err != nil {
		t.Fatalf("could not delete source definition: %v", err)
	}
}
