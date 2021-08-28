package element

import (
	"github.com/wedo-workflow/wedo/element/bpmn"
	"github.com/wedo-workflow/wedo/store"
	"github.com/wedo-workflow/xmltree"
)

type Element interface {
	// Parse parse xml dom element to element object.
	Parse(element *xmltree.Element) error
	// Store element to database.
	Store(store store.Store) error
}

func DefaultRegister() map[string]Element {
	return map[string]Element{
		BPMN_ELEMENT_DEFINITIONS: bpmn.NewDefinitions(),
		BPMN_ELEMENT_PROCESS:     bpmn.NewProcess(),
		BPMN_ELEMENT_MESSAGE:     bpmn.NewMessage(),
		BPMN_ELEMENT_GROUP:       bpmn.NewGroup(),
		BPMN_ELEMENT_SIGNAL:      bpmn.NewSignal(),
	}
}
