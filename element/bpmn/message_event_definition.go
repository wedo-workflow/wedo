package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type MessageEventDefinition struct {
	ID string `json:"id"`

	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewMessageEventDefinition() *MessageEventDefinition {
	return &MessageEventDefinition{}
}

func (e *MessageEventDefinition) EID() string {
	return e.ID
}

func (e *MessageEventDefinition) TypeName() string {
	return e.TName
}

func (e *MessageEventDefinition) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *MessageEventDefinition) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.parsed = true
	return nil
}
