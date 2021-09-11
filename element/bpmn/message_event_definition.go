package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type MessageEventDefinition struct {
	ID string `json:"id"`

	parsed bool
}

func NewMessageEventDefinition() *MessageEventDefinition {
	return &MessageEventDefinition{}
}

func (e *MessageEventDefinition) EID() string {
	return e.ID
}

func (e *MessageEventDefinition) RootID() string {
	panic("implement me")
}

func (e *MessageEventDefinition) SetRootID(s string) error {
	panic("implement me")
}

func (e *MessageEventDefinition) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.parsed = true
	return nil
}
