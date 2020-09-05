package apiserver

type Config struct {
	BindAddres string `toml:"bind_addres"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Config {

	return &Config{
		BindAddres: ":8000",
		LogLevel: "debug",
	}
}
