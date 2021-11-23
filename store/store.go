package store

import (
	"context"

	"github.com/wedo-workflow/wedo/element"
	"github.com/wedo-workflow/wedo/model"
	"github.com/wedo-workflow/wedo/runtime/config"
	"github.com/wedo-workflow/wedo/store/redis"
)

type Store interface {
	Ping(ctx context.Context) error

	ProcessDefinition(ctx context.Context, processID string) (*model.Deploy, error)
	ProcessDefinitionAdd(ctx context.Context, deploy *model.Deploy) error
	// ProcessDefinitionStart return the process instance id
	ProcessDefinitionStart(ctx context.Context, pd *model.ProcessDefinition) (string, error)

	Deploy(ctx context.Context, deployID string) (*model.Deploy, error)
	DeploySet(ctx context.Context, deploy *model.Deploy) error
	DeployList(ctx context.Context, opts *model.DeploymentListOptions) ([]*model.Deploy, error)
	DeployDelete(ctx context.Context, deployID string) error

	Element(ctx context.Context, element element.Element) (element.Element, error)
	ElementSet(ctx context.Context, element element.Element, rootID string) error
	ElementDelete(ctx context.Context, element element.Element) error
	ElementsByRootID(ctx context.Context, rootID string) ([]element.Element, error)

	NamespaceCreate(ctx context.Context, namespace *model.Namespace) error
	NamespaceCheckExist(ctx context.Context, namespace string) (bool, error)
	NamespaceGet(ctx context.Context, namespaceID string) (*model.Namespace, error)
	NamespaceDelete(ctx context.Context, namespaceID string) error
	NamespaceList(ctx context.Context, opts *model.NamespaceQueryOptions) ([]*model.Namespace, error)
	NamespaceListCount(ctx context.Context, opts *model.NamespaceQueryOptions) (int64, error)

	UserCreate(ctx context.Context, user *model.User) error
	UserGet(ctx context.Context, userID string) (*model.User, error)
	UserCheckExist(ctx context.Context, email string) (bool, error)
	UserDelete(ctx context.Context, userID string) error
	UserUpdate(ctx context.Context, user *model.User) error
	UserList(ctx context.Context, opts *model.UserListOptions) ([]*model.User, error)
}

func NewStore(config *config.Config) (Store, error) {
	switch config.Store.Driver {
	case "redis":
		return redis.NewRedis(config)
	case "mysql":
		// return mysql.NewMySQL(config)
	case "postgresql":
		// return postgresql.NewPG(config)
	}
	return nil, nil
}
