package runtime

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/wedo-workflow/wedo/model"
)

func (r *Runtime) ProcessDefinitionGet(ctx context.Context, pd *model.ProcessDefinition) (*model.Deployment, error) {
	pdID, err := r.store.ProcessDefinition(ctx, fmt.Sprintf(deployNameKey, pd.ProcessDefinitionKey))
	if err != nil {
		return nil, err
	}
	if pdID == "" {
		return nil, errors.New("not found process definition")
	}
	pdDetail, err := r.store.ProcessDefinition(ctx, pdID+"_detail")
	if err != nil {
		return nil, err
	}
	deploy := &model.Deployment{}
	if err := json.Unmarshal([]byte(pdDetail), deploy); err != nil {
		return nil, err
	}
	return deploy, nil
}

// GET /process-definition/invoice:1:c3a63aaa-2046-11e7-8f94-34f39ab71d4e

// ProcessDefinitionStart Request is the request body for starting a process.
func (r *Runtime) ProcessDefinitionStart(ctx context.Context, pd *model.ProcessDefinition) (string, error) {
	did, err := r.store.ProcessDefinition(ctx, fmt.Sprintf(deployNameKey, pd.ProcessDefinitionKey))
	if err != nil {
		return "", err
	}
	pd.Id = did
	return r.store.ProcessDefinitionStart(ctx, pd)
}
