package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/zlobste/mint-rest-api/internal/app/apiserver"
	"log"
)

var(
	ConfigPath string
)

func init() {
	flag.StringVar(&ConfigPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(ConfigPath, config)
	if err != nil {
		log.Fatal(err)
	}

	server := apiserver.New(config)
	if err := server.Start(); err != nil {
		log.Fatal(server)
	}
}

