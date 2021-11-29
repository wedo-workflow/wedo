package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Message struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewMessage() *Message {
	return &Message{}
}

func (e *Message) EID() string {
	return e.ID
}

func (e *Message) TypeName() string {
	return e.TName
}

func (e *Message) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *Message) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	return nil
}
