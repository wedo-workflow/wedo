package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type StartEvent struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	FormKey   string      `json:"formKey"`
	Outgoings []*Outgoing `json:"outgoings"`

	TName   string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	Content []byte `json:"content"`
	parsed  bool
}

func NewStartEvent() *StartEvent {
	return &StartEvent{}
}

func (e *StartEvent) EID() string {
	return e.ID
}

func (e *StartEvent) TypeName() string {
	return e.TName
}

func (e *StartEvent) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *StartEvent) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	e.FormKey = element.Attr("http://camunda.org/schema/1.0/bpmn", "formKey")

	for _, child := range element.Children {
		if child.StartElement.Name.Local == "outgoing" {
			outgoing := NewOutgoing()
			if err := outgoing.Parse(&child); err != nil {
				return err
			}
			e.Outgoings = append(e.Outgoings, outgoing)
		}
	}
	e.Content = element.Content

	e.parsed = true
	return nil
}
