package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Outgoing struct {
	ID string `json:"id"`

	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewOutgoing() *Outgoing {
	return &Outgoing{}
}

func (e *Outgoing) EID() string {
	return e.ID
}

func (e *Outgoing) TypeName() string {
	return e.TName
}

func (e *Outgoing) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *Outgoing) Parse(element *xmltree.Element) error {
	e.ID = string(element.Content)
	e.parsed = true
	return nil
}
