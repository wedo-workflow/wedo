package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type ExtensionElements struct {
	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewExtensionElements() *ExtensionElements {
	return &ExtensionElements{}
}

func (e *ExtensionElements) EID() string {
	return ""
}

func (e *ExtensionElements) TypeName() string {
	return e.TName
}

func (e *ExtensionElements) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *ExtensionElements) Parse(element *xmltree.Element) error {
	e.parsed = true
	return nil
}
