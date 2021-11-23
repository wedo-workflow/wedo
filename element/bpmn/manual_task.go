package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type ManualTask struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	RID    string `json:"rid"` // Root element id
	parsed bool
}

func NewManualTask() *ManualTask {
	return &ManualTask{}
}

func (e *ManualTask) EID() string {
	return e.ID
}

func (e *ManualTask) RootID() string {
	return e.RID
}

func (e *ManualTask) SetRootID(s string) error {
	e.RID = s
	return nil
}

func (e *ManualTask) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.Name = element.Attr("", "name")
	e.parsed = true
	return nil
}
