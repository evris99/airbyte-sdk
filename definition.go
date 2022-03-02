package airbytesdk

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ReleaseStageEnum int

const (
	Alpha ReleaseStageEnum = iota
	Beta
	GenerallyAvailable
	Custom
)

// Unmarshaler for json
func (r *ReleaseStageEnum) UnmarshalJSON(b []byte) error {
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
func (r ReleaseStageEnum) MarshalJSON() ([]byte, error) {
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
	Name             string           `json:"name,omitempty"`
	DockerRepository string           `json:"dockerRepository,omitempty"`
	DockerImageTag   string           `json:"dockerImageTag,omitempty"`
	DocumentationURL string           `json:"documentationUrl,omitempty"`
	Icon             string           `json:"icon,omitempty"`
	ReleaseStage     ReleaseStageEnum `json:"releaseStage,omitempty"`
	ReleaseDate      string           `json:"releaseDate,omitempty"`
}

const (
	OAuth AuthenticationTypeEnum = iota
)

// Unmarshaler for json
func (a *AuthenticationTypeEnum) UnmarshalJSON(b []byte) error {
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
func (a AuthenticationTypeEnum) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	case OAuth:
		s = "oath2.0"
	default:
		return nil, fmt.Errorf("unknown authentication type")
	}

	return json.Marshal(s)
}

type Oauth2SpecificationType struct {
	RootObject                interface{} `json:"rootObject,omitempty"`
	OauthFlowInitParameters   [][]string  `json:"oauthFlowInitParameters,omitempty"`
	OauthFlowOutputParameters [][]string  `json:"oauthFlowOutputParameters,omitempty"`
}

type AuthSpecificationType struct {
	AuthType            *AuthenticationTypeEnum  `json:"auth_type,omitempty"`
	Oauth2Specification *Oauth2SpecificationType `json:"oauth2Specification,omitempty"`
}

type AuthFlowTypeEnum int

const (
	OAuth2 AuthFlowTypeEnum = iota
	OAuth1
)

// Unmarshaler for json
func (a *AuthFlowTypeEnum) UnmarshalJSON(b []byte) error {
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
func (a AuthFlowTypeEnum) MarshalJSON() ([]byte, error) {
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

type OauthConfigSpecificationType struct {
	OauthUserInputFromConnectorConfigSpecification []byte `json:"oauthUserInputFromConnectorConfigSpecification,omitempty"`
	CompleteOAuthOutputSpecification               []byte `json:"completeOAuthOutputSpecification,omitempty"`
	CompleteOAuthServerInputSpecification          []byte `json:"completeOAuthServerInputSpecification,omitempty"`
	CompleteOAuthServerOutputSpecification         []byte `json:"completeOAuthServerOutputSpecification,omitempty"`
}

type AdvancedAuthType struct {
	AuthFlowType             AuthFlowTypeEnum              `json:"authFlowType,omitempty"`
	PredicateKey             []string                      `json:"predicateKey,omitempty"`
	PredicateValue           string                        `json:"predicateValue,omitempty"`
	OauthConfigSpecification *OauthConfigSpecificationType `json:"oauthConfigSpecification,omitempty"`
}

type DefinitionSpecification struct {
	DocumentationUrl        string                 `json:"documentationUrl,omitempty"`
	ConnectionSpecification map[string]interface{} `json:"connectionSpecification,omitempty"`
	AuthSpecification       AuthSpecificationType  `json:"authSpecification,omitempty"`
	AdvancedAuth            AdvancedAuthType       `json:"advancedAuth,omitempty"`
	JobInfo                 *JobInfoType           `json:"jobInfo,omitempty"`
}
