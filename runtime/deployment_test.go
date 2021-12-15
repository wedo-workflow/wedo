package runtime

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/wedo-workflow/wedo/model"
)

func TestRuntime_Deploy(t *testing.T) {
	doc, err := os.ReadFile("testdata/invoice.v1.bpmn")
	if err != nil {
		t.Fatal(err)
	}
	rt := initRuntime()
	ns, err := rt.NamespaceGetByName(context.Background(), "default")
	if err != nil {
		t.Fatal(err)
	}
	deploy, err := initRuntime().Deploy(context.Background(), &model.Deployment{
		NamespaceID: ns.ID,
		Name:        "formal",
		Content:     doc,
		CreateTime:  time.Now(),
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(deploy)
}

func TestRuntime_DeployList(t *testing.T) {
	list, err := initRuntime().DeployList(context.Background(), &model.DeploymentListOptions{})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	for _, deployment := range list {
		t.Log(deployment)
	}
}

func TestRuntime_DeployDelete(t *testing.T) {
	err := initRuntime().DeployDelete(context.Background(), "69191eb5-dfa4-464a-9b6b-71451d34e933")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}
