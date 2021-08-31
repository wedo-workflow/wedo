package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/wedo-workflow/wedo/runtime/config"
)

// Config of service.
type Config struct {
	LogLevel      string         `json:"log_level"`
	GRPCEndpoint  *Endpoint      `json:"grpc_endpoint"`
	HTTPEndpoint  *Endpoint      `json:"http_endpoint"`
	RuntimeConfig *config.Config `json:"runtime_config"`
}

type Endpoint struct {
	Address string `json:"address"`
	Port    uint16 `json:"port"`
}

func (e *Endpoint) String() string {
	return fmt.Sprintf("%v:%v", e.Address, e.Port)
}

func LoadConfig(configPath string) (*Config, error) {
	if configPath == "" {
		return nil, errors.New("file path empty")
	}

	data, err := ioutil.ReadFile(configPath) // nolint: gosec
	if err != nil {
		return nil, fmt.Errorf("open file(%v) with err %v", configPath, err)
	}

	cfg := &Config{}
	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
