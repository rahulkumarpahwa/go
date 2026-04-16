package config

import (
	"flag"
	"log"
	"os"
)

type HTTPServer struct {
	Address string
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func (c *Config) MustLoad() {
	var configPath string

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
	

}
