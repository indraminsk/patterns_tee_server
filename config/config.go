package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log/slog"
)

type Config struct {
	App struct {
		HTTP struct {
			Host string `env:"APP_HTTP_HOST" env-required:"true"`
			Port string `env:"APP_HTTP_PORT" env-required:"true"`
		}

		Logger struct {
			Level int `env:"APP_LOGGER_LEVEL" env-default:"0"`
			Type  int `env:"APP_LOGGER_TYPE" env-default:"0"`
		}
	}
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		slog.Warn("[WARN] no .env file found")
	}

	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
