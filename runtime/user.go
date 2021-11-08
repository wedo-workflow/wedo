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

// UserDelete deletes a user by ID.
func (r *Runtime) UserDelete(ctx context.Context, id string) error {
	return r.store.UserDelete(ctx, id)
}

// UserList returns a list of users.
func (r *Runtime) UserList(ctx context.Context, opts *model.UserListOptions) ([]*model.User, error) {
	return r.store.UserList(ctx, opts)
}
