package store

import (
	"github.com/wedo-workflow/wedo/configs"
	"github.com/wedo-workflow/wedo/pkg/store/mysql"
	"github.com/wedo-workflow/wedo/pkg/store/postgresql"
)

// Store is an interface to perform operations on store
type Store interface {
	BulkStore
	Init(metadata Metadata) error
	Features() []Feature
	Delete(req *DeleteRequest) error
	Get(req *GetRequest) (*GetResponse, error)
	Set(req *SetRequest) error

	Ping() error
}

// BulkStore is an interface to perform bulk operations on store
type BulkStore interface {
	BulkDelete(req []DeleteRequest) error
	BulkGet(req []GetRequest) (bool, []BulkGetResponse, error)
	BulkSet(req []SetRequest) error
}

// TransactionalStore is an interface for initialization and support multiple transactional requests
type TransactionalStore interface {
	Init(metadata Metadata) error
	Multi(request *TransactionalStateRequest) error
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
