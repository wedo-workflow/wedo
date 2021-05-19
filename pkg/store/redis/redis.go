package redis

import (
	"github.com/wedo-workflow/wedo/pkg/store/metadata"
)

type Redis struct{}

func NewRedis() (*Redis, error) {
	return &Redis{}, nil
}

func (r *Redis) BulkDelete(req []metadata.DeleteRequest) error {
	panic("implement me")
}

func (r *Redis) BulkGet(req []metadata.GetRequest) (bool, []metadata.BulkGetResponse, error) {
	panic("implement me")
}

func (r *Redis) BulkSet(req []metadata.SetRequest) error {
	panic("implement me")
}

func (r *Redis) Init(metadata metadata.Metadata) error {
	panic("implement me")
}

func (r *Redis) Features() []metadata.Feature {
	panic("implement me")
}

func (r *Redis) Delete(req *metadata.DeleteRequest) error {
	panic("implement me")
}

func (r *Redis) Get(req *metadata.GetRequest) (*metadata.GetResponse, error) {
	panic("implement me")
}

func (r *Redis) Set(req *metadata.SetRequest) error {
	panic("implement me")
}

func (r *Redis) Ping() error {
	panic("implement me")
}
