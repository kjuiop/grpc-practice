package config

import (
	"github.com/naoina/toml"
	"log"
	"os"
)

type Config struct {
}

func NewConfig(path string) *Config {
	c := new(Config)

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed open config file, err : %v", err)
	}
	defer file.Close()

	if err := toml.NewDecoder(file).Decode(c); err != nil {
		log.Fatalf("failed decode config file, err : %v", err)
	}

	return c
}
