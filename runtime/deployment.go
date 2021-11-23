package runtime

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/wedo-workflow/wedo/element"
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

func (r *Runtime) Deploy(ctx context.Context, deploy *model.Deployment) (string, error) {
	if deploy.NamespaceID == "" {
		return "", errors.New("namespace id is empty")
	}
	uuidV4, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	deploy.DID = uuidV4.String()
	// 0. store deployment
	if err := r.store.DeploymentCreate(ctx, deploy); err != nil {
		return "", err
	}

	// 1. store process definition
	// if err := r.store.ProcessDefinitionAdd(ctx, deploy); err != nil {
	// 	return "", err
	// }

	// 3.0 parse deploy content
	tree, err := xmltree.Parse(deploy.Content)
	if err != nil {
		return "", fmt.Errorf("parse deploy content error: %s", err)
	}
	// 3.1 deploy element
	parsers := r.RootParsers()
	if err := r.deploy(ctx, tree, parsers, deploy.DID); err != nil {
		return "", err
	}
	return deploy.DID, nil
}

func (r *Runtime) deploy(ctx context.Context, tree *xmltree.Element, parsers map[string]element.Element, rootID string) error {
	newRootID, err := r.parseAndStore(ctx, tree, parsers, rootID)
	if err != nil {
		return err
	}
	if len(tree.Children) == 0 {
		return nil
	}
	for _, child := range tree.Children {
		fmt.Println("deal with", tree.Name.Local, child.Name.Local)
		if err := r.deploy(ctx, &child, parsers, newRootID); err != nil {
			return err
		}
	}
	return nil
}

func (r *Runtime) parseAndStore(ctx context.Context, e *xmltree.Element, parsers map[string]element.Element, rootID string) (string, error) {
	log.Debug("start to deal", rootID, e.Name.Local)
	eleLocal := e.Name.Local
	_, skip := parsersWhitelist[eleLocal]
	if skip {
		return "", nil
	}
	ele, ok := parsers[eleLocal]
	if !ok {
		return "", fmt.Errorf("ele %s's parser not found", e.Name.Local)
	}
	if err := ele.SetRootID(rootID); err != nil {
		return "", err
	}
	// 1. Parse
	if err := ele.Parse(e); err != nil {
		return "", err
	}
	// 2. store
	if err := r.store.ElementSet(ctx, ele, rootID); err != nil {
		return "", err
	}
	log.Debug(e.Name.Local, "saved")
	return ele.EID(), nil
}

// DeployList returns a list of deployments
func (r *Runtime) DeployList(ctx context.Context, opts *model.DeploymentListOptions) ([]*model.Deployment, error) {
	return r.store.DeploymentList(ctx, opts)
}

// DeployDelete deletes a deployment
func (r *Runtime) DeployDelete(ctx context.Context, deploymentID string) error {
	return r.store.DeploymentDelete(ctx, deploymentID)
}
