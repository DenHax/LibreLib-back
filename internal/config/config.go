package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server HTTPserver
	PSQL   Postgres
}

type HTTPserver struct {
	Port int
}

type Postgres struct {
	Host     int
	Port     int
	User     string
	Password string
	Name     string
	SSLmode  string
}

type Logger struct {
	Dev   bool
	Level string
}

func MustConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable is not set")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("error opening config file: %s", err)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	return &cfg
}
