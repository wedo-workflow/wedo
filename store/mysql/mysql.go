package mysql

import (
	"database/sql"

	"github.com/wedo-workflow/wedo/runtime/config"
	metadata2 "github.com/wedo-workflow/wedo/store/metadata"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(config *config.Config) (*MySQL, error) {
	return &MySQL{}, nil
}

func (m *MySQL) BulkDelete(req []metadata2.DeleteRequest) error {
	panic("implement me")
}

func (m *MySQL) BulkGet(req []metadata2.GetRequest) (bool, []metadata2.BulkGetResponse, error) {
	panic("implement me")
}

func (m *MySQL) BulkSet(req []metadata2.SetRequest) error {
	panic("implement me")
}

func (m *MySQL) Init(metadata metadata2.Metadata) error {
	panic("implement me")
}

func (m *MySQL) Features() []metadata2.Feature {
	panic("implement me")
}

func (m *MySQL) Delete(req *metadata2.DeleteRequest) error {
	panic("implement me")
}

func (m *MySQL) Get(req *metadata2.GetRequest) (*metadata2.GetResponse, error) {
	panic("implement me")
}

func (m *MySQL) Set(req *metadata2.SetRequest) error {
	panic("implement me")
}

func (m *MySQL) Ping() error {
	// return m.db.Ping()
	// fmt.Println("It's OK")
	return nil
}
