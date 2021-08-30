package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Task struct {
	ID        string      `json:"id"`
	Incomings []*Incoming `json:"incomings"`
	Outgoings []*Outgoing `json:"outgoings"`

	parsed bool
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

func (e *Task) RootID() string {
	panic("implement me")
}

func (e *Task) SetRootID(s string) error {
	panic("implement me")
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

	e.parsed = true
	return nil
}
