package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Group struct {
	ID               string `json:"id"`
	CategoryValueRef string `json:"categoryValueRef"`

	TName  string `json:"type_name"` // Element's Type Name, "task" "process" etc.
	parsed bool
}

func NewGroup() *Group {
	return &Group{}
}

func (e *Group) EID() string {
	return e.ID
}

func (e *Group) TypeName() string {
	return e.TName
}

func (e *Group) SetTypeName(s string) error {
	e.TName = s
	return nil
}

func (e *Group) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.CategoryValueRef = element.Attr("", "categoryValueRef")
	e.parsed = true
	return nil
}
