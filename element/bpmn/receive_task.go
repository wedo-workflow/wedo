package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type ReceiveTask struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewReceiveTask() *ReceiveTask {
	return &ReceiveTask{}
}

func (e *ReceiveTask) EID() string {
	return e.ID
}

func (e *ReceiveTask) TypeName() string {
	return e.TName
}

func (e *ReceiveTask) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *ReceiveTask) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	e.parsed = true
	return nil
}
