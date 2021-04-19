package bpmn_parse

import (
	"os"
	"testing"

	"github.com/wedo-workflow/wedo/pkg/bpmn-parse/xmltree"
)

func TestXML(t *testing.T) {
	f, err := os.ReadFile("testdata/shenpi.bpmn")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	e, err := xmltree.Parse(f)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(e)
}
