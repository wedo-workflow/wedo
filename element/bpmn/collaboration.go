package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Collaboration struct {
	ID string `json:"id"`

	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewCollaboration() *Collaboration {
	return &Collaboration{}
}

func (e *Collaboration) EID() string {
	return e.ID
}

func (e *Collaboration) TypeName() string {
	return e.TName
}

func (e *Collaboration) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *Collaboration) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.parsed = true
	return nil
}
