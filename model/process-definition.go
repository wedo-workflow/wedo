package model

import (
	"encoding/json"
	"errors"
)

// ProcessDefinition is a struct that contains the information of a process definition
type ProcessDefinition struct {
	Id                   string `json:"id"`
	ProcessDefinitionKey string `json:"process_definition_key"`
	NamespaceId          string `json:"namespace_id"`
	BusinessName         string `json:"business_name"`  // process name
	ProcessName          string `json:"process_name"`   // process name
	BusinessID           string `json:"business_id"`    // process id
	ProcessID            string `json:"process_id"`     // process id
	DefinitionsID        string `json:"definitions_id"` // definitions id
}

func (pd *ProcessDefinition) ToMapping() map[string]interface{} {
	mp := map[string]interface{}{}
	if pd == nil {
		return mp
	}
	pdb, err := json.Marshal(pd)
	if err != nil {
		return mp
	}
	if err := json.Unmarshal(pdb, &mp); err != nil {
		return mp
	}
	return mp
}

// MarshalBinary ProcessDefinition implements the BinaryMarshaler interface
func (n *ProcessDefinition) MarshalBinary() ([]byte, error) {
	if n == nil {
		return nil, errors.New("ProcessDefinition is nil")
	}
	pdBytes, err := json.Marshal(n)
	if err != nil {
		return nil, err
	}
	return pdBytes, nil
}

// UnmarshalBinary ProcessDefinition implements the BinaryUnmarshaler interface
func (n *ProcessDefinition) UnmarshalBinary(data []byte) error {
	if n == nil {
		return errors.New("ProcessDefinition is nil")
	}
	err := json.Unmarshal(data, n)
	if err != nil {
		return err
	}
	return nil
}
