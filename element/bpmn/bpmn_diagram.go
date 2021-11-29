package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type BPMNDiagram struct {
	ID      string `json:"id"`
	Content string `json:"content"`

	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewBPMNDiagram() *BPMNDiagram {
	return &BPMNDiagram{}
}

func (e *BPMNDiagram) EID() string {
	return e.ID
}

func (e *BPMNDiagram) TypeName() string {
	return e.TName
}

func (e *BPMNDiagram) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *BPMNDiagram) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Content = string(element.Content)
	e.parsed = true
	return nil
}
