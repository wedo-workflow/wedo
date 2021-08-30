package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Incoming struct {
	ID string `json:"id"`

	parsed bool
}

func NewIncoming() *Incoming {
	return &Incoming{}
}

func (e *Incoming) EID() string {
	return e.ID
}

func (e *Incoming) Parse(element *xmltree.Element) error {
	e.ID = string(element.Content)
	e.parsed = true
	return nil
}
