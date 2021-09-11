package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Collaboration struct {
	ID string `json:"id"`

	parsed bool
}

func NewCollaboration() *Collaboration {
	return &Collaboration{}
}

func (e *Collaboration) EID() string {
	return e.ID
}

func (e *Collaboration) RootID() string {
	panic("implement me")
}

func (e *Collaboration) SetRootID(s string) error {
	panic("implement me")
}

func (e *Collaboration) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.parsed = true
	return nil
}
