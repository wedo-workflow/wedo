package bpmn

import (
	"os"
	"testing"
)

func TestXML(t *testing.T) {
	doc, err := os.ReadFile("testdata/diagram.bpmn")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	b := NewB()
	b.ParseDoc(doc)
}
