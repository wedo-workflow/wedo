package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type BPMNPlane struct {
	ID          string `json:"id"`
	BPMNElement string `json:"bpmnElement"`

	Content string `json:"content"`

	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewBPMNPlane() *BPMNPlane {
	return &BPMNPlane{}
}

func (e *BPMNPlane) EID() string {
	return e.ID
}

func (e *BPMNPlane) TypeName() string {
	return e.TName
}

func (e *BPMNPlane) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *BPMNPlane) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.BPMNElement = element.Attr("", "bpmnElement")
	e.Content = string(element.Content)
	e.parsed = true
	return nil
}
