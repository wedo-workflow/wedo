package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Participant struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ProcessRef string `json:"processRef"`

	RID    string `json:"rid"` // Root element id
	parsed bool
}

func NewParticipant() *Participant {
	return &Participant{}
}

func (e *Participant) EID() string {
	return e.ID
}

func (e *Participant) RootID() string {
	return e.RID
}

func (e *Participant) SetRootID(s string) error {
	e.RID = s
	return nil
}

func (e *Participant) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	e.ProcessRef = element.Attr("", "processRef")
	e.parsed = true
	return nil
}
