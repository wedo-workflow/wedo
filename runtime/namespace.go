package runtime

import (
	"context"

	"github.com/wedo-workflow/wedo/model"
)

func (r *Runtime) NamespaceCreate(ctx context.Context, namespace *model.Namespace) error {
	return r.store.NamespaceCreate(ctx, namespace)
}

func (r *Runtime) NamespaceGet(ctx context.Context, namespaceID string) (*model.Namespace, error) {
	return r.store.NamespaceGet(ctx, namespaceID)
}

// NamespaceList returns a list of namespaces
func (r *Runtime) NamespaceList(ctx context.Context, opts *model.NamespaceQueryOptions) ([]*model.Namespace, error) {
	return r.store.NamespaceList(ctx, opts)
}

// NamespaceListCount returns a count of namespaces
func (r *Runtime) NamespaceListCount(ctx context.Context, opts *model.NamespaceQueryOptions) (int64, error) {
	return r.store.NamespaceListCount(ctx, opts)
}

// NamespaceDelete deletes a namespace
func (r *Runtime) NamespaceDelete(ctx context.Context, namespaceID string) error {
	return r.store.NamespaceDelete(ctx, namespaceID)
}
