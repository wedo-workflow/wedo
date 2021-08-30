package runtime

import (
	"context"
	"os"
	"testing"

	"github.com/wedo-workflow/wedo/element/wedo_model"
	"github.com/wedo-workflow/wedo/runtime/config"
)

func TestXML(t *testing.T) {
	doc, err := os.ReadFile("testdata/diagram.bpmn")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	rt, err := NewRuntime(config.NewDefaultConfig())
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if err := rt.Deploy(context.Background(), &wedo_model.Deploy{
		DID:     "",
		Name:    "test",
		Content: doc,
	}); err != nil {
		t.Log(err)
		t.FailNow()
	}
}
