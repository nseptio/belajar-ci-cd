package app

import (
	"log"

	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Server ConfigServer
	DB     ConfigDB
}

type ConfigServer struct {
	Port int `env:"SERVER_PORT,required"`
}

type ConfigDB struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	Username string `env:"DB_USERNAME,required"`
	Password string `env:"DB_PASSWORD,required"`
	DBName   string `env:"DB_NAME,required"`
}

func NewConfig() *Config {
	var c Config
	if err := env.Parse(&c); err != nil {
		log.Panicf("Failed to parse .env: %s", err)
	}
	return &c
}
