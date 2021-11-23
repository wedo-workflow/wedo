package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Collaboration struct {
	ID string `json:"id"`

	RID    string `json:"rid"` // Root element id
	parsed bool
}

func NewCollaboration() *Collaboration {
	return &Collaboration{}
}

func (e *Collaboration) EID() string {
	return e.ID
}

func (e *Collaboration) RootID() string {
	return e.RID
}

func (e *Collaboration) SetRootID(s string) error {
	e.RID = s
	return nil
}

func (e *Collaboration) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.parsed = true
	return nil
}
