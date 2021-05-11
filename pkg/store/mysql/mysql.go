package mysql

import (
	"database/sql"
	"fmt"

	"github.com/wedo-workflow/wedo/configs"
	"github.com/wedo-workflow/wedo/pkg/store/metadata"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(config *configs.Config) (*MySQL, error) {
	return &MySQL{}, nil
}

func (m *MySQL) BulkDelete(req []metadata.DeleteRequest) error {
	panic("implement me")
}

func (m *MySQL) BulkGet(req []metadata.GetRequest) (bool, []metadata.BulkGetResponse, error) {
	panic("implement me")
}

func (m *MySQL) BulkSet(req []metadata.SetRequest) error {
	panic("implement me")
}

func (m *MySQL) Init(metadata metadata.Metadata) error {
	panic("implement me")
}

func (m *MySQL) Features() []metadata.Feature {
	panic("implement me")
}

func (m *MySQL) Delete(req *metadata.DeleteRequest) error {
	panic("implement me")
}

func (m *MySQL) Get(req *metadata.GetRequest) (*metadata.GetResponse, error) {
	panic("implement me")
}

func (m *MySQL) Set(req *metadata.SetRequest) error {
	panic("implement me")
}

func (m *MySQL) Ping() error {
	// return m.db.Ping()
	fmt.Println("It's OK")
	return nil
}
