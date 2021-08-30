package bpmn

import (
	"github.com/wedo-workflow/xmltree"
)

type Group struct {
	ID               string `json:"id"`
	CategoryValueRef string `json:"categoryValueRef"`

	parsed bool
}

func NewGroup() *Group {
	return &Group{}
}

func (e *Group) EID() string {
	return e.ID
}

func (e *Group) Parse(element *xmltree.Element) error {
	e.ID = element.Attr("", "id")
	e.CategoryValueRef = element.Attr("", "categoryValueRef")
	e.parsed = true
	return nil
}
