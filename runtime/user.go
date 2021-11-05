package runtime

import (
	"context"

	"github.com/wedo-workflow/wedo/model"
)

func (r *Runtime) UserCreate(ctx context.Context, user *model.User) error {
	return r.store.UserCreate(ctx, user)
}
