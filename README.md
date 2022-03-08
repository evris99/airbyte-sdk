# Airbyte Golang SDK

A Golang SDK for interacting with the [Airbyte](https://github.com/airbytehq/airbyte) API. Documentation for the underlying API can be found [here](https://airbyte-public-api-docs.s3.us-east-2.amazonaws.com/rapidoc-api-docs.html).

[![API Reference](
https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667
)](https://pkg.go.dev/github.com/evris99/airbyte-sdk)

## Installation

Have Go version 1.17.6 or higher and run 
```
go get -u github.com/evris99/airbyte-sdk
```

## Usage

An example that creates a new workspace, destination and source. The source uses the OpenAPI connector and the destination uses the local JSON file connector. Then it creates a connection between them with the name "Test Connection".

```go
package main

import (
	"context"
	"fmt"

	airbytesdk "github.com/evris99/airbyte-sdk"
	"github.com/evris99/airbyte-sdk/types"
	"github.com/google/uuid"
)

func main() {
	client, err := airbytesdk.New("http://localhost:8000/api")
	if err != nil {
		panic(err)
	}

	// Get all source definitions and search for the ID of PokeAPI
	sourceDefinitions, err := client.ListSourceDefinitions(context.Background())
	if err != nil {
		panic(err)
	}

	var srcDefinitionID uuid.NullUUID
	for _, def := range sourceDefinitions {
		if def.Name == "PokeAPI" {
			srcDefinitionID.UUID = *def.SourceDefinitionId
			srcDefinitionID.Valid = true
		}
	}

	if !srcDefinitionID.Valid {
		panic("Could not find pokeAPI source definition")
	}

	// Create a new workspace for the connection
	workspace := &types.Workspace{
		Name: "Test",
	}

	newWorkspace, err := client.CreateWorkspace(context.Background(), workspace)
	if err != nil {
		panic(err)
	}

	// Create new Source
	source := &types.Source{
		SourceDefinitionId:      &srcDefinitionID.UUID,
		WorkspaceId:             newWorkspace.WorkspaceId,
		Name:                    "PokeAPI",
		ConnectionConfiguration: make(map[string]interface{}),
	}
	source.ConnectionConfiguration["pokemon_name"] = "snorlax"

	newSource, err := client.CreateSource(context.Background(), source)
	if err != nil {
		panic(err)
	}

	// Get all destination definitions and search for the ID of Local JSON
	destinationDefinitions, err := client.ListDestinationDefinitions(context.Background())
	if err != nil {
		panic(err)
	}

	var destDefinitionID uuid.NullUUID
	for _, def := range destinationDefinitions {
		if def.Name == "Local JSON" {
			destDefinitionID.UUID = *def.DestinationDefinitionId
			destDefinitionID.Valid = true
		}
	}

	if !destDefinitionID.Valid {
		panic("Could not find local JSON destination definition")
	}

	// Create new Destination
	dest := &types.Destination{
		DestinationDefinitionId: &destDefinitionID.UUID,
		WorkspaceId:             newWorkspace.WorkspaceId,
		Name:                    "Local JSON",
		ConnectionConfiguration: make(map[string]interface{}),
	}
	dest.ConnectionConfiguration["destination_path"] = "/json_data"

	newDest, err := client.CreateDestination(context.Background(), dest)
	if err != nil {
		panic(err)
	}

	conn := &types.Connection{
		Name:          "Test Connection",
		SourceID:      newSource.SourceId,
		DestinationId: newDest.DestinationId,
		Status:        types.Active,
	}

	// Create new connection between the 2 connectors
	newConn, err := client.CreateConnection(context.Background(), conn)
	if err != nil {
		panic(err)
	}

	fmt.Println(newConn.Name)
	// Output: Test Connection
}

```

## Contributing

All contributions are welcome and we are grateful for even the smallest of fixes! 