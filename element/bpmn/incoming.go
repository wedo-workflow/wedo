package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Incoming struct {
	ID string `json:"id"`

	RID    string `json:"rid"` // Root element id
	parsed bool
}

func NewIncoming() *Incoming {
	return &Incoming{}
}

func (e *Incoming) EID() string {
	return e.ID
}

func (e *Incoming) RootID() string {
	return e.RID
}

func (e *Incoming) SetRootID(s string) error {
	e.RID = s
	return nil
}

func (e *Incoming) Parse(element *xmltree.Element) error {
	e.ID = string(element.Content)
	e.parsed = true
	return nil
}
