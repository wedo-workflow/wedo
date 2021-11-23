package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type MessageEventDefinition struct {
	ID string `json:"id"`

	RID    string `json:"rid"` // Root element id
	parsed bool
}

func NewMessageEventDefinition() *MessageEventDefinition {
	return &MessageEventDefinition{}
}

func (e *MessageEventDefinition) EID() string {
	return e.ID
}

func (e *MessageEventDefinition) RootID() string {
	return e.RID
}

func (e *MessageEventDefinition) SetRootID(s string) error {
	e.RID = s
	return nil
}

func (e *MessageEventDefinition) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.parsed = true
	return nil
}
