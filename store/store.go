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

	ProcessDefinition(ctx context.Context, key string) (*model.ProcessDefinition, error)
	ProcessDefinitionAdd(ctx context.Context, pd *model.ProcessDefinition) error
	// ProcessDefinitionStart return the process instance id
	ProcessDefinitionStart(ctx context.Context, pi *model.ProcessInstance) error

	// ProcessInstanceGet return the process instance
	ProcessInstanceGet(ctx context.Context, id string) (*model.ProcessInstance, error)
	// ProcessInstanceUpdateStatus update the process instance status
	ProcessInstanceUpdateStatus(ctx context.Context, pi *model.ProcessInstance) error

	// Deployment read a deployment
	Deployment(ctx context.Context, deployID string) (*model.Deployment, error)
	DeploymentCreate(ctx context.Context, deploy *model.Deployment) error
	DeploymentList(ctx context.Context, opts *model.DeploymentListOptions) ([]*model.Deployment, error)
	DeploymentDelete(ctx context.Context, deployID string) error

	// Element(ctx context.Context, element element.Element) (element.Element, error)
	ElementSet(ctx context.Context, deploy *model.Deployment, element element.Element) error
	ElementSetAnchor(ctx context.Context, deploy *model.Deployment, element element.Element) error
	// ElementDelete(ctx context.Context, element element.Element) error
	// ElementsByRootID(ctx context.Context, rootID string) ([]element.Element, error)

	NamespaceCreate(ctx context.Context, namespace *model.Namespace) error
	NamespaceCheckExist(ctx context.Context, namespace string) (bool, error)
	NamespaceGetByID(ctx context.Context, namespaceID string) (*model.Namespace, error)
	NamespaceGetByName(ctx context.Context, namespaceName string) (*model.Namespace, error)
	NamespaceDelete(ctx context.Context, namespaceID string) error
	NamespaceList(ctx context.Context, opts *model.NamespaceQueryOptions) ([]*model.Namespace, error)
	NamespaceListCount(ctx context.Context, opts *model.NamespaceQueryOptions) (int64, error)

	UserCreate(ctx context.Context, user *model.User) error
	UserGet(ctx context.Context, userID string) (*model.User, error)
	UserCheckExist(ctx context.Context, email string) (bool, error)
	UserDelete(ctx context.Context, userID string) error
	UserUpdate(ctx context.Context, user *model.User) error
	UserList(ctx context.Context, opts *model.UserListOptions) ([]*model.User, error)

	// TaskCreate create a task
	TaskCreate(ctx context.Context, task *model.Task) error
	TaskGet(ctx context.Context, taskID string) (*model.Task, error)
	TaskDelete(ctx context.Context, taskID string, processInstanceID string) error
	TaskList(ctx context.Context, opts *model.TaskListOptions) ([]*model.Task, error)
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
