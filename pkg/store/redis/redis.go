package redis

import (
	"github.com/wedo-workflow/wedo/pkg/store"
)

type Redis struct{}

func NewRedis() (*Redis, error) {
	return &Redis{}, nil
}

func (r *Redis) BulkDelete(req []store.DeleteRequest) error {
	panic("implement me")
}

func (r *Redis) BulkGet(req []store.GetRequest) (bool, []store.BulkGetResponse, error) {
	panic("implement me")
}

func (r *Redis) BulkSet(req []store.SetRequest) error {
	panic("implement me")
}

func (r *Redis) Init(metadata store.Metadata) error {
	panic("implement me")
}

func (r *Redis) Features() []store.Feature {
	panic("implement me")
}

func (r *Redis) Delete(req *store.DeleteRequest) error {
	panic("implement me")
}

func (r *Redis) Get(req *store.GetRequest) (*store.GetResponse, error) {
	panic("implement me")
}

func (r *Redis) Set(req *store.SetRequest) error {
	panic("implement me")
}

func (r *Redis) Ping() error {
	panic("implement me")
}
