package api

import "github.com/babaev2023/GO/GO_lvl2/lec8/storage"

//General instance for API server of REST application
type Config struct {
	//Port
	BindAddr string `toml:"bind_addr"`
	//Logger Level
	LoggerLevel string `toml:"logger_level"`
	//Store configs
	Storage *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8081",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
