package bpmn

import (
	"github.com/wedo-workflow/wedo/pkg/store"
	"github.com/wedo-workflow/xmltree"
)

type Definitions struct {
	ID              string
	TargetNamespace string
	SchemaLocation  string

	parsed bool
}

func NewDefinitions() *Definitions {
	return &Definitions{}
}

func (d *Definitions) Parse(element *xmltree.Element) error {
	d.ID = element.Attr("", "id")
	d.TargetNamespace = element.Attr("", "targetNamespace")
	d.SchemaLocation = element.Attr("http://www.w3.org/2001/XMLSchema-instance", "schemaLocation")
	d.parsed = true
	return nil
}

func (d *Definitions) Store(store store.Store) {
	store.Ping()
}
