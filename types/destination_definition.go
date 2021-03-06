package types

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/google/uuid"
)

type DestinationDefinition struct {
	Definition
	DestinationDefinitionId *uuid.UUID `json:"destinationDefinitionId,omitempty"`
}

type SupportedDestinationSyncModesType int

const (
	Append SupportedDestinationSyncModesType = iota + 1
	Overwrite
	AppendDedup
)

// Unmarshaler for json
func (sup *SupportedDestinationSyncModesType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch strings.ToLower(s) {
	case "append":
		*sup = Append
	case "overwrite":
		*sup = Overwrite
	case "append_dedup":
		*sup = AppendDedup
	}

	return nil
}

// Marshaler for json
func (sup SupportedDestinationSyncModesType) MarshalJSON() ([]byte, error) {
	var s string
	switch sup {
	case Append:
		s = "append"
	case Overwrite:
		s = "overwrite"
	case AppendDedup:
		s = "append_dedup"
	}

	return json.Marshal(s)
}

type DestinationDefinitionSpecification struct {
	DefinitionSpecification
	DestinationDefinitionId       *uuid.UUID                        `json:"destinationDefinitionId,omitempty"`
	SupportedDestinationSyncModes SupportedDestinationSyncModesType `json:"supportedDestinationSyncModes,omitempty"`
	SupportsDbt                   bool                              `json:"supportsDbt,omitempty"`
	SupportsNormalization         bool                              `json:"supportsNormalization,omitempty"`
}

// DestinationDefinitionFromJSON reads json data from a Reader and returns a destination definition
func DestinationDefinitionFromJSON(r io.Reader) (*DestinationDefinition, error) {
	destinationDef := new(DestinationDefinition)
	err := json.NewDecoder(r).Decode(destinationDef)
	return destinationDef, err
}

// DestinationDefinitionsFromJSON reads json data from a Reader and returns a slice of destinations definitions
func DestinationDefinitionsFromJSON(r io.Reader) ([]DestinationDefinition, error) {
	var destinationDefs struct {
		DestinationDefinitions []DestinationDefinition `json:"destinationDefinitions"`
	}

	// Decode JSON
	err := json.NewDecoder(r).Decode(&destinationDefs)
	return destinationDefs.DestinationDefinitions, err
}

// DestinationDefinitionSpecificationFromJSON reads json data from a Reader and returns a destination definition specification
func DestinationDefinitionSpecificationToJSON(r io.Reader) (*DestinationDefinitionSpecification, error) {
	destinationDefSpec := new(DestinationDefinitionSpecification)
	err := json.NewDecoder(r).Decode(destinationDefSpec)
	return destinationDefSpec, err
}
