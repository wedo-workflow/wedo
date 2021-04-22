package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Process struct {
	ID           string
	Name         string
	IsExecutable bool
	Version      string

	parsed bool
}

func NewProcess() *Process {
	return &Process{}
}

func (p *Process) Parse(element *xmltree.Element) error {
	p.ID = element.Attr("", "id")
	p.Name = element.Attr("", "name")
	p.IsExecutable = element.Attr("", "isExecutable") == "true"
	p.Version = element.Attr("http://camunda.org/schema/1.0/bpmn", "versionTag")
	p.parsed = true
	return nil
}

func (p *Process) Store() {
	if !p.parsed {
		return
	}
}
