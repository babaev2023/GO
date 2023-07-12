package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/babaev2023/GO/GO_lvl2/lec8/internal/app/api"
)

var (
	configPath string
	// string = "configs/api.toml"
)

func init() {
	//Скажем, что наше приложение будет на этапе запуска получать путь до конфиг файла из внешнего мира
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")

}

func main() {
	//В этот момент происходит инициализация переменной configPath значением
	flag.Parse()

	log.Println("Works")
	//server instance initialization
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config) // Десериалзиуете содержимое .toml файла
	if err != nil {
		log.Println("can not find configs file. using default values:", err)
	}

	server := api.New(config)

	//api server start
	log.Fatal(server.Start())

}
