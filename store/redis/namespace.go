package redis

import (
	"context"
	"encoding/json"

	"github.com/wedo-workflow/wedo/model"
)

const (
	NamespaceAll = "namespace_all"
)

func (r *Redis) NamespaceCreate(ctx context.Context, namespace *model.Namespace) error {
	namespaceBytes, err := json.Marshal(namespace)
	if err != nil {
		return err
	}
	return r.db.HSet(ctx, NamespaceAll, namespace.ID, string(namespaceBytes)).Err()
}

func (r *Redis) NamespaceGet(ctx context.Context, namespaceID string) (*model.Namespace, error) {
	namespaceBytes, err := r.db.HGet(ctx, NamespaceAll, namespaceID).Bytes()
	if err != nil {
		return nil, err
	}
	var namespace model.Namespace
	if err := json.Unmarshal(namespaceBytes, &namespace); err != nil {
		return nil, err
	}
	return &namespace, nil
}

// NamespaceList returns all namespaces
func (r *Redis) NamespaceList(ctx context.Context, opts *model.NamespaceQueryOptions) ([]*model.Namespace, error) {
	// todo : deal with opts
	namespaces := make([]*model.Namespace, 0)
	results, err := r.db.HGetAll(ctx, NamespaceAll).Result()
	if err != nil {
		return nil, err
	}
	for _, result := range results {
		var namespace model.Namespace
		if err := json.Unmarshal([]byte(result), &namespace); err != nil {
			return nil, err
		}
		namespaces = append(namespaces, &namespace)
	}
	return namespaces, nil
}

// NamespaceListCount returns the count of all namespaces
func (r *Redis) NamespaceListCount(ctx context.Context, opts *model.NamespaceQueryOptions) (int64, error) {
	// todo : deal with opts
	return r.db.HLen(ctx, NamespaceAll).Result()
}

// NamespaceDelete deletes a namespace
func (r *Redis) NamespaceDelete(ctx context.Context, namespaceID string) error {
	return r.db.HDel(ctx, NamespaceAll, namespaceID).Err()
}
