package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MessageSender(bot *tgbotapi.BotAPI, ids []int64, message string) {

	for _, id := range ids {
		bot.Send(
			tgbotapi.NewMessage(
				id, message,
			),
		)
	}

	log.Println("Message sent successfully")
}
