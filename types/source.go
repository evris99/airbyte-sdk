package types

import (
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

type Source struct {
	SourceId                *uuid.UUID  `json:"sourceId,omitempty"`
	SourceDefinitionId      *uuid.UUID  `json:"sourceDefinitionId,omitempty"`
	WorkspaceId             *uuid.UUID  `json:"workspaceId,omitempty"`
	ConnectionConfiguration interface{} `json:"connectionConfiguration,omitempty"`
	Name                    string      `json:"name,omitempty"`
	SourceName              string      `json:"sourceName,omitempty"`
}

// SourceFromJSON reads json data from a Reader and returns a source
func SourceFromJSON(r io.Reader) (*Source, error) {
	source := new(Source)
	err := json.NewDecoder(r).Decode(source)

	return source, err
}

// SourcesFromJSON reads json data from a Reader and returns a slice of sources
func SourcesFromJSON(r io.Reader) ([]Source, error) {
	var sources struct {
		Sources []Source `json:"sources"`
	}

	// Decode JSON
	err := json.NewDecoder(r).Decode(&sources)
	return sources.Sources, err
}
