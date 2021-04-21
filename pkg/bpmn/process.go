package bpmn

import "github.com/wedo-workflow/xmltree"

type Process struct {
	ID           string
	Name         string
	IsExecutable bool
}

func NewProcess() *Process {
	return &Process{}
}

func (p *Process) Parse(*xmltree.Element) error {
	panic("implement me")
	return nil
}

func (p *Process) Store() {
	panic("implement me")
}

func (p *Process) List() ([]BPMN, error) {
	panic("implement me")
	return nil, nil
}
