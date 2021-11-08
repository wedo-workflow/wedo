package runtime

import (
	"context"

	"github.com/wedo-workflow/wedo/model"
)

func (r *Runtime) UserCreate(ctx context.Context, user *model.User) error {
	return r.store.UserCreate(ctx, user)
}

// UserGet ByID returns a user by ID.
func (r *Runtime) UserGet(ctx context.Context, id string) (*model.User, error) {
	return r.store.UserGet(ctx, id)
}
