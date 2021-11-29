package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Definitions struct {
	ID              string `json:"id"`
	TargetNamespace string `json:"targetNamespace"`
	SchemaLocation  string `json:"schemaLocation"`

	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewDefinitions() *Definitions {
	return &Definitions{}
}

func (e *Definitions) EID() string {
	return e.ID
}

func (e *Definitions) TypeName() string {
	return e.TName
}

func (e *Definitions) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *Definitions) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.TargetNamespace = element.Attr("", "targetNamespace")
	e.SchemaLocation = element.Attr("http://www.w3.org/2001/XMLSchema-instance", "schemaLocation")
	e.parsed = true
	return nil
}
