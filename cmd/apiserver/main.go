package main

import (
	"flag"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"

	"github.com/zlobste/mint-rest-api/internal/app/apiserver/server"
)

var (
	ConfigPath string
)

func init() {
	flag.StringVar(&ConfigPath, "config-path", "configs/apiserver.yaml", "path to config file")
}

func main() {
	flag.Parse()

	config := server.NewConfig()

	yamlConfig, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal([]byte(yamlConfig), &config)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Start(config); err != nil {
		log.Fatal(err)
	}
}
