package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Process struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	IsExecutable bool   `json:"isExecutable"`
	Version      string `json:"versionTag"`

	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewProcess() *Process {
	return &Process{}
}

func (e *Process) EID() string {
	return e.ID
}

func (e *Process) TypeName() string {
	return e.TName
}

func (e *Process) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *Process) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	e.IsExecutable = element.Attr("", "isExecutable") == "true"
	e.Version = element.Attr("http://camunda.org/schema/1.0/bpmn", "versionTag")
	e.parsed = true
	return nil
}
