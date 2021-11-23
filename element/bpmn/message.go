package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Message struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	RID    string `json:"rid"` // Root element id
	parsed bool
}

func NewMessage() *Message {
	return &Message{}
}

func (e *Message) EID() string {
	return e.ID
}

func (e *Message) RootID() string {
	return e.RID
}

func (e *Message) SetRootID(s string) error {
	e.RID = s
	return nil
}

func (e *Message) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	return nil
}
