package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ReleaseStage int

const (
	Alpha ReleaseStage = iota
	Beta
	GenerallyAvailable
	Custom
)

// Unmarshaler for json
func (r *ReleaseStage) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch strings.ToLower(s) {
	case "alpha":
		*r = Alpha
	case "beta":
		*r = Beta
	case "generally_available":
		*r = GenerallyAvailable
	case "custom":
		*r = Custom
	default:
		return fmt.Errorf("unknown release stage")
	}

	return nil
}

// Marshaler for json
func (r ReleaseStage) MarshalJSON() ([]byte, error) {
	var s string
	switch r {
	case Alpha:
		s = "alpha"
	case Beta:
		s = "beta"
	case GenerallyAvailable:
		s = "generally_available"
	case Custom:
		s = "custom"
	default:
		return nil, fmt.Errorf("unknown release stage")
	}

	return json.Marshal(s)
}

type Definition struct {
	Name             string       `json:"name,omitempty"`
	DockerRepository string       `json:"dockerRepository,omitempty"`
	DockerImageTag   string       `json:"dockerImageTag,omitempty"`
	DocumentationURL string       `json:"documentationUrl,omitempty"`
	Icon             string       `json:"icon,omitempty"`
	ReleaseStage     ReleaseStage `json:"releaseStage,omitempty"`
	ReleaseDate      string       `json:"releaseDate,omitempty"`
}

type AuthenticationType int

const (
	OAuth AuthenticationType = iota
)

// Unmarshaler for json
func (a *AuthenticationType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch strings.ToLower(s) {
	case "oath2.0":
		*a = OAuth
	default:
		return fmt.Errorf("unknown authentication type")
	}

	return nil
}

// Marshaler for json
func (a AuthenticationType) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	case OAuth:
		s = "oath2.0"
	default:
		return nil, fmt.Errorf("unknown authentication type")
	}

	return json.Marshal(s)
}

type Oauth2Specification struct {
	RootObject                interface{} `json:"rootObject,omitempty"`
	OauthFlowInitParameters   [][]string  `json:"oauthFlowInitParameters,omitempty"`
	OauthFlowOutputParameters [][]string  `json:"oauthFlowOutputParameters,omitempty"`
}

type AuthSpecification struct {
	AuthType            *AuthenticationType  `json:"auth_type,omitempty"`
	Oauth2Specification *Oauth2Specification `json:"oauth2Specification,omitempty"`
}

type AuthFlowType int

const (
	OAuth2 AuthFlowType = iota
	OAuth1
)

// Unmarshaler for json
func (a *AuthFlowType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch strings.ToLower(s) {
	case "oath2.0":
		*a = OAuth2
	case "oath1.0":
		*a = OAuth1
	default:
		return fmt.Errorf("unknown auth flow type")
	}

	return nil
}

// Marshaler for json
func (a AuthFlowType) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	case OAuth2:
		s = "oath2.0"
	case OAuth1:
		s = "oath1.0"
	default:
		return nil, fmt.Errorf("unknown auth flow type")
	}

	return json.Marshal(s)
}

type OauthConfigSpecification struct {
	OauthUserInputFromConnectorConfigSpecification []byte `json:"oauthUserInputFromConnectorConfigSpecification,omitempty"`
	CompleteOAuthOutputSpecification               []byte `json:"completeOAuthOutputSpecification,omitempty"`
	CompleteOAuthServerInputSpecification          []byte `json:"completeOAuthServerInputSpecification,omitempty"`
	CompleteOAuthServerOutputSpecification         []byte `json:"completeOAuthServerOutputSpecification,omitempty"`
}

type AdvancedAuth struct {
	AuthFlowType             AuthFlowType              `json:"authFlowType,omitempty"`
	PredicateKey             []string                  `json:"predicateKey,omitempty"`
	PredicateValue           string                    `json:"predicateValue,omitempty"`
	OauthConfigSpecification *OauthConfigSpecification `json:"oauthConfigSpecification,omitempty"`
}

type DefinitionSpecification struct {
	DocumentationUrl        string                 `json:"documentationUrl,omitempty"`
	ConnectionSpecification map[string]interface{} `json:"connectionSpecification,omitempty"`
	AuthSpecification       AuthSpecification      `json:"authSpecification,omitempty"`
	AdvancedAuth            *AdvancedAuth          `json:"advancedAuth,omitempty"`
	JobInfo                 *JobInfo               `json:"jobInfo,omitempty"`
}
