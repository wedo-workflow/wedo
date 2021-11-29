package model

import (
	"encoding/json"
	"errors"
	"time"
)

type Deployment struct {
	DID           string
	NamespaceID   string    `json:"namespace_id"`
	Name          string    `json:"name"`
	BusinessName  string    `json:"business_name"`  // process name
	BusinessID    string    `json:"business_id"`    // process id
	DefinitionsID string    `json:"definitions_id"` // definitions id
	Content       []byte    `json:"content"`
	CreateTime    time.Time `json:"create_time"`
}

// DeploymentListOptions specifies the optional parameters to various
// methods that list deployments.
type DeploymentListOptions struct {
	// For pagination.
	PageOpts PageOpts
}

// PageOpts is used for specifying pagination options.
type PageOpts struct {
	Page int
	Size int
}

// MarshalBinary Deployment implements the BinaryMarshaler interface
func (d *Deployment) MarshalBinary() ([]byte, error) {
	if d == nil {
		return nil, errors.New("Deployment is nil")
	}
	DeployBytes, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	return DeployBytes, nil
}

// UnmarshalBinary Deployment implements the BinaryUnmarshaler interface
func (d *Deployment) UnmarshalBinary(data []byte) error {
	if d == nil {
		return errors.New("Deployment is nil")
	}
	err := json.Unmarshal(data, d)
	if err != nil {
		return err
	}
	return nil
}
