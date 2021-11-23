package model

import (
	"encoding/json"
	"errors"
)

type Namespace struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// NamespaceQueryOptions is the query options for a namespace
type NamespaceQueryOptions struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	NameLike string `json:"name_like"`
	// Limit is the maximum number of namespaces to return
	Limit int `json:"limit"`
	// Offset is the offset from the first namespace to return
	Offset int `json:"offset"`
}

// MarshalBinary Namespace implements the BinaryMarshaler interface
func (n *Namespace) MarshalBinary() ([]byte, error) {
	if n == nil {
		return nil, errors.New("Namespace is nil")
	}
	namespaceBytes, err := json.Marshal(n)
	if err != nil {
		return nil, err
	}
	return namespaceBytes, nil
}

// UnmarshalBinary Namespace implements the BinaryUnmarshaler interface
func (n *Namespace) UnmarshalBinary(data []byte) error {
	if n == nil {
		return errors.New("Namespace is nil")
	}
	err := json.Unmarshal(data, n)
	if err != nil {
		return err
	}
	return nil
}
