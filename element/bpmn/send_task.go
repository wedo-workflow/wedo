package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type SendTask struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewSendTask() *SendTask {
	return &SendTask{}
}

func (e *SendTask) EID() string {
	return e.ID
}

func (e *SendTask) TypeName() string {
	return e.TName
}

func (e *SendTask) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *SendTask) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	e.parsed = true
	return nil
}
