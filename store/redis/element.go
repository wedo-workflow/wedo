package redis

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/wedo-workflow/wedo/element"
)

func (r *Redis) ElementSave(ctx context.Context, element element.Element, rootID string) error {
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
