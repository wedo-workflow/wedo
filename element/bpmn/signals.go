package bpmn

import (
	"github.com/wedo-workflow/wedo/store"
	"github.com/wedo-workflow/xmltree"
)

type Signal struct {
	ID   string
	Name string

	parsed bool
}

func NewSignal() *Signal {
	return &Signal{}
}

func (s *Signal) Parse(element *xmltree.Element) error {
	s.ID = element.Attr("", "id")
	s.Name = element.Attr("", "name")
	s.parsed = true
	return nil
}

func (Signal) Store(store store.Store) error {
	store.Ping()
	return nil
}
