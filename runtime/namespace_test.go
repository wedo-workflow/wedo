package runtime

import (
	"context"
	"testing"

	"github.com/wedo-workflow/wedo/model"
)

func TestRuntime_NamespaceCreate(t *testing.T) {
	np, err := initRuntime().NamespaceCreate(context.Background(), &model.Namespace{
		Name: "test",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(np)
}
