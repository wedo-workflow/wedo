package bpmn

import (
	"log"

	"github.com/wedo-workflow/wedo/configs"
	"github.com/wedo-workflow/wedo/pkg/store"
	"github.com/wedo-workflow/xmltree"
)

type BPMN interface {
	Parse(element *xmltree.Element) error
	Store(store store.Store) error
}

type Wedo struct {
	rootParsers map[string]BPMN
	store       store.Store
}

func NewWedo(config *configs.Config) (*Wedo, error) {
	newStore, err := store.NewStore(config)
	if err != nil {
		return nil, err
	}
	wedo := &Wedo{
		rootParsers: defaultParser(),
		store:       newStore,
	}
	return wedo, nil
}

func (w *Wedo) ParseDoc(doc []byte) error {
	tree, err := xmltree.Parse(doc)
	if err != nil {
		return err
	}
	for _, element := range tree.Flatten() {
		eleLocal := element.Name.Local
		parser, ok := w.rootParsers[eleLocal]
		if !ok {
			continue
		}
		if err := parser.Parse(element); err != nil {
			log.Print(err)
			continue
		}
		if err := parser.Store(w.store); err != nil {
			return err
		}
	}
	return nil
}

func defaultParser() map[string]BPMN {
	return map[string]BPMN{
		BPMN_ELEMENT_DEFINITIONS: NewDefinitions(),
		BPMN_ELEMENT_PROCESS:     NewProcess(),
		BPMN_ELEMENT_MESSAGE:     NewMessage(),
		BPMN_ELEMENT_GROUP:       NewGroup(),
	}
}
