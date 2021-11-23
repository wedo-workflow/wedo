package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type SequenceFlow struct {
	ID        string `json:"id"`
	SourceRef string `json:"sourceRef"`
	TargetRef string `json:"targetRef"`

	RID    string `json:"rid"` // Root element id
	parsed bool
}

func NewSequenceFlow() *SequenceFlow {
	return &SequenceFlow{}
}

func (e *SequenceFlow) EID() string {
	return e.ID
}

func (e *SequenceFlow) RootID() string {
	return e.RID
}

func (e *SequenceFlow) SetRootID(s string) error {
	e.RID = s
	return nil
}

func (e *SequenceFlow) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.SourceRef = element.Attr("", "sourceRef")
	e.TargetRef = element.Attr("", "targetRef")
	e.parsed = true
	return nil
}
