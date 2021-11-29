package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Task struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Exclusive bool        `json:"exclusive"`
	Incomings []*Incoming `json:"incomings"`
	Outgoings []*Outgoing `json:"outgoings"`

	TName   string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	Content []byte `json:"content"`
	parsed  bool
}

func NewTask() *Task {
	return &Task{
		Incomings: []*Incoming{},
		Outgoings: []*Outgoing{},
	}
}

func (e *Task) EID() string {
	return e.ID
}

func (e *Task) TypeName() string {
	return e.TName
}

func (e *Task) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *Task) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	for _, child := range element.Children {
		if child.StartElement.Name.Local == "incoming" {
			incoming := NewIncoming()
			if err := incoming.Parse(&child); err != nil {
				return err
			}
			e.Incomings = append(e.Incomings, incoming)
		}
		if child.StartElement.Name.Local == "outgoing" {
			outgoing := NewOutgoing()
			if err := outgoing.Parse(&child); err != nil {
				return err
			}
			e.Outgoings = append(e.Outgoings, outgoing)
		}
	}
	e.Name = element.Attr("", "name")
	e.Exclusive = element.Attr("http://camunda.org/schema/1.0/bpmn", "exclusive") == "true"
	e.Content = element.Content

	e.parsed = true
	return nil
}
