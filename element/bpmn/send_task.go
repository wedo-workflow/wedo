package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type SendTask struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	RID    string `json:"rid"` // Root element id
	parsed bool
}

func NewSendTask() *SendTask {
	return &SendTask{}
}

func (e *SendTask) EID() string {
	return e.ID
}

func (e *SendTask) RootID() string {
	return e.RID
}

func (e *SendTask) SetRootID(s string) error {
	e.RID = s
	return nil
}

func (e *SendTask) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	e.parsed = true
	return nil
}
