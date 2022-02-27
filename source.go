package airbytesdk

import "github.com/google/uuid"

type Source struct {
	SourceId                *uuid.UUID  `json:"sourceId"`
	SourceDefinitionId      *uuid.UUID  `json:"sourceDefinitionId"`
	WorkspaceId             *uuid.UUID  `json:"workspaceId"`
	ConnectionConfiguration interface{} `json:"connectionConfiguration"`
	Name                    string      `json:"name"`
	SourceName              string      `json:"sourceName"`
}
