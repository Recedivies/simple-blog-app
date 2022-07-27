package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type AppConfig struct {
	App struct {
		Port string
	}
	DB struct {
		Driver   string
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
}

var Config AppConfig

func InitConfig() *AppConfig {
	if err := godotenv.Load(); err != nil {
		log.Error().Err(err).Msg("Error loading .env file")
	}

	Config.App.Port = os.Getenv("PORT")

	Config.DB.Driver = os.Getenv("DB_DRIVER")
	Config.DB.Host = os.Getenv("DB_HOST")
	Config.DB.Port = os.Getenv("DB_PORT")
	Config.DB.User = os.Getenv("DB_USER")
	Config.DB.Password = os.Getenv("DB_PASSWORD")
	Config.DB.Name = os.Getenv("DB_NAME")

	return &Config
}
