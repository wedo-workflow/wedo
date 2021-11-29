package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Incoming struct {
	ID string `json:"id"`

	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewIncoming() *Incoming {
	return &Incoming{}
}

func (e *Incoming) EID() string {
	return e.ID
}

func (e *Incoming) TypeName() string {
	return e.TName
}

func (e *Incoming) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *Incoming) Parse(element *xmltree.Element) error {
	e.ID = string(element.Content)
	e.parsed = true
	return nil
}
