package types

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/google/uuid"
)

type SourceDefinition struct {
	Definition
	SourceDefinitionId *uuid.UUID `json:"sourceDefinitionId,omitempty"`
}

type SourceDefinitionSpecification struct {
	DefinitionSpecification
	SourceDefinitionId *uuid.UUID `json:"sourceDefinitionId,omitempty"`
}

// SourceDefinitionFromJSON reads json data from a Reader and returns a source definition
func SourceDefinitionFromJSON(r io.Reader) (*SourceDefinition, error) {
	sourceDefinition := new(SourceDefinition)
	err := json.NewDecoder(r).Decode(sourceDefinition)

	return sourceDefinition, fmt.Errorf("could not decode JSON: %w", err)
}

// SourcesFromJSON reads json data from a Reader and returns a slice of source definitions
func SourceDefinitionsFromJSON(r io.Reader) ([]SourceDefinition, error) {
	var sourceDefinitions struct {
		SourceDefinitions []SourceDefinition `json:"sourceDefinitions"`
	}

	// Decode JSON
	err := json.NewDecoder(r).Decode(&sourceDefinitions)
	return sourceDefinitions.SourceDefinitions, fmt.Errorf("could not decode JSON: %w", err)
}

// SourceDefinitionSpecificationFromJSON reads json data from a Reader and returns a source definition specification
func SourceDefinitionSpecificationFromJSON(r io.Reader) (*SourceDefinitionSpecification, error) {
	sourceDefinitionSpecification := new(SourceDefinitionSpecification)
	err := json.NewDecoder(r).Decode(sourceDefinitionSpecification)

	return sourceDefinitionSpecification, fmt.Errorf("could not decode JSON: %w", err)
}
