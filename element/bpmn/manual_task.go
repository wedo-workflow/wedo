package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type ManualTask struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewManualTask() *ManualTask {
	return &ManualTask{}
}

func (e *ManualTask) EID() string {
	return e.ID
}

func (e *ManualTask) TypeName() string {
	return e.TName
}

func (e *ManualTask) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *ManualTask) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	e.parsed = true
	return nil
}
