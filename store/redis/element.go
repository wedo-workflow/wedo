package redis

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/wedo-workflow/wedo/element"
)

func (r *Redis) Element(ctx context.Context, element element.Element) (element.Element, error) {
	if element == nil {
		return nil, errors.New("element is nil")
	}
	if element.RootID() == "" {
		return nil, errors.New("rootID is empty")
	}
	if element.EID() == "" {
		return nil, errors.New("elementID is empty")
	}
	if err := r.db.HGet(ctx, element.RootID(), element.EID()).Scan(element); err != nil {
		return nil, err
	}
	return element, nil
}

func (r *Redis) ElementSet(ctx context.Context, element element.Element, rootID string) error {
	if element == nil {
		return errors.New("element is nil")
	}
	if rootID == "" {
		return errors.New("rootID is empty")
	}
	elementBytes, err := json.Marshal(element)
	if err != nil {
		return err
	}
	if err := r.db.HSet(ctx, rootID, element.EID(), elementBytes).Err(); err != nil {
		return err
	}
	return nil
}

func (r Redis) ElementDelete(ctx context.Context, element element.Element) error {
	if element == nil {
		return errors.New("element is nil")
	}
	if element.RootID() == "" || element.EID() == "" {
		return nil
	}
	return r.db.HDel(ctx, element.RootID(), element.EID()).Err()
}

func (r *Redis) ElementsByRootID(ctx context.Context, rootID string) ([]element.Element, error) {
	if rootID == "" {
		return nil, errors.New("rootID is empty")
	}
	elements := []element.Element{}
	if err := r.db.HGetAll(ctx, rootID).Scan(elements); err != nil {
		return nil, err
	}
	return elements, nil
}
