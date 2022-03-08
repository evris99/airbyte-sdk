package types

import (
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

type Destination struct {
	DestinationId           *uuid.UUID             `json:"destinationId,omitempty"`
	DestinationDefinitionId *uuid.UUID             `json:"destinationDefinitionId,omitempty"`
	WorkspaceId             *uuid.UUID             `json:"workspaceId,omitempty"`
	ConnectionConfiguration map[string]interface{} `json:"connectionConfiguration,omitempty"`
	Name                    string                 `json:"name,omitempty"`
	DestinationName         string                 `json:"destinationName,omitempty"`
}

// DestinationFromJSON reads json data from a Reader and returns a destination
func DestinationFromJSON(r io.Reader) (*Destination, error) {
	destination := new(Destination)
	err := json.NewDecoder(r).Decode(destination)

	return destination, err
}

// DestinationsFromJSON reads json data from a Reader and returns a slice of destinations
func DestinationsFromJSON(r io.Reader) ([]Destination, error) {
	var destinations struct {
		Destinations []Destination `json:"destinations"`
	}

	// Decode JSON
	err := json.NewDecoder(r).Decode(&destinations)
	return destinations.Destinations, err
}
