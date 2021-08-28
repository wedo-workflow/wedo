package config

type Config struct {
	StoreDriver string
	StoreDSN    string
}

func NewDefaultConfig() *Config {
	return &Config{
		StoreDriver: "mysql",
		StoreDSN:    "",
	}
}
