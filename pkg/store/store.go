package store

import (
	"github.com/wedo-workflow/wedo/configs"
	"github.com/wedo-workflow/wedo/pkg/store/mysql"
	"github.com/wedo-workflow/wedo/pkg/store/postgresql"
)

type Store interface {
	Ping() error
}

func NewStore(config *configs.Config) (Store, error) {
	switch config.StoreDriver {
	case "mysql":
		return mysql.NewMySQL(config)
	case "postgresql":
		return postgresql.NewPG(config)
	}
	return nil, nil
}
