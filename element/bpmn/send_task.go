package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type SendTask struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	parsed bool
}

func NewSendTask() *SendTask {
	return &SendTask{}
}

func (e *SendTask) EID() string {
	return e.ID
}

func (e *SendTask) RootID() string {
	panic("implement me")
}

func (e *SendTask) SetRootID(s string) error {
	panic("implement me")
}

func (e *SendTask) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	e.parsed = true
	return nil
}
