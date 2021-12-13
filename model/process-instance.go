package model

import (
	"encoding/json"
	"errors"
	wedo "github.com/wedo-workflow/wedo/api"
)

type ProcessInstance struct {
	Id                string                    `json:"id"`
	Status            wedo.ProcessInstanceState `json:"status"`
	ProcessDefinition *ProcessDefinition        `json:"processDefinition"`
}

// MarshalBinary ProcessDefinition implements the BinaryMarshaler interface
func (pi *ProcessInstance) MarshalBinary() ([]byte, error) {
	if pi == nil {
		return nil, errors.New("ProcessInstance is nil")
	}
	pdBytes, err := json.Marshal(pi)
	if err != nil {
		return nil, err
	}
	return pdBytes, nil
}

// UnmarshalBinary ProcessInstance implements the BinaryUnmarshaler interface
func (pi *ProcessInstance) UnmarshalBinary(data []byte) error {
	if pi == nil {
		return errors.New("ProcessDefinition is nil")
	}
	err := json.Unmarshal(data, pi)
	if err != nil {
		return err
	}
	return nil
}
