package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type EndEvent struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewEndEvent() *EndEvent {
	return &EndEvent{}
}

func (e *EndEvent) EID() string {
	return e.ID
}

func (e *EndEvent) TypeName() string {
	return e.TName
}

func (e *EndEvent) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *EndEvent) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	e.parsed = true
	return nil
}
