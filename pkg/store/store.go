package store

import (
	"github.com/wedo-workflow/wedo/configs"
	"github.com/wedo-workflow/wedo/pkg/store/metadata"
	"github.com/wedo-workflow/wedo/pkg/store/mysql"
	"github.com/wedo-workflow/wedo/pkg/store/postgresql"
)

// Store is an interface to perform operations on store
type Store interface {
	BulkStore
	Init(metadata metadata.Metadata) error
	Features() []metadata.Feature
	Delete(req *metadata.DeleteRequest) error
	Get(req *metadata.GetRequest) (*metadata.GetResponse, error)
	Set(req *metadata.SetRequest) error

	Ping() error
}

// BulkStore is an interface to perform bulk operations on store
type BulkStore interface {
	BulkDelete(req []metadata.DeleteRequest) error
	BulkGet(req []metadata.GetRequest) (bool, []metadata.BulkGetResponse, error)
	BulkSet(req []metadata.SetRequest) error
}

// TransactionalStore is an interface for initialization and support multiple transactional requests
type TransactionalStore interface {
	Init(metadata metadata.Metadata) error
	Multi(request *metadata.TransactionalStateRequest) error
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
