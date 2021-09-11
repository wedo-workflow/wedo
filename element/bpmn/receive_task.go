package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type ReceiveTask struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	parsed bool
}

func NewReceiveTask() *ReceiveTask {
	return &ReceiveTask{}
}

func (e *ReceiveTask) EID() string {
	return e.ID
}

func (e *ReceiveTask) RootID() string {
	panic("implement me")
}

func (e *ReceiveTask) SetRootID(s string) error {
	panic("implement me")
}

func (e *ReceiveTask) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	e.parsed = true
	return nil
}
