package runtime

import (
	"os"
	"testing"

	"github.com/wedo-workflow/wedo/runtime/config"
)

func TestXML(t *testing.T) {
	doc, err := os.ReadFile("testdata/diagram.bpmn")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	wedo, err := NewRuntime(config.NewDefaultConfig())
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if err := wedo.Deploy(doc); err != nil {
		t.Log(err)
		t.FailNow()
	}
}
