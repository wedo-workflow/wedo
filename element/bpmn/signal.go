package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Signal struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	parsed bool
}

func NewSignal() *Signal {
	return &Signal{}
}
func (e *Signal) EID() string {
	return e.ID
}
func (e *Signal) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	e.parsed = true
	return nil
}
