package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Address  string `yaml:"address"`
		TLS struct {
			Cert  string `yaml:"cert_file"`
			Key   string `yaml:"key_file"`
		} `yaml:"tls"`
	} `yaml:"server"`

	Auth struct {
		JWTPrivate string `yaml:"jwt_private_key_file"`
		JWTPublic  string `yaml:"jwt_public_key_file"`	
	} `yaml:"auth"`
}

func Load() *Config {
	data, err := os.ReadFile("config/config.yml")
	if err != nil {
		log.Fatalf("config read error: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("config parse error: %v", err)
	}

	return &cfg
}