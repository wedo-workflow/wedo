package postgresql

import (
	"github.com/wedo-workflow/wedo/configs"
	"github.com/wedo-workflow/wedo/pkg/store/metadata"
)

type PostgreSQL struct {
}

func NewPG(config *configs.Config) (*PostgreSQL, error) {
	return &PostgreSQL{}, nil

}

func (pg *PostgreSQL) BulkDelete(req []metadata.DeleteRequest) error {
	panic("implement me")
}

func (pg *PostgreSQL) BulkGet(req []metadata.GetRequest) (bool, []metadata.BulkGetResponse, error) {
	panic("implement me")
}

func (pg *PostgreSQL) BulkSet(req []metadata.SetRequest) error {
	panic("implement me")
}

func (pg *PostgreSQL) Init(metadata metadata.Metadata) error {
	panic("implement me")
}

func (pg *PostgreSQL) Features() []metadata.Feature {
	panic("implement me")
}

func (pg *PostgreSQL) Delete(req *metadata.DeleteRequest) error {
	panic("implement me")
}

func (pg *PostgreSQL) Get(req *metadata.GetRequest) (*metadata.GetResponse, error) {
	panic("implement me")
}

func (pg *PostgreSQL) Set(req *metadata.SetRequest) error {
	panic("implement me")
}

func (pg *PostgreSQL) Ping() error {
	return nil
}
