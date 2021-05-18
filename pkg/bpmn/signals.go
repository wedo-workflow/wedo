package bpmn

import (
	"github.com/wedo-workflow/wedo/pkg/store"
	"github.com/wedo-workflow/xmltree"
)

type Signals struct {
	ID   string
	Name string

	parsed bool
}

func (s *Signals) Parse(element *xmltree.Element) error {
	s.ID = element.Attr("", "id")
	s.Name = element.Attr("", "name")
	s.parsed = true
	return nil
}

func (Signals) Store(store store.Store) error {
	store.Ping()
	return nil
}
