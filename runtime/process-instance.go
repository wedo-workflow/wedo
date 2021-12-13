package runtime

import (
	"context"
	"github.com/wedo-workflow/wedo/model"
)

func (r *Runtime) ProcessInstanceGet(ctx context.Context, pi *model.ProcessInstance) (*model.ProcessInstance, error) {
	return r.store.ProcessInstanceGet(ctx, pi.Id)
}
