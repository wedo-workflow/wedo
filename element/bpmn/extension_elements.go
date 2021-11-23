package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type ExtensionElements struct {
	RID    string `json:"rid"` // Root element id
	parsed bool
}

func NewExtensionElements() *ExtensionElements {
	return &ExtensionElements{}
}

func (e *ExtensionElements) EID() string {
	return ""
}

func (e *ExtensionElements) RootID() string {
	return e.RID
}

func (e *ExtensionElements) SetRootID(s string) error {
	e.RID = s
	return nil
}

func (e *ExtensionElements) Parse(element *xmltree.Element) error {
	e.parsed = true
	return nil
}
