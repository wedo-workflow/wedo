package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type SequenceFlow struct {
	ID        string `json:"id"`
	SourceRef string `json:"sourceRef"`
	TargetRef string `json:"targetRef"`

	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewSequenceFlow() *SequenceFlow {
	return &SequenceFlow{}
}

func (e *SequenceFlow) EID() string {
	return e.ID
}

func (e *SequenceFlow) TypeName() string {
	return e.TName
}

func (e *SequenceFlow) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *SequenceFlow) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.SourceRef = element.Attr("", "sourceRef")
	e.TargetRef = element.Attr("", "targetRef")
	e.parsed = true
	return nil
}
