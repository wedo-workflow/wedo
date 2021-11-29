package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Signal struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewSignal() *Signal {
	return &Signal{}
}

func (e *Signal) EID() string {
	return e.ID
}

func (e *Signal) TypeName() string {
	return e.TName
}

func (e *Signal) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *Signal) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	e.parsed = true
	return nil
}
