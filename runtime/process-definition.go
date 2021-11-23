package runtime

import (
	"context"

	"github.com/wedo-workflow/wedo/model"
)

func (r *Runtime) ProcessDiagram(ctx context.Context, processID string) (*model.Deploy, error) {
	return r.store.ProcessDefinition(ctx, processID)
}

// GET /process-definition/invoice:1:c3a63aaa-2046-11e7-8f94-34f39ab71d4e

// ProcessDefinitionStart Request is the request body for starting a process.
func (r *Runtime) ProcessDefinitionStart(ctx context.Context, pf *model.ProcessDefinition) (string, error) {
	return r.store.ProcessDefinitionStart(ctx, pf)
}
