package bpmn

import (
	"log"

	"github.com/wedo-workflow/xmltree"
)

type BPMN interface {
	Parse(element *xmltree.Element) error
	Store()
}

type B struct {
	rootParsers map[string]BPMN
}

func NewB() *B {
	rootParsers := map[string]BPMN{
		BPMN_ELEMENT_PROCESS: NewProcess(),
	}
	return &B{
		rootParsers: rootParsers,
	}
}

func (b *B) Parse(doc []byte) {
	tree, err := xmltree.Parse(doc)
	if err != nil {
		log.Print(err)
		return
	}
	for _, element := range tree.Flatten() {
		eleLocal := element.Name.Local
		parser, ok := b.rootParsers[eleLocal]
		if !ok {
			continue
		}
		if err := parser.Parse(element); err != nil {
			log.Print(err)
			continue
		}
		parser.Store()
	}
}
