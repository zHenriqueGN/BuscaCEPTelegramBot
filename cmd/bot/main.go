package main

import (
	"fmt"
	"log"

	"github.com/zHenriqueGN/BuscaCEPTelegramBot/internal/config"
	"github.com/zHenriqueGN/BuscaCEPTelegramBot/internal/controller"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func init() {
	config.LoadEnv()
}

func main() {
	bot, err := tgbotapi.NewBotAPI(config.BotAPIToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 5

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			var msgTxt string
			switch update.Message.Command() {
			case "start":
				userName := update.Message.From.UserName
				if userName == "" {
					userName = "usuário"
				}
				msgTxt = fmt.Sprintf("Olá, %s!\nSeja bem-vindo(a) ao buscador de CEP!\nPara buscar um CEP, digite-o com o comando: /cep XXXXX-XXX", userName)
			case "cep":
				addr, err := controller.SearchCEP(update.Message.CommandArguments())
				if err != nil {
					msgTxt = err.Error()
				} else {
					msgTxt = fmt.Sprintf(
						"CEP: %s\nLogradouro: %s\nComplemento: %s\nBairro: %s\nLocalidade: %s\nUF: %s\nDDD: %s\nCódigo IBGE: %s\nCódigo GIA: %s\nCódigo SIAFI: %s",
						addr.CEP, addr.Logradouro, addr.Complemento, addr.Bairro, addr.Localidade, addr.UF, addr.DDD, addr.IBGE, addr.GIA, addr.SIAFI,
					)
				}
			default:
				msgTxt = "Digite um comando inválido!"
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgTxt)

			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}
