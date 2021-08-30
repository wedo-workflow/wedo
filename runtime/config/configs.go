package config

import "time"

type Config struct {
	Store *Store
}

type Store struct {
	Driver      string
	PingTimeout time.Duration
	Redis       *Redis
}

type Redis struct {
	DSN      string
	Username string
	Password string
	DB       int
}

func NewDefaultConfig() *Config {
	return &Config{
		Store: &Store{
			Driver:      "redis",
			PingTimeout: 5 * time.Second,
			Redis: &Redis{
				DSN:      "localhost:6379",
				Username: "",
				Password: "",
				DB:       0,
			},
		},
	}
}
