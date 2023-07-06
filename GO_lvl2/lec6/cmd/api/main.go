package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/babaev2023/GO/GO_lvl2/lec6/internal/app/api"
)

var (
	configPath string = "configs/api.toml"
)

func init() {

}

func main() {
	log.Println("Works")
	//server instance initialization
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config) // Десериалзиуете содержимое .toml файла
	if err != nil {
		log.Println("can not find configs file. using default values:", err)
	}

	server := api.New(config)

	//api server start
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
