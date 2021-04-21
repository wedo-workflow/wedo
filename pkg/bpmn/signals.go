package bpmn

import "github.com/wedo-workflow/xmltree"

type Signals struct{}

func (Signals) Parse(*xmltree.Element) error {
	panic("implement me")
	return nil
}

func (Signals) Store() {
	panic("implement me")
}

func (Signals) List() ([]BPMN, error) {
	panic("implement me")
}
