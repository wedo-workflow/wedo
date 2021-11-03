package model

import "time"

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