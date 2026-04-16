package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address" env-required:"true" env-default:"localhost:8080"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	// checking the config path in os env vars
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		// we will check in runtime environment variable such as flags
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()
		configPath = *flags
		if configPath == "" {
			log.Fatal("Config Path is not set!")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config File Does not exist: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("Cannot Read the config File: %v", err)
	}

	return &cfg
}
