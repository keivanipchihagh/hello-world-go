package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Http struct {
	Host string
	Port int
}

func NewConfig() *Http {
	if err := godotenv.Load(); err != nil {
		log.Panic("error loading .env")
	}
	viper.AutomaticEnv()

	return &Http{
		Host: viper.GetString("HTTP_HOST"),
		Port: viper.GetInt("HTTP_PORT"),
	}
}
