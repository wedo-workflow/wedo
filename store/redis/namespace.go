package redis

import (
	"context"

	"github.com/wedo-workflow/wedo/model"
)

func (r *Redis) NamespaceCreate(ctx context.Context, namespace *model.Namespace) error {
	if err := r.db.SAdd(ctx, NamespacesNameUniqueSet, namespace.Name).Err(); err != nil {
		return err
	}
	if err := r.db.HSet(ctx, NamespacesByName, namespace.Name, namespace).Err(); err != nil {
		return err
	}
	return r.db.HSet(ctx, NamespacesByID, namespace.ID, namespace).Err()
}

func (r *Redis) NamespaceCheckExist(ctx context.Context, namespace string) (bool, error) {
	return r.db.SIsMember(ctx, NamespacesNameUniqueSet, namespace).Result()
}

func (r *Redis) NamespaceGetByID(ctx context.Context, namespaceID string) (*model.Namespace, error) {
	var namespace model.Namespace
	err := r.db.HGet(ctx, NamespacesByID, namespaceID).Scan(&namespace)
	return &namespace, err
}

func (r *Redis) NamespaceGetByName(ctx context.Context, namespaceName string) (*model.Namespace, error) {
	var namespace model.Namespace
	err := r.db.HGet(ctx, NamespacesByName, namespaceName).Scan(&namespace)
	return &namespace, err
}

// NamespaceList returns all namespaces
func (r *Redis) NamespaceList(ctx context.Context, opts *model.NamespaceQueryOptions) ([]*model.Namespace, error) {
	// todo : deal with opts
	namespaces := make([]*model.Namespace, 0)
	results, err := r.db.HGetAll(ctx, NamespacesByID).Result()
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
	return r.db.HLen(ctx, NamespacesByID).Result()
}

// NamespaceDelete deletes a namespace
func (r *Redis) NamespaceDelete(ctx context.Context, namespaceID string) error {
	return r.db.HDel(ctx, NamespacesByID, namespaceID).Err()
}
