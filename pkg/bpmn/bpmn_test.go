package bpmn

import (
	"os"
	"testing"

	"github.com/wedo-workflow/wedo/configs"
)

func TestXML(t *testing.T) {
	doc, err := os.ReadFile("testdata/diagram.bpmn")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	wedo, err := NewWedo(configs.NewDefaultConfig())
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	wedo.ParseDoc(doc)
}
