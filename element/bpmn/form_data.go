package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type FormData struct {
	BusinessKey string `json:"businessKey"`

	RID    string `json:"rid"` // Root element id
	parsed bool
}

func NewFormData() *FormData {
	return &FormData{}
}

func (e *FormData) EID() string {
	return e.BusinessKey
}

func (e *FormData) RootID() string {
	return e.RID
}

func (e *FormData) SetRootID(s string) error {
	e.RID = s
	return nil
}

func (e *FormData) Parse(element *xmltree.Element) error {
	e.BusinessKey = element.Attr("", "businessKey")
	e.parsed = true
	return nil
}
