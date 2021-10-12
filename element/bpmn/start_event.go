package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type StartEvent struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	FormKey string `json:"formKey"`

	parsed bool
}

func NewStartEvent() *StartEvent {
	return &StartEvent{}
}

func (e *StartEvent) EID() string {
	return e.ID
}

func (e *StartEvent) RootID() string {
	panic("implement me")
}

func (e *StartEvent) SetRootID(s string) error {
	panic("implement me")
}

func (e *StartEvent) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	e.FormKey = element.Attr("http://camunda.org/schema/1.0/bpmn", "formKey")
	e.parsed = true
	return nil
}
