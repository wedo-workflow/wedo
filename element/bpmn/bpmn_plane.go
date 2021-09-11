package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type BPMNPlane struct {
	ID          string `json:"id"`
	BPMNElement string `json:"bpmnElement"`

	Content string `json:"content"`

	parsed bool
}

func NewBPMNPlane() *BPMNPlane {
	return &BPMNPlane{}
}

func (e *BPMNPlane) EID() string {
	return e.ID
}

func (e *BPMNPlane) RootID() string {
	panic("implement me")
}

func (e *BPMNPlane) SetRootID(s string) error {
	panic("implement me")
}

func (e *BPMNPlane) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.BPMNElement = element.Attr("", "bpmnElement")
	e.Content = string(element.Content)
	e.parsed = true
	return nil
}
