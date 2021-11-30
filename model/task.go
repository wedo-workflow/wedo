package model

import (
	"encoding/json"
	"errors"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// Task
type Task struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Id                string
	Name              string
	Description       string
	Assignee          string
	Owner             string
	DelegationState   string
	Due               *timestamppb.Timestamp
	Priority          string
	ParentTaskId      string
	ProcessInstanceId string
	TaskDefinitionKey string
	NamespaceId       string
	FormKey           string
	FollowUpDate      *timestamppb.Timestamp
}

// MarshalBinary Task implements the BinaryMarshaler interface
func (t *Task) MarshalBinary() ([]byte, error) {
	if t == nil {
		return nil, errors.New("Task is nil")
	}
	userBytes, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return userBytes, nil
}

// UnmarshalBinary Task implements the BinaryUnmarshaler interface
func (t *Task) UnmarshalBinary(data []byte) error {
	if t == nil {
		return errors.New("Task is nil")
	}
	err := json.Unmarshal(data, t)
	if err != nil {
		return err
	}
	return nil
}

// TaskListOptions task list options
type TaskListOptions struct {
	ProcessInstanceID string
}
