package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type IntermediateThrowEvent struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	RID    string `json:"rid"` // Root element id
	parsed bool
}

func NewIntermediateThrowEvent() *IntermediateThrowEvent {
	return &IntermediateThrowEvent{}
}

func (e *IntermediateThrowEvent) EID() string {
	return e.ID
}

func (e *IntermediateThrowEvent) RootID() string {
	return e.RID
}

func (e *IntermediateThrowEvent) SetRootID(s string) error {
	e.RID = s
	return nil
}

func (e *IntermediateThrowEvent) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	e.parsed = true
	return nil
}
