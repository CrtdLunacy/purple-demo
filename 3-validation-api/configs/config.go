package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Email         EmailConfig
	EmailSettings EmailSettings
}

type EmailConfig struct {
	Email    string
	Password string
	Address  string
}

type EmailSettings struct {
	SMTP_HOST     string
	SMTP_PORT     string
	SMTP_EMAIL    string
	SMTP_PASSWORD string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("ERROR: Failed to load config, using default config")
	}

	return &Config{
		EmailSettings: EmailSettings{
			SMTP_HOST:     os.Getenv("SMTP_HOST"),
			SMTP_PORT:     os.Getenv("SMTP_PORT"),
			SMTP_EMAIL:    os.Getenv("SMTP_EMAIL"),
			SMTP_PASSWORD: os.Getenv("SMTP_PASSWORD"),
		},
	}
}
