package config

import "time"

type Config struct {
	Store *Store `json:"store"`
}

type Store struct {
	Driver      string        `json:"driver"`
	Redis       *Redis        `json:"redis"`
	PingTimeout time.Duration `json:"ping_timeout_second"`
}

type Redis struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBID     int    `json:"dbid"`
}

func NewDefaultConfig() *Config {
	return &Config{
		Store: &Store{
			Driver:      "redis",
			PingTimeout: 5 * time.Second,
			Redis: &Redis{
				Addr:     "10.151.144.75",
				Port:     "6379",
				Username: "",
				Password: "",
				DBID:     1,
			},
		},
	}
}
