package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server  ServerConfig `yaml:"server"`
	Storage Postgres     `yaml:"postgres"`
}

type ServerConfig struct {
	Port    int    // `yaml:"port"`
	SSLMode string // `yaml:"ssl_mode"`
	// AppVersion string // `yaml:"app_version"`
	// ReadTimeout  time.Duration // `yaml:"read_timeout"`
	// WriteTimeout time.Duration // `yaml:"write_timeout"`
}

type Postgres struct {
	URL      string // `yaml:"url"`
	Host     int    // `yaml:"host"`
	Port     int    // `yaml:"port"`
	User     string // `yaml:"user"`
	Password string // `yaml:"password"`
	Name     string // `yaml:"name"`
	SSLMode  string // `yaml:"ssl_mode"`
}

// type Logger struct {
// 	Env   string `yaml:"env"`
// 	Level string `yaml:"level"`
// }

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
