package model

import (
	"encoding/json"
	"errors"
	"time"
)

type Deploy struct {
	DID        string
	Name       string    `json:"name"`
	Content    []byte    `json:"content"`
	CreateTime time.Time `json:"create_time"`
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

// MarshalBinary Deploy implements the BinaryMarshaler interface
func (d *Deploy) MarshalBinary() ([]byte, error) {
	if d == nil {
		return nil, errors.New("Deploy is nil")
	}
	DeployBytes, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	return DeployBytes, nil
}

// UnmarshalBinary Deploy implements the BinaryUnmarshaler interface
func (d *Deploy) UnmarshalBinary(data []byte) error {
	if d == nil {
		return errors.New("Deploy is nil")
	}
	err := json.Unmarshal(data, d)
	if err != nil {
		return err
	}
	return nil
}
