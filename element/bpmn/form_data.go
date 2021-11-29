package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type FormData struct {
	BusinessKey string `json:"businessKey"`

	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewFormData() *FormData {
	return &FormData{}
}

func (e *FormData) EID() string {
	return e.BusinessKey
}

func (e *FormData) TypeName() string {
	return e.TName
}

func (e *FormData) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *FormData) Parse(element *xmltree.Element) error {
	e.BusinessKey = element.Attr("", "businessKey")
	e.parsed = true
	return nil
}
