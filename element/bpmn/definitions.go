package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Definitions struct {
	ID              string `json:"id"`
	TargetNamespace string `json:"targetNamespace"`
	SchemaLocation  string `json:"schemaLocation"`

	parsed bool
}

func NewDefinitions() *Definitions {
	return &Definitions{}
}

func (e *Definitions) EID() string {
	return e.ID
}

func (e *Definitions) RootID() string {
	panic("implement me")
}

func (e *Definitions) SetRootID(s string) error {
	panic("implement me")
}

func (e *Definitions) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.TargetNamespace = element.Attr("", "targetNamespace")
	e.SchemaLocation = element.Attr("http://www.w3.org/2001/XMLSchema-instance", "schemaLocation")
	e.parsed = true
	return nil
}
