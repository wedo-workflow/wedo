package postgresql

import "github.com/wedo-workflow/wedo/configs"

type PostgreSQL struct {
}

func NewPG(config *configs.Config) (*PostgreSQL, error) {
	return &PostgreSQL{}, nil

}

func (pg *PostgreSQL) Ping() error {
	return nil
}
