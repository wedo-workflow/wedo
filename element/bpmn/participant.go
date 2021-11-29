package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Participant struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ProcessRef string `json:"processRef"`

	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewParticipant() *Participant {
	return &Participant{}
}

func (e *Participant) EID() string {
	return e.ID
}

func (e *Participant) TypeName() string {
	return e.TName
}

func (e *Participant) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *Participant) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	e.ProcessRef = element.Attr("", "processRef")
	e.parsed = true
	return nil
}
