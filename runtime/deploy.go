package runtime

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/wedo-workflow/wedo/model"
	"github.com/wedo-workflow/xmltree"
)

var (
	parsersWhitelist = map[string]bool{
		"BPMNShape":        true,
		"Bounds":           true,
		"BPMNPlane":        true,
		"BPMNEdge":         true,
		"BPMNLabel":        true,
		"waypoint":         true,
		"exclusiveGateway": true,
	}
)

func (r *Runtime) Deploy(ctx context.Context, deploy *model.Deploy) (string, error) {
	uuidV4, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	deploy.DID = uuidV4.String()
	// 1. store process definition
	if err := r.store.ProcessDefinitionAdd(ctx, deploy); err != nil {
		return "", err
	}
	// 2. store deploy
	if err := r.store.DeploySet(ctx, deploy); err != nil {
		return "", err
	}
	// 3.0 parse deploy content
	tree, err := xmltree.Parse(deploy.Content)
	if err != nil {
		return "", fmt.Errorf("parse deploy content error: %s", err)
	}
	// 3.1 deploy element
	return deploy.DID, r.deploy(tree, deploy.DID)
}

func (r *Runtime) deploy(tree *xmltree.Element, rootID string) error {
	if err := r.parseAndStore(tree, rootID); err != nil {
		return err
	}
	if len(tree.Children) == 0 {
		return nil
	}
	for _, child := range tree.Children {
		fmt.Println("deal with", tree.Name.Local, child.Name.Local)
		if err := r.deploy(&child, rootID); err != nil {
			return err
		}
	}
	return nil
}

func (r *Runtime) parseAndStore(e *xmltree.Element, rootID string) error {
	fmt.Println(e.Name.Local, string(e.Content))
	eleLocal := e.Name.Local
	_, skip := parsersWhitelist[eleLocal]
	if skip {
		return nil
	}
	parser, ok := r.rootParsers[eleLocal]
	if !ok {
		return fmt.Errorf("element %s's parser not found", e.Name.Local)
	}
	// 1. Parse
	if err := parser.Parse(e); err != nil {
		return err
	}
	// 2. store
	if err := r.store.ElementSet(context.Background(), parser, rootID); err != nil {
		return err
	}
	fmt.Println(e.Name.Local, "saved")
	return nil
}

// DeployList returns a list of deployments
func (r *Runtime) DeployList(ctx context.Context, opts *model.DeploymentListOptions) ([]*model.Deploy, error) {
	return r.store.DeployList(ctx, opts)
}

// DeployDelete deletes a deployment
func (r *Runtime) DeployDelete(ctx context.Context, deploymentID string) error {
	return r.store.DeployDelete(ctx, deploymentID)
}
