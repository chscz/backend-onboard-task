package config

import (
	"encoding/json"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	MySQL          MySQL `envPrefix:"MYSQL_"`
	JWT            JWT   `envPrefix:"JWT_"`
	LocalDebugMode bool  `env:"LOCAL_DEBUG_MODE"`
}

func (c Config) String() string {
	b, _ := json.MarshalIndent(c, "", "  ")
	return string(b)
}

func LoadFromEnv() (Config, error) {
	cfg := Config{
		MySQL: MySQL{
			UserName: "",
			Password: "",
			Host:     "",
			Port:     "",
			DB:       "",
		},
		JWT: JWT{
			SecretKey:     "",
			ExpiredMinute: 1,
		},
		LocalDebugMode: false,
	}

	_ = godotenv.Load(".env")

	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
