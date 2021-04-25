package mysql

import (
	"database/sql"
	"fmt"

	"github.com/wedo-workflow/wedo/configs"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(config *configs.Config) (*MySQL, error) {
	return &MySQL{}, nil
}

func (m *MySQL) Ping() error {
	// return m.db.Ping()
	fmt.Println("It's OK")
	return nil
}
