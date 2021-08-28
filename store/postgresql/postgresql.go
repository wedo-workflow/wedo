package postgresql

import (
	"github.com/wedo-workflow/wedo/runtime/config"
	metadata2 "github.com/wedo-workflow/wedo/store/metadata"
)

type PostgreSQL struct {
}

func NewPG(config *config.Config) (*PostgreSQL, error) {
	return &PostgreSQL{}, nil

}

func (pg *PostgreSQL) BulkDelete(req []metadata2.DeleteRequest) error {
	panic("implement me")
}

func (pg *PostgreSQL) BulkGet(req []metadata2.GetRequest) (bool, []metadata2.BulkGetResponse, error) {
	panic("implement me")
}

func (pg *PostgreSQL) BulkSet(req []metadata2.SetRequest) error {
	panic("implement me")
}

func (pg *PostgreSQL) Init(metadata metadata2.Metadata) error {
	panic("implement me")
}

func (pg *PostgreSQL) Features() []metadata2.Feature {
	panic("implement me")
}

func (pg *PostgreSQL) Delete(req *metadata2.DeleteRequest) error {
	panic("implement me")
}

func (pg *PostgreSQL) Get(req *metadata2.GetRequest) (*metadata2.GetResponse, error) {
	panic("implement me")
}

func (pg *PostgreSQL) Set(req *metadata2.SetRequest) error {
	panic("implement me")
}

func (pg *PostgreSQL) Ping() error {
	return nil
}
