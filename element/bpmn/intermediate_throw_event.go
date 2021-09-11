package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type IntermediateThrowEvent struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	parsed bool
}

func NewIntermediateThrowEvent() *IntermediateThrowEvent {
	return &IntermediateThrowEvent{}
}

func (e *IntermediateThrowEvent) EID() string {
	return e.ID
}

func (e *IntermediateThrowEvent) RootID() string {
	panic("implement me")
}

func (e *IntermediateThrowEvent) SetRootID(s string) error {
	panic("implement me")
}

func (e *IntermediateThrowEvent) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	e.parsed = true
	return nil
}
