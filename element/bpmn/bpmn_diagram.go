package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type BPMNDiagram struct {
	ID      string `json:"id"`
	Content string `json:"content"`

	RID    string `json:"rid"` // Root element id
	parsed bool
}

func NewBPMNDiagram() *BPMNDiagram {
	return &BPMNDiagram{}
}

func (e *BPMNDiagram) EID() string {
	return e.ID
}

func (e *BPMNDiagram) RootID() string {
	return e.RID
}

func (e *BPMNDiagram) SetRootID(s string) error {
	e.RID = s
	return nil
}

func (e *BPMNDiagram) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Content = string(element.Content)
	e.parsed = true
	return nil
}
