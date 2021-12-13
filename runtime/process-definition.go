package runtime

import (
	"context"
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
	return r.store.ProcessDefinitionStart(ctx, pd)
}
