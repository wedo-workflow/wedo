package redis

import (
	metadata2 "github.com/wedo-workflow/wedo/store/metadata"
)

type Redis struct{}

func NewRedis() (*Redis, error) {
	return &Redis{}, nil
}

func (r *Redis) BulkDelete(req []metadata2.DeleteRequest) error {
	panic("implement me")
}

func (r *Redis) BulkGet(req []metadata2.GetRequest) (bool, []metadata2.BulkGetResponse, error) {
	panic("implement me")
}

func (r *Redis) BulkSet(req []metadata2.SetRequest) error {
	panic("implement me")
}

func (r *Redis) Init(metadata metadata2.Metadata) error {
	panic("implement me")
}

func (r *Redis) Features() []metadata2.Feature {
	panic("implement me")
}

func (r *Redis) Delete(req *metadata2.DeleteRequest) error {
	panic("implement me")
}

func (r *Redis) Get(req *metadata2.GetRequest) (*metadata2.GetResponse, error) {
	panic("implement me")
}

func (r *Redis) Set(req *metadata2.SetRequest) error {
	panic("implement me")
}

func (r *Redis) Ping() error {
	panic("implement me")
}
