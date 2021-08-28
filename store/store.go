package store

import (
	"github.com/wedo-workflow/wedo/runtime/config"
	metadata2 "github.com/wedo-workflow/wedo/store/metadata"
	"github.com/wedo-workflow/wedo/store/mysql"
	"github.com/wedo-workflow/wedo/store/postgresql"
)

// Store is an interface to perform operations on store
type Store interface {
	BulkStore
	Init(metadata metadata2.Metadata) error
	Features() []metadata2.Feature
	Delete(req *metadata2.DeleteRequest) error
	Get(req *metadata2.GetRequest) (*metadata2.GetResponse, error)
	Set(req *metadata2.SetRequest) error

	Ping() error
}

// BulkStore is an interface to perform bulk operations on store
type BulkStore interface {
	BulkDelete(req []metadata2.DeleteRequest) error
	BulkGet(req []metadata2.GetRequest) (bool, []metadata2.BulkGetResponse, error)
	BulkSet(req []metadata2.SetRequest) error
}

// TransactionalStore is an interface for initialization and support multiple transactional requests
type TransactionalStore interface {
	Init(metadata metadata2.Metadata) error
	Multi(request *metadata2.TransactionalStateRequest) error
}

func NewStore(config *config.Config) (Store, error) {
	switch config.StoreDriver {
	case "mysql":
		return mysql.NewMySQL(config)
	case "postgresql":
		return postgresql.NewPG(config)
	}
	return nil, nil
}
