package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type ExtensionElements struct {
	parsed bool
}

func NewExtensionElements() *ExtensionElements {
	return &ExtensionElements{}
}

func (e *ExtensionElements) EID() string {
	return ""
}

func (e *ExtensionElements) RootID() string {
	panic("implement me")
}

func (e *ExtensionElements) SetRootID(s string) error {
	panic("implement me")
}

func (e *ExtensionElements) Parse(element *xmltree.Element) error {
	e.parsed = true
	return nil
}
