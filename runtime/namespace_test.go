package runtime

import (
	"context"
	"testing"

	"github.com/wedo-workflow/wedo/model"
)

func TestRuntime_NamespaceCreate(t *testing.T) {
	np, err := initRuntime().NamespaceCreate(context.Background(), &model.Namespace{
		Name: "default",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(np.ID, np.Name)
}

func TestRuntime_NamespaceGet(t *testing.T) {
	rt := initRuntime()
	np, err := rt.NamespaceGetByName(context.Background(), "default")
	if err != nil {
		t.Fatal(err)
	}
	np2, err := rt.NamespaceGetByID(context.Background(), np.ID)
	if err != nil {
		t.Fatal(err)
	}
	if np.ID != np2.ID {
		t.Fatal("namespace id not equal")
	}
	t.Log(np.ID, np.Name)
}

func TestRuntime_NamespaceDelete(t *testing.T) {

}

func TestRuntime_NamespaceList(t *testing.T) {

}

func TestRuntime_NamespaceListCount(t *testing.T) {

}
