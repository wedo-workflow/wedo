package runtime

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/wedo-workflow/wedo/model"
)

func TestRuntime_Deploy(t *testing.T) {
	doc, err := os.ReadFile("testdata/diagram.BPMN")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	deployID, err := initRuntime().Deploy(context.Background(), &model.Deployment{
		NamespaceID: "dffd661d-978d-4ed8-803d-a54ac0fed8ee",
		Name:        "test",
		Content:     doc,
		CreateTime:  time.Now(),
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(deployID)
}
