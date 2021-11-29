package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type IntermediateThrowEvent struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewIntermediateThrowEvent() *IntermediateThrowEvent {
	return &IntermediateThrowEvent{}
}

func (e *IntermediateThrowEvent) EID() string {
	return e.ID
}

func (e *IntermediateThrowEvent) TypeName() string {
	return e.TName
}

func (e *IntermediateThrowEvent) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *IntermediateThrowEvent) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	e.parsed = true
	return nil
}
