package config

import (
	"errors"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Postgres struct {
		Port    int    `yaml:"port"`
		Host    string `yaml:"host"`
		User    string `yaml:"user"`
		Pass    string `yaml:"pass"`
		Name    string `yaml:"name"`
		Sslmode string `yaml:"sslmode"`
	}
}

func GetConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("config.yml", cfg)
	if err != nil {
		help, _ := cleanenv.GetDescription(cfg, nil)
		return nil, errors.New(help)
	}
	return cfg, nil
}
