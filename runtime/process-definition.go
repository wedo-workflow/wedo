package runtime

import (
	"context"
	wedo "github.com/wedo-workflow/wedo/api"
	"github.com/wedo-workflow/wedo/model"
)

func (r *Runtime) ProcessDefinitionGet(ctx context.Context, pd *model.ProcessDefinition) (*model.ProcessDefinition, error) {
	pd, err := r.store.ProcessDefinition(ctx, pd.Id)
	if err != nil {
		return nil, err
	}
	return pd, nil
}

// ProcessDefinitionStart Request is the request body for starting a process.
func (r *Runtime) ProcessDefinitionStart(ctx context.Context, pd *model.ProcessDefinition) (string, error) {
	pi := &model.ProcessInstance{
		Id:                r.generateUUID(),
		Status:            wedo.ProcessInstanceState_STARTED,
		ProcessDefinition: pd,
	}
	return pi.Id, r.store.ProcessDefinitionStart(ctx, pi)
}
