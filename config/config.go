package config

import (
	"fmt"

	"github.com/joho/godotenv"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	FormParameter = "url"
	ShortAddr     = "http://sh.com/"
)

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		PG
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		//ShortAddr string `env-required:"true" yaml:"short_addr" env:"SHORT_ADDR"` //todo
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	PG struct {
		Host     string `env:"DB_HOST"`
		Port     string `env:"DB_PORT"`
		Username string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		DBName   string `env:"DB_NAME"`
		DBDriver string `env:"DB_DRIVER"`
		SSLMode  string `env:"SSL_MODE"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = godotenv.Load()

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
