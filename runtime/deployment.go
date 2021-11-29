package runtime

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
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
		"properties":       true,
		"property":         true,
	}

	anchorElements = map[string]bool{
		"startEvent": true,
		"endEvent":   true,
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
	if err := r.deploy(ctx, deploy, tree); err != nil {
		return "", err
	}
	return deploy.DID, nil
}

func (r *Runtime) deploy(ctx context.Context, deploy *model.Deployment, tree *xmltree.Element) error {
	_, err := r.parseAndStore(ctx, deploy, tree)
	if err != nil {
		return err
	}
	if len(tree.Children) == 0 {
		return nil
	}
	for _, child := range tree.Children {
		fmt.Println("deal with", tree.Name.Local, child.Name.Local)
		if err := r.deploy(ctx, deploy, &child); err != nil {
			return err
		}
	}
	return nil
}

func (r *Runtime) parseAndStore(ctx context.Context, deploy *model.Deployment, e *xmltree.Element) (string, error) {
	log.Debug("start to deal", deploy.DID, deploy.NamespaceID, deploy.Name, e.Name.Local)
	eleLocal := e.Name.Local
	_, skip := parsersWhitelist[eleLocal]
	if skip {
		return "", nil
	}
	ele, ok := r.RootParsers()[eleLocal]
	if !ok {
		return "", fmt.Errorf("ele %s's parser not found", e.Name.Local)
	}

	if err := ele.SetTypeName(eleLocal); err != nil {
		return "", err
	}

	// 1. Parse
	if err := ele.Parse(e); err != nil {
		return "", err
	}

	// 2. store
	if err := r.store.ElementSet(ctx, deploy, ele); err != nil {
		return "", err
	}

	if _, ok := anchorElements[eleLocal]; ok {
		if err := r.store.ElementSetAnchor(ctx, deploy, ele); err != nil {
			return "", err
		}
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
