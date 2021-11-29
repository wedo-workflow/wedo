package model

type ProcessIntance struct {
	Status ProcessInstanceStatus `json:"status"`
}

type ProcessInstanceStatus int

const (
	PiOK ProcessInstanceStatus = iota
)
