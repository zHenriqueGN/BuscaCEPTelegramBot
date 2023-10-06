package config

import (
	"os"
)

var BotAPIToken string

func LoadEnv() {
	BotAPIToken = os.Getenv("BOTAPITOKEN")
}
