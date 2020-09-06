package apiserver

import "github.com/zlobste/mint-rest-api/internal/app/store"

type Config struct {
	BindAddres string `toml:"bind_addres"`
	LogLevel string `toml:"log_level"`
	Store *store.Config
}

func NewConfig() *Config {

	return &Config{
		BindAddres: ":8000",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}
