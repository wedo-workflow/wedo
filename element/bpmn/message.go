package bpmn

import (
	"github.com/wedo-workflow/wedo/store"
	"github.com/wedo-workflow/xmltree"
)

type Message struct {
	ID   string
	Name string
}

func NewMessage() *Message {
	return &Message{}
}

func (m *Message) Parse(element *xmltree.Element) error {
	m.ID = element.Attr("", "id")
	m.Name = element.Attr("", "name")
	return nil
}

func (m *Message) Store(store store.Store) error {
	store.Ping()
	return nil
}
