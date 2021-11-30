package runtime

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/wedo-workflow/wedo/element"
	"github.com/wedo-workflow/wedo/runtime/config"
	"github.com/wedo-workflow/wedo/store"
)

type Runtime struct {
	store   store.Store
	parsers map[string]element.Element
}

func NewRuntime(config *config.Config) (*Runtime, error) {
	newStore, err := store.NewStore(config)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.Store.PingTimeout*time.Second)
	defer cancel()
	if err := newStore.Ping(ctx); err != nil {
		return nil, err
	}
	r := &Runtime{
		store:   newStore,
		parsers: element.DefaultRegister(),
	}

	return r, nil
}

func (r *Runtime) RootParsers() map[string]element.Element {
	return r.parsers
}

func (r *Runtime) Store() store.Store {
	return r.store
}

func (r *Runtime) Run(opts ...Option) error {

	return nil
}

func (r *Runtime) generateUUID() string {
	v4, err := uuid.NewRandom()
	if err != nil {
		return ""
	}
	return v4.String()
}
