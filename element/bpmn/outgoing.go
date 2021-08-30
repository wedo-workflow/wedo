package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Outgoing struct {
	ID string `json:"id"`

	parsed bool
}

func NewOutgoing() *Outgoing {
	return &Outgoing{}
}

func (e *Outgoing) EID() string {
	return e.ID
}

func (e *Outgoing) Parse(element *xmltree.Element) error {
	e.ID = string(element.Content)
	e.parsed = true
	return nil
}
