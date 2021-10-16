package runtime

import (
	"context"
	"os"
	"testing"

	"github.com/wedo-workflow/wedo/model"
	"github.com/wedo-workflow/wedo/runtime/config"
)

func TestXML(t *testing.T) {
	doc, err := os.ReadFile("testdata/diagram.BPMN")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	rt, err := NewRuntime(config.NewDefaultConfig())
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if err := rt.Deploy(context.Background(), &model.Deploy{
		DID:     "",
		Name:    "test",
		Content: doc,
	}); err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log("test2")
}