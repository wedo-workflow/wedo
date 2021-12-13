package runtime

import (
	"context"
	"github.com/wedo-workflow/wedo/model"
	"testing"
)

func TestRuntime_ProcessDefinitionGet(t *testing.T) {
	rt := initRuntime()
	pd, err := rt.ProcessDefinitionGet(context.Background(), &model.ProcessDefinition{
		Id: "944e5937-ee35-4a9a-9c5a-f77c37a9a7f0",
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
		Id: "944e5937-ee35-4a9a-9c5a-f77c37a9a7f0",
		//ProcessDefinitionKey: request.ProcessDefinitionKey,
		//NamespaceId:          request.NamespaceId,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pdi)
}
