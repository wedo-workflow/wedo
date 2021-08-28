package bpmn

import (
	"github.com/wedo-workflow/wedo/store"
	"github.com/wedo-workflow/xmltree"
)

type Group struct {
	ID               string
	CategoryValueRef string

	parsed bool
}

func NewGroup() *Group {
	return &Group{}
}

func (g *Group) Parse(element *xmltree.Element) error {
	g.ID = element.Attr("", "id")
	g.CategoryValueRef = element.Attr("", "categoryValueRef")
	g.parsed = true
	return nil
}

func (g *Group) Store(store store.Store) error {
	if !g.parsed {
		return nil
	}
	store.Ping()
	return nil
}
