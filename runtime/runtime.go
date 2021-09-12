package runtime

import (
	"context"
	"time"

	"github.com/wedo-workflow/wedo/element"
	"github.com/wedo-workflow/wedo/runtime/config"
	"github.com/wedo-workflow/wedo/store"
)

type Runtime struct {
	rootParsers map[string]element.Element

	store store.Store
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
		store:       newStore,
		rootParsers: element.DefaultRegister(),
	}

	return r, nil
}

func (r *Runtime) Run(opts ...Option) error {

	return nil
}
