package config

type Config struct {
	Store *Store
}

type Store struct {
	Driver   string
	DSN      string
	Username string
	Password string
}

func NewDefaultConfig() *Config {
	return &Config{
		Store: &Store{
			Driver:   "redis",
			DSN:      "localhost:6379",
			Username: "",
			Password: "",
		},
	}
}
