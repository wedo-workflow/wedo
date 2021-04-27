package postgresql

import (
	"github.com/wedo-workflow/wedo/configs"
	"github.com/wedo-workflow/wedo/pkg/store"
)

type PostgreSQL struct {
}

func NewPG(config *configs.Config) (*PostgreSQL, error) {
	return &PostgreSQL{}, nil

}

func (pg *PostgreSQL) BulkDelete(req []store.DeleteRequest) error {
	panic("implement me")
}

func (pg *PostgreSQL) BulkGet(req []store.GetRequest) (bool, []store.BulkGetResponse, error) {
	panic("implement me")
}

func (pg *PostgreSQL) BulkSet(req []store.SetRequest) error {
	panic("implement me")
}

func (pg *PostgreSQL) Init(metadata store.Metadata) error {
	panic("implement me")
}

func (pg *PostgreSQL) Features() []store.Feature {
	panic("implement me")
}

func (pg *PostgreSQL) Delete(req *store.DeleteRequest) error {
	panic("implement me")
}

func (pg *PostgreSQL) Get(req *store.GetRequest) (*store.GetResponse, error) {
	panic("implement me")
}

func (pg *PostgreSQL) Set(req *store.SetRequest) error {
	panic("implement me")
}

func (pg *PostgreSQL) Ping() error {
	return nil
}
