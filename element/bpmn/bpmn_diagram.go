package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type BPMNDiagram struct {
	ID      string `json:"id"`
	Content string `json:"content"`

	parsed bool
}

func NewBPMNDiagram() *BPMNDiagram {
	return &BPMNDiagram{}
}

func (e *BPMNDiagram) EID() string {
	return e.ID
}

func (e *BPMNDiagram) RootID() string {
	panic("implement me")
}

func (e *BPMNDiagram) SetRootID(s string) error {
	panic("implement me")
}

func (e *BPMNDiagram) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Content = string(element.Content)
	e.parsed = true
	return nil
}
