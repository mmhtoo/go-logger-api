package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	PORT string `env:"PORT" default:":8080"`
	DB_HOST string `env:"DB_HOST" default:"localhost"`
	DB_PORT string `env:"DB_PORT" default:"5433"`
	DB_NAME string `env:"DB_NAME" default:"logger"`
	DB_USERNAME string `env:"DB_USERNAME" default:"root"`
	DB_PASSWORD string `env:"DB_PASSWORD" default:"root"`
	GIN_MODE string `env:"GIN_MODE" default:"debug"`
}

func LoadEnv() Environment {
	godotenv.Load()
	return Environment{
		PORT: os.Getenv("PORT"),
		DB_HOST: os.Getenv("DB_HOST"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_NAME: os.Getenv("DB_NAME"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),	
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),	
		GIN_MODE: os.Getenv("GIN_MODE"),
	}
}