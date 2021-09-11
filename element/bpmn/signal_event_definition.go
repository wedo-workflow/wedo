package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type SignalEventDefinition struct {
	ID string `json:"id"`

	parsed bool
}

func NewSignalEventDefinition() *SignalEventDefinition {
	return &SignalEventDefinition{}
}

func (e *SignalEventDefinition) EID() string {
	return e.ID
}

func (e *SignalEventDefinition) RootID() string {
	panic("implement me")
}

func (e *SignalEventDefinition) SetRootID(s string) error {
	panic("implement me")
}

func (e *SignalEventDefinition) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.parsed = true
	return nil
}
