package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type ReceiveTask struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	RID    string `json:"rid"` // Root element id
	parsed bool
}

func NewReceiveTask() *ReceiveTask {
	return &ReceiveTask{}
}

func (e *ReceiveTask) EID() string {
	return e.ID
}

func (e *ReceiveTask) RootID() string {
	return e.RID
}

func (e *ReceiveTask) SetRootID(s string) error {
	e.RID = s
	return nil
}

func (e *ReceiveTask) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	e.parsed = true
	return nil
}
