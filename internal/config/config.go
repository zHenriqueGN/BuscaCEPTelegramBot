package config

import (
	"os"

	"github.com/joho/godotenv"
)

var BotAPIToken string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	BotAPIToken = os.Getenv("BOTAPITOKEN")
}
