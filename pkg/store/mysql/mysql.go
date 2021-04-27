package mysql

import (
	"database/sql"
	"fmt"

	"github.com/wedo-workflow/wedo/configs"
	"github.com/wedo-workflow/wedo/pkg/store"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(config *configs.Config) (*MySQL, error) {
	return &MySQL{}, nil
}

func (m *MySQL) BulkDelete(req []store.DeleteRequest) error {
	panic("implement me")
}

func (m *MySQL) BulkGet(req []store.GetRequest) (bool, []store.BulkGetResponse, error) {
	panic("implement me")
}

func (m *MySQL) BulkSet(req []store.SetRequest) error {
	panic("implement me")
}

func (m *MySQL) Init(metadata store.Metadata) error {
	panic("implement me")
}

func (m *MySQL) Features() []store.Feature {
	panic("implement me")
}

func (m *MySQL) Delete(req *store.DeleteRequest) error {
	panic("implement me")
}

func (m *MySQL) Get(req *store.GetRequest) (*store.GetResponse, error) {
	panic("implement me")
}

func (m *MySQL) Set(req *store.SetRequest) error {
	panic("implement me")
}

func (m *MySQL) Ping() error {
	// return m.db.Ping()
	fmt.Println("It's OK")
	return nil
}
