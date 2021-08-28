package store

import (
	"reflect"
	"testing"

	"github.com/wedo-workflow/wedo/runtime/config"
)

func TestNewStore(t *testing.T) {
	type args struct {
		config *config.Config
	}
	tests := []struct {
		name    string
		args    args
		want    Store
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStore(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStore() got = %v, want %v", got, tt.want)
			}
		})
	}
}
