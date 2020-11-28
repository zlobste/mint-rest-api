package main

import (
	"flag"
	"log"
	
	"github.com/BurntSushi/toml"
	
	"github.com/zlobste/mint-rest-api/internal/app/apiserver/server"
)

var (
	ConfigPath string
)

func init() {
	flag.StringVar(&ConfigPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()
	
	config := server.NewConfig()
	
	_, err := toml.DecodeFile(ConfigPath, config)
	if err != nil {
		log.Fatal(err)
	}
	
	if err := server.Start(config); err != nil {
		log.Fatal(err)
	}
}
