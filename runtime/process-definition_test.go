package runtime

import (
	"context"
	"github.com/wedo-workflow/wedo/model"
	"testing"
)

func TestRuntime_ProcessDefinitionGet(t *testing.T) {
	rt := initRuntime()
	pd, err := rt.ProcessDefinitionGet(context.Background(), &model.ProcessDefinition{
		Id: "2d08f563-bfc4-4c45-bdd9-22f6afb6d14d",
		//ProcessDefinitionKey: request.ProcessDefinitionKey,
		//NamespaceId:          request.NamespaceId,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pd)
}

func TestRuntime_ProcessDefinitionStart(t *testing.T) {
	rt := initRuntime()
	pdi, err := rt.ProcessDefinitionStart(context.Background(), &model.ProcessDefinition{
		Id: "2d08f563-bfc4-4c45-bdd9-22f6afb6d14d",
		//ProcessDefinitionKey: request.ProcessDefinitionKey,
		//NamespaceId:          request.NamespaceId,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pdi)
}
