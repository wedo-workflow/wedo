package redis

import (
	"context"

	"github.com/wedo-workflow/wedo/model"
)

func (r *Redis) NamespaceCreate(ctx context.Context, namespace *model.Namespace) error {
	if err := r.db.SAdd(ctx, NamespacesSet, namespace.Name).Err(); err != nil {
		return err
	}
	return r.db.HSet(ctx, Namespaces, namespace.ID, namespace).Err()
}

func (r *Redis) NamespaceCheckExist(ctx context.Context, namespace string) (bool, error) {
	return r.db.SIsMember(ctx, NamespacesSet, namespace).Result()
}

func (r *Redis) NamespaceGet(ctx context.Context, namespaceID string) (*model.Namespace, error) {
	var namespace model.Namespace
	err := r.db.HGet(ctx, Namespaces, namespaceID).Scan(&namespace)
	return &namespace, err
}

// NamespaceList returns all namespaces
func (r *Redis) NamespaceList(ctx context.Context, opts *model.NamespaceQueryOptions) ([]*model.Namespace, error) {
	// todo : deal with opts
	namespaces := make([]*model.Namespace, 0)
	results, err := r.db.HGetAll(ctx, Namespaces).Result()
	if err != nil {
		return nil, err
	}
	for _, result := range results {
		namespace := &model.Namespace{}
		if err := namespace.UnmarshalBinary([]byte(result)); err != nil {
			return nil, err
		}
		namespaces = append(namespaces, namespace)
	}
	return namespaces, nil
}

// NamespaceListCount returns the count of all namespaces
func (r *Redis) NamespaceListCount(ctx context.Context, opts *model.NamespaceQueryOptions) (int64, error) {
	// todo : deal with opts
	return r.db.HLen(ctx, Namespaces).Result()
}

// NamespaceDelete deletes a namespace
func (r *Redis) NamespaceDelete(ctx context.Context, namespaceID string) error {
	return r.db.HDel(ctx, Namespaces, namespaceID).Err()
}
