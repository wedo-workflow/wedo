package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type SequenceFlow struct {
	ID        string `json:"id"`
	SourceRef string `json:"sourceRef"`
	TargetRef string `json:"targetRef"`

	parsed bool
}

func NewSequenceFlow() *SequenceFlow {
	return &SequenceFlow{}
}

func (e *SequenceFlow) EID() string {
	return e.ID
}

func (e *SequenceFlow) RootID() string {
	panic("implement me")
}

func (e *SequenceFlow) SetRootID(s string) error {
	panic("implement me")
}

func (e *SequenceFlow) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.SourceRef = element.Attr("", "sourceRef")
	e.TargetRef = element.Attr("", "targetRef")
	e.parsed = true
	return nil
}
