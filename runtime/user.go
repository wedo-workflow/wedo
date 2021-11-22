package runtime

import (
	"context"
	"errors"

	"github.com/wedo-workflow/wedo/model"
)

var (
	ErrUserExist = errors.New("user email already exist")
)

func (r *Runtime) UserCreate(ctx context.Context, user *model.User) error {
	exist, err := r.store.UserCheckExist(ctx, user.Email)
	if err != nil {
		return err
	}
	if exist {
		return ErrUserExist
	}
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

// UserUpdate updates a user.
func (r *Runtime) UserUpdate(ctx context.Context, user *model.User) error {
	oldUser, err := r.store.UserGet(ctx, user.Id)
	if err != nil {
		return err
	}
	user.Password = oldUser.Password
	user.Email = oldUser.Email
	return r.store.UserUpdate(ctx, user)
}
