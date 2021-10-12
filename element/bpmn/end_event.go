package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type EndEvent struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	parsed bool
}

func NewEndEvent() *EndEvent {
	return &EndEvent{}
}

func (e *EndEvent) EID() string {
	return e.ID
}

func (e *EndEvent) RootID() string {
	panic("implement me")
}

func (e *EndEvent) SetRootID(s string) error {
	panic("implement me")
}

func (e *EndEvent) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	e.parsed = true
	return nil
}
