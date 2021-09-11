package runtime

import (
	"context"
	"fmt"
	"time"

	"github.com/wedo-workflow/wedo/element"
	"github.com/wedo-workflow/wedo/model"
	"github.com/wedo-workflow/wedo/runtime/config"
	"github.com/wedo-workflow/wedo/store"
	"github.com/wedo-workflow/xmltree"

	"github.com/google/uuid"
)

type Runtime struct {
	rootParsers map[string]element.Element

	store store.Store
}

func NewRuntime(config *config.Config) (*Runtime, error) {
	newStore, err := store.NewStore(config)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), config.Store.PingTimeout*time.Second)
	defer cancel()
	if err := newStore.Ping(ctx); err != nil {
		return nil, err
	}
	r := &Runtime{
		store:       newStore,
		rootParsers: element.DefaultRegister(),
	}

	return r, nil
}

var (
	parsersWhitelist = map[string]bool{
		"BPMNShape": true,
		"Bounds":    true,
		"BPMNPlane": true,
		"BPMNEdge":  true,
		"BPMNLabel": true,
		"waypoint":  true,
	}
)

func (r *Runtime) Run(opts ...Option) error {

	return nil
}

func (r *Runtime) Deploy(ctx context.Context, deploy *model.Deploy) error {
	uuidV4, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	deploy.DID = uuidV4.String()
	if err := r.store.ProcessDefinitionAdd(ctx, deploy); err != nil {
		return err
	}
	if err := r.store.DeploySet(ctx, deploy); err != nil {
		return err
	}
	tree, err := xmltree.Parse(deploy.Content)
	if err != nil {
		return err
	}

	return r.deploy(tree, deploy.DID)
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
