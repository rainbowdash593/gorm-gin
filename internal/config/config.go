package config

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"sync"
)

type Config struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"INFO"`
	Database struct {
		Host     string `env:"DATABASE_HOST" envDefault:"0.0.0.0"`
		Port     int    `env:"DATABASE_PORT" envDefault:"5432"`
		Name     string `env:"DATABASE_NAME"`
		User     string `env:"DATABASE_USER"`
		Password string `env:"DATABASE_PASSWORD"`
	}
	HTTP struct {
		Host string `env:"HTTP_HOST" envDefault:"127.0.0.1"`
		Port int    `env:"HTTP_PORT" envDefault:"8080"`
	}
}

var instance = Config{}
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		if err := env.Parse(&instance); err != nil {
			log.Fatal(err)
		}
	})

	return &instance
}
