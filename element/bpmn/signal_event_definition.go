package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type SignalEventDefinition struct {
	ID string `json:"id"`

	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewSignalEventDefinition() *SignalEventDefinition {
	return &SignalEventDefinition{}
}

func (e *SignalEventDefinition) EID() string {
	return e.ID
}

func (e *SignalEventDefinition) TypeName() string {
	return e.TName
}

func (e *SignalEventDefinition) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *SignalEventDefinition) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.parsed = true
	return nil
}
