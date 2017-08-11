package v5

import (
	"encoding/json"
	"fmt"

	"github.com/DataDog/datadog-agent/pkg/metadata/common"
	"github.com/DataDog/datadog-agent/pkg/metadata/host"
	"github.com/DataDog/datadog-agent/pkg/metadata/resources"
)

// CommonPayload wraps Payload from the common package
type CommonPayload struct {
	common.Payload
}

// HostPayload wraps Payload from the host package
type HostPayload struct {
	host.Payload
}

// ResourcesPayload wraps Payload from the resources package
type ResourcesPayload struct {
	resources.Payload `json:"resources"`
}

// MarshalJSON serialization a Payload to JSON
func (p *Payload) MarshalJSON() ([]byte, error) {
	// use an alias to avoid infinite recursion while serializing
	type PayloadAlias Payload

	return json.Marshal((*PayloadAlias)(p))
}

// Marshal not implemented
func (p *Payload) Marshal() ([]byte, error) {
	return nil, fmt.Errorf("V5 Payload serialization is not implemented")
}
