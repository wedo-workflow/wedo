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

	Deploy(ctx context.Context, deployID string) (*model.Deploy, error)
	DeploySet(ctx context.Context, deploy *model.Deploy) error

	Element(ctx context.Context, element element.Element) (element.Element, error)
	ElementSet(ctx context.Context, element element.Element, rootID string) error
	ElementDelete(ctx context.Context, element element.Element) error
	ElementsByRootID(ctx context.Context, rootID string) ([]element.Element, error)
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