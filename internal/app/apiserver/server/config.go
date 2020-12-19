package server

type Config struct {
	BindAddres  string `yaml:"bind_addres"`
	LogLevel    string `yaml:"log_level"`
	DatabaseURL string `yaml:"database_url"`
}

func NewConfig() *Config {

	return &Config{
		BindAddres: ":8000",
		LogLevel:   "debug",
	}
}
